package zippopotam

type Place struct {
	PlaceName         string `json:"place name"`
	Longitude         string `json:"longitude"`
	State             string `json:"state"`
	StateAbbreviation string `json:"state abbreviation"`
	Latitude          string `json:"latitude"`
}

type Response struct {
	PostCode            string  `json:"post code"`
	Country             string  `json:"country"`
	CountryAbbreviation string  `json:"country abbreviation"`
	Places              []Place `json:"places"`
}

type ErrNoResults struct{ Zip string }

func (e ErrNoResults) Error() string {
	return "no results for zip: " + e.Zip
}
