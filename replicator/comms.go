package replicator

import "time"

type CommsObservation struct {
	State         string `json:"state"`
	Outcome       string `json:"outcome"`
	Endpoint      string `json:"endpoint,omitempty"`
	Error         string `json:"error,omitempty"`
	ObservedAt    string `json:"observed_at,omitempty"`
	ActivityAt    string `json:"activity_at,omitempty"`
	LastSuccessAt string `json:"last_success_at,omitempty"`
	ExceptionCode *uint8 `json:"exception_code,omitempty"`
}

type CycleComms struct {
	Network CommsObservation `json:"network"`
	TCP     CommsObservation `json:"tcp"`
	Modbus  CommsObservation `json:"modbus"`
	MMA2    CommsObservation `json:"mma2"`
}

func observation(state, outcome, endpoint string) CommsObservation {
	return CommsObservation{State: state, Outcome: outcome, Endpoint: endpoint, ObservedAt: time.Now().UTC().Format(time.RFC3339Nano)}
}

func aggregateComms(observations []CommsObservation) string {
	result := "OK"
	if len(observations) == 0 {
		return "UNKNOWN"
	}
	for _, observed := range observations {
		if observed.State == "ERROR" {
			return "ERROR"
		}
		if observed.State == "WARNING" {
			result = "WARNING"
		} else if observed.State != "OK" && result != "WARNING" {
			result = "UNKNOWN"
		}
	}
	return result
}
