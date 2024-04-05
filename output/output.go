package output

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"

	"github.com/EmmanuelStan12/web-scraper/rottentomatoes"
)

func Export(items []rottentomatoes.TVShow, exportType string, filename string) {
    switch exportType {
    case "csv":
        export_csv(items, filename)
    case "json": 
        export_json(items, filename)
    default:
        log.Fatalln("Cannot export type: ", exportType)
    }
}

func export_csv(items []rottentomatoes.TVShow, filename string) {
    file, err := os.Create(filename)
    if err != nil {
        log.Fatalln("failed to create file", err)
    }
    defer file.Close()
    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write(items[0].GetHeaders())
    for _, show := range items {
        row := show.ToArray()
        if err := writer.Write(row); err != nil {
            log.Fatalln("Error writing record to file", err)
        }
    }
}

func export_json(items []rottentomatoes.TVShow, filename string) {
    file, err := os.Create(filename)
    if err != nil {
        log.Fatalln("Error creating file", err) 
    }
    defer file.Close()
    shows := [][]string {}
    for _, show := range items {
        shows = append(shows, show.ToArray())
    }
    data, err := json.Marshal(shows)
    if err != nil {
        log.Fatalln("Failed to create json", err)
    }

    _, err = file.Write(data)
    if err != nil {
        log.Fatalln("Failed to write a file", err)
    }
}
