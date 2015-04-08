#!/bin/bash

set -e # fail on error

go test ./... 
go install 
echo "Built and tested"
