package scrapper

import (
	"fmt"
	"sync"

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
		baseURL:           baseURL,
		OnScrapedCallback: onScraped,
		parsedData:        make([]string, 0),
		collector:         colly.NewCollector(),
	}

	return scrapper
}

func (s *Scrapper) Parse(path string) {
	var wg sync.WaitGroup

	s.collector.OnHTML("a.pagination-holder-next", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	s.collector.OnHTML("div.holder-body", func(h *colly.HTMLElement) {
		s.parsedData = append(s.parsedData, h.ChildTexts("p")...)
	})

	s.collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	s.collector.Visit(s.baseURL + path)

	wg.Wait()

	s.OnScrapedCallback(path, s.parsedData)
	s.parsedData = make([]string, 0)
}
