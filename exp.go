// Copyright 2011 AUTHORS. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package main

// local
import tdb "./_obj/tdb"
// {GCIMPORTS=-I,LDIMPORTS=-L}${GOPATH}/pkg/${GOOS}_${GOARCH} make exp
// import tdb "github.com/ypb/gotdb"

func main() {
	var Or tdb.Error // LULZ
	foo, Or := tdb.New("exp.tdb")
	// tdb.Error can be nil per the "convention"
	// println(Or.Errno()) // a "nil, no-no"
	// bar := foo
	// exp.go:14: implicit assignment of unexported field 'db' of tdb.DB in assignment
	// good!
	foo.Debug()
	goo, _ := tdb.Open("exp.tdb", 256, tdb.NOSYNC, tdb.O_RDWR, tdb.USR_RW|tdb.GRP_R|tdb.OTH_R)
	println(goo.String())
	// println("argh!", goo)
	// goo.Debug() // here we "were" turning off DEBUG since goo is the same obj
	if Or = foo.Close(); Or != nil {
		println(Or.String())
	} // safe
	println(foo.Close().String()) // unsafe;-(
	goo.Close()
	boo, _ := tdb.New("exp.tdb")
	// boo.Debug() // now this will turn off foo.Debug()'s on
	boo.Close()
	// Pure Insanity! gotta fix this!
	// foo,_ = tdb.New("exp.tdb")
	// foo.Debug()
	// FIXED
	goo.Close()
	boo.Close()
	foo.Close()
}

// Local Variables:
// mode: Go
// End:
