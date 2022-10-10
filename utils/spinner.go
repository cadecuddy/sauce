package utils

import (
	"github.com/briandowns/spinner"
	"time"
)

func ConfigSpinner() *spinner.Spinner {
	someSet := []string{"[🥫🌸🌸🌸🌸]", "[🌸🥫🌸🌸🌸]", "[🌸🌸🥫🌸🌸]", "[🌸🌸🌸🥫🌸]", "[🌸🌸🌸🌸🥫]"}
	s := spinner.New(someSet, 100*time.Millisecond)
	s.Suffix = " 🔍 searching for sauce..."

	return s
}
