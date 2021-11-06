package common

import "github.com/labstack/echo/v4"

type Response struct {
	Code int `json:"code" example:"200"`
}

type IResponse interface {
	GetHttpStatusCode() int
}

func (req *Response)GetHttpStatusCode() int{
	if req.Code >1000{
		return req.Code/100
	}
	return req.Code
}


func JSON(ctx echo.Context,req IResponse)error{
	return ctx.JSONPretty(req.GetHttpStatusCode(),req," ")
}

func NotFoundHandler(ctx echo.Context)error{
	return ErrNoRouterFounded.Abort(ctx)
}