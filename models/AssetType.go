package models

import (
	"gorm.io/gorm"
)

type AssetType struct {
	gorm.Model
	Name      string  `json:"name"`
	AccountID int     `json:"account_id"`
	Assets    []Asset `gorm:"has_many:assettype_has_many_assets"`
	Notes     string  `json:"notes"`
}

// create a asset type
func CreateAssetType(db *gorm.DB, AssetType *AssetType) (err error) {
	err = db.Create(AssetType).Error
	if err != nil {
		return err
	}
	return nil
}

// get asset type

func GetAssetType(db *gorm.DB, AssetType *[]AssetType) (err error) {
	err = db.Find(AssetType).Error
	if err != nil {
		return err
	}
	return nil
}

//update asset type

func UpdateAssetType(db *gorm.DB, AssetType *AssetType, id string) (err error) {
	//db.where("id=?",id).Save(AssetType)
	db.Save(AssetType)
	return nil
}

//Delete AssetType
func DeleteAssetType(db *gorm.DB, Asset *Asset, id string) (err error) {
	db.Where("id =?", id).Delete(Asset)
	return nil
}
