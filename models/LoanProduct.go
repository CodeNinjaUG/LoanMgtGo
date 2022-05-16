package models

import (
	"gorm.io/gorm"
)

type Fund string

const (
	Donorfund Fund = "Donorfund"
	Clientfund    Fund = "Clientfund"
)

type AutoDisburse string

 const (
   Yes AutoDisburse = "Yes"
   No  AutoDisburse = "No"
)


type AccountingRule string

const (
	None AccountingRule = "None"
	Cash AccountingRule = "Cash"
)

type InterestMethodology string

const (
	Flat             InterestMethodology = "Flat"
	DecliningBalance InterestMethodology = "DecliningBalance"
)

type AmortizationMethod string

const (
	EqualInstallments      AmortizationMethod = "EqualInstallments"
	EqualPrincipalPayments AmortizationMethod = "EqualPrincipalPayments"
)


type RepaymentType string

const (
	Days   RepaymentType = "Days"
	Weeks  RepaymentType = "Weeks"
	Months RepaymentType = "Months"
)

type Per string

const (
	Month Per = "Month"
	Year  Per = "Year"
)

type LoanProduct struct {
	gorm.Model
	Name                              string   `json:"name"`
	ShortName                         string   `json:"shortname"`
	Description                       string   `json:"description"`
	Fund                              Fund     `json:"fund"`
	Currency                          Currency `json:"currency"`
	DecimalPlaces                     float64  `json:"decimal"`
	DefaultPrincipal                  float64  `json:"default_principal"`
	MinimumPrincipal                  float64  `json:"minimum_princial"`
	MaximumPrincipal                  float64  `json:"maximum_principal"`
	DefaultLoanTerm                   float64  `json:"default_loan_term"`
	MinimumLoanTerm                   float64  `json:"minimum_loan_term"`
	MaximumLoanTerm                   float64  `json:"maximum_loan_term"`
	RepaymentFrequency                float64  `json:"repayment_frequency"`
	Type                              RepaymentType `json:"repayment_type"`
	DefaultInterestRate               float64  `json:"default_interest_rate"`
	MinimumInterestRate               float64  `json:"minimum_interest_rate"`
	MaximumInterestRate               float64  `json:"maximum_interest_rate"`
	Per                               Per `json:"per"`
	GraceOnPrincipalPayment           float64  `json:"grace_on_principal_payment"`
	GraceOnInterestPayment            float64  `json:"grace_on_interest_payment"`
	GraceOnInterestCharged            float64  `json:"grace_on_interest_charged"`
	InterestMethodology               InterestMethodology `json:"interest_methodology"`
	AmortizationMethod                AmortizationMethod  `json:"amortization_method"`
	LoanTransactionProcessingStrategy string  `json:"loan_transaction_processing_strategy"`
	LoanCharge                        []LoanCharge `gorm:"many2many:loanproduct_loancharges;"`
	CreditChecks                      string  `json:"credit_checks"`
	AccountingRule                    AccountingRule `json:"accounting_rule"`
	AutoDisburse                      AutoDisburse  `json:"auto_disburse"`
	Status                            Active `json:"product_status"`
}




//create a loan product 
func CreateLoanProduct(db *gorm.DB , LoanProduct *LoanProduct)(err error){
	err = db.Create(LoanProduct).Error
	if err!=nil{
		return err
	}
	return nil
}
//get loan products

func GetLoanProducts(db *gorm.DB,LoanProduct *[]LoanProduct)(err error){
	err = db.Find(LoanProduct).Error
	if err!=nil{
      return nil
	}
	return nil
}

//get loan product by ID
func GetLoanProduct(db *gorm.DB, LoanProduct *LoanProduct, id string)(err error){
	err = db.Where("id = ?", id).First(LoanProduct).Error
	if err!=nil{
       return err
	}
	return nil
}

//Update loan product 
// still pending research
func UpdateLoanProduct(db *gorm.DB,LoanProduct *LoanProduct,id string)(err error){
	db.Where("id =?",id).Save(LoanProduct)
	//db.Save(LoanProduct)
	return nil
}

//Delete loan Product
func DeleteLoanProduct(db *gorm.DB,LoanProduct *LoanProduct,id string)(err error){
   db.Where("id =?",id).Delete(LoanProduct)
   return nil
}