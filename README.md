# Indigo

The Indigo is a full text search and indexing server written in [Go](https://golang.org) based on [Bleve](http://www.blevesearch.com).  

Indigo includes a full text search and indexing server, it also includes a web server that provides a [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) interface for Indigo gRPC server.  

Indigo gRPC Server communicates with the client using [gRPC](http://www.grpc.io) ([HTTP/2](https://en.wikipedia.org/wiki/HTTP/2) + [Protocol Buffers](https://developers.google.com/protocol-buffers/)). You can access to the Indigo gRPC Server using gRPC directly.  

Indigo REST Server provides JSON API ([HTTP/1.1](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) + [JSON](http://www.json.org)) that communicate with Indigo gRPC Server. If you want to access to Indigo gRPC Server using traditional JSON API, you can access via Indigo REST Server.


## Features

- TODO


## The Indigo Command Line Interface

The Indigo provides some commands for controlling the Indigo Server.


### Start Indigo gRPC Server

The `start grpc` command starts the Indigo gRPC Server.

```sh
$ ./indigo start grpc -l trace -f color
```


### Create the index to the Indigo gRPC Server

```sh
$ ./indigo create index -n example -m example/index_mapping.json -s boltdb -t upside_down -f json
```

The result of the above `create index` command is:

```json
{
  "index_name": "example",
  "index_dir": "data/example"
}
```


### Open the index to the Indigo gRPC Server

```sh
$ ./indigo open index -n example -f json
```

The result of the above `open index` command is:

```json
{
  "index_name": "example",
  "index_dir": "data/example"
}
```


### Get the index information from the Indigo gRPC Server

```sh
$ ./indigo get index -n example -f json
```

The result of the above `get index` command is:

```json
{
  "document_count": 0,
  "index_stats": {
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
  },
  "index_mapping": {
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


### Close the index from the Indigo gRPC Server

```sh
$ ./indigo close index -n example -f json
```

The result of the above `close index` command is:

```json
{
  "index_name": "example"
}
```


### Delete the index from the Indigo gRPC Server

```sh
$ ./indigo delete index -n example -f json
```

The result of the above `delete index` command is:

```json
{
  "index_name": "example"
}
```


### List the indices from the Indigo gRPC Server

```sh
$ ./indigo list index -f json
```

The result of the above `list index` command is:

```json
{
  "indices": [
    "example"
  ]
}
```


### Put the document to the Indigo gRPC Server

```sh
$ ./indigo put document -n example -i 1 -F example/document_1.json -f json
```

The result of the above `put document` command is:

```json
{
  "success": true
}
```


### Get the document from the Indigo gRPC Server

```sh
$ ./indigo get document -n example -i 1 -f json
```

The result of the above `get document` command is:

```json
{
  "id": "1",
  "fields": {
    "category": "Library",
    "description": "Bleve is a full-text search and indexing library for Go.",
    "name": "Bleve",
    "popularity": 3,
    "release": "2014-04-18T00:00:00Z",
    "type": "document"
  }
}
```


### Delete the document from the Indigo gRPC Server

```sh
$ ./indigo delete document -n example -i 1 -f json
```

The result of the above `delete document` command is:

```json
{
  "success": true
}
```


### Index the documents in bulk to the Indigo gRPC Server

```sh
$ ./indigo bulk -n example -b example/bulk_put.json -f json
```

The result of the above `bulk` command is:

```json
{
  "put_count": 7,
  "put_error_count": 0,
  "delete_count": 0
}
```


### Search the documents frmo the Indigo gRPC Server

```sh
$ ./indigo search -n example -s example/simple_query.json -f json
```

The result of the above `search` command is:

```json
{
  "search_result": {
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
    "took": 7725035,
    "total_hits": 7
  }
}
```

See [type SearchResult](https://godoc.org/github.com/blevesearch/bleve#SearchResult) for more details.


### Start Indigo REST Server

The `start rest` command starts the Indigo REST Server.

```sh
$ indigo start rest
```

### Create the index to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -s -X PUT -H "Content-Type: application/json" --data-binary @example/index_mapping.json "http://localhost:2289/api/example?indexType=upside_down&indexStore=boltdb"
```

The result of the above command is:

```json
{
  "index_name": "example",
  "index_dir": "data/example"
}
```

### Open the index to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -s -H "Content-Type: application/json" --data-binary @example/runtime_config.json -X POST "http://localhost:2289/api/example/_open"
```

The result of the above command is:

```json
{
  "index_name": "example",
  "index_dir": "data/example"
}
```

### Close the index to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -s -X POST "http://localhost:2289/api/example/_close"
```

The result of the above command is:

```json
{
  "index_name": "example"
}
```

### Delete the index to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -s -X DELETE "http://localhost:2289/api/example"
```

The result of the above command is:

```json
{
  "index_name": "example"
}
```

### List the index to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -s -X GET "http://localhost:2289/api/_list"
```

The result of the above command is:

```json
{
  "indices": [
    "example"
  ]
}
```

### Get the index information from the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -s -X GET "http://localhost:2289/api/example"
```

The result of the above command is:

```json
{
  "document_count": 0,
  "index_stats": {
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
  },
  "index_mapping": {
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
$ curl -s -X PUT -H "Content-Type: application/json" --data-binary @example/document_1.json "http://localhost:2289/api/example/1"
```

The result of the above command is:

```json
{
  "success": true
}
```

### Get the document to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -s -X GET "http://localhost:2289/api/example/1"
```

The result of the above command is:

```json
{
  "id": "1",
  "fields": {
    "category": "Library",
    "description": "Bleve is a full-text search and indexing library for Go.",
    "name": "Bleve",
    "popularity": 3,
    "release": "2014-04-18T00:00:00Z",
    "type": "document"
  }
}
```

### Delete the document from the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -s -X DELETE "http://localhost:2289/api/example/1"
```

The result of the above command is:

```json
{
  "success": true
}
```

### Index the documents in bulk to the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -s -X POST -H "Content-Type: application/json" --data-binary @example/bulk_put.json "http://localhost:2289/api/example/_bulk"
```

The result of the above command is:

```text
{
  "put_count": 7
}
```


### Search the documents from the Indigo gRPC Server via the Indigo REST Server

```sh
$ curl -s -X POST -H "Content-Type: application/json" --data-binary @example/simple_query.json "http://localhost:2289/api/example/_search"
```

The result of the above command is:

```json
{
  "search_result": {
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
    "took": 38819452,
    "total_hits": 7
  }
}
```


## The index mapping

The index mapping describes how to your data model should be indexed. See following example.  
The index_mapping.json file contains all of the details about which fields your documents can contain, and how those fields should be dealt with when adding documents to the index, or when querying those fields.  
See [Introduction to Index Mappings](http://www.blevesearch.com/docs/Index-Mapping/) and [type IndexMappingImpl](https://godoc.org/github.com/blevesearch/bleve/mapping#IndexMappingImpl) for more details.  

### example

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

## document mapping

### example

```json
{
  "name": "Bleve",
  "description": "Bleve is a full-text search and indexing library for Go.",
  "category": "Library",
  "popularity": 3.0,
  "release": "2014-04-18T00:00:00Z",
  "type": "document"
}
```

## Bulk request format

### example

```json
[
  {
    "method" : "put",
    "id": "1",
    "fields": {
      "name": "Bleve",
      "description": "Bleve is a full-text search and indexing library for Go.",
      "category": "Library",
      "popularity": 3.0,
      "release": "2014-04-18T00:00:00Z",
      "type": "document"
    }
  },
  {
    "method" : "delete",
    "id": "2"
  },
  {
    "method" : "delete",
    "id": "3"
  },
  {
    "method" : "delete",
    "id": "4"
  },
  {
    "method" : "delete",
    "id": "5"
  },
  {
    "method" : "delete",
    "id": "6"
  },
  {
    "method" : "put",
    "id": "7",
    "fields": {
      "name": "Indigo",
      "description": "Indigo is a full-text search and indexing server written in Go, built on top of Bleve.",
      "category": "Server",
      "popularity": 1.0,
      "release": "2017-01-13T00:00:00Z",
      "type": "document"
    }
  }
]
```


## Search request format

See [Queries](http://www.blevesearch.com/docs/Query/), [Query String Query](http://www.blevesearch.com/docs/Query-String-Query/) and [type SearchRequest](https://godoc.org/github.com/blevesearch/bleve#SearchRequest) for more details.

### example

```json
{
  "query": {
    "query": "description:search"
  },
  "size": 10,
  "from": 0,
  "fields": [
    "name",
    "description",
    "category",
    "popularity",
    "release",
    "type"
  ],
  "facets": {
    "Category count": {
      "size": 10,
      "field": "category"
    },
    "Popularity range": {
      "size": 10,
      "field": "popularity",
      "numeric_ranges": [
        {
          "name": "less than 1",
          "max": 1
        },
        {
          "name": "more than or equal to 1 and less than 2",
          "min": 1,
          "max": 2
        },
        {
          "name": "more than or equal to 2 and less than 3",
          "min": 2,
          "max": 3
        },
        {
          "name": "more than or equal to 3 and less than 4",
          "min": 3,
          "max": 4
        },
        {
          "name": "more than or equal to 4 and less than 5",
          "min": 4,
          "max": 5
        },
        {
          "name": "more than or equal to 5",
          "min": 5
        }
      ]
    },
    "Release date range": {
      "size": 10,
      "field": "release",
      "date_ranges": [
        {
          "name": "2001 - 2010",
          "start": "2001-01-01T00:00:00Z",
          "end": "2010-12-31T23:59:59Z"
        },
        {
          "name": "2011 - 2020",
          "start": "2011-01-01T00:00:00Z",
          "end": "2020-12-31T23:59:59Z"
        }
      ]
    }
  },
  "highlight": {
    "style": "html",
    "fields": [
      "name",
      "description"
    ]
  }
}
```

## License

Apache License Version 2.0
