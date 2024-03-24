package auth

import (
	"time"

	"github.com/google/uuid"
)

// an interface for managing tokens
type TokenIssuer interface {
	// creates a new token for a specific usernmame and duration
	// *Payload is neccessary to refresh token (cuz it needs token ID)
	Create(userID uuid.UUID, tokenType TokenType, duration time.Duration) (string, *Payload, error)
	Verify(token string) (*Payload, error)
}
