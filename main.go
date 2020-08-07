package main

import (
	"github.com/CodyGuo/logtest/cmd"

	"github.com/CodyGuo/glog"
)

func main() {
	if err := cmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}
