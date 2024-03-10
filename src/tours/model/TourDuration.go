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
		return errors.New("invalid Duration. Duration cannot be empty")
	}
	if tourDuration.TransportType < 0 || tourDuration.TransportType > 2 {
		return errors.New("invalid Transport Type. Transport Type's value must be in range of 0 to 2")
	}
	return nil
}
