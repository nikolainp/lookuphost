# consul-template-plugin

## About
This is a simple little plugin to let you lookup ip address by host for [consul-template](https://github.com/hashicorp/consul-template). 
Download it from [releases](https://github.com/nikolainp/lookuphost/releases) and it's ready to use.

## Examples

```
# master postgres
{{ range service "master.postgres|passing" }}{{ range $d := plugin "/usr/local/bin/lookuphost" .Address | parseJSON}}{{ $d }} master.postgres
{{ end }}{{ end }}
```
