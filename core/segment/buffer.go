package segment

import (
	inits2 "github.com/yungsem/goleaf/inits"
	"sync"
)

// bufferSize 表示 buffer 的号段 size
// 默认情况下，每个 buffer 拥有两个号段
const bufferSize = 2

// BizBuffer 代表业务的 Buffer
// 每个业务都有自己独立的 Buffer ，业务与业务之间获取 ID 是独立的，
// 所以它们持有各自独立的锁
type BizBuffer struct {
	BizTag     string
	Mutex      sync.Mutex
	expandChan chan error
	Segments   []*Segment
}

// NewBizBuffer 创建一个 BizBuffer
// 新创建的 BizBuffer 没有号段
func NewBizBuffer(bizTag string) *BizBuffer {
	inits2.Log.Info("新建buffer")
	var segments []*Segment

	return &BizBuffer{
		BizTag:     bizTag,
		expandChan: make(chan error),
		Segments:   segments,
	}
}

// removeFirstSegment 删除 buffer 的第一个 segment
func (b *BizBuffer) removeFirstSegment() {
	b.Segments = b.Segments[1:]
}

// isEmpty 判断 buffer 的 segments 是否为空
func (b *BizBuffer) isEmpty() bool {
	return b.Segments == nil || len(b.Segments) == 0
}

// expand 扩充该 buffer 的号段
func (b *BizBuffer) expand() {
	seg, err := LoadSegment(b.BizTag)
	if err != nil {
		b.expandChan <- err
		return
	}
	b.Segments = append(b.Segments, seg)
	b.expandChan <- nil
}

// NextId 获取下一个 ID
func (b *BizBuffer) NextId() (int, error) {
	if b.isEmpty() || (b.Segments[0].UsedPercent() >= UsedPercentMax && len(b.Segments) < bufferSize) {
		inits2.Log.Info("开启新的goroutine，加载新的segment")
		go b.expand()

		select {
		case err := <-b.expandChan:
			if err != nil {
				return 0, err
			}
		}
	}

	seg0 := b.Segments[0]
	nextId := seg0.NextId()

	if seg0.IsEmpty() {
		b.removeFirstSegment()
	}

	return nextId, nil
}
