// Copyright 2011 AUTHORS. All rights reserved.
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

// Package gotdb implements Go bindings to C's libtdb (Trivial DataBase)
//
// TDB is a filesystem or in-memory key/value store in the vein of
// (G/B)DB/M API, a common DB abstraction layer used by the SAMBA project
package tdb

// #cgo LDFLAGS: -ltdb
// #cgo pkg-config: tdb
// #include <stdlib.h>
// #include <tdb.h>
import "C"
import "unsafe"

// DB type is a struct...
type DB struct {
	pth string                // path name
	Dbg bool                  // to DEBUG or not to DEBUG
	ctx *C.struct_TDB_CONTEXT // handle me gently
}

// String returns string representation of *DB
func (db *DB) String() (s string) {
	s = "DB{pth:\"" + db.pth + "\""
	if db.Dbg {
		s += ", Dbg:true"
	} else {
		s += ", Dbg:false"
	}
	if db.ctx == nil {
		s += ", ctx:#f}"
	} else {
		s += ", ctx:#t}"
	}
	return
}

// New given a string returns *DB. One must keep in mind that given string
// of the same value New will return fresh *DB refering to the same underlying
// DB. Performing Close() on any of the *DB opened with the same, unique path
// will thereafter fail with ERR_NOEXIST status, (until another succesful
// New or Open is executed?)...
func New(path string) *DB {
	return Open(path, 0, DEFAULT, O_RDWR|O_CREAT, USR_RW)
}

// Open is used by New with some reasonably default initial values besides
// path name. This is Open's signature in a Go caricature of libtdb
// original C tdb_open function:
// func Open(name const *C.char, hash_size, tdb_flags, open_flags C.int, mode C.mode_t) *C.struct_tdb_context
func Open(path string, hash_size, tdb_flags, open_flags C.int, mode C.mode_t) (db *DB) {
	name := C.CString(path)
	defer C.free(unsafe.Pointer(name))
	return &DB{path, false,
		C.tdb_open(name, hash_size, tdb_flags, open_flags, mode)}
}

// Close calls tdb_close on the C's ctx pointer contained in DB struct,
// rendering it invalid in all other instances of the same name. see New
// and here is trivially meaningless C signature:
// func Close() C.int
func (db *DB) Close() int {
	if db.Dbg {
		println("tdb.Close()", db.String())
	}
	if db.ctx == nil {
		return ERR_NOEXIST
	}
	return int(C.tdb_close(db.ctx))
}

// Local Variables:
// mode: Go
// End:
