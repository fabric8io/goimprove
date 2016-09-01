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
package cmds

import (
	"strconv"
	"strings"

	"github.com/fabric8io/goimprove/util"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat = "2006-01-02 15:04 MST"
)

// NewCmdGitHubDownloads retrives the number of downloads of a GitHub project release
func NewCmdGitHubDownloads() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gh-downloads",
		Short: "retrives the number of downloads of a GitHub project release",
		Long:  `retrives the number of downloads of a GitHub project release`,
		PreRun: func(cmd *cobra.Command, args []string) {
			showBanner()
		},
		Run: func(cmd *cobra.Command, args []string) {
			repo := cmd.Flags().Lookup("repository").Value.String()
			util.Infof("Getting release downloads numbers for GitHub project %s\n", repo)

			org := strings.Split(repo, "/")[0]
			project := strings.Split(repo, "/")[1]

			client := github.NewClient(nil)

			opt := &github.ListOptions{
				Page:    0,
				PerPage: 100,
			}
			// get all pages of results
			var allReleases []*github.RepositoryRelease
			for {
				releases, resp, err := client.Repositories.ListReleases(org, project, opt)
				if err != nil {
					util.Errorf("Unable to list repositories by org %s %v", "fabric8io", err)
				}
				allReleases = append(allReleases, releases...)
				if resp.NextPage == 0 {
					break
				}
				opt.Page = resp.NextPage
			}
			grandTotal := 0
			var previousReleaseTimeStamp github.Timestamp
			for v := range allReleases {
				release := allReleases[v]
				tag := *release.Name
				releaseDate := *release.PublishedAt
				duration := previousReleaseTimeStamp.Time.Sub(releaseDate.Time)
				totalDownloadCount := 0
				if tag != "" {
					for w := range release.Assets {
						asset := release.Assets[w]
						totalDownloadCount = totalDownloadCount + *asset.DownloadCount
					}

					d := duration.Hours() / 24
					days := strconv.FormatFloat(d, 'f', 6, 64)

					util.Infof("Tag %s published had %v downloads and was available for %s days\n", tag, totalDownloadCount, days)

				}
				previousReleaseTimeStamp = releaseDate
				grandTotal = grandTotal + totalDownloadCount
			}
			util.Infof("\nGrand total of %v downloads\n", grandTotal)
		},
	}
	cmd.PersistentFlags().StringP("repository", "r", "", "the GitHub repository to get the release download numbers e.g. fabric8io/gofabric8")

	return cmd
}
