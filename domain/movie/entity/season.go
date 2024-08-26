package movieentity

import (
	commonentity "org.idev.koala/backend/domain/common/entity"
	commonenum "org.idev.koala/backend/domain/common/enum"
	movieenum "org.idev.koala/backend/domain/movie/enum"
)

type Season struct {
	commonentity.BaseEntity
	ProgressionStatus movieenum.ProgressionStatus `json:"progressionStatus,omitempty"`
	Description       string                      `json:"description,omitempty"`
	ThumbnailUrl      string                      `json:"thumbnailUrl,omitempty"`
	Name              string                      `json:"name,omitempty"`
	AvailableStatus   commonenum.AvailableStatus  `json:"availableStatus,omitempty"`
	Espiodes          []Episode                   `json:"espiodes,omitempty"`
	ReleaseYear       int32                       `json:"releaseYear,omitempty"`
	Rate              float32                     `json:"rate,omitempty"`
	TotalRate         int32                       `json:"totalRate,omitempty"`
	TotalViewer       int32                       `json:"totalViewer,omitempty"`
	TotalUpvote       int32                       `json:"totalUpvote,omitempty"`
	TotalDownvote     int32                       `json:"totalDownvote,omitempty"`
	Order             int32                       `json:"order,omitempty"`
	TotalEpisode      int32                       `json:"totalEpisode,omitempty"`
}
