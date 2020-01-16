package validators

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/vladimirdotk/weather/types"
)

const (
	// todo: Use some kind of validation library
	latRegExp  string = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	longRegExp string = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
)

// ValidateCoords checks latitude and longitude
func ValidateCoords(weather *types.Weather) url.Values {
	errors := url.Values{}

	var validLat = regexp.MustCompile(latRegExp)
	var validLong = regexp.MustCompile(longRegExp)

	var errorsMsgPart = "You should provide a valid %s"

	if !validLat.MatchString(weather.Lat) {
		errors.Add("lat", fmt.Sprintf(errorsMsgPart, "Latitude"))
	}

	if !validLong.MatchString(weather.Long) {
		errors.Add("long", fmt.Sprintf(errorsMsgPart, "Longitude"))
	}

	return errors
}
