package signal

import (
	"log"
	"os"
	"syscall"
)

type SignHandler func(s os.Signal)

// TermHandler deal with the signal that should terminate the process
func TermHandler(s os.Signal) {
	log.Fatalf("termHandler(): receive signal[%v], terminate.", s)
}

// IgnoreHandler deal with the signal that should be ignored
func IgnoreHandler(s os.Signal) {
	log.Printf("ignoreHandler(): receive signal[%v], ignore.", s)
}

// RegisterSignalHandlers register signal handlers
func RegisterSignalHandlers(signTable *SignTable) {
	// term handlers
	signTable.Register(syscall.SIGTERM, TermHandler)

	// ignore handlers
	signTable.Register(syscall.SIGHUP, IgnoreHandler)
	signTable.Register(syscall.SIGQUIT, IgnoreHandler)
	signTable.Register(syscall.SIGILL, IgnoreHandler)
	signTable.Register(syscall.SIGTRAP, IgnoreHandler)
	signTable.Register(syscall.SIGABRT, IgnoreHandler)
}
