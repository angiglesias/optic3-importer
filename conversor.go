package optic3importer

import (
	"io"

	"github.com/angiglesias/optic3-importer/pkg/alge"
	"github.com/angiglesias/optic3-importer/pkg/mantis"
)

type Converter interface {
	// Parse  will read raw data from a io.Reader and generate an alge XML meeting
	Parse(data io.Reader) (alge.Meet, error)
	// Conver will process a series of mantis Heats and will generate an alge XML meeting
	Convert(series []mantis.Heat) (alge.Meet, error)
}

type Config struct {
	ExtendedHeatName bool
	GroupDays        bool
}
