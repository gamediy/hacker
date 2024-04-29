package register

import (
	"attack/utility/xetcd"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	clientv3 "go.etcd.io/etcd/client/v3"
	"net"
)

type node struct {
	Ip string
}

// 客户端注册
func KeepOnline() error {

	var (
		err           error
		keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
		keepAliveResp *clientv3.LeaseKeepAliveResponse
	)
	ipv4, err := getLocalIP()
	if err != nil {
		return err
	}
	cancelCtx, cancelFunc := context.WithCancel(context.TODO())
	regKey := "/worker/" + ipv4
	// 创建租约
	leaseGrantResp, err := xetcd.Client.Grant(cancelCtx, 10)
	if err != nil {
		cancelFunc()
		return err
	}
	if keepAliveChan, err = xetcd.Client.KeepAlive(cancelCtx, leaseGrantResp.ID); err != nil {
		cancelFunc()
		return err
	}
	// 注册到etcd
	if _, err = xetcd.Client.Put(cancelCtx, regKey, ipv4, clientv3.WithLease(leaseGrantResp.ID)); err != nil {
		cancelFunc()
		return err
	}

	fmt.Println("register success")

	get, err := xetcd.Client.Get(cancelCtx, "/worker/", clientv3.WithPrefix())
	if err != nil {
		g.Log().Error(cancelCtx, err)
		return err
	}
	g.Dump(get.Kvs)
	// 处理续租应答
	for {
		select {
		case keepAliveResp = <-keepAliveChan:
			if keepAliveResp == nil { // 续租失败
				cancelFunc()
			}
		}
	}

	return nil
}

// 获取本机网卡IP
func getLocalIP() (ipv4 string, err error) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
	)
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}

	err = fmt.Errorf("未找倒IP")
	return
}
