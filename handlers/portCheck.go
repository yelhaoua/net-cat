package handlers

import (
	"strconv"
)
func PortCheck(port string) bool  {
	num, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	if num < 1024 {
		return false
	}
	return true
}