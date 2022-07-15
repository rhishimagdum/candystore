package entity

/*
 This entity top customer details which is expected in the putput
*/

type TopCustomerDetails struct {
	CustomerName   string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks    int    `json:"totalSnacks"`
}

type TopCustomers []TopCustomerDetails

func (c TopCustomers) Len() int {
	return len(c)
}

func (c TopCustomers) Less(i, j int) bool {
	return c[i].TotalSnacks > c[j].TotalSnacks
}

func (c TopCustomers) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
