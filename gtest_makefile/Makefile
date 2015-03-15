include config.mk

PROGNAME=square

SRC = ${PROGNAME}.cc
OBJ = ${SRC:.cc=.o}

TESTSRC = ${PROGNAME}_test.cc
TESTOBJ = ${TESTSRC:.cc=.o}

all: options unittest # ${PROGNAME}

options:
	@echo ${PROGNAME} build options:
	@echo "CFLAGS   = ${CFLAGS}"
	@echo "LDFLAGS  = ${LDFLAGS}"
	@echo "CC       = ${CC}"

.cc.o:
	@echo CC $<
	@${CC} -c -I${GTEST_INCLUDE} ${CFLAGS} $<

${OBJ}: config.mk

${PROGNAME}: ${OBJ}
	@echo CC -o $@
	@${CC} -o $@ ${OBJ} ${LDFLAGS}

unittest: ${OBJ} ${TESTOBJ} libgtest.a
	@${CC} -o $@ ${OBJ} ${TESTOBJ} libgtest.a -lpthread ${LDFLAGS}

libgtest.a: config.mk
	g++ -I${GTEST_ROOT} -I${GTEST_INCLUDE} ${CFLAGS} -c ${GTEST_ROOT}/src/gtest-all.cc
	ar -rv libgtest.a gtest-all.o

clean:
	@echo cleaning
	@rm -f ${PROGNAME} ${OBJ} ${TESTOBJ} gtest-all.o libgtest.a unittest

install: all
	@echo installing executable file to ${DESTDIR}${PREFIX}/bin
	@mkdir -p ${DESTDIR}${PREFIX}/bin
	@cp -f ${PROGNAME} ${DESTDIR}${PREFIX}/bin
	@chmod 755 ${DESTDIR}${PREFIX}/bin/${PROGNAME}

uninstall:
	@echo removing executable file from ${DESTDIR}${PREFIX}/bin
	@rm -f ${DESTDIR}${PREFIX}/bin/${PROGNAME}

.PHONY: all options clean install uninstall
