# \EnergyAccountBillingAPI

All URIs are relative to *https://mtls.dh.example.com/cds-au/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBillingForEnergyAccount**](EnergyAccountBillingAPI.md#GetBillingForEnergyAccount) | **Get** /energy/accounts/{accountId}/billing | Get Billing For Account
[**GetEnergyAccountInvoices**](EnergyAccountBillingAPI.md#GetEnergyAccountInvoices) | **Get** /energy/accounts/{accountId}/invoices | Get Invoices For Account
[**ListEnergyAccountBillingBulk**](EnergyAccountBillingAPI.md#ListEnergyAccountBillingBulk) | **Get** /energy/accounts/billing | Get Bulk Billing
[**ListEnergyAccountBillingForSpecificAccounts**](EnergyAccountBillingAPI.md#ListEnergyAccountBillingForSpecificAccounts) | **Post** /energy/accounts/billing | Get Billing For Specific Accounts
[**ListEnergyAccountInvoicesBulk**](EnergyAccountBillingAPI.md#ListEnergyAccountInvoicesBulk) | **Get** /energy/accounts/invoices | Get Bulk Invoices
[**ListEnergyInvoicesForSpecificAccounts**](EnergyAccountBillingAPI.md#ListEnergyInvoicesForSpecificAccounts) | **Post** /energy/accounts/invoices | Get Invoices For Specific Accounts



## GetBillingForEnergyAccount

> EnergyBillingListResponseV3 GetBillingForEnergyAccount(ctx, accountId).XV(xV).NewestTime(newestTime).OldestTime(oldestTime).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()

Get Billing For Account



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
	newestTime := "newestTime_example" // string | Constrain the request to records with effective time at or before this date/time. If absent defaults to current date/time. Format is aligned to DateTimeString common type. (optional)
	oldestTime := "oldestTime_example" // string | Constrain the request to records with effective time at or after this date/time. If absent defaults to _newest-time_ minus 12 months. Format is aligned to DateTimeString common type. (optional)
	page := int32(56) // int32 | Page of results to request (standard pagination). (optional)
	pageSize := int32(56) // int32 | Page size to request. Default is 25 (standard pagination). (optional)
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)
	xFapiInteractionId := "xFapiInteractionId_example" // string | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. (optional)
	xFapiAuthDate := "xFapiAuthDate_example" // string | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. (optional)
	xFapiCustomerIpAddress := "xFapiCustomerIpAddress_example" // string | The customer's original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. (optional)
	xCdsClientHeaders := "xCdsClientHeaders_example" // string | The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnergyAccountBillingAPI.GetBillingForEnergyAccount(context.Background(), accountId).XV(xV).NewestTime(newestTime).OldestTime(oldestTime).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnergyAccountBillingAPI.GetBillingForEnergyAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingForEnergyAccount`: EnergyBillingListResponseV3
	fmt.Fprintf(os.Stdout, "Response from `EnergyAccountBillingAPI.GetBillingForEnergyAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The _accountId_ to obtain data for. _accountId_ values are returned by account list endpoints. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingForEnergyAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **newestTime** | **string** | Constrain the request to records with effective time at or before this date/time. If absent defaults to current date/time. Format is aligned to DateTimeString common type. | 
 **oldestTime** | **string** | Constrain the request to records with effective time at or after this date/time. If absent defaults to _newest-time_ minus 12 months. Format is aligned to DateTimeString common type. | 
 **page** | **int32** | Page of results to request (standard pagination). | 
 **pageSize** | **int32** | Page size to request. Default is 25 (standard pagination). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 
 **xFapiInteractionId** | **string** | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. | 
 **xFapiAuthDate** | **string** | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. | 
 **xFapiCustomerIpAddress** | **string** | The customer&#39;s original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. | 
 **xCdsClientHeaders** | **string** | The customer&#39;s original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. | 

### Return type

[**EnergyBillingListResponseV3**](EnergyBillingListResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnergyAccountInvoices

> EnergyInvoiceListResponse GetEnergyAccountInvoices(ctx, accountId).XV(xV).NewestDate(newestDate).OldestDate(oldestDate).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()

Get Invoices For Account



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
	newestDate := "newestDate_example" // string | Constrain the request to records with issue date at or before this date. If absent defaults to current date. Format is aligned to DateString common type. (optional)
	oldestDate := "oldestDate_example" // string | Constrain the request to records with issue date at or after this date. If absent defaults to _newest-date_ minus 24 months. Format is aligned to DateString common type. (optional)
	page := int32(56) // int32 | Page of results to request (standard pagination). (optional)
	pageSize := int32(56) // int32 | Page size to request. Default is 25 (standard pagination). (optional)
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)
	xFapiInteractionId := "xFapiInteractionId_example" // string | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. (optional)
	xFapiAuthDate := "xFapiAuthDate_example" // string | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. (optional)
	xFapiCustomerIpAddress := "xFapiCustomerIpAddress_example" // string | The customer's original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. (optional)
	xCdsClientHeaders := "xCdsClientHeaders_example" // string | The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnergyAccountBillingAPI.GetEnergyAccountInvoices(context.Background(), accountId).XV(xV).NewestDate(newestDate).OldestDate(oldestDate).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnergyAccountBillingAPI.GetEnergyAccountInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnergyAccountInvoices`: EnergyInvoiceListResponse
	fmt.Fprintf(os.Stdout, "Response from `EnergyAccountBillingAPI.GetEnergyAccountInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The _accountId_ to obtain data for. _accountId_ values are returned by account list endpoints. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnergyAccountInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **newestDate** | **string** | Constrain the request to records with issue date at or before this date. If absent defaults to current date. Format is aligned to DateString common type. | 
 **oldestDate** | **string** | Constrain the request to records with issue date at or after this date. If absent defaults to _newest-date_ minus 24 months. Format is aligned to DateString common type. | 
 **page** | **int32** | Page of results to request (standard pagination). | 
 **pageSize** | **int32** | Page size to request. Default is 25 (standard pagination). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 
 **xFapiInteractionId** | **string** | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. | 
 **xFapiAuthDate** | **string** | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. | 
 **xFapiCustomerIpAddress** | **string** | The customer&#39;s original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. | 
 **xCdsClientHeaders** | **string** | The customer&#39;s original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. | 

### Return type

[**EnergyInvoiceListResponse**](EnergyInvoiceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnergyAccountBillingBulk

> EnergyBillingListResponseV3 ListEnergyAccountBillingBulk(ctx).XV(xV).NewestTime(newestTime).OldestTime(oldestTime).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()

Get Bulk Billing



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
	newestTime := "newestTime_example" // string | Constrain the request to records with effective time at or before this date/time. If absent defaults to current date/time. Format is aligned to DateTimeString common type. (optional)
	oldestTime := "oldestTime_example" // string | Constrain the request to records with effective time at or after this date/time. If absent defaults to _newest-time_ minus 12 months. Format is aligned to DateTimeString common type. (optional)
	page := int32(56) // int32 | Page of results to request (standard pagination). (optional)
	pageSize := int32(56) // int32 | Page size to request. Default is 25 (standard pagination). (optional)
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)
	xFapiInteractionId := "xFapiInteractionId_example" // string | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. (optional)
	xFapiAuthDate := "xFapiAuthDate_example" // string | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. (optional)
	xFapiCustomerIpAddress := "xFapiCustomerIpAddress_example" // string | The customer's original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. (optional)
	xCdsClientHeaders := "xCdsClientHeaders_example" // string | The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnergyAccountBillingAPI.ListEnergyAccountBillingBulk(context.Background()).XV(xV).NewestTime(newestTime).OldestTime(oldestTime).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnergyAccountBillingAPI.ListEnergyAccountBillingBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnergyAccountBillingBulk`: EnergyBillingListResponseV3
	fmt.Fprintf(os.Stdout, "Response from `EnergyAccountBillingAPI.ListEnergyAccountBillingBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEnergyAccountBillingBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **newestTime** | **string** | Constrain the request to records with effective time at or before this date/time. If absent defaults to current date/time. Format is aligned to DateTimeString common type. | 
 **oldestTime** | **string** | Constrain the request to records with effective time at or after this date/time. If absent defaults to _newest-time_ minus 12 months. Format is aligned to DateTimeString common type. | 
 **page** | **int32** | Page of results to request (standard pagination). | 
 **pageSize** | **int32** | Page size to request. Default is 25 (standard pagination). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 
 **xFapiInteractionId** | **string** | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. | 
 **xFapiAuthDate** | **string** | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. | 
 **xFapiCustomerIpAddress** | **string** | The customer&#39;s original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. | 
 **xCdsClientHeaders** | **string** | The customer&#39;s original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. | 

### Return type

[**EnergyBillingListResponseV3**](EnergyBillingListResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnergyAccountBillingForSpecificAccounts

> EnergyBillingListResponseV3 ListEnergyAccountBillingForSpecificAccounts(ctx).XV(xV).RequestAccountIdListV1(requestAccountIdListV1).NewestTime(newestTime).OldestTime(oldestTime).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()

Get Billing For Specific Accounts



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
	requestAccountIdListV1 := *openapiclient.NewRequestAccountIdListV1(*openapiclient.NewRequestAccountIdListV1Data([]string{"AccountIds_example"})) // RequestAccountIdListV1 | Request payload containing a list of _accountId_ values to obtain data for.
	newestTime := "newestTime_example" // string | Constrain the request to records with effective time at or before this date/time. If absent defaults to current date/time. Format is aligned to DateTimeString common type. (optional)
	oldestTime := "oldestTime_example" // string | Constrain the request to records with effective time at or after this date/time. If absent defaults to _newest-time_ minus 12 months. Format is aligned to DateTimeString common type. (optional)
	page := int32(56) // int32 | Page of results to request (standard pagination). (optional)
	pageSize := int32(56) // int32 | Page size to request. Default is 25 (standard pagination). (optional)
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)
	xFapiInteractionId := "xFapiInteractionId_example" // string | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. (optional)
	xFapiAuthDate := "xFapiAuthDate_example" // string | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. (optional)
	xFapiCustomerIpAddress := "xFapiCustomerIpAddress_example" // string | The customer's original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. (optional)
	xCdsClientHeaders := "xCdsClientHeaders_example" // string | The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnergyAccountBillingAPI.ListEnergyAccountBillingForSpecificAccounts(context.Background()).XV(xV).RequestAccountIdListV1(requestAccountIdListV1).NewestTime(newestTime).OldestTime(oldestTime).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnergyAccountBillingAPI.ListEnergyAccountBillingForSpecificAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnergyAccountBillingForSpecificAccounts`: EnergyBillingListResponseV3
	fmt.Fprintf(os.Stdout, "Response from `EnergyAccountBillingAPI.ListEnergyAccountBillingForSpecificAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEnergyAccountBillingForSpecificAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **requestAccountIdListV1** | [**RequestAccountIdListV1**](RequestAccountIdListV1.md) | Request payload containing a list of _accountId_ values to obtain data for. | 
 **newestTime** | **string** | Constrain the request to records with effective time at or before this date/time. If absent defaults to current date/time. Format is aligned to DateTimeString common type. | 
 **oldestTime** | **string** | Constrain the request to records with effective time at or after this date/time. If absent defaults to _newest-time_ minus 12 months. Format is aligned to DateTimeString common type. | 
 **page** | **int32** | Page of results to request (standard pagination). | 
 **pageSize** | **int32** | Page size to request. Default is 25 (standard pagination). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 
 **xFapiInteractionId** | **string** | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. | 
 **xFapiAuthDate** | **string** | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. | 
 **xFapiCustomerIpAddress** | **string** | The customer&#39;s original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. | 
 **xCdsClientHeaders** | **string** | The customer&#39;s original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. | 

### Return type

[**EnergyBillingListResponseV3**](EnergyBillingListResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnergyAccountInvoicesBulk

> EnergyInvoiceListResponse ListEnergyAccountInvoicesBulk(ctx).XV(xV).NewestDate(newestDate).OldestDate(oldestDate).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()

Get Bulk Invoices



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
	newestDate := "newestDate_example" // string | Constrain the request to records with issue date at or before this date. If absent defaults to current date. Format is aligned to DateString common type. (optional)
	oldestDate := "oldestDate_example" // string | Constrain the request to records with issue date at or after this date. If absent defaults to _newest-date_ minus 24 months. Format is aligned to DateString common type. (optional)
	page := int32(56) // int32 | Page of results to request (standard pagination). (optional)
	pageSize := int32(56) // int32 | Page size to request. Default is 25 (standard pagination). (optional)
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)
	xFapiInteractionId := "xFapiInteractionId_example" // string | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. (optional)
	xFapiAuthDate := "xFapiAuthDate_example" // string | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. (optional)
	xFapiCustomerIpAddress := "xFapiCustomerIpAddress_example" // string | The customer's original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. (optional)
	xCdsClientHeaders := "xCdsClientHeaders_example" // string | The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnergyAccountBillingAPI.ListEnergyAccountInvoicesBulk(context.Background()).XV(xV).NewestDate(newestDate).OldestDate(oldestDate).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnergyAccountBillingAPI.ListEnergyAccountInvoicesBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnergyAccountInvoicesBulk`: EnergyInvoiceListResponse
	fmt.Fprintf(os.Stdout, "Response from `EnergyAccountBillingAPI.ListEnergyAccountInvoicesBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEnergyAccountInvoicesBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **newestDate** | **string** | Constrain the request to records with issue date at or before this date. If absent defaults to current date. Format is aligned to DateString common type. | 
 **oldestDate** | **string** | Constrain the request to records with issue date at or after this date. If absent defaults to _newest-date_ minus 24 months. Format is aligned to DateString common type. | 
 **page** | **int32** | Page of results to request (standard pagination). | 
 **pageSize** | **int32** | Page size to request. Default is 25 (standard pagination). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 
 **xFapiInteractionId** | **string** | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. | 
 **xFapiAuthDate** | **string** | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. | 
 **xFapiCustomerIpAddress** | **string** | The customer&#39;s original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. | 
 **xCdsClientHeaders** | **string** | The customer&#39;s original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. | 

### Return type

[**EnergyInvoiceListResponse**](EnergyInvoiceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnergyInvoicesForSpecificAccounts

> EnergyInvoiceListResponse ListEnergyInvoicesForSpecificAccounts(ctx).XV(xV).RequestAccountIdListV1(requestAccountIdListV1).NewestDate(newestDate).OldestDate(oldestDate).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()

Get Invoices For Specific Accounts



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
	requestAccountIdListV1 := *openapiclient.NewRequestAccountIdListV1(*openapiclient.NewRequestAccountIdListV1Data([]string{"AccountIds_example"})) // RequestAccountIdListV1 | Request payload containing a list of _accountId_ values to obtain data for.
	newestDate := "newestDate_example" // string | Constrain the request to records with issue date at or before this date. If absent defaults to current date. Format is aligned to DateString common type. (optional)
	oldestDate := "oldestDate_example" // string | Constrain the request to records with issue date at or after this date. If absent defaults to _newest-date_ minus 24 months. Format is aligned to DateString common type. (optional)
	page := int32(56) // int32 | Page of results to request (standard pagination). (optional)
	pageSize := int32(56) // int32 | Page size to request. Default is 25 (standard pagination). (optional)
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)
	xFapiInteractionId := "xFapiInteractionId_example" // string | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. (optional)
	xFapiAuthDate := "xFapiAuthDate_example" // string | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. (optional)
	xFapiCustomerIpAddress := "xFapiCustomerIpAddress_example" // string | The customer's original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. (optional)
	xCdsClientHeaders := "xCdsClientHeaders_example" // string | The customer's original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnergyAccountBillingAPI.ListEnergyInvoicesForSpecificAccounts(context.Background()).XV(xV).RequestAccountIdListV1(requestAccountIdListV1).NewestDate(newestDate).OldestDate(oldestDate).Page(page).PageSize(pageSize).XMinV(xMinV).XFapiInteractionId(xFapiInteractionId).XFapiAuthDate(xFapiAuthDate).XFapiCustomerIpAddress(xFapiCustomerIpAddress).XCdsClientHeaders(xCdsClientHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnergyAccountBillingAPI.ListEnergyInvoicesForSpecificAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnergyInvoicesForSpecificAccounts`: EnergyInvoiceListResponse
	fmt.Fprintf(os.Stdout, "Response from `EnergyAccountBillingAPI.ListEnergyInvoicesForSpecificAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEnergyInvoicesForSpecificAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **requestAccountIdListV1** | [**RequestAccountIdListV1**](RequestAccountIdListV1.md) | Request payload containing a list of _accountId_ values to obtain data for. | 
 **newestDate** | **string** | Constrain the request to records with issue date at or before this date. If absent defaults to current date. Format is aligned to DateString common type. | 
 **oldestDate** | **string** | Constrain the request to records with issue date at or after this date. If absent defaults to _newest-date_ minus 24 months. Format is aligned to DateString common type. | 
 **page** | **int32** | Page of results to request (standard pagination). | 
 **pageSize** | **int32** | Page size to request. Default is 25 (standard pagination). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 
 **xFapiInteractionId** | **string** | An **[[RFC4122]](#nref-RFC4122)** UUID used as a correlation id. If provided, the data holder **MUST** play back this value in the _x-fapi-interaction-id_ response header. If not provided a **[[RFC4122]](#nref-RFC4122)** UUID value is required to be provided in the response header to track the interaction. | 
 **xFapiAuthDate** | **string** | The time when the customer last logged in to the Data Recipient Software Product as described in **[[FAPI-1.0-Baseline]](#nref-FAPI-1-0-Baseline)**. Required for all resource calls (customer present and unattended). Not required for unauthenticated calls. | 
 **xFapiCustomerIpAddress** | **string** | The customer&#39;s original IP address if the customer is currently logged in to the data recipient. The presence of this header indicates that the API is being called in a customer present context. Not to be included for unauthenticated calls. | 
 **xCdsClientHeaders** | **string** | The customer&#39;s original standard http headers [Base64](#common-field-types) encoded, including the original User-Agent header, if the customer is currently logged in to the data recipient. Mandatory for customer present calls. Not required for unattended or unauthenticated calls. | 

### Return type

[**EnergyInvoiceListResponse**](EnergyInvoiceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

