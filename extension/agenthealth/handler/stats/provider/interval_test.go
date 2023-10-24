// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package provider

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
)

func TestIntervalStats(t *testing.T) {
	s := newIntervalStats(time.Millisecond)
	s.stats.ThreadCount = aws.Int32(2)
	got := s.Stats("")
	assert.NotNil(t, got.ThreadCount)
	got = s.Stats("")
	assert.Nil(t, got.ThreadCount)
	time.Sleep(2 * time.Millisecond)
	got = s.Stats("")
	assert.NotNil(t, got.ThreadCount)
	got = s.Stats("")
	assert.Nil(t, got.ThreadCount)
}
