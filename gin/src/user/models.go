package user

import (
	"gin/gin/base"
)

func (s *Service) getUser(id int) *base.User {
	u := base.User{}
	result := s.db.Where("id = ?",id)
	result.Last(&u)
	//if err:=result.Last(&u).Error;err!=nil{
	//	fmt.Println("222")
	//	return &u
	//}
	//fmt.Println("123")
	return &u
}