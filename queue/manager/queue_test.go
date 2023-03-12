package manager

import (
	"testing"
)

func TestAddMessage(t *testing.T) {
	content := "some content"
	q := NewQueueRaw()
	message := q.AddMessage(content)
	if len(q.messages) != 1 {
		t.Log("queue should have 1 message")
		t.Fail()
	}
	if q.messages[0].GetContent() != content || message.GetContent() != content {
		t.Logf("content of message should be %s", content)
		t.Fail()
	}
}

func TestPoll_fifo(t *testing.T) {
	q := NewQueueRaw()
	q.messages = append(q.messages, NewMessage("message 0"), NewMessage("message 1"))

	firstMessages := q.Poll()
	if len(firstMessages) != len(q.messages) {
		t.Logf("polled messages should retrieve all messages from queue but received %d", len(firstMessages))
		t.FailNow()
	}
	if (firstMessages)[0].content != "message 0" || (firstMessages)[1].content != "message 1" {
		t.Log("polled message content did not match original content")
		t.Fail()
	}

	secondMessages := q.Poll()
	if len(secondMessages) != 0 {
		t.Logf("second poll attempt should receive zero messages but received %d", len(secondMessages))
		t.FailNow()
	}
}

func TestPoll_filo(t *testing.T) {
	q := NewQueueRaw()
	q.messages = append(q.messages, NewMessage("message 0"), NewMessage("message 1"))
	q.order = "filo"

	firstMessages := q.Poll()
	if len(firstMessages) != len(q.messages) {
		t.Logf("polled messages should retrieve all messages from queue but received %d", len(firstMessages))
		t.FailNow()
	}
	if (firstMessages)[0].content != "message 1" || (firstMessages)[1].content != "message 0" {
		t.Log("polled message content did not match original content")
		t.Fail()
	}

	secondMessages := q.Poll()
	if len(secondMessages) != 0 {
		t.Logf("second poll attempt should receive zero messages but received %d", len(secondMessages))
		t.FailNow()
	}
}

func TestPollByCount_fifo(t *testing.T) {
	q := NewQueueRaw()
	q.messages = append(q.messages, NewMessage("message 0"), NewMessage("message 1"), NewMessage("message 2"))

	firstMessages := q.PollByCount(1)
	if len(firstMessages) != 1 {
		t.Logf("polling should have received only 1 message but received %d", len(firstMessages))
		t.FailNow()
	}
	if (firstMessages)[0].content != "message 0" {
		t.Log("polled message content did not match original content")
		t.Fail()
	}

	secondMessages := q.PollByCount(2)
	if len(secondMessages) != 2 {
		t.Logf("second poll attempt should receive 2 messages but received %d", len(secondMessages))
		t.FailNow()
	}
	if (secondMessages)[0].content != "message 1" || (secondMessages)[1].content != "message 2" {
		t.Log("polled message content did not match original content")
		t.Fail()
	}
}

func TestPollByCount_filo(t *testing.T) {
	q := NewQueueRaw()
	q.messages = append(q.messages, NewMessage("message 0"), NewMessage("message 1"), NewMessage("message 2"))
	q.order = "filo"

	firstMessages := q.PollByCount(1)
	if len(firstMessages) != 1 {
		t.Logf("polling should have received only 1 message but received %d", len(firstMessages))
		t.FailNow()
	}
	if (firstMessages)[0].content != "message 2" {
		t.Log("polled message content did not match original content")
		t.Fail()
	}

	secondMessages := q.PollByCount(2)
	if len(secondMessages) != 2 {
		t.Logf("second poll attempt should receive 2 messages but received %d", len(secondMessages))
		t.FailNow()
	}
	if (secondMessages)[0].content != "message 1" || (secondMessages)[1].content != "message 0" {
		t.Log("polled message content did not match original content")
		t.Fail()
	}
}
