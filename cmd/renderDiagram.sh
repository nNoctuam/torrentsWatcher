#!/bin/bash

# https://github.com/kisielk/godepgraph + graphviz
godepgraph -p golang.org,gopkg.in,github.com,google.golang.org,go.uber.org -s ./cmd/torrentsWatcher | dot -Tpng -o graph.png
