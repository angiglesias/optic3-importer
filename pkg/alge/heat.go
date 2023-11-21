package alge

import "time"

type Heat struct {
	Name         string       `xml:",attr,omitempty"`
	Start        StartTime    `xml:"ScheduledStarttime,attr,omitempty"`
	ID           string       `xml:"Id,attr,omitempty"`
	Number       int          `xml:"Nr,attr,omitempty"`
	Distance     int          `xml:",attr,omitempty"`
	WindMeas     WindMeas     `xml:",attr,omitempty"`
	UserDistance string       `xml:",attr,omitempty"`
	DistanceType Distance     `xml:",attr,omitempty"`
	Competitors  []Competitor `xml:",omitempty"`
}

type Competitor struct {
	Bib         int       `xml:",attr,omitempty"`
	Lane        int       `xml:",attr,omitempty"`
	Start       time.Time `xml:"Starttime,attr,omitempty"`
	Firsname    string    `xml:",attr,omitempty"`
	Lastname    string    `xml:",attr,omitempty"`
	Country     string    `xml:"Nation,attr,omitempty"`
	Code        string    `xml:",attr,omitempty"`
	Gender      string    `xml:",attr,omitempty"`
	City        string    `xml:",attr,omitempty"`
	Team        string    `xml:",attr,omitempty"`
	Class       string    `xml:",attr,omitempty"`
	Horse       string    `xml:",attr,omitempty"`
	HorseDriver string    `xml:",attr,omitempty"`
	HorserOwner string    `xml:",attr,omitempty"`
}
