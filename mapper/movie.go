package mapper

import (
	"log"

	"github.com/jackc/pgx/v5/pgtype"
	commonenum "org.idev.koala/backend/domain/common/enum"
	movieentity "org.idev.koala/backend/domain/movie/entity"
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
