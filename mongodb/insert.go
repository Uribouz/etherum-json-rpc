package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	collectionName = "transactions"
)
type Inserter struct {
	ctx context.Context
	databaseName string
}
func NewInserter (databaseName string, ctx context.Context) Inserter {
	return Inserter{ctx:ctx, databaseName:databaseName}
}

func (i Inserter) InsertJsonDataTransactions(data []string) error {
	bsonData := make([]interface{}, len(data))  
	for i, each := range data {
		var val interface{}
		err := bson.UnmarshalExtJSON([]byte(each), true, &val)
		if err != nil {	
			return fmt.Errorf("cannot do UnmarshalExtJSON, %v", err)
		}
		bsonData[i] = val
	}
	err := i.InsertBsonDataTransactions(bsonData)
	if err != nil {
		return fmt.Errorf("cannot do InsertJsonDataTransaction, %v", err)
	}
	return nil
}


func (i Inserter) InsertBsonDataTransactions(data []interface{}) error {
	//TODO: Must bulk limit insert ...
	db := dbClient.Database(i.databaseName)
	_, err := db.Collection(collectionName).InsertMany(i.ctx,data);
	if err != nil {
		return fmt.Errorf("cannot do InsertMany, %v", err)
	}
	return nil
}