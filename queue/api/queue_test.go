package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"queue/manager"
	mock_manager "queue/manager/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetAllQueues(t *testing.T) {
	ctrl := gomock.NewController(t)
	man := mock_manager.NewMockIQueueManager(ctrl)

	man.EXPECT().GetQueues().Return(getSampleQueues(ctrl)).Times(1)

	qa := NewQueueAPI(man)

	const expected = `["a"]`
	bdata, err := qa.GetAllQueues()
	if err != nil {
		t.Log("should not return an error")
		t.FailNow()
	}
	if string(bdata) != expected {
		t.Errorf("expected %s but received %s", expected, string(bdata))
	}
}

func TestCreateQueue(t *testing.T) {
	ctrl := gomock.NewController(t)
	man := mock_manager.NewMockIQueueManager(ctrl)
	qa := NewQueueAPI(man)

	qreq := QueueRequest{
		Name:    "queue",
		Order:   "fifo",
		Timeout: 100,
	}
	man.EXPECT().AddQueue(qreq.Name, qreq.Order, qreq.Timeout).Times(1)
	jData, _ := json.Marshal(qreq)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(jData))
	byteData, err := qa.CreateQueue(request)
	if err != nil {
		t.Log("unexepcted error returned from CreateQueue")
		t.FailNow()
	}
	if string(byteData) != "{}" {
		t.Error(`expected return value to be "{}"`)
	}

	qreq = QueueRequest{
		Name:    "",
		Order:   "fifo",
		Timeout: 100,
	}
	man.EXPECT().AddQueue(qreq.Name, qreq.Order, qreq.Timeout).Times(0)
	jData, _ = json.Marshal(qreq)
	request, _ = http.NewRequest(http.MethodPost, "", bytes.NewReader(jData))
	_, err = qa.CreateQueue(request)
	if err == nil {
		t.Log("expected error due to blank name, but none recieved")
		t.FailNow()
	}
}

func TestDeleteQueue(t *testing.T) {
	ctrl := gomock.NewController(t)
	man := mock_manager.NewMockIQueueManager(ctrl)
	qa := NewQueueAPI(man)

	qreq := QueueRequest{
		Name:    "queue",
		Order:   "fifo",
		Timeout: 100,
	}
	man.EXPECT().DeleteQueue(qreq.Name).Times(1)
	jData, _ := json.Marshal(qreq)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(jData))
	byteData, err := qa.DeleteQueue(request)
	if err != nil {
		t.Log("unexepcted error returned from DeleteQueue")
		t.FailNow()
	}
	if string(byteData) != "{}" {
		t.Error(`expected return value to be "{}"`)
	}

	qreq = QueueRequest{
		Name:    "",
		Order:   "fifo",
		Timeout: 100,
	}
	man.EXPECT().DeleteQueue(qreq.Name).Times(0)
	jData, _ = json.Marshal(qreq)
	request, _ = http.NewRequest(http.MethodPost, "", bytes.NewReader(jData))
	_, err = qa.DeleteQueue(request)
	if err == nil {
		t.Log("expected error due to blank name, but none recieved")
		t.FailNow()
	}
}

func getSampleQueues(ctrl *gomock.Controller) map[string]manager.IQueue {
	m := make(map[string]manager.IQueue)
	m["a"] = mock_manager.NewMockIQueue(ctrl)
	return m
}
