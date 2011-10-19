// Copyright 2011 AUTHORS. All rights reserved.
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

/*
 Binding of TDB (a Trivial DataBase, in the vein of (G/B)DB/M, a common
 DB abstraction layer used by the SAMBA project)
*/
package tdb

// #cgo LDFLAGS: -ltdb
// #cgo pkg-config: tdb
// #include <stdlib.h>
// #include <tdb.h>
import "C"
import "unsafe"

type DB struct {
	pth string                // path name
	Dbg bool                  // to DEBUG or not to DEBUG
	ctx *C.struct_TDB_CONTEXT // handle me gently
}

func (db *DB) String() (s string) {
	s = "DB{pth:\"" + db.pth + "\""
	if db.Dbg {
		s += ", Dbg:true"
	} else {
		s += ", Dbg:false"
	}
	if db.ctx == nil {
		s += ", nil}"
	} else {
		s += "}"
	}
	return
}

// just D0 it!
func New(path string) *DB {
	return Open(path, 0, DEFAULT, O_RDWR|O_CREAT, USR_RW)
}

// func Open(name const *C.char, hash_size, tdb_flags, open_flags C.int, mode C.mode_t) *C.struct_tdb_context
// so it beguns...
func Open(path string, hash_size, tdb_flags, open_flags C.int, mode C.mode_t) (db *DB) {
	name := C.CString(path)
	defer C.free(unsafe.Pointer(name))
	return &DB{path, false,
		C.tdb_open(name, hash_size, tdb_flags, open_flags, mode)}
}

// func Close() C.int
// so it unds...
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
