package user

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/linshenqi/sptty"
)

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