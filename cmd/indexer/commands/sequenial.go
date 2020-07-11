package commands

import (
	"context"
	"fmt"

	"github.com/mailchain/mailchain/cmd/indexer/internal/processor"
	"github.com/spf13/cobra"
)

func doSequential(cmd *cobra.Command, p *processor.Sequential) {
	for {
		err := p.NextBlock(context.Background())

		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "%+v", err)
		}
	}
}
