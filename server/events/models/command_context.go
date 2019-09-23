// Copyright 2017 HootSuite Media Inc.
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Modified hereafter by contributors to runatlantis/atlantis.

package models

import (
	"os/exec"

	"github.com/runatlantis/atlantis/server/logging"
)

// CommandContext represents the context of a command that should be executed
// for a pull request.
type CommandContext struct {
	// BaseRepo is the repository that the pull request will be merged into.
	BaseRepo Repo
	// HeadRepo is the repository that is getting merged into the BaseRepo.
	// If the pull request branch is from the same repository then HeadRepo will
	// be the same as BaseRepo.
	// See https://help.github.com/articles/about-pull-request-merges/.
	HeadRepo Repo
	Pull     PullRequest
	// User is the user that triggered this command.
	User User
	Log  *logging.SimpleLogger
	// PullMergeable is true if Pull is able to be merged. This is available in
	// the CommandContext because we want to collect this information before we
	// set our own build statuses which can affect mergeability if users have
	// required the Atlantis status to be successful prior to merging.
	PullMergeable bool

	// Cancelled signifies that the command should stop running
	Cancelled bool
	// LastCmd is the command to cancel
	LastCmd *exec.Cmd
}

// Cancel the running command
func (ctx *CommandContext) Cancel() {
	ctx.Cancelled = true
	if ctx.LastCmd != nil && !ctx.LastCmd.ProcessState.Exited() {
		ctx.LastCmd.Process.Kill()
	}
}