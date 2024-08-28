package movieentity

import (
	commonentity "org.idev.koala/backend/domain/common/entity"
	commonenum "org.idev.koala/backend/domain/common/enum"
)

type Movie struct {
	commonentity.BaseEntity
	Id              int32                      `json:"id,omitempty"`
	ItemStatus      string                     `json:"itemStatus,omitempty"`
	Name            string                     `json:"name,omitempty"`
	Description     string                     `json:"description,omitempty"`
	ThumbnailUrl    string                     `json:"thumbnailUrl,omitempty"`
	AvailableStatus commonenum.AvailableStatus `json:"availableStatus,omitempty"`
	Types           []string                   `json:"types,omitempty"`
	Seasons         []Season                   `json:"seasons,omitempty"`
	ReleaseYear     int32                      `json:"releaseYear,omitempty"`
	TotalViewer     int32                      `json:"totalViewer,omitempty"`
	TotalUpvote     int32                      `json:"totalUpvote,omitempty"`
	TotalDownvote   int32                      `json:"totalDownvote,omitempty"`
	TotalRate       int32                      `json:"totalRate,omitempty"`
	Rate            float32                    `json:"rate,omitempty"`
}
