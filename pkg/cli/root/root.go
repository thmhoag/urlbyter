package root

import (
	"github.com/spf13/cobra"
	"github.com/thmhoag/clif"
	"log"
	"os"
	"runtime"
)

type rootOpts struct {
	showVersion  bool
	formatString string
}

func NewRootCmd(ctx Ctx) *cobra.Command {
	opts := &rootOpts{}
	client := ctx.BatchRunner()

	cmd := &cobra.Command{
		Use:   "urlbyter [path]",
		Short: "gets the sizes of pages at a list of URLs",
		Args: cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			// set numCPU to half the available logical cores by default
			setNumCPU()
		},
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			results, err := client.ProcessFile(path)
			if err != nil {
				cmd.PrintErrf("unable to process URL list:\n%s\n", err)
				os.Exit(1)
			}

			if err := clif.New(opts.formatString).
				Output(cmd.OutOrStdout()).
				Write(results); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.SetVersionTemplate("{{.Version}}\n")
	cmd.Flags().BoolVarP(&opts.showVersion, "version", "v", false, "return the version of the executable")
	cmd.Flags().StringVar(&opts.formatString, "format", "table {{ .Host }} {{ .URL }} {{ .Bytes }}", "provide a template to format output")

	return cmd
}

func setNumCPU() {
	totalLogicalCores := runtime.NumCPU()
	if totalLogicalCores < 2 {
		return
	}

	numCoresToUse := totalLogicalCores / 2
	runtime.GOMAXPROCS(numCoresToUse)
}