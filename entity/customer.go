package entity

/*
 This entity customer and thier snacks
*/

type Customer struct {
	Name   string
	Snacks map[string]int
}

func NewCustomer(name string) *Customer {
	return &Customer{
		Name:   name,
		Snacks: make(map[string]int),
	}
}

func (c *Customer) AddCandy(name string, newQty int) {
	if currentQty, ok := c.Snacks[name]; ok {
		c.Snacks[name] = currentQty + newQty
	} else {
		c.Snacks[name] = newQty
	}
}

func (c *Customer) GetFavouriteCandy() string {
	maxEatonCandies := 0
	favouriteCandy := ""
	for k, v := range c.Snacks {
		if v > maxEatonCandies {
			maxEatonCandies = v
			favouriteCandy = k
		}
	}
	return favouriteCandy
}

func (c *Customer) GetTotalCandies() int {
	sum := 0
	for _, val := range c.Snacks {
		sum += val
	}
	return sum
}
