package postgres_repository

import (
	"dynamo-test/connection/dynamodb"
	models "dynamo-test/model"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateShopperProfile(t *testing.T) {
	client, err := dynamodb.DynamoConnection()
	if err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
	}

	var ShopperProfile models.ShopperProfile

	for i := 0; i < 100; i++ {
		if err := gofakeit.Struct(&ShopperProfile); err != nil {
			os.Exit(2)
		}
		err = CreateShopperProfile(ShopperProfile, client)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("done!")
}
func TestCreateUpdateShopperProfile(t *testing.T) {
	client, err := dynamodb.DynamoConnection()
	if err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
	}

	err = CreateUpdateShopperProfile(models.ShopperProfile{ID: 7206, OutstandingBalance: 666}, client)
	if err != nil {
		fmt.Println(err)
	}
}
func TestDeleteShopperProfileTable(t *testing.T) {
	client, err := dynamodb.DynamoConnection()
	if err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
	}
	err = dynamodb.DeleteTable("shopper_profile", client)
	if err != nil {
		fmt.Println(err)
	}
}
func TestCreateShopperProfileTable(t *testing.T) {
	client, err := dynamodb.DynamoConnection()
	if err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
	}
	dynamodb.CreateTableIfNotExists(client, "shopper_profile")
}
func TestReadAllShopperProfile(t *testing.T) {
	client, err := dynamodb.DynamoConnection()
	if err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
	}
	out, err := ReadAllShopperProfile(client)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.NotNil(t, out)
	for _, shoper := range out {
		fmt.Println(shoper)
	}
}
func TestRead(t *testing.T) {
	client, err := dynamodb.DynamoConnection()
	if err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
	}
	data, err := Read(7206, client)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.NotNil(t, data)
	fmt.Println(data)
}
func TestReadShopperProfile(t *testing.T) {
	client, err := dynamodb.DynamoConnection()
	if err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
	}
	data, err := ReadShopperProfile(7206, client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
func TestUpdate(t *testing.T) {
	client, err := dynamodb.DynamoConnection()
	if err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
	}
	err = Update(7206, 666, client)
	fmt.Println(err)
}
