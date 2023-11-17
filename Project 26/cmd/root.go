package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Execute() {
	cmd := &cobra.Command{
		Use: "go-grpc",
	}

	cmd.AddCommand(
		Api,
		Client,
	)

	if err :=cmd.Execute(); err != nil {
        fmt.Println(err)
    }
}