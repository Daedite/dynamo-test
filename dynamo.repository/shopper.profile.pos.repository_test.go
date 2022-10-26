package dynamo_repository

import (
	"dynamo-test/connection/postgressdb"
	models "dynamo-test/model"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadShopperProfile(t *testing.T) {

}
func TestCount(t *testing.T) {
	sp, err := postgressdb.Connect()
	if err != nil {
		assert.Fail(t, err.Error())
	}
	fmt.Println(Count(sp))
}
func TestReadAllShopperProfile(t *testing.T) {
	sp, err := postgressdb.Connect()
	if err != nil {
		assert.Fail(t, err.Error())
	}
	profiles, err := ReadAllShopperProfile(sp)
	if err != nil {
		fmt.Println(err)
	}
	for _, profile := range profiles {
		fmt.Println(profile)
	}
}
func TestCreateShopperProfile(t *testing.T) {
	var ShopperProfile models.ShopperProfile
	sp, err := postgressdb.Connect()
	if err != nil {
		assert.Fail(t, err.Error())
	}

	for i := 0; i < 100; i++ {
		if err := gofakeit.Struct(&ShopperProfile); err != nil {
			os.Exit(2)
		}
		err := CreateShopperProfile(ShopperProfile, sp)
		if err != nil {
			fmt.Println(err)
		}
	}

}
func TestUpdateShopperProfile(t *testing.T) {

}
