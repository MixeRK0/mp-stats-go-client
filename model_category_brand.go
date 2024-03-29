/*
mpstats

MPStats API

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CategoryBrand 
type CategoryBrand struct {
	// Название бренда
	Brand string `json:"brand"`
	// Число товаров
	Items int32 `json:"items"`
	// Число товаров с продажами
	ItemsWithSells int32 `json:"items_with_sells"`
	// Число товаров с продажами в процентах
	ItemsWithSellsPercent float32 `json:"items_with_sells_percent"`
	// Число продавцов
	Sellers int32 `json:"sellers"`
	// Число продавцов с продажами
	SellersWithSells int32 `json:"sellers_with_sells"`
	// Число продавцов с продажами в процентах
	SellersWithSellsPercent float32 `json:"sellers_with_sells_percent"`
	// Число зафиксированных продаж (единицы)
	Sales int32 `json:"sales"`
	// Сумма произведений числа проданных товаров на их стоимость
	Revenue float32 `json:"revenue"`
	// Среднее количество продаж на товарную позицию
	SalesPerItemsAverage float32 `json:"sales_per_items_average"`
	// Среднее количество продаж на товарную позицию с продажами
	SalesPerItemsWithSellsAverage float32 `json:"sales_per_items_with_sells_average"`
	// Средняя выручка за товар
	RevenuePerItemsAverage float32 `json:"revenue_per_items_average"`
	// Средняя выручка за товар с продажами
	RevenuePerItemsWithSellsAverage float32 `json:"revenue_per_items_with_sells_average"`
	// Название бренда
	Name string `json:"name"`
	// Баланс
	Balance float32 `json:"balance"`
	// Средняя стоимость товара
	AvgPrice float32 `json:"avg_price"`
	// Средний рейтинг
	Rating float32 `json:"rating"`
	// Среднее число комментариев
	Comments float32 `json:"comments"`
	// Позиция в рейтинге
	Position int32 `json:"position"`
	// Граф
	Graph []int32 `json:"graph"`
}

// NewCategoryBrand instantiates a new CategoryBrand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryBrand(brand string, items int32, itemsWithSells int32, itemsWithSellsPercent float32, sellers int32, sellersWithSells int32, sellersWithSellsPercent float32, sales int32, revenue float32, salesPerItemsAverage float32, salesPerItemsWithSellsAverage float32, revenuePerItemsAverage float32, revenuePerItemsWithSellsAverage float32, name string, balance float32, avgPrice float32, rating float32, comments float32, position int32, graph []int32) *CategoryBrand {
	this := CategoryBrand{}
	this.Brand = brand
	this.Items = items
	this.ItemsWithSells = itemsWithSells
	this.ItemsWithSellsPercent = itemsWithSellsPercent
	this.Sellers = sellers
	this.SellersWithSells = sellersWithSells
	this.SellersWithSellsPercent = sellersWithSellsPercent
	this.Sales = sales
	this.Revenue = revenue
	this.SalesPerItemsAverage = salesPerItemsAverage
	this.SalesPerItemsWithSellsAverage = salesPerItemsWithSellsAverage
	this.RevenuePerItemsAverage = revenuePerItemsAverage
	this.RevenuePerItemsWithSellsAverage = revenuePerItemsWithSellsAverage
	this.Name = name
	this.Balance = balance
	this.AvgPrice = avgPrice
	this.Rating = rating
	this.Comments = comments
	this.Position = position
	this.Graph = graph
	return &this
}

// NewCategoryBrandWithDefaults instantiates a new CategoryBrand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryBrandWithDefaults() *CategoryBrand {
	this := CategoryBrand{}
	return &this
}

// GetBrand returns the Brand field value
func (o *CategoryBrand) GetBrand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Brand
}

// GetBrandOk returns a tuple with the Brand field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Brand, true
}

// SetBrand sets field value
func (o *CategoryBrand) SetBrand(v string) {
	o.Brand = v
}

// GetItems returns the Items field value
func (o *CategoryBrand) GetItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *CategoryBrand) SetItems(v int32) {
	o.Items = v
}

// GetItemsWithSells returns the ItemsWithSells field value
func (o *CategoryBrand) GetItemsWithSells() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemsWithSells
}

// GetItemsWithSellsOk returns a tuple with the ItemsWithSells field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetItemsWithSellsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsWithSells, true
}

// SetItemsWithSells sets field value
func (o *CategoryBrand) SetItemsWithSells(v int32) {
	o.ItemsWithSells = v
}

// GetItemsWithSellsPercent returns the ItemsWithSellsPercent field value
func (o *CategoryBrand) GetItemsWithSellsPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ItemsWithSellsPercent
}

// GetItemsWithSellsPercentOk returns a tuple with the ItemsWithSellsPercent field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetItemsWithSellsPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsWithSellsPercent, true
}

// SetItemsWithSellsPercent sets field value
func (o *CategoryBrand) SetItemsWithSellsPercent(v float32) {
	o.ItemsWithSellsPercent = v
}

// GetSellers returns the Sellers field value
func (o *CategoryBrand) GetSellers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sellers
}

// GetSellersOk returns a tuple with the Sellers field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetSellersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sellers, true
}

// SetSellers sets field value
func (o *CategoryBrand) SetSellers(v int32) {
	o.Sellers = v
}

// GetSellersWithSells returns the SellersWithSells field value
func (o *CategoryBrand) GetSellersWithSells() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SellersWithSells
}

// GetSellersWithSellsOk returns a tuple with the SellersWithSells field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetSellersWithSellsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellersWithSells, true
}

// SetSellersWithSells sets field value
func (o *CategoryBrand) SetSellersWithSells(v int32) {
	o.SellersWithSells = v
}

// GetSellersWithSellsPercent returns the SellersWithSellsPercent field value
func (o *CategoryBrand) GetSellersWithSellsPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SellersWithSellsPercent
}

// GetSellersWithSellsPercentOk returns a tuple with the SellersWithSellsPercent field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetSellersWithSellsPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellersWithSellsPercent, true
}

// SetSellersWithSellsPercent sets field value
func (o *CategoryBrand) SetSellersWithSellsPercent(v float32) {
	o.SellersWithSellsPercent = v
}

// GetSales returns the Sales field value
func (o *CategoryBrand) GetSales() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sales
}

// GetSalesOk returns a tuple with the Sales field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetSalesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sales, true
}

// SetSales sets field value
func (o *CategoryBrand) SetSales(v int32) {
	o.Sales = v
}

// GetRevenue returns the Revenue field value
func (o *CategoryBrand) GetRevenue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revenue
}

// GetRevenueOk returns a tuple with the Revenue field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetRevenueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revenue, true
}

// SetRevenue sets field value
func (o *CategoryBrand) SetRevenue(v float32) {
	o.Revenue = v
}

// GetSalesPerItemsAverage returns the SalesPerItemsAverage field value
func (o *CategoryBrand) GetSalesPerItemsAverage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SalesPerItemsAverage
}

// GetSalesPerItemsAverageOk returns a tuple with the SalesPerItemsAverage field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetSalesPerItemsAverageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SalesPerItemsAverage, true
}

// SetSalesPerItemsAverage sets field value
func (o *CategoryBrand) SetSalesPerItemsAverage(v float32) {
	o.SalesPerItemsAverage = v
}

// GetSalesPerItemsWithSellsAverage returns the SalesPerItemsWithSellsAverage field value
func (o *CategoryBrand) GetSalesPerItemsWithSellsAverage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SalesPerItemsWithSellsAverage
}

// GetSalesPerItemsWithSellsAverageOk returns a tuple with the SalesPerItemsWithSellsAverage field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetSalesPerItemsWithSellsAverageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SalesPerItemsWithSellsAverage, true
}

// SetSalesPerItemsWithSellsAverage sets field value
func (o *CategoryBrand) SetSalesPerItemsWithSellsAverage(v float32) {
	o.SalesPerItemsWithSellsAverage = v
}

// GetRevenuePerItemsAverage returns the RevenuePerItemsAverage field value
func (o *CategoryBrand) GetRevenuePerItemsAverage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RevenuePerItemsAverage
}

// GetRevenuePerItemsAverageOk returns a tuple with the RevenuePerItemsAverage field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetRevenuePerItemsAverageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevenuePerItemsAverage, true
}

// SetRevenuePerItemsAverage sets field value
func (o *CategoryBrand) SetRevenuePerItemsAverage(v float32) {
	o.RevenuePerItemsAverage = v
}

// GetRevenuePerItemsWithSellsAverage returns the RevenuePerItemsWithSellsAverage field value
func (o *CategoryBrand) GetRevenuePerItemsWithSellsAverage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RevenuePerItemsWithSellsAverage
}

// GetRevenuePerItemsWithSellsAverageOk returns a tuple with the RevenuePerItemsWithSellsAverage field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetRevenuePerItemsWithSellsAverageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevenuePerItemsWithSellsAverage, true
}

// SetRevenuePerItemsWithSellsAverage sets field value
func (o *CategoryBrand) SetRevenuePerItemsWithSellsAverage(v float32) {
	o.RevenuePerItemsWithSellsAverage = v
}

// GetName returns the Name field value
func (o *CategoryBrand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CategoryBrand) SetName(v string) {
	o.Name = v
}

// GetBalance returns the Balance field value
func (o *CategoryBrand) GetBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *CategoryBrand) SetBalance(v float32) {
	o.Balance = v
}

// GetAvgPrice returns the AvgPrice field value
func (o *CategoryBrand) GetAvgPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetAvgPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgPrice, true
}

// SetAvgPrice sets field value
func (o *CategoryBrand) SetAvgPrice(v float32) {
	o.AvgPrice = v
}

// GetRating returns the Rating field value
func (o *CategoryBrand) GetRating() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetRatingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *CategoryBrand) SetRating(v float32) {
	o.Rating = v
}

// GetComments returns the Comments field value
func (o *CategoryBrand) GetComments() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetCommentsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *CategoryBrand) SetComments(v float32) {
	o.Comments = v
}

// GetPosition returns the Position field value
func (o *CategoryBrand) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *CategoryBrand) SetPosition(v int32) {
	o.Position = v
}

// GetGraph returns the Graph field value
func (o *CategoryBrand) GetGraph() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Graph
}

// GetGraphOk returns a tuple with the Graph field value
// and a boolean to check if the value has been set.
func (o *CategoryBrand) GetGraphOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Graph, true
}

// SetGraph sets field value
func (o *CategoryBrand) SetGraph(v []int32) {
	o.Graph = v
}

func (o CategoryBrand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["brand"] = o.Brand
	}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["items_with_sells"] = o.ItemsWithSells
	}
	if true {
		toSerialize["items_with_sells_percent"] = o.ItemsWithSellsPercent
	}
	if true {
		toSerialize["sellers"] = o.Sellers
	}
	if true {
		toSerialize["sellers_with_sells"] = o.SellersWithSells
	}
	if true {
		toSerialize["sellers_with_sells_percent"] = o.SellersWithSellsPercent
	}
	if true {
		toSerialize["sales"] = o.Sales
	}
	if true {
		toSerialize["revenue"] = o.Revenue
	}
	if true {
		toSerialize["sales_per_items_average"] = o.SalesPerItemsAverage
	}
	if true {
		toSerialize["sales_per_items_with_sells_average"] = o.SalesPerItemsWithSellsAverage
	}
	if true {
		toSerialize["revenue_per_items_average"] = o.RevenuePerItemsAverage
	}
	if true {
		toSerialize["revenue_per_items_with_sells_average"] = o.RevenuePerItemsWithSellsAverage
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if true {
		toSerialize["avg_price"] = o.AvgPrice
	}
	if true {
		toSerialize["rating"] = o.Rating
	}
	if true {
		toSerialize["comments"] = o.Comments
	}
	if true {
		toSerialize["position"] = o.Position
	}
	if true {
		toSerialize["graph"] = o.Graph
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryBrand struct {
	value *CategoryBrand
	isSet bool
}

func (v NullableCategoryBrand) Get() *CategoryBrand {
	return v.value
}

func (v *NullableCategoryBrand) Set(val *CategoryBrand) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryBrand) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryBrand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryBrand(val *CategoryBrand) *NullableCategoryBrand {
	return &NullableCategoryBrand{value: val, isSet: true}
}

func (v NullableCategoryBrand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryBrand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


