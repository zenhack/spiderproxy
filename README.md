Spiderproxy builds a tree of SSH connections, and offers a socks proxy
on the local machine which intelligently routes traffic to the right
places.

Spiderproxy uses your ssh agent to authenticate; it must be running in
order for this to work.

# Installing

    go get zenhack.net/go/spiderproxy

# Configuration File Reference

Spiderproxy reads a configuration file (by default `spiderproxy.json` in
the current directory) which specifies a tree of connections to build,
and rules for routing traffic based on destination addresses.

The file must contain a single [JSON][1] array, where each element of
the array specifies a host to connect to via ssh. Each host is an object
with the following fields:

* `"comment"` (string): A comment describing the host; meant for humans.
* `"host"` (string): The host name or IP address of the host to connect
  to.
* `"port"` (number, optional): The TCP port to connect to. Defaults to
  `22`.
* `"user"` (string, optional): The name of the user to log in as.
  Defaults to the value of the `USER` environment variable.
* `"match"` (array of strings): An array of shell glob patterns. When a
  client asks the socks proxy to establish a connection, the connection
  will be tunneled through the first host who's `"match"` field contains a
  glob pattern matching the destination host.
* `"next"` (array of host objects): An array of the same form as the
  top-level object. When building the tree, connections any hosts
  specified in a `"next"` field will be tunneled through the host to which
  the field belongs. `"next"` fields may be nested.

The file `example.spiderproxy.json` in the root of this repository provides
an example.

# Invocation

If run with no arguments, Spiderproxy will read the file `spiderproxy.json`
in the current directory, build a tree of ssh connections according to the
rules above, and start a socks listening on port 1080, which will route
traffic accordingly. The `-addr` command line option can be used to
override the port (and control what ip addresses the server listens on),
while the option `-config` specifies the path to an alternate config
file.

# LICENSE

Free/Open Source under the MIT license (see `COPYING`)

[1]: http://json.org
