package storage

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/dihmuzikien/smallurl/goapp"
)

type dynamoAPI interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

var (
	ErrNotFound = errors.New("cannot find URL with matching key")
)

type Db struct {
	client dynamoAPI
}

func (d Db) Put(url string) (*goapp.Url, error) {
	panic("implement me")
}

func (d Db) PutWithAlias(url, alias string) (*goapp.Url, error) {
	panic("implement me")
}

func NewDB() (*Db, error){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil{
		return nil, err
	}
	return &Db{
		client: dynamodb.NewFromConfig(cfg),
	}, nil
}

func (d Db) Get(id string) (string, error) {
	panic("implement me")
}



