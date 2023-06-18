#!/bin/bash
gcc -o run test.c minioc.dylib
./run
rm -rf run
