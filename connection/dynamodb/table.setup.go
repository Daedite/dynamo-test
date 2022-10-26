package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

func tableExists(d *dynamodb.Client, name string) bool {
	tables, err := d.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatal("ListTables failed", err)
	}
	for _, n := range tables.TableNames {
		if n == name {
			return true
		}
	}
	return false
}

func CreateTableIfNotExists(d *dynamodb.Client, tableName string) {
	if tableExists(d, tableName) {
		log.Printf("table=%v already exists\n", tableName)
		return
	}
	_, err := d.CreateTable(context.TODO(), buildCreateTableInput(tableName))
	if err != nil {
		log.Fatal("CreateTable failed", err)
	}
	log.Printf("created table=%v\n", tableName)
}

func buildCreateTableInput(tableName string) *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeN,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}
}

func DeleteTable(tableName string, client *dynamodb.Client) error {
	_, err := client.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	})
	return err
}
