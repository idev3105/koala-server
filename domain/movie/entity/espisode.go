package movieentity

import (
	commonentity "org.idev.koala/backend/domain/common/entity"
	commonenum "org.idev.koala/backend/domain/common/enum"
)

type Episode struct {
	commonentity.BaseEntity
	Id               int32                      `json:"id,omitempty"`
	Name             string                     `json:"name,omitempty"`
	ThumbnailUrl     string                     `json:"thumbnailUrl,omitempty"`
	OriginalVideoUrl string                     `json:"originalVideoUrl,omitempty"`
	StreamVideoUrl   string                     `json:"streamVideoUrl,omitempty"`
	AvailableStatus  commonenum.AvailableStatus `json:"availableStatus,omitempty"`
	Order            int32                      `json:"order,omitempty"`
	Rate             float32                    `json:"rate,omitempty"`
	TotalRate        int32                      `json:"totalRate,omitempty"`
	TotalViewer      int32                      `json:"totalViewer,omitempty"`
	TotalUpvote      int32                      `json:"totalUpvote,omitempty"`
	TotalDownvote    int32                      `json:"totalDownvote,omitempty"`
	Duration         int32                      `json:"duration,omitempty"`
}
