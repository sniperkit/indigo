# Bleve Server

Bleve Server is a search and indexing server built on top of [Bleve](http://www.blevesearch.com).


## bleve-server command

The `bleve-server` command provide following sub-commands.

```
$ bleve-server --help
Bleve Server is a search server built on top of Bleve.

Usage:
  bleve-server [flags]
  bleve-server [command]

Available Commands:
  start       start Bleve Server
  version     show version number

Use "bleve-server [command] --help" for more information about a command.
```

### start command

The `start` command starts the Bleve Server.

```
$ bleve-server start
Start Bleve Server localhost:10000
```


### bleve-server.yaml

The bleve-server.yaml file is the configuration file with the parameters affecting Bleve Server itself.

```
#
# Logging configuration
#
log:
    file: ./bleve-server.log
    level: info
    format: text

#
# Blever Server configuration
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

## bleve-cli command

The `bleve-cli` command provide following sub-commands.

```
$ ./bleve-cli/bleve-cli --help
Bleve Command Line Interface is controlls the Bleve Server.

Usage:
  bleve-cli [flags]
  bleve-cli [command]

Available Commands:
  search      searches the documents from Bleve Server
  index       indexes the documents to Bleve Server
  mapping     prints the mapping used for Bleve Server
  version     shows version number

Use "bleve-cli [command] --help" for more information about a command.
```

### mapping command

The `mapping` command shows IndexMapping used by the Bleve Server.

```
$ bleve-cli mapping
```

The result of the above mapping command is:

```
{
  "status": 0,
  "message": "success",
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

See [Introduction to Index Mappings](http://www.blevesearch.com/docs/Index-Mapping/) and [type IndexMappingImpl](https://godoc.org/github.com/blevesearch/bleve/mapping#IndexMappingImpl) for more details.


### index command

#### index documents

The `index` command indexes the documents to specified Bleve Server.

```
$ bleve-cli index '{
  "1": {
    "name": "Bleve",
    "description": "Full-text search library written in Go.リ",
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
    "name": "Bleve Server",
    "description": "Full-text search server built on Bleve.",
    "category": "Server",
    "popularity": 5.0,
    "release": "2017-01-13T00:00:00Z",
    "type": "document"
  }
}'
```

The result of the above index command is:

```
{
  "status": 0,
  "message": "success",
  "document_count": 5
}
```

### deletes documents

The `-d` option deletes the documents from specified Bleve Server.

```
$ bleve-cli index -d '[
  "2",
  "3",
  "4"
]'
```

The result of the above delete command is:

```
{
  "status": 0,
  "message": "success",
  "document_count": 3
}
```


### search command

The `search` command queries the documents to specified Bleve Server.
See [Queries](http://www.blevesearch.com/docs/Query/), [Query String Query](http://www.blevesearch.com/docs/Query-String-Query/) and [type SearchRequest](https://godoc.org/github.com/blevesearch/bleve#SearchRequest) for more details.

#### Simple query

```
$ bleve-cli query '{
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
}'
```

The result of the above simple query command is:

```
{
  "status": 0,
  "message": "success",
  "search_result": {
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
        "score": 0.7823224437894639,
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
    "max_score": 0.7823224437894639,
    "took": 166357,
    "facets": {}
  }
}
```

### Faceted query

```
$ bleve-cli query '{
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
}'
```

The result of the above faceted query command is:

```
{
  "status": 0,
  "message": "success",
  "search_result": {
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
        "score": 0.7823224437894639,
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
    "max_score": 0.7823224437894639,
    "took": 253853,
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
}
```

### Highlighted query

```
$ bleve-cli query '{
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
}'
```

The result of the above highlighted query command is:

```
{
  "status": 0,
  "message": "success",
  "search_result": {
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
        "score": 0.7823224437894639,
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
            "Full-text search library written in \u003cmark\u003eGo\u003c/mark\u003e."
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
    "max_score": 0.7823224437894639,
    "took": 169730,
    "facets": {}
  }
}
```

### Faceted and highlighted query

```
$ bleve-cli query '{
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
}'
```

The result of the above faceted and highlighted query command is:

```
{
  "status": 0,
  "message": "success",
  "search_result": {
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
        "score": 0.5121502312722896,
        "locations": {
          "description": {
            "Go": [
              {
                "pos": 13,
                "start": 36,
                "end": 38,
                "array_positions": null
              }
            ]
          }
        },
        "fragments": {
          "description": [
            "Full-text search library written in \u003cmark\u003eGo\u003c/mark\u003e."
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
    "max_score": 0.5121502312722896,
    "took": 4562801,
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
}
```

See [type SearchResult](https://godoc.org/github.com/blevesearch/bleve#SearchResult) for more details.



## License

Apache License Version 2.0
