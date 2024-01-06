package tool

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Ocyss/OI/generator/utils"
)

func OpenGoLand(dir string, files ...string) {
	dirAbs, _ := filepath.Abs(dir)
	args := []string{dirAbs}
	for _, file := range files {
		fileAbs, _ := filepath.Abs(file)
		args = append(args, fileAbs)
	}
	err := os.MkdirAll(dirAbs, os.ModePerm)
	if err == nil {
		cmd := exec.Command(utils.Config.GoLandPath, args...)
		_, _ = cmd.CombinedOutput()
	}
}

func Er(err error) {
	if err != nil {
		panic(err)
	}
}
