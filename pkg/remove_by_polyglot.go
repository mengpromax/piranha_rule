package pkg

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path"
	"runtime"

	utils "github.com/mengpromax/piranha_rule/pkg/util"
)

type osArch string

func (v osArch) valid() bool {
	return utils.CreateStringSet(supportedOSArch).Contains(string(v))
}

var supportedOSArch = []string{
	string(arm64),
	string(amd64),
}

const (
	arm64 osArch = "arm64"
	amd64 osArch = "amd64"
)

func RemoveGrayKey(absoluteSrcDir string, constantGrayKeyMatchPattern, grayKeyMatchPattern string,
	executableDir string) (err error) {
	goArch := runtime.GOARCH
	if !osArch(goArch).valid() {
		log.Fatal("current darwin architecture is not supported\n")
	}

	executor := path.Join(executableDir, "executors", goArch)

	c := exec.Command(executor,
		"--path-to-codebase", absoluteSrcDir,
		"--path-to-configurations", path.Join(executableDir, "rules"),
		"-l", "go",
		"-s", fmt.Sprintf("constant_gray_key=%s", constantGrayKeyMatchPattern),
		"-s", fmt.Sprintf("plain_gray_key=%s", grayKeyMatchPattern),
	)

	var stderr bytes.Buffer
	c.Stderr = &stderr

	err = c.Run()
	if err != nil {
		return fmt.Errorf(stderr.String())
	}

	return nil
}
