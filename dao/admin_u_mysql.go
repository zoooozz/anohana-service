package dao

import (
	"anohana-service/model"
	"database/sql"
	"github.com/donnie4w/go-logger/logger"
)

const (
	_findLoginSQL = "select id,ctime,mtime,username,email,type,password,photo,ip,phone,state from admin_u where phone = ?"
)

func (d *Dao) GetUserByPhone(phone string) (r *model.AdminUModel, err error) {
	r = &model.AdminUModel{}
	row := d.db.QueryRow(_findLoginSQL, phone)
	if err = row.Scan(
		&r.Id,
		&r.Ctime,
		&r.Mtime,
		&r.Username,
		&r.Email,
		&r.Type,
		&r.Password,
		&r.Photo,
		&r.Ip,
		&r.Phone,
		&r.State,
	); err != nil {
		if err == sql.ErrNoRows {
			r = nil
			err = nil
		}
		logger.Info(err)
	}
	return
}

// func (d *Dao) SetByUserPhone(phone string) (id int64, err error) {
// 	model, err := d.db.Prepare(_insertUserSQL)
// 	if err != nil {
// 		return
// 	}
// 	result, err := model.Exec(phone)
// 	id, err = result.LastInsertId()

// 	if err != nil {
// 		return
// 	}
// 	return
// }

// func (d *Dao) UpdateLoginPass(phone, password string) (err error) {

// 	model, err := d.db.Prepare(_editUserPassSQL)
// 	if err != nil {
// 		return
// 	}
// 	_, err = model.Exec(password, phone)
// 	if err != nil {
// 		return
// 	}
// 	return nil
// }

// func (d *Dao) UpdateUserInfo(phone, username, photo string) (err error) {
// 	var (
// 		sqls   string
// 		params string
// 	)

// 	if username != "" {
// 		sqls = _editUserNameSQL
// 		params = username
// 	}

// 	if photo != "" {
// 		sqls = _editUserPhotoSQL
// 		params = photo

// 	}
// 	model, err := d.db.Prepare(sqls)

// 	if err != nil {
// 		return
// 	}
// 	_, err = model.Exec(params, phone)
// 	if err != nil {
// 		logger.Info(err)
// 		return
// 	}
// 	return nil
// }
