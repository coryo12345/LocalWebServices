package manager

import "testing"

func TestNewQueueManager(t *testing.T) {
	qm := NewQueueManager()
	if qm.queues == nil {
		t.Log("queues is not defined")
		t.Fail()
	}
}

func TestAddQueue(t *testing.T) {
	qm := NewQueueManager()
	q, _ := qm.AddQueue("myqueue", "fifo", 100)
	if q != qm.queues["myqueue"] {
		t.Log("returned queue should be added to internal queue map")
		t.Fail()
	}
}

func TestAddQueueWithMessages(t *testing.T) {
	qm := NewQueueManager()
	messages := make([]string, 0)
	q, _ := qm.AddQueueWithMessages("a", "fifo", 100, messages)
	if q != qm.queues["a"] {
		t.Log("returned queue should be added to internal queue map")
		t.FailNow()
	}

	_, err := qm.AddQueueWithMessages("b", "abc", 100, messages)
	if err == nil {
		t.Log("should fail if order is not a valid value")
		t.Fail()
	}

	_, err = qm.AddQueueWithMessages("", "fifo", 100, messages)
	if err == nil {
		t.Log("should fail when provided a non-alphanumeric queue name")
		t.Fail()
	}

	_, err = qm.AddQueueWithMessages("b", "abc", 100, messages)
	if err == nil {
		t.Log("should fail when provided a duplicate queue name")
		t.Fail()
	}
}

func TestGetQueues(t *testing.T) {
	qm := NewQueueManager()
	qs := qm.GetQueues()
	if qs == nil {
		t.Log("should return a map of queues")
		t.Fail()
	}
}

func TestGetQueue(t *testing.T) {
	qm := NewQueueManager()
	qm.queues["queue1"] = &Queue{}

	queue := qm.GetQueue("queue1")
	if queue == nil {
		t.Error("expected queue but received nil")
	}

	queue = qm.GetQueue("queue2")
	if queue != nil {
		t.Errorf("expected nil but received %s", queue)
	}
}

func TestDeleteQueue(t *testing.T) {
	qm := NewQueueManager()
	err := qm.DeleteQueue("myqueue")
	if err == nil {
		t.Log("should return an error when deleting a queue that does not exist")
		t.Fail()
	}

	q := Queue{}
	qm.queues["myqueue"] = &q

	err = qm.DeleteQueue("myqueue")
	if err != nil || qm.queues["myqueue"] != nil {
		t.Log("should delete a queue from queue map")
		t.Fail()
	}
}
