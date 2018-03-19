// Copyright © 2018 Matthias Diester
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/HeavyWombat/dyff/core"
	"github.com/spf13/cobra"
)

// betweenCmd represents the between command
var betweenCmd = &cobra.Command{
	Use:   "between",
	Short: "Compares differences between documents",
	Long: `
Compares differences between documents and displays the delta. Supported
document types are: YAML (http://yaml.org/) and JSON (http://json.org/).

`,
	Args:    cobra.ExactArgs(2),
	Aliases: []string{"bw"},
	Run: func(cmd *cobra.Command, args []string) {
		fromLocation := args[0]
		toLocation := args[1]

		from, err := core.LoadYAMLFromLocation(fromLocation)
		if err != nil {
			panic(err)
		}

		to, err := core.LoadYAMLFromLocation(toLocation)
		if err != nil {
			panic(err)
		}

		diffs := core.CompareDocuments(from, to)

		fmt.Printf("Difference between %s and %s ...\n", core.Bold(fromLocation), core.Bold(toLocation))
		for i, diff := range diffs {
			fmt.Printf("%s\n%v\n\n", core.Bold(fmt.Sprintf("diff #%d:", i)), diff)
		}
	},
}

func init() {
	rootCmd.AddCommand(betweenCmd)
}