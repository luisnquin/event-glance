package cinema

import "time"

type Movie struct {
	LocaleTitle string    `json:"localeTitle"`
	Date        time.Time `json:"date"`
	ImageLink   string    `json:"imageLink"`
	LocaleGenre string    `json:"localeGenre"`
	Director    string    `json:"director"`
}
