package mantis

import (
	"encoding/csv"
	"strings"
	"testing"
	"time"

	"github.com/jszwec/csvutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testDateStr = "12/11/2023"
	testHourStr = "09:46:12"
)

func TestDateMarshal(t *testing.T) {
	d := Date(time.Date(2023, time.November, 12, 0, 0, 0, 0, time.UTC))

	ser, err := d.MarshalText()
	require.NoError(t, err, "Unexpected error serializing date")
	assert.EqualValues(t, testDateStr, string(ser), "Mismatched date serialization result")
}

func TestDateUnmarshal(t *testing.T) {
	var (
		d        Date
		expected = Date(time.Date(2023, time.November, 12, 0, 0, 0, 0, time.UTC))
	)

	err := d.UnmarshalText([]byte(testDateStr))
	require.NoError(t, err, "Unexpected error parsing date")
	assert.EqualValues(t, expected, d, "Mismatched parsed date")

}

func TestHourMarshal(t *testing.T) {
	h := Hour(time.Date(0, time.January, 1, 9, 46, 12, 0, time.UTC))
	ser, err := h.MarshalText()
	require.NoError(t, err, "Unexpected error serializaing date")
	assert.EqualValues(t, testHourStr, string(ser), "Mismatched hour serialization result")
}

func TestHourUnmarshal(t *testing.T) {
	var (
		h        Hour
		expected = Hour(time.Date(0, time.January, 1, 9, 46, 12, 0, time.UTC))
	)

	err := h.UnmarshalText([]byte(testHourStr))
	require.NoError(t, err, "Unexpected error parsing hour")
	assert.EqualValues(t, expected, h, "Mismatched parsed hour")
}

const (
	testMantisCSVStr = `
Fecha;Hora;Prueba;Fase;Serie;Observaciones
11/11/2023; 09:18:00; H INF-A K1 500; Eliminatoria; Eliminatoria 1; Primera eliminación
`
)

func TestCSVUnmarshal(t *testing.T) {
	var h []Heat

	rd := csv.NewReader(strings.NewReader(testMantisCSVStr))
	rd.Comma = ';'
	r, err := csvutil.NewDecoder(rd)
	require.NoError(t, err, "Error preparing csv data")

	err = r.Decode(&h)
	require.NoError(t, err, "Error parsing csv data")

	require.Len(t, h, 1, "Unexpected slice size")
	dec := h[0]

	assert.EqualValues(t, time.Date(2023, time.November, 11, 0, 0, 0, 0, time.UTC), dec.Date, "Unexpected date")
	assert.EqualValues(t, time.Date(0, time.January, 1, 9, 18, 0, 0, time.UTC), dec.Hour, "Unexpected hour")
	assert.EqualValues(t, time.Date(2023, time.November, 11, 9, 18, 0, 0, time.UTC), dec.Time())

	assert.EqualValues(t, "H INF-A K1 500", dec.Event, "Unexpected event name")
	assert.EqualValues(t, "Eliminatoria", dec.Phase, "Unexpected event phase")
	assert.EqualValues(t, "Eliminatoria 1", dec.Series, "Unexpected event series")
	assert.EqualValues(t, "Primera eliminación", dec.Comments, "Unexpected event comments")

}
