package task

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/expression"
	"github.com/davecgh/go-spew/spew"
)

type Store struct {
	db    dynamodbiface.ClientAPI
	table string
}

func NewStore(db dynamodbiface.ClientAPI, table string) *Store {
	return &Store{
		db:    db,
		table: table,
	}
}

func (s Store) Save(ctx context.Context, conn Connection) error {
	dbItem, err := dynamodbattribute.MarshalMap(&conn)
	if err != nil {
		return err
	}
	spew.Dump(dbItem)

	cond := expression.AttributeNotExists(expression.Name("taskID"))
	expr, err := expression.NewBuilder().WithCondition(cond).Build()
	if err != nil {
		return err
	}
	req := s.db.PutItemRequest(&dynamodb.PutItemInput{
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		Item:                      dbItem,
		TableName:                 aws.String(s.table),
	})

	_, err = req.Send(ctx)
	if err != nil {
		if strings.Contains(err.Error(), dynamodb.ErrCodeConditionalCheckFailedException) {
			return nil
		}
		return err
	}
	return err
}

func (s Store) Retrieve(ctx context.Context, taskID string) (*Connection, error) {
	req := s.db.GetItemRequest(&dynamodb.GetItemInput{
		Key: map[string]dynamodb.AttributeValue{
			"taskID": {
				S: aws.String(taskID),
			},
		},
		TableName: aws.String(s.table),
	})

	out, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}

	var conn Connection
	err = dynamodbattribute.UnmarshalMap(out.Item, &conn)
	if err != nil {
		return nil, err
	}

	return &conn, nil
}
