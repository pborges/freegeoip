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
