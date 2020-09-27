package internal

import (
	"fmt"
	"runtime"
)

var (
	buildDate     = "1970-01-01T00:00:00Z"
	gitBranch     = "unknown"
	gitCommit     = "unknown"
	gitState      = "unknown"
	versionNumber = "unknown"
)

// VersionInfo represents a CLI version
type VersionInfo struct {
	BuildDate string
	Compiler  string
	GitBranch string
	GitCommit string
	GitState  string
	GoVersion string
	Platform  string
	Version   string
}

// GetVersionInfo returns the curent CLI version
func GetVersionInfo() VersionInfo {
	return VersionInfo{
		BuildDate: buildDate,
		Compiler:  runtime.Compiler,
		GitBranch: gitBranch,
		GitCommit: gitCommit,
		GitState:  gitState,
		GoVersion: runtime.Version(),
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		Version:   versionNumber,
	}
}

func (v VersionInfo) String() string {
	return fmt.Sprintf("%#+v", v)
}
