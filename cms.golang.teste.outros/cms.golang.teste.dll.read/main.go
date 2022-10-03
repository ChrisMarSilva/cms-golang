package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

// go mod init github.com/ChrisMarSilva/cms-golang-teste-dll
// go mod tidy
// go run main.go

func main() {

	libname := "SPB_SecDll.dll"
	procname := "EncryptMsgV3"

	// dll, err := syscall.LoadDLL(libname)
	// if err != nil {
	// 	fmt.Println("erro LoadDLL: ", err)
	// 	return
	// }
	// _, err = dll.FindProc(procname)
	// if err != nil {
	// 	fmt.Println("erro FindProc: ", err)
	// 	return
	// }
	//
	// ret, _, _ := proc.Call()

	dll := syscall.NewLazyDLL(libname)
	dll.NewProc(procname)
	// _, _, _ = lazyProc.Call(2, uintptr(unsafe.Pointer(&args)))

	// dll2 := windows.NewLazySystemDLL("libname)
	// lazyProc := dll2.NewProc("EncryptMsgV3")
	// lazyProc.Call()

	dll2, err := syscall.LoadLibrary(libname) //Make sure this DLL follows Golang machine bit architecture (64-bit in my case)
	if err != nil {
		fmt.Println("erro LoadLibrary: ", err)
		//return
	}
	defer syscall.FreeLibrary(dll2)

	_, err = syscall.GetProcAddress(dll2, procname)
	if err != nil {
		fmt.Println("erro GetProcAddress: ", err)
		//return
	}

	kernel32, _ := syscall.LoadLibrary(libname)
	_, err = syscall.GetProcAddress(kernel32, procname)
	if err != nil {
		fmt.Println("erro GetProcAddress 2222: ", err)
		//return
	}

	// r1, r2, errno := syscall.Syscall12(readCard,
	// 	11,
	// 	uintptr(unsafe.Pointer(&room)),
	// 	uintptr(unsafe.Pointer(&gate)),
	// 	uintptr(unsafe.Pointer(&stime)),
	// 	uintptr(unsafe.Pointer(&guestname)),
	// 	uintptr(unsafe.Pointer(&guestid)),
	// 	uintptr(unsafe.Pointer(&lift)),
	// 	uintptr(unsafe.Pointer(&track1)),
	// 	uintptr(unsafe.Pointer(&track2)),
	// 	uintptr(cardno),
	// 	uintptr(st),
	// 	uintptr(Breakfast),
	// 	0)

	// 	const maxSize = 100 //Adjust the value to your need

	// room := make([]byte, maxSize)
	// gate := make([]byte, maxSize)
	// stime := make([]byte, maxSize)
	// guestname := make([]byte, maxSize)
	// guestid := make([]byte, maxSize)
	// lift := make([]byte, maxSize)
	// track1 := make([]byte, maxSize)
	// track2 := make([]byte, maxSize)

	// //The following variable type is correct for 64-bit DLL.
	// //For 32-bit DLL, instead of int64, use int32.
	// cardno := int64(0)
	// st := int64(0)
	// Breakfast := int64(0)

	// MAINDLL, _ := syscall.LoadLibrary("xxxxx.dll")
	// defer syscall.FreeLibrary(MAINDLL)
	// readCard, _ := syscall.GetProcAddress(MAINDLL, "ReadCard")

	// r1, r2, errno := syscall.Syscall12(readCard,
	//     11,
	//     uintptr(unsafe.Pointer(&room[0])),
	//     uintptr(unsafe.Pointer(&gate[0])),
	//     uintptr(unsafe.Pointer(&stime[0])),
	//     uintptr(unsafe.Pointer(&guestname[0])),
	//     uintptr(unsafe.Pointer(&guestid[0])),
	//     uintptr(unsafe.Pointer(&lift[0])),
	//     uintptr(unsafe.Pointer(&track1[0])),
	//     uintptr(unsafe.Pointer(&track2[0])),
	//     uintptr(unsafe.Pointer(&cardno)),
	//     uintptr(unsafe.Pointer(&st)),
	//     uintptr(unsafe.Pointer(&Breakfast)),
	//     0)

	// fmt.Println(r1, "\n", r2, "\n", errno)

	// func cstr(buf []byte) string {
	// 	str := string(buf)
	// 	for i, r := range str {
	// 		if r == 0 {
	// 			return string(buf[:i])
	// 		}
	// 	}
	// 	return str
	// }

	// //usage example
	// sRoom := cstr(room)
	// //see the difference
	// fmt.Printf("`%s` => `%s`\n", string(room), sRoom)

	fmt.Println("FindProc ok")

}

// uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("mono.exe"))),
// uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(app)))  )
// uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("This test is Done."))),
// 		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Done Title"))),
// 		uintptr(MB_YESNOCANCEL))
// if ret, _, callErr := syscall.Syscall(uintptr(getModuleHandle), nargs, 0, 0, 0); callErr != 0 {
// uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
// 	uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),

// function AddIntegers(const _a, _b: integer): integer; stdcall;
// function CountChars(_s: Pchar): integer; StdCall;
// function MeuLowerCase(_s: PAnsiChar): PAnsiChar; stdcall;

func main_old() {

	dll1, _ := syscall.LoadDLL("DelphiDLL.dll")

	proc1, _ := dll1.FindProc("AddIntegers")
	ret1, _, _ := proc1.Call(8, 9)
	fmt.Println("ret1", ret1)

	proc11, _ := dll1.FindProc("CountChars")
	ret11, _, _ := proc11.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("ABC123"))))
	fmt.Println("ret11", ret11)

	var dll2 = syscall.NewLazyDLL("DelphiDLL.dll")

	var proc2 = dll2.NewProc("AddIntegers")
	ret2, _, _ := proc2.Call(8, 9)
	fmt.Println("ret2", ret2)

	var proc22 = dll2.NewProc("CountChars")
	ret22, _, _ := proc22.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("ABC123"))))
	fmt.Println("ret22", ret22)

	// dll3, _ := syscall.LoadLibrary("DelphiDLL.dll")
	// defer syscall.FreeLibrary(dll3)
	// proc3, e := syscall.GetProcAddress(dll3, "AddIntegers")
	// ret3, _, _ := syscall.Syscall12(uintptr(proc3), uintptr(unsafe.Pointer(8)), uintptr(unsafe.Pointer(9)))
	// fmt.Println("ret3", ret3)

}

/*

go get -u golang.org/x/sys
golang.org/x/sys/windows


func signalDaemonDump(pid int) {
	modkernel32 := windows.NewLazySystemDLL("kernel32.dll")
	procOpenEvent := modkernel32.NewProc("OpenEventW")
	procPulseEvent := modkernel32.NewProc("PulseEvent")

	ev := "Global\\docker-daemon-" + strconv.Itoa(pid)
	h2, _ := openEvent(0x0002, false, ev, procOpenEvent)
	if h2 == 0 {
		return
	}
	pulseEvent(h2, procPulseEvent)
}

vmcompute := windows.NewLazySystemDLL("vmcompute.dll")
	if vmcompute.Load() != nil {
		return fmt.Errorf("Failed to load vmcompute.dll. Ensure that the Containers role is installed.")
	}


// getFileSystemType obtains the type of a file system through GetVolumeInformation
// https://msdn.microsoft.com/en-us/library/windows/desktop/aa364993(v=vs.85).aspx
func getFileSystemType(drive string) (fsType string, hr error) {
	var (
		modkernel32              = windows.NewLazySystemDLL("kernel32.dll")
		procGetVolumeInformation = modkernel32.NewProc("GetVolumeInformationW")
		buf                      = make([]uint16, 255)
		size                     = syscall.MAX_PATH + 1
	)
	if len(drive) != 1 {
		hr = errors.New("getFileSystemType must be called with a drive letter")
		return
	}
	drive += `:\`
	n := uintptr(unsafe.Pointer(nil))
	r0, _, _ := syscall.Syscall9(procGetVolumeInformation.Addr(), 8, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(drive))), n, n, n, n, n, uintptr(unsafe.Pointer(&buf[0])), uintptr(size), 0)
	if int32(r0) < 0 {
		hr = syscall.Errno(win32FromHresult(r0))
	}
	fsType = syscall.UTF16ToString(buf)
	return
}

*/
