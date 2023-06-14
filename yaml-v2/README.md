# YAML V2 Delayed unmarshalling


Unlike the json package the yaml package doesn't have a type similar to the RawMessage


There was actually a discussion on a feature add of RawMessage to yaml, and why it wouldn't be able to work as indended with json - https://github.com/go-yaml/yaml/issues/13#issuecomment-135407098


Basically it came down to the differences between yaml and json, that we highlight in the json directory README.md.  YAML marshaling is context sensitive, and you can run into edge case issues when marshalling data back.  It isn't as strait forward as JSON. 


## YAML Marshaller interface

To mimic what JSON was doing with the RawMessage in YAML V2 you have a few options, but here we will detail how to create type that implements the UnMarshUnmarshaler interface.  

The core of this example uses the type - 
```
type RawMessage struct {
	unmarshal func(interface{}) error
}

func (msg *RawMessage) UnmarshalYAML(unmarshal func(interface{}) error) error {
	msg.unmarshal = unmarshal
	return nil
}

func (msg *RawMessage) Unmarshal(v interface{}) error {
	return msg.unmarshal(v)
}
```

Then in a golang structure you can define an object to that type - 
```
type Plugin struct {
	Type   string      `yaml:"type"`
	Config *RawMessage `yaml:"config"`
}
```

After you read in the entire yaml config you can then unmarshal that RawMessage seperately later to the correct struct-
```
c := &models.HTTPPluginConfig{}
p.Config.Unmarshal(c)
```

