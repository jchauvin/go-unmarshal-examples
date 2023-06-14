# Yaml Commmon Unmarshalling

This example deals with the typical yaml scenario when you know how the entire document should resolve ahead of time.  In this case you can define the go structures and populate them using common design patterns

## Example Config - 
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

## Example Struct Objects - 

```
type Plugin struct {
	URI            string `yaml:"uri"`
	ExpectedResult uint   `yaml:"expectedResult"`
}

type Config struct {
	HostPort     int      `yaml:"hostPort"`
	MetricsPorts int      `yaml:"metricsPort"`
	MetricsPath  string   `yaml:"metricsPath"`
	LogLevel     string   `yaml:"logLevel"`
	Plugins      []Plugin `yaml:"plugins"`
}

```

## Unmarshalling -

```
var (
		cfg   models.Config
		bytes []byte
		err   error
)

bytes, err = os.ReadFile(filepath.Clean(path))
err = yaml.Unmarshal(bytes, &cfg)
	
```

## Running the example - 

```
cd yaml-normal
go run cmd/main.go -conf config/config.yaml
```