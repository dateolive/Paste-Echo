package middleware

import (
	"Paste-Echo/app/model/paste"
	"Paste-Echo/common"
	"regexp"
)

var (
	validLang  = []string{"plain", "cpp", "java", "python", "bash", "markdown", "json", "go"}
	keyPattern = regexp.MustCompile("^[0-9a-z]{8}$")
)

type CreateRequest struct {
	*paste.AbstructPaste

}

type CreateResponse struct {
	*common.Response
	Key string `json:"key" example:"abcdefgh"`
}

type GetResponse struct {
	*common.Response
	Lang    string `json:"lang" example:"plain"`
	Content string `json:"content" example:"Hello World!"`
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Validator(body *paste.AbstructPaste) *common.ErrorResponse {
	if body.Content == "" {
		return common.ErrEmptyContent // 内容为空，返回错误信息 "empty content"
	}
	if body.Lang == "" {
		return common.ErrEmptyLang // 语言类型为空，返回错误信息 "empty lang"
	}
	if !contains(validLang, body.Lang) {
		return common.ErrInvalidLang
	}


	return nil
}



func KeyValidator(key string) *common.ErrorResponse {
	if len(key) != 8 {
		return common.ErrInvalidKeyLength // key's length should at least 3 and at most 8
	}
	if flag := keyPattern.MatchString(key); !flag {
		return common.ErrInvalidKeyFormat
	}
	return nil
}