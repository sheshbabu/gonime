package anilist

import "html/template"

type AnimeList = []AnimeListItem

type AnimeListItem struct {
	ID    int
	Cover string
}

type AnimeDetail struct {
	ID                    int
	Title                 string
	Description           template.HTML
	Cover                 string
	AverageScore          int
	TotalEpisodeCount     int
	NextEpisodeNumber     int
	NextEpisodeAiringDate int
	Genres                []string
}

type AnimeListRaw struct {
	Data struct {
		Page struct {
			Media []AnimeListItemRaw
		}
	}
}

type AnimeDetailRaw struct {
	Data struct {
		Media AnimeListItemRaw
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
	Description       string
	Genres            []string
	AverageScore      int
	Episodes          int
	NextAiringEpisode struct {
		Episode  int
		AiringAt int
	}
}
