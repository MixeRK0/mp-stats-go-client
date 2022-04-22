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

// CategorySubcategory 
type CategorySubcategory struct {
	// Вложенная рубрика
	Name string `json:"name"`
	// Число товаров
	Items int32 `json:"items"`
	// Число товаров с продажами
	ItemsWithSells int32 `json:"items_with_sells"`
	// Число брендов
	Brands int32 `json:"brands"`
	// Число брендов с продажами
	BrandsWithSells int32 `json:"brands_with_sells"`
	// Число продавцов
	Sellers *int32 `json:"sellers,omitempty"`
	// Число продавцов с продажами
	SellersWithSells *int32 `json:"sellers_with_sells,omitempty"`
	// Число зафиксированных продаж (единицы)
	Sales int32 `json:"sales"`
	// Сумма произведений числа проданных товаров на их стоимость
	Revenue float32 `json:"revenue"`
	// Средняя стоимость товара
	AvgPrice float32 `json:"avg_price"`
	// Среднее число комментариев
	Comments float32 `json:"comments"`
	// Средний рейтинг
	Rating float32 `json:"rating"`
	// Число товаров с продажами в процентах
	ItemsWithSellsPercent float32 `json:"items_with_sells_percent"`
	// Число брендов с продажами в процентах
	BrandsWithSellsPercent float32 `json:"brands_with_sells_percent"`
	// Число продавцов с продажами в процентах
	SellersWithSellsPercent float32 `json:"sellers_with_sells_percent"`
	// Среднее количество продаж на товар
	SalesPerItemsAverage float32 `json:"sales_per_items_average"`
	// Среднее количество продаж на товар с продажами
	SalesPerItemsWithSellsAverage float32 `json:"sales_per_items_with_sells_average"`
	// Среднее выручка на товар
	RevenuePerItemsAverage float32 `json:"revenue_per_items_average"`
	// Среднее выручка на товар с продажами
	RevenuePerItemsWithSellsAverage float32 `json:"revenue_per_items_with_sells_average"`
}

// NewCategorySubcategory instantiates a new CategorySubcategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategorySubcategory(name string, items int32, itemsWithSells int32, brands int32, brandsWithSells int32, sales int32, revenue float32, avgPrice float32, comments float32, rating float32, itemsWithSellsPercent float32, brandsWithSellsPercent float32, sellersWithSellsPercent float32, salesPerItemsAverage float32, salesPerItemsWithSellsAverage float32, revenuePerItemsAverage float32, revenuePerItemsWithSellsAverage float32) *CategorySubcategory {
	this := CategorySubcategory{}
	this.Name = name
	this.Items = items
	this.ItemsWithSells = itemsWithSells
	this.Brands = brands
	this.BrandsWithSells = brandsWithSells
	this.Sales = sales
	this.Revenue = revenue
	this.AvgPrice = avgPrice
	this.Comments = comments
	this.Rating = rating
	this.ItemsWithSellsPercent = itemsWithSellsPercent
	this.BrandsWithSellsPercent = brandsWithSellsPercent
	this.SellersWithSellsPercent = sellersWithSellsPercent
	this.SalesPerItemsAverage = salesPerItemsAverage
	this.SalesPerItemsWithSellsAverage = salesPerItemsWithSellsAverage
	this.RevenuePerItemsAverage = revenuePerItemsAverage
	this.RevenuePerItemsWithSellsAverage = revenuePerItemsWithSellsAverage
	return &this
}

// NewCategorySubcategoryWithDefaults instantiates a new CategorySubcategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategorySubcategoryWithDefaults() *CategorySubcategory {
	this := CategorySubcategory{}
	return &this
}

// GetName returns the Name field value
func (o *CategorySubcategory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CategorySubcategory) SetName(v string) {
	o.Name = v
}

// GetItems returns the Items field value
func (o *CategorySubcategory) GetItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *CategorySubcategory) SetItems(v int32) {
	o.Items = v
}

// GetItemsWithSells returns the ItemsWithSells field value
func (o *CategorySubcategory) GetItemsWithSells() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemsWithSells
}

// GetItemsWithSellsOk returns a tuple with the ItemsWithSells field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetItemsWithSellsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsWithSells, true
}

// SetItemsWithSells sets field value
func (o *CategorySubcategory) SetItemsWithSells(v int32) {
	o.ItemsWithSells = v
}

// GetBrands returns the Brands field value
func (o *CategorySubcategory) GetBrands() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetBrandsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Brands, true
}

// SetBrands sets field value
func (o *CategorySubcategory) SetBrands(v int32) {
	o.Brands = v
}

