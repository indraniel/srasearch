A SRA Search utility.

    NO - go-bindata -o assets.go -ignore=web/static/js/bootstrap.js -ignore=web/static/js/jquery.js web/static/... web/views/...
    NO - go-bindata -o assets.go -pkg="assets" -ignore=web/static/js/bootstrap.js -ignore=web/static/js/jquery.js web/static/... web/views/...
    YES! - go-bindata -o assets/assets.go -pkg="assets" -ignore=web/static/js/bootstrap.js -ignore=web/static/js/jquery.js web/static/... web/views/...
    YES (developing) - go-bindata -debug -o assets/assets.go -pkg="assets" -ignore=web/static/js/bootstrap.js -ignore=web/static/js/jquery.js web/static/... web/views/...
    go build -o srasearch main.go

    srasearch init-dump -m data/SRA/NCBI_SRA_Metadata_Full_WUGSC_20150501.tar.gz -u data/SRA/NCBI_SRA_Files_Full_WUGSC_20150501.gz -o 2015-05-01.dump.dat.gz
    srasearch increment-dump -i 2015-05-01.dump.dat.gz -m /path/to/NCBI_SRA_Metadata_WUGSC_20150502.tar.gz -u data/SRA/NCBI_SRA_Files_WUGSC_20150502.gz -o 2015-05-02.dump.dat.gz
    srasearch increment-dump -i 2015-05-02.dump.dat.gz -m /path/to/NCBI_SRA_Metadata_WUGSC_20150503.tar.gz -u data/SRA/NCBI_SRA_Files_WUGSC_20150503.gz -o 2015-05-03.dump.dat.gz

    # ... and so on...

    srasearch make-index -i 2015-05-03.dump.dat.gz -o /path/to/db/srasearch0503.idx
    srasearch-noweb make-uploads --ncbi-uploads="/path/to/SRA/NCBI_SRA_Files_WUGSC_20150503.gz" --output-dir="/path/to/db/srasearch0503.idx" --threshold=4000 
    srasearch web --host="0.0.0.0" --port=9999 --index-path="/path/to/db//srasearch0503.idx"
