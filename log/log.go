/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/4/9 2:23 下午
 * @Description:日志文件切分设置、通用日志输出函数接口封装，全项目使用自己封装的日志输出函数
 */

package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"strings"
	"time"
)
func InitLog(){
	LogFile()
	//日志写入文件时，禁用控制台颜色
	gin.DisableConsoleColor()

	//写入日志文件
	f, _ := os.OpenFile("./mikasa_ufs.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend)//0644?
	log.SetOutput(f)
	//gin.DefaultWriter = io.MultiWriter(f)

	//如果需要同时写入日志文件和控制台上显示，使用下面代码
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)
}


//一、日志输出格式重新封装
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//一、普通信息Info
func Info(f interface{}, v ...interface{}){
	opttime := time.Now().Format("2006/01/02 - 15:04:05")
	fmt.Fprintln(gin.DefaultWriter, "[INFO]"+opttime + " [I] " + formatLog(f, v...))
}

//二、错误信息Error
func Error(f interface{}, v ...interface{}){
	opttime := time.Now().Format("2006/01/02 - 15:04:05")
	fmt.Fprintln(gin.DefaultWriter, "[Erro]"+opttime + " [E] " + formatLog(f, v...))
}

//三、Debug信息
func Debug(f interface{}, v ...interface{}){
	opttime := time.Now().Format("2006/01/02 - 15:04:05")
	fmt.Fprintln(gin.DefaultWriter, "[Debg]"+opttime + " [D] " + formatLog(f, v...))
}

//日志信息格式化输出函数
func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}

//二、日志文件创建及控制
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
var logClient = log.New()
//日志文件控制器
func LogFile() {
	var logPath = "./"// 日志打印到指定的目录
	// 目录不存在则创建
	//if !util.PathExists(logPath) {
	//	os.MkdirAll(logPath, os.ModePerm)
	//}
	fileName := path.Join(logPath, "mikasa_ufs.log")
	//禁止logrus的输出
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err!= nil{
		fmt.Println("err", err)
	}
	// 设置日志输出的路径
	logClient.Out = src
	logClient.SetLevel(log.DebugLevel)
	logWriter, err := rotatelogs.New(
		"mikasa_ufs-%Y-%m-%d.log",
		rotatelogs.WithLinkName(fileName), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})   //json格式打印日志
	logClient.AddHook(lfHook)

	//logClient.SetFormatter(&log.TextFormatter{})              //文本格式打印日志
	logClient.Logf(log.InfoLevel,"------------------------------日志文件生成完毕-------------------------------")
	//logClient.Logf(log.ErrorLevel,"============================================================================================")
}