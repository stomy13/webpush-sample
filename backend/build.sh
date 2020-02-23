#!/bin/bash

if [ -e build.env ]; then
	source build.env
fi

pushd webpush
go build
popd