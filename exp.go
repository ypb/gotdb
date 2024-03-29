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
	// "ma kota" or "nie ma \"\""
	if Or = foo.Store("ala", "", tdb.INSERT); Or != nil {
		println("Error0:", Or.String())
	}
	// var val tdb.DATA
	var sval string
	if sval, Or = foo.Fetch("ala"); Or == nil {
		println("Value0: \"" + sval + "\"")
	} else {
		println("Error1:", Or.String())
	}
	sval, Or = foo.Fetch("nokey")
	println("Insanity continues \""+sval+"\"", Or.String())
	// trivalence of the bivalence AHOY!
	foo.Store("", "shit my foot is off", tdb.INSERT)
	sval, Or = foo.Fetch("")
	if Or != nil {
		print("Error2: \"" + Or.String() + "\" ")
	}
	println("Value1: \"" + sval + "\"")

	goo, _ := tdb.Open("exp.tdb", 256, tdb.NOSYNC, tdb.O_RDWR, tdb.USR_RW|tdb.GRP_R|tdb.OTH_R)
	println("goo:", goo.String())
	// println("argh!", goo)
	// goo.Debug() // here we "were" turning off DEBUG since goo is the same obj
	if Or = foo.Close(); Or != nil {
		println("Safe Error:", Or.String())
	} // safe
	println("Unsafe Error:", foo.Close().String()) // unsafe;-(
	// and breaks intended indentation )-;
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
