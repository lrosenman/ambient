// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright 2018 Larry Rosenman, LERCTR Consulting, larryrtx@gmail.com
//

// Example Program for the ambient package.
package main

import (
	"fmt"
	"github.com/lrosenman/ambient"
	"time"
)

const applicationKey = "<your application key here>"
const apiKey = "<your API key here>"

func main() {
	// create a Key object
	key := ambient.NewKey(applicationKey, apiKey)
	// Get a list of the Devices for this key pair
	dr, err := ambient.Device(key)
	if err != nil {
		panic(err)
	}
	// walk the list of Mac Addresses, and print the latest temperature
	for i, macRec := range dr.DeviceRecord {
		fmt.Printf("MacAddress[%d]=%s\n", i, macRec.Macaddress)
		// API Rate Limit (1/second)
		time.Sleep(1 * time.Second)
		// Get the latest ambient.AmbientRecord and print Date and Tempf
		ar, err := ambient.DeviceMac(key, macRec.Macaddress, time.Now(), 1)
		if err != nil {
			panic(err)
		}
		if ar.HTTPResponseCode == 200 {
			fmt.Printf("Date=%v, Tempf=%f\n", ar.AmbientRecord[0].Date, ar.AmbientRecord[0].Tempf)
		} else {
			fmt.Printf("Bad HTTPResponseCode=%d\n", ar.HTTPResponseCode)
		}
	}
}