#!/bin/bash

echo inline
time go run inline.go

echo .
echo concurrent
time go run concurrancy.go
