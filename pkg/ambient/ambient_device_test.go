package ambient

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func Test_Device_OkResponse_ReturnsListOfDevices(t *testing.T) {
	expectedResult := []*DeviceRecord{
		getValidDeviceRecord(),
		getValidDeviceRecord(),
	}

	httpGet = getMockHttpGetRequest(http.StatusOK, expectedResult)

	key := NewKey("application-key", "api-key")
	devices, err := Device(key)

	require.Nil(t, err)

	require.Equal(t, http.StatusOK, devices.HTTPResponseCode)
	require.NotNil(t, devices.JSONResponse)
	require.Greater(t, devices.ResponseTime, time.Second*0)
	requireDeviceRecordsEqualValues(t, expectedResult, devices.DeviceRecord)
}

func Test_Device_TooManyRequestsResponse_ReturnsNoData(t *testing.T) {
	expectedResult := []*DeviceRecord{
		getValidDeviceRecord(),
		getValidDeviceRecord(),
	}

	httpGet = getMockHttpGetRequest(http.StatusTooManyRequests, expectedResult)

	key := NewKey("application-key", "api-key")
	devices, err := Device(key)

	require.Nil(t, err)

	require.Equal(t, http.StatusTooManyRequests, devices.HTTPResponseCode)
	require.Empty(t, devices.DeviceRecord)
	require.NotContains(t, string(devices.JSONResponse), "errormessage")
}

func Test_Device_ServiceUnavailableResponse_ReturnsNoDataWithErrorBody(t *testing.T) {
	expectedResult := []*DeviceRecord{
		getValidDeviceRecord(),
		getValidDeviceRecord(),
	}

	httpGet = getMockHttpGetRequest(http.StatusServiceUnavailable, expectedResult)

	key := NewKey("application-key", "api-key")
	devices, err := Device(key)

	require.Nil(t, err)

	require.Equal(t, http.StatusServiceUnavailable, devices.HTTPResponseCode)
	require.Empty(t, devices.DeviceRecord)
	require.Contains(t, string(devices.JSONResponse), "errormessage")
}

func Test_Device_BadGatewayResponse_ReturnsNoDataNoDataWithErrorBody(t *testing.T) {
	expectedResult := []*DeviceRecord{
		getValidDeviceRecord(),
		getValidDeviceRecord(),
	}

	httpGet = getMockHttpGetRequest(http.StatusBadGateway, expectedResult)

	key := NewKey("application-key", "api-key")
	devices, err := Device(key)

	require.Nil(t, err)

	require.Equal(t, http.StatusBadGateway, devices.HTTPResponseCode)
	require.Empty(t, devices.DeviceRecord)
	require.Contains(t, string(devices.JSONResponse), "errormessage")
}

func Test_Device_NonSupportedResponse_ReturnsError(t *testing.T) {
	expectedResult := []*DeviceRecord{
		getValidDeviceRecord(),
		getValidDeviceRecord(),
	}

	httpGet = getMockHttpGetRequest(http.StatusForbidden, expectedResult)

	key := NewKey("application-key", "api-key")
	devices, err := Device(key)

	require.NotNil(t, err)

	require.Equal(t, http.StatusForbidden, devices.HTTPResponseCode)
	require.Empty(t, devices.DeviceRecord)
}
