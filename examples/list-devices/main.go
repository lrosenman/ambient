package main

import (
	"flag"
	"github.com/lrosenman/ambient"
	"log"
)

/*
This example lists all devices (weather stations) that are registered for the account the application and api keys are associated with.
To generate an application and api key for your own account, do so at https://ambientweather.net/account in the API Keys section.

Sample Usage:
go run main.go -applicationKey AFEA804E-9AB8-4E4F-BBCC-276C413E8B84 -apiKey F362D94E-FB4C-434F-A9B3-D4A2694CF6A4
*/

var (
	applicationKey = flag.String("applicationKey", "", "Ambient Weather Application Key")
	apiKey         = flag.String("apiKey", "", "Ambient Weather API Key")
)

func main() {
	flag.Parse()

	key := ambient.NewKey(*applicationKey, *apiKey)
	devices, err := ambient.Device(key)
	if err != nil {
		log.Panicln("unable to retrieve devices")
	}

	log.Printf("%v devices found", len(devices.DeviceRecord))
	for _, item := range devices.DeviceRecord {
		log.Printf("Mac:%s Name:%s Location: %s", item.Macaddress, item.Info.Name, item.Info.Location)
	}
}
