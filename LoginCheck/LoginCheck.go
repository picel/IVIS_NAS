package loginCheck

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func LoginCheck(w http.ResponseWriter, r *http.Request) {

	// get id, password from post data
	id := r.FormValue("id")
	pw := r.FormValue("pw")

	// prevent sql injection
	if strings.Contains(id, "'") || strings.Contains(pw, "'") {
		http.Error(w, "Login fail", http.StatusUnauthorized)
		return
	}

	// if id or pw contains space or sudo command, return error
	if id == " " || pw == " " || id == "sudo" || pw == "sudo" {
		// Sign up fail with 401
		http.Error(w, "Sign up fail", http.StatusUnauthorized)
		return
	}

	// get /etc/shadow by id
	cmd := exec.Command("sudo", "cat", "/etc/shadow", "|", "grep", "\""+id+"\"")
	out, _ := cmd.Output()

	var shadow []string
	// find id in /etc/shadow
	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(line, id) {
			shadow = strings.Split(line, "$")
		}
	}

	// if id is not in /etc/shadow
	if len(shadow) == 0 || shadow[0] == "" {
		http.Error(w, "Login fail", http.StatusUnauthorized)
		return
	}

	salt := shadow[2]
	hashed := strings.Split(shadow[3], ":")[0]

	// get hashed password
	cmd = exec.Command("openssl", "passwd", "-6", "-salt", salt, pw)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out), pw)

	realPw := strings.Split(string(out), "$")[3]
	realPw = strings.Split(realPw, "\n")[0]

	// compare hashed password
	if hashed == realPw {
		w.WriteHeader(http.StatusOK)
	} else {
		// login fail with 401
		http.Error(w, "Login fail", http.StatusUnauthorized)
	}
}
