package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/kardianos/service"
)

var serviceConfig = &service.Config{
	Name:        serviceName,
	DisplayName: serviceDisplayName,
	Description: serviceDescription,
}
var logger service.Logger

func main() {

	// 构建服务对象
	prog := &Program{}
	s, err := service.New(prog, serviceConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 用于记录系统日志
	var errlog error
	logger, errlog = s.Logger(nil)
	if errlog != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
		return
	}

	cmd := os.Args[1]

	if cmd == "install" {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("安装成功")
		s.Start()
	}
	if cmd == "uninstall" {
		s.Stop()
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("卸载成功")
	}

	// install, uninstall, start, stop 的另一种实现方式
	// err = service.Control(s, os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	log.Println("Startting service ....")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	log.Println("Stopping service ...")
	cupath, _ := getCurrentPath()
	// log.Println("path:", fmt.Sprintf("%sEtaxHelper.exe", p), fmt.Sprintf("%sEtaxHelper.exe run", p), p)
	lockFile := fmt.Sprintf("%slock.pid", cupath)
	lock, err := os.Open(lockFile)
	defer lock.Close()
	if err == nil {
		filePid, err := ioutil.ReadAll(lock)

		if err == nil {
			pidStr := fmt.Sprintf("%s", filePid)

			pid, _ := strconv.Atoi(pidStr)
			x, err := os.FindProcess(pid)
			if err == nil {
				fmt.Printf("[ERROR] 工具已启动[%s].", pidStr)
				if err := x.Kill(); err != nil {
					logger.Error("err kill", err)
				} else {
					logger.Info("killed pid", pid)
				}
			} else {
				logger.Warning("err FindProcess", err)
			}
		} else {
			logger.Warning("not read pid file", err)
		}
	} else {
		logger.Warning("not open pid file", err)
	}
	return nil
}

func (p *Program) run() {
	// 此处编写具体的服务代码
	runit()
}
