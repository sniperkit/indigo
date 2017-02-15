# Indigo

The Indigo is an index server written in [Go](https://golang.org), built on top of the [Bleve](http://www.blevesearch.com).


## Indigo gRPC Server

*WIP:*

### Start Indigo gRPC Server

The `indigo start grpc` command starts the Indigo gRPC Server.

```
$ indigo start grpc
```

### mapping.json

The mapping.json file contains all of the details about which fields your documents can contain, and how those fields should be dealt with when adding documents to the index, or when querying those fields.
See [Introduction to Index Mappings](http://www.blevesearch.com/docs/Index-Mapping/) and [type IndexMappingImpl](https://godoc.org/github.com/blevesearch/bleve/mapping#IndexMappingImpl) for more details.

```
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


## Indigo Client

*WIP:*

### Create index

```
$ ./indigo/indigo create index example "$(cat mapping.json)" -s boltdb -t upside_down
```


### Get index mapping

```
$ ./indigo/indigo get mapping example | jq .
$ indigo client mapping | jq .
```

The result of the above mapping command is:

```
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

### Index documents

```
$ indigo client index '{
  "1": {
    "name": "Bleve",
    "description": "Full-text search library written in Go.",
    "category": "Library",
    "popularity": 1.0,
    "release": "2014-04-18T00:00:00Z",
    "type": "document"
  },
  "2": {
    "name": "Lucene",
    "description": "Full-text search library written in Java.",
    "category": "Library",
    "popularity": 2.5,
    "release": "2000-03-30T00:00:00Z",
    "type": "document"
  },
  "3": {
    "name": "Solr",
    "description": "Full-text search server built on Lucene.",
    "category": "Server",
    "popularity": 4.5,
    "release": "2006-12-22T00:00:00Z",
    "type": "document"
  },
  "4": {
    "name": "Elasticsearch",
    "description": "Full-text search server built on Lucene.",
    "category": "Server",
    "popularity": 5.0,
    "release": "2010-02-08T00:00:00Z",
    "type": "document"
  },
  "5": {
    "name": "Indigo",
    "description": "Full-text search server built on Bleve.",
    "category": "Server",
    "popularity": 5.0,
    "release": "2017-01-13T00:00:00Z",
    "type": "document"
  }
}' | jq .
```


The result of the above index command is:

```
{
  "document_count": 5
}
```


### Delete documents

```
$ indigo client index -d '[
  "2",
  "3",
  "4"
]' | jq .
```

The result of the above delete command is:

```
{
  "document_count": 3
}
```


### Search documents

#### Simple query

```
$ ./indigo/indigo client search '{
  "query": {
    "query": "description:Go"
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
}' | jq .
```

The result of the above simple query command is:

```
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


## Indigo REST Server

The Indigo REST Server provides a RESTful JSON API into Indigo gRPC Server.

### Start Indigo REST Server

The `indigo_rest start` command starts the Indigo REST Server.

```
$ indigo_rest start
```

#### indigo.yaml

The indigo.yaml file is the configuration file with the parameters affecting the Indigo REST Server itself.

```
#
# Indigo gRPC Server configuration
#
grpc:
    server:
        name: localhost
        port: 10000

#
# Indigo REST Server configuration
#
rest:
    log:
        file: ./indigo_rest.log
        level: info
        format: text
    server:
        name: localhost
        port: 20000
        base_path: /api
```

#### Configuration parameters

| Name             | Description   |
| ---------------- | ------------- |
| rest.log.file         | Log file path |
| rest.log.level        | Log level. You can choose `trace`, `debug`, `info`, `warn`, `error`, `alert`. default is `info` |
| rest.log.format       | Log format. You can choose `text` or `json`. default is `text` |
| rest.server.name      | Server name |
| rest.server.port      | Server port |
| rest.server.base_path   | Server URI path |
| grpc.server.name | Indigo gRPC Server name |
| grpc.server.port | Indigo gRPC Server port |


### Get mapping via Indigo REST Server

```
$ curl -X GET -s http://localhost:20000/api/mapping | jq .
```

The result of the above mapping command is:

```
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


### Index documents via Indigo REST Server

```
$ curl -X POST -s http://localhost:20000/api/index -d '{
  "1": {
    "name": "Bleve",
    "description": "Full-text search library written in Go.",
    "category": "Library",
    "popularity": 1.0,
    "release": "2014-04-18T00:00:00Z",
    "type": "document"
  },
  "2": {
    "name": "Lucene",
    "description": "Full-text search library written in Java.",
    "category": "Library",
    "popularity": 2.5,
    "release": "2000-03-30T00:00:00Z",
    "type": "document"
  },
  "3": {
    "name": "Solr",
    "description": "Full-text search server built on Lucene.",
    "category": "Server",
    "popularity": 4.5,
    "release": "2006-12-22T00:00:00Z",
    "type": "document"
  },
  "4": {
    "name": "Elasticsearch",
    "description": "Full-text search server built on Lucene.",
    "category": "Server",
    "popularity": 5.0,
    "release": "2010-02-08T00:00:00Z",
    "type": "document"
  },
  "5": {
    "name": "Indigo",
    "description": "Full-text search server built on Bleve.",
    "category": "Server",
    "popularity": 5.0,
    "release": "2017-01-13T00:00:00Z",
    "type": "document"
  }
}' | jq .
```

The result of the above index command is:

```
{
  "document_count": 5
}
```

### Delete documents via Indigo REST Server

```
$ curl -X DELETE -s http://localhost:20000/api/index -d '[
  "2",
  "3",
  "4"
]' | jq .
```

The result of the above delete command is:

```
{
  "document_count": 3
}
```


### Search documents via Indigo REST Server

The `search` command queries the documents to specified Indigo Server.
See [Queries](http://www.blevesearch.com/docs/Query/), [Query String Query](http://www.blevesearch.com/docs/Query-String-Query/) and [type SearchRequest](https://godoc.org/github.com/blevesearch/bleve#SearchRequest) for more details.

#### Simple query

```
$ curl -X POST -s http://localhost:20000/api/search -d '{
  "query": {
    "query": "description:Go"
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
}' | jq .
```

The result of the above simple query command is:

```
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

### Faceted query

```
$ curl -X POST -s http://localhost:20000/api/search -d '{
  "query": {
    "query": "description:Go"
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
  }
}' | jq .
```

The result of the above faceted query command is:

```
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
            "end": "2010-12-31T23:59:59Z",
            "name": "2001 - 2010",
            "start": "2001-01-01T00:00:00Z"
          },
          {
            "end": "2020-12-31T23:59:59Z",
            "name": "2011 - 2020",
            "start": "2011-01-01T00:00:00Z"
          }
        ]
      }
    },
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
  "took": 1185465,
  "facets": {
    "Category count": {
      "field": "category",
      "total": 1,
      "missing": 0,
      "other": 0,
      "terms": [
        {
          "term": "Library",
          "count": 1
        }
      ]
    },
    "Popularity range": {
      "field": "popularity",
      "total": 1,
      "missing": 0,
      "other": 0,
      "numeric_ranges": [
        {
          "name": "more than or equal to 1 and less than 2",
          "min": 1,
          "max": 2,
          "count": 1
        }
      ]
    },
    "Release date range": {
      "field": "release",
      "total": 1,
      "missing": 0,
      "other": 0,
      "date_ranges": [
        {
          "name": "2011 - 2020",
          "start": "2011-01-01T00:00:00Z",
          "end": "2020-12-31T23:59:59Z",
          "count": 1
        }
      ]
    }
  }
}
```

### Highlighted query

```
$ curl -X POST -s http://localhost:20000/api/search -d '{
  "query": {
    "query": "description:Go"
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
  "highlight": {
    "style": "html",
    "fields": [
      "name",
      "description"
    ]
  }
}' | jq .
```

The result of the above highlighted query command is:

```
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
    "highlight": {
      "style": "html",
      "fields": [
        "name",
        "description"
      ]
    },
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
      "locations": {
        "description": {
          "go": [
            {
              "pos": 7,
              "start": 36,
              "end": 38,
              "array_positions": null
            }
          ]
        }
      },
      "fragments": {
        "description": [
          "Full-text search library written in <mark>Go</mark>."
        ],
        "name": [
          "Bleve"
        ]
      },
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
  "took": 708885,
  "facets": {}
}
```

### Faceted and highlighted query

```
$ curl -X POST -s http://localhost:20000/api/search -d '{
  "query": {
    "query": "description:Go"
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
}' | jq .
```

The result of the above faceted and highlighted query command is:

```
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
    "highlight": {
      "style": "html",
      "fields": [
        "name",
        "description"
      ]
    },
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
            "end": "2010-12-31T23:59:59Z",
            "name": "2001 - 2010",
            "start": "2001-01-01T00:00:00Z"
          },
          {
            "end": "2020-12-31T23:59:59Z",
            "name": "2011 - 2020",
            "start": "2011-01-01T00:00:00Z"
          }
        ]
      }
    },
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
      "locations": {
        "description": {
          "go": [
            {
              "pos": 7,
              "start": 36,
              "end": 38,
              "array_positions": null
            }
          ]
        }
      },
      "fragments": {
        "description": [
          "Full-text search library written in <mark>Go</mark>."
        ],
        "name": [
          "Bleve"
        ]
      },
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
  "took": 626486,
  "facets": {
    "Category count": {
      "field": "category",
      "total": 1,
      "missing": 0,
      "other": 0,
      "terms": [
        {
          "term": "Library",
          "count": 1
        }
      ]
    },
    "Popularity range": {
      "field": "popularity",
      "total": 1,
      "missing": 0,
      "other": 0,
      "numeric_ranges": [
        {
          "name": "more than or equal to 1 and less than 2",
          "min": 1,
          "max": 2,
          "count": 1
        }
      ]
    },
    "Release date range": {
      "field": "release",
      "total": 1,
      "missing": 0,
      "other": 0,
      "date_ranges": [
        {
          "name": "2011 - 2020",
          "start": "2011-01-01T00:00:00Z",
          "end": "2020-12-31T23:59:59Z",
          "count": 1
        }
      ]
    }
  }
}
```

See [type SearchResult](https://godoc.org/github.com/blevesearch/bleve#SearchResult) for more details.



## License

Apache License Version 2.0
