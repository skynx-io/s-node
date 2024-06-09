//go:build linux
// +build linux

package hsec

func reportFile() string {
	return "/var/lib/skynx/report.hsr"
}

func rootTargetDir() string {
	return "/"
}

func skipDirs() []string {
	return []string{
		"/proc/",
		"/run/",
		"/srv/",
		"/mnt/",
		"/var/lib/docker/",
	}
}

func globalCacheDir() string {
	return "/var/cache/skynx"
}

/*
func globalCacheDir() string {
	tmpDir, err := os.UserCacheDir()
	if err != nil {
		tmpDir = os.TempDir()
	}

	return filepath.Join(tmpDir, "skynx", "cache")
}
*/
