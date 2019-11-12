
package main

import (
"fmt"
"log"
"os"
	"runtime"
	"strings"
"golang.org/x/sys/windows/svc"
"service/infra"

)

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"usage: %s <command>\n"+
			"       where <command> is one of\n"+
			"       install, remove, debug, start, stop, pause or continue.\n",
		errmsg, os.Args[0])
	os.Exit(2)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("hello from windows")
	}
	const svcName = "myservice"

	isIntSess, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Fatalf("failed to determine if we are running in an interactive session: %v", err)
	}
	if !isIntSess {
		infra.RunService(svcName, false)
		return
	}

	if len(os.Args) < 2 {
		usage("no command specified")
	}

	cmd := strings.ToLower(os.Args[1])
	switch cmd {
	case "debug":
		infra.RunService(svcName, true)
		return
	case "install":
		err = infra.InstallService(svcName, "my service")
	case "remove":
		err = infra.RemoveService(svcName)
	case "start":
		err = infra.StartService(svcName)
	case "stop":
		err = infra.ControlService(svcName, svc.Stop, svc.Stopped)
	case "pause":
		err = infra.ControlService(svcName, svc.Pause, svc.Paused)
	case "continue":
		err = infra.ControlService(svcName, svc.Continue, svc.Running)
	default:
		usage(fmt.Sprintf("invalid command %s", cmd))
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, svcName, err)
	}
	return
}
