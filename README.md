# FreeGeoIP.net API
It's a simple API, but someone should do it :)
## Install:
```
go get github.com/pborges/freegeoip
```
##Example:
```
package main
import (
	"fmt"
	"github.com/pborges/freegeoip"
	"log"
)

func main() {
	l, err := freegeoip.Lookup("github.com")
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("%- 11s: %s\n","Ip",l.Ip)
	fmt.Printf("%- 11s: %s\n","CountryCode",l.CountryCode)
	fmt.Printf("%- 11s: %s\n","CountryName",l.CountryName)
	fmt.Printf("%- 11s: %s\n","RegionCode",l.RegionCode)
	fmt.Printf("%- 11s: %s\n","RegionName",l.RegionName)
	fmt.Printf("%- 11s: %s\n","City",l.City)
	fmt.Printf("%- 11s: %s\n","TimeZone",l.TimeZone)
	fmt.Printf("%- 11s: %f\n","Latitude",l.Latitude)
	fmt.Printf("%- 11s: %f\n","Longitude",l.Longitude)
	fmt.Printf("%- 11s: %d\n","MetroCode",l.MetroCode)
}

```
##Output:
```
Ip         : 192.30.252.129
CountryCode: US
CountryName: United States
RegionCode : CA
RegionName : California
City       : San Francisco
TimeZone   : America/Los_Angeles
Latitude   : 37.769700
Longitude  : -122.393300
MetroCode  : 807
```