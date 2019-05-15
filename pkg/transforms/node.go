
//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package transforms

import (
	"os/exec"
	"os"
	
	"errors"
	

	"github.com/edgexfoundry/app-functions-sdk-go/appcontext"
)

// NodeExecutor ...
type NodeExecutor struct {
	FilePath      string
}

// NodeExecutor ...
func (sender NodeExecutor) NodeExecutor(edgexcontext *appcontext.Context, params ...interface{}) (bool, interface{}) {
	if len(params) < 1 {
		// We didn't receive a result
		return false, errors.New("No Data Received")
	}
	
	edgexcontext.LoggingClient.Trace("EVENT PROCESSED BY NODE")

	env := os.Environ()
	funcPath := sender.FilePath
	cmd := exec.Command("node", funcPath, params[0].(string))
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Start()
	if err != nil {
		edgexcontext.LoggingClient.Error(err.Error())
	}
	//return cmd.Process, cmd.Start()

	return true, nil
}
