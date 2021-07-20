package user

import (
	"github.com/jinzhu/gorm"
	"github.com/linshenqi/sptty"
	"sptty/base"
)

type Service struct {
	sptty.BaseService

	db *gorm.DB
}

func (s *Service) Init(app sptty.ISptty) error {
	s.db = app.Model().(*sptty.ModelService).DB()

	app.AddRoute("GET", "/hello", s.helloworld)
	app.AddRoute("POST", "/add", s.addUser)
	app.AddRoute("GET", "/sel/{id:string}", s.selUser)
	app.AddRoute("GET", "/verify/{id:string}", s.verifyEmail)
	app.AddRoute("GET", "/verify/{id:string}", s.verifyStatus)

	return nil
}

func (s *Service) verifyemail(id int, code string) (bool, error) {
	_, err := s.verifyEmailStatus(id, code)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) ServiceName() string {
	return base.ServiceUser
}
