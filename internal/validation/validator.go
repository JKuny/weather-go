package validation

import (
	"errors"
	"strconv"
)

// ValidateLatLong Checks if the passed in string are validate latitude/longitude values.
func ValidateLatLong(lat, long string) (bool, error) {

	// Check values and return an error if not floats
	convLat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return false, err
	}

	convLong, err := strconv.ParseFloat(long, 64)
	if err != nil {
		return false, err
	}

	// Check if the convLat and convLong values are valid for a lat/long coordinate
	if convLat > -90 && convLat < 90 && convLong > -180 && convLong < 180 {
		return true, nil
	}

	err = errors.New("latitude and longitude must be between -90 and 90 and -180 and 180")
	return false, err
}
