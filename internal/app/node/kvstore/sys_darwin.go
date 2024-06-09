//go:build darwin
// +build darwin

package kvstore

func dbDir() string {
	return "/opt/skynx/var/lib/db"
}
