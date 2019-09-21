// +build linux
package fdnotifier

import (
	"syscall"
	"unsafe"

	"gvisor.dev/gvisor/pkg/waiter"
)

// TODO: Any specific reason we are usuing an anonymous struct here?
func NonBlockingPoll(fd int32, mask waiter.EventMask) waiter.EventMask {
	e := struct {
		fd      int32
		events  int16
		revents int16
	}{
		fd:     fd,
		events: int16(mask.ToLinux()),
	}

	ts := syscall.Timespec{
		Sec:  0,
		Nsec: 0,
	}

	for {
		n, _, err := syscall.RawSyscall6(syscall.SYS_PPOLL, uintptr(unsafe.Pointer(&e)), 1,
			uintptr(unsafe.Pointer(&ts)), 0, 0, 0)
		// Interrupted by signal, try again.
		if err == syscall.EINTR {
			continue
		}
		// If an error occur we'll conservatively say the FD is ready for
		// whatever is being checked.
		if err != 0 {
			return mask
		}
		if n == 0 {
			return 0
		}
		return waiter.EventMaskFromLinux(uint32(e.revents))
	}
}

// epollWait performs a blocking wait on epfd.
// Preconditions:
//  * len(events) > 0
func epollWait(epfd int, events []syscall.EpollEvent, msec int) (int, error) {
	if len(events) == 0 {
		panic("Empty events passed to EpollWait")
	}
	r, _, e := syscall.Syscall6(syscall.SYS_EPOLL_PWAIT, uintptr(epfd), uintptr(unsafe.Pointer(&events[0])), uintptr(len(events)), uintptr(msec), 0, 0)
	if e != 0 {
		return 0, e
	}
	return int(r), nil
}
