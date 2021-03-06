# hermes-utils

Some amazing & helpful utils for **Hermes** project.

## Usage

`import "github.com/mikevel2955/hermes-utils"`

## Content

#### ReadConfig
It reads config from environment variables into passed struct using Golang reflection

##### Example
```
type config struct {
	Addr     string        `env:"ADDR"`
	Port     int           `env:"PORT" def:"8081"`
	ExtPort  uint32        `env:"EXT_PORT" def:"777"`
	Timeout  time.Duration `env:"TIMEOUT" required:"true"`
	IsRemote bool          `env:"IS_REMOTE" def:"true"`
	Rate     float64       `env:"RATE" def:"0.05"`
}

func main() {
	config := config{}
	if err := ReadConfig(&config); err != nil {
		panic(err)
	}
}
```

## Enjoy!