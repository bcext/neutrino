#!/bin/bash

git checkout $(cat $GOPATH/src/github.com/bcext/neutrino/glide.yaml | grep -A1 gcash | tail -n1 | awk '{ print $2}')
