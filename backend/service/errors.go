package service

import "errors"

var (
	// 時間割（Registration）関連エラー
	ErrAlreadyRegistered = errors.New("a course is already registered in this time slot")
	
	// 友達（Friendship）関連エラー
	ErrCannotAddSelf = errors.New("cannot add yourself as a friend")
	ErrAlreadyFriend = errors.New("user is already your friend")
)
