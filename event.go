package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shankusu2017/proto_pb/go/proto"
	pb "google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func eventPost(c *gin.Context) {
	ip := c.RemoteIP()

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("0x2dd56b9d read request body fail:%s, ip:%s", err.Error(), ip)
		return
	}

	var msg proto.MsgEventPost
	err = pb.Unmarshal(bodyBytes, &msg)
	if err != nil {
		log.Printf("0x4842fc43 Invalid request body(%v), ip:%s", bodyBytes, ip)
		return
	}

	{
		jBuf, _ := json.Marshal(msg)
		log.Printf("0x09d8bb7d recv a event:[%s], cli:%s", string(jBuf), ip)
	}

	machine := msg.GetMachine()
	if machine == nil {
		log.Printf("0x630d0ded client(ip:%s) req repeater server list, machine.id is nil", ip)
		return
	}
	log.Printf("client(ip:%s, id:%s) req repeater server list", ip, machine.GetUUID())

	event := msg.GetEvent()
	if event == proto.Event_STARTED ||
		event == proto.Event_KEEPALIVE {
		nodeEvent(c, &msg)
	} else {
		log.Printf("0x1ae4262b recv invalid event(%s)", event)
		return
	}
}
