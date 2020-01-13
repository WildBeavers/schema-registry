package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getConfigCmd = &cobra.Command{
	Use:   "get-config [subject]",
	Short: "retrieves global or subject specific configuration",
	Long: `Configuration can be requested for all or a specific subject. 
When "compatibility-type" is not defined for a specific subject,
then it's using global compatibility type. To check global setting
just call "get-config" without arguments.
Compatibility types in Schema-Registry may be: "NONE", "BACKWARD",
"BACKWARD_TRANSITIVE", "FORWARD", "FORWARD_TRANSITIVE", "FULL" and
"FULL_TRANSITIVE".
Please consider official documentation for more details.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch {
		case len(args) > 1:
			return fmt.Errorf("only one subject allowed")
		case len(args) == 0:
			if err := getConfig(""); err != nil {
				return err
			}
		case len(args) == 1:
			if err := getConfig(args[0]); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(getConfigCmd)
}
