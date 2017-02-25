# Indigo

The Indigo is a full text search and indexing server written in [Go](https://golang.org), built on top of the [Bleve](http://www.blevesearch.com).


## Indigo gRPC Server

The Indigo gRPC Server is an index server over [gRPC](http://www.grpc.io).


## Indigo REST Server

The Indigo REST Server is a [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) Web Server that communicates with The Indigo gRPC Server.


## The Indigo Command Line Interface

The Indigo provides some commands for control the Indigo Server.


### Start Indigo gRPC Server

The `indigo start grpc` command starts the Indigo gRPC Server.

```sh
$ indigo start grpc
```


### Create the index to the Indigo gRPC Server

```sh
$ indigo create index -n example -m example/index_mapping.json -s boltdb -t upside_down -f json
```

The result of the above `create index` command is:

```json
{
  "indexName": "example"
}
```

### Open the index to the Indigo gRPC Server

```sh
$ indigo open index -n example -f json
```

The result of the above `close index` command is:

```json
{
  "indexName": "example"
}
```

### Close the index from the Indigo gRPC Server

```sh
$ indigo close index -n example -f json
```

The result of the above `close index` command is:

```json
{
  "indexName": "example"
}
```

### Delete the index from the Indigo gRPC Server

```sh
$ indigo delete index -n example -f json
```

The result of the above `delete index` command is:

```json
{
  "indexName": "example"
}
```

### Get the index stats from the Indigo gRPC Server

```sh
$ indigo get stats -n example -f json
```

The result of the above `get stats` command is:

```json
{
  "indexStats": {
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

### Get the index mapping from the Indigo gRPC Server

```sh
$ indigo get mapping -n example -f
```

The result of the above `get mapping` command is:

```json
{
  "indexMapping": {
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

### Put the document to the Indigo gRPC Server

```sh
$ indigo put document -n example -i 1 -d example/document_1.json -f json
```

The result of the above `put document` command is:

```json
{
  "success": true
}
```

### Get the document from the Indigo gRPC Server

```sh
$ indigo get document -n example -i 1 -f json
```

The result of the above `get document` command is:

```json
{
  "document": {
    "category": "Library",
    "description": "Full-text search library written in Go.",
    "name": "Bleve",
    "popularity": 1,
    "release": "2014-04-18T00:00:00Z",
    "type": "document"
  }
}
```

### Delete the document from the Indigo gRPC Server

```sh
$ indigo delete document -n example -i 1 -f json
```

The result of the above `delete document` command is:

```json
{
  "success": true
}
```


### Index the documents in bulk to the Indigo gRPC Server

```sh
$ indigo bulk -n example -b example/bulk_put.json -f json
```

The result of the above `index bulk` command is:

```text
{
  "putCount": 7
}
```


### Search the documents frmo the Indigo gRPC Server

See [Queries](http://www.blevesearch.com/docs/Query/), [Query String Query](http://www.blevesearch.com/docs/Query-String-Query/) and [type SearchRequest](https://godoc.org/github.com/blevesearch/bleve#SearchRequest) for more details.

#### Simple query

```sh
$ indigo search -n example -s example/simple_query.json -f json
```

The result of the above `search` command is:

```json
{
  "searchResult": {
    "facets": {},
    "hits": [
      {
        "fields": {
          "category": "Library",
          "description": "Apache Lucene is a high-performance, full-featured text search engine library written entirely in Java.",
          "name": "Lucene",
          "popularity": 4,
          "release": "2000-03-30T00:00:00Z",
          "type": "document"
        },
        "id": "2",
        "index": "data/example",
        "score": 0.28598991738818746,
        "sort": [
          "_score"
        ]
      },
      {
        "fields": {
          "category": "Server",
          "description": "Solr is an open source enterprise search platform, written in Java, from the Apache Lucene project.",
          "name": "Solr",
          "popularity": 5,
          "release": "2006-12-22T00:00:00Z",
          "type": "document"
        },
        "id": "5",
        "index": "data/example",
        "score": 0.2842565476963312,
        "sort": [
          "_score"
        ]
      },
      {
        "fields": {
          "category": "Library",
          "description": "Whoosh is a fast, featureful full-text indexing and searching library implemented in pure Python. ",
          "name": "Whoosh",
          "popularity": 3,
          "release": "2008-02-20T00:00:00Z",
          "type": "document"
        },
        "id": "3",
        "index": "data/example",
        "score": 0.2484309575477134,
        "sort": [
          "_score"
        ]
      },
      {
        "fields": {
          "category": "Server",
          "description": "Indigo is a full-text search and indexing server written in Go, built on top of Bleve.",
          "name": "Indigo",
          "popularity": 1,
          "release": "2017-01-13T00:00:00Z",
          "type": "document"
        },
        "id": "7",
        "index": "data/example",
        "score": 0.24526800905441196,
        "sort": [
          "_score"
        ]
      },
      {
        "fields": {
          "category": "Library",
          "description": "Ferret is a super fast, highly configurable search library written in Ruby.",
          "name": "Ferret",
          "popularity": 2,
          "release": "2005-10-01T00:00:00Z",
          "type": "document"
        },
        "id": "4",
        "index": "data/example",
        "score": 0.2057485301587168,
        "sort": [
          "_score"
        ]
      },
      {
        "fields": {
          "category": "Server",
          "description": "Elasticsearch is a search engine based on Lucene, written in Java.",
          "name": "Elasticsearch",
          "popularity": 5,
          "release": "2010-02-08T00:00:00Z",
          "type": "document"
        },
        "id": "6",
        "index": "data/example",
        "score": 0.11396329383474207,
        "sort": [
          "_score"
        ]
      },
      {
        "fields": {
          "category": "Library",
          "description": "Bleve is a full-text search and indexing library for Go.",
          "name": "Bleve",
          "popularity": 3,
          "release": "2014-04-18T00:00:00Z",
          "type": "document"
        },
        "id": "1",
        "index": "data/example",
        "score": 0.0853843602235094,
        "sort": [
          "_score"
        ]
      }
    ],
    "max_score": 0.28598991738818746,
    "request": {
      "explain": false,
      "facets": null,
      "fields": [
        "name",
        "description",
        "category",
        "popularity",
        "release",
        "type"
      ],
      "from": 0,
      "highlight": null,
      "includeLocations": false,
      "query": {
        "query": "description:*"
      },
      "size": 10,
      "sort": [
        "-_score"
      ]
    },
    "status": {
      "failed": 0,
      "successful": 1,
      "total": 1
    },
    "took": 2.543262e+06,
    "total_hits": 7
  }
}
```

See [type SearchResult](https://godoc.org/github.com/blevesearch/bleve#SearchResult) for more details.


### Start Indigo REST Server

The `indigo start rest` command starts the Indigo REST Server.

```sh
$ indigo start rest
```

### Create the index to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X PUT "http://localhost:2289/api/example?indexType=upside_down&indexStore=boltdb" -H "Content-Type: application/json" --data-binary @example/index_mapping.json -s
```

The result of the above command is:

```json
{
  "indexName": "example"
}
```

### Open the index to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X POST "http://localhost:2289/api/example/_open" -H "Content-Type: application/json" --data-binary @example/runtime_config.json -s
```

The result of the above command is:

```json
{
  "indexName": "example"
}
```

### Close the index to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X POST "http://localhost:2289/api/example/_close" -s
```

The result of the above command is:

```json
{
  "indexName": "example"
}
```

### Delete the index to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X DELETE "http://localhost:2289/api/example" -s
```

The result of the above command is:

```json
{
  "name": "example"
}
```

### Get the index stats from the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X GET "http://localhost:2289/api/example/_stats" -s
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
$ curl -X GET "http://localhost:2289/api/example/_mapping" -s | jq .
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

### Put the document to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X PUT "http://localhost:2289/api/example/1" -H "Content-Type: application/json" --data-binary @example/document_1.json -s | jq .
```

The result of the above command is:

```json
{
  "count": 1
}
```

### Get the document to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X GET "http://localhost:2289/api/example/1" -s | jq .
```

The result of the above command is:

```json
{
  "document": {
    "category": "Library",
    "description": "Full-text search library written in Go.",
    "name": "Bleve",
    "popularity": 1,
    "release": "2014-04-18T00:00:00Z",
    "type": "document"
  }
}
```

### Delete the document from the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X DELETE "http://localhost:2289/api/example/1" -s | jq .
```

The result of the above command is:

```json
{
  "count": 1
}
```

### Index the documents in bulk to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -X POST "http://localhost:2289/api/example/_bulk" -H "Content-Type: application/json" --data-binary @example/bulk.json -s | jq .
```

The result of the above command is:

```text
{
  "delete_count": 5,
  "put_count": 2,
  "put_error_count": 0
}
```


### Search the documents from the Indigo gRPC Server via the Indigo REST Server

#### Simple query

```sh
$ curl -X POST "http://localhost:2289/api/example/_search" -H "Content-Type: application/json" --data-binary @example/simple_query.json -s | jq .
```

The result of the above `search documents` command is:

```json
{
  "searchResult": {
    "facets": {},
    "hits": [
      {
        "fields": {
          "category": "Library",
          "description": "Full-text search library written in Go.",
          "name": "Bleve",
          "popularity": 1,
          "release": "2014-04-18T00:00:00Z",
          "type": "document"
        },
        "id": "1",
        "index": "data/example",
        "score": 0.41589955606724527,
        "sort": [
          "_score"
        ]
      },
      {
        "fields": {
          "category": "Server",
          "description": "Full-text search server built on Bleve.",
          "name": "Indigo",
          "popularity": 5,
          "release": "2017-01-13T00:00:00Z",
          "type": "document"
        },
        "id": "7",
        "index": "data/example",
        "score": 0.41589955606724527,
        "sort": [
          "_score"
        ]
      }
    ],
    "max_score": 0.41589955606724527,
    "request": {
      "explain": false,
      "facets": null,
      "fields": [
        "name",
        "description",
        "category",
        "popularity",
        "release",
        "type"
      ],
      "from": 0,
      "highlight": null,
      "includeLocations": false,
      "query": {
        "query": "description:*"
      },
      "size": 10,
      "sort": [
        "-_score"
      ]
    },
    "status": {
      "failed": 0,
      "successful": 1,
      "total": 1
    },
    "took": 7488585,
    "total_hits": 2
  }
}
```






## The index mapping

The index mapping describes how to your data model should be indexed. See following example.

#### index_mapping.json

The index_mapping.json file contains all of the details about which fields your documents can contain, and how those fields should be dealt with when adding documents to the index, or when querying those fields.
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


## License

Apache License Version 2.0
