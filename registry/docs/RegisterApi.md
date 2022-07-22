# {{classname}}

All URIs are relative to *https://&lt;register-base-url&gt;/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataHolderBrands**](RegisterApi.md#GetDataHolderBrands) | **Get** /cdr-register/v1/{industry}/data-holders/brands | Get Data Holder Brands
[**GetDataHolderBrandsSummary**](RegisterApi.md#GetDataHolderBrandsSummary) | **Get** /cdr-register/v1/{industry}/data-holders/brands/summary | Get Data Holder Brands Summary
[**GetDataHolderStatuses**](RegisterApi.md#GetDataHolderStatuses) | **Get** /cdr-register/v1/{industry}/data-holders/status | Get Data Holder Statuses
[**GetDataRecipients**](RegisterApi.md#GetDataRecipients) | **Get** /cdr-register/v1/{industry}/data-recipients | Get Data Recipients
[**GetDataRecipientsStatuses**](RegisterApi.md#GetDataRecipientsStatuses) | **Get** /cdr-register/v1/{industry}/data-recipients/status | Get Data Recipients Statuses
[**GetJWKS**](RegisterApi.md#GetJWKS) | **Get** /jwks | Get JWKS
[**GetOpenIdProviderConfig**](RegisterApi.md#GetOpenIdProviderConfig) | **Get** /.well-known/openid-configuration | Get OpenId Provider Config
[**GetSoftwareProductsStatuses**](RegisterApi.md#GetSoftwareProductsStatuses) | **Get** /cdr-register/v1/{industry}/data-recipients/brands/software-products/status | Get Software Products Statuses
[**GetSoftwareStatementAssertion**](RegisterApi.md#GetSoftwareStatementAssertion) | **Get** /cdr-register/v1/{industry}/data-recipients/brands/{dataRecipientBrandId}/software-products/{softwareProductId}/ssa | Get Software Statement Assertion (SSA)

# **GetDataHolderBrands**
> ResponseRegisterDataHolderBrandList GetDataHolderBrands(ctx, industry, authorization, optional)
Get Data Holder Brands

Allows Data Recipients to discover Data Holder Brands available in the CDR ecosystem.  Obsolete versions: [v1](includes/obsolete/get-data-holder-brands-v1.html)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **industry** | **string**| The industry the participant is retrieving data for (Banking, etc) | 
  **authorization** | **string**| An Authorisation Token as per **[[RFC6750]](#nref-RFC6750)**. | 
 **optional** | ***RegisterApiGetDataHolderBrandsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegisterApiGetDataHolderBrandsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xV** | **optional.String**| The version of the API end point requested by the client. Must be set to a positive integer. | 
 **xMinV** | **optional.String**| The [minimum version](https://consumerdatastandardsaustralia.github.io/standards/#http-headers) of the API end point requested by the client. Must be set to a positive integer if provided. | 
 **updatedSince** | **optional.Time**| query filter returns results updated since the specified date-time | 
 **page** | **optional.Int32**| the page number to return | 
 **pageSize** | **optional.Int32**| the number of records to return per page | 

### Return type

[**ResponseRegisterDataHolderBrandList**](ResponseRegisterDataHolderBrandList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDataHolderBrandsSummary**
> ResponseDataHoldersBrandSummaryList GetDataHolderBrandsSummary(ctx, industry, xV, optional)
Get Data Holder Brands Summary

Endpoint used by participants to discover public details of Data Holder Brands from the CDR Register

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **industry** | **string**| The industry the participant is retrieving data for (Banking, etc) | 
  **xV** | **string**| The version of the API end point requested by the client. Must be set to a positive integer. | 
 **optional** | ***RegisterApiGetDataHolderBrandsSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegisterApiGetDataHolderBrandsSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xMinV** | **optional.String**| The [minimum version](https://consumerdatastandardsaustralia.github.io/standards/#http-headers) of the API end point requested by the client. Must be set to a positive integer if provided. | 
 **ifNoneMatch** | **optional.String**| Makes the request method conditional on a recipient cache or origin server not having any current representation of the target resource with an entity-tag that does not match any of those listed in the field-value. | 

### Return type

[**ResponseDataHoldersBrandSummaryList**](ResponseDataHoldersBrandSummaryList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDataHolderStatuses**
> DataHoldersStatusList GetDataHolderStatuses(ctx, industry, optional)
Get Data Holder Statuses

Endpoint used by participants to discover the statuses for Data Holders from the CDR Register

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **industry** | **string**| The industry the participant is retrieving data for (Banking, etc) | 
 **optional** | ***RegisterApiGetDataHolderStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegisterApiGetDataHolderStatusesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xV** | **optional.String**| The version of the API end point requested by the client. Must be set to a positive integer. | 
 **xMinV** | **optional.String**| The [minimum version](https://consumerdatastandardsaustralia.github.io/standards/#http-headers) of the API end point requested by the client. Must be set to a positive integer if provided. | 
 **ifNoneMatch** | **optional.String**| Makes the request method conditional on a recipient cache or origin server not having any current representation of the target resource with an entity-tag that does not match any of those listed in the field-value. | 

### Return type

[**DataHoldersStatusList**](DataHoldersStatusList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDataRecipients**
> ResponseRegisterDataRecipientList GetDataRecipients(ctx, industry, optional)
Get Data Recipients

Endpoint used by participants to discover data recipients and associated brands and software products, available in the CDR ecosystem.  Obsolete versions: [v2](includes/obsolete/get-data-recipients-v2.html)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **industry** | **string**| The industry the participant is retrieving data for (Banking, etc) | 
 **optional** | ***RegisterApiGetDataRecipientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegisterApiGetDataRecipientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xV** | **optional.String**| The version of the API end point requested by the client. Must be set to a positive integer. | 
 **xMinV** | **optional.String**| The [minimum version](https://consumerdatastandardsaustralia.github.io/standards/#http-headers) of the API end point requested by the client. Must be set to a positive integer if provided. | 
 **ifNoneMatch** | **optional.String**| Makes the request method conditional on a recipient cache or origin server not having any current representation of the target resource with an entity-tag that does not match any of those listed in the field-value. | 

### Return type

[**ResponseRegisterDataRecipientList**](ResponseRegisterDataRecipientList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDataRecipientsStatuses**
> DataRecipientsStatusList GetDataRecipientsStatuses(ctx, industry, optional)
Get Data Recipients Statuses

Endpoint used by participants to discover the statuses for Data Recipients from the CDR Register.  Obsolete versions: [v1](includes/obsolete/get-data-recipient-statuses-v1.html)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **industry** | **string**| The industry the participant is retrieving data for (Banking, etc) | 
 **optional** | ***RegisterApiGetDataRecipientsStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegisterApiGetDataRecipientsStatusesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xV** | **optional.String**| The version of the API end point requested by the client. Must be set to a positive integer. | 
 **xMinV** | **optional.String**| The [minimum version](https://consumerdatastandardsaustralia.github.io/standards/#http-headers) of the API end point requested by the client. Must be set to a positive integer if provided. | 
 **ifNoneMatch** | **optional.String**| Makes the request method conditional on a recipient cache or origin server not having any current representation of the target resource with an entity-tag that does not match any of those listed in the field-value. | 

### Return type

[**DataRecipientsStatusList**](DataRecipientsStatusList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJWKS**
> ResponseJwks GetJWKS(ctx, )
Get JWKS

JWKS endpoint containing the public keys used by the CDR Register to validate the signature of issued SSAs and authenticate outbound calls to participants in the CDR.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseJwks**](ResponseJWKS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpenIdProviderConfig**
> ResponseOpenIdProviderConfigMetadata GetOpenIdProviderConfig(ctx, )
Get OpenId Provider Config

Endpoint used by participants to discover the CDR Register OpenID configuration and obtain information needed to interact with it, including its OAuth 2.0 endpoint locations.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseOpenIdProviderConfigMetadata**](ResponseOpenIDProviderConfigMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSoftwareProductsStatuses**
> SoftwareProductsStatusList GetSoftwareProductsStatuses(ctx, industry, optional)
Get Software Products Statuses

Endpoint used by participants to discover the statuses for software products from the CDR Register.  Obsolete versions: [v1](includes/obsolete/get-software-product-statuses-v1.html)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **industry** | **string**| The industry the participant is retrieving data for (Banking, etc) | 
 **optional** | ***RegisterApiGetSoftwareProductsStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegisterApiGetSoftwareProductsStatusesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xV** | **optional.String**| The version of the API end point requested by the client. Must be set to a positive integer. | 
 **xMinV** | **optional.String**| The [minimum version](https://consumerdatastandardsaustralia.github.io/standards/#http-headers) of the API end point requested by the client. Must be set to a positive integer if provided. | 
 **ifNoneMatch** | **optional.String**| Makes the request method conditional on a recipient cache or origin server not having any current representation of the target resource with an entity-tag that does not match any of those listed in the field-value. | 

### Return type

[**SoftwareProductsStatusList**](SoftwareProductsStatusList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSoftwareStatementAssertion**
> string GetSoftwareStatementAssertion(ctx, industry, dataRecipientBrandId, softwareProductId, authorization, optional)
Get Software Statement Assertion (SSA)

Get a Software Statement Assertion (SSA) for a software product on the CDR Register to be used for Dynamic Client Registration with a Data Holder Brand.  Obsolete versions: [v2](includes/obsolete/get-software-statement-assertion-v2.html)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **industry** | **string**| The industry the participant is retrieving data for (Banking, etc) | 
  **dataRecipientBrandId** | **string**| Unique id for the Accredited Data Recipient Brand that the Software Product is associated with in the CDR Register | 
  **softwareProductId** | **string**| Unique id for the Accredited Data Recipient Software Product in the CDR Register | 
  **authorization** | **string**| An Authorisation Token as per **[[RFC6750]](#nref-RFC6750)**. | 
 **optional** | ***RegisterApiGetSoftwareStatementAssertionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegisterApiGetSoftwareStatementAssertionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xV** | **optional.String**| The version of the API end point requested by the client. Must be set to a positive integer. | 
 **xMinV** | **optional.String**| The [minimum version](https://consumerdatastandardsaustralia.github.io/standards/#http-headers) of the API end point requested by the client. Must be set to a positive integer if provided. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

