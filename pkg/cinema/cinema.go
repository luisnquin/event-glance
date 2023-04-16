package cinema

import (
	"fmt"
	"regexp"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/goodsign/monday"
)

var rxDateSeparator = regexp.MustCompile(`[0-9]{1,2}\s[a-z]{1,2}\s[A-z]+`)

func NextReleases() ([]Movie, error) {
	c := colly.NewCollector()

	now := time.Now()

	var (
		movies []Movie
		dates  []time.Time
	)

	c.OnHTML("body form .wrapper .contentWrapper .proxEstrenos .first .listProxEstreno .estrenoFecha", func(h *colly.HTMLElement) {
		if releases := h.ChildText(".diaEstreno"); releases != "" {
			for _, d := range rxDateSeparator.FindAllString(h.ChildText(" .diaEstreno"), -1) {
				date, err := monday.Parse(monday.DefaultFormatCaESLong, fmt.Sprintf("%s de %d", d, now.Year()), monday.LocaleEsES)
				if err != nil {
					return
				}

				dates = append(dates, date)
			}
		}

		var moviesInDay []Movie

		for _, link := range h.ChildAttrs(".listProxEstreno .slides li figure a img", "src") {
			moviesInDay = append(moviesInDay, Movie{ImageLink: link, Date: dates[h.Index]})
		}

		for i, director := range h.ChildAttrs(".listProxEstreno .slides li span", "data-director") {
			moviesInDay[i].Director = director
		}

		for i, title := range h.ChildTexts(".listProxEstreno .slides li figure a figcaption") {
			moviesInDay[i].LocaleTitle = title
		}

		for i, genre := range h.ChildAttrs(".listProxEstreno .slides li span", "data-genero") {
			moviesInDay[i].LocaleGenre = genre
		}

		movies = append(movies, moviesInDay...)
	})

	c.Visit("https://cinepolis.com.sv/proximos-estrenos")

	return movies, nil
}
