// Copyright 2018 charro (Miguel Rivero). All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package joroutines contains the Go implementation of the lib.
package joroutines

// BackgroundTask interface should be implemented in Java
type BackgroundTask interface {
	Run()
	LogDebug(message string)
}

// RunIntoGoroutine is a static method that launches a BackgroundTask into a GoRoutine
func RunIntoGoroutine(taskToRun BackgroundTask) {

	go func(methodToRun BackgroundTask) {
		printLog := "Joroutine: Starting a new BackgroundTask into a goroutine"
		methodToRun.LogDebug(printLog)
		methodToRun.Run()
	}(taskToRun)

}
