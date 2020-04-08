package securepass

var insecurePasswords = []string{"password", "pass123", "password123", "1234"}

func IsInsecure(password string) bool {
	return isStringInSlice(insecurePasswords, password)
}

func isStringInSlice(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
