/*
Copyright © 2021 kubelt

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package util

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func CheckIfKtrlIsInstalledAndRunning(cmd *cobra.Command, args []string) error {
	// check if the ktrl executable exists
	if _, err := exec.LookPath("ktrl"); err != nil {
		return fmt.Errorf("ktrl is not installed.")
	}
	// check if ktrl is running
	if !checkKtrlProcess() {
		return fmt.Errorf("ktrl is not running")
	}
	return nil
}
