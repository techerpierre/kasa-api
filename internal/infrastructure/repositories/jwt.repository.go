package repositories

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/techerpierre/kasa-api/internal/domain/entities"
)

type JwtRepository struct{}

func CreateJwtRepository() *JwtRepository {
	return &JwtRepository{}
}

func (*JwtRepository) Sign(payloads entities.Payloads) (string, *entities.Exception) {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":               payloads.ID,
			"email":            payloads.Email,
			"firstname":        payloads.Firstname,
			"lastname":         payloads.Lastname,
			"authorizationsId": payloads.AuthorizationsID,
			"exp":              time.Now().Add(time.Hour * 168).Unix(),
		},
	)
	stringifyToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	return stringifyToken, nil
}

func (r *JwtRepository) Verify(tokenString string) (entities.Payloads, *entities.Exception) {
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return secret, nil
	})

	if err != nil {
		return entities.Payloads{}, entities.CreateException(
			entities.ExceptionCode_Unauthorized,
			entities.ExceptionMessage_Unauthorized,
		)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, found := claims["exp"].(int64); found {
			if time.Now().Unix() > exp {
				return entities.Payloads{}, entities.CreateException(
					entities.ExceptionCode_Unauthorized,
					entities.ExceptionMessage_Unauthorized,
				)
			}
			var payloads entities.Payloads
			r.mapClaimsInPayloads(claims, &payloads)
			return payloads, nil
		}
	}

	return entities.Payloads{}, entities.CreateException(
		entities.ExceptionCode_Unauthorized,
		entities.ExceptionMessage_Unauthorized,
	)
}

func (*JwtRepository) mapClaimsInPayloads(claims jwt.MapClaims, payloads *entities.Payloads) {
	for key, value := range claims {
		switch key {
		case "id":
			payloads.ID = value.(string)
		case "email":
			payloads.Email = value.(string)
		case "firstname":
			payloads.Firstname = value.(string)
		case "lastname":
			payloads.Lastname = value.(string)
		case "AuthorizationsID":
			payloads.AuthorizationsID = value.(string)
		}
	}
}
