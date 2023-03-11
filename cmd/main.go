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
	scrapper := scrapper.New("https://anekdoty.ru/", onScraped)

	scrapper.Parse("pro-programmistov")
	scrapper.Parse("pro-mobilizaciyu")
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
