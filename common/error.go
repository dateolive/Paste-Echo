package common

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrExpireSecondGreaterThanMonth   = New(http.StatusBadRequest, 3, "expire minute greater than a month")
	ErrExpireCountGreaterThanMaxCount = New(http.StatusBadRequest, 4, "expire count greater than max count")
	ErrEmptyContent                   = New(http.StatusBadRequest, 5, "empty content")
	ErrEmptyLang                      = New(http.StatusBadRequest, 6, "empty lang")
	ErrInvalidLang                    = New(http.StatusBadRequest, 7, "invalid lang")
	ErrWrongParamType                 = New(http.StatusBadRequest, 8, "wrong param type")
	ErrInvalidKeyLength               = New(http.StatusBadRequest, 9, "invalid key length")
	ErrInvalidKeyFormat               = New(http.StatusBadRequest, 10, "invalid key format")

	ErrFrequentRequest = New(http.StatusTooManyRequests, 1, "frequent request")

	ErrDeserialization = New(http.StatusInternalServerError, 1, "deserialization error")

	ErrUnauthorized = New(http.StatusUnauthorized, 1, "unauthorized")

	ErrWrongPassword = New(http.StatusForbidden, 1, "wrong password")

	ErrNoRouterFounded = New(http.StatusNotFound, 1, "no router founded")
	ErrRecordNotFound  = New(http.StatusNotFound, 2, "record not found")

	ErrQueryDBFailed = New(http.StatusInternalServerError, 1, "query from db failed")
	ErrSaveFailed    = New(http.StatusInternalServerError, 2, "save failed")
)

type ErrorResponse struct {
	*Response
	Message string `json:"message" example:"ok"`
}

func (response *ErrorResponse) Error() string {
	return response.Message
}

func (req *ErrorResponse) Abort(ctx echo.Context) error {
	return ctx.JSONPretty(req.GetHttpStatusCode(), req, " ")
}

func New(code int, index int, message string) *ErrorResponse {
	return &ErrorResponse{
		Response: &Response{
			Code: code*100 + index,
		},
		Message: message,
	}
}
