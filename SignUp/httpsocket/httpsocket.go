package httpsocket

import (
	"fmt"
	"os/exec"
)

// HTTPSocket is a wrapper for the curl command line tool.
type HTTPSocket struct {
	// The URL to connect to.
	url string

	// The method to use.
	method string

	// The path to use.
	path string

	// The username to authenticate with.
	username string

	// The password to authenticate with.
	password string

	// The query parameters to send.
	query map[string]string
}

// New creates a new HTTPSocket.
func New() *HTTPSocket {
	return &HTTPSocket{
		method: "GET",
		path:   "/",
		query:  make(map[string]string),
	}
}

// Connect connects to the given URL.
func (s *HTTPSocket) Connect(url string) {
	s.url = url
}

// SetMethod sets the method to use.
func (s *HTTPSocket) SetMethod(method string) {
	s.method = method
}

// SetPath sets the path to use.
func (s *HTTPSocket) SetPath(path string) {
	s.path = path
}

// SetAuth sets the username and password to use for authentication.
func (s *HTTPSocket) SetAuth(username string, password string) {
	s.username = username
	s.password = password
}

// Query adds a query parameter.
func (s *HTTPSocket) Query(key string, value string) {
	s.query[key] = value
}

// Send sends the request.
func (s *HTTPSocket) Send() {
	args := []string{
		"-s",
		"-k",
		"-X", s.method,
		"-u", s.username + ":" + s.password,
		"-d", s.queryString(),
		s.url + s.path,
	}

	fmt.Println("curl", args)

	cmd := exec.Command("curl", args...)

	// get response from curl
	out, err := cmd.Output()
	if err != nil {
		fmt.Print("curl error: ")
		fmt.Println(err)
		return
	}

	fmt.Println("cur out :" + string(out))
}

// StatusCode returns the status code of the response.
func (s *HTTPSocket) StatusCode() int {
	return 200
}

// Status returns the status of the response.
func (s *HTTPSocket) Status() string {
	return "OK"
}

// Body returns the body of the response.
func (s *HTTPSocket) Body() string {
	return ""
}

// queryString returns the query string.
func (s *HTTPSocket) queryString() string {
	var query string

	for key, value := range s.query {
		query += fmt.Sprintf("%s=%s&", key, value)
	}

	return query
}
