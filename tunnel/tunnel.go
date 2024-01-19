package tunnel

// 这里主要是创建tun接口，我们直接使用tailscale的库即可，它相对比较成熟.
// 这里的tunnel与clash中的tunnel并不相同，那里实际上类似于relay???


import (
	"fmt"
	"log"
	"runtime"

	"tailscale.com/net/tstun"
	"tailscale.com/wgengine"
	"tailscale.com/wgengine/netstack"
)

func defaultTunName() string {
	switch runtime.GOOS {
	case "windows":
		return "clash-tun"
	case "darwin":
		// "utun" is recognized by wireguard-go/tun/tun_darwin.go
		// as a magic value that uses/creates any free number.
		return "utun"
	case "linux":
		return "clash-tun"

	}

	return "clash-tun"
}

func InitTunnel() {
	_, devName, err := tstun.New(log.Printf, defaultTunName())
	if err != nil {
		fmt.Println(fmt.Sprintf("init clash-tun failed, err %v"), err)
	}
	fmt.Println("tun create ok with name: " + devName)

	// 或许我们可以在这里初始化ts的用户态netstack，从而最终于clash
	// 的tunnel（clash处理用户连接的核心逻辑，即需要被代理的流量）
	// 关联在一起。
	//
	// 对于分流流量，clash通过embed dns server + fake ip引流。

	// 2401:
	// TODO: 可以考虑用 wgengine.NewUserspaceEngine来实现。
	// 请参考tailscale tsnet逻辑实现。
}
