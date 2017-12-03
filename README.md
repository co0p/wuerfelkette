[![Build Status](https://travis-ci.org/co0p/wuerfelkette.svg?branch=master)](https://travis-ci.org/co0p/wuerfelkette) [![Coverage Status](https://coveralls.io/repos/github/co0p/wuerfelkette/badge.svg)](https://coveralls.io/github/co0p/wuerfelkette)

Wuerfelkette
============

Demonstration implementation of a block chain in golang.

There is a CHAIN which contains BLOCKs. To add a BLOCK to the CHAIN you have to MINE a new BLOCK. There are multiple strategies for mining a block. A SimpleMiner and a ParallelMiner.

build
-----

    go build -o wuerfelclient client/main.go


testing
-------

    go test ./...


