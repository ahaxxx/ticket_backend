package utils

// Mp Save login user
var Mp map[int]string

func init() {
	Mp = make(map[int]string, 20)
}

// Set Save or Update token
func Set(id int, token string) {
	Mp[id] = token
}

// Get Get the login information
func Get(id int) (string, bool) {
	v, ok := Mp[id]
	return v, ok
}
