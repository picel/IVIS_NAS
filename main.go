package main

import (
	"net/http"

	loginCheck "./LoginCheck"
	signUp "./SignUp"
)

func main() {
	http.HandleFunc("/", http.FileServer(http.Dir("./")).ServeHTTP)
	http.HandleFunc("/process", signUp.SignUpProcess)
	http.HandleFunc("/loginCheck", loginCheck.LoginCheck)
	http.ListenAndServe(":8080", nil)
}
