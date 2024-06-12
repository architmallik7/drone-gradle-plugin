// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"os/exec"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	// Execute Gradle build command
	cmd := exec.Command("gradle", "build")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error executing Gradle build command:", err)
		return fmt.Errorf("error executing gradle build command")
	}

	fmt.Println(string(output))
	return nil
}
