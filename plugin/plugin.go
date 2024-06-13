// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	// Goals defines the gradle targets/goals to execute.
	Goals string `envconfig:"PLUGIN_GOALS"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	// Split the goals into individual targets
	goals := strings.Fields(args.Goals)

	// Run `gradle` command with specified goals
	gradleCmd := exec.Command("gradle", goals...)
	gradleOutput, gradleErr := gradleCmd.CombinedOutput()
	if gradleErr != nil {
		fmt.Println("Error running 'gradle "+args.Goals+"':", gradleErr)
		fmt.Println(string(gradleOutput))
		return fmt.Errorf("error running 'gradle %s': %w", args.Goals, gradleErr)
	}
	fmt.Println("Output of 'gradle "+args.Goals+"':", string(gradleOutput))

	return nil
}
