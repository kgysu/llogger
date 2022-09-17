package main

import "github.com/kgysu/llogger/pkg/llogger"

func main() {
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

}
