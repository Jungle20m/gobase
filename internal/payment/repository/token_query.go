package repository

import (
	"context"
	"gobase/internal/payment/model"
)

func (r *Repo) GetTokenByID(ctx context.Context, id int) (*model.Token, error) {
	conn := r.getDB(ctx)

	var record model.Token
	sql := `
			SELECT id, user_id, client_id, id_token, access_token, refresh_token
			FROM user_token
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