// GetBrandsWithSells returns the BrandsWithSells field value
func (o *CategorySubcategory) GetBrandsWithSells() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BrandsWithSells
}

// GetBrandsWithSellsOk returns a tuple with the BrandsWithSells field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetBrandsWithSellsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandsWithSells, true
}

// SetBrandsWithSells sets field value
func (o *CategorySubcategory) SetBrandsWithSells(v int32) {
	o.BrandsWithSells = v
}

// GetSellers returns the Sellers field value if set, zero value otherwise.
func (o *CategorySubcategory) GetSellers() int32 {
	if o == nil || o.Sellers == nil {
		var ret int32
		return ret
	}
	return *o.Sellers
}

// GetSellersOk returns a tuple with the Sellers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetSellersOk() (*int32, bool) {
	if o == nil || o.Sellers == nil {
		return nil, false
	}
	return o.Sellers, true
}

// HasSellers returns a boolean if a field has been set.
func (o *CategorySubcategory) HasSellers() bool {
	if o != nil && o.Sellers != nil {
		return true
	}

	return false
}

// SetSellers gets a reference to the given int32 and assigns it to the Sellers field.
func (o *CategorySubcategory) SetSellers(v int32) {
	o.Sellers = &v
}

// GetSellersWithSells returns the SellersWithSells field value if set, zero value otherwise.
func (o *CategorySubcategory) GetSellersWithSells() int32 {
	if o == nil || o.SellersWithSells == nil {
		var ret int32
		return ret
	}
	return *o.SellersWithSells
}

// GetSellersWithSellsOk returns a tuple with the SellersWithSells field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetSellersWithSellsOk() (*int32, bool) {
	if o == nil || o.SellersWithSells == nil {
		return nil, false
	}
	return o.SellersWithSells, true
}

// HasSellersWithSells returns a boolean if a field has been set.
func (o *CategorySubcategory) HasSellersWithSells() bool {
	if o != nil && o.SellersWithSells != nil {
		return true
	}

	return false
}

// SetSellersWithSells gets a reference to the given int32 and assigns it to the SellersWithSells field.
func (o *CategorySubcategory) SetSellersWithSells(v int32) {
	o.SellersWithSells = &v
}

// GetSales returns the Sales field value
func (o *CategorySubcategory) GetSales() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sales
}

// GetSalesOk returns a tuple with the Sales field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetSalesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sales, true
}

// SetSales sets field value
func (o *CategorySubcategory) SetSales(v int32) {
	o.Sales = v
}

// GetRevenue returns the Revenue field value
func (o *CategorySubcategory) GetRevenue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revenue
}

// GetRevenueOk returns a tuple with the Revenue field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetRevenueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revenue, true
}

// SetRevenue sets field value
func (o *CategorySubcategory) SetRevenue(v float32) {
	o.Revenue = v
}

// GetAvgPrice returns the AvgPrice field value
func (o *CategorySubcategory) GetAvgPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetAvgPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgPrice, true
}

// SetAvgPrice sets field value
func (o *CategorySubcategory) SetAvgPrice(v float32) {
	o.AvgPrice = v
}

// GetComments returns the Comments field value
func (o *CategorySubcategory) GetComments() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetCommentsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *CategorySubcategory) SetComments(v float32) {
	o.Comments = v
}

// GetRating returns the Rating field value
func (o *CategorySubcategory) GetRating() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetRatingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *CategorySubcategory) SetRating(v float32) {
	o.Rating = v
}

// GetItemsWithSellsPercent returns the ItemsWithSellsPercent field value
func (o *CategorySubcategory) GetItemsWithSellsPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ItemsWithSellsPercent
}

// GetItemsWithSellsPercentOk returns a tuple with the ItemsWithSellsPercent field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetItemsWithSellsPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsWithSellsPercent, true
}

// SetItemsWithSellsPercent sets field value
func (o *CategorySubcategory) SetItemsWithSellsPercent(v float32) {
	o.ItemsWithSellsPercent = v
}

// GetBrandsWithSellsPercent returns the BrandsWithSellsPercent field value
func (o *CategorySubcategory) GetBrandsWithSellsPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BrandsWithSellsPercent
}

// GetBrandsWithSellsPercentOk returns a tuple with the BrandsWithSellsPercent field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetBrandsWithSellsPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandsWithSellsPercent, true
}

// SetBrandsWithSellsPercent sets field value
func (o *CategorySubcategory) SetBrandsWithSellsPercent(v float32) {
	o.BrandsWithSellsPercent = v
}

// GetSellersWithSellsPercent returns the SellersWithSellsPercent field value
func (o *CategorySubcategory) GetSellersWithSellsPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SellersWithSellsPercent
}

