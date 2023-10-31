package main

import (
	"fmt"

	"github.com/fortitudepub/clash-gioui/ui"
	"github.com/fortitudepub/clash-gioui/clash"
	"github.com/fortitudepub/clash-gioui/tunnel"
)


func main() {
	clash.InitClash()
	tunnel.InitTunnel()

	fmt.Println("staring ui main...")
	ui.UIMain()
}
