/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/10/31
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 * Desc: compatible log framework
 *******************************************************************************/
package log

import (
	//"os"
	"os"
	log "github.com/sirupsen/logrus"
	"time"
	"path"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/pkg/errors"
)

var format = &log.TextFormatter{}

var logger = NewLogger("aliens", format, true)

//调试版本日志带颜色
func Init(debug bool) {
	format.ForceColors = debug
	format.DisableTimestamp = debug
}

func NewLogger(name string, formatter log.Formatter, local bool) *log.Logger {
	logger := log.New()
	logger.Formatter = formatter
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.Out = os.Stdout
	// Only log the warning severity or above.
	logger.Level = log.DebugLevel
	if local {
		configLocalFilesystemLogger(name, logger, 30 * 24 * time.Hour, 24 * time.Hour)
	}
	return logger
}

// config logrus log to amqp  rabbitMQ
//func ConfigAmqpLogger(server, username, password, exchange, exchangeType, virtualHost, routingKey string) {
//	hook := logrus_amqp.NewAMQPHookWithType(server, username, password, exchange, exchangeType, virtualHost, routingKey)
//	log.AddHook(hook)
//}

// config logrus log to elasticsearch
//func ConfigESLogger(esUrl string, esHOst string, index string) {
//	client, err := elastic.NewClient(elastic.SetURL(esUrl))
//	if err != nil {
//		log.Errorf("config es logger error. %+v", errors.WithStack(err))
//	}
//	esHook, err := elogrus.NewElasticHook(client, esHOst, log.DebugLevel, index)
//	if err != nil {
//		log.Errorf("config es logger error. %+v", errors.WithStack(err))
//	}
//	log.AddHook(esHook)
//}

//config logrus log to local file
func configLocalFilesystemLogger(name string, logger *log.Logger, maxAge time.Duration, rotationTime time.Duration) {
	logPath := ""
	logFileName := name + ".log"
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, logger.Formatter)
	logger.AddHook(lfHook)
}

//Debugf Printf Infof Warnf Warningf Errorf Panicf Fatalf

//做一层适配，方便后续切换到其他日志框架或者自己写
func Debug(arg ...interface{}) {
	logger.Debug(arg...)
}

func Print(arg ...interface{}) {
	logger.Print(arg...)
}

func Info(arg ...interface{}) {
	logger.Info(arg...)
}

func Warn(arg ...interface{}) {
	logger.Warn(arg...)
}

func Error(arg ...interface{}) {
	logger.Error(arg...)
}

func Panic(arg ...interface{}) {
	logger.Panic(arg...)
}

func Fatal(arg ...interface{}) {
	logger.Fatal(arg...)
}


//-----------format

func Debugf(format string, arg ...interface{}) {
	logger.Debugf(format, arg...)
}

func Printf(format string, arg ...interface{}) {
	logger.Printf(format, arg...)
}

func Infof(format string, arg ...interface{}) {
	logger.Infof(format, arg...)
}

func Warnf(format string, arg ...interface{}) {
	logger.Warnf(format, arg...)
}

func Errorf(format string, arg ...interface{}) {
	logger.Errorf(format, arg...)
}

func Panicf(format string, arg ...interface{}) {
	logger.Panicf(format, arg...)
}

func Fatalf(format string, arg ...interface{}) {
	logger.Fatalf(format, arg...)
}

