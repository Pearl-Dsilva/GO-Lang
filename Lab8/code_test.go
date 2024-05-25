package main

import (
	"testing"
)

func TestMarshalAndUnMarshal(t *testing.T) {
	var testProduct, testProduct2 Product
	testProduct.Name = "Onion"
	testProduct.Price = 100.5
	testProduct.Stock = 20
	res, err := testProduct.MarshalJSON()
	if err != nil {
		t.Error("Something went Wrong while Marshal", err)
	}
	err2 := testProduct2.UnmarshalJSON(res)
	if err2 != nil {
		t.Error("Something went Wrong while unMarshal", err2)
	}
	if !(testProduct.Name == testProduct2.Name && testProduct.Price == testProduct2.Price && testProduct.Stock == testProduct2.Stock) {
		t.Error("Data Modified", err)
	}

}
