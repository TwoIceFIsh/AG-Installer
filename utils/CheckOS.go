package utils

import "runtime"

func CheckOS() string {
	var osv string
	v := runtime.GOOS
	switch v {
	case "windows":
		osv = "windows"
	case "linux":
		osv = "linux"
	}
	return osv
}
