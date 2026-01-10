package functions

import (
	"net"
	"strings"
)

func ValidUsername(user string, conn net.Conn) bool {
	if len(user) == 0 {
		return false
	}
	if strings.HasPrefix(user, "/nick"){
		return false
	}
	for _, char := range user {
		if char < 32 || char > 126{
			return false
		}
	}
	return true
}
