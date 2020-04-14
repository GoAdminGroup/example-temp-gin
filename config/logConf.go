package config

import (
	"github.com/GoAdminGroup/example-temp-gin/pkg/zlog"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// Initialization log
//	zap:
//    AtomicLevel: -1 # DebugLevel:-1 InfoLevel:0 WarnLevel:1 ErrorLevel:2
//    FieldsAuto: false # is use auto Fields key set
//    Fields:
//      Key: key
//      Val: val
//    Development: true # is open Open file and line number
//    Encoding: console # output format, only use console or json, default is console
//    rotate:
//      Filename: log/example-temp-gin.log # Log file path
//      MaxSize: 16 # Maximum size of each zlog file, Unit: M
//      MaxBackups: 10 # How many backups are saved in the zlog file
//      MaxAge: 7 # How many days can the file be keep, Unit: day
//      Compress: true # need compress
//    EncoderConfig:
//      TimeKey: time
//      LevelKey: level
//      NameKey: logger
//      CallerKey: caller
//      MessageKey: msg
//      StacktraceKey: stacktrace
//      TimeEncoder: ISO8601TimeEncoder # ISO8601TimeEncoder EpochMillisTimeEncoder EpochNanosTimeEncoder EpochTimeEncoder default is ISO8601TimeEncoder
//      EncodeDuration: SecondsDurationEncoder # NanosDurationEncoder SecondsDurationEncoder StringDurationEncoder default is SecondsDurationEncoder
//      EncodeLevel: CapitalColorLevelEncoder # CapitalLevelEncoder CapitalColorLevelEncoder LowercaseColorLevelEncoder LowercaseLevelEncoder default is CapitalLevelEncoder
//      EncodeCaller: ShortCallerEncoder # ShortCallerEncoder FullCallerEncoder default is FullCallerEncoder
func (c *ConfFile) initLog() error {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        viper.GetString("zap.EncoderConfig.TimeKey"),
		LevelKey:       viper.GetString("zap.EncoderConfig.LevelKey"),
		NameKey:        viper.GetString("zap.EncoderConfig.NameKey"),
		CallerKey:      viper.GetString("zap.EncoderConfig.CallerKey"),
		MessageKey:     viper.GetString("zap.EncoderConfig.MessageKey"),
		StacktraceKey:  viper.GetString("zap.EncoderConfig.StacktraceKey"),
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    filterZapEncodeLevel(),
		EncodeTime:     filterZapTimeEncoder(), // ISO8601TimeEncoder ISO8601 UTC time
		EncodeDuration: filterZapDurationEncoder(),
		EncodeCaller:   filterZapCallerEncoder(),
	}

	atomicLevel := zap.NewAtomicLevelAt(filterZapAtomicLevelByViper()) // log Level

	rotateLogger := lumberjack.Logger{
		Filename:   viper.GetString("zap.rotate.Filename"), // Log file path
		MaxSize:    viper.GetInt("zap.rotate.MaxSize"),     // Maximum size of each log file Unit: M
		MaxBackups: viper.GetInt("zap.rotate.MaxBackups"),  // How many backups are saved in the log file
		MaxAge:     viper.GetInt("zap.rotate.MaxAge"),      // How many days can the file be keep
		Compress:   viper.GetBool("zap.rotate.Compress"),   // need compress
	}

	encoder := filterZapEncoder(encoderConfig)

	core := zapcore.NewCore(
		encoder, // Encoder configuration
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&rotateLogger),
		), // Print to console and file
		atomicLevel, // Log level
	)
	var filed zap.Option
	if viper.GetBool("zap.FieldsAuto") {
		filed = zap.Fields( //the initialization field
			zap.String(viper.GetString("zap.Fields.Key"), viper.GetString("zap.Fields.Val")),
		)
	}

	var logZap *zap.Logger
	if viper.GetBool("zap.Development") {
		if filed != nil {
			logZap = zap.New(core, zap.AddCaller(), zap.Development(), filed)
		} else {
			logZap = zap.New(core, zap.AddCaller(), zap.Development())
		}
	} else {
		if filed != nil {
			logZap = zap.New(core, filed)
		} else {
			logZap = zap.New(core)
		}
	}
	zlog.NewZapLog(logZap, logZap.Sugar())
	return nil
}

// default ISO8601TimeEncoder
func filterZapTimeEncoder() func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	switch viper.GetString("zap.EncoderConfig.TimeEncoder") {
	default:
		return zapcore.ISO8601TimeEncoder
	case "ISO8601TimeEncoder":
		return zapcore.ISO8601TimeEncoder
	case "EpochMillisTimeEncoder":
		return zapcore.EpochMillisTimeEncoder
	case "EpochNanosTimeEncoder":
		return zapcore.EpochNanosTimeEncoder
	case "EpochTimeEncoders":
		return zapcore.EpochTimeEncoder
	}
}

// default SecondsDurationEncoder
func filterZapDurationEncoder() func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	switch viper.GetString("zap.EncoderConfig.EncodeDuration") {
	default:
		return zapcore.SecondsDurationEncoder
	case "SecondsDurationEncoder":
		return zapcore.SecondsDurationEncoder
	case "NanosDurationEncoder":
		return zapcore.NanosDurationEncoder
	case "StringDurationEncoder":
		return zapcore.StringDurationEncoder
	}
}

// default FullCallerEncoder
func filterZapCallerEncoder() func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	switch viper.GetString("zap.EncoderConfig.EncodeCaller") {
	default:
		return zapcore.FullCallerEncoder
	case "FullCallerEncoder":
		return zapcore.FullCallerEncoder
	case "ShortCallerEncoder":
		return zapcore.ShortCallerEncoder
	}

}

// default CapitalLevelEncoder
func filterZapEncodeLevel() func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch viper.GetString("zap.EncoderConfig.EncodeLevel") {
	default:
		return zapcore.CapitalLevelEncoder
	case "CapitalLevelEncoder":
		return zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder":
		return zapcore.CapitalColorLevelEncoder
	case "LowercaseLevelEncoder":
		return zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder":
		return zapcore.LowercaseColorLevelEncoder
	}
}

func filterZapEncoder(encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
	var encoder zapcore.Encoder
	zapEncoding := viper.GetString("zap.Encoding")
	switch zapEncoding {
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	return encoder
}

func filterZapAtomicLevelByViper() zapcore.Level {
	var atomViper zapcore.Level
	switch viper.GetInt("zap.AtomicLevel") {
	default:
		atomViper = zap.InfoLevel
	case -1:
		atomViper = zap.DebugLevel
	case 0:
		atomViper = zap.InfoLevel
	case 1:
		atomViper = zap.WarnLevel
	case 2:
		atomViper = zap.ErrorLevel
	}
	return atomViper
}
