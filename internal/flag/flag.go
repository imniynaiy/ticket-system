package flag

import (
	"os"

	"github.com/spf13/pflag"
)

var (
	help = pflag.BoolP("help", "h", false, "Show this help message")
)

func ParseAndHandleHelpFlag() {
	pflag.Parse()
	if *help {
		pflag.Usage()
		os.Exit(0)
	}
}
