package cmd

import (
	"fmt"

	schemaregistry "github.com/coursehero/schema-registry"
	"github.com/spf13/cobra"
)

// constants for parameter names
const (
	Compatibility = "compatibility-type"
)

var (
	compatibility string
)

var setConfigCmd = &cobra.Command{
	Use:   "set-config [subject]",
	Short: "set global or subject specific configuration",
	Long: `Configuration currently contains only the compatibility type.
It can be set for all or a specific subject. 
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
			if err := setConfig("", compatibility); err != nil {
				return err
			}
		case len(args) == 1:
			if err := setConfig(args[0], compatibility); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(setConfigCmd)
	setConfigCmd.PersistentFlags().StringVar(&compatibility, Compatibility, "", "compatibility level to set")
	setConfigCmd.MarkFlagRequired(Compatibility)
}

func setConfig(subj string, compatibilityType string) error {
	client := assertClient()
	return client.SetConfig(subj, schemaregistry.ConfigPut{CompatibilityType: schemaregistry.CompatibilityType(compatibilityType)})
}
