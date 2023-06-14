# YAML V3 Delayed unmarshalling

In moving from yaml v2 to v3 a new object called yaml.Node was introduced to mimc the behavior of the JSON RawMessage.  It is the prefered object to use now when delaying yaml unmarshalling, as opposed to the custom type with an unmarshaller, or yaml.MapSlice object. The MapSlice object was actually removed from the package and replaced by the yaml.Node object 

## yaml.Node

Node represents an element in the YAML document hierarchy. While documents are typically encoded and decoded into higher level types, such as structs and maps, Node is an intermediate representation that allows detailed control over the content being decoded or encoded.  You can use the Encode and Decode receiver functions to marshal and unmarshal 

You utilize the type similar to the json RawMessage -
```
type Plugin struct {
	Type   string      `yaml:"type"`
	Config *yaml.Node `yaml:"config"`
}
```

After you read in the entire yaml config you can then unmarshal that yaml.Node seperately to the correct struct using the Decode function -
```
c := &models.HTTPPluginConfig{}
err = p.Config.Decode(c)
```



