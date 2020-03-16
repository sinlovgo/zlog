package zlog

import "go.uber.org/zap"

type ZapLogger struct {
	ConfigPath string
	Log        *zap.Logger
	Sugar      *zap.SugaredLogger
}

// Logger is the global variable for the object of lager.Logger
var Logger LoggerGlobal

type logger struct {
	ConfigPath string
	Log        *zap.Logger
	Sugar      *zap.SugaredLogger
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.Sugar.Debugf(format, args)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.Sugar.Infof(format, args)
}

func (l *logger) Debug(args ...interface{}) {
	l.Sugar.Debug(args)
}

func (l *logger) Info(args ...interface{}) {
	l.Sugar.Info(args)
}

func (l *logger) Warn(args ...interface{}) {
	l.Sugar.Warn(args)
}

func (l *logger) Error(args ...interface{}) {
	l.Sugar.Error(args)
}

func (l *logger) DPanic(args ...interface{}) {
	l.Sugar.DPanic(args)
}

func (l *logger) Panic(args ...interface{}) {
	l.Sugar.Panic(args)
}

func (l *logger) Fatal(args ...interface{}) {
	l.Sugar.Fatal(args)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.Sugar.Warnf(template, args)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.Sugar.Errorf(template, args)
}

func (l *logger) DPanicf(template string, args ...interface{}) {
	l.Sugar.DPanicf(template, args)
}

func (l *logger) Panicf(template string, args ...interface{}) {
	l.Sugar.Panicf(template, args)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.Sugar.Fatalf(template, args)
}

func (l *logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.Sugar.Debugw(msg, keysAndValues)
}

func (l *logger) Infow(msg string, keysAndValues ...interface{}) {
	l.Sugar.Infow(msg, keysAndValues)
}

func (l *logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.Sugar.Warnw(msg, keysAndValues)
}

func (l *logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.Sugar.Errorw(msg, keysAndValues)
}

func (l *logger) DPanicw(msg string, keysAndValues ...interface{}) {
	l.Sugar.DPanicw(msg, keysAndValues)
}

func (l *logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.Sugar.Panicw(msg, keysAndValues)
}

func (l *logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.Sugar.Fatalw(msg, keysAndValues)
}

func (l *logger) Sync() error {
	panic("implement me")
}

//Logger is a interface
type LoggerGlobal interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	DPanic(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync() error
}

// newLog new log
func newLogger(configPath string, log *zap.Logger, sugar *zap.SugaredLogger) LoggerGlobal {
	return &logger{
		ConfigPath: configPath,
		Log:        log,
		Sugar:      sugar,
	}
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	Logger.Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	Logger.Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	Logger.Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	Logger.Error(args)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See zapcore.DPanicLevel for details.)
func DPanic(args ...interface{}) {
	Logger.DPanic(args)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	Logger.Panic(args)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	Logger.Fatal(args)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	Logger.Warnf(template, args)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	Logger.Errorf(template, args)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See zapcore.DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	Logger.DPanicf(template, args)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	Logger.Panicf(template, args)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	Logger.Fatalf(template, args)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	Logger.Debugw(msg, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	Logger.Infow(msg, keysAndValues)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	Logger.Warnw(msg, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	Logger.Errorw(msg, keysAndValues)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	Logger.DPanicw(msg, keysAndValues)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	Logger.Panicw(msg, keysAndValues)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	Logger.Fatalw(msg, keysAndValues)
}

// Sync flushes any buffered log entries.
func Sync() error {
	return Logger.Sync()
}
