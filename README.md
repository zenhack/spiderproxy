Spiderproxy builds a tree of SSH connections, and offers a socks proxy 
on the local machine which intelligently routes traffic to the right 
places.

Detailed usage documentation is in the pipeline, but for now the file 
`example.spiderproxy.json` can provide an intuition for what the 
configuration looks like; pass a file like that to Spiderproxy, and it 
will build the tree, and route traffic based on the "matches" rules.

Spiderproxy uses your ssh agent to authenticate; it must be running in 
order for this to work.

Install with:

    go get zenhack.net/go/spiderproxy

# LICENSE

Free/Open Source under the MIT license (see `COPYING`)
