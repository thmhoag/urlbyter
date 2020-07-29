package cli

import (
	"github.com/thmhoag/urlbyter/pkg/cli/completion"
	"github.com/thmhoag/urlbyter/pkg/cli/root"
	"github.com/thmhoag/urlbyter/pkg/cli/version"
	"github.com/thmhoag/urlbyter/pkg/urlbyter"
	"net/http"
)

const name = "urlbyter"

func Execute() error {
	client := urlbyter.NewClient(http.DefaultClient)
	batchRunner := urlbyter.NewBatchRunner(&urlbyter.BatchOpts{
		Client: client,
	})

	ctx := &globalCtx{
		batchRunner: batchRunner,
		version: version.Properties{
			Semver: Semver,
			Commit: Commit,
			Built: Built,
			GoVersion: GoVersion,
		},
	}

	rootCmd := root.NewRootCmd(ctx)
	rootCmd.Version = ctx.CurrentVersion().Semver

	rootCmd.AddCommand(completion.NewCompletionCmd(ctx))
	rootCmd.AddCommand(version.NewVersionCmd(ctx))

	return rootCmd.Execute()
}