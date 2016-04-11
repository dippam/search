# DIPPAM Search API

This project provides the search API for DIPPAM, replacing the Rails/SOLR combination.

## Libraries

- gorilla: URI router
- bolt: KV store for index.
- bleve: Indexer

## Notes

Collection Objects are modeled as go structs.

Document metadata is stored in a RDBMS and accessed using database/sql.

bleve indexes the text of documents.
