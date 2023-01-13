package signal

import (
	"log"
	"os"
	"os/signal"
)

type SignTable struct {
	table map[os.Signal]SignHandler // signal table handler
	ch    chan struct{}
}

// NewSignTable 信号表
func NewSignTable() *SignTable {
	return &SignTable{
		table: make(map[os.Signal]SignHandler),
		ch:    make(chan struct{}),
	}
}

func (st *SignTable) Shutdown() {
	close(st.ch)
}

// Register 将信号量和对应的操作注入到信号表
func (st *SignTable) Register(s os.Signal, handler SignHandler) {
	if _, ok := st.table[s]; !ok {
		st.table[s] = handler
	} else {
		log.Fatalf("请勿重复注入相同信号量(%s)操作", s.String())
	}
}

func (st *SignTable) StartSignalHandler() {
	go st.signalHandler()
}

func (st *SignTable) signalHandler() {
	var signals []os.Signal
	for s, _ := range st.table {
		signals = append(signals, s)
	}
	c := make(chan os.Signal, len(signals))
	signal.Notify(c, signals...)

	for {
		select {
		case sig := <-c: // 系统传递的信号
			handler := st.table[sig]
			handler(sig)
		case <-st.ch: // 信号量表能传过来的信号
			return
		}
	}
}
