# DataHolderBrandSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataHolderBrandId** | **string** | Unique id of the Data Holder Brand issued by the CDR Register | [optional] [default to null]
**InterimId** | **string** | Interim id of the Data Holder Brand issued by the CDR Register. This is to be used to uniquely identify the record when dataHolderBrandId is not populated and is not to be reused | [optional] [default to null]
**BrandName** | **string** | The name of Data Holder Brand | [default to null]
**PublicBaseUri** | **string** | Base URI for the Data Holder&#x27;s Consumer Data Standard public endpoints | [default to null]
**LogoUri** | **string** | Brand logo URI | [default to null]
**Industries** | **[]string** | The industries the Data Holder Brand belongs to. Please note that the CDR Register entity model is constrained to one industry per brand which is planned to be relaxed in the future. | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) | The date/time that the Data Holder Brand data was last updated in the Register | [default to null]
**Abn** | **string** | Australian Business Number for the organisation | [optional] [default to null]
**Acn** | **string** | Australian Company Number for the organisation | [optional] [default to null]
**Arbn** | **string** | Australian Registered Body Number.  ARBNs are issued to registrable Australian bodies and foreign companies | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

