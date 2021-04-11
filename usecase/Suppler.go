package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)

var (
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)


type Unit struct {
	Cid	string 
	

}

type Session interface {

}

type Session struct {

}

func (u *Unit) qr_encode() (, error) {

}

func (u *Unit) Trace_Unit() (, error) {
	
}

func (s *Session) Authenticate() () {

}








func main() {
}
