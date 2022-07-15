package service

import (
	"sort"

	"github.com/rhishimagdum/candystore/entity"
	"github.com/rhishimagdum/candystore/repository"
)

type Customer struct {
	customerRepo repository.Customer
}

func NewCustomerService(repo repository.Customer) *Customer {
	return &Customer{
		customerRepo: repo,
	}
}

func (c *Customer) GetTopCustomersAndTheirFavourites() []entity.TopCustomerDetails {

	customerData := c.getCustomerData()

	var topCustomers entity.TopCustomers

	for _, v := range customerData {
		c := entity.TopCustomerDetails{
			CustomerName:   v.Name,
			FavouriteSnack: v.GetFavouriteCandy(),
			TotalSnacks:    v.GetTotalCandies(),
		}
		topCustomers = append(topCustomers, c)
	}
	sort.Sort(topCustomers)
	return topCustomers
}

func (c *Customer) getCustomerData() map[string]entity.Customer {
	customerData := make(map[string]entity.Customer)
	data := c.customerRepo.GetAll()
	for _, val := range data {
		if customer, ok := customerData[val.Name]; ok {
			customer.AddCandy(val.Candy, val.Eaton)
		} else {
			cust := *entity.NewCustomer(val.Name)
			cust.AddCandy(val.Candy, val.Eaton)
			customerData[val.Name] = cust
		}
	}
	return customerData
}
