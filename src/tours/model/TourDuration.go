package model

import (
	"errors"
)

type TransportType string

const (
	Walking TransportType = "Walking"
	Bicycle TransportType = "Bicycle"
	Car     TransportType = "Car"
)

type TourDuration struct {
	Duration      int           `json:"duration"`
	TransportType TransportType `json:"transportType"`
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
	if tourDuration.TransportType == "" {
		return errors.New("invalid Transport type")
	}

	return nil
}
