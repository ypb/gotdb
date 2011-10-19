# Copyright 2011 AUTHORS. All rights reserved.
# Use of this source code is governed by a LGPL-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=tdb

GOFILES=tdb.h.c.go

CGOFILES=tdb.go
#CGO_LDFLAGS=$(shell pkg-config --libs $(TARG))

CLEANFILES=tdb.h.c.go *~ exp*.tdb exp

include $(GOROOT)/src/Make.pkg

tdb.h.c.go: tdb.h.c
	godefs -g $(TARG) $< > $@
	gofmt -w $@

# tentative testing...
exp: exp.go _obj/$(TARG).a
	$(GC) $(GCIMPORTS) $<
	$(LD) $(LDIMPORTS) -o $@ $@.$(O)

fmt:
	gofmt -d $(CGOFILES)

# Local Variables:
# mode: Makefile
# End:
