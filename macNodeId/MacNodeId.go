package macNodeId

import (
	"net"
	"bytes"
	"crypto/sha1"
	"errors"
)

// node id length is 6 bytes
const nodeIdLength = 6

type MacNodeId struct {
}

func (n *MacNodeId) Id() ([]byte,error) {

	macNodeId := make([]byte,nodeIdLength)
	err := assignMacNodeId(&macNodeId)
	
	return macNodeId,err 
}


func assignMacNodeId(ptrId *[]byte) error {
	inets,err := net.Interfaces()

	if err != nil {
		return errors.New("Unable to obtain mac address: "+err.Error())
	}
        macs := make([][]byte,len(inets))
	for i:=0;i<len(inets);i++ {
		macs[i] = inets[i].HardwareAddr
	}
	m := bytes.Join(macs,nil)
	s :=  sha1.Sum(m)
	copy((*ptrId),s[0:nodeIdLength])
	return nil
}

