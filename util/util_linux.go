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
	"os/exec"
	"regexp"
)

func checkKtrlProcess() bool {
	// cmd := exec.Command("pgrep", "-f", "ktrl")
	// result, err := cmd.Output()
	// if err != nil {
	// 	return false
	// }
	// return len(result) != 0
	output, err := exec.Command("systemctl", "status", "ktrl.service").Output()
	if err == nil {
		if matched, err := regexp.MatchString("Active: active", string(output)); err == nil && matched {
			return true
		}
	}

	return false
}
