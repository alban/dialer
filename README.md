# Dialer

Dialer is a simple Golang program that, given a `hostname:port` tuple, establishes *n* TCP connections to it and then sleeps forever.

We use it as a way to benchmark long-lasting connections in Scope.

## USAGE

```
$ dialer github.com:80 10
```
