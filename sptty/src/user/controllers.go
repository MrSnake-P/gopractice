package user

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/linshenqi/sptty"
)

type Result struct {
	status	int
	msg	string
}

func (s *Service) helloworld(ctx iris.Context) {
	//
	//err := sptty.SimpleResponse(ctx, iris.StatusOK, "成功")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	user :=s.getUser(1)
	fmt.Println("user:",user)
	_ = sptty.SimpleResponse(ctx, iris.StatusOK, user)
}

func (s *Service) addUser(ctx iris.Context) {
	err := s.createUser("daffy", 23)
	if err != nil {
		return
	}
	result := Result{status: 200, msg: "成功"}
	_ = sptty.SimpleResponse(ctx, iris.StatusOK, &result)
}