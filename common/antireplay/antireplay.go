package antireplay


/*
	@desc antie replay
	@question 这个模块是干嘛的
*/

import (
	"sync"
	"time"

	cuckoo "github.com/seiflotfy/cuckoofilter"
	/*
	@desc cuckoo 布隆过滤器
	*/
)

func NewAntiReplayWindow(antiReplayTime int64) *AntiReplayWindow {
	arw := &AntiReplayWindow{}
	arw.AntiReplayTime = antiReplayTime
	return arw
}

type AntiReplayWindow struct {
	lock           sync.Mutex
	poolA          *cuckoo.Filter
	poolB          *cuckoo.Filter
	/*
	@desc poolA poolB是干嘛的
	*/
	lastSwapTime   int64
	PoolSwap       bool
	AntiReplayTime int64
}

func (aw *AntiReplayWindow) Check(sum []byte) bool {
	aw.lock.Lock()

	if aw.lastSwapTime == 0 {
		aw.lastSwapTime = time.Now().Unix()
		aw.poolA = cuckoo.NewFilter(100000)
		aw.poolB = cuckoo.NewFilter(100000)
	}

	tnow := time.Now().Unix()
	timediff := tnow - aw.lastSwapTime

	if timediff >= aw.AntiReplayTime {
		if aw.PoolSwap {
			aw.PoolSwap = false
			aw.poolA.Reset()
		} else {
			aw.PoolSwap = true
			aw.poolB.Reset()
		}
		aw.lastSwapTime = tnow
	}

	ret := aw.poolA.InsertUnique(sum) && aw.poolB.InsertUnique(sum)
	aw.lock.Unlock()
	return ret
}
