//go:build darwin
// +build darwin

package cmd

func ConsoleInit() error {
	return nil
}

func defaultConfigFile() string {
	return "/opt/skynx/etc/skynx-node.yml"
}

/*
func logFile() string {
	return "/opt/skynx/var/log/skynx-node.log"
}
*/
