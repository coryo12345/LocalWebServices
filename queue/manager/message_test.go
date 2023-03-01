package manager

import (
	"testing"
	"time"
)

const (
	content = "some message content"
)

func TestGetContent(t *testing.T) {
	m := NewMessage(content)
	actual := m.GetContent()
	if actual != content {
		t.Logf("expected content to be '%s' but was '%s'", content, actual)
		t.Fail()
	}
}

func TestGetTimestamp(t *testing.T) {
	m := NewMessage(content)
	expected := time.Now().String()
	m.timestamp = expected
	actual := m.GetTimestamp()
	if actual != expected {
		t.Logf("expected content to be '%s' but was '%s'", expected, actual)
		t.Fail()
	}
}

func TestLock(t *testing.T) {
	m := NewMessage(content)
	if m.visible != true {
		t.Log("visible should be true")
		t.FailNow()
	}
	m.Lock(10000)
	if m.visible != false {
		t.Log("after locking, visible should be false")
		t.FailNow()
	}
}
