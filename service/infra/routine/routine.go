package routine

import (
	"context"
	"e-commerce/service/infra/log"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

const (
	DEFAULT_CACHE_WAITING_ROUTINE          = 80
	DEFAULT_MAX_CONCURRENT_WORKING_ROUTINE = 20
	DEFAULT_RESET_INTERVAL                 = time.Millisecond

	MQTT_ESL_SERVICE           = "mqttservice-esl"
	PRODUCT_MANAGEMENT_SERVICE = "product-management"
	DEV_ESL_MANAGEMENT_SERVICE = "device-management-esl"
)

var (
	ctx          context.Context
	canceller    context.CancelFunc
	waitGroup    sync.WaitGroup
	routineCount int32
)

func init() {
	routineCount = 0
	waitGroup = sync.WaitGroup{}
	ctx, canceller = context.WithCancel(context.Background())
}

type Runner struct {
	tb *TokenBucket
}

func NewRunner(interval time.Duration, maxWorkingNum, maxWaitingNum int64) *Runner {
	tb := NewTokenBucket(interval, maxWorkingNum, maxWaitingNum)
	return &Runner{
		tb: tb,
	}
}

func (r *Runner) LimitRun(exec func()) {
	tb := r.tb
	if ok, err := tb.ApplyToken(); !ok {
		w := &WaitingJob{}
		tb.AddWaitingJob(w)
		log.Debug("Apply token err: ", err)
		select {
		case <-tb.QueueChan:
			atomic.AddInt64(&tb.Num, 1)
			SyncRun(exec, tb)
		}
	} else {
		atomic.AddInt64(&tb.Num, 1)
		SyncRun(exec, tb)
	}
}

func SyncRun(exec func(), tb *TokenBucket) {
	syncExec := func() {
		defer atomic.AddInt64(&tb.Num, -1)
		defer catchPanic()

		exec()
	}

	go syncExec()
}

func Run(exec func()) {
	go loop(exec)
}

func loop(exec func()) {
	log.Debug("Routine: Anonymous Go Routine started")
	defer catchPanic()

	// This is for safe exit
	AddWait()
	defer WaitDone()

	exec()
	log.Debug("Routine: Anonymous Go Routine Left")
}

func Run3(name string, exec func(), cleanup func()) {
	log.Debug("Routine: Starting Go routine,", name)
	go loop3(name, exec, cleanup)
}

func loop3(name string, exec func(), cleanup func()) {
	log.Info("Routine: Go Routine started,", name)
	defer catchPanic()

	// This is for safe exit
	AddWait()
	defer WaitDone()

	exec()

	log.Info("Routine: Go Routine Leaving & Cleanup,", name)
	cleanup()

	log.Info("Routine: Go Routine Left,", name)
}

func RunAsDaemon(exec func()) {
	go loopAsDaemon(exec)
}

func loopAsDaemon(exec func()) {
	log.Debug("Routine: Anonymous Go Routine started")
	defer catchPanic()

	// This is for safe exit
	AddWait()
	defer WaitDone()

	exec()
	log.Debug("Routine: Anonymous Go Routine Left")
}

func catchPanic() {
	if err := recover(); err != nil {
		log.Warn("capture panic for goroutine:", err)
	}
}

func LoopUntilExit() {
	waitSignal()
	ExitAll()
}

func waitSignal() {
	ch := make(chan os.Signal)
	//signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGTSTP,
	//	syscall.SIGQUIT, syscall.SIGKILL)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP,
		syscall.SIGQUIT, syscall.SIGKILL)
	sig := <-ch
	log.Info("Received Exit Signal:", sig)
}

func ExitAll() {
	log.Info("Routine: Exiting all routines")
	canceller()

	log.Info("Routine: Wait for all routines exit")
	Wait()
	log.Info("Routine: Wait Finish.")
}

func AddWait() {
	atomic.AddInt32(&routineCount, 1)
	waitGroup.Add(1)
}

func WaitDone() {
	waitGroup.Done()
	atomic.AddInt32(&routineCount, -1)
}

func Wait() {
	waitGroup.Wait()
}

func GetChildContext() context.Context {
	return ctx
}
