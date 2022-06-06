package models

import "gorm.io/gorm"

type AccountType string

const (
	Assets    AccountType = "Asset"
	Expense   AccountType = "Expense"
	Income    AccountType = "Income"
	Equity    AccountType = "Equity"
	Liability AccountType = "Liability"
)

type Parent string

const (
	Test                 Parent = "Test"
	Investment           Parent = "Investment"
	xxx                  Parent = "xxx"
	xxxIncome            Parent = "xxxIncome"
	ReturnFromUnitTrusts Parent = "ReturnFromUnitTrusts"
	ROI_Business         Parent = "ROI_Business"
)

type EntriesAllowed string

const (
	Y EntriesAllowed = "Yes"
	N EntriesAllowed = "No"
)

type AccountStatus string

const (
	Live      AccountStatus = "Live"
	NotActive AccountStatus = "NotActive"
)

type ChartAccount struct {
	Account_type   AccountType    `json:"account_type" binding:"required"`
	Parent         Parent         `json:"parent" binding:"required"`
	AccountName    string         `json:"account_name" binding:"required"`
	GL_Code        string         `json:"gl_code" binding:"required"`
	EntriesAllowed EntriesAllowed `json:"entries_allowed" binding:"required"`
	Notes          string         `json:"notes" binding:"required"`
}

//create account
func CreateAccount(db *gorm.DB, ChartAccount *ChartAccount) (err error) {
	err = db.Create(ChartAccount).Error
	if err != nil {
		return err
	}
	return nil
}

//get account
func GetAccounts(db *gorm.DB, ChartAccount *[]ChartAccount) (err error) {
	err = db.Find(ChartAccount).Error
	if err != nil {
		return err
	}
	return nil
}

///get account by id
func GetAccount(db *gorm.DB, ChartAccount *ChartAccount, id string) (err error) {
	err = db.Where("id=?", id).First(ChartAccount).Error
	if err != nil {
		return err
	}
	return nil
}

///update account
func UpdateAccount(db *gorm.DB, ChartAccount *ChartAccount) (err error) {
	db.Save(ChartAccount)
	return nil
}

///delete account
func DeleteAccount(db *gorm.DB, ChartAccount *ChartAccount, id string) (err error) {
	db.Where("id=?", id).Delete(ChartAccount)
	return nil
}
