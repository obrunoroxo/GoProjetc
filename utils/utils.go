package utils

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"time"

	"github.com/fatih/color"
)

func Sleep(wait int) {
	duration := time.Duration(wait) * time.Second
	time.Sleep(duration)
}

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls", "clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Close(value bool) {
	if !value {
		os.Exit(1)
	} else {
		fmt.Println("See ya! :)")
		os.Exit(0)
	}
}

func Type(value interface{}) interface{} {
	return reflect.TypeOf(value)
}

func GreenColor(value string) string {
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	return green(value)
}

func TimeNow() string {
	timeNow := time.Now()                       // Getting a current date/time
	timeZone := time.FixedZone("BRT", -3*60*60) // Definning a Brazil time zone
	timeBrazil := timeNow.In(timeZone)          // Transform a current date/time in a Brazil time zone
	layout := "2006-01-02 15:04:05 MST"         // Custom layout example
	formatHour := timeBrazil.Format(layout)     // Formatting the date/time in the layout passed

	return formatHour
}
