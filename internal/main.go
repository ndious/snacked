package internal

import (
	"os"
	"path/filepath"
)

func BaseDir() string {
	baseDir := os.Getenv("BASEDIR")

	return filepath.Join(baseDir, "..")
}

func GetDir(path string) string {
	base := BaseDir()

	switch path {
	case "migrations":
		return filepath.Join(base, "config", "migration")
	case "config":
		return filepath.Join(base, "config")
	case "assets":
		return filepath.Join(base, "assets", "generated")
	default:
		return base
	}
}
