package InfrastructureConcrete

import (
	"Domic.Domain/Commons/DTOs"
	"errors"
	"github.com/golang-jwt/jwt"
	"maps"
)

type JsonWebToken struct {
}

func (jsonWebToken *JsonWebToken) Generate(claims map[string]interface{}) DomainCommonDTO.Result[string] {

	jwtClaims := jwt.MapClaims{}

	//todo: should be readed from config json file
	secretKey := "3137511375313753"

	maps.Copy(claims, jwtClaims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return DomainCommonDTO.Result[string]{
			Error:  err,
			Result: "",
		}
	}

	return DomainCommonDTO.Result[string]{
		Error:  nil,
		Result: tokenString,
	}

}

func (jsonWebToken *JsonWebToken) GetClaims(tokenString string) DomainCommonDTO.Result[map[string]interface{}] {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return "3137511375313753", nil
	})

	if err != nil {

	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return DomainCommonDTO.Result[map[string]interface{}]{
			Error:  nil,
			Result: claims,
		}
	}

	return DomainCommonDTO.Result[map[string]interface{}]{
		Error:  nil,
		Result: nil,
	}
}

func (jsonWebToken *JsonWebToken) Verify(tokenString string) DomainCommonDTO.Result[bool] {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return "3137511375313753", nil
	})

	if err != nil {
		return DomainCommonDTO.Result[bool]{
			Error:  err,
			Result: false,
		}
	}

	if !token.Valid {
		return DomainCommonDTO.Result[bool]{
			Error:  errors.New("token is invalid"),
			Result: false,
		}
	}

	return DomainCommonDTO.Result[bool]{
		Error:  nil,
		Result: true,
	}
}

func NewJsonWebToken() *JsonWebToken {
	return &JsonWebToken{}
}
