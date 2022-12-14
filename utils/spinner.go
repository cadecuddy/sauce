package utils

import (
	"github.com/briandowns/spinner"
	"time"
)

func ConfigSpinner() *spinner.Spinner {
	someSet := []string{"[π₯«πΈπΈπΈπΈ]", "[πΈπ₯«πΈπΈπΈ]", "[πΈπΈπ₯«πΈπΈ]", "[πΈπΈπΈπ₯«πΈ]", "[πΈπΈπΈπΈπ₯«]"}
	s := spinner.New(someSet, 100*time.Millisecond)
	s.Suffix = " π searching for sauce..."

	return s
}
