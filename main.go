package main

import (
	"os"

	"github.com/EmmanuelStan12/web-scraper/output"
	"github.com/EmmanuelStan12/web-scraper/rottentomatoes"
)

func main() {
    filename := os.Args[1]
    exportType := os.Args[2]
    rottentomatoes.ScrapePopularTVShows(func(shows []rottentomatoes.TVShow) {
        output.Export(shows, exportType, filename)
    })
}
