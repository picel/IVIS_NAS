package SignUp

import (
	"fmt"
	"net/http"
	"os/exec"

	key "../key"
	httpsocket "./httpsocket"
)

func SignUpProcess(w http.ResponseWriter, r *http.Request) {
	// get id, password from post data
	id := r.FormValue("id")
	pw := r.FormValue("pass")

	// sql injection check
	if id == "" || pw == "" {
		// Sign up fail with 401
		http.Error(w, "Sign up fail", http.StatusUnauthorized)
		return
	}

	// if id or pw contains space or sudo command, return error
	if id == " " || pw == " " || id == "sudo" || pw == "sudo" {
		// Sign up fail with 401
		http.Error(w, "Sign up fail", http.StatusUnauthorized)
		return
	}

	// make command
	//tmp := "sudo" + " " + "useradd" + " " + id + " " + pw
	// execute command
	cmd := exec.Command("sh", "usradd.sh", id, pw)
	// check error
	err := cmd.Run()
	if err != nil {
		// Sign up fail with 401
		http.Error(w, "Sign up fail", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	// run mail useradd function with id, pw
	UserAdd(id, pw)

	// redirect to SignUp page
	http.Redirect(w, r, "/SignUp", http.StatusFound)
}

func UserAdd(username string, password string) {
	// set mail server account
	var serverid string = key.MxRouteID
	var serverpw string = key.MxRoutePW
	var serverURL string = key.MxRouteURL

	// setup to add mail user
	socket := httpsocket.New()
	socket.Connect(serverURL)
	socket.SetAuth(serverid, serverpw)
	socket.SetMethod("POST")
	socket.SetPath("/CMD_API_POP")
	socket.Query("action", "create")
	socket.Query("domain", "ivis.dev") // domain setting
	socket.Query("user", username)
	socket.Query("passwd", password)
	socket.Query("passwd2", password)
	socket.Query("quota", "0")    // mail box size, 0 is unlimited
	socket.Query("limit", "7200") // limit of sending mail. don't change this value
	socket.Query("create", "Create")

	socket.Send()

	if socket.StatusCode() != 200 {
		fmt.Println("Error: " + socket.Status())
		return
	}

	fmt.Println("scoket body: " + socket.Body())
	fmt.Println("Success: " + socket.Status())
}
