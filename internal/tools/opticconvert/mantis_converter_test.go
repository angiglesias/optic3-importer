package opticconvert

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/angiglesias/optic3-importer/pkg/alge"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConversion(t *testing.T) {

	tests := []struct {
		settings Config
		name     string
		data     string
		result   alge.Meet
	}{
		{
			settings: Config{
				ExtendedHeatName: true,
				GroupDays:        false,
			},
			name: "plain conversion",
			data: `Fecha;Hora;Prueba;Fase;Serie;Observaciones
11/11/2023; 08:30:00; H CAD-A K1 500; Eliminatoria; Eliminatoria 1; 
11/11/2023; 08:34:00; H CAD-A K1 500; Eliminatoria; Eliminatoria 2; 
12/11/2023; 09:00:00; H CAD-A K1 500; Final; Final B; 
12/11/2023; 09:04:00; H CAD-A K1 500; Final; Final A; 
`,
			result: alge.Meet{
				Sessions: []alge.Session{
					{
						Date:   alge.Date(time.Date(2023, time.November, 11, 8, 30, 0, 0, time.UTC)),
						Name:   "Sesións",
						ID:     "Sesións",
						Number: 1,
						Events: []alge.Event{
							{
								Name:   "Carreiras",
								ID:     "Carreiras",
								Number: 1,
								Heats: []alge.Heat{
									{
										Name:   "H CAD-A K1 500 | Eliminatoria 1",
										ID:     "H CAD-A K1 500 | Eliminatoria 1",
										Number: 1,
										Start:  alge.StartTime(time.Date(2023, time.November, 11, 8, 30, 0, 0, time.UTC)),
									},
									{
										Name:   "H CAD-A K1 500 | Eliminatoria 2",
										ID:     "H CAD-A K1 500 | Eliminatoria 2",
										Number: 2,
										Start:  alge.StartTime(time.Date(2023, time.November, 11, 8, 34, 0, 0, time.UTC)),
									},
									{
										Name:   "H CAD-A K1 500 | Final B",
										ID:     "H CAD-A K1 500 | Final B",
										Number: 3,
										Start:  alge.StartTime(time.Date(2023, time.November, 12, 9, 0, 0, 0, time.UTC)),
									},
									{
										Name:   "H CAD-A K1 500 | Final A",
										ID:     "H CAD-A K1 500 | Final A",
										Number: 4,
										Start:  alge.StartTime(time.Date(2023, time.November, 12, 9, 4, 0, 0, time.UTC)),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			settings: Config{
				ExtendedHeatName: true,
				GroupDays:        true,
			},
			name: "grouped days",
			data: `Fecha;Hora;Prueba;Fase;Serie;Observaciones
11/11/2023; 08:30:00; H CAD-A K1 500; Eliminatoria; Eliminatoria 1; 
11/11/2023; 08:34:00; H CAD-A K1 500; Eliminatoria; Eliminatoria 2; 
12/11/2023; 09:00:00; H CAD-A K1 500; Final; Final B; 
12/11/2023; 09:04:00; H CAD-A K1 500; Final; Final A; 
`,
			result: alge.Meet{
				Sessions: []alge.Session{
					{
						Date:   alge.Date(time.Date(2023, time.November, 11, 8, 30, 0, 0, time.UTC)),
						Name:   "Día 1",
						ID:     "Día 1",
						Number: 1,
						Events: []alge.Event{
							{
								Name:   "Carreiras",
								ID:     "Carreiras",
								Number: 1,
								Heats: []alge.Heat{
									{
										Name:   "H CAD-A K1 500 | Eliminatoria 1",
										ID:     "H CAD-A K1 500 | Eliminatoria 1",
										Number: 1,
										Start:  alge.StartTime(time.Date(2023, time.November, 11, 8, 30, 0, 0, time.UTC)),
									},
									{
										Name:   "H CAD-A K1 500 | Eliminatoria 2",
										ID:     "H CAD-A K1 500 | Eliminatoria 2",
										Number: 2,
										Start:  alge.StartTime(time.Date(2023, time.November, 11, 8, 34, 0, 0, time.UTC)),
									},
								},
							},
						},
					},
					{
						Date:   alge.Date(time.Date(2023, time.November, 12, 9, 0, 0, 0, time.UTC)),
						Name:   "Día 2",
						ID:     "Día 2",
						Number: 2,
						Events: []alge.Event{
							{
								Name:   "Carreiras",
								ID:     "Carreiras",
								Number: 1,
								Heats: []alge.Heat{
									{
										Name:   "H CAD-A K1 500 | Final B",
										ID:     "H CAD-A K1 500 | Final B",
										Number: 3,
										Start:  alge.StartTime(time.Date(2023, time.November, 12, 9, 0, 0, 0, time.UTC)),
									},
									{
										Name:   "H CAD-A K1 500 | Final A",
										ID:     "H CAD-A K1 500 | Final A",
										Number: 4,
										Start:  alge.StartTime(time.Date(2023, time.November, 12, 9, 4, 0, 0, time.UTC)),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			settings: Config{},
			name:     "non extended names",
			data: `Fecha;Hora;Prueba;Fase;Serie;Observaciones
			11/11/2023; 08:30:00; H CAD-A K1 500; Eliminatoria; Eliminatoria 1; 
			11/11/2023; 08:34:00; H CAD-A K1 500; Eliminatoria; Eliminatoria 2; 
`,
			result: alge.Meet{
				Sessions: []alge.Session{
					{
						Date:   alge.Date(time.Date(2023, time.November, 11, 8, 30, 0, 0, time.UTC)),
						Name:   "Sesións",
						ID:     "Sesións",
						Number: 1,
						Events: []alge.Event{
							{
								Name:   "Carreiras",
								ID:     "Carreiras",
								Number: 1,
								Heats: []alge.Heat{
									{
										Name:   "H CAD-A K1 500",
										ID:     "H CAD-A K1 500",
										Number: 1,
										Start:  alge.StartTime(time.Date(2023, time.November, 11, 8, 30, 0, 0, time.UTC)),
									},
									{
										Name:   "H CAD-A K1 500",
										ID:     "H CAD-A K1 500",
										Number: 2,
										Start:  alge.StartTime(time.Date(2023, time.November, 11, 8, 34, 0, 0, time.UTC)),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cvt := mantisConverter{cfg: test.settings}
			data := strings.NewReader(test.data)
			res, err := cvt.Parse(data)
			require.NoError(t, err, "Unexpected error converting formats")
			assert.EqualValues(t, test.result, res, "Mismatched meeting after conversion")
			// t.Logf("%#v", res)
			xml, err := xml.Marshal(res)
			t.Log(string(xml))
		})
	}
}
