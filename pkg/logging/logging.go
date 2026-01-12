// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the Licenses
package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	infolog  *log.Logger
	errorlog *log.Logger
	fatallog *log.Logger
	// Exit is a function that exits the program. It is overridden in tests.
	Exit = os.Exit
)

func init() {
	// keep simple writers (we'll format timestamps in our wrapper functions so
	// the format is consistent and in RFC3339Nano UTC)
	infolog = log.New(os.Stdout, "", 0)
	errorlog = log.New(os.Stderr, "", 0)
	fatallog = log.New(os.Stderr, "", 0)
}

// timestamp returns current time in RFC3339Nano UTC
func timestamp() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}

// Info prints info to stdout with timestamp and level
func Info(f string, a ...any) {
	msg := fmt.Sprintf(f, a...)
	coloredTimestamp := "\x1b[35m" + timestamp() + "\x1b[0m"
	infolog.Printf("%s: %s", coloredTimestamp, msg)
}

// Error prints info to stderr but does not end the program
func Error(f string, a ...any) {
	msg := fmt.Sprintf(f, a...)
	coloredTimestamp := "\x1b[35m" + timestamp() + "\x1b[0m"
	errorlog.Printf("%s ERROR: %s", coloredTimestamp, msg)
}

// Fatal prints info to stderr and ends the program
func Fatal(f string, a ...any) {
	msg := fmt.Sprintf(f, a...)
	coloredTimestamp := "\x1b[35m" + timestamp() + "\x1b[0m"
	fatallog.Printf("%s FATAL: %s", coloredTimestamp, msg)
	Exit(1)
}
