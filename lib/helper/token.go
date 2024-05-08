package helper

import (
	"time"
	"user-service/config"
	"user-service/domain/user/model"

	"github.com/golang-jwt/jwt"
)

var jwtSigningMethod = jwt.SigningMethodHS256

type JWTToken struct {
	Id          string `json:"id"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	jwt.StandardClaims
}

func GenerateTokenAndRefreshToken(login model.Users, cfg *config.TokenConfig) (tokenData model.UserToken, err error) {
	tokenData, err = GenerateJWT(login, cfg)
	if err != nil {
		return
	}

	refreshTokenData, err := GenerateRefresh(login, cfg)

	if err != nil {
		return
	}

	tokenData.RefreshToken = refreshTokenData.RefreshToken
	tokenData.RefreshTokenExpiredAt = refreshTokenData.RefreshTokenExpiredAt

	return
}

func GenerateToken(login model.Users, cfg *config.TokenConfig) (tokenData model.UserToken, err error) {
	tokenData, err = GenerateJWT(login, cfg)
	if err != nil {
		return
	}

	return
}

func GenerateJWT(user model.Users, cfg *config.TokenConfig) (tokenData model.UserToken, err error) {

	exp := time.Now().UTC().Add(time.Duration(cfg.Expiry) * time.Minute)
	claims := JWTToken{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().UTC().Unix(),
			ExpiresAt: exp.Unix(),
		},
		Id:          user.Id,
		FullName:    user.FullName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString([]byte(cfg.Key))
	if err != nil {
		return
	}

	tokenData = model.UserToken{
		Id:             user.Id,
		Token:          signedToken,
		TokenExpiredAt: exp,
	}

	return
}

func GenerateRefresh(user model.Users, cfg *config.TokenConfig) (refreshTokenData model.UserToken, err error) {

	exp := time.Now().UTC().Add(time.Duration(cfg.RefreshTokenExpiry) * time.Hour)
	claims := JWTToken{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().UTC().Unix(),
			ExpiresAt: exp.Unix(),
		},
		Id: user.Id,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString([]byte(cfg.Key))
	if err != nil {
		return
	}

	refreshTokenData = model.UserToken{
		RefreshToken:          signedToken,
		RefreshTokenExpiredAt: exp,
	}

	return
}

func ParseToken(tokenString string, cfg *config.Config) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, jwt.ErrInvalidKeyType
		}

		return []byte(cfg.Token.Key), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return token, err
}

func GetDataFromToken(token string, cfg *config.Config) (userData model.Users, err error) {

	resToken, err := ParseToken(token, cfg)
	if err != nil {
		return
	}

	claims := resToken.Claims.(jwt.MapClaims)

	userData.Id = claims["id"].(string)

	return
}
