package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func NewVendor6Command() *cobra.Command {

	var rootCmd = &cobra.Command{
		Use:                   "vendor6",
		Short:                 "",
		Long:                  ``,
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := Vendor6(os.Stdin, os.Stdout, os.Stderr)
			if err != nil {
				return err
			}
			return nil
		},
	}

	

	return rootCmd
}

func Vendor6(in io.Reader, out, eout io.Writer) error {
	fmt.Fprintln(out, "Hello, IoT Vendors!")
	return nil
}

func Execute() {
	err := NewVendor6Command().Execute()
	if err != nil {
		os.Exit(1)
	}
}
