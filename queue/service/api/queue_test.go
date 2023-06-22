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

	resp, apierr := qa.GetAllQueues()
	if apierr != nil {
		t.Log("should not return an error")
		t.FailNow()
	}

	if len(resp) != 1 {
		t.Error("expected response length to be 1 queue")
	}

	if resp[0].Name != "a" || resp[0].Order != "fifo" {
		t.Error("expected name = 'a' and order = 'fifo'")
	}
}

func TestCreateQueue(t *testing.T) {
	ctrl := gomock.NewController(t)
	man := mock_manager.NewMockIQueueManager(ctrl)
	qa := NewQueueAPI(man)

	qreq := CreateQueueRequest{
		Name:    "queue",
		Order:   "fifo",
		Timeout: 100,
	}
	man.EXPECT().AddQueue(qreq.Name, qreq.Order, qreq.Timeout).Times(1)
	jData, _ := json.Marshal(qreq)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(jData))
	_, err := qa.CreateQueue(request)
	if err != nil {
		t.Error("expected response to be defined")
	}

	qreq = CreateQueueRequest{
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

	qreq := CreateQueueRequest{
		Name:    "queue",
		Order:   "fifo",
		Timeout: 100,
	}
	man.EXPECT().DeleteQueue(qreq.Name).Times(1)
	jData, _ := json.Marshal(qreq)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(jData))
	_, err := qa.DeleteQueue(request)
	if err != nil {
		t.Log("unexepcted error returned from DeleteQueue")
		t.FailNow()
	}

	qreq = CreateQueueRequest{
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
	mq := mock_manager.NewMockIQueue(ctrl)
	mq.EXPECT().GetName().Return("a")
	mq.EXPECT().GetOrder().Return("fifo")
	mq.EXPECT().GetVisibilityTimeout().Return(100)
	m["a"] = mq
	return m
}
