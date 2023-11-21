package alge

import (
	"encoding/xml"
	"fmt"
	"time"
)

func ExampleMeeting() {
	meeting := Meet{
		ID:        "CESA2023",
		Name:      "CESA 2023",
		City:      "Verducido",
		Country:   "España",
		Organizer: "FEGAPI",
		Sessions: []Session{
			{
				ID:     "Pruebas",
				Name:   "Probas",
				Number: 1,
				Date:   Date(time.Date(2023, time.November, 11, 0, 0, 0, 0, time.Local)),
			},
			{
				ID:     "Dia 1",
				Name:   "Dia Un",
				Number: 2,
				Date:   Date(time.Date(2023, time.November, 11, 0, 0, 0, 0, time.Local)),
				Events: []Event{
					{
						Name:   "Carreiras",
						ID:     "Carreras",
						Number: 1,
						Heats: []Heat{
							{
								Number: 1,
								ID:     "H CAD-A K1 500",
								Name:   "HCAD K1 500",
								Start:  StartTime(time.Date(2023, time.November, 11, 8, 30, 0, 0, time.Local)),
							},
						},
					},
				},
			},
		},
	}

	data, _ := xml.Marshal(meeting)
	fmt.Println(string(data))

	// Output: <Meet Name="CESA 2023" Id="CESA2023" City="Verducido" Nation="España" Organizer="FEGAPI"><Session Name="Probas" Id="Pruebas" Nr="1" Date="2023-11-11"></Session><Session Name="Dia Un" Id="Dia 1" Nr="2" Date="2023-11-11"><Event Name="Carreiras" Id="Carreras" Nr="1"><Heat Name="HCAD K1 500" ScheduledStarttime="08:30" Id="H CAD-A K1 500" Nr="1"></Heat></Event></Session></Meet>
}
