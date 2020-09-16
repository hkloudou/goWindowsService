package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
func runit() {

	p, _ := getCurrentPath()
	// 这里循环主要是避免用户还没登录的时候，无法运行。每5秒尝试一次启动app
	go func() {
		for {
			if err := StartProcessAsCurrentUser(fmt.Sprintf("%s%s", p, appPath), fmt.Sprintf("%s%s run", p, appPath), p, true); err == nil {
				break
			}
			time.Sleep(5 * time.Second)
		}
	}()
}
