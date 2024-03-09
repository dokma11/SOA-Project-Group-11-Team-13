package model

import (
	"errors"
)

type TransportType int

const (
	Walking TransportType = iota
	Bicycle
	Car
)

type TourDuration struct {
	Duration      int           `json:"Duration"`
	TransportType TransportType `json:"TransportType"`
}

func NewTourDuration(duration int, transportType TransportType) (*TourDuration, error) {
	tourDuration := &TourDuration{
		Duration:      duration,
		TransportType: transportType,
	}

	if err := tourDuration.Validate(); err != nil {
		return nil, err
	}

	return tourDuration, nil
}

func (tourDuration *TourDuration) Validate() error {
	if tourDuration.Duration < 0 {
		return errors.New("invalid Duration")
	}

	/*
		if tourDuration.TransportType == "" {
			return errors.New("invalid Transport type")
		}
	*/
	return nil
}

/*
func (tourDuration TourDuration) Value() (driver.Value, error) {
	return json.Marshal(tourDuration)
}

func (tourDuration *TourDuration) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, tourDuration)
	case string:
		return json.Unmarshal([]byte(v), tourDuration)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
*/
/*
func (td TourDuration) Value() (driver.Value, error) {
	return json.Marshal(td)
}

func (td *TourDuration) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), td)
}
*/
