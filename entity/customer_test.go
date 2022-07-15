package entity

import (
	"reflect"
	"testing"
)

func TestGetAddCandy(t *testing.T) {

	got := getMockCustomer()

	want := &Customer{
		Name: "Ross",
		Snacks: map[string]int{
			"candy1": 10,
			"candy2": 5,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetFavouriteCandy(t *testing.T) {

	got := getMockCustomer().GetFavouriteCandy()

	want := "candy1"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetTotalCandies(t *testing.T) {

	got := getMockCustomer().GetTotalCandies()

	want := 15

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func getMockCustomer() *Customer {
	cust := NewCustomer("Ross")
	cust.AddCandy("candy1", 3)
	cust.AddCandy("candy2", 5)
	cust.AddCandy("candy1", 7)
	return cust
}
