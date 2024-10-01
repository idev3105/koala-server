package mapper

import (
	"log"

	"github.com/jackc/pgx/v5/pgtype"
	commonentity "org.idev.koala/backend/domain/common/entity"
	commonenum "org.idev.koala/backend/domain/common/enum"
	movieentity "org.idev.koala/backend/domain/movie/entity"
	movieenum "org.idev.koala/backend/domain/movie/enum"
	sqlc_generated "org.idev.koala/backend/generated/sqlc"
	"org.idev.koala/backend/utils"
)

func MapToDomainMovie(m sqlc_generated.Movie) *movieentity.Movie {
	return &movieentity.Movie{
		MovieID:         m.MovieID,
		ItemStatus:      m.ItemStatus,
		Name:            m.Name,
		Description:     m.Description,
		ThumbnailUrl:    m.ThumbnailUrl.String,
		AvailableStatus: commonenum.AvailableStatus(m.AvailableStatus),
		Types:           m.Types,
		ReleaseYear:     m.ReleaseYear,
		TotalViewer:     m.TotalViewer.Int32,
		TotalUpvote:     m.TotalUpvote.Int32,
		TotalDownvote:   m.TotalDownvote.Int32,
		TotalRate:       m.TotalRate.Int32,
		Rate:            float32(utils.ConvertNumericToFloat(m.Rate)),
	}
}

func MapToSQLModelMovie(m movieentity.Movie) *sqlc_generated.Movie {
	var rateNumeric pgtype.Numeric
	err := rateNumeric.Scan(m.Rate)
	if err != nil {
		// Handle the error appropriately. For now, we'll log it.
		log.Printf("Error scanning rate: %v", err)
	}
	return &sqlc_generated.Movie{
		MovieID:         m.MovieID,
		ItemStatus:      m.ItemStatus,
		Name:            m.Name,
		Description:     m.Description,
		ThumbnailUrl:    pgtype.Text{String: m.ThumbnailUrl, Valid: m.ThumbnailUrl != ""},
		AvailableStatus: m.AvailableStatus.String(),
		Types:           m.Types,
		ReleaseYear:     m.ReleaseYear,
		TotalViewer:     pgtype.Int4{Int32: m.TotalViewer, Valid: m.TotalViewer != 0},
		TotalUpvote:     pgtype.Int4{Int32: m.TotalUpvote, Valid: m.TotalUpvote != 0},
		TotalDownvote:   pgtype.Int4{Int32: m.TotalDownvote, Valid: m.TotalDownvote != 0},
		TotalRate:       pgtype.Int4{Int32: m.TotalRate, Valid: m.TotalRate != 0},
		Rate:            rateNumeric,
	}
}

func MapToDomainSeason(s sqlc_generated.DSeason) *movieentity.Season {
	return &movieentity.Season{
		Id:                s.ID,
		Name:              s.Name,
		Description:       s.Description.String,
		ThumbnailUrl:      s.ThumbnailUrl.String,
		ProgressionStatus: movieenum.ProgressionStatus(s.ProgressionStatus.String),
		AvailableStatus:   commonenum.AvailableStatus(s.AvailableStatus.String),
		ReleaseYear:       s.ReleaseYear.Int32,
		Order:             s.Order.Int32,
		TotalEpisode:      s.TotalEpisode.Int32,
		BaseEntity: commonentity.BaseEntity{
			CreatedBy: &s.CreatedBy.String,
			CreatedAt: &s.CreatedAt.Time,
			UpdatedBy: &s.UpdatedBy.String,
			UpdatedAt: &s.UpdatedAt.Time,
		},
	}
}

func MapToSQLModelSeason(s movieentity.Season) *sqlc_generated.DSeason {
	return &sqlc_generated.DSeason{
		ID:                s.Id,
		Name:              s.Name,
		Description:       pgtype.Text{String: s.Description, Valid: s.Description != ""},
		ThumbnailUrl:      pgtype.Text{String: s.ThumbnailUrl, Valid: s.ThumbnailUrl != ""},
		ProgressionStatus: pgtype.Text{String: s.ProgressionStatus.String(), Valid: s.ProgressionStatus != ""},
		AvailableStatus:   pgtype.Text{String: s.AvailableStatus.String(), Valid: s.AvailableStatus != ""},
		ReleaseYear:       pgtype.Int4{Int32: s.ReleaseYear, Valid: s.ReleaseYear != 0},
		Order:             pgtype.Int4{Int32: s.Order, Valid: s.Order != 0},
		TotalEpisode:      pgtype.Int4{Int32: s.TotalEpisode, Valid: s.TotalEpisode != 0},
		CreatedBy:         pgtype.Text{String: *s.CreatedBy, Valid: s.CreatedBy != nil},
		CreatedAt:         pgtype.Timestamp{Time: *s.CreatedAt, Valid: s.CreatedAt != nil},
		UpdatedBy:         pgtype.Text{String: *s.UpdatedBy, Valid: s.UpdatedBy != nil},
		UpdatedAt:         pgtype.Timestamp{Time: *s.UpdatedAt, Valid: s.UpdatedAt != nil},
	}
}

func MapToDomainEpisode(e sqlc_generated.DEpisode) *movieentity.Episode {
	return &movieentity.Episode{
		Id:               e.ID,
		Name:             e.Name,
		OriginalVideoUrl: e.OriginalVideoUrl.String,
		StreamVideoUrl:   e.StreamVideoUrl.String,
		BaseEntity: commonentity.BaseEntity{
			CreatedBy: &e.CreatedBy.String,
			CreatedAt: &e.CreatedAt.Time,
			UpdatedBy: &e.UpdatedBy.String,
			UpdatedAt: &e.UpdatedAt.Time,
		},
	}
}

func MapToSQLModelEpisode(e movieentity.Episode) *sqlc_generated.DEpisode {
	return &sqlc_generated.DEpisode{
		ID:               e.Id,
		Name:             e.Name,
		OriginalVideoUrl: pgtype.Text{String: e.OriginalVideoUrl, Valid: e.OriginalVideoUrl != ""},
		StreamVideoUrl:   pgtype.Text{String: e.StreamVideoUrl, Valid: e.StreamVideoUrl != ""},
	}
}
