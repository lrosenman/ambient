package ambient

import (
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func Test_DeviceMac_OkResponse_ReturnsQueryResults(t *testing.T) {
	key := NewKey("application-key", "api-key")
	deviceMac := faker.MacAddress()
	endTime := time.Now().Add(time.Minute * -1)
	resultLimit := 15
	expectedResult := getValidRecordSlice(resultLimit)

	httpGet = getMockHttpGetRequest(http.StatusOK, expectedResult)

	queryResults, err := DeviceMac(key, deviceMac, endTime, int64(resultLimit))

	require.Nil(t, err)

	require.Equal(t, http.StatusOK, queryResults.HTTPResponseCode)
	require.NotNil(t, queryResults.JSONResponse)
	require.Greater(t, queryResults.ResponseTime, time.Second*0)
	requireRecordsEqualValues(t, expectedResult, queryResults.Record)
}

func Test_DeviceMac_TooManyRequestsResponse_ReturnsNoData(t *testing.T) {
	key := NewKey("application-key", "api-key")
	deviceMac := faker.MacAddress()
	endTime := time.Now().Add(time.Minute * -1)
	resultLimit := 15
	expectedResult := getValidRecordSlice(resultLimit)

	httpGet = getMockHttpGetRequest(http.StatusTooManyRequests, expectedResult)

	queryResults, err := DeviceMac(key, deviceMac, endTime, int64(resultLimit))

	require.Nil(t, err)

	require.Equal(t, http.StatusTooManyRequests, queryResults.HTTPResponseCode)
	require.NotNil(t, queryResults.JSONResponse)
	require.Empty(t, queryResults.Record)
}

func Test_DeviceMac_ServiceUnavailableResponse_ReturnsNoDataAndErrorMessage(t *testing.T) {
	key := NewKey("application-key", "api-key")
	deviceMac := faker.MacAddress()
	endTime := time.Now().Add(time.Minute * -1)
	resultLimit := 15
	expectedResult := getValidRecordSlice(resultLimit)

	httpGet = getMockHttpGetRequest(http.StatusServiceUnavailable, expectedResult)

	queryResults, err := DeviceMac(key, deviceMac, endTime, int64(resultLimit))

	require.Nil(t, err)

	require.Equal(t, http.StatusServiceUnavailable, queryResults.HTTPResponseCode)
	require.NotNil(t, queryResults.JSONResponse)
	require.Empty(t, queryResults.Record)
	require.Contains(t, string(queryResults.JSONResponse), "errormessage")
}

func Test_DeviceMac_BadGatewayResponse_ReturnsNoDataAndErrorMessage(t *testing.T) {
	key := NewKey("application-key", "api-key")
	deviceMac := faker.MacAddress()
	endTime := time.Now().Add(time.Minute * -1)
	resultLimit := 15
	expectedResult := getValidRecordSlice(resultLimit)

	httpGet = getMockHttpGetRequest(http.StatusBadGateway, expectedResult)

	queryResults, err := DeviceMac(key, deviceMac, endTime, int64(resultLimit))

	require.Nil(t, err)

	require.Equal(t, http.StatusBadGateway, queryResults.HTTPResponseCode)
	require.NotNil(t, queryResults.JSONResponse)
	require.Empty(t, queryResults.Record)
	require.Contains(t, string(queryResults.JSONResponse), "errormessage")
}

func Test_DeviceMac_NonSupportedResponse_ReturnsError(t *testing.T) {
	key := NewKey("application-key", "api-key")
	deviceMac := faker.MacAddress()
	endTime := time.Now().Add(time.Minute * -1)
	resultLimit := 15
	expectedResult := getValidRecordSlice(resultLimit)

	httpGet = getMockHttpGetRequest(http.StatusForbidden, expectedResult)

	queryResults, err := DeviceMac(key, deviceMac, endTime, int64(resultLimit))

	require.NotNil(t, err)

	require.Equal(t, http.StatusForbidden, queryResults.HTTPResponseCode)
	require.Empty(t, queryResults.Record)
}
