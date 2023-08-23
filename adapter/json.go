package adapter

import (
	"encoding/json"
	"fmt"
)

type DataGetter interface {
	GetData() (string, error)
}

type InputFormat struct {
	Addresseses []string `json:"addresseses"`
}

func JsonDataToAddresses(dataGetter DataGetter) ([]string,error) {
	input, err := dataGetter.GetData()
	if err != nil {
		return nil, fmt.Errorf("cannot GetData: %v", err)
	}
	var data InputFormat
	json.Unmarshal([]byte(input), &data)
	return data.Addresseses, nil
}