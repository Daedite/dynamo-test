package postgres_repository

import (
	"context"
	models "dynamo-test/model"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"strconv"
	"time"
)

var ShopperProfileTable = "shopper_profile"

func Update(ID uint, balance int64, client *dynamodb.Client) error {

	upd := expression.
		Set(expression.Name("balance"), expression.Value(balance))

	cond := expression.Equal(
		expression.Name("anyIntField"),
		expression.Value(1))

	expr, err := expression.NewBuilder().WithUpdate(upd).WithCondition(cond).Build()

	out, err := client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(ShopperProfileTable),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{
				Value: strconv.Itoa(int(ID)),
			},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ConditionExpression:       expr.Condition(),
	})
	if err != nil {
		return fmt.Errorf("TrasnacitonWrite: %v\n", err)
	}
	fmt.Println(out)
	return nil
}
func Read(ID uint, client *dynamodb.Client) (models.ShopperProfile, error) {
	var sp models.ShopperProfileSample
	data, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(ShopperProfileTable),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{
				Value: strconv.Itoa(int(ID)),
			},
		},
	})
	if err != nil {
		return models.ShopperProfile{}, err
	}
	if data.Item == nil {
		return models.ShopperProfile{}, fmt.Errorf("GetItem: Data not found.\n")
	}

	err = attributevalue.UnmarshalMap(data.Item, &sp)
	if err != nil {
		return models.ShopperProfile{}, fmt.Errorf("UnmarshalMap: %v\n", err)
	}
	return cleanShopperProfileStruct(sp), nil
}
func CreateShopperProfile(shopperProfile models.ShopperProfile, client *dynamodb.Client) error {
	fmt.Println(int(shopperProfile.ID))
	_, err := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(ShopperProfileTable),
		Item: map[string]types.AttributeValue{
			"id":        &types.AttributeValueMemberN{Value: strconv.Itoa(int(shopperProfile.ID))},
			"email":     &types.AttributeValueMemberS{Value: shopperProfile.Email},
			"currency":  &types.AttributeValueMemberS{Value: shopperProfile.Currency},
			"balance":   &types.AttributeValueMemberN{Value: strconv.FormatInt(int64(shopperProfile.OutstandingBalance), 10)},
			"createdAt": &types.AttributeValueMemberS{Value: strconv.FormatInt(shopperProfile.CreatedAt.Unix(), 10)},
			"updatedAt": &types.AttributeValueMemberS{Value: strconv.FormatInt(shopperProfile.UpdatedAt.Unix(), 10)},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
func CreateUpdateShopperProfile(shopperProfile models.ShopperProfile, client *dynamodb.Client) error {
	fmt.Println(int(shopperProfile.ID))
	_, err := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(ShopperProfileTable),
		Item: map[string]types.AttributeValue{
			"id":      &types.AttributeValueMemberN{Value: strconv.Itoa(int(shopperProfile.ID))},
			"balance": &types.AttributeValueMemberN{Value: strconv.FormatInt(int64(shopperProfile.OutstandingBalance), 10)},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

// CreateShopperProfile2 Todo test this method.
func CreateShopperProfile2(shopperProfile models.ShopperProfile, client *dynamodb.Client) error {
	av, err := attributevalue.MarshalMap(shopperProfile)
	if err != nil {
		return err
	}
	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(ShopperProfileTable),
		Item:      av,
	})
	if err != nil {
		return fmt.Errorf("failed to put Record, %w", err)
	}
	return nil
}

// todo failed.
func ReadShopperProfile(ID uint, client *dynamodb.Client) (models.ShopperProfileSample, error) {
	var sp models.ShopperProfileSample
	concatenated := "hash#" + strconv.Itoa(int(ID))
	//requestKey := fmt.Sprintf("hk#", ID)
	fmt.Println(concatenated)
	selectedKeys := map[string]string{
		"id": concatenated,
	}
	key, err := attributevalue.MarshalMap(selectedKeys)
	if err != nil {
		fmt.Println(err)
	}
	data, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(ShopperProfileTable),
		Key:       key,
	})
	if err != nil {
		return sp, fmt.Errorf("GetItem: %v\n", err)
	}

	if data.Item == nil {
		return sp, fmt.Errorf("GetItem: Data not found.\n")
	}

	err = attributevalue.UnmarshalMap(data.Item, &sp)
	if err != nil {
		return sp, fmt.Errorf("UnmarshalMap: %v\n", err)
	}

	return sp, nil
}
func ReadAllShopperProfile(client *dynamodb.Client) ([]models.ShopperProfile, error) {
	out, err := client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(ShopperProfileTable),
	})
	if err != nil {
		return nil, err
	}
	if len(out.Items) == 0 {
		return nil, nil
	}

	var spsl []models.ShopperProfileSample
	err = attributevalue.UnmarshalListOfMaps(out.Items, &spsl)
	if err != nil {
		fmt.Println(err)
	}
	var spl []models.ShopperProfile
	for _, sps := range spsl {
		spl = append(spl, cleanShopperProfileStruct(sps))
	}
	return spl, err
}

// cleanShopperProfileStruct
func cleanShopperProfileStruct(sp models.ShopperProfileSample) models.ShopperProfile {
	return models.ShopperProfile{ID: sp.ID, Email: sp.Email, Currency: sp.Currency, OutstandingBalance: sp.Balance, CreatedAt: getTime(sp.CreatedAt), UpdatedAt: getTime(sp.UpdatedAt)}
}

// getTime
func getTime(int642 string) time.Time {
	timeUnti, _ := strconv.ParseInt(int642, 0, 64)
	return time.Unix(timeUnti, 0)
}
