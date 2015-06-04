# SRASearch

An NCBI Sequence Read Archive (SRA) upload management search utility.

The (SRA) regularly provides a set of [batch telemetry file][2] for laboratories that submit data to its repository.  The telemetry files can help correlate the data a submitter has sent with what the SRA has received and processed.

This utility processes those telemetry files and presents that data through a "Google"-esque search interface.

## BUILDING

_The instructions assume that you have access to a [Go][7] compiler._

### Setup a Go workspace

Initialize a [proper Go development workspace][1]:

    mkdir /path/to/project
    cd /path/to/project
    export GOPATH=$PWD
    export GOROOT=$(go env GOROOT)

Setup the git repository:

    mkdir -p $GOPATH/src/github.com/indraniel/
    cd $GOPATH/src/github.com/indraniel/
    git clone git@github.com:indraniel/srasearch.git
    cd srasearch/

### Initialize the depdencies

    make prepare

### Build the app

    make

You should now see a `srasearch` executable inside the `$GOPATH/src/github.com/indraniel/srasearch` directory.  You can move that file to wherever you please.

## USAGE

For the examples below we are using the SRA batch telemetry files available to the [McDonnell Genome Institute][4]; which has a NCBI center abbreviation name `WUGSC`.

The SRA provides monthly full and daily incremental versions of the batch telemetry files.  Using these we can create an intermediary file, an "SRA Dump", which is collection of JSON documents.

`srasearch` will incrementally build upon prior "SRA Dumps" as new information is collected.

### Initializing a SRA Dump

This command uses the full metadata and data transfer telemetry files produced at the beginning of the month.  For example, here is how the command would run on May 1, 2015:

    srasearch init-dump -m data/SRA/NCBI_SRA_Metadata_Full_WUGSC_20150501.tar.gz -u data/SRA/NCBI_SRA_Files_Full_WUGSC_20150501.gz -o 2015-05-01.dump.dat.gz

Here we initialized a SRA Dump file named `2015-05-01.dump.dat.gz`.

### Incrementing an existing SRA Dump

This command uses the incremental metadata and data transfer telemetry files and a prior existing SRA Dump file.  For example, this is how the command would run on May 2, 2015:

    srasearch increment-dump -i 2015-05-01.dump.dat.gz -m /path/to/NCBI_SRA_Metadata_WUGSC_20150502.tar.gz -u data/SRA/NCBI_SRA_Files_WUGSC_20150502.gz -o 2015-05-02.dump.dat.gz

Here we initialized a SRA Dump file named `2015-05-02.dump.dat.gz`.

On May 3, 2015, the command to create a new updated "SRA Dump" file would look like so:

    srasearch increment-dump -i 2015-05-02.dump.dat.gz -m /path/to/NCBI_SRA_Metadata_WUGSC_20150503.tar.gz -u data/SRA/NCBI_SRA_Files_WUGSC_20150503.gz -o 2015-05-03.dump.dat.gz

We can proceed onwards simliarly through the rest of the month.  Once the next month arrives we can initialize a brand new SRA Dump again.

### Creating a search index from a given SRA Dump

Given an SRA Dump file, a primary search index database can be constructed.  This sub-command creates the primary search index database:

    srasearch make-index -i 2015-05-03.dump.dat.gz -o /path/to/db/srasearch0503.idx

### Creating a recent uploads file (optional)

This sub-command creates an abbreviated "recent uploads" TSV file:

    srasearch-noweb make-uploads --ncbi-uploads="/path/to/SRA/NCBI_SRA_Files_WUGSC_20150503.gz" --output-dir="/path/to/db/srasearch0503.idx" --threshold=4000 
    
In this example, we've placed the last recent 4000 uploads (from May 3, 2015) as a file inside the primary search index database directory called `/path/to/db/srasearch0503.idx/recent-4000-sra-uploads-20150523.tsv`.

_This file is generally placed within a search index directory.  It provides the data for the "Recent Uploads" link in the web app._

### Start up the web app

    srasearch web --host="0.0.0.0" --port=9999 --index-path="/path/to/db/srasearch0503.idx"

## NOTES

`srasearch` is using the [bleve][3] text indexing library for the underlying search engine.  [BoltDB][5] is being used for bleve's underlying key/value store.

All the dependecies to this app are stored within this repository and are managed by [godep][6].

## LICENSE

ISC

[1]: http://golang.org/doc/code.html
[2]: http://www.ncbi.nlm.nih.gov/books/NBK242623/#SRA_Submission_Telemetry_BK.Batch_Teleme
[3]: https://github.com/blevesearch/bleve
[4]: http://genome.wustl.edu
[5]: https://github.com/boltdb/bolt
[6]: https://github.com/tools/godep
[7]: https://golang.org/doc/install
