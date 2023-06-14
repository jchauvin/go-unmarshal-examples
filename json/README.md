# JSON Delayed unmarshalling

## RawMessage

The json package of golang has a RawMessage type.  When you are unmarshalling json, typically you will fully resolve the entire JSON document into a single go structure, or multiple structures.  They can by maps, or other predefined custom structs.  However there will be situations where you don't know what part of a document should look like, or don't care.  If you use json.RawMessage it acts as a place holder and you don't have to define it.  Effectivley it is a byte array alias


## Yaml vs JSON

YAML is a superset of JSON. It was developed around the same time to handle more kinds of data and offer a more complex but still readable syntax.  Because of this all JSON is actually fully valid YAML

A typical yaml file could look like - 
```
- name: Sam
  age: 22
- name: Jessica
  age: 40
- name: Jeff
  age: OLD
```

But it is also valid YAML to be written as -
```
- {name: Sam, age: 22}
- {name: Jessica, age: 40}
- {name: Jeff, age: OLD}
```

Or - 
```
[{name: "Sam", age: "22"}, {name: "Jessica", age: "40"}, {name: "Jeff", age: "OLD"}]
```

Key Differences - 

- Yaml allows comments, JSON does not
- Yaml uses indentation and spaces to denote a heirarchy and context, Json uses braces and brackets
- Yaml supports more data types then JSON
- Yaml is more flexible with string quotation.  They are optional, and single or double quotes are allowed.  Json requires strings to be double quoted
- Yaml allows the root node to be any valid data type, where as JSON requires it to be an array or object

