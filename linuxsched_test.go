package linuxsched

import (
	"testing"
)

func TestGetSchedAttrFromCurrentProcess(t *testing.T) {
	attr := SchedAttr{}
	err := attr.GetFrom(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", attr)
}

func TestGetSchedAttrFromBadProcess(t *testing.T) {
	attr := SchedAttr{}
	err := attr.GetFrom(-1, 0)
	t.Logf("error (perhaps expected): %s", err)
	if err == nil {
		t.Fatal("error is expected")
	}
}

func TestSetSchedAttrToCurrentProcess(t *testing.T) {
	attr := SchedAttr{}
	err := attr.GetFrom(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("old attr: %#v", attr)
	if !(attr.Policy == Normal && attr.Nice == 0) {
		t.Skip("Skipping due to unexpected environment")
	}
	attr = SchedAttr{
		// range  -20 (high priority) to +19 (low priority)
		Nice: 1,
	}
	err = attr.SetTo(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	attr = SchedAttr{}
	err = attr.GetFrom(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("new attr: %#v", attr)
	if !(attr.Policy == Normal && attr.Nice == 1) {
		t.Fatalf("unexpected: %#v", attr)
	}
}
