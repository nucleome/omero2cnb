# Running title: OMERO2CNB

## Introduction
OMERO server use PostgreSQL database to store annottation information of images.

Omero2cnb read and monitor the key-value map stored in PostgreSQL annotation_mapvalue table.
and identified key "regions" and value format such as "genome:chromosome:start-end;..."
parse the genome coordinates and put them into a memory binning index structure
provide a web data service of query genome choordinates for Nucleome Browser

## Usage
```
omero2cnb [db_host] [dbname] [user] [passwd] [omero_server]
```
