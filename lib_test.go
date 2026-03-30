// Copyright (c) 2026 Jane Doe
// Licensed under the MIT License. See LICENSE file in the project root.
package libuflog_test

import (
	"testing"

	"github.com/zelferion/libuflog.go"
)

func TestPrints(t *testing.T) {
	logger := libuflog.NewDefaultLogger()
	logger.Debug("zxc")
	logger.Info("zxc")
	logger.Warn("zxc")
	logger.Error("zxc")
	logger.Fatal("zxc")
}
