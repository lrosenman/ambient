package main

import (
	"flag"
	"github.com/lrosenman/ambient"
	"log"
	"time"
)

/*
This example queries all devices (weather stations) that are registered for the account the application and api keys are associated with.
To generate an application and api key for your own account, do so at https://ambientweather.net/account in the API Keys section.

API Docs:
https://ambientweather.docs.apiary.io/#reference/0/devices/list-user's-devices

Sample Usage:
go run main.go -applicationKey AFEA804E-9AB8-4E4F-BBCC-276C413E8B84 -apiKey F362D94E-FB4C-434F-A9B3-D4A2694CF6A4
*/

var (
	applicationKey     = flag.String("applicationKey", "", "Ambient Weather Application Key")
	apiKey             = flag.String("apiKey", "", "Ambient Weather API Key")
	maxNumberOfResults = flag.Int64("maxResults", 10, "Maximum number of results returned from the query.  Maximum allowed is 288")
)

func main() {
	flag.Parse()

	key := ambient.NewKey(*applicationKey, *apiKey)
	devices, err := ambient.Device(key)
	if err != nil {
		log.Panicln("unable to retrieve devices")
	}

	endDate := time.Now().UTC()
	for _, device := range devices.DeviceRecord {
		// Ensuring the rate limit is not exceeded per https://ambientweather.docs.apiary.io/#introduction/rate-limiting
		time.Sleep(1 * time.Second)

		log.Printf("Querying device '%s'", device.Macaddress)
		queryResults, queryErr := ambient.DeviceMac(key, device.Macaddress, endDate, *maxNumberOfResults)
		if queryErr != nil {
			log.Panicf("error when querying device '%s' %v", device.Macaddress, queryErr)
		}

		log.Printf("%v records found for device '%s' with response code %v", len(queryResults.Record), device.Macaddress, queryResults.HTTPResponseCode)
		for _, record := range queryResults.Record {
			log.Printf("Mac: '%s' Recorded At: %v Temperature: %v", device.Macaddress, record.Date, record.Tempf)
		}
	}
}
