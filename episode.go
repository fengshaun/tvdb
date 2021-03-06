package tvdb

// Episode struct store all data of an episode.
type Episode struct {
	AbsoluteNumber     int      `json:"absoluteNumber"`
	AiredEpisodeNumber int      `json:"airedEpisodeNumber"`
	AiredSeason        int      `json:"airedSeason"`
	AirsAfterSeason    int      `json:"airsAfterSeason"`
	AirsBeforeEpisode  int      `json:"airsBeforeEpisode"`
	AirsBeforeSeason   int      `json:"airsBeforeSeason"`
	Director           string   `json:"director"`
	Directors          []string `json:"directors"`
	DvdChapter         int      `json:"dvdChapter"`
	DvdDiscid          string   `json:"dvdDiscid"`
	DvdEpisodeNumber   float64  `json:"dvdEpisodeNumber"`
	DvdSeason          int      `json:"dvdSeason"`
	EpisodeName        string   `json:"episodeName"`
	Filename           string   `json:"filename"`
	FirstAired         string   `json:"firstAired"`
	GuestStars         []string `json:"guestStars"`
	ID                 int      `json:"id"`
	ImdbID             string   `json:"imdbId"`
	LastUpdated        int      `json:"lastUpdated"`
	LastUpdatedBy      int      `json:"lastUpdatedBy"`
	Overview           string   `json:"overview"`
	ProductionCode     string   `json:"productionCode"`
	SeriesID           int      `json:"seriesId"`
	ShowURL            string   `json:"showURL"`
	SiteRating         float32  `json:"siteRating"`
	SiteRatingCount    int      `json:"siteRatingCount"`
	ThumbAdded         string   `json:"thumbAdded"`
	ThumbAuthor        int      `json:"thumbAuthor"`
	ThumbHeight        string   `json:"thumbHeight"`
	ThumbWidth         string   `json:"thumbWidth"`
	Writers            []string `json:"writers"`
}

// Empty verify if the episode's fields are empty and don't are filled by an api
// response.
func (e *Episode) Empty() bool {
	return e.ID == 0 && e.EpisodeName == ""
}
