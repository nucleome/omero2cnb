# Running title: OMERO2CNB

## Introduction
OMERO use PostgreSQL database to store annotation information of images.
Omero2cnb extracts and monitors the key-value map stored in annotation_mapvalue table in PostgreSQL database.
If any key is "regions",  and value format is "genome:chromosome:start-end;...", omero2cnb will
parse this values into genome coordinates and put them into a memory binning index structure and provides a web data service of query genome choordinates for Nucleome Browser.


## Usage
```
omero2cnb [db_host] [dbname] [db_user] [db_passwd] [omero_server]
```
