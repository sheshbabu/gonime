package anilist

import (
	"encoding/json"
	"time"
)

var getPopularAnimeQuery = `
	query {
		Page(page: 1, perPage: 50) {
			media(sort: POPULARITY_DESC, season: {{ .Season}}, seasonYear: {{ .SeasonYear}}, type: ANIME, isAdult: false, averageScore_greater: 60) {
				id
				title {
					english
					userPreferred
				}
				coverImage {
					large
				}
				description
				genres
				averageScore
				episodes
				nextAiringEpisode {
					airingAt
					episode
				}
			}
		}
	}
`

type getPopularAnimeQueryVariables struct {
	Season     string
	SeasonYear int
}

func GetPopularAnime() AnimeList {
	var rawList AnimeListRaw
	var list []AnimeListItem

	season := GetCurrentSeasonName()
	year := time.Now().Year()

	body := Request(getPopularAnimeQuery, getPopularAnimeQueryVariables{season, year})

	json.Unmarshal(body, &rawList)

	for _, item := range rawList.Data.Page.Media {
		list = append(list, AnimeListItem{ID: item.ID, Cover: item.CoverImage.Large})
	}

	return list
}
