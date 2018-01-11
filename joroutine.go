// Copyright 2018 Miguel Rivero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package joroutine contains the Go implementation of the lib.
package joroutine

// BackgroundTask interface should be implemented in Java
type BackgroundTask interface {
	Run()
	LogDebug(message string)
}

// RunOnGoRoutine is a static method that launches a BackgroundTask into a GoRoutine
func RunOnGoRoutine(taskToRun BackgroundTask) {

	go func(methodToRun BackgroundTask) {
		printLog := "Joroutine: Starting a new BackgroundTask into a goroutine"
		methodToRun.LogDebug(printLog)
		methodToRun.Run()
	}(taskToRun)

}
