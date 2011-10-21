// Copyright 2011 AUTHORS. All rights reserved.
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

// Package gotdb implements Go bindings to C libtdb (Trivial DataBase).
//
// TDB is a filesystem or in-memory key/value store in the vein of
// (G/B)DB/M API, a common DB abstraction layer used by the SAMBA project.
package tdb

// #cgo LDFLAGS: -ltdb
// #cgo pkg-config: tdb
// #include <stdlib.h> //free...
// #include <tdb.h>
import "C"
import "unsafe"

// ns is internal NameSpace.
var ns map[string]*db

// DB type is a pointer wrapping exposed to the final user of the library.
// The aim being prevention of limbs being shot off.
type DB struct {
	db *db // hiding bleak reality
}

// convenience typedef.
type tdb_CTX *C.struct_TDB_CONTEXT

// db type is an actual data structure holding pertinent metadata.
type db struct {
	pth *string // path name
	dbg bool    // to DEBUG or not to DEBUG?
	cld bool    // if it's closed (testing ctx's behaviour for now)
	ctx tdb_CTX // handle me gently
}

type Error struct{ msg string }

func init() {
	ns = make(map[string]*db)
}

// String returns string representation of db struct underlying DB.
func (o DB) String() (s string) {
	s = "db{pth:\"" + *o.db.pth + "\""
	if o.db.dbg {
		s += ", Dbg:true"
	} else {
		s += ", Dbg:false"
	}
	if o.db.cld {
		s += ", cld:true"
	} else {
		s += ", cld:false"
	}
	if o.db.ctx == nil {
		s += ", ctx:#f}"
	} else {
		s += ", ctx:#t}"
	}
	return
}

// New given a string representation of a path name always returns DB value
// along with Error status. In case of the latter being non-nil the former
// is probably unusable, should be considered "closed" (see further) and
// can be safely discarded.
//
// It's inadvisable to attempt opening already opened paths unless previous
// initial attempts failed and one considers conditions suitably improved.
// New will return the same DB value connected with already "touched" path.
// And although Go will prevent one from rebinding variable containing DB
// instance to a new name one can contravene this limitation by calling
// New and binding its return value to a freshly declared variable name. One
// should feel dully warned.
//
// Performing successful Close on any of the various DB instances of the
// same, unique path will thereafter cause any operation on them to fail with
// ERR_EINVAL status, hopefully only until another successful New or Open is
// executed...?
//
// At the moment above "functionality" is still under developmental
// investigation.
func New(path string) (DB, *Error) {
	return Open(path, 0, DEFAULT, O_RDWR|O_CREAT, USR_RW)
}

// Open is used by New with some reasonable default initial values apart from
// path name. Following is a signature of libtdb's original C tdb_open() function
// written in cgo convention:
//
// func tdb_open(name const *C.char, hash_size, tdb_flags, open_flags C.int, mode C.mode_t) *C.struct_tdb_context
func Open(path string, hash_size, tdb_flags, open_flags int, mode uint32) (DB, *Error) {
	name := C.CString(path)
	defer C.free(unsafe.Pointer(name))
	var ctx tdb_CTX
	if old := ns[path]; old != nil { // now, what do we do?
		// if db is still "here" in the ns but closed we
		if old.cld {
			ctx = C.tdb_open(name, C.int(hash_size), C.int(tdb_flags), C.int(open_flags), C.mode_t(mode))
			if ctx == nil {
				return DB{old}, &Error{"tdb_open failed"}
			} else {
				old.cld = false
				old.ctx = ctx
				return DB{old}, nil
			}
			// if it's not closed perhaps we should to something "more"
			// intelligent, like closing and reopening with new params
			// TODO: later?
		} else {
			return DB{old}, nil
		}
	} else {
		var fresh *db
		ctx = C.tdb_open(name, C.int(hash_size), C.int(tdb_flags), C.int(open_flags), C.mode_t(mode))
		if ctx == nil {
			println("Open() new ctx == nil")
			fresh = &db{&path, false, true, ctx}
			ns[path] = fresh
			return DB{fresh}, &Error{"tdb_open failed"}
		} else {
			fresh = &db{pth: &path, cld: false, dbg: false, ctx: ctx}
			ns[path] = fresh
			return DB{fresh}, nil
		}
	}
	panic("unreachable")
	// return &DB{path, false, ctx}
}

// Close calls tdb_close() on the C ctx pointer contained in DB struct,
// rendering it invalid in all other instances of the same path name (see New).
// Only on success does it return nil Error along with integer SUCCESS status.
// And here is trivially meaningless cgo signature of the original C function:
//
// func tdb_close() C.int
func (o DB) Close() (int, *Error) {
	dbg := o.db.dbg
	if dbg {
		println("tdb.Close()", o.String())
	}
	if o.db.cld {
		if dbg {
			println("tdb.Close()", "db.ctx =", o.db.ctx)
		}
		return ERR_EINVAL, &Error{"already closed"}
	}
	var status = int(C.tdb_close(o.db.ctx))
	if dbg {
		println("tdb.Close()", "tdb_close() ->", status)
	}
	if status == 0 {
		o.db.cld = true
		// for now, while testing let us hold on with that one
		// o.db.ctx = nil // argh! this does not stack up!
		return status, nil
	}
	// TODO: extract proper error string
	return status, &Error{"non-zero tdb_close status"}
}

// Debug toggles debugging setting on/off. One must be careful not to
// become casualty of the schizophrenia of detoggling this setting via
// different variable instances of the same DB.
func (o DB) Debug() {
	if o.db.dbg {
		o.db.dbg = false
	} else {
		o.db.dbg = true
	}
}

// Local Variables:
// mode: Go
// End:
