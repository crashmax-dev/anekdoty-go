package scrapper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type OnScrapedCallback func(path string, scrapedData []string)

type Scrapper struct {
	collector         *colly.Collector
	baseURL           string
	parsedData        []string
	OnScrapedCallback OnScrapedCallback
}

func New(baseURL string, onScraped OnScrapedCallback) *Scrapper {
	scrapper := &Scrapper{
		collector:         colly.NewCollector(),
		baseURL:           baseURL,
		OnScrapedCallback: onScraped,
		parsedData:        make([]string, 0),
	}

	scrapper.collector.OnHTML("a.pagination-holder-next", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	scrapper.collector.OnHTML("div.holder-body", func(h *colly.HTMLElement) {
		scrapper.parsedData = append(scrapper.parsedData, h.ChildTexts("p")...)
	})

	scrapper.collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	return scrapper
}

func (s *Scrapper) Parse(path string) {
	s.collector.Visit(s.baseURL + path)
	s.OnScrapedCallback(path, s.parsedData)
	s.parsedData = make([]string, 0)
}
