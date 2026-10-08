package simulator

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"

)

// PollPersistenceRestore is the persistence manager watchdog.
//
// It intentionally has one simple rule for each persistence-enabled memory:
// a normal Modbus probe returning the configured sealed exception (normally
// 0x06 Server Device Busy) means MMA2 is sealed, so restore the latest snapshot
// through Raw Ingest and let the existing restore contract perform the final
// lock-coil=1 commit. Any normal response means no work.
//
// This makes MMA2 restart recovery independent of Electron and independent of
// whether the Simulator Windows service itself restarted.
func (a *SchedulerApplier) PollPersistenceRestore() error {
	a.mu.Lock()
	devices := make([]DeviceDefinition, 0, len(a.devices))
	for _, device := range a.devices {
		devices = append(devices, device)
	}
	a.mu.Unlock()

	for _, device := range devices {
		if !device.Enabled || !persistenceConfigured(device) {
			continue
		}
		sealed, err := persistenceProbeSealed(device)
		if err != nil {
			// Endpoint-down and unrelated Modbus errors are not sealing events.
			// The next watchdog tick will probe again.
			continue
		}
		if !sealed {
			continue
		}

		results, err := persistenceStartupContext(a.store.Root, []DeviceDefinition{device})
		if err != nil {
			return fmt.Errorf("device %q persistence restore: %w", device.Name, err)
		}
		a.recordPersistenceStartup(results)
		for _, result := range results {
			if !result.Result.Committed {
				return fmt.Errorf(
					"device %q persistence restore did not commit (%s); memory remains sealed",
					device.Name, result.Result.Failure,
				)
			}
		}
	}
	return nil
}

// persistenceProbeSealed performs one ordinary Modbus read against the first
// configured area. State Sealing rejects all ordinary Modbus requests with the
// configured exception, so one address is enough to detect the sealed state.
func persistenceProbeSealed(device DeviceDefinition) (bool, error) {
	var function byte
	var start uint16
	switch {
	case device.MMA2.FC1.Count > 0:
		function, start = 1, device.MMA2.FC1.Start
	case device.MMA2.FC2.Count > 0:
		function, start = 2, device.MMA2.FC2.Start
	case device.MMA2.FC3.Count > 0:
		function, start = 3, device.MMA2.FC3.Start
	case device.MMA2.FC4.Count > 0:
		function, start = 4, device.MMA2.FC4.Start
	default:
		return false, fmt.Errorf("device %q has no Modbus area to probe", device.Name)
	}

	conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", fmt.Sprint(device.MMA2.Port)), 500*time.Millisecond)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(500 * time.Millisecond))

	req := make([]byte, 12)
	binary.BigEndian.PutUint16(req[0:2], 1)
	binary.BigEndian.PutUint16(req[4:6], 6)
	req[6] = byte(device.MMA2.UnitID)
	req[7] = function
	binary.BigEndian.PutUint16(req[8:10], start)
	binary.BigEndian.PutUint16(req[10:12], 1)
	if _, err := conn.Write(req); err != nil {
		return false, err
	}

	header := make([]byte, 7)
	if _, err := io.ReadFull(conn, header); err != nil {
		return false, err
	}
	pduLen := int(binary.BigEndian.Uint16(header[4:6])) - 1
	if pduLen < 1 {
		return false, fmt.Errorf("device %q returned an empty Modbus PDU", device.Name)
	}
	pdu := make([]byte, pduLen)
	if _, err := io.ReadFull(conn, pdu); err != nil {
		return false, err
	}

	if pdu[0]&0x80 == 0 {
		return false, nil
	}
	if len(pdu) < 2 {
		return false, fmt.Errorf("device %q returned a short Modbus exception", device.Name)
	}

	exception := byte(0x06)
	if block, ok := device.MMA2.StateSealing["exception"]; ok {
		switch v := block.(type) {
		case int:
			exception = byte(v)
		case uint8:
			exception = byte(v)
		case uint16:
			exception = byte(v)
		}
	}
	return pdu[1] == exception, nil
}
