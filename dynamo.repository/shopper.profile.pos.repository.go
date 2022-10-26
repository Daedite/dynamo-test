package dynamo_repository

import (
	models "dynamo-test/model"
	"gorm.io/gorm"
)

// CreateShopperProfile Creates a shopper profile in a db.
func CreateShopperProfile(shopperProfile models.ShopperProfile, gormDB *gorm.DB) error {
	if err := gormDB.Create(&shopperProfile).Error; err != nil {
		return err
	}
	return nil
}

// UpdateShopperProfile updates a shopper profile in a db.
func UpdateShopperProfile(shopperId uint, balance *float64, gormDB *gorm.DB) (models.ShopperProfile, error) {
	gormShopperProfile := models.ShopperProfile{}
	if err := gormDB.Where("id = ?", shopperId).First(&gormShopperProfile).Update("outstanding_balance", balance).Error; err != nil {
		return gormShopperProfile, nil
	}
	return gormShopperProfile, nil
}

// ReadShopperProfile reads shopper table state based on shopperIdentityId.
func ReadShopperProfile(shopperIdentityId string, gormDB *gorm.DB) (models.ShopperProfile, error) {
	gormShopperProfile := models.ShopperProfile{}
	if err := gormDB.First(&gormShopperProfile, shopperIdentityId).Error; err != nil {
		return gormShopperProfile, nil
	}
	return gormShopperProfile, nil
}

func ReadAllShopperProfile(gormDB *gorm.DB) ([]models.ShopperProfile, error) {
	gormShopperProfile := []models.ShopperProfile{}
	if err := gormDB.Find(&gormShopperProfile).Error; err != nil {
		return gormShopperProfile, nil
	}
	return gormShopperProfile, nil
}
func Count(gormDB *gorm.DB) int64 {
	var shopper int64
	gormDB.Table("shopper_profiles").Count(&shopper)
	return shopper
}
