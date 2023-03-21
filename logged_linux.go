package users

// #include "cgo-utmp.h"
import "C"
import (
	"errors"
	"fmt"
)

type LoggedUser struct {
	Username  string
	Tty       string
	LoginTime string
}

func ListLogged() ([]*LoggedUser, error) {
	fp := C.cgo_utmp_open()
	if fp == nil {
		return nil, errors.New("error opening utmp file")
	}
	defer C.cgo_utmp_close(fp)

	var ut C.struct_utmp
	var loggedUsers []*LoggedUser

	for C.cgo_utmp_next(fp, &ut) == 1 {
		// 7 为 user session
		if ut.ut_type != 7 {
			continue
		}

		loggedUser := &LoggedUser{
			Username:  C.GoString(&ut.ut_user[0]),
			Tty:       C.GoString(&ut.ut_line[0]),
			LoginTime: fmt.Sprintf("%d", ut.ut_tv.tv_sec),
		}
		loggedUsers = append(loggedUsers, loggedUser)
	}

	return loggedUsers, nil
}
