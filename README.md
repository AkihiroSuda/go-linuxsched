go-linuxsched
==================
[![Build Status](https://travis-ci.org/AkihiroSuda/go-linuxsched.svg)](https://travis-ci.org/AkihiroSuda/go-linuxsched)
[![GoDoc](https://godoc.org/github.com/AkihiroSuda/go-linuxsched?status.svg)](https://godoc.org/github.com/AkihiroSuda/go-linuxsched)

Go binding for [`sched_setattr(2)` and `sched_getattr(2)`](http://man7.org/linux/man-pages/man2/sched_getattr.2.html).

```go
type SchedAttr struct {
	Policy   SchedPolicy
	Flags    SchedFlag
	Nice     int32
	Priority uint32
	Runtime  time.Duration
	Deadline time.Duration
	Period   time.Duration
}

func (a *SchedAttr) SetTo(pid int, flags uint) error {..}

func (a *SchedAttr) GetFrom(pid int, flags uint) error {..}
```
