# \EnergyPlansAPI

All URIs are relative to *https://mtls.dh.example.com/cds-au/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnergyPlanDetail**](EnergyPlansAPI.md#GetEnergyPlanDetail) | **Get** /energy/plans/{planId} | Get Generic Plan Detail
[**ListEnergyPlans**](EnergyPlansAPI.md#ListEnergyPlans) | **Get** /energy/plans | Get Generic Plans



## GetEnergyPlanDetail

> EnergyPlanResponseV3 GetEnergyPlanDetail(ctx, planId).XV(xV).XMinV(xMinV).Execute()

Get Generic Plan Detail



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
	planId := "planId_example" // string | ID of the specific plan requested.
	xV := "xV_example" // string | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. See [HTTP Headers](#request-headers).
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnergyPlansAPI.GetEnergyPlanDetail(context.Background(), planId).XV(xV).XMinV(xMinV).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnergyPlansAPI.GetEnergyPlanDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnergyPlanDetail`: EnergyPlanResponseV3
	fmt.Fprintf(os.Stdout, "Response from `EnergyPlansAPI.GetEnergyPlanDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | ID of the specific plan requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnergyPlanDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 

### Return type

[**EnergyPlanResponseV3**](EnergyPlanResponseV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnergyPlans

> EnergyPlanListResponse ListEnergyPlans(ctx).XV(xV).Type_(type_).FuelType(fuelType).Effective(effective).UpdatedSince(updatedSince).Brand(brand).Page(page).PageSize(pageSize).XMinV(xMinV).Execute()

Get Generic Plans



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
	type_ := "type__example" // string | Used to filter results on the _type_ field. Any one of the valid values for this field can be supplied plus `ALL`. If absent defaults to `ALL`. (optional) (default to "ALL")
	fuelType := "fuelType_example" // string | Used to filter results on the _fuelType_ field. Any one of the valid values for this field can be supplied plus `ALL`. If absent defaults to `ALL`. (optional) (default to "ALL")
	effective := "effective_example" // string | Allows for the filtering of plans based on whether the current time is within the period of time defined as effective by the _effectiveFrom_ and _effectiveTo_ fields. Valid values are `CURRENT`, `FUTURE` and `ALL`. If absent defaults to `CURRENT`. (optional) (default to "CURRENT")
	updatedSince := "updatedSince_example" // string | Only include plans that have been updated after the specified date and time. If absent defaults to include all plans. (optional)
	brand := "brand_example" // string | Used to filter results on the _brand_ field. If absent defaults to include all plans. (optional)
	page := int32(56) // int32 | Page of results to request (standard pagination). (optional)
	pageSize := int32(56) // int32 | Page size to request. Default is 25 (standard pagination). (optional)
	xMinV := "xMinV_example" // string | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a `406 Not Acceptable`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnergyPlansAPI.ListEnergyPlans(context.Background()).XV(xV).Type_(type_).FuelType(fuelType).Effective(effective).UpdatedSince(updatedSince).Brand(brand).Page(page).PageSize(pageSize).XMinV(xMinV).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnergyPlansAPI.ListEnergyPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnergyPlans`: EnergyPlanListResponse
	fmt.Fprintf(os.Stdout, "Response from `EnergyPlansAPI.ListEnergyPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEnergyPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xV** | **string** | Version of the API endpoint requested by the client. Must be set to a positive integer. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If the value of [_x-min-v_](#request-headers) is equal to or higher than the value of [_x-v_](#request-headers) then the [_x-min-v_](#request-headers) header should be treated as absent. If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. See [HTTP Headers](#request-headers). | 
 **type_** | **string** | Used to filter results on the _type_ field. Any one of the valid values for this field can be supplied plus &#x60;ALL&#x60;. If absent defaults to &#x60;ALL&#x60;. | [default to &quot;ALL&quot;]
 **fuelType** | **string** | Used to filter results on the _fuelType_ field. Any one of the valid values for this field can be supplied plus &#x60;ALL&#x60;. If absent defaults to &#x60;ALL&#x60;. | [default to &quot;ALL&quot;]
 **effective** | **string** | Allows for the filtering of plans based on whether the current time is within the period of time defined as effective by the _effectiveFrom_ and _effectiveTo_ fields. Valid values are &#x60;CURRENT&#x60;, &#x60;FUTURE&#x60; and &#x60;ALL&#x60;. If absent defaults to &#x60;CURRENT&#x60;. | [default to &quot;CURRENT&quot;]
 **updatedSince** | **string** | Only include plans that have been updated after the specified date and time. If absent defaults to include all plans. | 
 **brand** | **string** | Used to filter results on the _brand_ field. If absent defaults to include all plans. | 
 **page** | **int32** | Page of results to request (standard pagination). | 
 **pageSize** | **int32** | Page size to request. Default is 25 (standard pagination). | 
 **xMinV** | **string** | Minimum version of the API endpoint requested by the client. Must be set to a positive integer if provided. The endpoint should respond with the highest supported version between [_x-min-v_](#request-headers) and [_x-v_](#request-headers). If all versions requested are not supported then the endpoint **MUST** respond with a &#x60;406 Not Acceptable&#x60;. | 

### Return type

[**EnergyPlanListResponse**](EnergyPlanListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

