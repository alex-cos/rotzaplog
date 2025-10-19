package rotzaplog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kjk/common/filerotate"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type UTCClock struct{}

func (UTCClock) Now() time.Time {
	return time.Now().UTC()
}

func (UTCClock) NewTicker(duration time.Duration) *time.Ticker {
	return time.NewTicker(duration)
}

func InitFileLogger(logpath string, level zapcore.LevelEnabler, utc bool, verbose bool) *zap.Logger {
	var logger *zap.Logger

	dir := filepath.Dir(logpath)
	filename := filepath.Base(logpath)
	ext := filepath.Ext(filename)
	basename := strings.TrimSuffix(filename, ext)
	fileconfig := filerotate.Config{
		DidClose: func(path string, didRotate bool) {
			// By default do noting
		},
		PathIfShouldRotate: func(creationTime time.Time, now time.Time) string {
			if creationTime.YearDay() == now.YearDay() {
				return ""
			}
			name := fmt.Sprintf("%s_%s%s", basename, now.Format("2006-01-02"), ext)
			return filepath.Join(dir, name)
		},
	}
	file, err := filerotate.New(&fileconfig)
	if err != nil {
		panic(err)
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	writer := zapcore.AddSync(file)

	if verbose {
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		core := zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, level),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
		)
		logger = zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	} else {
		core := zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, level),
		)
		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	}
	if utc {
		logger.WithOptions(zap.WithClock(UTCClock{}))
	}

	return logger
}

func InitConsoleLogger(level zapcore.LevelEnabler, utc bool) *zap.Logger {
	var (
		logger *zap.Logger
		core   zapcore.Core
	)

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	core = zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	if utc {
		logger.WithOptions(zap.WithClock(UTCClock{}))
	}

	return logger
}
