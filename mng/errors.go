package mng

import "gopkg.in/mgo.v2"

func IsErrNotFound(err error) bool {
	return err == mgo.ErrNotFound
}
