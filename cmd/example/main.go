package main

import (
	"github.com/kgysu/llogger/pkg/llogger"
)

func main() {
	llogger.Infoln("info", 1)

	llogger.Traceln("trace")
	llogger.Debugln("debug")
	llogger.Infoln("info")
	llogger.Warnln("warn")
	llogger.Errorln("error")

	ll := llogger.NewLogger(llogger.TraceLevel)
	ll.Traceln("trace")
	ll.Debugln("debug")
	ll.Infoln("info")
	ll.Warnln("warn")
	ll.Errorln("error")

	llogger.Traceln("trace")
	llogger.Debugln("debug")
	llogger.Infoln("info")
	llogger.Warnln("warn")
	llogger.Errorln("error")

	testValue := 1

	llogger.Infoln("info", testValue)
	llogger.Infoln("info", testValue, testValue)
	llogger.Infof("T %d : %d", testValue, testValue)

	test()
}

func test() {
	llogger.Infoln("Test")
}
