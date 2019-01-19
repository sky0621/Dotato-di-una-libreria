package logger

import (
	"Dotato-di-una-libreria/backend/util"
	"os"
	"time"

	"github.com/jinzhu/gorm"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger ...
func ZapLogger() *zap.Logger {
	logLvl := zapcore.InfoLevel
	if util.IsLocal() {
		logLvl = zapcore.DebugLevel
	}
	encCfg := zap.NewProductionEncoderConfig()
	encCfg.EncodeTime = jstTimeEncoder
	lgr := zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), os.Stdout, logLvl))
	return lgr
}

func jstTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	const layout = "2006-01-02 15:04:05"
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	enc.AppendString(t.In(jst).Format(layout))
}

// AppLogger ...
type AppLogger interface {
	Path(path string) AppLogger
	RequestID(requestID string) AppLogger
	// TODO: 他にも各コントローラー内で固定で追加したい要素が発生したら、ここにセッターを追加し、merge()内でも処理を追加！

	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})

	// gorm.loggerインタフェースを実装することで、Gormのログ出力をAppLoggerでフックする。
	Print(v ...interface{})

	merge(keysAndValues ...interface{}) []interface{}
}

type appLogger struct {
	sugar     *zap.SugaredLogger
	requestID string
	path      string
	severity  string // StackdriverLoggingでのLogSeverityの表示用	https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity
}

// NewLogger ...
func NewLogger(sugar *zap.SugaredLogger) AppLogger {
	return &appLogger{sugar: sugar, requestID: "-", path: "-"}
}

func (l *appLogger) Path(path string) AppLogger {
	l.path = path
	return l
}

func (l *appLogger) RequestID(requestID string) AppLogger {
	l.requestID = requestID
	return l
}

// Debugw ...
func (l *appLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sugar.Debugw(msg, l.setSeverity("DEBUG").merge(keysAndValues...)...)
}

// Infow ...
func (l *appLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.sugar.Infow(msg, l.setSeverity("INFO").merge(keysAndValues...)...)
}

// Warnw ...
func (l *appLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.sugar.Warnw(msg, l.setSeverity("WARNING").merge(keysAndValues...)...)
}

// Errorw ...
func (l *appLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sugar.Errorw(msg, l.setSeverity("ERROR").merge(keysAndValues...)...)
}

// Fatalw ...
func (l *appLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.sugar.Fatalw(msg, l.setSeverity("CRITICAL").merge(keysAndValues...)...)
}

// Panicw ...
func (l *appLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.sugar.Panicw(msg, l.setSeverity("EMERGENCY").merge(keysAndValues...)...)
}

func (l *appLogger) merge(keysAndValues ...interface{}) []interface{} {
	kvs := []interface{}{"severity", l.severity, "path", l.path, "request_id", l.requestID}
	kvs = append(kvs, keysAndValues...)
	return kvs
}

// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity
func (l *appLogger) setSeverity(s string) AppLogger {
	l.severity = s
	return l
}

// Print ...
func (l *appLogger) Print(v ...interface{}) {
	if v == nil || len(v) == 0 {
		return
	}
	l.setSeverity("INFO").Infow("[GORM]", gorm.LogFormatter(v...)...)
}
