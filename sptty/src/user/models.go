package user

import (
	"encoding/json"
	"fmt"
	"github.com/linshenqi/sptty"
	"sptty/base"
)

func (s *Service) getUser(id int) *base.User {
	u := base.User{}
	result := s.db.Where("id = ?", id)
	result.Last(&u)
	//if err:=result.Last(&u).Error;err!=nil{
	//	fmt.Println("222")
	//	return &u
	//}
	//fmt.Println("123")
	return &u
}

func (s *Service) createUser(user *base.User) error {
	//u := base.User{Name: name, Age: age}
	userJson, _ := json.Marshal(user)
	sptty.Log(sptty.DebugLevel, "data"+string(userJson), s.ServiceName())
	s.db.Create(user)
	return nil
}

func (s *Service) selUserDetail(id int) *base.User {
	u := base.User{}
	result := s.db.Where("id = ?", id)
	if err := result.Last(&u).Error; err != nil {
		sptty.Log(sptty.DebugLevel, "msg: "+err.Error(), s.ServiceName())
		return &u
	}
	fmt.Println(u)
	return &u
}

func (s *Service) verifyEmailCode(id int) *base.EmailStatus {
	u := base.User{}
	e := base.EmailStatus{}
	result := s.db.Where("id = ?", id)
	result.Last(&u).Scan(&e)
	return &e
	//fmt.Println(e)
	//if err := result.Last(&u).Error; err != nil {
	//	sptty.Log(sptty.DebugLevel, "msg: "+err.Error(), s.ServiceName())
	//	return &u
	//}
	//return &u
}

func (s *Service) verifyEmailStatus(id int, code string) (bool, error) {
	u := base.User{}
	result := s.db.Where("id = ? and verify_status = ?", id, false)
	if err := result.Last(&u).Error; err != nil {
		sptty.Log(sptty.DebugLevel, "msg: "+err.Error(), s.ServiceName())
		return false, err
	}
	if u.VerifyCode == code {
		fmt.Println(u)
		u.VerifyStatus = true
		s.db.Save(&u)
		return true, nil
	}

	return false, fmt.Errorf("验证失败")
}
