package mantis

import (
	"bytes"
	"time"
)

type Heat struct {
	DateFields

	Event    String `csv:"Prueba"`
	Phase    String `csv:"Fase"`
	Series   String `csv:"Serie"`
	Comments String `csv:"Observaciones"`
}

func (h Heat) Headers() []string {
	return []string{"Fecha", "Hora", "Prueba", "Fase", "Serie", "Observaciones"}
}

func (h Heat) KnownHeaders() [][]string {
	return [][]string{
		{"Fecha", "Hora", "Prueba", "Fase", "Serie", "Observaciones"},
		{"Data", "Hora", "Proba", "Fase", "Serie", "Observacións"},
	}
}

type String string

func (s *String) UnmarshalCSV(data []byte) error {
	*s = String(bytes.TrimSpace(data))
	return nil
}

func (s String) MarshalCSV() ([]byte, error) {
	return append([]byte(" "), []byte(s.String())...), nil
}

func (s String) String() string {
	return string(s)
}

type DateFields struct {
	Date Date `csv:"Fecha"`
	Hour Hour `csv:"Hora"`
}

func (df DateFields) Time() time.Time {
	d := time.Time(df.Date)
	h := time.Time(df.Hour)
	return time.Date(
		d.Year(), d.Month(), d.Day(),
		h.Hour(), h.Minute(), h.Second(), h.Nanosecond(), h.Location(),
	)
}

func (df *DateFields) SetTime(t time.Time) {
	df.Date = Date(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()))
	df.Hour = Hour(time.Date(0, time.January, 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location()))
}

type Date time.Time

const dateFmt = "02/01/2006"

func (d Date) MarshalCSV() ([]byte, error) {
	return d.MarshalText()
}

func (d *Date) UnmarshalCSV(data []byte) error {
	return d.UnmarshalText(bytes.TrimSpace(data))
}

func (d Date) MarshalText() ([]byte, error) {
	return []byte(time.Time(d).Format(dateFmt)), nil
}

func (d *Date) UnmarshalText(text []byte) error {
	t, err := time.Parse(dateFmt, string(text))
	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}

func (d Date) Time() time.Time {
	return time.Time(d)
}

type Hour time.Time

const hourFmt = "15:04:05"

func (h Hour) MarshalCSV() ([]byte, error) {
	return h.MarshalText()
}

func (h *Hour) UnmarshalCSV(data []byte) error {
	return h.UnmarshalText(bytes.TrimSpace(data))
}

func (h Hour) MarshalText() ([]byte, error) {
	return []byte(time.Time(h).Format(hourFmt)), nil
}

func (h *Hour) UnmarshalText(text []byte) error {
	t, err := time.Parse(hourFmt, string(text))
	if err != nil {
		return err
	}

	*h = Hour(t)
	return nil
}

func (d Hour) Time() time.Time {
	return time.Time(d)
}
