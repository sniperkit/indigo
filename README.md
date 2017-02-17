# Indigo

The Indigo is an index server written in [Go](https://golang.org), built on top of the [Bleve](http://www.blevesearch.com).


## Indigo gRPC Server

The Indigo gRPC Server is an index server over [gRPC](http://www.grpc.io).

### Start Indigo gRPC Server

The `indigo start grpc` command starts the Indigo gRPC Server.

```sh
$ indigo start grpc
```


## Indigo REST Server

The Indigo REST Server is a [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) Web Server that communicates with The Indigo gRPC Server.

### Start Indigo gRPC Server

The `indigo start rest` command starts the Indigo REST Server.

```sh
$ indigo start rest
```


## The index mapping

The index mapping describes how to your data model should be indexed. See following example.

### mapping.json

The mapping.json file contains all of the details about which fields your documents can contain, and how those fields should be dealt with when adding documents to the index, or when querying those fields.
See [Introduction to Index Mappings](http://www.blevesearch.com/docs/Index-Mapping/) and [type IndexMappingImpl](https://godoc.org/github.com/blevesearch/bleve/mapping#IndexMappingImpl) for more details.

```json
{
  "types": {
    "document": {
      "enabled": true,
      "dynamic": true,
      "properties": {
        "category": {
          "enabled": true,
          "dynamic": true,
          "fields": [
            {
              "type": "text",
              "analyzer": "keyword",
              "store": true,
              "index": true,
              "include_term_vectors": true,
              "include_in_all": true
            }
          ],
          "default_analyzer": ""
        },
        "description": {
          "enabled": true,
          "dynamic": true,
          "fields": [
            {
              "type": "text",
              "analyzer": "en",
              "store": true,
              "index": true,
              "include_term_vectors": true,
              "include_in_all": true
            }
          ],
          "default_analyzer": ""
        },
        "name": {
          "enabled": true,
          "dynamic": true,
          "fields": [
            {
              "type": "text",
              "analyzer": "en",
              "store": true,
              "index": true,
              "include_term_vectors": true,
              "include_in_all": true
            }
          ],
          "default_analyzer": ""
        },
        "popularity": {
          "enabled": true,
          "dynamic": true,
          "fields": [
            {
              "type": "number",
              "store": true,
              "index": true,
              "include_in_all": true
            }
          ],
          "default_analyzer": ""
        },
        "release": {
          "enabled": true,
          "dynamic": true,
          "fields": [
            {
              "type": "datetime",
              "store": true,
              "index": true,
              "include_in_all": true
            }
          ],
          "default_analyzer": ""
        },
        "type": {
          "enabled": true,
          "dynamic": true,
          "fields": [
            {
              "type": "text",
              "analyzer": "keyword",
              "store": true,
              "index": true,
              "include_term_vectors": true,
              "include_in_all": true
            }
          ],
          "default_analyzer": ""
        }
      },
      "default_analyzer": ""
    }
  },
  "default_mapping": {
    "enabled": true,
    "dynamic": true,
    "default_analyzer": ""
  },
  "type_field": "type",
  "default_type": "document",
  "default_analyzer": "standard",
  "default_datetime_parser": "dateTimeOptional",
  "default_field": "_all",
  "store_dynamic": true,
  "index_dynamic": true,
  "analysis": {}
}
```

## The Indigo Command Line Interface


### Create the index to the Indigo gRPC Server via CLI

```sh
$ indigo create index example "$(cat example/mapping.json)" -s boltdb -t upside_down
```

### Delete the index from the Indigo gRPC Server via CLI

```sh
$ indigo delete index example
```

### Get the index mapping from the Indigo gRPC Server via CLI

```sh
$ indigo get mapping example | jq .
```

The result of the above mapping command is:

```json
{
  "types": {
    "document": {
      "enabled": true,
      "dynamic": true,
      "properties": {
        "category": {
          "enabled": true,
          "dynamic": true,
          "fields": [
            {
              "type": "text",
              "analyzer": "keyword",
              "store": true,
              "index": true,
              "include_term_vectors": true,
              "include_in_all": true
            }
          ],
          "default_analyzer": ""
        },
        "description": {
          "enabled": true,
          "dynamic": true,
          "fields": [
            {
              "type": "text",
              "analyzer": "en",
              "store": true,
              "index": true,
              "include_term_vectors": true,
              "include_in_all": true
            }
          ],
          "default_analyzer": ""
        },
        "name": {
          "enabled": true,
          "dynamic": true,
          "fields": [
            {
              "type": "text",
              "analyzer": "en",
              "store": true,
              "index": true,
              "include_term_vectors": true,
              "include_in_all": true
            }
          ],
          "default_analyzer": ""
        },
        "popularity": {
          "enabled": true,
          "dynamic": true,
          "fields": [
            {
              "type": "number",
              "store": true,
              "index": true,
              "include_in_all": true
            }
          ],
          "default_analyzer": ""
        },
        "release": {
          "enabled": true,
          "dynamic": true,
          "fields": [
            {
              "type": "datetime",
              "store": true,
              "index": true,
              "include_in_all": true
            }
          ],
          "default_analyzer": ""
        },
        "type": {
          "enabled": true,
          "dynamic": true,
          "fields": [
            {
              "type": "text",
              "analyzer": "keyword",
              "store": true,
              "index": true,
              "include_term_vectors": true,
              "include_in_all": true
            }
          ],
          "default_analyzer": ""
        }
      },
      "default_analyzer": ""
    }
  },
  "default_mapping": {
    "enabled": true,
    "dynamic": true,
    "default_analyzer": ""
  },
  "type_field": "type",
  "default_type": "document",
  "default_analyzer": "standard",
  "default_datetime_parser": "dateTimeOptional",
  "default_field": "_all",
  "store_dynamic": true,
  "index_dynamic": true,
  "analysis": {}
}
```

### Index the documents to the Indigo gRPC Server via CLI

```sh
$ indigo index documents example "$(cat example/index_documents.json)"
```

The result of the above index command is:

```json
{
  "document_count": 5
}
```

### Delete the documents from the Indigo gRPC Server via CLI

```sh
$ indigo delete documents example "$(cat example/delete_documents.json)"
```

The result of the above delete command is:

```json
{
  "document_count": 3
}
```


### Search the documents frmo the Indigo gRPC Server via CLI

See [Queries](http://www.blevesearch.com/docs/Query/), [Query String Query](http://www.blevesearch.com/docs/Query-String-Query/) and [type SearchRequest](https://godoc.org/github.com/blevesearch/bleve#SearchRequest) for more details.

#### Simple query

```sh
$ indigo search document example "$(cat example/simple_query.json)" | jq .
```

The result of the above simple query command is:

```json
{
  "status": {
    "total": 1,
    "failed": 0,
    "successful": 1
  },
  "request": {
    "query": {
      "query": "description:Go"
    },
    "size": 10,
    "from": 0,
    "highlight": null,
    "fields": [
      "name",
      "description",
      "category",
      "popularity",
      "release",
      "type"
    ],
    "facets": null,
    "explain": false,
    "sort": [
      "-_score"
    ],
    "includeLocations": false
  },
  "hits": [
    {
      "index": "./index",
      "id": "1",
      "score": 0.40824830532073975,
      "sort": [
        "_score"
      ],
      "fields": {
        "category": "Library",
        "description": "Full-text search library written in Go.",
        "name": "Bleve",
        "popularity": 1,
        "release": "2014-04-18T00:00:00Z",
        "type": "document"
      }
    }
  ],
  "total_hits": 1,
  "max_score": 0.40824830532073975,
  "took": 220795,
  "facets": {}
}
```

See [type SearchResult](https://godoc.org/github.com/blevesearch/bleve#SearchResult) for more details.


## License

Apache License Version 2.0
