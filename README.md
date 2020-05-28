# Running title: OMERO2CNB

## Introduction
OMERO use PostgreSQL database to store annotation information of images.
Omero2cnb extracts and monitors the key-value map stored in annotation_mapvalue table in PostgreSQL database.
If any key is "regions",  and value format is "genome:chromosome:start-end;...", omero2cnb will
parse this values into genome coordinates and put them into a memory binning index structure and provides a web data service http://127.0.0.1:3721 of query genome choordinates for Nucleome Browser.

Omero2cnb monitors the change of annotation_mapvalue table every 90 seconds.


## Install
```
go get github.com/nimezhu/omero2cnb
```

## Usage
```
omero2cnb [db_host] [dbname] [db_user] [db_passwd] [omero_web_server]
```

## Docker
- [omero-cnb in Docker Hub](https://hub.docker.com/repository/docker/nimezhu/cnb-index-omero)
- [Dockerfile in GitHub](https://github.com/nimezhu/cnb-index-omero-docker)

## Docker Compose with OMERO
User can start a docker compose along with OMERO and OMERO-web.
- [Docker compose file](https://gist.github.com/nimezhu/920130590d9a288be61d35971e11857f)

