// Package modified snowflake provides a very simple 32bits id generator, the unique only can last 9 hours .
package utils

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	// Epoch is set to the twitter snowflake epoch of Nov 04 2010 01:42:54 UTC in milliseconds
	// You may customize this to set a different epoch for your application.
	Epoch int64 = 1288834974657

	// NodeBits holds the number of bits to use for Node
	// Remember, you have 3 bits for node
	NodeBits uint8 = 3

	// StepBits holds the number of bits to use for Step
	// Remember, you have 3 bits for Step
	StepBits uint8 = 3

	//shift keep 25 bits to keep million second, only in 9 hours not repeat
	TimeKeepShift uint8 = 25

	// DEPRECATED: the below four variables will be removed in a future release.
	mu        sync.Mutex
	nodeMax   int32 = -1 ^ (-1 << NodeBits)
	nodeMask        = nodeMax << StepBits
	stepMask  int32 = -1 ^ (-1 << StepBits)
	timeMask  int64 = -1 ^ (-1 << TimeKeepShift)
	timeShift       = NodeBits + StepBits
	nodeShift       = StepBits
)

// A Node struct holds the basic information needed for a snowball generator
// node
type Node struct {
	mu    sync.Mutex
	epoch time.Time
	time  int32
	node  int32
	step  int32

	timeMask  int64
	nodeMax   int32
	nodeMask  int32
	stepMask  int32
	timeShift uint8
	nodeShift uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int32

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(node int32) (*Node, error) {
	// re-calc in case custom NodeBits or StepBits were set
	// DEPRECATED: the below block will be removed in a future release.
	mu.Lock()
	nodeMax = -1 ^ (-1 << NodeBits)
	nodeMask = nodeMax << StepBits
	stepMask = -1 ^ (-1 << StepBits)
	timeShift = NodeBits + StepBits
	nodeShift = StepBits
	timeMask = -1 ^ (-1 << TimeKeepShift)
	mu.Unlock()

	n := Node{}
	n.node = node
	n.nodeMax = -1 ^ (-1 << NodeBits)
	n.nodeMask = n.nodeMax << StepBits
	n.stepMask = -1 ^ (-1 << StepBits)
	n.timeMask = -1 ^ (-1 << TimeKeepShift)
	n.timeShift = NodeBits + StepBits
	n.nodeShift = StepBits

	if n.node < 0 || n.node > n.nodeMax {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(int64(n.nodeMax), 10))
	}

	var curTime = time.Now()
	// add time.Duration to curTime to make sure we use the monotonic clock if available
	n.epoch = curTime.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(curTime))

	return &n, nil
}

// Generate creates and returns a unique 32bits snowball ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (n *Node) Generate() ID {

	n.mu.Lock()

	now := (time.Since(n.epoch).Nanoseconds() / 1000000) & n.timeMask
	if int32(now) == n.time {
		fmt.Println("the same")
		fmt.Println(int32(now))
		n.step = (n.step + 1) & n.stepMask

		if n.step == 0 {
			for int32(now) <= n.time {
				now = time.Since(n.epoch).Nanoseconds() / 1000000 & n.timeMask
			}
		}
	} else {
		n.step = 0
	}

	n.time = int32(now)

	r := ID((int32(now))<<n.timeShift |
		(n.node << n.nodeShift) |
		(n.step),
	)

	n.mu.Unlock()
	return r
}
