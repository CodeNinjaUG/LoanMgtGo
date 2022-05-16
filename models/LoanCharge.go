package models

import "gorm.io/gorm"


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

type Currency string

const (
	UGX Currency = "UGX"
	KES Currency = "KES"
	USD Currency = "USD"
)



type Penalty string

const (
	Nop  Penalty = "Nop"
	Yep  Penalty = "Yep"
)

type Active string 

const (
	InActive Active = "InActive"
	Activated  Active = "Activated"
)


type LoanCharge struct {
	gorm.Model
	Name         string       `json:"name"`
	ChargeType   ChargeType   `json:"charge_type"`
	Amount       float64      `json:"amount"`
	ChargeOption ChargeOption `json:"charge_option"`
	Currency     Currency     `json:"currency"`
	Penalty      Penalty      `json:"penalty"`
	Override     string       `json:"overide_status"`
	Active       Active       `json:"status"`
}


func CreateLoanCharge(db *gorm.DB, LoanCharge *LoanCharge) (err error) {
	err = db.Create(LoanCharge).Error
	if err != nil {
		return err
	}
	return nil
}

//get loan products

func GetLoanCharges(db *gorm.DB, LoanCharge *[]LoanCharge) (err error) {
	err = db.Find(LoanCharge).Error
	if err != nil {
		return nil
	}
	return nil
}

//get loan product by ID
func GetLoanCharge(db *gorm.DB, LoanCharge *LoanCharge, id string) (err error) {
	err = db.Where("id = ?", id).First(LoanCharge).Error
	if err != nil {
		return err
	}
	return nil
}

//Update loan product
// still pending research
func UpdateLoanCharge(db *gorm.DB, LoanCharge *LoanCharge) (err error) {
	//db.Where("id =?",id).Save(LoanProduct)
	db.Save(LoanCharge)
	return nil
}

//Delete loan Product
func DeleteLoanCharge(db *gorm.DB, LoanCharge *LoanCharge, id string) (err error) {
	db.Where("id =?", id).Delete(LoanCharge)
	return nil
}