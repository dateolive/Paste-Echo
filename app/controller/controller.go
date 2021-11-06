package controller

import (
	"Paste-Echo/app/middleware"
	"Paste-Echo/app/model"
	"Paste-Echo/app/model/paste"
	"Paste-Echo/app/service"
	"Paste-Echo/common"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
)


func Index(ctx echo.Context)error{
	return ctx.JSON(http.StatusOK,"user")
}

func Login(ctx echo.Context)error{
	user:=new(model.User)
	if err := ctx.Bind(user); err!=nil{
		return err
	}
	fmt.Println(user)

	return ctx.JSONPretty(http.StatusOK,user," ")
}

func GetUserByUserId(ctx echo.Context)error{

	userid,convErr:=strconv.Atoi(ctx.Param("userid"))

	user,err:= service.UserDao.GetUserByUserId(userid)

	if err!=nil||convErr!=nil{
		return err
	}
	return ctx.JSONPretty(http.StatusOK,user," ")

}

func CreateUser(ctx echo.Context)error{
	user:=new(model.User)
	if err :=ctx.Bind(user);err!=nil{
		return err
	}
	if createerr:=service.UserDao.Save(user);createerr!=nil{
		return createerr
	}

	return ctx.JSONPretty(http.StatusOK,user," ")
}

//创建复制贴——永久性帖子
func CreateForeverPaste(ctx echo.Context)error{
	body:=&paste.AbstructPaste{
		ClientIP:ctx.Request().RemoteAddr,
		CreatedAt: time.Now(),
	}
	if err:=ctx.Bind(&body);err!=nil{
		return common.ErrWrongParamType.Abort(ctx)
	}

	if err := middleware.Validator(body); err != nil {
		return err.Abort(ctx)
	}



	paste:=&paste.LongTimePaste{
		AbstructPaste:body,
	}
	if err:=paste.Save();err!=nil{
		return common.ErrSaveFailed.Abort(ctx)
	}
	return common.JSON(ctx,middleware.CreateResponse{
		Response:&common.Response{
			Code: http.StatusCreated,
		},
		Key: paste.GetKey(),
	})
}

func GetLongTimePasteByKey(ctx echo.Context)error{
	key:=strings.ToLower(ctx.Param("key"))

	if err:=middleware.KeyValidator(key);err!=nil{
		return err.Abort(ctx)
	}

	abstructPaste:=&paste.AbstructPaste{Key: key}
	paste:=paste.LongTimePaste{
		AbstructPaste:abstructPaste,
	}
	if err:=paste.Get(ctx.QueryParam("password"));err!=nil{
		var errorResponse *common.ErrorResponse
		switch err {
		case gorm.ErrRecordNotFound:
			errorResponse = common.ErrRecordNotFound
		case common.ErrWrongPassword:
			errorResponse = err.(*common.ErrorResponse)
		default:
			errorResponse = common.ErrQueryDBFailed
		}

		return errorResponse.Abort(ctx)

	}

	return common.JSON(ctx,middleware.GetResponse{
		Response:&common.Response{
			Code: http.StatusOK,
		},
		Lang: paste.GetLang(),
		Content: paste.GetContent(),
	})
}

func CreateExpirePaste(ctx echo.Context)error{
	body:=&paste.AbstructPaste{
		ClientIP:ctx.Request().RemoteAddr,
		CreatedAt: time.Now(),
	}
	paste:=&paste.ExpireTimePaste{
		AbstructPaste:body,
	}

	if err:=ctx.Bind(&paste);err!=nil{
		return common.ErrWrongParamType.Abort(ctx)
	}


	if err := middleware.Validator(body); err != nil {
		return err.Abort(ctx)
	}



	if err:=paste.Save();err!=nil{
		return common.ErrSaveFailed.Abort(ctx)
	}
	return common.JSON(ctx,middleware.CreateResponse{
		Response:&common.Response{
			Code: http.StatusCreated,
		},
		Key: paste.GetKey(),
	})
}

func GetExpirePasteByKey(ctx echo.Context)error{
	key:=strings.ToLower(ctx.Param("key"))

	if err:=middleware.KeyValidator(key);err!=nil{
		return err.Abort(ctx)
	}

	abstructPaste:=&paste.AbstructPaste{Key: key}
	paste:=paste.ExpireTimePaste{
		AbstructPaste:abstructPaste,
	}
	if err:=paste.Get(ctx.QueryParam("password"));err!=nil{
		var errorResponse *common.ErrorResponse
		switch err {
		case gorm.ErrRecordNotFound:
			errorResponse = common.ErrRecordNotFound
		case common.ErrWrongPassword:
			errorResponse = err.(*common.ErrorResponse)
		default:
			errorResponse = common.ErrQueryDBFailed
		}

		return errorResponse.Abort(ctx)

	}

	return common.JSON(ctx,middleware.GetResponse{
		Response:&common.Response{
			Code: http.StatusOK,
		},
		Lang: paste.GetLang(),
		Content: paste.GetContent(),
	})
}