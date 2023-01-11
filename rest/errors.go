package rest

import (
	"errors"
	"strconv"
)

// RestError is a struct to separate errors with same messages and
// different codes or vise-versa.
type RestError struct {
	Message string
	Code    int
}

func (e *RestError) Error() string {
	return strconv.Itoa(e.Code) + ":" + e.Message
}

var (
	ErrRestNoLoginMeth  = errors.New("no login method")
	ErrRestNotLoggedIn  = errors.New("not logged in")
	ErrRestExpired      = errors.New("session expired")
	ErrInvalidPersid    = errors.New("persid must be a semicolon separated <factory:value>")
	ErrRestBadRequest   = errors.New("bad request")
	ErrRestUnathorized  = errors.New("unauthorized")
	ErrRestNotFound     = errors.New("not found")
	ErrRestConflict     = errors.New("conflict")
	ErrRestInternal     = errors.New("internal")
	ErrRestUploadFailed = errors.New("file upload failed")
)
