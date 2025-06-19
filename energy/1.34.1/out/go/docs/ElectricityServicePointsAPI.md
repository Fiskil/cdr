# \ElectricityServicePointsAPI

All URIs are relative to *https://mtls.dh.example.com/cds-au/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetElectricityServicePointDetail**](ElectricityServicePointsAPI.md#GetElectricityServicePointDetail) | **Get** /energy/electricity/servicepoints/{servicePointId} | Get Service Point Detail
[**ListElectricityServicePoints**](ElectricityServicePointsAPI.md#ListElectricityServicePoints) | **Get** /energy/electricity/servicepoints | Get Service Points



## GetElectricityServicePointDetail

> EnergyServicePointDetailResponseV2 GetElectricityServicePointDetail(ctx, servicePointId).XV(xV).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()

Get Service Point Detail



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	servicePointId := "servicePointId_example" // string | The _servicePointId_ to obtain data for. _servicePointId_ values are returned by service point list endpoints. Note that it is not a _nationalMeteringId_.
	xV := "xV_example" // string | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. See [HTTP Headers](#request-headers).
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)
	xFapiInteractionId := "xFapiInteractionId_example" // string | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. (optional)
	xFapiAuthDate := "xFapiAuthDate_example" // string | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. (optional)
	xFapiCustomerIpAddress := "xFapiCustomerIpAddress_example" // string | The customer's original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. (optional)
	xCdsClientHeaders := "xCdsClientHeaders_example" // string | The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElectricityServicePointsAPI.GetElectricityServicePointDetail(context.Background(), servicePointId).XV(xV).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElectricityServicePointsAPI.GetElectricityServicePointDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetElectricityServicePointDetail`: EnergyServicePointDetailResponseV2
	fmt.Fprintf(os.Stdout, "Response from `ElectricityServicePointsAPI.GetElectricityServicePointDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePointId** | **string** | The _servicePointId_ to obtain data for. _servicePointId_ values are returned by service point list endpoints. Note that it is not a _nationalMeteringId_. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElectricityServicePointDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 
 **xFapiInteractionId** | **string** | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. | 
 **xFapiAuthDate** | **string** | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. | 
 **xFapiCustomerIpAddress** | **string** | The customer&#39;s original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. | 
 **xCdsClientHeaders** | **string** | The customer&#39;s original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. | 

### Return type

[**EnergyServicePointDetailResponseV2**](EnergyServicePointDetailResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListElectricityServicePoints

> EnergyServicePointListResponseV2 ListElectricityServicePoints(ctx).XV(xV).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()

Get Service Points



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xV := "xV_example" // string | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. See [HTTP Headers](#request-headers).
	page := int32(56) // int32 | Page of results to request (standard pagination). (optional)
	pageSize := int32(56) // int32 | Page size to request. Default is 25 (standard pagination). (optional)
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)
	xFapiInteractionId := "xFapiInteractionId_example" // string | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. (optional)
	xFapiAuthDate := "xFapiAuthDate_example" // string | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. (optional)
	xFapiCustomerIpAddress := "xFapiCustomerIpAddress_example" // string | The customer's original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. (optional)
	xCdsClientHeaders := "xCdsClientHeaders_example" // string | The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElectricityServicePointsAPI.ListElectricityServicePoints(context.Background()).XV(xV).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElectricityServicePointsAPI.ListElectricityServicePoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListElectricityServicePoints`: EnergyServicePointListResponseV2
	fmt.Fprintf(os.Stdout, "Response from `ElectricityServicePointsAPI.ListElectricityServicePoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListElectricityServicePointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **page** | **int32** | Page of results to request (standard pagination). | 
 **pageSize** | **int32** | Page size to request. Default is 25 (standard pagination). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 
 **xFapiInteractionId** | **string** | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. | 
 **xFapiAuthDate** | **string** | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. | 
 **xFapiCustomerIpAddress** | **string** | The customer&#39;s original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. | 
 **xCdsClientHeaders** | **string** | The customer&#39;s original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. | 

### Return type

[**EnergyServicePointListResponseV2**](EnergyServicePointListResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