// GetSellersWithSellsPercentOk returns a tuple with the SellersWithSellsPercent field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetSellersWithSellsPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellersWithSellsPercent, true
}

// SetSellersWithSellsPercent sets field value
func (o *CategorySubcategory) SetSellersWithSellsPercent(v float32) {
	o.SellersWithSellsPercent = v
}

// GetSalesPerItemsAverage returns the SalesPerItemsAverage field value
func (o *CategorySubcategory) GetSalesPerItemsAverage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SalesPerItemsAverage
}

// GetSalesPerItemsAverageOk returns a tuple with the SalesPerItemsAverage field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetSalesPerItemsAverageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SalesPerItemsAverage, true
}

// SetSalesPerItemsAverage sets field value
func (o *CategorySubcategory) SetSalesPerItemsAverage(v float32) {
	o.SalesPerItemsAverage = v
}

// GetSalesPerItemsWithSellsAverage returns the SalesPerItemsWithSellsAverage field value
func (o *CategorySubcategory) GetSalesPerItemsWithSellsAverage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SalesPerItemsWithSellsAverage
}

// GetSalesPerItemsWithSellsAverageOk returns a tuple with the SalesPerItemsWithSellsAverage field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetSalesPerItemsWithSellsAverageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SalesPerItemsWithSellsAverage, true
}

// SetSalesPerItemsWithSellsAverage sets field value
func (o *CategorySubcategory) SetSalesPerItemsWithSellsAverage(v float32) {
	o.SalesPerItemsWithSellsAverage = v
}

// GetRevenuePerItemsAverage returns the RevenuePerItemsAverage field value
func (o *CategorySubcategory) GetRevenuePerItemsAverage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RevenuePerItemsAverage
}

// GetRevenuePerItemsAverageOk returns a tuple with the RevenuePerItemsAverage field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetRevenuePerItemsAverageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevenuePerItemsAverage, true
}

// SetRevenuePerItemsAverage sets field value
func (o *CategorySubcategory) SetRevenuePerItemsAverage(v float32) {
	o.RevenuePerItemsAverage = v
}

// GetRevenuePerItemsWithSellsAverage returns the RevenuePerItemsWithSellsAverage field value
func (o *CategorySubcategory) GetRevenuePerItemsWithSellsAverage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RevenuePerItemsWithSellsAverage
}

// GetRevenuePerItemsWithSellsAverageOk returns a tuple with the RevenuePerItemsWithSellsAverage field value
// and a boolean to check if the value has been set.
func (o *CategorySubcategory) GetRevenuePerItemsWithSellsAverageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevenuePerItemsWithSellsAverage, true
}

// SetRevenuePerItemsWithSellsAverage sets field value
func (o *CategorySubcategory) SetRevenuePerItemsWithSellsAverage(v float32) {
	o.RevenuePerItemsWithSellsAverage = v
}

func (o CategorySubcategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["items_with_sells"] = o.ItemsWithSells
	}
	if true {
		toSerialize["brands"] = o.Brands
	}
	if true {
		toSerialize["brands_with_sells"] = o.BrandsWithSells
	}
	if o.Sellers != nil {
		toSerialize["sellers"] = o.Sellers
	}
	if o.SellersWithSells != nil {
		toSerialize["sellers_with_sells"] = o.SellersWithSells
	}
	if true {
		toSerialize["sales"] = o.Sales
	}
	if true {
		toSerialize["revenue"] = o.Revenue
	}
	if true {
		toSerialize["avg_price"] = o.AvgPrice
	}
	if true {
		toSerialize["comments"] = o.Comments
	}
	if true {
		toSerialize["rating"] = o.Rating
	}
	if true {
		toSerialize["items_with_sells_percent"] = o.ItemsWithSellsPercent
	}
	if true {
		toSerialize["brands_with_sells_percent"] = o.BrandsWithSellsPercent
	}
	if true {
		toSerialize["sellers_with_sells_percent"] = o.SellersWithSellsPercent
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
	return json.Marshal(toSerialize)
}

type NullableCategorySubcategory struct {
	value *CategorySubcategory
	isSet bool
}

func (v NullableCategorySubcategory) Get() *CategorySubcategory {
	return v.value
}

func (v *NullableCategorySubcategory) Set(val *CategorySubcategory) {
	v.value = val
	v.isSet = true
}

func (v NullableCategorySubcategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCategorySubcategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategorySubcategory(val *CategorySubcategory) *NullableCategorySubcategory {
	return &NullableCategorySubcategory{value: val, isSet: true}
}

func (v NullableCategorySubcategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategorySubcategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


