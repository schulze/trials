#! /bin/sh

cat $1 | tred | unflatten | dot -Tsvg > `basename $1 .dot`.svg