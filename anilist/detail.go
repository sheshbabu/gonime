package anilist

import (
	"encoding/json"
	"html/template"
	"strings"
)

var getAnimeDetailQuery = `
	query {
		Media(id: {{ .ID}}) {
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
`

type getAnimeDetailQueryVariables struct {
	ID int
}

func GetAnimeDetail(id int) AnimeDetail {
	var rawDetail AnimeDetailRaw
	var detail AnimeDetail

	body := Request(getAnimeDetailQuery, getAnimeDetailQueryVariables{ID: id})

	json.Unmarshal(body, &rawDetail)

	title := rawDetail.Data.Media.Title.English

	if title == "null" {
		title = rawDetail.Data.Media.Title.UserPreferred
	}

	description := template.HTML(rawDetail.Data.Media.Description)

	genres := strings.Join(rawDetail.Data.Media.Genres, ", ")

	detail = AnimeDetail{
		ID:                    rawDetail.Data.Media.ID,
		Cover:                 rawDetail.Data.Media.CoverImage.Large,
		Title:                 title,
		Description:           description,
		AverageScore:          rawDetail.Data.Media.AverageScore,
		TotalEpisodeCount:     rawDetail.Data.Media.Episodes,
		NextEpisodeNumber:     rawDetail.Data.Media.NextAiringEpisode.Episode,
		NextEpisodeAiringDate: rawDetail.Data.Media.NextAiringEpisode.AiringAt,
		Genres:                genres,
	}

	return detail
}
