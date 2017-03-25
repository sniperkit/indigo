# Indigo

Indigo is a full text search and indexing server written in [Go](https://golang.org) based on [Bleve](http://www.blevesearch.com), it also includes a web server that provides a [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) interface and control command. Indigo makes it easy for programmers to develop search applications with advanced features.  

The Indigo gRPC Server provides full text search and indexing functions through [gRPC](http://www.grpc.io) ([HTTP/2](https://en.wikipedia.org/wiki/HTTP/2) + [Protocol Buffers](https://developers.google.com/protocol-buffers/)).  
The Indigo REST Server is a gateway, it provides a traditional JSON API ([HTTP/1.1](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) + [JSON](http://www.json.org)) that communicating with the Indigo gRPC Server.  

![](./img/Indigo%20Architecture.png "Indigo")


## Features

- Full-text search and indexing
- Faceting
- Result highlighting
- Text analysis
- Multiple indices

For more detailed information, refer to the [Bleve document](http://www.blevesearch.com/docs/Home/).

## Configuration file

Indigo has a YAML format configuration file.

```yaml
log_output_format: text
log_output: ""
log_level: "info"

grpc:
  port: 1289
  data_dir: "/var/indigo/data"
  open_existing_index: true

rest:
  port: 2289
  base_uri: "/api"
  grpc_server: "localhost:1289"
```

| Name                     | Description                                                                                       |
| ------------------------ | ------------------------------------------------------------------------------------------------- |
| log_output_format        | The log output format of the Indigo gRPC or REST Server. Default is `text`                        |
| log_output               | The log output destination of the Indigo gRPC or REST Server. Default is `stdout`                 |
| log_level                | The log level of log output by Indigo gRPC or REST Server. Default is `info`                      |
| grpc.port                | Port number to be used when the Indigo gRPC Server starts up. default is `1289`                   |
| grpc.data_dir            | The path of the directory where Indigo gRPC Server stores the data. Default is `/var/indigo/data` |
| grpc.open_existing_index | Flag to open indices when started to Indigo gRPC Server. Default is `false`                       |
| rest.port                | Port number to be used when the Indigo REST Server starts up. default is `2289`                   |
| rest.base_uri            | The base URI of API endpoint on the Indigo REST Server. Default is `/api`                         |
| rest.grpc_server         | Indigo gRPC server to which Indigo REST Server connects. Default is `localhost:1289`              |


## Environment variables

The Indigo supports following environment variables.

| Name                            | Description                                                                                       |
| ------------------------------- | ------------------------------------------------------------------------------------------------- |
| INDIGO_LOG_OUTPUT_FORMAT        | The log output format of the Indigo gRPC or REST Server. Default is `text`                        |
| INDIGO_LOG_OUTPUT               | The log output destination of the Indigo gRPC or REST Server. Default is `stdout`                 |
| INDIGO_LOG_LEVEL                | The log level of log output by Indigo gRPC or REST Server. Default is `info`                      |
| INDIGO_GRPC_PORT                | Port number to be used when the Indigo gRPC Server starts up. default is `1289`                   |
| INDIGO_GRPC_DATA_DIR            | The path of the directory where Indigo gRPC Server stores the data. Default is `/var/indigo/data` |
| INDIGO_GRPC_OPEN_EXISTING_INDEX | Flag to open indices when started to Indigo gRPC Server. Default is `false`                       |
| INDIGO_REST_PORT                | Port number to be used when the Indigo REST Server starts up. default is `2289`                   |
| INDIGO_REST_BASE_URI            | The base URI of API endpoint on the Indigo REST Server. Default is `/api`                         |
| INDIGO_REST_GRPC_SERVER         | Indigo gRPC server to which Indigo REST Server connects. Default is `localhost:1289`              |


## Start Indigo gRPC Server

The `start grpc` command starts Indigo gRPC Server. You can display a help message by specifying `-h` or `--help` option.

```sh
$ indigo start grpc
```


## Create the index to the Indigo gRPC Server

The `create index` command creates the Index to the Indigo gRPC Server. Indigo provides support for multiple indices, including executing operations across several indices. You can display a help message by specifying `-h` or `--help` option.  
You can specify the index mapping describes how to your data model should be indexed. it contains all of the details about which fields your documents can contain, and how those fields should be dealt with when adding documents to the index, or when querying those fields. The example is following:

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

See [Introduction to Index Mappings](http://www.blevesearch.com/docs/Index-Mapping/) and [type IndexMappingImpl](https://godoc.org/github.com/blevesearch/bleve/mapping#IndexMappingImpl) for more details.  

```sh
$ indigoctl create index -i example example/index_mapping.json -f json
```

The result of the above `create index` command is:

```json
{
  "index_name": "example",
  "index_dir": "data/example"
}
```


## Open the index to the Indigo gRPC Server

The `open index` command opens an existing closed index. You can display a help message by specifying the `- h` or` --help` option.

```sh
$ indigoctl open index -i example -f json
```

The result of the above `open index` command is:

```json
{
  "index_name": "example",
  "index_dir": "data/example"
}
```


## Get the index information from the Indigo gRPC Server

The `get index` command retrieves an index information about existing opened index. You can display a help message by specifying the `- h` or` --help` option.

```sh
$ indigoctl get index -i example -f json
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


## Close the index from the Indigo gRPC Server

The `close index` command closes an existing opened index. You can display a help message by specifying the `- h` or` --help` option.

```sh
$ indigoctl close index -i example -f json
```

The result of the above `close index` command is:

```json
{
  "index_name": "example"
}
```


## Delete the index from the Indigo gRPC Server

The `delete index` command deletes an existing closed index. You can display a help message by specifying the `- h` or` --help` option.

```sh
$ indigoctl delete index -i example -f json
```

The result of the above `delete index` command is:

```json
{
  "index_name": "example"
}
```


## List the indices from the Indigo gRPC Server

The `list index` command lists opened indices. You can display a help message by specifying the `- h` or` --help` option.

```sh
$ indigoctl list index -f json
```

The result of the above `list index` command is:

```json
{
  "indices": [
    "example"
  ]
}
```


## Put the document to the Indigo gRPC Server

The `put document` command adds or updates a JSON formatted document in a specified index. You can display a help message by specifying the `- h` or` --help` option.  
The document example is following:

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

```sh
$ indigoctl put document -i example -d 1 -F example/document_1.json -f json
```

The result of the above `put document` command is:

```json
{
  "success": true
}
```


## Get the document from the Indigo gRPC Server

The `get document` command retrieves a JSON formatted document on its id from a specified index. You can display a help message by specifying the `- h` or` --help` option.

```sh
$ indigoctl get document -i example -d 1 -f json
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


## Delete the document from the Indigo gRPC Server

The `delete document` command deletes a document on its id from a specified index. You can display a help message by specifying the `- h` or` --help` option.

```sh
$ indigoctl delete document -i example -d 1 -f json
```

The result of the above `delete document` command is:

```json
{
  "success": true
}
```


## Index the documents in bulk to the Indigo gRPC Server

The `bulk` command makes it possible to perform many put/delete operations in a single command execution. This can greatly increase the indexing speed. You can display a help message by specifying the `- h` or` --help` option.
The bulk example is following:

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

```sh
$ indigoctl bulk -i example -b example/bulk_put.json -f json
```

The result of the above `bulk` command is:

```json
{
  "put_count": 7,
}
```


## Search the documents from the Indigo gRPC Server

The `search` command can be executed with a search request, which includes the Query, within its file. Here is an example:

```json
{
  "query": {
    "query": "description:*"
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
  ]
}
```

See [Queries](http://www.blevesearch.com/docs/Query/), [Query String Query](http://www.blevesearch.com/docs/Query-String-Query/) and [type SearchRequest](https://godoc.org/github.com/blevesearch/bleve#SearchRequest) for more details.

You can display a help message by specifying the `- h` or` --help` option.

```sh
$ indigoctl search -i example -s example/simple_query.json -f json
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








## Start Indigo REST Server

The `start rest` command starts Indigo REST Server. You can display a help message by specifying `-h` or `--help` option.

```sh
$ indigo start rest
```


## Create the index to the Indigo gRPC Server via the Indigo REST Server

The create index API creates the Index to the Indigo gRPC Server. Indigo provides support for multiple indices, including executing operations across several indices.

```sh
$ curl -s -X PUT -H "Content-Type: application/json" --data-binary @example/index_mapping.json "http://localhost:2289/api/example"
```

The result of the above command is:

```json
{
  "index_name": "example",
  "index_dir": "data/example"
}
```


## Open the index to the Indigo gRPC Server via the Indigo REST Server

The open index API opens an existing closed index.

```sh
$ curl -s -X POST -H "Content-Type: application/json" --data-binary @example/runtime_config.json "http://localhost:2289/api/example/_open"
```

The result of the above command is:

```json
{
  "index_name": "example",
  "index_dir": "data/example"
}
```


## Get the index information from the Indigo gRPC Server via the Indigo REST Server

The get index API retrieves an index information about existing opened index.

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


## Close the index to the Indigo gRPC Server via the Indigo REST Server

The close index API closes an existing opened index.

```sh
$ curl -s -X POST "http://localhost:2289/api/example/_close"
```

The result of the above command is:

```json
{
  "index_name": "example"
}
```


## Delete the index to the Indigo gRPC Server via the Indigo REST Server

The delete index API deletes an existing closed index.

```sh
$ curl -s -X DELETE "http://localhost:2289/api/example"
```

The result of the above command is:

```json
{
  "index_name": "example"
}
```


## List the index to the Indigo gRPC Server via the Indigo REST Server

The list index API lists opened indices.

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


## Put the document to the Indigo gRPC Server via the Indigo REST Server

The put document API adds or updates a JSON formatted document in a specified index.

```sh
$ curl -s -X PUT -H "Content-Type: application/json" --data-binary @example/document_1.json "http://localhost:2289/api/example/1"
```

The result of the above command is:

```json
{
  "success": true
}
```


## Get the document to the Indigo gRPC Server via the Indigo REST Server

The get document API retrieves a JSON formatted document on its id from a specified index.

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


## Delete the document from the Indigo gRPC Server via the Indigo REST Server

The delete document API deletes a document on its id from a specified index.

```sh
$ curl -s -X DELETE "http://localhost:2289/api/example/1"
```

The result of the above command is:

```json
{
  "success": true
}
```


## Index the documents in bulk to the Indigo gRPC Server via the Indigo REST Server

The bulk API makes it possible to perform many put/delete operations in a single command execution. This can greatly increase the indexing speed.

```sh
$ curl -s -X POST -H "Content-Type: application/json" --data-binary @example/bulk_put.json "http://localhost:2289/api/example/_bulk"
```

The result of the above command is:

```text
{
  "put_count": 7
}
```


## Search the documents from the Indigo gRPC Server via the Indigo REST Server

The search API can be executed with a search request, which includes the Query, within its file.

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


## License

Apache License Version 2.0
