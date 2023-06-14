# Yaml Delayed Unmarshalling


This repo serves to demonstrate a few scenarios of unmarshalling data in golang using yaml and json.  Specifically it shows how you can delay unmarshalling of certain portions of a document until a later time.  In the yaml-v2, yaml-v3, and json directories we want to delay to solve the problem where you might have a section of your document that will resolve differently per component

## Problem Description - 

Normally when you are unmarshalling data you need to know the type of objects that you are unmarshalling to ahead of time.  However there are a few situations where you might want to have sections of a document use a common field name but act differently and unmarshal to a different  struct.  In a situation where you have a list of plugins, each plugin could have a `config` field, however you might have each plugin config differ between types.  There could be an http plugin that has a radically different config structure then another plugin.  In that situation you need to delay unmarshalling that part of the document until you figure out what type of object to unmarshal to.  



## Examples - 

- **yaml-normal** - unmarshalling a common configuration.  All structures are known ahead of time
- **json** - umarshalling with a json raw message to delay config unmarshalling.  Each plugin has its own configuration type
- **yaml-v2** - umarshalling with a custom struct that implements the unmarshalling interface to delay config unmarshalling.  Each plugin has its own configuration type
- **yaml-v3** - umarshalling with a yaml.Node object that will delay config unmarshalling.  Each plugin has its own configuration type

Each directory should have it's own README.md file that explains how the example works in more detail
allow the unmarshalling to work


## Common Unmarshalling scenario - 

In the most common yaml/json unmarshallign scenario, you know how to fully resolve the whole document - 

Example - 
```
metricsPath: "/metrics"
metricsPort: 9121
hostPort: 3000
logLevel: "info"
plugins :
    - uri: "https://www.google.com"
      expectedResult: 200
    - uri: "https://www.ibm.com"
      expectedResult: 200
    - uri: "http://localhost:3000"
      expectedResult: 200
```

Each plugin above has the same config.  You can create a common go structure and then unmarshal the above into an array of those structures.


## Complex Unmarshalling Scenario - 

In a more interesting example, we could have a situation where you want to have each plugin have its own config fieldname, however the structure of that config will differ from plugin to plugin.  

Example - 
```
metricsPath: "/metrics"
metricsPort: 9121
hostPort: 3000
logLevel: "info"
plugins :
    - type: "http"
      config:
        uri: "https://www.google.com"
        expectedResult: 200
    - type: "secrets_manager"
      config:
        uri: "https://70e21385-ccf8-437b-8f05-b77db3821603.us-south.secrets-manager.appdomain.cloud"
        iamAuthType: "apikey"
        secrets:
            - secretID: "someid1"
              secretName: "someid2"
            - secretID: "anotherid1"

```

Notice how the fields of the two plugin configs differ quite a bit, and require different go structs to unmarshall to.  However at the time of loading the config and unmarshalling the whole document, you won't know the struct to use for each plugin until you evaluate it's type.

In this situation you need to define the config as a generic object that you can unmarshall at a later point. 