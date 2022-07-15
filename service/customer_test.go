package service

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rhishimagdum/candystore/entity"
	"github.com/rhishimagdum/candystore/repository/mock"
)

func TestGetTopCustomersAndTheirFavourites(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockCustomerRepository(mockCtrl)
	mockRepo.EXPECT().GetAll().Return([]entity.CustomerRecord{
		{
			Name:  "Ross",
			Candy: "Candy 1",
			Eaton: 10,
		},
		{
			Name:  "Joe",
			Candy: "Candy 2",
			Eaton: 15,
		},
		{
			Name:  "Monica",
			Candy: "Candy 1",
			Eaton: 5,
		},
		{
			Name:  "Ross",
			Candy: "Candy 3",
			Eaton: 25,
		},
	})

	customerService := NewCustomerService(mockRepo)
	want := []entity.TopCustomerDetails{
		{
			CustomerName:   "Ross",
			FavouriteSnack: "Candy 3",
			TotalSnacks:    35,
		},
		{
			CustomerName:   "Joe",
			FavouriteSnack: "Candy 2",
			TotalSnacks:    15,
		},
		{
			CustomerName:   "Monica",
			FavouriteSnack: "Candy 1",
			TotalSnacks:    5,
		},
	}
	got := customerService.GetTopCustomersAndTheirFavourites()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
