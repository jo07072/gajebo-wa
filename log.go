package main

import "github.com/fatih/color"

func Info(format string, a ...interface{}) {
	color.Blue("[INFO] : "+format, a...)
}

func Success(format string, a ...interface{}) {
	color.Green("[ OK ] : "+format, a...)
}

func Warn(format string, a ...interface{}) {
	color.Yellow("[WARN] : "+format, a...)
}

func Error(format string, a ...interface{}) {
	color.Red("[EROR] : "+format, a...)
}
