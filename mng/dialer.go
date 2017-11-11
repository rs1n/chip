package mng

import "gopkg.in/mgo.v2"

func MustDial(dialInfo *mgo.DialInfo) *mgo.Session {
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	return session
}
