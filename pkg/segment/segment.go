package segment

import (
	"context"
	"errors"
	"github.com/yungsem/goleaf/pkg/db"
	"github.com/yungsem/gotool/maths"
	"time"
)

// UsedPercentMax 表示号段使用比例达到该值时，就要扩充号段
const UsedPercentMax float64 = 40

// 号段 [0,1000) ，其中 Left = 0 ，Right = 1000
// Offset 表示已经分配的 ID 的偏移量
// 比如已经分配了 10 个 ID ，Offset 就是 10
// 本号段剩余 ID 的个数：Right - Left - Offset
type Segment struct {
	Offset int
	Left   int
	Right  int
}

// AllCount 返回该号段的总 ID 数
func (s *Segment) AllCount() int {
	return s.Right - s.Left + s.Offset
}

// RemainPercent 返回该号段的已使用 ID 数的百分比
func (s *Segment) UsedPercent() float64 {
	f := maths.Round(float64(s.Offset)/float64(s.AllCount()), 2) * 100
	return f
}

// IsEmpty 返回该号段是否已经分配完所有 ID ，没有剩余 ID 可以分配了
func (s *Segment) IsEmpty() bool {
	return s.Left == s.Right
}

// NextId 返回该号段的下一个 ID
func (s *Segment) NextId() int {
	s.Left += 1
	s.Offset += 1
	return s.Left
}

// LoadSegment 根据 bizTag 加载一个号段 segment
func LoadSegment(bizTag string) (s *Segment, err error) {
	maxId, size, err := updateSegment(bizTag)
	if err != nil {
		return
	}
	s = buildSegment(maxId, size)
	return
}

// buildSegment 构建一个 segment
// end 为新构建 segment 的上边界
// size 为新构建 segment 的大小
func buildSegment(end int, size int) (s *Segment) {
	// 获取下一个号段
	// 下一个号段的 startId 就是本号段的 endId
	s = &Segment{}
	s.Left = end - size
	s.Right = end

	return
}

// updateSegment 将 bizTag 对应号段的 max_id 更新为下一段的 max_id ，并返回新的 max_id
// 执行下面两句 sql
// update segments set max_id = max_id + size where biz_tag = 'xxx'
// select max_id, size from segments where biz_tag = 'xxx'
func updateSegment(bizTag string) (maxId int, size int, err error) {
	// 定义 ctx ，超时时间为 2 秒
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(2000)*time.Millisecond)
	defer cancelFunc()

	// 开启事务
	tx, err := db.DB.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// 更新 bizTag 的下一个号段的起始 ID
	// new_start_id = old_start_id + size
	stmt, err := tx.PrepareContext(ctx, "update segments set max_id = max_id + size where biz_tag = ?")
	if err != nil {
		tx.Rollback()
		return
	}

	result, err := stmt.ExecContext(ctx, bizTag)
	if err != nil {
		tx.Rollback()
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return
	}
	if rowsAffected == 0 {
		err = errors.New("biz_tag not found")
		tx.Rollback()
		return
	}

	// 查询刚更新的记录
	stmt, err = tx.PrepareContext(ctx, "select max_id, size from segments where biz_tag = ?")
	if err != nil {
		tx.Rollback()
		return
	}

	rows, err := stmt.QueryContext(ctx, bizTag)
	if err != nil {
		tx.Rollback()
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&maxId, &size)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return
	}

	return
}
