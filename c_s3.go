/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// S3Bucket : ...
type S3Bucket struct {
}

// Handle : ...
func (n *S3Bucket) Handle(subject string, c component, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "s3.create":
		lines = n.getSingleDetail(c, "S3 bucket created")
	case "s3.update":
		lines = n.getSingleDetail(c, "S3 bucket updated")
	case "s3.delete":
		lines = n.getSingleDetail(c, "S3 bucket deleted")
	case "s3s.find":
		lines = n.getSingleDetail(c, "S3 bucket imported")
	}
	return lines
}

func (n *S3Bucket) getSingleDetail(c component, prefix string) (lines []Message) {
	name, _ := c["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	acl, _ := c["acl"].(string)
	if acl == "" {
		acl = "by grantees"
	}
	status, _ := c["_state"].(string)
	level := "INFO"
	if status == "errored" {
		level = "ERROR"
	}
	lines = append(lines, Message{Body: " " + name, Level: level})
	lines = append(lines, Message{Body: "   ACL       : " + acl, Level: ""})
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if status == "errored" {
		err, _ := c["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: ""})
	}
	return lines
}
