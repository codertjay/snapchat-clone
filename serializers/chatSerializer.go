package serializers

import "github.com/google/uuid"

type AcceptRequestSerializer struct {
	ID       uuid.UUID `json:"id"`
	Accepted bool      `json:"accepted"`
}
