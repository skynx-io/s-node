//go:build linux
// +build linux

package kvstore

func dbDir() string {
	return "/var/lib/skynx/db"
}
