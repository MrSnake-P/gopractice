package user

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/linshenqi/sptty"
	"sptty/base"
	"strconv"
)

type Result struct {
	Status int         `json:"status""`
	Msg    string      `json"msg"`
	Data   interface{} `json:"data"`
}

func (s *Service) helloworld(ctx iris.Context) {
	//
	//err := sptty.SimpleResponse(ctx, iris.StatusOK, "成功")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	user := s.getUser(1)
	fmt.Println("user:", user)
	_ = sptty.SimpleResponse(ctx, iris.StatusOK, user)
}

func (s *Service) addUser(ctx iris.Context) {

	user := base.User{}
	if err := ctx.ReadJSON(&user); err != nil {
		_ = sptty.SimpleResponse(ctx, iris.StatusOK, Result{Status: 202, Msg: "失败"})
		return
	}
	err := s.createUser(&user)
	if err != nil {
		return
	}
	result := Result{Status: 200, Msg: "成功"}
	_ = sptty.SimpleResponse(ctx, iris.StatusOK, &result)
}

func (s *Service) selUser(ctx iris.Context) {
	id := ctx.Params().Get("id")
	userId, _ := strconv.Atoi(id)
	response := s.selUserDetail(userId)
	_ = sptty.SimpleResponse(ctx, iris.StatusOK, Result{Status: 200, Data: response})
}

func (s *Service) verifyEmail(ctx iris.Context) {
	//email := base.EmailStatus{}
	id := ctx.Params().Get("id")
	userId, _ := strconv.Atoi(id)
	response := s.verifyEmailCode(userId)
	//email.VerifyCode = response.VerifyCode
	//email.VerifyStatus = response.VerifyStatus
	_ = sptty.SimpleResponse(ctx, iris.StatusOK, Result{Status: 200, Data: response})
}

func (s *Service) verifyStatus(ctx iris.Context) {
	id := 	ctx.URLParam("id")
	userId, _ := strconv.Atoi(id)
	verifyCode := ctx.URLParam("verify_code")
	fmt.Println(userId, verifyCode)
	status, err := s.verifyemail(userId, verifyCode)
	if err != nil {
		_ = sptty.SimpleResponse(ctx, iris.StatusOK, Result{Status: 202, Msg: "未验证"})
		return
	}
	_ = sptty.SimpleResponse(ctx, iris.StatusOK, Result{Status: 200, Msg: "验证成功", Data: status})
}