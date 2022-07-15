package repository

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rhishimagdum/candystore/config"
	"github.com/rhishimagdum/candystore/entity"
	"github.com/rhishimagdum/candystore/logger"
)

type CustomerWebScraper struct {
}

//GetAll scrapes data from webpage
func (repository *CustomerWebScraper) GetAll() []entity.CustomerRecord {
	var records []entity.CustomerRecord
	res, err := http.Get(config.GetConfig().Url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		logger.Log().Error("Error while parsing webpage. Response code: %d. Error:  %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logger.Log().Error(err)
	}

	doc.Find("table").Each(func(tableIndex int, tablehtml *goquery.Selection) {
		// There are two tables in html page and we are interested in second table
		if (tableIndex) == 1 {
			// Read table rows one by one
			tablehtml.Find("tr").Each(func(rowIndex int, rowHtml *goquery.Selection) {
				// Create CustomerRecord entity for each row
				c := entity.CustomerRecord{}
				rowHtml.Find("td").Each(func(cellIndex int, cellHtml *goquery.Selection) {
					switch cellIndex {
					case 0:
						c.Name = cellHtml.Text()
					case 1:
						c.Candy = cellHtml.Text()
					case 2:
						var err error
						c.Eaton, err = strconv.Atoi(cellHtml.Text())
						if err != nil {
							panic(err)
						}
					}
				})
				if (c != entity.CustomerRecord{}) {
					records = append(records, c)
				}
			})
		}

	})
	logger.Log().Debug("Data pulled from web page")
	return records
}
