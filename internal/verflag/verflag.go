package verflag

import (
	"log"
	"os"

	"github.com/spf13/pflag"
)

var (
	// GitVersion 是语义化的版本号.
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate 是 ISO8601 格式的构建时间, $(date -u +'%Y-%m-%dT%H:%M:%SZ') 命令的输出.
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit 是 Git 的 SHA1 值，$(git rev-parse HEAD) 命令的输出.
	GitCommit = "$Format:%H$"
	// GitTreeState 代表构建时 Git 仓库的状态，可能的值有：clean, dirty.
	GitTreeState = ""
)

var versionFlag = pflag.BoolP("version", "v", false, "Show version info")

func HandleVersionFlagAndExit() {
	if *versionFlag {
		log.Printf("Version: %#v\n", GitVersion)
		os.Exit(0)
	}
}
