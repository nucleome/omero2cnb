# NucleIndexOMERO
  previous name : omero2cnb
## Introduction
OMERO server use PostgreSQL to store meta information for images.

NucleIndexOMERO read and monitor the key-value map stored in PostgreSQL table ...
and identified key "regions" and value format such as "genome:chromosome:start-end;..."
read them into a memory binning index structure
provide a web data service for query genome choordinates

## Usage
```
omero2cnb [db_host] [dbname] [user] [passwd] [omero_server]
```
