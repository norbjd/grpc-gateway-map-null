# Demonstrate gRPC gateway behavior with nullable values in maps

Using [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) v2.7.2.

This is a simple service just echoing a `map<string, google.protobuf.StringValue>`.

Generated Go code is a `map[string]*wrapperspb.StringValue`, so the value could be `nil`.

In grpc-gateway version 2.0.0 or greater, passing `null` value in a JSON fails and does not set `*wrapperspb.StringValue` to `nil` :

In grpc-gateway version 1.16.0, passing `null` value in a JSON set `*wrapperspb.StringValue` to empty string `""`.

## Reproduce issue

Checkout the right branch associated with the version of grpc-gateway (`git checkout`) :

- master (v2.7.2)
- v2.0.1
- v1.16.0

Run server and run gateway :

```shell
make server
make gateway
```

Send a HTTP request with body :

```json
{
  "aMap": {
    "someString": "test",
    "emptyString": "",
    "nullableString": null
  }
}
```

On version 2.7.2 and 2.0.1 :

```shell
$ make echo
{"code":3, "message":"proto: (line 1:70): invalid value for string type: null", "details":[]}
```

On version 1.16.0 :

```shell
$ make echo
{"aMap": {"someString": "test", "emptyString": "", "nullableString": ""}}
```
