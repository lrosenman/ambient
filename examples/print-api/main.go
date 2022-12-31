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
	"os"
	"time"
)

// This is a DEMO KEY, replace it with your own.
//
//goland:noinspection SpellCheckingInspection
const applicationKey = "21a439e927a84a25bb79ffe894fdd372b3e9d2e8bcef4167943b52cbe4530d9f"

// This is a DEMO KEY, replace it with your own.
//
//goland:noinspection SpellCheckingInspection
const apiKey = "78f9704baaab411a87edeed59052cbb687a4aa7764a44accbaf6447492b0ca8c"

func main() {
	key := ambient.NewKey(applicationKey, apiKey)
	dr, err := ambient.Device(key)
	if err != nil {
		panic(err)
	}
	switch dr.HTTPResponseCode {
	case 200:
	case 429, 502, 503:
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
			_, err := fmt.Fprintf(os.Stderr, "HTTPResponseCode=%d\n", dr.HTTPResponseCode)
			if err != nil {
				return
			}
			panic(dr)
		}
	}
	ar := make([]ambient.APIDeviceMacResponse, len(dr.DeviceRecord))
	for z := range dr.DeviceRecord {
		// API RateLimit
		time.Sleep(1 * time.Second)
		ar2, err := ambient.DeviceMac(key, dr.DeviceRecord[z].Macaddress, time.Now(), 1)
		if err != nil {
			panic(err)
		}
		ar[z] = ar2
		switch ar[z].HTTPResponseCode {
		case 200:
		case 429, 502, 503:
			{
				fmt.Printf("Error code %d, retrying.\n", ar[z].HTTPResponseCode)
				time.Sleep(1 * time.Second)
				ar2, err = ambient.DeviceMac(key, dr.DeviceRecord[z].Macaddress, time.Now(), 1)
				if err != nil {
					panic(err)
				}
				ar[z] = ar2
				switch ar[z].HTTPResponseCode {
				case 200:
				default:
					{
						_, err := fmt.Fprintf(os.Stderr, "HTTPResponseCode=%d\n", ar[z].HTTPResponseCode)
						if err != nil {
							return
						}
						panic(ar)
					}
				}
			}
		default:
			{
				_, err := fmt.Fprintf(os.Stderr, "HTTPResponseCode=%d\n", ar[z].HTTPResponseCode)
				if err != nil {
					return
				}
				panic(ar)
			}
		}
	}
	var drPrettyJSON bytes.Buffer
	err = json.Indent(&drPrettyJSON, dr.JSONResponse, "", "\t")
	if err != nil {
		return
	}
	drDeviceRecordJSON, _ := json.MarshalIndent(dr.DeviceRecord, "", "\t")
	fmt.Printf("DeviceResponse:\nHTTPResponseCode: %d, ResponseTime: %v\n", dr.HTTPResponseCode, dr.ResponseTime)
	fmt.Printf("Device Record:\n%+v\n", string(drDeviceRecordJSON))
	fmt.Printf("JSONResponse:\n%s\n\n", string(drPrettyJSON.Bytes()))
	for idx := range ar {
		var arPrettyJSON bytes.Buffer
		err := json.Indent(&arPrettyJSON, ar[idx].JSONResponse, "", "\t")
		if err != nil {
			return
		}
		arRecordJSON, _ := json.MarshalIndent(ar[idx].Record, "", "\t")
		arRecordFieldsJSON, _ := json.MarshalIndent(ar[idx].RecordFields, "", "\t")
		fmt.Printf(
			"DeviceMacResponse[%s(%s)]:\nHTTPResponseCode: %d, ResponseTime: %v\n", dr.DeviceRecord[idx].Info.Name,
			dr.DeviceRecord[idx].Macaddress, ar[idx].HTTPResponseCode, ar[idx].ResponseTime,
		)
		fmt.Printf("Record:\n%+v\n", string(arRecordJSON))
		fmt.Printf("RecordFields:\n%+v\n", string(arRecordFieldsJSON))
		fmt.Printf("JSONResponse:\n%s\n\n", string(arPrettyJSON.Bytes()))
	}
}
