package commands

import (
	"github.com/indraniel/srasearch/makeuploads"
	"github.com/indraniel/srasearch/utils"

	"github.com/spf13/cobra"

	"log"
	"os"
)

type MkUploadsCmdOpts struct {
	uploads   string
	threshold int
	outPath   string
}

var MkUploadsOpts MkUploadsCmdOpts

var cmdMakeUploads = &cobra.Command{
	Use:   "make-uploads -u <ncbi-uploads.gz> -t <threshold> -o <output-dir>",
	Short: "Generate an search index file from a SRA Dump file",
	Long:  `Generate an search index file from a SRA Dump file`,
	Run: func(cmd *cobra.Command, args []string) {
		MkUploadsOpts.main()
	},
}

func init() {
	cmdMakeUploads.Flags().StringVarP(
		&MkUploadsOpts.uploads,
		"ncbi-uploads",
		"u",
		"NCBI_SRA_Files_Full_WUGSC_20150501.gz",
		"the full NCBI SRA Uploads gzip file (NCBI_SRA_Files_Full_*.gz)",
	)

	cmdMakeUploads.Flags().IntVarP(
		&MkUploadsOpts.threshold,
		"threshold",
		"t",
		1000,
		"number of recent entries to select from the NCBI SRA Uploads gzip file",
	)

	cmdMakeUploads.Flags().StringVarP(
		&MkUploadsOpts.outPath,
		"output-dir",
		"o",
		os.Getenv("PWD"),
		"the directory path to place the pruned recent uploads file",
	)

}

func (opts MkUploadsCmdOpts) main() {
	opts.processOpts()
	makeuploads.CreateRecentUploadsFile(
		opts.uploads,
		opts.outPath,
		opts.threshold,
	)
}

func (opts MkUploadsCmdOpts) processOpts() {
	if opts.uploads == "" {
		log.Fatal(
			"Please supply the NCBI SRA Uploads file",
			"via --ncbi-uploads= !",
		)
	}

	if opts.outPath == "" {
		log.Fatal("Please supply an output directory via --output-dir !")
	}

	utils.CheckFileExists(opts.uploads)
	utils.CheckFileExists(opts.outPath)
}
