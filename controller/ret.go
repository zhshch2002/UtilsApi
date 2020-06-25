package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GinRet struct {
	ctx *gin.Context
}

type JsonResult struct {
	Code   int         `json:"code"`   // return code
	Msg    string      `json:"msg"`    // message
	Result interface{} `json:"result"` // data object
}

func Ret(ctx *gin.Context) *GinRet {
	return &GinRet{ctx: ctx}
}

func (s *GinRet) OK(d interface{}) {
	s.ctx.JSON(http.StatusOK, JsonResult{
		Code:   0,
		Msg:    "",
		Result: d,
	})
}

func (s *GinRet) Err(msg string) {
	s.ctx.JSON(http.StatusOK, JsonResult{
		Code:   -1,
		Msg:    msg,
		Result: nil,
	})
}

func (s *GinRet) ServerErr(err error) {
	log.Println(err)
	s.ctx.JSON(http.StatusOK, JsonResult{
		Code:   -1,
		Msg:    "Server error",
		Result: nil,
	})
}

func (s *GinRet) MissParam(msg string) {
	s.ctx.JSON(http.StatusOK, JsonResult{
		Code:   -2,
		Msg:    msg,
		Result: nil,
	})
}
