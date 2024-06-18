package service

import "errors"

var (
	ErrNoUpdates                   = errors.New("there are few arguments, at least one is needed")
	ErrUserNameExists        error = errors.New("user with this username already exists")
	ErrUserNamePassNotExists error = errors.New("user with this username or password not exists")
	ErrStudentIDExists       error = errors.New("such a student has already been registered")
	ErrTeacherIDExists       error = errors.New("such a teacher has already been registered")
	ErrHeadmanIDExists       error = errors.New("such a headman has already been registered")
)
