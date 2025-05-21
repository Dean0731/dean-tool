// signalhandler/signalhandler.go

package handler

import (
	"os"
	"os/signal"
	"sync"
)

// 定义信号处理函数类型
type SignalFunc func()

// 信号映射表，key是信号，value是对应的处理函数
var signals = make(map[os.Signal]SignalFunc)
var mutex sync.Mutex

// 注册新的信号处理函数
func RegisterSignal(sig os.Signal, action SignalFunc) {
	mutex.Lock()
	defer mutex.Unlock()
	signals[sig] = action
}

// 启动信号监听器，在独立的goroutine中运行
func StartListener() {
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, getRegisteredSignals()...)

		for sig := range sigChan {
			if action, exists := getActionForSignal(sig); exists {
				action()
			}
		}
	}()
}

// 获取已注册的所有信号
func getRegisteredSignals() []os.Signal {
	mutex.Lock()
	defer mutex.Unlock()
	sigs := make([]os.Signal, 0, len(signals))
	for sig := range signals {
		sigs = append(sigs, sig)
	}
	return sigs
}

// 根据信号获取对应的处理函数
func getActionForSignal(sig os.Signal) (SignalFunc, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	action, exists := signals[sig]
	return action, exists
}
