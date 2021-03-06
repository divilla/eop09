package interfaces

type Logger interface {
	Panic(i ...interface{})
	Panicf(format string, args ...interface{})
	Error(i ...interface{})
	Errorf(format string, args ...interface{})
	Info(i ...interface{})
	Infof(format string, args ...interface{})
	Fatal(i ...interface{})
	Fatalf(format string, args ...interface{})
	//Infoj(j log.JSON)
	//Warn(i ...interface{})
	//Warnf(format string, args ...interface{})
	//Warnj(j log.JSON)
	//Errorj(j log.JSON)
	//Fatalj(j log.JSON)
	//Panicj(j log.JSON)
}
