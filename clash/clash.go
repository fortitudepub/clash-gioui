package clash

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/Dreamacro/clash/adapter/outbound"
	"github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/listener/socks"
)

func InitClash() {
	// 以下从clash代码中直接拷贝得出。

	// 我们的架构方案是以创建tun设备后，通过路由引流到tun设备。
	// 然由利用用户态协议栈比如tailscale netstack（或直接用底层gVISOR)重组出tcp+udp socket，并模拟进行回复，
	// 同时，利用从中提取了的目标ip+port，利用clash的或规则引擎/outbound
	// 对目标ip+port进行代理。

	// clash的初始化应当是
	// 1. 加载配置（初期从固定目录中读取）
	// 2. 加载配置
	//   2.1 加载rule
	//   2.2 如果有inbound，则进行listen（此处我们并不需要）
	//   2.3 初始化outbound


	in := make(chan constant.ConnContext, 100)
	defer close(in)


	l, err := socks.New("127.0.0.1:10000", in)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	println("listen at:", l.Address())

	direct := outbound.NewDirect()

	// just a test...
	go func () {
		for c := range in {
			conn := c
			metadata := conn.Metadata()
			fmt.Printf("request incoming from %s to %s\n", metadata.SourceAddress(), metadata.RemoteAddress())
			go func () {
				remote, err := direct.DialContext(context.Background(), metadata)
				if err != nil {
					fmt.Printf("dial error: %s\n", err.Error())
					return
				}
				relay(remote, conn.Conn())
			}()
		}
	}()
}

func relay(l, r net.Conn) {
	go io.Copy(l, r)
	io.Copy(r, l)
}
