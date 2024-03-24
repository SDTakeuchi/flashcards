package auth

import (
	"fmt"

	"github.com/SDTakeuchi/go/src/flashcards/domain/model/auth"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ClaimsFromPayload(p auth.Payload) jwt.MapClaims {
	return jwt.MapClaims{
		"id":   p.ID(),
		"type": p.TokenType(),
		"sub":  p.UserID(),
		"exp":  p.ExpiresAt().Unix(),
		"iat":  p.IssuedAt().Unix(),
	}
}

func PayloadFromClaims(claims jwt.MapClaims) (*auth.Payload, error) {
	// id
	fetchedPayloadID, ok := claims["id"]
	if !ok {
		return nil, fmt.Errorf("payload id not found in claims")
	}
	payloadIDString, ok := fetchedPayloadID.(string)
	if !ok {
		return nil, fmt.Errorf("invalid payload id")
	}
	payloadID, err := uuid.Parse(payloadIDString)
	if err != nil {
		return nil, err
	}

	// token type
	fetchedTokenType, ok := claims["type"]
	if !ok {
		return nil, fmt.Errorf("token type not found in claims")
	}
	assertedTokenType, ok := fetchedTokenType.(float64)
	if !ok {
		return nil, fmt.Errorf("invalid token type value: %v, in type of: %T", fetchedTokenType, fetchedTokenType)
	}
	tokenType := auth.TokenType(int(assertedTokenType))

	// user id
	fetchedUserID, ok := claims["sub"]
	if !ok {
		return nil, fmt.Errorf("user id not found in claims")
	}
	userIDString, ok := fetchedUserID.(string)
	if !ok {
		return nil, fmt.Errorf("user id did not match")
	}
	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return nil, err
	}

	// issued at
	issuedAt, err := claims.GetIssuedAt()
	if err != nil {
		return nil, err
	}

	// expires at
	expiresAt, err := claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}
	return auth.PayloadFromPersistance(
		payloadID,
		tokenType,
		userID,
		issuedAt.Time,
		expiresAt.Time,
	), nil
}
