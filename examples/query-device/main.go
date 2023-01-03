package main

import (
	"flag"
	"github.com/lrosenman/ambient/pkg/ambient"
	"log"
	"time"
)

/*
This example queries a specific device for its historical data.

API Docs:
https://ambientweather.docs.apiary.io/#reference/0/device-data/query-device-data

Sample Usage:
go run main.go -applicationKey AFEA804E-9AB8-4E4F-BBCC-276C413E8B84 -apiKey F362D94E-FB4C-434F-A9B3-D4A2694CF6A4 -macAddress 00:0E:C6:10:01:86
*/

var (
	applicationKey     = flag.String("applicationKey", "", "Ambient Weather Application Key")
	apiKey             = flag.String("apiKey", "", "Ambient Weather API Key")
	macAddress         = flag.String("macAddress", "", "Mac Address for the device to query")
	maxNumberOfResults = flag.Int64("maxResults", 10, "Maximum number of results returned from the query.  Maximum allowed is 288")
)

func main() {
	flag.Parse()

	key := ambient.NewKey(*applicationKey, *apiKey)

	endDate := time.Now().UTC()
	queryResults, err := ambient.DeviceMac(key, *macAddress, endDate, *maxNumberOfResults)
	if err != nil {
		log.Panicln("unable to query device")
	}

	for _, record := range queryResults.Record {
		log.Printf("Recorded At: %v Temperature: %v", record.Date, record.Tempf)
	}
}
