# RegisterDataRecipient

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegalEntityId** | **string** | Unique id of the Data Recipient Legal Entity issued by the CDR Register. | [default to null]
**LegalEntityName** | **string** | Legal name of the Data Recipient | [default to null]
**AccreditationNumber** | **string** | CDR Register issued human readable unique number given to Data Recipients upon accreditation | [default to null]
**AccreditationLevel** | **string** | Accreditation level of the Data Recipient in the CDR Register | [default to null]
**LogoUri** | **string** | Legal Entity logo URI | [default to null]
**DataRecipientBrands** | [**[]DataRecipientBrandMetaData**](DataRecipientBrandMetaData.md) |  | [optional] [default to null]
**Status** | **string** | Data Recipient status in the CDR Register | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) | The date/time that the Legal Entity was last updated in the CDR Register | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

