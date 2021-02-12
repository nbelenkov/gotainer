package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//docker             run image <cmd> <params>
//go run main.go     run       <cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	
	case "child":
		child()
	default:
		panic("bad command")
	}
}

func run() {
	fmt.Print("Running ", os.Args[2:], " as ", os.Getpid(), "\n")

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:] ... ) ...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	cmd.Run()

}

func child() {
	fmt.Print("Running ", os.Args[2:], " as ", os.Getpid(), "\n")

	syscall.Sethostname([]byte("Container"))
	syscall.Chroot("/vagrant/ubuntu-fs")
	syscall.Chdir("/")

	cmd := exec.Command(os.Args[2], os.Args[3:] ...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
