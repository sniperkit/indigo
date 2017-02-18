# Indigo

The Indigo is an index server written in [Go](https://golang.org), built on top of the [Bleve](http://www.blevesearch.com).


## Indigo gRPC Server

The Indigo gRPC Server is an index server over [gRPC](http://www.grpc.io).

### Start Indigo gRPC Server

The `indigo start grpc` command starts the Indigo gRPC Server.

```sh
$ indigo start grpc
```

## The Indigo Command Line Interface

The Indigo provides some commands for control the Indigo Server.

### Create the index to the Indigo gRPC Server via CLI

```sh
$ indigo create index example "$(cat example/mapping.json)" -s boltdb -t upside_down
```

The result of the above `create index` command is:

```text
example created
```

### The index mapping

The index mapping describes how to your data model should be indexed. See following example.

#### mapping.json

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

### Delete the index from the Indigo gRPC Server via CLI

```sh
$ indigo delete index example
```

The result of the above `delete index` command is:

```text
example deleted
```

### Get the index stats from the Indigo gRPC Server via CLI

```sh
$ indigo get stats example | jq .
```

The result of the above `get stats` command is:

```json
{
  "index": {
    "analysis_time": 0,
    "batches": 0,
    "deletes": 0,
    "errors": 0,
    "index_time": 0,
    "num_plain_text_bytes_indexed": 0,
    "term_searchers_finished": 0,
    "term_searchers_started": 0,
    "updates": 0
  },
  "search_time": 0,
  "searches": 0
}
```

### Get the index mapping from the Indigo gRPC Server via CLI

```sh
$ indigo get mapping example | jq .
```

The result of the above `get mapping` command is:

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

### Index the document to the Indigo gRPC Server via CLI

```sh
$ indigo index document example "1" "$(cat example/index_document.json)"
```

The result of the above `index document` command is:

```text
1 document indexed
```


### Delete the document from the Indigo gRPC Server via CLI

```sh
$ indigo delete document example "1"
```

The result of the above `delete document` command is:

```text
1 document deleted
```


### Index the documents in bulk to the Indigo gRPC Server via CLI

```sh
$ indigo index bulk example "$(cat example/index_bulk.json)"
```

The result of the above `index bulk` command is:

```text
5 documents indexed in bulk
```

### Delete the documents in bulk from the Indigo gRPC Server via CLI

```sh
$ indigo delete bulk example "$(cat example/delete_bulk.json)"
```

The result of the above `delete bulk` command is:

```text
3 documents deleted in bulk
```


### Search the documents frmo the Indigo gRPC Server via CLI

See [Queries](http://www.blevesearch.com/docs/Query/), [Query String Query](http://www.blevesearch.com/docs/Query-String-Query/) and [type SearchRequest](https://godoc.org/github.com/blevesearch/bleve#SearchRequest) for more details.

#### Simple query

```sh
$ indigo search documents example "$(cat example/simple_query.json)" | jq .
```

The result of the above `search documents` command is:

```json
{
  "status": {
    "total": 1,
    "failed": 0,
    "successful": 1
  },
  "request": {
    "query": {
      "query": "description:*"
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
      "index": "data/example",
      "id": "1",
      "score": 0.2527852661670985,
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
    },
    {
      "index": "data/example",
      "id": "5",
      "score": 0.2527852661670985,
      "sort": [
        "_score"
      ],
      "fields": {
        "category": "Server",
        "description": "Full-text search server built on Bleve.",
        "name": "Indigo",
        "popularity": 5,
        "release": "2017-01-13T00:00:00Z",
        "type": "document"
      }
    }
  ],
  "total_hits": 2,
  "max_score": 0.2527852661670985,
  "took": 1991761,
  "facets": {}
}
```

See [type SearchResult](https://godoc.org/github.com/blevesearch/bleve#SearchResult) for more details.





## Indigo REST Server

The Indigo REST Server is a [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) Web Server that communicates with The Indigo gRPC Server.

### Start Indigo gRPC Server

The `indigo start rest` command starts the Indigo REST Server.

```sh
$ indigo start rest
```

### Create the index to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X PUT "http://localhost:2289/api/example" -H "Content-Type: application/json" -d @example/mapping.json -s | jq .
```

The result of the above command is:

```json
{
  "name": "example"
}
```

### Delete the index to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X DELETE "http://localhost:2289/api/example" -s | jq .
```

The result of the above command is:

```json
{
  "name": "example"
}
```

### Get the index stats from the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X GET "http://localhost:2289/api/example/stats" -s | jq .
```

The result of the above command is:

```json
{
  "stats": {
    "index": {
      "analysis_time": 0,
      "batches": 0,
      "deletes": 0,
      "errors": 0,
      "index_time": 0,
      "num_plain_text_bytes_indexed": 0,
      "term_searchers_finished": 0,
      "term_searchers_started": 0,
      "updates": 0
    },
    "search_time": 0,
    "searches": 0
  }
}
```

### Get the index mapping from the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X GET "http://localhost:2289/api/example/mapping" -s | jq .
```

The result of the above command is:

```json
{
  "mapping": {
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
}
```

### Index the document to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X PUT "http://localhost:2289/api/example/1" -H "Content-Type: application/json" -d @example/index_document.json -s | jq .
```

The result of the above command is:

```text
{
  "count": 1
}
```

### Delete the document from the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X DELETE "http://localhost:2289/api/example/1" -s | jq .
```

The result of the above command is:

```text
{
  "count": 1
}
```

### Index the documents in bulk to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X POST "http://localhost:2289/api/example/_bulk" -H "Content-Type: application/json" -d @example/index_bulk.json -s | jq .
```

The result of the above command is:

```text
{
  "count": 5
}
```

### Delete the documents in bulk to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X DELETE "http://localhost:2289/api/example/_bulk" -H "Content-Type: application/json" -d @example/delete_bulk.json -s | jq .
```

The result of the above command is:

```text
{
  "count": 5
}
```


### Search the documents from the Indigo gRPC Server via the Indigo REST Server

#### Simple query

```sh
$ curl -X POST "http://localhost:2289/api/example/_search" -H "Content-Type: application/json" -d @example/simple_query.json -s | jq .
```

The result of the above `search documents` command is:

```json
{
  "status": {
    "total": 1,
    "failed": 0,
    "successful": 1
  },
  "request": {
    "query": {
      "query": "description:*"
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
      "index": "data/example",
      "id": "1",
      "score": 0.2527852661670985,
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
    },
    {
      "index": "data/example",
      "id": "5",
      "score": 0.2527852661670985,
      "sort": [
        "_score"
      ],
      "fields": {
        "category": "Server",
        "description": "Full-text search server built on Bleve.",
        "name": "Indigo",
        "popularity": 5,
        "release": "2017-01-13T00:00:00Z",
        "type": "document"
      }
    }
  ],
  "total_hits": 2,
  "max_score": 0.2527852661670985,
  "took": 1991761,
  "facets": {}
}
```


## License

Apache License Version 2.0
