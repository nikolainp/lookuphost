# consul-template-plugin

## About
This is a simple little plugin to let you lookup ip address by host for [consul-template](https://github.com/hashicorp/consul-template). 

## Examples

```
{{ range $d = plugin "lookuphost" .Address }}
    {{ $d }} .Name
{{ end }}
```
