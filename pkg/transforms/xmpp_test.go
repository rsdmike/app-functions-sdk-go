//
// Copyright (c) 2018 Tencent
//
// SPDX-License-Identifier: Apache-2.0
//

package transforms

import (
	"flag"
	"strings"
	"testing"
	"time"

	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/mattn/go-xmpp"
)

const (
	TestMessage = "hello world"
)

var statusMessage = flag.String("status-msg", "I for one welcome our new codebot overlords.", "status message")

func getServerName(host string) string {
	return strings.Split(host, ":")[0]
}

func TestXmppSend(t *testing.T) {

	var err error

	addr := contract.Addressable{
		Address:  "talk.google.com:5222",
		User:     "", //your gmail account, eg: xxx@gmail.com
		Password: "", //your Gmail password
		Protocol: "", //set to TLS if desired
	}

	//talk, err = options.NewClient()
	sender, err := NewXMPPSender(addr)
	if err != nil {
		t.Fatal(err.Error())
	}
	go func() {
		for {
			chat, err := sender.client.Recv()
			if err != nil {
				t.Fatal(err)
			}
			switch v := chat.(type) {
			case xmpp.Chat:
				if v.Text != TestMessage {
					t.Errorf("Expected received message : %s, actual received message : %s", TestMessage, v.Text)
				}
			case xmpp.Presence:
			}
		}
	}()

	sender.XMPPSend(context, TestMessage)

	//waiting for receiving go routine
	time.Sleep(5 * time.Second)

}
