#! /bin/sh

cat $1 | tred | unflatten | dot -Tpng > `basename $1 .dot`.png
