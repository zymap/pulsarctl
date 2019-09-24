// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package info

import (
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLastMessageIdArgsError(t *testing.T) {
	args := []string{"last-message-id"}
	_, _, nameErr, _ := TestTopicCommands(GetLastMessageIdCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestGetLastMessageIdTopicNotExistError(t *testing.T) {
	args := []string{"last-message-id", "not-existent-topic"}
	_, execErr, _, _ := TestTopicCommands(GetLastMessageIdCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestGetLastMessageIdNotAllowedError(t *testing.T) {
	args := []string{"create",
		"non-persistent://public/default/last-message-id-non-persistent-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"last-message-id", "non-persistent://public/default/last-message-id-non-persistent-topic"}
	_, execErr, _, _ = TestTopicCommands(GetLastMessageIdCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t,
		"code: 405 reason: GetLastMessageId on a non-persistent topic is not allowed",
		execErr.Error())
}