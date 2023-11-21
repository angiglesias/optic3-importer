package alge

type Meet struct {
	Name      string    `xml:",attr,omitempty"`
	ID        string    `xml:"Id,attr,omitempty"`
	City      string    `xml:",attr,omitempty"`
	Country   string    `xml:"Nation,attr,omitempty"`
	Organizer string    `xml:",attr,omitempty"`
	Promoter  string    `xml:",attr,omitempty"`
	Timing    string    `xml:",attr,omitempty"`
	Sessions  []Session `xml:"Session,omitempty"`
}

type Session struct {
	Name     string  `xml:",attr,omitempty"`
	ID       string  `xml:"Id,attr,omitempty"`
	Number   int     `xml:"Nr,attr,omitempty"`
	Location string  `xml:",attr,omitempty"`
	Type     string  `xml:",attr,omitempty"`
	Date     Date    `xml:",attr,omitempty"`
	Events   []Event `xml:"Event,omitempty"`
}

type Event struct {
	Name         string   `xml:",attr,omitempty"`
	ID           string   `xml:"Id,attr,omitempty"`
	Number       int      `xml:"Nr,attr,omitempty"`
	Distance     int      `xml:",attr,omitempty"`
	WindMeas     WindMeas `xml:"WindMeasurement,attr,omitempty"`
	UserDistance string   `xml:",attr,omitempty"`
	DistanceType Distance `xml:",attr,omitempty"`
	Heats        []Heat   `xml:"Heat,omitempty"`
}
