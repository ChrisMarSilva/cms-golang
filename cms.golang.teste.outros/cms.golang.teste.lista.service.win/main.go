package main

import (
	"fmt"
	"log"
	"time"

	ps "github.com/mitchellh/go-ps"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"golang.org/x/sys/windows"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.lista.service.win
// go get github.com/mitchellh/go-ps
// go mod tidy

// go run main.go

const processEntrySize = 568

func main() {

	log.Println("INI")
	var start time.Time = time.Now()

	pids, _ := win.EnumProcesses()
	for _, pid := range pids {
		hSnap, _ := win.CreateToolhelp32Snapshot(co.TH32CS_SNAPPROCESS, pid)
		defer hSnap.CloseHandle()
		hSnap.EnumProcesses(func(pe32 *win.PROCESSENTRY32) {
			fmt.Printf("PID: %d @ %s\n", pe32.Th32ProcessID, pe32.SzExeFile())
		})
	}

	log.Println("FIM:", time.Since(start))
}

func whatever3() {

	pids, _ := win.EnumProcesses()
	for _, pid := range pids {
		hSnap, _ := win.CreateToolhelp32Snapshot(co.TH32CS_SNAPMODULE, pid)
		defer hSnap.CloseHandle()
		hSnap.EnumModules(func(me32 *win.MODULEENTRY32) {
			fmt.Printf("PID: %d, %s @ %s\n", me32.Th32ProcessID, me32.SzModule(), me32.SzExePath())
		})
	}
}

func whatever2() {

	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if e != nil {
		panic(e)
	}

	p := windows.ProcessEntry32{Size: processEntrySize}
	for {
		e := windows.Process32Next(h, &p)
		if e != nil {
			break
		}
		s := windows.UTF16ToString(p.ExeFile[:])
		println(s)
	}

}

func whatever() {

	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	for x := range processList {
		var process ps.Process
		process = processList[x]
		log.Printf("%d\t%s\n", process.Pid(), process.Executable())
	}

}
