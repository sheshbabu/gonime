package anilist

import (
	"encoding/json"
	"html/template"
	"strings"
	"time"
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
			bannerImage
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
	ID string
}

func GetAnimeDetail(id string) AnimeDetail {
	var rawDetail AnimeDetailRaw
	var detail AnimeDetail

	body := Request(getAnimeDetailQuery, getAnimeDetailQueryVariables{ID: id})

	json.Unmarshal(body, &rawDetail)

	title := rawDetail.Data.Media.Title.English

	if title == "null" {
		title = rawDetail.Data.Media.Title.UserPreferred
	}

	description := template.HTML(strings.ReplaceAll(rawDetail.Data.Media.Description, "’", "'"))

	genres := strings.Join(rawDetail.Data.Media.Genres, ", ")

	shouldDisplayNextEpisode := true

	if rawDetail.Data.Media.NextAiringEpisode.AiringAt == 0 || rawDetail.Data.Media.NextAiringEpisode.AiringAt == 0 {
		shouldDisplayNextEpisode = false
	}

	nextEpisodeAiringDate := time.Unix(rawDetail.Data.Media.NextAiringEpisode.AiringAt, 0).Format("02 Jan 2006 3:04PM MST")

	detail = AnimeDetail{
		ID:                       rawDetail.Data.Media.ID,
		Cover:                    rawDetail.Data.Media.CoverImage.Large,
		Banner:                   rawDetail.Data.Media.BannerImage,
		Title:                    title,
		Description:              description,
		AverageScore:             rawDetail.Data.Media.AverageScore,
		TotalEpisodeCount:        rawDetail.Data.Media.Episodes,
		ShouldDisplayNextEpisode: shouldDisplayNextEpisode,
		NextEpisodeNumber:        rawDetail.Data.Media.NextAiringEpisode.Episode,
		NextEpisodeAiringDate:    nextEpisodeAiringDate,
		Genres:                   genres,
	}

	return detail
}
