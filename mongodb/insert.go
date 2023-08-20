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
}
func NewInserter (ctx context.Context) Inserter {
	return Inserter{ctx:ctx}
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
	_, err := getDB().Collection(collectionName).InsertMany(i.ctx,data);
	if err != nil {
		return fmt.Errorf("cannot do InsertMany, %v", err)
	}
	return nil
}