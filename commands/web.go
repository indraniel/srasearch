package commands

import (
	"github.com/indraniel/srasearch/web"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type WebCmdOpts struct {
	Port      int
	Host      string
	IndexPath string
	Debug     bool
}

var WebOpts WebCmdOpts

var cmdWeb = &cobra.Command{
	Use:   "web",
	Short: "A web portal to inspect an SRA search index",
	Long:  `A web portal to inspect an SRA search index`,
	Run: func(cmd *cobra.Command, args []string) {
		WebOpts.processOpts()
		WebOpts.main()
	},
}

func init() {
	cmdWeb.Flags().IntVarP(
		&WebOpts.Port,
		"port",
		"p",
		9999,
		"the web port",
	)

	cmdWeb.Flags().StringVarP(
		&WebOpts.Host,
		"host",
		"H",
		"127.0.0.1",
		"the web host",
	)

	cmdWeb.Flags().StringVarP(
		&WebOpts.IndexPath,
		"index-path",
		"i",
		"",
		"path to the bleve index dir",
	)

	cmdWeb.Flags().BoolVarP(
		&WebOpts.Debug,
		"debug",
		"d",
		false,
		"enable debug mode on the web interface",
	)
}

func (opts WebCmdOpts) main() {
	w := web.NewWeb(opts.Port, opts.Host, opts.IndexPath, opts.Debug)
	w.Main()
}

func (opts WebCmdOpts) processOpts() {
	if opts.IndexPath == "" {
		log.Fatal(
			"Please supply an bleve index directory path ",
			"via --index-path !",
		)
	}

	checkDirExists(opts.IndexPath)
}

func checkDirExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf(
			"Could not find '%s' on file system: %s",
			dir, err,
		)
	}
}
