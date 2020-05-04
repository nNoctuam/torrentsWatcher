#!/bin/bash

# https://github.com/kisielk/godepgraph + graphviz
godepgraph -p golang.org,gopkg.in,github.com -s ./torrentsWatcher | dot -Tpng -o ../docs/graph.png
