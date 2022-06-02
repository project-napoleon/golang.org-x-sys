//go:build (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris) && !sw64
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
// +build !sw64

package unix

import (
	"unsafe"
)

func Pread(fd int, p []byte, offset int64) (n int, err error) {
	n, err = pread(fd, p, offset)
	if raceenabled {
		if n > 0 {
			raceWriteRange(unsafe.Pointer(&p[0]), n)
		}
		if err == nil {
			raceAcquire(unsafe.Pointer(&ioSync))
		}
	}
	return
}

func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
	if raceenabled {
		raceReleaseMerge(unsafe.Pointer(&ioSync))
	}
	n, err = pwrite(fd, p, offset)
	if raceenabled && n > 0 {
		raceReadRange(unsafe.Pointer(&p[0]), n)
	}
	return
}
