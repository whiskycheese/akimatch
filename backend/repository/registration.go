package repository

import (
	"database/sql"
	"backend/schema"
)

type RegistrationRepository struct {
	db *sql.DB
}

func NewRegistrationRepository(db *sql.DB) *RegistrationRepository {
	return &RegistrationRepository{db: db}
}

// 指定した user_id の登録済みレコード一覧を取得する
func (r *RegistrationRepository) GetByUserID(userID int) ([]schema.RegistrationResponse, error) {
	query := `
		SELECT 
			r.id, 
			r.user_id, 
			r.course_id, 
			r.day, 
			r.period,
			c.subject_name,
			COALESCE(c.room, '')
		FROM registrations r
		INNER JOIN courses c ON r.course_id = c.id
		WHERE r.user_id = $1
		ORDER BY r.day ASC, r.period ASC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []schema.RegistrationResponse
	for rows.Next() {
		var reg schema.RegistrationResponse
		if err := rows.Scan(
			&reg.ID,
			&reg.UserID,
			&reg.CourseID,
			&reg.Day,
			&reg.Period,
			&reg.SubjectName,
			&reg.Room,
		); err != nil {
			return nil, err
		}
		list = append(list, reg)
	}

	return list, nil
}

func (r *RegistrationRepository) Create(userID int, req schema.CreateRegistrationRequest) (*schema.RegistrationResponse, error) {
	query := `
		INSERT INTO registrations (user_id, course_id, day, period)
		VALUES ($1, $2, $3, $4)
		RETURNING id, user_id, course_id, day, period`

	var reg schema.RegistrationResponse
	err := r.db.QueryRow(query, userID, req.CourseID, req.Day, req.Period).Scan(
		&reg.ID,
		&reg.UserID,
		&reg.CourseID,
		&reg.Day,
		&reg.Period,
	)
	if err != nil {
		return nil, err
	}

	return &reg, nil
}

// 履修登録の削除
func (r *RegistrationRepository) Delete(userID, registrationID int) error {
	query := `DELETE FROM registrations WHERE id = $1 AND user_id = $2`
	
	result, err := r.db.Exec(query, registrationID, userID)
	if err != nil {
		return err
	}

	// 削除された行数を確認
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows // 削除対象が存在しなかった、または他のユーザーのデータ
	}

	return nil
}

// 指定したユーザー・曜日・時限に既に登録があるかチェックする
func (r *RegistrationRepository) Exists(userID int, day string, period int) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM registrations 
			WHERE user_id = $1 AND day = $2 AND period = $3
		)`

	var exists bool
	err := r.db.QueryRow(query, userID, day, period).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
