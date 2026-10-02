package service

import "errors"

var ErrAlreadyRegistered = errors.New("a course is already registered in this time slot")
