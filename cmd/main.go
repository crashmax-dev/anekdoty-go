package main

import (
	"anekdoty-go/internal/scrapper"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	categories := []string{"pro-1-aprelja", "pro-programmistov", "pro-mobilizaciyu"}
	scrapper := scrapper.New("https://anekdoty.ru/", onScraped)

	for _, category := range categories {
		scrapper.Parse(category)
	}
}

func onScraped(path string, scrapedData []string) {
	data, err := json.Marshal(scrapedData)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filepath.Join("data", fmt.Sprintf("%v.json", path)), data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Saved %d items to data/%v.json\n", len(scrapedData), path)
}
