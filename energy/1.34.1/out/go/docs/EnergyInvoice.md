# EnergyInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique identifier for the account. | 
**InvoiceNumber** | **string** | The number assigned to this invoice by the energy Retailer. | 
**IssueDate** | **string** | The date that the invoice was actually issued (as opposed to generated or calculated). | 
**DueDate** | Pointer to **string** | The date that the invoice is due to be paid. | [optional] 
**Period** | Pointer to [**EnergyInvoicePeriod**](EnergyInvoicePeriod.md) |  | [optional] 
**InvoiceAmount** | Pointer to **string** | The net amount due for this invoice regardless of previous balance. | [optional] 
**GstAmount** | Pointer to **string** | The total GST amount for this invoice. If absent then zero is assumed. | [optional] [default to "0.00"]
**PayOnTimeDiscount** | Pointer to [**EnergyInvoicePayOnTimeDiscount**](EnergyInvoicePayOnTimeDiscount.md) |  | [optional] 
**BalanceAtIssue** | **string** | The account balance at the time the invoice was issued. | 
**ServicePoints** | **[]string** | Array of service point IDs to which this invoice applies. May be empty if the invoice contains no electricity usage related charges. | 
**Gas** | Pointer to [**EnergyInvoiceGasUsageCharges**](EnergyInvoiceGasUsageCharges.md) | Object containing charges and credits related to gas usage. | [optional] 
**Electricity** | Pointer to [**EnergyInvoiceElectricityUsageCharges**](EnergyInvoiceElectricityUsageCharges.md) | Object containing charges and credits related to electricity usage. | [optional] 
**AccountCharges** | Pointer to [**EnergyInvoiceAccountCharges**](EnergyInvoiceAccountCharges.md) |  | [optional] 
**PaymentStatus** | **string** | Indicator of the payment status for the invoice. | 

## Methods

### NewEnergyInvoice

`func NewEnergyInvoice(accountId string, invoiceNumber string, issueDate string, balanceAtIssue string, servicePoints []string, paymentStatus string, ) *EnergyInvoice`

NewEnergyInvoice instantiates a new EnergyInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyInvoiceWithDefaults

`func NewEnergyInvoiceWithDefaults() *EnergyInvoice`

NewEnergyInvoiceWithDefaults instantiates a new EnergyInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *EnergyInvoice) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EnergyInvoice) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EnergyInvoice) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetInvoiceNumber

`func (o *EnergyInvoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *EnergyInvoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *EnergyInvoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetIssueDate

`func (o *EnergyInvoice) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *EnergyInvoice) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *EnergyInvoice) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetDueDate

`func (o *EnergyInvoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *EnergyInvoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *EnergyInvoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *EnergyInvoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetPeriod

`func (o *EnergyInvoice) GetPeriod() EnergyInvoicePeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EnergyInvoice) GetPeriodOk() (*EnergyInvoicePeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EnergyInvoice) SetPeriod(v EnergyInvoicePeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EnergyInvoice) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetInvoiceAmount

`func (o *EnergyInvoice) GetInvoiceAmount() string`

GetInvoiceAmount returns the InvoiceAmount field if non-nil, zero value otherwise.

### GetInvoiceAmountOk

`func (o *EnergyInvoice) GetInvoiceAmountOk() (*string, bool)`

GetInvoiceAmountOk returns a tuple with the InvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAmount

`func (o *EnergyInvoice) SetInvoiceAmount(v string)`

SetInvoiceAmount sets InvoiceAmount field to given value.

### HasInvoiceAmount

`func (o *EnergyInvoice) HasInvoiceAmount() bool`

HasInvoiceAmount returns a boolean if a field has been set.

### GetGstAmount

`func (o *EnergyInvoice) GetGstAmount() string`

GetGstAmount returns the GstAmount field if non-nil, zero value otherwise.

### GetGstAmountOk

`func (o *EnergyInvoice) GetGstAmountOk() (*string, bool)`

GetGstAmountOk returns a tuple with the GstAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGstAmount

`func (o *EnergyInvoice) SetGstAmount(v string)`

