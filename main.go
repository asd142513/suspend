package main

import (
	"fmt"
	"net/http"
	"syscall"
	"time"
)

var (
	powrprof, _        = syscall.LoadLibrary("powrprof.dll")
	setSuspendState, _ = syscall.GetProcAddress(powrprof, "SetSuspendState")
)

func suspend() {
	time.Sleep(5 * time.Second)
	syscall.SyscallN(uintptr(setSuspendState), 0, 0, 0)
}

func suspendHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Suspending...")
	go suspend()
}

func run() {
	http.HandleFunc("/suspend", suspendHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	run()
}
