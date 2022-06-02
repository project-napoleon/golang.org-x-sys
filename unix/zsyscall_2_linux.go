//go:build linux && sw64
// +build linux,sw64

package unix

// copied from zsyscall_linux.go for sw64 compiling
func ioctlPtr(fd int, req uint, arg unsafe.Pointer) (err error) {
	_, _, e1 := Syscall(SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(arg))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
