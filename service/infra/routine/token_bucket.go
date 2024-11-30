package routine

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

type TokenBucket struct {
	D                 time.Duration //refresh token pool timer
	Mu                *sync.Mutex
	Token             chan bool //took pool
	Num               int64     //the current goroutine num
	maxNum            int64     //the allowed goroutine num
	waitingQuqueMutex *sync.Mutex
	waitingQuque      *list.List
	QueueChan         chan bool
	maxWaitingNum     int64
}

type WaitingJob struct{} // placeholder

func (t *TokenBucket) AddWaitingJob(w *WaitingJob) error {
	if int64(len(t.QueueChan)) < t.maxWaitingNum {
		t.waitingQuqueMutex.Lock()
		t.waitingQuque.PushBack(w)
		t.waitingQuqueMutex.Unlock()
		return nil
	} else {
		return ErrWaitingQueueFull
	}
}

func (t *TokenBucket) getFrontWaitingJob() *list.Element {
	t.waitingQuqueMutex.Lock()
	e := t.waitingQuque.Front()
	t.waitingQuqueMutex.Unlock()

	return e
}

func (t *TokenBucket) removeWaitingJob(e *list.Element) {
	t.waitingQuqueMutex.Lock()
	t.waitingQuque.Remove(e)
	t.waitingQuqueMutex.Unlock()
}

var ErrApplyTimeout = errors.New("apply token time out")
var ErrWaitingQueueFull = errors.New("waiting job queue is full")

func NewTokenBucket(D time.Duration, maxNum, maxWaitingNum int64) *TokenBucket {
	instance := &TokenBucket{
		D:                 D,
		Mu:                &sync.Mutex{},
		Token:             make(chan bool, maxNum),
		Num:               0,
		maxNum:            maxNum,
		maxWaitingNum:     maxWaitingNum,
		waitingQuqueMutex: &sync.Mutex{},
		waitingQuque:      list.New(),
		QueueChan:         make(chan bool, maxWaitingNum),
	}
	go instance.reset()
	return instance
}

// every timer refresh token pool
func (t *TokenBucket) reset() {
	ticker := time.NewTicker(t.D) //loop to refresh token buket pool
	for _ = range ticker.C {
		if t.Num >= t.maxNum {
			continue
		}
		t.Mu.Lock()
		supply := t.maxNum - int64(len(t.Token)) - t.Num
		for supply > 0 {
			element := t.getFrontWaitingJob()
			if element != nil {
				t.removeWaitingJob(element)
				t.QueueChan <- true
			} else {
				t.Token <- true
			}
			supply--
		}
		t.Mu.Unlock()
	}
}

// //apply for token unless beyond two waiting time
func (t *TokenBucket) ApplyToken() (bool, error) {
	select {
	case <-t.Token:
		return true, nil
	case <-time.After(time.Second * 4):
		return false, ErrApplyTimeout
	}
}
