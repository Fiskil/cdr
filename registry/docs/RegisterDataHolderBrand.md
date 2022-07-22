# RegisterDataHolderBrand

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataHolderBrandId** | **string** | Unique id of the Data Holder Brand issued by the CDR Register | [default to null]
**BrandName** | **string** | The name of Data Holder Brand | [default to null]
**Industries** | **[]string** | The industries the Data Holder Brand belongs to. Please note that the CDR Register entity model is constrained to one industry per brand which is planned to be relaxed in the future. | [default to null]
**LogoUri** | **string** | Brand logo URI | [default to null]
**LegalEntity** | [***LegalEntityDetail**](LegalEntityDetail.md) |  | [default to null]
**Status** | **string** |  | [default to null]
**EndpointDetail** | [***RegisterDataHolderBrandServiceEndpoint**](RegisterDataHolderBrandServiceEndpoint.md) |  | [default to null]
**AuthDetails** | [**[]RegisterDataHolderAuth**](RegisterDataHolderAuth.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) | The date/time that the Data Holder Brand data was last updated in the Register | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

