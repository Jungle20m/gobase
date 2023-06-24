package repository

import (
	"context"
	"gobase/internal/payment/model"
)

func (r *Repo) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	conn := r.getDB(ctx)

	var record model.User
	sql := `
			SELECT id, user_name, phone_number, email, password, password_salt, password_hash_algorithms, user_status, registration_time, create_time, update_time
			FROM user_account
			WHERE id=?
   		   `
	err := conn.Raw(sql, id).First(&record).Error
	if err != nil {
		//if errors.Is(err, gorm.ErrRecordNotFound) {
		//	return nil, common.ErrRecordNotFound
		//}
		return nil, err
	}
	return &record, nil

}
