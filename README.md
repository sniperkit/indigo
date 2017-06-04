# Indigo

Indigo is a full text search and indexing server written in [Go](https://golang.org) based on [Bleve](http://www.blevesearch.com), it also includes a web server that provides a [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) interface and control command. Indigo makes it easy for programmers to develop search applications with advanced features.  

The Indigo gRPC Server provides full text search and indexing functions through [gRPC](http://www.grpc.io) ([HTTP/2](https://en.wikipedia.org/wiki/HTTP/2) + [Protocol Buffers](https://developers.google.com/protocol-buffers/)).  
The Indigo REST Server is a gateway, it provides a traditional JSON API ([HTTP/1.1](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) + [JSON](http://www.json.org)) that communicating with the Indigo gRPC Server.  


## Features

- Full-text search and indexing
- Faceting
- Result highlighting
- Text analysis
- Multiple indices

For more detailed information, refer to the [Bleve document](http://www.blevesearch.com/docs/Home/).

## Parameters and  Environment variables

Indigo parameters are described in indigo.yaml.

```yaml
log_format: text
log_output: ""
log_level: "info"

port: 1289
path: "/var/indigo/data"

index_mapping: ""
index_type: "upside_down"
kvstore: "boltdb"
kvconfig: ""

delete_index_at_startup: false
delete_index_at_shutdown: false
```

| Parameter name | Environment variable | Command line option | Description |
| --- | --- | --- | --- |
| log_format               | INDIGO_LOG_FORMAT               | --log-format | log format. `text`, `color` and `json` are available. Default is `text` |
| log_output               | INDIGO_LOG_OUTPUT               | --log-output | log output path. Default is `stdout` |
| log_level                | INDIGO_LOG_LEVEL                | --log-level | log level. `debug`, `info`, `warn`, `error`, `fatal` and `panic` are available. Default is `info` |
| port                     | INDIGO_PORT                     | --port | port number. default is `1289` |
| path                     | INDIGO_PATH                     | --path | index directory path. Default is `/var/indigo/data/index` |
| index_mapping            | INDIGO_INDEX_MAPPING            | --index-mapping |index mapping path. Default is `""` |
| index_type               | INDIGO_INDEX_TYPE               | --index-type | index type. `upside_down` is available. Default is `upside_down` |
| kvstore                  | INDIGO_KVSTORE                  | --kvstore | kvstore. `boltdb`, `goleveldb`, `gtreap` and `moss` are available. Default is `boltdb` |
| kvconfig                 | INDIGO_KVCONFIG                 | --kvconfig | kvconfig path. Default is `""` |
| delete_index_at_startup  | INDIGO_DELETE_INDEX_AT_STARTUP  | --delete-index-at-startup | delete index at startup. Default is `false` |
| delete_index_at_shutdown | INDIGO_DELETE_INDEX_AT_SHUTDOWN | --delete-index-at-shutdown | delete index at shutdown. Default is `false` |



## Start Indigo Server

The `start` command starts Indigo gRPC Server. You can display a help message by specifying `-h` or `--help` option.

```sh
$ indigo start
```


### Index Mapping

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


## Get the index information from the Indigo Server

The `get index` command retrieves an index information about existing opened index. You can display a help message by specifying the `- h` or` --help` option.

```sh
$ indigoctl get index
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


## Put the document to the Indigo Server

The `put document` command adds or updates a JSON formatted document in a specified index. You can display a help message by specifying the `- h` or` --help` option.  
The document example is following:

```json
{
  "id": "1",
  "fields": {
    "name": "Bleve",
    "description": "Bleve is a full-text search and indexing library for Go.",
    "category": "Library",
    "popularity": 3.0,
    "release": "2014-04-18T00:00:00Z",
    "type": "document"
  }
}
```

```sh
$ indigoctl put document --resource ../example/document_1.json
```

The result of the above `put document` command is:

```json
{
  "put_count": 1
}
```


## Get the document from the Indigo Server

The `get document` command retrieves a JSON formatted document on its id from a specified index. You can display a help message by specifying the `- h` or` --help` option.

```sh
$ indigoctl get document --id 1
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
$ indigoctl delete document --id 1
```

The result of the above `delete document` command is:

```json
{
  "delete_count": 1
}
```


## Index the documents in bulk to the Indigo Server

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
$ indigoctl bulk --resource ../example/bulk_put.json
```

The result of the above `bulk` command is:

```json
{
  "put_count": 7
}
```


## Search the documents from the Indigo Server

The `search` command can be executed with a search request, which includes the Query, within its file. Here is an example:

```json
{
  "query": {
    "query": "name:*"
  },
  "size": 10,
  "from": 0,
  "fields": [
    "*"
  ],
  "sort": [
    "-_score"
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

See [Queries](http://www.blevesearch.com/docs/Query/), [Query String Query](http://www.blevesearch.com/docs/Query-String-Query/) and [type SearchRequest](https://godoc.org/github.com/blevesearch/bleve#SearchRequest) for more details.

You can display a help message by specifying the `- h` or` --help` option.

```sh
$ indigoctl search --resource ../example/search_request.json
```

The result of the above `search` command is:

```json
{
  "search_result": {
    "facets": {
      "Category count": {
        "field": "category",
        "missing": 0,
        "other": 0,
        "terms": [
          {
            "count": 4,
            "term": "Library"
          },
          {
            "count": 3,
            "term": "Server"
          }
        ],
        "total": 7
      },
      "Popularity range": {
        "field": "popularity",
        "missing": 0,
        "numeric_ranges": [
          {
            "count": 2,
            "max": 4,
            "min": 3,
            "name": "more than or equal to 3 and less than 4"
          },
          {
            "count": 2,
            "min": 5,
            "name": "more than or equal to 5"
          },
          {
            "count": 1,
            "max": 2,
            "min": 1,
            "name": "more than or equal to 1 and less than 2"
          },
          {
            "count": 1,
            "max": 3,
            "min": 2,
            "name": "more than or equal to 2 and less than 3"
          },
          {
            "count": 1,
            "max": 5,
            "min": 4,
            "name": "more than or equal to 4 and less than 5"
          }
        ],
        "other": 0,
        "total": 7
      },
      "Release date range": {
        "date_ranges": [
          {
            "count": 4,
            "end": "2010-12-31T23:59:59Z",
            "name": "2001 - 2010",
            "start": "2001-01-01T00:00:00Z"
          },
          {
            "count": 2,
            "end": "2020-12-31T23:59:59Z",
            "name": "2011 - 2020",
            "start": "2011-01-01T00:00:00Z"
          }
        ],
        "field": "release",
        "missing": 0,
        "other": 0,
        "total": 6
      }
    },
    "hits": [
      {
        "fields": {
          "category": "Library",
          "description": "Bleve is a full-text search and indexing library for Go.",
          "name": "Bleve",
          "popularity": 3,
          "release": "2014-04-18T00:00:00Z",
          "type": "document"
        },
        "fragments": {
          "description": [
            "Bleve is a full-text search and indexing library for Go."
          ],
          "name": [
            "\u003cmark\u003eBleve\u003c/mark\u003e"
          ]
        },
        "id": "1",
        "index": "/var/indigo/data/index",
        "locations": {
          "name": {
            "bleve": [
              {
                "array_positions": null,
                "end": 5,
                "pos": 1,
                "start": 0
              }
            ]
          }
        },
        "score": 0.12163776688600772,
        "sort": [
          "_score"
        ]
      },
      {
        "fields": {
          "category": "Library",
          "description": "Apache Lucene is a high-performance, full-featured text search engine library written entirely in Java.",
          "name": "Lucene",
          "popularity": 4,
          "release": "2000-03-30T00:00:00Z",
          "type": "document"
        },
        "fragments": {
          "description": [
            "Apache Lucene is a high-performance, full-featured text search engine library written entirely in Java."
          ],
          "name": [
            "\u003cmark\u003eLucene\u003c/mark\u003e"
          ]
        },
        "id": "2",
        "index": "/var/indigo/data/index",
        "locations": {
          "name": {
            "lucen": [
              {
                "array_positions": null,
                "end": 6,
                "pos": 1,
                "start": 0
              }
            ]
          }
        },
        "score": 0.12163776688600772,
        "sort": [
          "_score"
        ]
      },
      {
        "fields": {
          "category": "Library",
          "description": "Whoosh is a fast, featureful full-text indexing and searching library implemented in pure Python.",
          "name": "Whoosh",
          "popularity": 3,
          "release": "2008-02-20T00:00:00Z",
          "type": "document"
        },
        "fragments": {
          "description": [
            "Whoosh is a fast, featureful full-text indexing and searching library implemented in pure Python."
          ],
          "name": [
            "\u003cmark\u003eWhoosh\u003c/mark\u003e"
          ]
        },
        "id": "3",
        "index": "/var/indigo/data/index",
        "locations": {
          "name": {
            "whoosh": [
              {
                "array_positions": null,
                "end": 6,
                "pos": 1,
                "start": 0
              }
            ]
          }
        },
        "score": 0.12163776688600772,
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
        "fragments": {
          "description": [
            "Ferret is a super fast, highly configurable search library written in Ruby."
          ],
          "name": [
            "\u003cmark\u003eFerret\u003c/mark\u003e"
          ]
        },
        "id": "4",
        "index": "/var/indigo/data/index",
        "locations": {
          "name": {
            "ferret": [
              {
                "array_positions": null,
                "end": 6,
                "pos": 1,
                "start": 0
              }
            ]
          }
        },
        "score": 0.12163776688600772,
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
        "fragments": {
          "description": [
            "Solr is an open source enterprise search platform, written in Java, from the Apache Lucene project."
          ],
          "name": [
            "\u003cmark\u003eSolr\u003c/mark\u003e"
          ]
        },
        "id": "5",
        "index": "/var/indigo/data/index",
        "locations": {
          "name": {
            "solr": [
              {
                "array_positions": null,
                "end": 4,
                "pos": 1,
                "start": 0
              }
            ]
          }
        },
        "score": 0.12163776688600772,
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
        "fragments": {
          "description": [
            "Elasticsearch is a search engine based on Lucene, written in Java."
          ],
          "name": [
            "\u003cmark\u003eElasticsearch\u003c/mark\u003e"
          ]
        },
        "id": "6",
        "index": "/var/indigo/data/index",
        "locations": {
          "name": {
            "elasticsearch": [
              {
                "array_positions": null,
                "end": 13,
                "pos": 1,
                "start": 0
              }
            ]
          }
        },
        "score": 0.12163776688600772,
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
        "fragments": {
          "description": [
            "Indigo is a full-text search and indexing server written in Go, built on top of Bleve."
          ],
          "name": [
            "\u003cmark\u003eIndigo\u003c/mark\u003e"
          ]
        },
        "id": "7",
        "index": "/var/indigo/data/index",
        "locations": {
          "name": {
            "indigo": [
              {
                "array_positions": null,
                "end": 6,
                "pos": 1,
                "start": 0
              }
            ]
          }
        },
        "score": 0.12163776688600772,
        "sort": [
          "_score"
        ]
      }
    ],
    "max_score": 0.12163776688600772,
    "request": {
      "explain": false,
      "facets": {
        "Category count": {
          "field": "category",
          "size": 10
        },
        "Popularity range": {
          "field": "popularity",
          "numeric_ranges": [
            {
              "max": 1,
              "name": "less than 1"
            },
            {
              "max": 2,
              "min": 1,
              "name": "more than or equal to 1 and less than 2"
            },
            {
              "max": 3,
              "min": 2,
              "name": "more than or equal to 2 and less than 3"
            },
            {
              "max": 4,
              "min": 3,
              "name": "more than or equal to 3 and less than 4"
            },
            {
              "max": 5,
              "min": 4,
              "name": "more than or equal to 4 and less than 5"
            },
            {
              "min": 5,
              "name": "more than or equal to 5"
            }
          ],
          "size": 10
        },
        "Release date range": {
          "date_ranges": [
            {
              "end": "2010-12-31T23:59:59Z",
              "name": "2001 - 2010",
              "start": "2001-01-01T00:00:00Z"
            },
            {
              "end": "2020-12-31T23:59:59Z",
              "name": "2011 - 2020",
              "start": "2011-01-01T00:00:00Z"
            }
          ],
          "field": "release",
          "size": 10
        }
      },
      "fields": [
        "*"
      ],
      "from": 0,
      "highlight": {
        "fields": [
          "name",
          "description"
        ],
        "style": "html"
      },
      "includeLocations": false,
      "query": {
        "query": "name:*"
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
    "took": 4103742,
    "total_hits": 7
  }
}
```

See [type SearchResult](https://godoc.org/github.com/blevesearch/bleve#SearchResult) for more details.











## License

Apache License Version 2.0
