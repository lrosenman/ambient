// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright 2018 Larry Rosenman, LERCTR Consulting, larryrtx@gmail.com
//

// Package ambient provides helper functions and Go types
// for accessing ambientweather.net's API which is documented at
//      https://ambientweather.docs.apiary.io/
package ambient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const ApiVer = "v1"                                      // API Version
const ApiEP = "https://api.ambientweather.net/" + ApiVer // API Endpoint

// AmbientRecord maps the data for a specific time
// as returned by the API.
type AmbientRecord struct {
	Date           time.Time
	Baromabsin     float64
	Baromrelin     float64
	Dailyrainin    float64
	Dewpoint       float64
	Feelslike      float64
	Hourlyrainin   float64
	Humidity       int
	Humidityin     int
	LastRain       time.Time
	Maxdailygust   float64
	Monthlyrainin  float64
	Solarradiation float64
	Tempf          float64
	Tempinf        float64
	Uv             int
	Weeklyrainin   float64
	Winddir        int
	Windgustmph    float64
	Windspeedmph   float64
	Yearlyrainin   float64
}

// DeviceInfo maps the info portion of the /devices API.
type DeviceInfo struct {
	Name     string
	Location string
}

//DeviceRecord maps one record of the /devices API.
type DeviceRecord struct {
	Macaddress string
	Info       DeviceInfo
	LastData   AmbientRecord
}

// ApiDeviceMacResponse returns the data from
// /devices/macaddr API.
type ApiDeviceMacResponse struct {
	AmbientRecord    []AmbientRecord
	JSONResponse     []byte
	HTTPResponseCode int
	ResponseTime     time.Duration
}

// ApiDeviceResponse returns the data from
// /devices API.
type ApiDeviceResponse struct {
	DeviceRecord     []DeviceRecord
	JSONResponse     []byte
	HTTPResponseCode int
	ResponseTime     time.Duration
}

// holds the keys.
type Key struct {
	applicationKey string
	apiKey         string
}

// returns Key stucture to be used
func NewKey(applicationKey string, apiKey string) Key {
	return Key{applicationKey: applicationKey, apiKey: apiKey}
}

// ApiKey returns the currently set ApiKey.
func (Key Key) ApiKey() string {
	return Key.apiKey
}

// ApplicationKey returns the currently set ApplicationKey.
func (Key Key) ApplicationKey() string {
	return Key.applicationKey
}

// SetApplicationKey sets the applicationKey
func (Key Key) SetApplicationKey(applicationKey string) {
	Key.applicationKey = applicationKey
}

// SetApiKey sets the aoiKey
func (Key Key) SetApiKey(apiKey string) {
	Key.apiKey = apiKey
}

// Issue a /devices call
func Device(key Key) ApiDeviceResponse {
	var ar ApiDeviceResponse

	url := ApiEP + "/devices?applicationKey=" + key.applicationKey +
		"&apiKey=" + key.apiKey
	startTime := time.Now()
	resp, err := http.Get(url)
	ar.ResponseTime = time.Since(startTime)
	if err != nil {
		panic(err)
	}
	ar.HTTPResponseCode = resp.StatusCode
	switch resp.StatusCode {
	case 200:
	case 503, 429:
		{
			return ar
		}
	default:
		{
			panic(resp)
		}
	}
	ar.JSONResponse, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(ar.JSONResponse, &ar.DeviceRecord)
	if err != nil {
		panic(err)
	}
	return ar

}

// issue a /devices/macaddr call.
func DeviceMac(key Key, macaddr string, endtime time.Time, limit int64) ApiDeviceMacResponse {
	var ar ApiDeviceMacResponse
	url := ApiEP + "/devices/" + macaddr + "?endDate=" + url.QueryEscape(endtime.Format(time.RFC3339)) +
		"&limit=" + fmt.Sprintf("%d", limit) + "&applicationKey=" + key.applicationKey +
		"&apiKey=" + key.apiKey
	startTime := time.Now()
	resp, err := http.Get(url)
	ar.ResponseTime = time.Since(startTime)
	if err != nil {
		panic(err)
	}
	ar.HTTPResponseCode = resp.StatusCode
	switch resp.StatusCode {
	case 200:
	case 503, 429:
		{
			return ar
		}
	default:
		{
			panic(resp)
		}
	}
	ar.JSONResponse, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(ar.JSONResponse, &ar.AmbientRecord)
	if err != nil {
		panic(err)
	}
	return ar
}
