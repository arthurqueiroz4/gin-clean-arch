package internal

import (
	"fmt"
	"gin-clean-arch/domain"
	"gin-clean-arch/models"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

func CreateAccessToken(employee *models.Employee, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &domain.JwtCustomClaims{
		Name: employee.Name,
		ID:   strconv.Itoa(int(employee.Model.ID)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return accessToken, err
}

func CreateRefreshToken(employee *models.Employee, secret string, expiry int) (refreshToken string, err error) {
	claimsRefresh := &domain.JwtCustomRefreshClaims{
		ID: strconv.Itoa(int(employee.Model.ID)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	refreshToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected siging method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	return claims["id"].(string), nil
}
