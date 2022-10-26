package main

import (
	"dynamo-test/connection/dynamodb"
	"dynamo-test/postgres.repository"
	"fmt"
	"os"
)

func main() {
	client, err := dynamodb.DynamoConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	//dynamodb.CreateTableIfNotExists(client, "shopper_profile")

	//err = dynamodb.DeleteTable("shopper_profile", client)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//tableName := "shopper_profile"

	//var ShopperProfile models.ShopperProfile
	//if err := gofakeit.Struct(&ShopperProfile); err != nil {
	//	os.Exit(2)
	//}
	//fmt.Println(ShopperProfile)
	//err = postgres.repository.CreateShopperProfile(ShopperProfile, client)
	//if err != nil {
	//	fmt.Println(err)
	//}

	out, err := postgres_repository.ReadAllShopperProfile(client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
}
