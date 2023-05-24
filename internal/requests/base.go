package requests

import "github.com/spf13/cobra"

// Base encapsulates the required information needed to make requests to the API
type Base struct {
	Cmd *cobra.Command
}
