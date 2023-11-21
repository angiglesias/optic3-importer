package alge

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subchen/go-xmldom"
)

const (
	heatTestStr string = `<Heat Name="H CAD K1 500" Id="H CAD K1 500" Nr="1" ScheduledStarttime="10:35"></Heat>`
)

func TestHeatUmarshal(t *testing.T) {
	var heat Heat
	cmp := Heat{
		Name:   "H CAD K1 500",
		ID:     "H CAD K1 500",
		Number: 1,
		Start:  StartTime(time.Date(0, time.January, 1, 10, 35, 0, 0, time.UTC)),
	}

	err := xml.Unmarshal([]byte(heatTestStr), &heat)
	require.NoError(t, err, "Error decoding heat xml")
	assert.EqualValues(t, cmp, heat, "Mismatched reference heat and decoded")
}

func TestHeatMarshal(t *testing.T) {
	heat := Heat{
		Name:   "H CAD K1 500",
		ID:     "H CAD K1 500",
		Number: 1,
		Start:  StartTime(time.Date(0, time.January, 1, 10, 35, 0, 0, time.UTC)),
	}

	enc, err := xml.Marshal(heat)
	require.NoError(t, err, "Error encoding heat series")

	orig, err := xmldom.ParseXML(heatTestStr)
	require.NoError(t, err, "Error generating map from test xml")
	test, err := xmldom.ParseXML(string(enc))
	require.NoError(t, err, "Error generating map from encoded xml")

	assert.EqualValues(t, test.Root.GetAttribute("Name"), orig.Root.GetAttribute("Name"), "Mismatched heat series xml values")
	assert.EqualValues(t, test.Root.GetAttribute("Id"), orig.Root.GetAttribute("Id"), "Mismatched heat series xml values")
	assert.EqualValues(t, test.Root.GetAttribute("Nr"), orig.Root.GetAttribute("Nr"), "Mismatched heat series xml values")
	assert.EqualValues(t, test.Root.GetAttribute("ScheduledStarttime"), orig.Root.GetAttribute("ScheduledStarttime"), "Mismatched heat series xml values")
}
