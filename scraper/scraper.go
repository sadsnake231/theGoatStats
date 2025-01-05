package scraper

import (
	"github.com/gocolly/colly"
	"time"
	"os"
	"fmt"
)


func InfoUpdate() {
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error){
		fmt.Println("Error:", err)
	})

	c.OnRequest(func(r *colly.Request) {
        r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")
    })

	c.SetRequestTimeout(30 * time.Second)

	c.OnHTML("div.stats_pullout", func(e *colly.HTMLElement) {
		years := ""
		e.ForEach("span strong", func(i int, el *colly.HTMLElement){
			if i == 0 {
				years = el.Text
			}
		})
		fmt.Println("Season:", years)

		tournaments := []string{}

		e.ForEach("p strong", func(_ int, el *colly.HTMLElement){
			tournaments = append(tournaments, el.Text)
		})
		fmt.Println("Tournaments:", tournaments)

		stats := make(map[string][]string)

		e.ForEach(".p1, .p2, .p3", func(_ int, block *colly.HTMLElement){
			block.ForEach("div", func(_ int, stat *colly.HTMLElement){
				category := stat.ChildText("span strong")
				if category == ""{
					return
				}

				values := []string{}
				stat.ForEach("p", func(_ int, value *colly.HTMLElement){
					values = append(values, value.Text)
				})

				stats[category] = values
			})
		})

		for category, values := range stats {
			fmt.Printf("%s: %v\n", category, values)	
		}

	})

	c.Visit(os.Getenv("SOURCE"))
}