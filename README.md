# rifter

This is a [MQTT](http://mqtt.org/) reverse proxy which will do the following:

* Authenticate incoming connections.
* Load balance across back-end MQTT servers.
* Rate limit clients messages/sec or bytes/sec.
* Record metrics on different messages types.

It uses REDIS to share the state between load balancers, and emits metrics to [statsd](https://github.com/etsy/statsd/).

This project is currently in the very early stages, I don't recommend using it till it is feature complete and tested.

# usage

```
Usage of ./rifter:
  -alsologtostderr=false: log to standard error as well as files
  -auth="disabled": Authentication method 'disabled' or 'redis'
  -log_backtrace_at=:0: when logging hits line file:N, emit a stack trace
  -log_dir="": If non-empty, write log files in this directory
  -logtostderr=false: log to standard error instead of files
  -mqtt=[]: MQTT endpoints
  -port=":1883": The IP Address and TCP port to listen
  -redis="": REDIS which is used for authentication
  -sslcert="": File containing SSL Certificates
  -sslkey="": File containing SSL Private Key
  -stderrthreshold=0: logs at or above this threshold go to stderr
  -v=0: log level for V logs
  -vmodule=: comma-separated list of pattern=N settings for file-filtered logging
 ```

# License

*rifter* is Copyright (c) 2014 Mark Wolfe and licensed under the MIT License. See the included LICENSE file for more details.