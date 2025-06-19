# \EnergyAccountsAPI

All URIs are relative to *https://mtls.dh.example.com/cds-au/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnergyAccountConcessions**](EnergyAccountsAPI.md#GetEnergyAccountConcessions) | **Get** /energy/accounts/{accountId}/concessions | Get Concessions
[**GetEnergyAccountDetail**](EnergyAccountsAPI.md#GetEnergyAccountDetail) | **Get** /energy/accounts/{accountId} | Get Energy Account Detail
[**GetEnergyAccountPaymentSchedule**](EnergyAccountsAPI.md#GetEnergyAccountPaymentSchedule) | **Get** /energy/accounts/{accountId}/payment-schedule | Get Agreed Payment Schedule
[**ListEnergyAccounts**](EnergyAccountsAPI.md#ListEnergyAccounts) | **Get** /energy/accounts | Get Energy Accounts



## GetEnergyAccountConcessions

> EnergyConcessionsResponse GetEnergyAccountConcessions(ctx, accountId).XV(xV).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()

Get Concessions



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
	accountId := "accountId_example" // string | The _accountId_ to obtain data for. _accountId_ values are returned by account list endpoints.
	xV := "xV_example" // string | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. See [HTTP Headers](#request-headers).
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)
	xFapiInteractionId := "xFapiInteractionId_example" // string | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. (optional)
	xFapiAuthDate := "xFapiAuthDate_example" // string | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. (optional)
	xFapiCustomerIpAddress := "xFapiCustomerIpAddress_example" // string | The customer's original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. (optional)
	xCdsClientHeaders := "xCdsClientHeaders_example" // string | The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnergyAccountsAPI.GetEnergyAccountConcessions(context.Background(), accountId).XV(xV).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnergyAccountsAPI.GetEnergyAccountConcessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnergyAccountConcessions`: EnergyConcessionsResponse
	fmt.Fprintf(os.Stdout, "Response from `EnergyAccountsAPI.GetEnergyAccountConcessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The _accountId_ to obtain data for. _accountId_ values are returned by account list endpoints. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnergyAccountConcessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 
 **xFapiInteractionId** | **string** | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. | 
 **xFapiAuthDate** | **string** | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. | 
 **xFapiCustomerIpAddress** | **string** | The customer&#39;s original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. | 
 **xCdsClientHeaders** | **string** | The customer&#39;s original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. | 

### Return type

[**EnergyConcessionsResponse**](EnergyConcessionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnergyAccountDetail

> EnergyAccountDetailResponseV4 GetEnergyAccountDetail(ctx, accountId).XV(xV).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()

Get Energy Account Detail



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
	accountId := "accountId_example" // string | The _accountId_ to obtain data for. _accountId_ values are returned by account list endpoints.
	xV := "xV_example" // string | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. See [HTTP Headers](#request-headers).
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)
	xFapiInteractionId := "xFapiInteractionId_example" // string | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. (optional)
	xFapiAuthDate := "xFapiAuthDate_example" // string | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. (optional)
	xFapiCustomerIpAddress := "xFapiCustomerIpAddress_example" // string | The customer's original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. (optional)
	xCdsClientHeaders := "xCdsClientHeaders_example" // string | The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnergyAccountsAPI.GetEnergyAccountDetail(context.Background(), accountId).XV(xV).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnergyAccountsAPI.GetEnergyAccountDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnergyAccountDetail`: EnergyAccountDetailResponseV4
	fmt.Fprintf(os.Stdout, "Response from `EnergyAccountsAPI.GetEnergyAccountDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The _accountId_ to obtain data for. _accountId_ values are returned by account list endpoints. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnergyAccountDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 
 **xFapiInteractionId** | **string** | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. | 
 **xFapiAuthDate** | **string** | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. | 
 **xFapiCustomerIpAddress** | **string** | The customer&#39;s original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. | 
 **xCdsClientHeaders** | **string** | The customer&#39;s original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. | 

### Return type

[**EnergyAccountDetailResponseV4**](EnergyAccountDetailResponseV4.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnergyAccountPaymentSchedule

> EnergyPaymentScheduleResponse GetEnergyAccountPaymentSchedule(ctx, accountId).XV(xV).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()

Get Agreed Payment Schedule



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
	accountId := "accountId_example" // string | The _accountId_ to obtain data for. _accountId_ values are returned by account list endpoints.
	xV := "xV_example" // string | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. See [HTTP Headers](#request-headers).
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)
	xFapiInteractionId := "xFapiInteractionId_example" // string | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. (optional)
	xFapiAuthDate := "xFapiAuthDate_example" // string | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. (optional)
	xFapiCustomerIpAddress := "xFapiCustomerIpAddress_example" // string | The customer's original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. (optional)
	xCdsClientHeaders := "xCdsClientHeaders_example" // string | The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnergyAccountsAPI.GetEnergyAccountPaymentSchedule(context.Background(), accountId).XV(xV).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnergyAccountsAPI.GetEnergyAccountPaymentSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnergyAccountPaymentSchedule`: EnergyPaymentScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `EnergyAccountsAPI.GetEnergyAccountPaymentSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The _accountId_ to obtain data for. _accountId_ values are returned by account list endpoints. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnergyAccountPaymentScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 
 **xFapiInteractionId** | **string** | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. | 
 **xFapiAuthDate** | **string** | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. | 
 **xFapiCustomerIpAddress** | **string** | The customer&#39;s original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. | 
 **xCdsClientHeaders** | **string** | The customer&#39;s original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. | 

### Return type

[**EnergyPaymentScheduleResponse**](EnergyPaymentScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnergyAccounts

> EnergyAccountListResponseV2 ListEnergyAccounts(ctx).XV(xV).OpenStatus(openStatus).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()

Get Energy Accounts



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
	openStatus := "openStatus_example" // string | Used to filter results according to open/closed status. Values can be `OPEN`, `CLOSED` or `ALL`. If absent then `ALL` is assumed. (optional) (default to "ALL")
	page := int32(56) // int32 | Page of results to request (standard pagination). (optional)
	pageSize := int32(56) // int32 | Page size to request. Default is 25 (standard pagination). (optional)
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)
	xFapiInteractionId := "xFapiInteractionId_example" // string | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. (optional)
	xFapiAuthDate := "xFapiAuthDate_example" // string | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. (optional)
	xFapiCustomerIpAddress := "xFapiCustomerIpAddress_example" // string | The customer's original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. (optional)
	xCdsClientHeaders := "xCdsClientHeaders_example" // string | The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnergyAccountsAPI.ListEnergyAccounts(context.Background()).XV(xV).OpenStatus(openStatus).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnergyAccountsAPI.ListEnergyAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnergyAccounts`: EnergyAccountListResponseV2
	fmt.Fprintf(os.Stdout, "Response from `EnergyAccountsAPI.ListEnergyAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEnergyAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **openStatus** | **string** | Used to filter results according to open/closed status. Values can be &#x60;OPEN&#x60;, &#x60;CLOSED&#x60; or &#x60;ALL&#x60;. If absent then &#x60;ALL&#x60; is assumed. | [default to &quot;ALL&quot;]
 **page** | **int32** | Page of results to request (standard pagination). | 
 **pageSize** | **int32** | Page size to request. Default is 25 (standard pagination). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 
 **xFapiInteractionId** | **string** | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. | 
 **xFapiAuthDate** | **string** | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. | 
 **xFapiCustomerIpAddress** | **string** | The customer&#39;s original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. | 
 **xCdsClientHeaders** | **string** | The customer&#39;s original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. | 

### Return type

[**EnergyAccountListResponseV2**](EnergyAccountListResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

