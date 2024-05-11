package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shankusu2017/constant"
	"github.com/shankusu2017/proto_pb/go/proto"
	pb "google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type NodeT struct {
	IP   string    `json:"ip,omitempty"`
	Role string    `json:"ip,omitempty"`
	Ver  string    `json:"ip,omitempty"`
	Ping time.Time `json:"ip,omitempty"` // 最后一次 ping 的时间
}

type nodeMgrT struct {
	nodeMap map[string]*NodeT // ip->node
	dataMtx sync.Mutex
}

var (
	nodeMgr *nodeMgrT
)

func InitNodeMgr() {
	nodeMgr = &nodeMgrT{}
	nodeMgr.nodeMap = make(map[string]*NodeT)
	go nodeMgr.loopScan()
}

func (mgr *nodeMgrT) addNode(node *NodeT) {
	mgr.dataMtx.Lock()
	defer mgr.dataMtx.Unlock()

	mgr.nodeMap[node.IP] = node
}

func (mgr *nodeMgrT) updateNode(ip string) {
	mgr.dataMtx.Lock()
	defer mgr.dataMtx.Unlock()

	node, ok := mgr.nodeMap[ip]
	if !ok {
		log.Printf("ERROR 0x5a43bf8d node is nil, cli.ip: %s", ip)
		return
	}

	node.Ping = time.Now()
}

func (mgr *nodeMgrT) getNode(role string) []string {
	mgr.dataMtx.Lock()
	defer mgr.dataMtx.Unlock()

	ipList := make([]string, 0)
	for _, node := range mgr.nodeMap {
		if node.Role == role {
			ipList = append(ipList, node.IP)
		}
	}

	return ipList
}

func GetNodeAll() []NodeT {
	nodeMgr.dataMtx.Lock()
	defer nodeMgr.dataMtx.Unlock()

	lst := make([]NodeT, 0)

	for _, node := range nodeMgr.nodeMap {
		lst = append(lst, *node)
	}

	return lst
}

func (mgr *nodeMgrT) loopScan() {
	for {
		time.Sleep(time.Second * 3)
		now := time.Now()

		mgr.dataMtx.Lock()
		for key, node := range mgr.nodeMap {
			if node.Ping.Before(now.Add(time.Second * -7)) {
				log.Printf("LOG 0x71bec216 node(%s) too old, delete it", key)
				delete(mgr.nodeMap, key)
			}
		}
		mgr.dataMtx.Unlock()
	}
}

func nodeEvent(c *gin.Context, msg *proto.MsgEventPost) {
	if msg.GetEvent() == proto.Event_STARTED {
		node := &NodeT{}
		node.IP = c.RemoteIP()
		node.Ping = time.Now()
		mNode := msg.GetNode()
		if mNode == nil {
			log.Printf("ERROR 0x12912233 node is nil")
			return
		}
		node.Role = mNode.GetRole()
		if node.Role != constant.ROLE_REPEATER && node.Role != constant.ROLE_PAC {
			log.Printf("ERROR 0x5454cb42 role invalid:%s", node.Role)
			return
		}
		node.Ver = mNode.GetVer()

		nodeMgr.addNode(node)
	} else if msg.GetEvent() == proto.Event_KEEPALIVE {
		nodeMgr.updateNode(c.RemoteIP())
	}
}

func repeaterServerGet(c *gin.Context) {
	ip := c.RemoteIP()

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("0x66dda745 read request body fail:%s, ip:%s", err.Error(), ip)
		return
	}

	var req proto.MsgRepeaterServerInfoReq
	err = pb.Unmarshal(bodyBytes, &req)
	if err != nil {
		log.Printf("0x4842fc43 Invalid request body(%v), ip:%s", bodyBytes, ip)
		return
	}

	machine := req.GetMachine()
	if machine == nil {
		log.Printf("0x630d0ded client(ip:%s) req repeater server list, machine.id is nil", ip)
		return
	} else {
		log.Printf("client(ip:%s, id:%s) req repeater server list", ip, machine.GetUUID())
	}

	var rsp proto.MsgRepeaterServerInfoRsp

	ipLst := nodeMgr.getNode(constant.ROLE_REPEATER)
	for _, iP := range ipLst {
		node := &proto.RepeaterServerNode{IPv4: iP}
		rsp.Servers = append(rsp.Servers, node)
	}

	c.ProtoBuf(http.StatusOK, &rsp)
}