SetGstAmount sets GstAmount field to given value.

### HasGstAmount

`func (o *EnergyInvoice) HasGstAmount() bool`

HasGstAmount returns a boolean if a field has been set.

### GetPayOnTimeDiscount

`func (o *EnergyInvoice) GetPayOnTimeDiscount() EnergyInvoicePayOnTimeDiscount`

GetPayOnTimeDiscount returns the PayOnTimeDiscount field if non-nil, zero value otherwise.

### GetPayOnTimeDiscountOk

`func (o *EnergyInvoice) GetPayOnTimeDiscountOk() (*EnergyInvoicePayOnTimeDiscount, bool)`

GetPayOnTimeDiscountOk returns a tuple with the PayOnTimeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOnTimeDiscount

`func (o *EnergyInvoice) SetPayOnTimeDiscount(v EnergyInvoicePayOnTimeDiscount)`

SetPayOnTimeDiscount sets PayOnTimeDiscount field to given value.

### HasPayOnTimeDiscount

`func (o *EnergyInvoice) HasPayOnTimeDiscount() bool`

HasPayOnTimeDiscount returns a boolean if a field has been set.

### GetBalanceAtIssue

`func (o *EnergyInvoice) GetBalanceAtIssue() string`

GetBalanceAtIssue returns the BalanceAtIssue field if non-nil, zero value otherwise.

### GetBalanceAtIssueOk

`func (o *EnergyInvoice) GetBalanceAtIssueOk() (*string, bool)`

GetBalanceAtIssueOk returns a tuple with the BalanceAtIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAtIssue

`func (o *EnergyInvoice) SetBalanceAtIssue(v string)`

SetBalanceAtIssue sets BalanceAtIssue field to given value.


### GetServicePoints

`func (o *EnergyInvoice) GetServicePoints() []string`

GetServicePoints returns the ServicePoints field if non-nil, zero value otherwise.

### GetServicePointsOk

`func (o *EnergyInvoice) GetServicePointsOk() (*[]string, bool)`

GetServicePointsOk returns a tuple with the ServicePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoints

`func (o *EnergyInvoice) SetServicePoints(v []string)`

SetServicePoints sets ServicePoints field to given value.


### GetGas

`func (o *EnergyInvoice) GetGas() EnergyInvoiceGasUsageCharges`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *EnergyInvoice) GetGasOk() (*EnergyInvoiceGasUsageCharges, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *EnergyInvoice) SetGas(v EnergyInvoiceGasUsageCharges)`

SetGas sets Gas field to given value.

### HasGas

`func (o *EnergyInvoice) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetElectricity

`func (o *EnergyInvoice) GetElectricity() EnergyInvoiceElectricityUsageCharges`

GetElectricity returns the Electricity field if non-nil, zero value otherwise.

### GetElectricityOk

`func (o *EnergyInvoice) GetElectricityOk() (*EnergyInvoiceElectricityUsageCharges, bool)`

GetElectricityOk returns a tuple with the Electricity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectricity

`func (o *EnergyInvoice) SetElectricity(v EnergyInvoiceElectricityUsageCharges)`

SetElectricity sets Electricity field to given value.

### HasElectricity

`func (o *EnergyInvoice) HasElectricity() bool`

HasElectricity returns a boolean if a field has been set.

### GetAccountCharges

`func (o *EnergyInvoice) GetAccountCharges() EnergyInvoiceAccountCharges`

GetAccountCharges returns the AccountCharges field if non-nil, zero value otherwise.

### GetAccountChargesOk

`func (o *EnergyInvoice) GetAccountChargesOk() (*EnergyInvoiceAccountCharges, bool)`

GetAccountChargesOk returns a tuple with the AccountCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCharges

`func (o *EnergyInvoice) SetAccountCharges(v EnergyInvoiceAccountCharges)`

SetAccountCharges sets AccountCharges field to given value.

### HasAccountCharges

`func (o *EnergyInvoice) HasAccountCharges() bool`

HasAccountCharges returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *EnergyInvoice) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *EnergyInvoice) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *EnergyInvoice) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


