package commands

import (
	"github.com/indraniel/srasearch/incrementdump"
	"github.com/indraniel/srasearch/utils"
	"github.com/spf13/cobra"
	"log"
)

type IncrementDumpCmdOpts struct {
	metadata   string
	uploads    string
	inputDump  string
	outputDump string
}

var IncrementDumpOpts IncrementDumpCmdOpts

var cmdIncrementDump = &cobra.Command{
	Use:   "increment-dump -i <prior-dump.gz> -m <ncbi-metadata.tar.gz> -u <ncbi-uploads.gz> -o <new-dump.gz>",
	Short: "Incorporate new NCBI Batch Telemetry incrementals files to an existing SRA Dump",
	Long:  `Update an existing SRA Dump file given new incremental NCBI Batch Telemetry file information`,
	Run: func(cmd *cobra.Command, args []string) {
		IncrementDumpOpts.main()
	},
}

func init() {
	cmdIncrementDump.Flags().StringVarP(
		&IncrementDumpOpts.metadata,
		"ncbi-metadata",
		"m",
		"NCBI_SRA_Metadata_WUGSC_20150502.tar.gz",
		"the incremental NCBI SRA Metadata tar.gz file (NCBI_SRA_Metadata_*.tar.gz)",
	)

	cmdIncrementDump.Flags().StringVarP(
		&IncrementDumpOpts.uploads,
		"ncbi-uploads",
		"u",
		"NCBI_SRA_Files_WUGSC_20150502.gz",
		"the incremental NCBI SRA uploads gzip file (NCBI_SRA_Files_*.gz)",
	)

	cmdIncrementDump.Flags().StringVarP(
		&IncrementDumpOpts.inputDump,
		"input-dump",
		"i",
		"2015-05-01-dump.dat.gz",
		"an existing SRA dump file to update",
	)

	cmdIncrementDump.Flags().StringVarP(
		&IncrementDumpOpts.outputDump,
		"output",
		"o",
		"2015-05-02.dump.dat.gz",
		"an updated SRA dump file to create",
	)
}

func (opts IncrementDumpCmdOpts) main() {
	opts.processOpts()
	incrementdump.Main(opts.metadata, opts.uploads, opts.inputDump, opts.outputDump)
}

func (opts IncrementDumpCmdOpts) processOpts() {
	if opts.outputDump == "" {
		log.Fatal(
			"Please supply an gzipped output file dump to create",
			"via --output !",
		)
	}

	if opts.metadata == "" {
		log.Fatal(
			"Please supply a NCBI incremental input tar file ",
			"via --ncbi-metadata (use --help for more info) !",
		)
	}

	if opts.uploads == "" {
		log.Fatal(
			"Please supply a NCBI incremental input uploads file ",
			"via --ncbi-uploads (use --help for more info) !",
		)
	}

	if opts.inputDump == "" {
		log.Fatal(
			"Please supply an existing SRA Dump file ",
			"to update against via --input-dump (use --help for more info) !",
		)
	}

	utils.CheckFileExists(opts.metadata)
	utils.CheckFileExists(opts.uploads)
	utils.CheckFileExists(opts.inputDump)
}
