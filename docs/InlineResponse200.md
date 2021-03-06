# InlineResponse200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartRow** | **int32** | Номер строки начала получения данных | 
**EndRow** | **int32** | Номер строки конца получения данных | 
**FilterModel** | **interface{}** |  | 
**SortModel** | [**[]SortModelItem**](SortModelItem.md) |  | 
**Total** | **int32** | Кол-во строк в результирующем запросе без учета пагинации | 
**Data** | [**[]SearchItemsElement**](SearchItemsElement.md) | Массив данных | 

## Methods

### NewInlineResponse200

`func NewInlineResponse200(startRow int32, endRow int32, filterModel interface{}, sortModel []SortModelItem, total int32, data []SearchItemsElement, ) *InlineResponse200`

NewInlineResponse200 instantiates a new InlineResponse200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200WithDefaults

`func NewInlineResponse200WithDefaults() *InlineResponse200`

NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartRow

`func (o *InlineResponse200) GetStartRow() int32`

GetStartRow returns the StartRow field if non-nil, zero value otherwise.

### GetStartRowOk

`func (o *InlineResponse200) GetStartRowOk() (*int32, bool)`

GetStartRowOk returns a tuple with the StartRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartRow

`func (o *InlineResponse200) SetStartRow(v int32)`

SetStartRow sets StartRow field to given value.


### GetEndRow

`func (o *InlineResponse200) GetEndRow() int32`

GetEndRow returns the EndRow field if non-nil, zero value otherwise.

### GetEndRowOk

`func (o *InlineResponse200) GetEndRowOk() (*int32, bool)`

GetEndRowOk returns a tuple with the EndRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndRow

`func (o *InlineResponse200) SetEndRow(v int32)`

SetEndRow sets EndRow field to given value.


### GetFilterModel

`func (o *InlineResponse200) GetFilterModel() interface{}`

GetFilterModel returns the FilterModel field if non-nil, zero value otherwise.

### GetFilterModelOk

`func (o *InlineResponse200) GetFilterModelOk() (*interface{}, bool)`

GetFilterModelOk returns a tuple with the FilterModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterModel

`func (o *InlineResponse200) SetFilterModel(v interface{})`

SetFilterModel sets FilterModel field to given value.


### SetFilterModelNil

`func (o *InlineResponse200) SetFilterModelNil(b bool)`

 SetFilterModelNil sets the value for FilterModel to be an explicit nil

### UnsetFilterModel
`func (o *InlineResponse200) UnsetFilterModel()`

UnsetFilterModel ensures that no value is present for FilterModel, not even an explicit nil
### GetSortModel

`func (o *InlineResponse200) GetSortModel() []SortModelItem`

GetSortModel returns the SortModel field if non-nil, zero value otherwise.

### GetSortModelOk

`func (o *InlineResponse200) GetSortModelOk() (*[]SortModelItem, bool)`

GetSortModelOk returns a tuple with the SortModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortModel

`func (o *InlineResponse200) SetSortModel(v []SortModelItem)`

SetSortModel sets SortModel field to given value.


### GetTotal

`func (o *InlineResponse200) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetData

`func (o *InlineResponse200) GetData() []SearchItemsElement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse200) GetDataOk() (*[]SearchItemsElement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse200) SetData(v []SearchItemsElement)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


