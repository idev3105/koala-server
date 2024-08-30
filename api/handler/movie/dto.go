package moviehandler

type CreateMovieRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	// TODO: Add more
} // @name CreateMovieRequest

type MovieDto struct {
	Id           int32  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	// TODO: Add more
} // @name MovieDto

type VoteMovieRequest struct {
	VoteType string `json:"voteType"`
} // @name VoteMovieRequest
