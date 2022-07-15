package main

import (
	"encoding/json"
	"fmt"

	"github.com/rhishimagdum/candystore/repository"
	"github.com/rhishimagdum/candystore/service"
)

func main() {
	customerService := service.NewCustomerService(&repository.CustomerWebScraper{})
	data := customerService.GetTopCustomersAndTheirFavourites()
	output, _ := json.Marshal(data)
	fmt.Println(string(output))
}
