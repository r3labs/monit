/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// Vpc : ...
type Vpc struct {
}

// Handle : ...
func (n *Vpc) Handle(subject string, c component, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "vpc.create":
		lines = n.getSingleDetail(c, "VPC created")
	case "vpc.update":
		lines = n.getSingleDetail(c, "VPC udpated")
	case "vpc.delete":
		lines = n.getSingleDetail(c, "VPC deleted")
	case "vpcs.find":
		lines = n.getSingleDetail(c, "VPC Found")
	}
	return lines
}

func (n *Vpc) getSingleDetail(c component, prefix string) (lines []Message) {
	id, _ := c["vpc_id"].(string)
	if prefix != "" {
		id = prefix + " " + id
	}
	subnet, _ := c["subnet"].(string)
	status, _ := c["_state"].(string)
	lines = append(lines, Message{Body: " - " + id, Level: ""})
	lines = append(lines, Message{Body: "   Subnet    : " + subnet, Level: ""})
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if status == "errored" {
		err, _ := c["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}
	return lines
}
