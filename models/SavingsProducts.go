package models

import "gorm.io/gorm"

type Category string

const (
	Voluntary  Category = "Voluntary"
	Compulsory Category = "Compulsory"
)

type AutoCreate string

const (
	True  AutoCreate = "True"
	False AutoCreate = "False"
)

type CompoundingPeriod string

const (
	Daily    CompoundingPeriod = "Daily"
	Weekly   CompoundingPeriod = "Weekly"
	Monthly  CompoundingPeriod = "Monthly"
	BiAnnual CompoundingPeriod = "BiAnnual"
	Annual   CompoundingPeriod = "Annual"
)

type InterestPostingPeriodType string

const (
	Monthly_  InterestPostingPeriodType = "Monthly"
	BiAnnual_ InterestPostingPeriodType = "BiAnnual"
	Annually_ InterestPostingPeriodType = "Annual"
)

type InterestCalculationType string

const (
	DailyBalance   InterestCalculationType = "DailyBalance"
	AverageBalance InterestCalculationType = "AverageBalance"
)

type LockinPeriodFrequencyType string

const(
	Days_   LockinPeriodFrequencyType = "Days"
	Weeks_  LockinPeriodFrequencyType = "Weeks"
	Months_ LockinPeriodFrequencyType = "Months"
	Years_  LockinPeriodFrequencyType = "Years" 
)

type AllowOverDraft string 

const (
	No__ AllowOverDraft = "No"
	Yes__ AllowOverDraft = "Yes"
)

type SavingsProduct struct {
	Name                      string                    `json:"savings_product_name" binding:"required"`
	ShortName                 string                    `json:"savings_product_short_name" binding:"required"`
	Description               string                    `json:"savings_product_description" binding:"required"`
	Currency                  Currency                  `json:"savings_product_currency" binding:"required"`
	Interest                  string                    `json:"savings_product_interest" binding:"required"`
	Category                  Category                  `json:"savings_product_category" binding:"required"`
	AutoCreate                AutoCreate                `json:"auto_create_savings_product" binding:"required"`
	CompoundingPeriod         CompoundingPeriod         `json:"compounding_period_savings_product" binding:"required"`
	InterestPostingPeriodType InterestPostingPeriodType `json:"interest_post_type" binding:"required"`
	InterestCalculationType   InterestCalculationType   `json:"interest_calculation_type" binding:"required"`
	LockinPeriodFrequency     int                       `json:"lockin_period_frequency" binding:"required"`
	LockinPeriodFrequencyType LockinPeriodFrequencyType `json:"lockin_period_frequency_type" binding:"required"`
	AutomaticOpeningBalance  float64                    `json:"automatic_opening_balance" binding:"required"`
	MinimumBalanceInterestCalculation float64           `json:"MinimumBalanceInterestCalculation" binding:"required"`
	AllowOverDraft AllowOverDraft                       `json:"allow_over_draft"  binding:"required"`
	AccountingRule   AccountingRule                     `json:"accounting_rule" binding:"required"`
	Active Active                                       `json:"ative" binding:"required"`
}



//create a saving product
func CreateSavingProduct(db *gorm.DB, SavingsProduct *SavingsProduct)(err error){
	err = db.Create(SavingsProduct).Error
	if err!=nil{
		return err
	}
	return nil
}

//getting saving products
func GetSavingProducts(db *gorm.DB, SavingsProduct *[]SavingsProduct)(err error){
	err = db.Find(SavingsProduct).Error
	if err!=nil{
		return err
	}
	return nil
}

//getting saving product
func GetSavingProduct(db *gorm.DB, SavingsProduct *SavingsProduct, id string)(err error){
    err = db.Where("id = ?",id).First(SavingsProduct).Error
	if err!=nil{
		return err
	}
	return nil
}
//updating a saving product 
func UpdateSavingProduct(db *gorm.DB, SavingsProduct *SavingsProduct, id string)(err error){
      db.Save(SavingsProduct)
	  return nil
}
//delete saving product 
func DeleteSavingProduct(db *gorm.DB, SavingsProduct *SavingsProduct, id string)(err error){
	db.Where("id=?",id).Delete(SavingsProduct)
	return nil
}
