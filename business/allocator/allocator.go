package allocator

import (
	"github.com/yungsem/goleaf/business/segment"
	"sync"
)

var (
	// 全局的 bizBufferMap ，存储各个业务对应的号段 buffer
	bizBufferMap = make(map[string]*segment.BizBuffer)
	// 全局的锁，用于全局加锁
	mutex        sync.Mutex
)

// AllocateId 根据 bizTag 对应业务的下一个 ID
func AllocateId(bizTag string) (int, error) {
	// 从全局的 bizBufferMap 中取出 bizTag 对应的 bizBuffer
	// 为避免重复创建新的 bizBuffer ，此处使用全局锁
	mutex.Lock()

	bizBuffer, ok := bizBufferMap[bizTag]

	// 该 bizTag 对应的 bizBuffer 尚不存在，则需要创建
	if !ok {
		bizBuffer = segment.NewBizBuffer(bizTag)
		bizBufferMap[bizTag] = bizBuffer
	}
	mutex.Unlock()

	// 根据 bizTag 获取业务的下一个 ID
	// 各个业务使用各自独立的锁
	bizBuffer.Mutex.Lock()
	defer bizBuffer.Mutex.Unlock()
	nextId, err := bizBuffer.NextId()
	if err != nil {
		return 0, err
	}

	return nextId, nil
}

// Info 表示号段信息，用于监控查询
type Info struct {
	BizTag string `json:"bizTag"`
	Left   int    `json:"left"`
	Right  int    `json:"right"`
	Offset int    `json:"offset"`
}

// BizBufferMapInfo 获取号段信息，用于监控查询
func BizBufferMapInfo(bizTag string) []*Info {
	var list []*Info
	for k, v := range bizBufferMap {
		if bizTag != "" && k != bizTag{
			continue
		}
		for _, seg := range v.Segments {
			info := Info{
				BizTag: k,
				Left:   seg.Left,
				Right:  seg.Right,
				Offset: seg.Offset,
			}
			list = append(list, &info)
		}
	}
	return list
}
