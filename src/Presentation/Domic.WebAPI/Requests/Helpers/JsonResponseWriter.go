package WebAPIRequestHelper

import (
	"Domic.Domain/Commons/Contracts"
	"Domic.WebAPI/DTOs"
	"net/http"
)

func WriteJsonResponse(serializer DomainCommonContract.ISerializer, w http.ResponseWriter, jsonResponse WebAPIDTO.JsonResponseDto) {
	stringifyResponse, err := serializer.Serialize(jsonResponse)

	if err != nil {
	}

	w.WriteHeader(jsonResponse.Code)
	w.Write([]byte(stringifyResponse))
}
