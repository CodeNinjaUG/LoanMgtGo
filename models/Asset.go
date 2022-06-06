package models

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	BranchAccountID int    `json:"branch_account_id"`
	Name            string `json:"name"`
	AssetTypeID     int    `json:"asset_type_id"`
	AssetType       AssetType
	PurchaseDate    time.Time
}

// create an asset
func CreateAsset(db *gorm.DB, Asset *Asset) (err error) {
	err = db.Create(Asset).Error
	if err != nil {
		return err
	}
	return nil
}

//get Asset by ID
func GetAsset(db *gorm.DB, Asset *Asset, id string) (err error) {
	err = db.Where("id = ?", id).First(Asset).Error
	if err != nil {
		return err
	}
	return nil
}

//Update Asset
func UpdateAsset(db *gorm.DB, Asset *Asset) (err error) {
	//db.Where("id =?",id).Save(Asset)
	db.Save(Asset)
	return nil
}

//Delete Asset
func DeleteAsset(db *gorm.DB, Asset *Asset, id string) (err error) {
	db.Where("id =?", id).Delete(Asset)
	return nil
}
