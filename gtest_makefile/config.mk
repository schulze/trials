# Customize below to fit your system

# paths
PREFIX = /usr/local

GTEST_ROOT=/usr/src/gtest
GTEST_INCLUDE=/usr/include
#GTEST_ROOT=${HOME}/src/googlecode.com/googletest
#GTEST_INCLUDE=${GTEST_ROOT}/include

# includes and libs
INCS = -I/usr/include
LIBS = -L/usr/lib

# flags
CPPFLAGS = -D_SVID_SOURCE
CFLAGS = -Wall -Wextra -ansi -Os ${INCS} ${CPPFLAGS}
LDFLAGS = -s ${LIBS}

# compiler and linker
CC = g++
