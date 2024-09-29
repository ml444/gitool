package conf

import log "github.com/ml444/glog"

func InitLogger() {
	var err error

	// setting logger
	opts := []log.OptionFunc{
		log.SetLoggerName("gitool"),
		log.SetLoggerLevel(log.DebugLevel),
		log.SetColorRender(false),
		log.SetRecordCaller(0),
	}
	err = log.InitLog(opts...)
	if err != nil {
		println(err.Error())
		return
	}
}
