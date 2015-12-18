// Package freegeoip provides a method for getting the physical location of an IP address using the freegeoip.net service
package freegeoip
import (
	"net/http"
	"fmt"
	"encoding/json"
	"errors"
	"strconv"
)

const freeGeoIpHostLookupUrl string = "http://freegeoip.net/json/%s"

type Location struct {
	Ip          string `json:"ip"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	RegionCode  string `json:"region_code"`
	RegionName  string `json:"region_name"`
	City        string `json:"city"`
	TimeZone    string `json:"time_zone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int `json:"metro_code"`
}

// Lookup returns the GeoIP data for the given hostname or IP address.
func Lookup(host string) (l Location, err error) {
	res, err := http.Get(fmt.Sprintf(freeGeoIpHostLookupUrl, host))
	if err != nil { return }
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = errors.New("Invalid status code, expected 200 got " + strconv.Itoa(res.StatusCode))
		return
	}
	err = json.NewDecoder(res.Body).Decode(&l)
	return
}