package utils

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"time"
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
