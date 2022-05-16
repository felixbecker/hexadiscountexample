package cli

import (
	"fmt"
	"os"

	"github.com/felixbecker/hexadiscountexample/application"
	"github.com/spf13/cobra"
)

type Cli struct {
	application *application.Application
}

func New(app *application.Application) *Cli {
	c := Cli{}
	c.application = app
	return &c
}

type CalculationOptions struct {
	Amount float32
}

func (c *CalculationOptions) Validate() error {

	if c.Amount == 0 {
		return fmt.Errorf("error Amount cannot be 0")
	}
	return nil
}

func makeCalculateCommand(app *application.Application) *cobra.Command {

	opts := CalculationOptions{}
	var cmd = cobra.Command{
		Use:   "calculate",
		Short: "does the actual calculation",
		Long:  "does the actual calculation",
		PreRunE: func(cmd *cobra.Command, args []string) error {

			err := opts.Validate()
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			result := app.Discount(opts.Amount)
			fmt.Printf("The result: %f is calculated based on the amount: %f\n", result, opts.Amount)
			return nil
		},
	}
	cmd.Flags().Float32VarP(&opts.Amount, "amount", "a", 1.0, "please specify the amount for what you want to calculate your discount. [Default = 1.0]")
	return &cmd
}

func makeRootCommand(app *application.Application) *cobra.Command {
	var cmd = cobra.Command{
		Use:   "discounter",
		Short: "calculates the discount for a given amount",
		Long:  "calculates the discount for a given amount",
	}

	cmd.AddCommand(makeCalculateCommand(app))
	return &cmd
}

func (c *Cli) Execute() {

	rootCmd := makeRootCommand(c.application)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
