package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	var quote string
	
	url := "https://www.bing.com/search?q=Quote%20of%20the%20day"

	c := colly.NewCollector(
		colly.AllowedDomains("www.bing.com"),
	)

		// extract the quote from the page
		c.OnHTML(".bt_qotd", func(e *colly.HTMLElement) {
			res := e.ChildText("div.bt_quoteText")
			
			//if quote not found
			if len(res) < 1 {
				return
			}

			quote = res
	
			fmt.Println("\n"+"“"+quote+"”"+"\n")
		})
	
	c.Visit(url)
}
