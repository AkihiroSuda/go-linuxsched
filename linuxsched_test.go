package linuxsched

import (
	"syscall"
	"testing"
)

func skipOnENOSYS(t *testing.T, err error) {
	if err == syscall.ENOSYS {
		t.Skip("Skipping due to ENOSYS")
	}
}

func TestGetAttrFromCurrentProcess(t *testing.T) {
	attr, err := GetAttr(0)
	skipOnENOSYS(t, err)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", attr)
}

func TestGetAttrFromBadProcess(t *testing.T) {
	attr, err := GetAttr(-1)
	skipOnENOSYS(t, err)
	if err != syscall.EINVAL {
		t.Fatalf("EINVAL expected: err=%s, attr=%#v", err, attr)
	}
}

func TestSetAttrToCurrentProcess(t *testing.T) {
	attr, err := GetAttr(0)
	skipOnENOSYS(t, err)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("old attr: %#v", attr)
	if !(attr.Policy == Normal && attr.Nice == 0) {
		t.Skip("Skipping due to unexpected environment")
	}
	// range  -20 (high priority) to +19 (low priority)
	attr.Nice = 1
	err = SetAttr(0, attr)
	if err != nil {
		t.Fatal(err)
	}
	attr, err = GetAttr(0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("new attr: %#v", attr)
	if !(attr.Policy == Normal && attr.Nice == 1) {
		t.Fatalf("unexpected: %#v", attr)
	}
}
