// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package util

import (
	"sync"
)

var logBlockerSingleton *LogBlocker

type LogBlocker struct {
	maxLogBufferSize int64
	logsBufferSize int64
	logsBufferSizeMutex sync.Mutex
}

func GetLogBlocker() *LogBlocker {
	if logBlockerSingleton == nil {
		logBlockerSingleton = &LogBlocker{
			maxLogBufferSize: int64(-1),
			logsBufferSize: 0,
			logsBufferSizeMutex: sync.Mutex{},
		}
	}
	return logBlockerSingleton
}

func (l *LogBlocker) Add(v int64) {
	l.logsBufferSizeMutex.Lock()
	defer l.logsBufferSizeMutex.Unlock()
	l.logsBufferSize = l.logsBufferSize + v
}

func (l *LogBlocker) Subtract(v int64) {
	l.logsBufferSizeMutex.Lock()
	defer l.logsBufferSizeMutex.Unlock()
	l.logsBufferSize = l.logsBufferSize - v
}

// Reset this is test only code
// do not use for application
func (l *LogBlocker) Reset() {
	l.logsBufferSizeMutex.Lock()
	defer l.logsBufferSizeMutex.Unlock()
	l.logsBufferSize = 0
}

func (l *LogBlocker) SetMaxLogBuffer(v int64) {
	l.maxLogBufferSize = v
}

func (l *LogBlocker) Block() (block bool, bufferSize int64, maxBufferSize int64) {
	l.logsBufferSizeMutex.Lock()
	defer l.logsBufferSizeMutex.Unlock()
	block = l.maxLogBufferSize != -1 && l.logsBufferSize >= l.maxLogBufferSize
	bufferSize = l.logsBufferSize
	maxBufferSize = l.maxLogBufferSize
	return block, bufferSize, maxBufferSize
}
