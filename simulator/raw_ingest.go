package simulator

import (
	"fmt"
	"net"
	"strconv"

	"github.com/tamzrod/MCS.OSJS/mma2raw"
)

const (
	rawRespOK    = mma2raw.ResponseOK
	rawHeaderLen = mma2raw.HeaderLen
	rawVersion   = mma2raw.Version
)

type RawIngestClient struct {
	client *mma2raw.Client
	addr   string
}

func NewRawIngestClient(def DeviceDefinition) *RawIngestClient {
	addr := net.JoinHostPort("127.0.0.1", strconv.Itoa(int(def.MMA2.Port)))
	return &RawIngestClient{
		addr: addr,
		client: mma2raw.NewClient(addr, def.MMA2.UnitID, map[mma2raw.Area]mma2raw.Range{
			mma2raw.Coils:            {Start: def.MMA2.FC1.Start, Count: def.MMA2.FC1.Count},
			mma2raw.DiscreteInputs:   {Start: def.MMA2.FC2.Start, Count: def.MMA2.FC2.Count},
			mma2raw.HoldingRegisters: {Start: def.MMA2.FC3.Start, Count: def.MMA2.FC3.Count},
			mma2raw.InputRegisters:   {Start: def.MMA2.FC4.Start, Count: def.MMA2.FC4.Count},
		}),
	}
}

func (c *RawIngestClient) Send(v Values) error {
	area, err := sharedRawArea(v.FC)
	if err != nil {
		return err
	}
	c.client.Addr = c.addr
	return c.client.Send(area, mma2raw.Values{Bits: v.Coils, Registers: v.Regs})
}

func sharedRawArea(fc FCKind) (mma2raw.Area, error) {
	switch fc {
	case FC1:
		return mma2raw.Coils, nil
	case FC2:
		return mma2raw.DiscreteInputs, nil
	case FC3:
		return mma2raw.HoldingRegisters, nil
	case FC4:
		return mma2raw.InputRegisters, nil
	default:
		return 0, fmt.Errorf("unsupported FC kind in values: %d", fc)
	}
}
