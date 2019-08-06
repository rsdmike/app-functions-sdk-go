//
// Copyright (c) 2018
// Tencent
// IOTech
//
// SPDX-License-Identifier: Apache-2.0
//

package transforms

import (
	"crypto/tls"
	"errors"
	"strings"
	"time"

	"github.com/edgexfoundry/app-functions-sdk-go/appcontext"
	"github.com/edgexfoundry/app-functions-sdk-go/pkg/util"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"

	"github.com/mattn/go-xmpp"
)

// XMPPSender ...
type XMPPSender struct {
	client  *xmpp.Client
	remote  string
	msgType string
	subject string
	thread  string
	other   []string
	stamp   time.Time
}

// NewXMPPSender ...
func NewXMPPSender(addr contract.Addressable) *XMPPSender {
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
		// LoggingClient.Error(err.Error())
	}

	sender := &XMPPSender{
		client: xmppClient,
	}

	return sender
}

// XMPPSend ...
func (sender *XMPPSender) XMPPSend(edgexcontext *appcontext.Context, params ...interface{}) (bool, interface{}) {
	if len(params) < 1 {
		// We didn't receive a result
		return false, errors.New("No Data Received")
	}

	edgexcontext.LoggingClient.Debug("Setting output data")

	data, err := util.CoerceType(params[0])
	if err != nil {
		return false, err
	}
	stringData := string(data)

	sender.client.Send(xmpp.Chat{
		Text:    stringData,
		Remote:  sender.remote,
		Subject: sender.subject,
		Thread:  sender.thread,
		Other:   sender.other,
		Stamp:   sender.stamp,
	})

	return true, nil
}

func serverName(host string) string {
	return strings.Split(host, ":")[0]
}
