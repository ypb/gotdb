// Copyright 2011 AUTHORS. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package main

// local
import tdb "./_obj/tdb"
// {GCIMPORTS=-I,LDIMPORTS=-L}${GOPATH}/pkg/${GOOS}_${GOARCH} make exp
// import tdb "github.com/ypb/gotdb"

func main() {
	foo := tdb.New("exp.tdb")
	foo.Dbg = true
	goo := tdb.Open("exp.tdb", 256, tdb.NOSYNC, tdb.O_RDWR, tdb.USR_RW | tdb.GRP_R | tdb.OTH_R)
	goo.Dbg = true
	println(foo.Close())
	println(goo.Close())
}

// Local Variables:
// mode: Go
// End:
