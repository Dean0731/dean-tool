package utils

import (
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

type GoroutinePool struct {
	pool *ants.Pool
	wg   sync.WaitGroup
	mu   sync.Mutex
}

// NewGoroutinePool 创建一个新的GoroutinePool实例
func NewGoroutinePool(maxWorkers int, expiryDuration time.Duration) (*GoroutinePool, error) {
	pool, err := ants.NewPool(maxWorkers, func(opts *ants.Options) {
		opts.ExpiryDuration = expiryDuration
	})
	if err != nil {
		return nil, err
	}
	return &GoroutinePool{pool: pool}, nil
}

// Submit 提交一个任务到池中，并增加WaitGroup计数
func (gp *GoroutinePool) Submit(task func()) error {
	gp.mu.Lock()
	defer gp.mu.Unlock()

	gp.wg.Add(1)
	err := gp.pool.Submit(func() {
		defer gp.wg.Done()
		task()
	})
	if err != nil {
		gp.wg.Done() // 如果提交失败，需要手动减少计数
	}
	return err
}

// Wait 等待所有任务完成
func (gp *GoroutinePool) Wait() {
	gp.wg.Wait()
}

// Release 释放资源
func (gp *GoroutinePool) Release() {
	gp.pool.Release()
}
