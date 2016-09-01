/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	commands "github.com/fabric8io/goimprove/cmds"
	"github.com/spf13/cobra"
)

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func main() {
	cmds := &cobra.Command{
		Use:   "goimprove",
		Short: "goimprove is used to gather stats and metrics",
		Long: `goimprove is used to gather stats and metrics of various types and expose via rest to be scraped by a metrics tool
								Find more information at http://fabric8.io.`,
		Run: runHelp,
	}

	cmds.PersistentFlags().BoolP("yes", "y", false, "assume yes")

	cmds.AddCommand(commands.NewCmdGitHubDownloads())
	cmds.AddCommand(commands.NewCmdVersion())

	cmds.Execute()
}
