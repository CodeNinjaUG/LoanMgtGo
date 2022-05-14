package models

import (
	"gorm.io/gorm"
)

type Fund string

const (
	Donorfund Fund = "Donorfund"
	Client    Fund = "Client"
)

type AutoDisburse string

 const (
   Yes AutoDisburse = "Yes"
   No  AutoDisburse = "No"
)

type Active string 

const (
	InActive Active = "InActive"
	Activated  Active = "Activated"
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

type Currency string

const (
	UGX Currency = "UGX"
	KES Currency = "KES"
	USD Currency = "USD"
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

type ChargeType string

const (
	Disbursement ChargeType = "Disbursement"
    SpecificDueDate ChargeType = "Specific Due Date"
	InstallmentFees ChargeType = "Installment Fees"
	OverDueInstallmentFees ChargeType = "OverDue Installment Fees"
    DisbursementPaidWithRepayment ChargeType = "Disbursement Paid With Repayment"
	LoanSchedulingFee ChargeType = "LoanScheduling Fee"
	OverDueOnLoanMaturity ChargeType =  "OverDue On LoanMaturity"
	LastInstallmentFee ChargeType = "Last Installment Fee"
)

type ChargeOption string

const (
	FlatCharge ChargeOption = "Flat Charge"
	PrincipalDueOnInstallment ChargeOption = "Principal Due On Installment"
	PrincipalInterestDueOnInstallment ChargeOption = "Principal Interest Due On Installment"
	InterestDueOnInstallment ChargeOption = "Interest Due On Installment"
	TotalOustandingLoanPrincipal ChargeOption = "Total Oustanding Loan Principal"
	PercentageOfOriginalLoanPrincipalPerInstallment ChargeOption = "Percentage Of Original Loan Principal Per Installment"
)


type Penalty string

const (
	Nop  Penalty = "Nop"
	Yep  Penalty = "Yep"
)



type LoanProduct struct {
	gorm.Model
	Name                              string   `json:"name"`
	ShortName                         string   `json:"shortname"`
	Description                       string   `json:"description"`
	Fund                              Fund     `json:"fund"`
	Currency                          Currency `json:"currency"`
	DecimalPlaces                     float64
	DefaultPrincipal                  float64
	MinimumPrincipal                  float64
	MaximumPrincipal                  float64
	DefaultLoanTerm                   float64
	MinimumLoanTerm                   float64
	MaximumLoanTerm                   float64
	RepaymentFrequency                float64
	Type                              RepaymentType `json:"repayment_type"`
	DefaultInterestRate               float64
	MinimumInterestRate               float64
	MaximumInterestRate               float64
	Per                               Per `json:"per"`
	GraceOnPrincipalPayment           float64
	GraceOnInterestPayment            float64
	GraceOnInterestCharged            float64
	InterestMethodology               InterestMethodology `json:"interest_methodology"`
	AmortizationMethod                AmortizationMethod  `json:"amortization_method"`
	LoanTransactionProcessingStrategy string
	LoanCharge                        []LoanCharge `gorm:"many2many:loanproduct_loancharges;"`
	CreditChecks                      string
	AccountingRule                    AccountingRule `json:"accounting_rule"`
	AutoDisburse                      AutoDisburse  `json:"auto_disburse"`
	Status                            Active `json:"product_status"`
}

type LoanCharge struct {
	gorm.Model
	Name         string `json:"name"`
	ChargeType   ChargeType `json:"charge_type"`
	Amount       float64     `json:"amount"`
	ChargeOption ChargeOption `json:"charge_option"`
	Currency     Currency `json:"currency"`
	Penalty      Penalty  `json:"penalty"`
	Override     string    `json:"overide_status"`
	Active       Active   `json:"status"`
}
