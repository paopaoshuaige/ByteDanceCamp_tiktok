package utils

const maxUserPassSaltLen = 32

func CheckName(username string) bool {
	return len(username) != 0 && len(username) <= maxUserPassSaltLen
}

func CheckPass(pass string) bool {
	return len(pass) != 0 && len(pass) <= maxUserPassSaltLen
}
