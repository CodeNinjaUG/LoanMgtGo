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
