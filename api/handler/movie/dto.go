package moviehandler

import "org.idev.koala/backend/common/message"

type CreateMovieRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	// TODO: Add more
} // @name CreateMovieRequest

type MovieDto struct {
	MovieID      string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	// TODO: Add more
} // @name MovieDto

type VoteMovieRequest struct {
	VoteType string `json:"voteType"`
} // @name VoteMovieRequest

type UpdateMovieRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ThumbnailUrl string `json:"thumbnailUrl"`
} // @name UpdateMovieRequest

type MovieVotedMessage struct {
	message.BaseMessage
	MovieID  string `json:"movie_id"`
	UserID   string `json:"user_id"`
	VoteType string `json:"vote_type"`
} // @name MovieVotedMessage
