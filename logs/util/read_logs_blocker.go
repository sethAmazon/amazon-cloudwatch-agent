// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package util

import (
	"log"
	"sync"
)

type LogBlocker struct {
	maxLogBufferSize int64
	logsBufferSize int64
	logsBufferSizeMutex sync.Mutex
}

// DefaultLogBlocker will not block logs
func DefaultLogBlocker() *LogBlocker {
	return &LogBlocker{
		maxLogBufferSize:    -1,
		logsBufferSize:      0,
		logsBufferSizeMutex: sync.Mutex{},
	}
}

func NewLogBlocker(maxLogBufferSize int64) *LogBlocker {
	return &LogBlocker{
		maxLogBufferSize:    maxLogBufferSize,
		logsBufferSize:      0,
		logsBufferSizeMutex: sync.Mutex{},
	}
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

func (l *LogBlocker) SetMaxLogBuffer(v int64) {
	l.maxLogBufferSize = v
}

func (l *LogBlocker) Block() (block bool, bufferSize int64, maxBufferSize int64) {
	l.logsBufferSizeMutex.Lock()
	defer l.logsBufferSizeMutex.Unlock()
	block = l.maxLogBufferSize != -1 && l.logsBufferSize >= l.maxLogBufferSize
	bufferSize = l.logsBufferSize
	maxBufferSize = l.maxLogBufferSize
	log.Printf("D! [log blocker] log blocker block %t, bs %d, mbs %d", block, bufferSize, maxBufferSize)
	return block, bufferSize, maxBufferSize
}
