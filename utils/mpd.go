package utils

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var noTaskExit = time.Second * time.Duration(10)

type taskStatus struct {
	ID       int
	Name     string
	Finished bool
}

type BarTask struct {
	MaxWorker       int
	PrintLineNum    int
	taskIndex       int
	pool            *GoroutinePool
	TaskStatusArray []*taskStatus
	Refresh         time.Duration
	firstPrint      bool
	hashTask        bool
	mutex           sync.Mutex
	startTime       *time.Time
}

func (bar *BarTask) getProcess() ([]string, int) {
	var message []string
	var unfinished int
	for _, status := range bar.TaskStatusArray {
		if status.Finished {
			message = append(message, fmt.Sprintf("任务：%s, \t已完成", status.Name))
		} else {
			unfinished++
			message = append(message, fmt.Sprintf("任务：%s, \t执行中...", status.Name))
		}
	}
	return message, unfinished
}
func (bar *BarTask) clearLines() {
	for i := 0; i < bar.PrintLineNum; i++ {
		fmt.Print("\033[A\033[2K") // 先上移一行然后清除该行
	}
}
func (bar *BarTask) getTaskIndex() int {
	bar.mutex.Lock()
	defer bar.mutex.Unlock()
	bar.taskIndex++
	return bar.taskIndex
}

func (bar *BarTask) Add(fun func(), taskName string) {
	bar.pool.Submit(func() {
		status := &taskStatus{ID: bar.getTaskIndex(), Name: taskName, Finished: false}
		bar.TaskStatusArray = append(bar.TaskStatusArray, status)
		fun()
		status.Finished = true
	})
	bar.hashTask = true
}

func (bar *BarTask) Finish() {
	bar.pool.Wait()
	bar.pool.Release()
}
func NewBarTask(maxWorker int) *BarTask {
	pool, _ := NewGoroutinePool(maxWorker+1, time.Minute)
	t := time.Now()
	bar := &BarTask{
		MaxWorker:    maxWorker + 1,
		PrintLineNum: 0,
		taskIndex:    0,
		pool:         pool,
		Refresh:      time.Second,
		firstPrint:   true,
		startTime:    &t,
	}
	pool.Submit(func() {
		for {
			time.Sleep(bar.Refresh)
			if !bar.hashTask {
				if time.Now().After(bar.startTime.Add(noTaskExit)) {
					break
				}
				continue
			}
			message, unfinished := bar.getProcess()

			if !bar.firstPrint {
				bar.clearLines()
			}
			fmt.Println(strings.Join(message, "\n"))
			if unfinished == 0 {
				break
			}
			bar.PrintLineNum = len(message)
			if bar.firstPrint {
				bar.firstPrint = false
			}
		}
	})
	return bar
}
