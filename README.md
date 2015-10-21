rfc3339log
===========

A drop-in replacement for the standard logger, with the following features:

- RFC3339 are always formatted timestamps (unfortunately no other options for now)
- The zero-value pointer (*Logger) is usable. All output methods will no-op instead of panic.

For documentation, check [godoc](https://godoc.org/github.com/tomwans/rfc3339log).
