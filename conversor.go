package optic3importer

import (
	"io"

	"github.com/angiglesias/optic3-importer/pkg/alge"
	"github.com/angiglesias/optic3-importer/pkg/mantis"
)

type Converter interface {
	// Parse  will read raw data from a io.Reader and generate an alge XML meeting
	Parse(data io.Reader) ([]alge.Meet, error)
	// Conver will process a series of mantis Heats and will generate an alge XML meeting
	Convert(series []mantis.Heat) ([]alge.Meet, error)
}

type GroupDayMode int

const (
	None       GroupDayMode = iota // All series in a single file
	SingleFile                     // Group heats by day under single file
	MultiFile                      // Group heats by day, with each day on a separate file
)

type Config struct {
	IncludeIndexInHeatName bool
	ExtendedHeatName       bool
	GroupDays              GroupDayMode
}
