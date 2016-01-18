package main

import (
	"fmt"
	"net/http"

	"github.com/jamal/flannel"
)

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	login := &LoginForm{}
	flannel.DecodeForm(r, login)

	flannel.Write(w, http.StatusForbidden, fmt.Sprintf("User %s does not exist", login.Username))
}
