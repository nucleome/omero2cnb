# OMERO-NB-index - An OMERO plugin for connecting Nucleome Browser with OMERO

## Introduction
OMERO-NB-index (formerly known as omero2cnb) allows users to interactively explore imaging datasets stored on a local OMERO server in a genome browser panel of Nucleome Browser. 
OMERO server uses a PostgreSQL database to store the annotation information of images.
OMERO-NB-index extracts and monitors the key-value map stored in the annotation_mapvalue table in PostgreSQL database.
If a key is named as `regions`,  and the format of value follows `genome:chromosome:start-end;...` , such as `hg38:chr1:1-20000;hg38:chr1:30000-40000`, OMERO-NB-index will parse these values into genome coordinates and store them into a [binning index](http://genomewiki.ucsc.edu/index.php/Bin_indexing_system) data structure and create a web service at http://127.0.0.1:3721.
Use then can add this web service in the genome browser of Nucleome Browser, and query images based genomic coordinates through the navigation on the genome browser.
By default, OMERO-NB-index monitors the change of annotation_mapvalue table every 90 seconds.

If you want to use this web service other than a local OMERO server, please use [reverse proxy](https://en.wikipedia.org/wiki/Reverse_proxy) software such [nginx](https://www.nginx.com/) or [traefik](https://github.com/containous/traefik).

## Install
```
go get github.com/nucleome/omero2cnb
```

## Usage
```
omero2cnb [db_host] [dbname] [db_user] [db_passwd] [omero_web_server]
```

## Docker
- [OMERO-NB-index in Docker Hub](https://hub.docker.com/repository/docker/nimezhu/cnb-index-omero)
- [Dockerfile in GitHub](https://github.com/nimezhu/cnb-index-omero-docker)
This docker file requires a binary executable file. 
You can generate it using the following command.
```
git clone https://github.com/nucleome/omero2cnb
cd omero2cnb
env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build
```

## Docker Compose with OMERO
User can start a docker compose along with OMERO and OMERO.web.
- [Docker compose file](https://gist.github.com/nimezhu/920130590d9a288be61d35971e11857f)

