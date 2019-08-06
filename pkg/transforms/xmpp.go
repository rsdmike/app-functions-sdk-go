//
// Copyright (c) 2018
// Tencent
// IOTech
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"context"
	"crypto/tls"
	"strings"
	"time"

	contract "github.com/edgexfoundry/go-mod-core-contracts/models"

	"github.com/mattn/go-xmpp"
)

type XMPPSender struct {
	client  *xmpp.Client
	remote  string
	msgType string
	subject string
	thread  string
	other   []string
	stamp   time.Time
}

func NewXMPPSender(addr contract.Addressable) XMPPSender {
	protocol := strings.ToLower(addr.Protocol)

	if protocol == "tls" {
		xmpp.DefaultConfig = tls.Config{
			ServerName:         serverName(addr.Address),
			InsecureSkipVerify: false,
		}
	}

	options := xmpp.Options{
		Host:     addr.Address,
		User:     addr.User,
		Password: addr.Password,
		NoTLS:    protocol == "tls",
		Debug:    false,
		Session:  false,
	}

	xmppClient, err := options.NewClient()
	if err != nil {
		LoggingClient.Error(err.Error())
	}

	sender := &xmppSender{
		client: xmppClient,
	}

	return sender
}

func (sender *xmppSender) XMPPSend(edgexcontext *appcontext.Context, params ...interface{}) (bool, interface{}) {
	if len(params) < 1 {
		// We didn't receive a result
		return false, errors.New("No Data Received")
	}
	
	edgexcontext.LoggingClient.Debug("Setting output data")

	stringData := string(data)

	sender.client.Send(xmpp.Chat{
		Text:    stringData,
		Remote:  sender.remote,
		Subject: sender.subject,
		Thread:  sender.thread,
		Other:   sender.other,
		Stamp:   sender.stamp,
	})

	return true
}

func serverName(host string) string {
	return strings.Split(host, ":")[0]
}
