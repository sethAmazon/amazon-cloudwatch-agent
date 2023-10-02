// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

//go:build !windows
// +build !windows

package security

import (
	"github.com/aws/amazon-cloudwatch-agent/translator/translate/metrics/util"
)

func CheckNvidiaSMIBinaryRights() error {
	if err := CheckFileRights(util.Default_Unix_Smi_Path); err != nil {
		return err
	}
	return nil
}
