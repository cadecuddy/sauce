package types

// Holds unmarshalled trace.moe JSON response data
type MoeResponse struct {
	Framecount int      `json:"frameCount"`
	Error      string   `json:"error"`
	Result     []Result `json:"result"`
}

// Body of relevant data when a successful response goes through
type Result struct {
	Anilist    AnilistData `json:"anilist"`
	Filename   string      `json:"filename"`
	Episode    int         `json:"episode"`
	From       float64     `json:"from"`
	To         float64     `json:"to"`
	Similarity float64     `json:"similarity"`
	Video      string      `json:"video"`
	Image      string      `json:"image"`
}

// Anilist information related to show/episode
type AnilistData struct {
	ID    int `json:"id"`
	MalID int `json:"idMal"`
	Title struct {
		Native  string `json:"native"`
		Romaji  string `json:"romaji"`
		English string `json:"english"`
	} `json:"title"`
	Synonyms []string `json:"synonyms"`
	IsAdult  bool     `json:"isAdult"`
}
