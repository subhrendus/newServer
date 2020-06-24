### Intend
1. simple interface, only debug, info, warn, error
2. always use json format, every log must be json parsable
3. add meta information one-time, call them default fields/context
4. always add timestamp in the utc format
5. always log to stdout or a file
6. uniform config struct
7. auto inject 'environment', from the config. This will be part of the default fields

### Notes:
1. logs should be discoverable, inject meta (context) information to every log
2. by choice, no 'fatal' & 'panic'. The caller should explicitly call os.Exit, or panic()
3. always log with default fields (context)


### TODO:
 	1. use the logstash hook, if any issues with parsing occur after filebeat flushes out log
 	2. log to file, controlled via config

### Example:

```go
package main

import (
	"fmt"
	"os"

	"github.com/subhrendus/newServer/logging"
)

func main() {
	// initiate log with common config
	logConfig := logging.LogConfig{AppName: "DemoApp", AppVersion: "1.1.1", EngGroup: "video-infra", Environment: "development", Level: "INFO"}

	logger, err := logging.New(&logConfig)

	if err != nil {
		fmt.Printf("something went wrong. Error: %v", err)
		os.Exit(1)
	}

	// the logger should add context fields by default, without any additional instrumentation
	logger.Info("hola mundo, first log")

	// this shouldn't be printed if the level is set to 'INFO'
	logger.Debug("here is a debug statement")

	// log one more time for context fields
	additionalFields := logging.DataFields{"foo": "bar"}
	logger.Info("still going strong", additionalFields)

	// log a simple statement without additional fields
	logger.Info("once more, without any context fields")

	// log an error
	logger.Error("deal with it")
}
```

