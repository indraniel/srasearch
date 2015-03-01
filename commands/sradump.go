package commands

import (
	"github.com/indraniel/srasearch/sradump"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

type SraDumpCmdOpts struct {
	output string
}

var SraDumpOpts SraDumpCmdOpts

var cmdSraDump = &cobra.Command{
	Use:   "sra-dump [path/to/NCBIDownloadTarFile]",
	Short: "Transform the NCBI Batch Telemetry tar files to a set of JSON Docs",
	Long: `This command transforms the raw NCBI Batch Telemetry tar file
         contents into a intermediary file of custom JSON Documents`,
	Run: func(cmd *cobra.Command, args []string) {
		tarfile := getTarFile(args)
		checkTarExists(tarfile)
		makeSraDump(tarfile)
	},
}

func init() {
	cmdSraDump.Flags().StringVarP(
		&SraDumpOpts.output,
		"output",
		"o",
		"sradump.sjd.gz",
		"the output file to dump the serialized JSON Documents to",
	)
}

func makeSraDump(tarfile string) {
	log.Println("Collecting Accession Stats")
	db := sradump.CollectAccessionStats(tarfile)

	log.Println("Processing XMLs / Creating Dump File")

	tmpdir, tmpfile := makeTmpFile()
	defer os.Remove(tmpfile)
	defer os.Remove(tmpdir)
	log.Println("Tmp Dump File is:", tmpfile)
	sradump.ProcessTarXMLs(tarfile, db, tmpfile)

	log.Println("Compressing Dump File")
	err := sradump.CompressDumpFile(tmpfile, SraDumpOpts.output)
	if err != nil {
		log.Print("Trouble making gzip file:", err)
		return
	}
	log.Println("All Done!")
}

func makeTmpFile() (tmpdir, tmpfile string) {
	tmpdir, err := ioutil.TempDir(os.TempDir(), "sra-dump")
	if err != nil {
		log.Fatal("Trouble making temp dir:", err)
	}

	f, err := ioutil.TempFile(tmpdir, "sra-tmp-dump")
	if err != nil {
		log.Fatal("Trouble making temp file:", err)
	}
	defer f.Close()

	tmpfile = f.Name()
	return
}

func getTarFile(args []string) (tarfile string) {
	if len(args) == 0 {
		log.Fatal("Please supply a tar file as an argument!")
	}

	tarfile = args[0]

	if tarfile == "" {
		log.Fatal("Please supply a tar file as an argument!")
	}

	return
}

func checkTarExists(tarfile string) {
	if _, err := os.Stat(tarfile); os.IsNotExist(err) {
		log.Fatalf(
			"Could not find '%s' on file system: %s",
			tarfile, err,
		)
	}
}
