package bq

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMarshalling(t *testing.T) {
	tran := &Transaction{
		Created: DateTime{time.Now()},
		Items: []*Item{
			{SKU: "123", Created: DateTime{time.Now()}},
		},
		Employee: &Employee{
			EmployeeID: "123",
		},
	}

	v, _, err := tran.Save()
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(v)
	fmt.Println(string(b))
}
