package optic3importer

import (
	"encoding/csv"
	"fmt"
	"io"
	"slices"
	"sort"
	"time"

	"github.com/angiglesias/optic3-importer/pkg/alge"
	"github.com/angiglesias/optic3-importer/pkg/mantis"
	"github.com/jszwec/csvutil"

	"golang.org/x/exp/maps"
)

type mantisConverter struct {
	cfg Config
}

func NewMantisConverter(cfg Config) Converter {
	return &mantisConverter{cfg: cfg}
}

func (m *mantisConverter) Parse(data io.Reader) (alge.Meet, error) {
	series, err := m.parseCsvStream(data)
	if err != nil {
		return alge.Meet{}, err
	}
	return m.Convert(series)
}

func (m *mantisConverter) parseCsvStream(data io.Reader) ([]mantis.Heat, error) {
	reader := csv.NewReader(data)
	reader.Comma = ';'
	dec, err := csvutil.NewDecoder(reader)
	if err != nil {
		return nil, err
	}
	var heats []mantis.Heat
	err = dec.Decode(&heats)
	if err != nil {
		return nil, err
	}
	return heats, nil
}

func (m *mantisConverter) Convert(series []mantis.Heat) (alge.Meet, error) {
	return m.simpleConversion(series)
}

func (m *mantisConverter) heatName(h mantis.Heat) string {
	if m.cfg.ExtendedHeatName {
		return h.Event.String() + " | " + h.Series.String()
	}
	return h.Event.String()
}

func (m *mantisConverter) simpleConversion(series []mantis.Heat) (alge.Meet, error) {
	// Pre-allocate series
	heats := make([]alge.Heat, len(series))
	meet := alge.Meet{}

	for idx, entry := range series {
		// Fill data
		heats[idx].Name = m.heatName(entry)
		heats[idx].ID = m.heatName(entry)
		heats[idx].Start = alge.StartTime(entry.DateFields.Time())
		heats[idx].Number = idx + 1
	}

	if m.cfg.GroupDays {
		days := make(map[time.Time]alge.Session)
		dayCtr := 1

		for i, h := range series {
			d, ok := days[h.Date.Time()]
			if !ok {
				d = alge.Session{
					Date:   alge.Date(h.DateFields.Time()),
					Number: dayCtr,
					Name:   fmt.Sprintf("Día %d", len(days)+1),
					ID:     fmt.Sprintf("Día %d", len(days)+1),
					Events: []alge.Event{
						{
							Name:   "Carreiras",
							ID:     "Carreiras",
							Number: 1,
						},
					},
				}
				dayCtr++
			}

			// Add heats to event
			d.Events[0].Heats = slices.Insert(d.Events[0].Heats, len(d.Events[0].Heats), heats[i])
			days[h.Date.Time()] = d
		}

		// fill meeting (sorted by day)
		keys := maps.Keys(days)
		sort.Slice(keys, func(i, j int) bool { return keys[i].Before(keys[j]) })
		for _, day := range keys {
			meet.Sessions = slices.Insert(meet.Sessions, len(meet.Sessions), days[day])
		}
	} else {
		// Create only session containing all events
		d := alge.Session{
			Date:   alge.Date(series[0].DateFields.Time()),
			Number: 1,
			Name:   "Sesións",
			ID:     "Sesións",
			Events: []alge.Event{
				{
					Name:   "Carreiras",
					ID:     "Carreiras",
					Number: 1,
					Heats:  heats,
				},
			},
		}

		// Add session to meet
		meet.Sessions = slices.Insert(meet.Sessions, len(meet.Sessions), d)
	}

	return meet, nil
}
