package repository

import (
	"database/sql"
	"backend/schema"
)

type FriendshipRepository struct {
	db *sql.DB
}

func NewFriendshipRepository(db *sql.DB) *FriendshipRepository {
	return &FriendshipRepository{db: db}
}

// 既に友達か確認
func (r *FriendshipRepository) Exists(userID, friendID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM friendships WHERE user_id = $1 AND friend_id = $2)`
	var exists bool
	err := r.db.QueryRow(query, userID, friendID).Scan(&exists)
	return exists, err
}

// 友達追加（相互登録）
func (r *FriendshipRepository) Create(userID, friendID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO friendships (user_id, friend_id) VALUES ($1, $2), ($2, $1)`
	_, err = tx.Exec(query, userID, friendID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// 友達一覧を取得
func (r *FriendshipRepository) GetFriendsByUserID(userID int) ([]schema.FriendResponse, error) {
	query := `
		SELECT u.id, u.name, f.created_at
		FROM friendships f
		INNER JOIN users u ON f.friend_id = u.id
		WHERE f.user_id = $1
		ORDER BY f.created_at DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []schema.FriendResponse
	for rows.Next() {
		var f schema.FriendResponse
		if err := rows.Scan(&f.ID, &f.Name, &f.CreatedAt); err != nil {
			return nil, err
		}
		friends = append(friends, f)
	}

	return friends, nil
}
