package utils

import (
	"strings"
	"goServer/models"
	"goServer/mlog"
	"path/filepath"
	"os"
	"fmt"
)

const (
	LOG_ROTATE_SIZE = 10000
	LOG_FILENAME    = "monthlyExhibition.logs"
)

var (
	commonBaseSearchPaths = []string{
		".",
		"..",
		"../..",
		"../../..",
	}
)

func MloggerConfigFromLoggerConfig(s *models.LogSettings) *mlog.LoggerConfiguration {
	fmt.Println("과연",s)
	return &mlog.LoggerConfiguration{
		EnableConsole: s.EnableConsole,
		ConsoleJson:   *s.ConsoleJson,
		ConsoleLevel:  strings.ToLower(s.ConsoleLevel),
		EnableFile:    s.EnableFile,
		FileJson:      *s.FileJson,
		FileLevel:     strings.ToLower(s.FileLevel),
		FileLocation:  GetLogFileLocation(s.FileLocation),
	}
}
func FindPath(path string, baseSearchPaths []string,  filter func(info os.FileInfo) bool) string {
	if filepath.IsAbs(path) {
		if _, err := os.Stat(path); err == nil {
			return path
		}
		return ""
	}
	searchPaths := []string{}
	for _, baseSearchPaths := range baseSearchPaths {
		searchPaths = append(searchPaths, baseSearchPaths)
	}
	var binaryDir string
	if exe, err := os.Executable(); err == nil {
		if exe, err = filepath.EvalSymlinks(exe); err == nil {
			if exe, err = filepath.Abs(exe); err == nil {
				binaryDir = filepath.Dir(exe)
			}
		}
	}
	if binaryDir != "" {
		for _, baseSearchPaths := range baseSearchPaths {
			searchPaths = append(
				searchPaths, filepath.Join(binaryDir, baseSearchPaths),
			)
		}
	}

	for _, parent := range searchPaths {
		found, err := filepath.Abs(filepath.Join(parent, path))
		if err != nil {
			continue
		} else if fileInfo, err := os.Stat(found); err == nil {
			if filter != nil {
				if filter(fileInfo) {
					return found
				}
			} else {
				return found
			}
		}
	}
	return ""
}

// directory as well as the directory of the executable.
func FindFile(path string) string {
	return FindPath(path, commonBaseSearchPaths, func(fileInfo os.FileInfo) bool {
		return !fileInfo.IsDir()
	})
}

func FindDir(dir string) (string,bool) {
found := FindPath(dir, commonBaseSearchPaths, func(fileInfo os.FileInfo) bool {
	return fileInfo.IsDir()
})
if found == ""{
	return "./", false
}
	return found, true
}
func GetLogFileLocation(fileLocation string) string {
	if fileLocation == ""{
		fileLocation, _ = FindDir("logs")
	}
	return filepath.Join(fileLocation, LOG_FILENAME)
}