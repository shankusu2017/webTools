package main

import (
	"github.com/gin-gonic/gin"
	proto "github.com/shankusu2017/proto_pb/go/proto"
	pb "google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func repeaterServerListGet(c *gin.Context) {
	ip := c.RemoteIP()

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("0x66dda745 read request body fail:%s, ip:%s", err.Error(), ip)
		return
	}

	var req proto.RepeaterServerInfoReq
	err = pb.Unmarshal(bodyBytes, &req)
	if err != nil {
		log.Printf("0x4842fc43 Invalid request body(%v), ip:%s", bodyBytes, ip)
		return
	}

	machine := req.GetId()
	if machine == nil {
		log.Printf("0x630d0dedclient(ip:%s) req repeater server list, machine.id is nil", ip)
		return
	} else {
		log.Printf("client(ip:%s, id:%s) req repeater server list", ip, machine.GetUUID())
	}

	// TODO
	var rsp proto.RepeaterServerInfoRsp
	c.ProtoBuf(501, &rsp)

	// 解析 域名
	// ping 一下看看是否存活
	// 存档上述信息
	// 返回存活的server
}
