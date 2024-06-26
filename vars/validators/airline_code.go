package validators

import "regexp"

var (
	AirlineCodeMatcher          = regexp.MustCompile(`^[A-Z0-9]{2}[A-Z0-9]?$`)
	UnboundedAirlineCodeMatcher = regexp.MustCompile(`[A-Z0-9]{2}[A-Z0-9]?`)
	IataMatcher                 = regexp.MustCompile(`^[A-Z]{3}$`)
	UnboundedIataMatcher        = regexp.MustCompile(`[A-Z]{3}`)
)
