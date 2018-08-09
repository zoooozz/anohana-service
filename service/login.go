package service

import (
	"anohana-service/common"
	"anohana-service/model"
	"encoding/json"
)

func (s *service) PassLogin(phone, password string) (r *model.AdminUModel, token string, err error) {

	r, err = s.dao.GetUserByPhone(phone)
	if r == nil {
		return
	}
	password = common.PassProduction(phone, password)
	if password == r.Password {
		r.Password = ""
		if json, err := json.Marshal(r); err == nil {
			s.dao.SetByUserRedis(phone, string(json))
		}
		token = common.TokenProduction(phone)
	}
	return
}
