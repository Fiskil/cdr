# EnergyUsageReadIntervalRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadIntervalLength** | Pointer to **int32** | Read interval length in minutes. Required when _interval-reads_ query parameter equals &#x60;FULL&#x60; or &#x60;MIN_30&#x60;. | [optional] 
**AggregateValue** | **float32** | The aggregate sum of the interval read values. If positive then it means net consumption, if negative it means net export. | 
**IntervalReads** | Pointer to **[]float32** | Array of Interval read values. If positive then it means consumption, if negative it means export. Required when _interval-reads_ query parameter equals &#x60;FULL&#x60; or &#x60;MIN_30&#x60;.&lt;br&gt;Each read value indicates the read for the interval specified by _readIntervalLength_ beginning at midnight of _readStartDate_ (for example 00:00 to 00:30 would be the first reading in a 30 minute Interval). | [optional] 
**ReadQualities** | Pointer to [**[]EnergyUsageReadIntervalReadReadQualities**](EnergyUsageReadIntervalReadReadQualities.md) |  Specifies quality of reads that are not &#x60;ACTUAL&#x60;. For read indices that are not specified, quality is assumed to be &#x60;ACTUAL&#x60;. If not present, all quality of all reads are assumed to be actual. Required when _interval-reads_ query parameter equals &#x60;FULL&#x60; or &#x60;MIN_30&#x60;. | [optional] 

## Methods

### NewEnergyUsageReadIntervalRead

`func NewEnergyUsageReadIntervalRead(aggregateValue float32, ) *EnergyUsageReadIntervalRead`

NewEnergyUsageReadIntervalRead instantiates a new EnergyUsageReadIntervalRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnergyUsageReadIntervalReadWithDefaults

`func NewEnergyUsageReadIntervalReadWithDefaults() *EnergyUsageReadIntervalRead`

NewEnergyUsageReadIntervalReadWithDefaults instantiates a new EnergyUsageReadIntervalRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadIntervalLength

`func (o *EnergyUsageReadIntervalRead) GetReadIntervalLength() int32`

GetReadIntervalLength returns the ReadIntervalLength field if non-nil, zero value otherwise.

### GetReadIntervalLengthOk

`func (o *EnergyUsageReadIntervalRead) GetReadIntervalLengthOk() (*int32, bool)`

GetReadIntervalLengthOk returns a tuple with the ReadIntervalLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIntervalLength

`func (o *EnergyUsageReadIntervalRead) SetReadIntervalLength(v int32)`

SetReadIntervalLength sets ReadIntervalLength field to given value.

### HasReadIntervalLength

`func (o *EnergyUsageReadIntervalRead) HasReadIntervalLength() bool`

HasReadIntervalLength returns a boolean if a field has been set.

### GetAggregateValue

`func (o *EnergyUsageReadIntervalRead) GetAggregateValue() float32`

GetAggregateValue returns the AggregateValue field if non-nil, zero value otherwise.

### GetAggregateValueOk

`func (o *EnergyUsageReadIntervalRead) GetAggregateValueOk() (*float32, bool)`

GetAggregateValueOk returns a tuple with the AggregateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateValue

`func (o *EnergyUsageReadIntervalRead) SetAggregateValue(v float32)`

SetAggregateValue sets AggregateValue field to given value.


### GetIntervalReads

`func (o *EnergyUsageReadIntervalRead) GetIntervalReads() []float32`

GetIntervalReads returns the IntervalReads field if non-nil, zero value otherwise.

### GetIntervalReadsOk

`func (o *EnergyUsageReadIntervalRead) GetIntervalReadsOk() (*[]float32, bool)`

GetIntervalReadsOk returns a tuple with the IntervalReads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalReads

`func (o *EnergyUsageReadIntervalRead) SetIntervalReads(v []float32)`

SetIntervalReads sets IntervalReads field to given value.

### HasIntervalReads

`func (o *EnergyUsageReadIntervalRead) HasIntervalReads() bool`

HasIntervalReads returns a boolean if a field has been set.

### GetReadQualities

`func (o *EnergyUsageReadIntervalRead) GetReadQualities() []EnergyUsageReadIntervalReadReadQualities`

GetReadQualities returns the ReadQualities field if non-nil, zero value otherwise.

### GetReadQualitiesOk

`func (o *EnergyUsageReadIntervalRead) GetReadQualitiesOk() (*[]EnergyUsageReadIntervalReadReadQualities, bool)`

GetReadQualitiesOk returns a tuple with the ReadQualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadQualities

`func (o *EnergyUsageReadIntervalRead) SetReadQualities(v []EnergyUsageReadIntervalReadReadQualities)`

SetReadQualities sets ReadQualities field to given value.

### HasReadQualities

`func (o *EnergyUsageReadIntervalRead) HasReadQualities() bool`

HasReadQualities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


