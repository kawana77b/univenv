package sysutil

import "os"

func FileExists(p string) bool {
	fi, err := os.Stat(p)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

func DirExists(p string) bool {
	fi, err := os.Stat(p)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func HomeDir() string {
	d, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return d
}
