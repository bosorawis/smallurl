package storage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/dihmuzikien/smallurl/goapp"
)

type dynamoAPI interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}



type DB struct {
	client dynamoAPI
}

func NewDB() (*DB, error){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil{
		return nil, err
	}
	return &DB{
		client: dynamodb.NewFromConfig(cfg),
	}, nil
}

func (D DB) Get(id string) (string, error) {
	panic("implement me")
}

func (D DB) Put(url string) (*goapp.Url, error) {
	panic("implement me")
}

