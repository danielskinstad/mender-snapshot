// Copyright 2024 Northern.tech AS
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//	    http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.
package cli

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/creack/pty"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestDumpNoRedirect(t *testing.T) {
	// Create a pseudo terminal to capture stdout
	ptmx, tty, err := pty.Open()
	if err != nil {
		log.Error(err)
	}

	defer ptmx.Close()
	defer tty.Close()

	os.Stdout = tty

	args := []string{"snapshot", "dump"}
	err = SetupCLI(args)
	if err == nil {
		log.Error("Expected error")
	}
	assert.ErrorContains(t, err, "Refusing to write to terminal")
}

func TestVersion(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdout = w

	args := []string{"snapshot", "--version"}
	err := SetupCLI(args)
	if err != nil {
		log.Error(err)
	}

	w.Close()
	out, _ := io.ReadAll(r)

	read_line := strings.TrimSpace(string(out))
	assert.Equal(t, ShowVersion(), read_line)
}
