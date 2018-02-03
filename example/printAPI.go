// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright 2018 Larry Rosenman, LERCTR Consulting, larryrtx@gmail.com
//

// Example Program for the ambient package.
// printAPI shows all API calls and the responses to them
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lrosenman/ambient"
	"time"
)

// This is a DEMO KEY, replace it with your own.
const applicationKey = "21a439e927a84a25bb79ffe894fdd372b3e9d2e8bcef4167943b52cbe4530d9f"

// This is a DEMO KEY, replace it with your own.
const apiKey = "78f9704baaab411a87edeed59052cbb687a4aa7764a44accbaf6447492b0ca8c"

func main() {
	key := ambient.NewKey(applicationKey, apiKey)
	dr, err := ambient.Device(key)
	if err != nil {
		panic(err)
	}
	switch dr.HTTPResponseCode {
	case 200:
	case 429, 503:
		{
			fmt.Printf("Error code %d, retrying.\n", dr.HTTPResponseCode)
			time.Sleep(1 * time.Second)
			dr, err = ambient.Device(key)
			if err != nil {
				panic(err)
			}
			switch dr.HTTPResponseCode {
			case 200:
			default:
				{
					panic(dr)
				}
			}
		}
	default:
		{
			panic(dr)
		}
	}
	// API RateLimit
	time.Sleep(1 * time.Second)
	ar, err := ambient.DeviceMac(key, dr.DeviceRecord[0].Macaddress, time.Now(), 1)
	if err != nil {
		panic(err)
	}
	switch ar.HTTPResponseCode {
	case 200:
	case 429, 503:
		{
			fmt.Printf("Error code %d, retrying.\n", ar.HTTPResponseCode)
			time.Sleep(1 * time.Second)
			ar, err = ambient.DeviceMac(key, dr.DeviceRecord[0].Macaddress, time.Now(), 1)
			if err != nil {
				panic(err)
			}
			switch ar.HTTPResponseCode {
			case 200:
			default:
				{
					panic(ar)
				}
			}
		}
	default:
		{
			panic(ar)
		}
	}
	var arPrettyJSON bytes.Buffer
	var drPrettyJSON bytes.Buffer
	json.Indent(&drPrettyJSON, dr.JSONResponse, "", "\t")
	json.Indent(&arPrettyJSON, ar.JSONResponse, "", "\t")
	arRecordJSON, _ := json.MarshalIndent(ar.Record, "", "\t")
	drDeviceRecordJSON, _ := json.MarshalIndent(dr.DeviceRecord, "", "\t")
	fmt.Printf("DeviceResponse:\nHTTPResponseCode: %d, ResponseTime: %v\n", dr.HTTPResponseCode, dr.ResponseTime)
	fmt.Printf("Device Record:\n%+v\n", string(drDeviceRecordJSON))
	fmt.Printf("JSONResponse:\n%s\n\n", string(drPrettyJSON.Bytes()))
	fmt.Printf("DeviceMacResponse:\nHTTPResponseCode: %d, ResponseTime: %v\n", ar.HTTPResponseCode, ar.ResponseTime)
	fmt.Printf("Record:\n%+v\n", string(arRecordJSON))
	fmt.Printf("JSONResponse:\n%s\n\n", string(arPrettyJSON.Bytes()))
}
