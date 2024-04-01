package rottentomatoes

import (
	"log"

	"github.com/gocolly/colly"
)

type TVShowsResponse struct {
    TVShows []TVShow
    Url string
    Title string
}

type TVShow struct {
    Title string
    ImageUrl string
    Index int
    MeterScore string
    CriticsConcensus string
    ReleaseDate string
    Starring []string
    Directors []string
    Url string
}

type OnHTMLScrapeCallback func(TVShowsResponse)

func ScrapePopularTVShows() {
    url := "https://editorial.rottentomatoes.com/guide/popular-tv-shows/"
    collector := colly.NewCollector()

    collector.OnRequest(on_request)

    collector.OnResponse(on_response)

    collector.OnError(on_error)

    scrape(collector, func(response TVShowsResponse) {
        log.Println(response)  
    })

    collector.Visit(url)
}

func scrape(collector *colly.Collector, callback OnHTMLScrapeCallback) {
    response := TVShowsResponse{}
    collector.OnHTML("body", func(body *colly.HTMLElement) {
        response.Title = body.ChildText("h1")
        shows := make([]TVShow, 0, 200)
        body.ForEach(".row.countdown-item", func(i int, element *colly.HTMLElement) {
            imageUrl := element.ChildAttr(".article_poster", "src")
            tvShowUrl := element.ChildAttr(".article_movie_title a", "href")
            title := element.ChildText(".article_movie_title a")
            meterScore := element.ChildText(".article_movie_title .tMeterScore")
            year := element.ChildText(".subtle.start-year")

            critics := element.ChildText(".info.critics-consensus")
            casts := make([]string, 0, 10)
            element.ForEach(".info.cast a", func(_ int, el *colly.HTMLElement) {
                cast := el.Text
                casts = append(casts, cast)
            })
            directors := make([]string, 0, 10)
            element.ForEach(".info.director a", func(_ int, el *colly.HTMLElement) {
                director := el.Text
                directors = append(directors, director)
            })

            tvShow := TVShow{}
            tvShow.Title = title
            tvShow.ImageUrl = imageUrl
            tvShow.CriticsConcensus = critics
            tvShow.Starring = casts
            tvShow.Directors = directors
            tvShow.ReleaseDate = year
            tvShow.MeterScore = meterScore
            tvShow.Url = tvShowUrl
            tvShow.Index = i
            shows = append(shows, tvShow)
        })

        response.TVShows = shows
        callback(response)
    })
}

func on_request(request *colly.Request) {
    log.Println("OnRequest: ", request.URL)
}

func on_response(response *colly.Response) {
    log.Println("OnResponse: ", response.StatusCode)
}

func on_error(response *colly.Response, e error) {
    log.Println("OnError: ", response.StatusCode, e)
}
