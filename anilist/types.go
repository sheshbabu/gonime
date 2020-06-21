package anilist

type AnimeList = []AnimeListItem

type AnimeListItem struct {
	ID    int
	Cover string
}

type AnimeDetail struct {
	ID                    int
	Title                 string
	Description           string
	Cover                 string
	AverageScore          int
	TotalEpisodeCount     int
	NextEpisodeNumber     int
	NextEpisodeAiringDate string
	Genres                []string
}

type AnimeListRaw struct {
	Data struct {
		Page struct {
			Media []AnimeListItemRaw
		}
	}
}

type AnimeListItemRaw struct {
	ID    int
	Title struct {
		English       string
		UserPreferred string
	}
	CoverImage struct {
		Large string
	}
}
