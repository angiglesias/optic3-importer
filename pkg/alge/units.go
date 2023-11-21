package alge

import (
	"errors"
	"fmt"
	"time"
)

type WindMeas int

const (
	None WindMeas = iota
	TenSecWithDelay
	ThirteenSec
	FiveSec
	TenSec
)

func (w WindMeas) MarshalText() ([]byte, error) {
	switch w {
	case TenSecWithDelay:
		return []byte("10SecondsWith10SecondsDelay"), nil
	case ThirteenSec:
		return []byte("13Seconds"), nil
	case FiveSec:
		return []byte("5Seconds"), nil
	case TenSec:
		return []byte("10Seconds"), nil
	case None:
		return []byte("None"), nil
	}
	return nil, errors.New("unknown wind measurement type")
}

func (w *WindMeas) UnmarshalText(text []byte) error {
	switch string(text) {
	case "10SecondsWith10SecondsDelay":
		*w = TenSecWithDelay
	case "13Seconds":
		*w = ThirteenSec
	case "10Seconds":
		*w = TenSec
	case "5Seconds":
		*w = FiveSec
	case "None":
		*w = None
	default:
		return fmt.Errorf("unkown wind measurement type %s", string(text))
	}
	return nil
}

type Distance int

const (
	Regular Distance = iota
	Steeplechase
	Hurdles
	Relay
)

func (d Distance) MarshalText() ([]byte, error) {
	switch d {
	case Steeplechase:
		return []byte("Steeplechase"), nil
	case Regular:
		return []byte("Regular"), nil
	case Hurdles:
		return []byte("Hurdles"), nil
	case Relay:
		return []byte("Relay"), nil
	}
	return nil, errors.New("unknown distance type")
}

func (d *Distance) UnmarshalText(text []byte) error {
	switch string(text) {
	case "Steeplechase":
		*d = Steeplechase
	case "Regular":
		*d = Regular
	case "Hurdles":
		*d = Hurdles
	case "Relay":
		*d = Relay
	default:
		return fmt.Errorf("unkown distance type %s", string(text))
	}
	return nil
}

type Date time.Time

func (d Date) MarshalText() ([]byte, error) {
	return []byte(time.Time(d).Format(time.DateOnly)), nil
}

func (d *Date) UnmarshalText(text []byte) error {
	t, err := time.Parse(time.DateOnly, string(text))
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

type StartTime time.Time

func (s StartTime) MarshalText() ([]byte, error) {
	return []byte(time.Time(s).Format("15:04")), nil
}

func (s *StartTime) UnmarshalText(text []byte) error {
	t, err := time.Parse("15:04", string(text))
	if err != nil {
		return nil
	}
	*s = StartTime(t)
	return nil
}
