# Indigo gRPC Server

Indigo gRPC Server is a search and indexing server built on top of [Bleve](http://www.blevesearch.com).


## Start Indigo gRPC Server

The `indigo_grpc start` command starts the Indigo gRPC Server.

```
$ indigo_grpc start
```

### indigo_grpc.yaml

The indigo_grpc.yaml file is the configuration file with the parameters affecting Indigo gRPC Server itself.

```
#
# Logging configuration
#
log:
    file: ./indigo_grpc.log
    level: info
    format: text

#
# Indigo gRPC Server configuration
#
server:
    name: localhost
    port: 10000

#
# Index configuration
#
index:
    dir: ./index
    type: "upside_down"
    store: "boltdb"
    mapping: ./mapping.json
```

#### Configuration parameters

| Name          | Description   |
| ------------- | ------------- |
| log.file      | Log file path |
| log.level     | Log level. You can choose `trace`, `debug`, `info`, `warn`, `error`, `alert`. default is `info` |
| log.format    | Log format. You can choose `text` or `json`. default is `text` |
| server.name   | Server name |
| server.port   | Server port |
| index.dir     | Creates index at the specified path, if not exist |
| index.type    | Index type. You can choose `smolder` or `upside_down`. Default is `upside_down` |
| index.store   | Index store. You can choose `boltdb`, `goleveldb`, `gtreap`, `metrics` or `moss`. Default is `boltdb` |
| index.mapping | IndexMapping file path |


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

## Start Indigo REST Server

The `indigo_rest start` command starts the Indigo REST Server.

```
$ indigo_rest start
```

### indigo_rest.yaml

The indigo_rest.yaml file is the configuration file with the parameters affecting Indigo REST Server itself.

```
#
# Logging configuration
#
log:
    file: ./indigo_rest.log
    level: info
    format: text

#
# Indigo REST Server configuration
#
server:
    name: localhost
    port: 20000
    uripath: /api

#
# Indigo gRPC Server configuration
#
grpc:
    server:
        name: localhost
        port: 10000
```

#### Configuration parameters

| Name             | Description   |
| ---------------- | ------------- |
| log.file         | Log file path |
| log.level        | Log level. You can choose `trace`, `debug`, `info`, `warn`, `error`, `alert`. default is `info` |
| log.format       | Log format. You can choose `text` or `json`. default is `text` |
| server.name      | Server name |
| server.port      | Server port |
| server.uripath   | Server URI path |
| grpc.server.name | |
| grpc.server.port | |


### Get mapping via HTTP

```
$ curl -XGET http://localhost:20000/api/mapping | jq .
```

The result of the above mapping command is:

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1392  100  1392    0     0   176k      0 --:--:-- --:--:-- --:--:--  194k
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


### Index documents via HTTP

```
$ curl -XPOST http://localhost:20000/api/index -d '{
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
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1087  100    21  100  1066   1128  57274 --:--:-- --:--:-- --:--:-- 59222
{
  "document_count": 5
}
```

### Delete documents via HTTP

The `-d` option deletes the documents from specified Indigo Server.

```
$ curl -XDELETE http://localhost:20000/api/index -d '[
  "2",
  "3",
  "4"
]' | jq .
```

The result of the above delete command is:

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    44  100    21  100    23   1767   1936 --:--:-- --:--:-- --:--:--  2090
{
  "document_count": 3
}
```


### Search documents via HTTP

The `search` command queries the documents to specified Indigo Server.
See [Queries](http://www.blevesearch.com/docs/Query/), [Query String Query](http://www.blevesearch.com/docs/Query-String-Query/) and [type SearchRequest](https://godoc.org/github.com/blevesearch/bleve#SearchRequest) for more details.

#### Simple query

```
$ curl -XPOST http://localhost:20000/api/search -d '{
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
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   792  100   605  100   187  81701  25253 --:--:-- --:--:-- --:--:-- 86428
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
$ curl -XPOST http://localhost:20000/api/search -d '{
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
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  3105  100  1757  100  1348   172k   132k --:--:-- --:--:-- --:--:--  190k
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
$ curl -XPOST http://localhost:20000/api/search -d '{
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
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1126  100   861  100   265    99k  31435 --:--:-- --:--:-- --:--:--  105k
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
$ curl -XPOST http://localhost:20000/api/search -d '{
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
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  3452  100  2012  100  1440  52959  37903 --:--:-- --:--:-- --:--:-- 54378
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
