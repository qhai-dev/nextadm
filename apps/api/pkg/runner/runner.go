package runner

import (
	"fmt"
	"os"

	"github.com/nextadm/pkg/logger"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command) int {
	if nlogInitialized, err := run(cmd); err != nil {
		if !nlogInitialized {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		} else {
			// TODO command 相关 error
		}
		return 1
	}
	return 0
}

func run(cmd *cobra.Command) (nlogInitialized bool, err error) {
	defer logger.FlushLog()

	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		for _, arg := range args {
			if len(arg) > 0 {
				return fmt.Errorf("%q 不需要任何参数 %q", cmd.CommandPath(), args)
			}
		}
		return nil
	}

	switch {
	case cmd.PersistentPreRun != nil:
		pre := cmd.PersistentPreRun
		cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
			logger.InitLog()
			nlogInitialized = true
			pre(cmd, args)
		}
	default:
		cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
			logger.InitLog()
			nlogInitialized = true
		}
	}

	err = cmd.Execute()

	return
}
