package message

import "time"

type BaseMessage struct {
	CreatedAt time.Time `json:"created_at"`
}
