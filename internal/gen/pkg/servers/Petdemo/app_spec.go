// Code generated by sysl DO NOT EDIT.
package petdemo

const AppSpec = `{
 "apps": {
  "Petdemo": {
   "name": {
    "part": [
     "Petdemo"
    ]
   },
   "longName": "Petdemo",
   "attrs": {
    "package": {
     "s": "Petdemo"
    }
   },
   "endpoints": {
    "GET /pet": {
     "name": "GET /pet",
     "docstring": "Get a random pet",
     "attrs": {
      "patterns": {
       "a": {
        "elt": [
         {
          "s": "rest"
         }
        ]
       }
      }
     },
     "stmt": [
      {
       "call": {
        "target": {
         "part": [
          "Petstore"
         ]
        },
        "endpoint": "GET /pet"
       },
       "sourceContext": {
        "file": "specs/petdemo.sysl",
        "start": {
         "line": 10,
         "col": 12
        },
        "end": {
         "line": 10,
         "col": 24
        }
       },
       "sourceContexts": [
        {
         "file": "specs/petdemo.sysl",
         "start": {
          "line": 10,
          "col": 12
         },
         "end": {
          "line": 10,
          "col": 24
         }
        }
       ]
      },
      {
       "cond": {
        "test": "len(Pet) < 50",
        "stmt": [
         {
          "call": {
           "target": {
            "part": [
             "PokeAPI"
            ]
           },
           "endpoint": "GET /pokemon/{id}"
          },
          "sourceContext": {
           "file": "specs/petdemo.sysl",
           "start": {
            "line": 12,
            "col": 16
           },
           "end": {
            "line": 12,
            "col": 27
           }
          },
          "sourceContexts": [
           {
            "file": "specs/petdemo.sysl",
            "start": {
             "line": 12,
             "col": 16
            },
            "end": {
             "line": 12,
             "col": 27
            }
           }
          ]
         }
        ]
       },
       "sourceContext": {
        "file": "specs/petdemo.sysl",
        "start": {
         "line": 11,
         "col": 12
        },
        "end": {
         "line": 13,
         "col": 12
        }
       },
       "sourceContexts": [
        {
         "file": "specs/petdemo.sysl",
         "start": {
          "line": 11,
          "col": 12
         },
         "end": {
          "line": 13,
          "col": 12
         }
        }
       ]
      },
      {
       "ret": {
        "payload": "ok <: Pet"
       },
       "sourceContext": {
        "file": "specs/petdemo.sysl",
        "start": {
         "line": 13,
         "col": 12
        },
        "end": {
         "line": 13,
         "col": 18
        }
       },
       "sourceContexts": [
        {
         "file": "specs/petdemo.sysl",
         "start": {
          "line": 13,
          "col": 12
         },
         "end": {
          "line": 13,
          "col": 18
         }
        }
       ]
      }
     ],
     "restParams": {
      "method": "GET",
      "path": "/pet"
     },
     "sourceContext": {
      "file": "specs/petdemo.sysl",
      "start": {
       "line": 8,
       "col": 8
      },
      "end": {
       "line": 15,
       "col": 4
      }
     },
     "sourceContexts": [
      {
       "file": "specs/petdemo.sysl",
       "start": {
        "line": 8,
        "col": 8
       },
       "end": {
        "line": 15,
        "col": 4
       }
      }
     ]
    }
   },
   "types": {
    "Pet": {
     "tuple": {
      "attrDefs": {
       "breed": {
        "primitive": "STRING",
        "sourceContext": {
         "file": "specs/petdemo.sysl",
         "start": {
          "line": 16,
          "col": 8
         },
         "end": {
          "line": 16,
          "col": 17
         }
        },
        "sourceContexts": [
         {
          "file": "specs/petdemo.sysl",
          "start": {
           "line": 16,
           "col": 8
          },
          "end": {
           "line": 16,
           "col": 17
          }
         }
        ]
       },
       "pokemon": {
        "primitive": "STRING",
        "sourceContext": {
         "file": "specs/petdemo.sysl",
         "start": {
          "line": 17,
          "col": 8
         },
         "end": {
          "line": 17,
          "col": 19
         }
        },
        "sourceContexts": [
         {
          "file": "specs/petdemo.sysl",
          "start": {
           "line": 17,
           "col": 8
          },
          "end": {
           "line": 17,
           "col": 19
          }
         }
        ]
       }
      }
     },
     "sourceContext": {
      "file": "specs/petdemo.sysl",
      "start": {
       "line": 15,
       "col": 4
      },
      "end": {
       "line": 17,
       "col": 19
      }
     },
     "sourceContexts": [
      {
       "file": "specs/petdemo.sysl",
       "start": {
        "line": 15,
        "col": 4
       },
       "end": {
        "line": 17,
        "col": 19
       }
      }
     ]
    }
   },
   "sourceContext": {
    "file": "specs/petdemo.sysl",
    "start": {
     "line": 5
    },
    "end": {
     "line": 17,
     "col": 19
    }
   },
   "sourceContexts": [
    {
     "file": "specs/petdemo.sysl",
     "start": {
      "line": 5
     },
     "end": {
      "line": 17,
      "col": 19
     }
    }
   ]
  },
  "Petstore": {
   "name": {
    "part": [
     "Petstore"
    ]
   },
   "longName": "Pet Service",
   "attrs": {
    "description": {
     "s": "No description.\n"
    },
    "host": {
     "s": "australia-southeast1-innate-rite-238510.cloudfunctions.net"
    },
    "license": {
     "s": "MIT"
    },
    "package": {
     "s": "petstore"
    },
    "version": {
     "s": "1.0.0"
    }
   },
   "endpoints": {
    "GET /pet": {
     "name": "GET /pet",
     "docstring": "No description.",
     "attrs": {
      "patterns": {
       "a": {
        "elt": [
         {
          "s": "rest"
         }
        ]
       }
      }
     },
     "stmt": [
      {
       "ret": {
        "payload": "error <: Error"
       },
       "sourceContext": {
        "file": "specs/petstore.yaml",
        "start": {
         "line": 16,
         "col": 12
        },
        "end": {
         "line": 16,
         "col": 18
        }
       },
       "sourceContexts": [
        {
         "file": "specs/petstore.yaml",
         "start": {
          "line": 16,
          "col": 12
         },
         "end": {
          "line": 16,
          "col": 18
         }
        }
       ]
      },
      {
       "ret": {
        "payload": "ok <: Pet"
       },
       "sourceContext": {
        "file": "specs/petstore.yaml",
        "start": {
         "line": 17,
         "col": 12
        },
        "end": {
         "line": 17,
         "col": 18
        }
       },
       "sourceContexts": [
        {
         "file": "specs/petstore.yaml",
         "start": {
          "line": 17,
          "col": 12
         },
         "end": {
          "line": 17,
          "col": 18
         }
        }
       ]
      }
     ],
     "restParams": {
      "method": "GET",
      "path": "/pet"
     },
     "sourceContext": {
      "file": "specs/petstore.yaml",
      "start": {
       "line": 14,
       "col": 8
      },
      "end": {
       "line": 22,
       "col": 4
      }
     },
     "sourceContexts": [
      {
       "file": "specs/petstore.yaml",
       "start": {
        "line": 14,
        "col": 8
       },
       "end": {
        "line": 22,
        "col": 4
       }
      }
     ]
    }
   },
   "types": {
    "Error": {
     "tuple": {
      "attrDefs": {
       "code": {
        "primitive": "INT",
        "attrs": {
         "json_tag": {
          "s": "code"
         }
        },
        "constraint": [
         {
          "range": {
           "min": {
            "i": "-2147483648"
           },
           "max": {
            "i": "2147483647"
           }
          },
          "bitWidth": 32
         }
        ],
        "sourceContext": {
         "file": "specs/petstore.yaml",
         "start": {
          "line": 23,
          "col": 8
         },
         "end": {
          "line": 25,
          "col": 8
         }
        },
        "sourceContexts": [
         {
          "file": "specs/petstore.yaml",
          "start": {
           "line": 23,
           "col": 8
          },
          "end": {
           "line": 25,
           "col": 8
          }
         }
        ]
       },
       "message": {
        "primitive": "STRING",
        "attrs": {
         "json_tag": {
          "s": "message"
         }
        },
        "sourceContext": {
         "file": "specs/petstore.yaml",
         "start": {
          "line": 25,
          "col": 8
         },
         "end": {
          "line": 28,
          "col": 4
         }
        },
        "sourceContexts": [
         {
          "file": "specs/petstore.yaml",
          "start": {
           "line": 25,
           "col": 8
          },
          "end": {
           "line": 28,
           "col": 4
          }
         }
        ]
       }
      }
     },
     "sourceContext": {
      "file": "specs/petstore.yaml",
      "start": {
       "line": 22,
       "col": 4
      },
      "end": {
       "line": 28,
       "col": 4
      }
     },
     "sourceContexts": [
      {
       "file": "specs/petstore.yaml",
       "start": {
        "line": 22,
        "col": 4
       },
       "end": {
        "line": 28,
        "col": 4
       }
      }
     ]
    },
    "Pet": {
     "primitive": "STRING",
     "sourceContext": {
      "file": "specs/petstore.yaml",
      "start": {
       "line": 28,
       "col": 4
      },
      "end": {
       "line": 30
      }
     },
     "sourceContexts": [
      {
       "file": "specs/petstore.yaml",
       "start": {
        "line": 28,
        "col": 4
       },
       "end": {
        "line": 30
       }
      }
     ]
    }
   },
   "sourceContext": {
    "file": "specs/petstore.yaml",
    "start": {
     "line": 6
    },
    "end": {
     "line": 30
    }
   },
   "sourceContexts": [
    {
     "file": "specs/petstore.yaml",
     "start": {
      "line": 6
     },
     "end": {
      "line": 30
     }
    }
   ]
  },
  "PokeAPI": {
   "name": {
    "part": [
     "PokeAPI"
    ]
   },
   "longName": "PokeAPI Service",
   "attrs": {
    "description": {
     "s": "No description.\n"
    },
    "host": {
     "s": "pokeapi.co"
    },
    "license": {
     "s": "MIT"
    },
    "package": {
     "s": "pokeapi"
    },
    "version": {
     "s": "1.0.0"
    }
   },
   "endpoints": {
    "GET /pokemon/{id}": {
     "name": "GET /pokemon/{id}",
     "docstring": "No description.",
     "attrs": {
      "patterns": {
       "a": {
        "elt": [
         {
          "s": "rest"
         }
        ]
       }
      }
     },
     "stmt": [
      {
       "ret": {
        "payload": "error <: Error"
       },
       "sourceContext": {
        "file": "specs/pokeapi.yaml",
        "start": {
         "line": 16,
         "col": 12
        },
        "end": {
         "line": 16,
         "col": 18
        }
       },
       "sourceContexts": [
        {
         "file": "specs/pokeapi.yaml",
         "start": {
          "line": 16,
          "col": 12
         },
         "end": {
          "line": 16,
          "col": 18
         }
        }
       ]
      },
      {
       "ret": {
        "payload": "ok <: Pokemon"
       },
       "sourceContext": {
        "file": "specs/pokeapi.yaml",
        "start": {
         "line": 17,
         "col": 12
        },
        "end": {
         "line": 17,
         "col": 18
        }
       },
       "sourceContexts": [
        {
         "file": "specs/pokeapi.yaml",
         "start": {
          "line": 17,
          "col": 12
         },
         "end": {
          "line": 17,
          "col": 18
         }
        }
       ]
      }
     ],
     "restParams": {
      "method": "GET",
      "path": "/pokemon/{id}",
      "urlParam": [
       {
        "name": "id",
        "type": {
         "primitive": "INT",
         "sourceContext": {
          "file": "specs/pokeapi.yaml",
          "start": {
           "line": 13,
           "col": 13
          },
          "end": {
           "line": 13,
           "col": 21
          }
         },
         "sourceContexts": [
          {
           "file": "specs/pokeapi.yaml",
           "start": {
            "line": 13,
            "col": 13
           },
           "end": {
            "line": 13,
            "col": 21
           }
          }
         ]
        }
       }
      ]
     },
     "sourceContext": {
      "file": "specs/pokeapi.yaml",
      "start": {
       "line": 14,
       "col": 8
      },
      "end": {
       "line": 22,
       "col": 4
      }
     },
     "sourceContexts": [
      {
       "file": "specs/pokeapi.yaml",
       "start": {
        "line": 14,
        "col": 8
       },
       "end": {
        "line": 22,
        "col": 4
       }
      }
     ]
    }
   },
   "types": {
    "Error": {
     "tuple": {
      "attrDefs": {
       "code": {
        "primitive": "INT",
        "attrs": {
         "json_tag": {
          "s": "code"
         }
        },
        "constraint": [
         {
          "range": {
           "min": {
            "i": "-2147483648"
           },
           "max": {
            "i": "2147483647"
           }
          },
          "bitWidth": 32
         }
        ],
        "sourceContext": {
         "file": "specs/pokeapi.yaml",
         "start": {
          "line": 23,
          "col": 8
         },
         "end": {
          "line": 25,
          "col": 8
         }
        },
        "sourceContexts": [
         {
          "file": "specs/pokeapi.yaml",
          "start": {
           "line": 23,
           "col": 8
          },
          "end": {
           "line": 25,
           "col": 8
          }
         }
        ]
       },
       "message": {
        "primitive": "STRING",
        "attrs": {
         "json_tag": {
          "s": "message"
         }
        },
        "sourceContext": {
         "file": "specs/pokeapi.yaml",
         "start": {
          "line": 25,
          "col": 8
         },
         "end": {
          "line": 28,
          "col": 4
         }
        },
        "sourceContexts": [
         {
          "file": "specs/pokeapi.yaml",
          "start": {
           "line": 25,
           "col": 8
          },
          "end": {
           "line": 28,
           "col": 4
          }
         }
        ]
       }
      }
     },
     "sourceContext": {
      "file": "specs/pokeapi.yaml",
      "start": {
       "line": 22,
       "col": 4
      },
      "end": {
       "line": 28,
       "col": 4
      }
     },
     "sourceContexts": [
      {
       "file": "specs/pokeapi.yaml",
       "start": {
        "line": 22,
        "col": 4
       },
       "end": {
        "line": 28,
        "col": 4
       }
      }
     ]
    },
    "Pokemon": {
     "tuple": {
      "attrDefs": {
       "height": {
        "primitive": "INT",
        "attrs": {
         "json_tag": {
          "s": "height"
         }
        },
        "opt": true,
        "sourceContext": {
         "file": "specs/pokeapi.yaml",
         "start": {
          "line": 29,
          "col": 8
         },
         "end": {
          "line": 31,
          "col": 8
         }
        },
        "sourceContexts": [
         {
          "file": "specs/pokeapi.yaml",
          "start": {
           "line": 29,
           "col": 8
          },
          "end": {
           "line": 31,
           "col": 8
          }
         }
        ]
       },
       "id": {
        "primitive": "INT",
        "attrs": {
         "json_tag": {
          "s": "id"
         }
        },
        "opt": true,
        "sourceContext": {
         "file": "specs/pokeapi.yaml",
         "start": {
          "line": 31,
          "col": 8
         },
         "end": {
          "line": 33,
          "col": 8
         }
        },
        "sourceContexts": [
         {
          "file": "specs/pokeapi.yaml",
          "start": {
           "line": 31,
           "col": 8
          },
          "end": {
           "line": 33,
           "col": 8
          }
         }
        ]
       },
       "name": {
        "primitive": "STRING",
        "attrs": {
         "json_tag": {
          "s": "name"
         }
        },
        "opt": true,
        "sourceContext": {
         "file": "specs/pokeapi.yaml",
         "start": {
          "line": 33,
          "col": 8
         },
         "end": {
          "line": 35
         }
        },
        "sourceContexts": [
         {
          "file": "specs/pokeapi.yaml",
          "start": {
           "line": 33,
           "col": 8
          },
          "end": {
           "line": 35
          }
         }
        ]
       }
      }
     },
     "sourceContext": {
      "file": "specs/pokeapi.yaml",
      "start": {
       "line": 28,
       "col": 4
      },
      "end": {
       "line": 35
      }
     },
     "sourceContexts": [
      {
       "file": "specs/pokeapi.yaml",
       "start": {
        "line": 28,
        "col": 4
       },
       "end": {
        "line": 35
       }
      }
     ]
    }
   },
   "sourceContext": {
    "file": "specs/pokeapi.yaml",
    "start": {
     "line": 6
    },
    "end": {
     "line": 35
    }
   },
   "sourceContexts": [
    {
     "file": "specs/pokeapi.yaml",
     "start": {
      "line": 6
     },
     "end": {
      "line": 35
     }
    }
   ]
  }
 }
}`