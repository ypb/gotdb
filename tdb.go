// Copyright 2011 AUTHORS. All rights reserved.
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

// Package gotdb implements C bindings to the tdb <http://tdb.samba.org/>
// (Trivial Data Base) library.
//
// TDB is a filesystem or in-memory key/value store in the vein of
// (G/B)DB/M API, and a common DB abstraction layer used by the SAMBA project.
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

// Error typo humun. Possible Errno values:
//
//  /* error codes */
//  SUCCESS
//  ERR_CORRUPT
//  ERR_IO
//  ERR_LOCK
//  ERR_OOM
//  ERR_EXISTS
//  ERR_NOLOCK
//  ERR_LOCK_TIMEOUT
//  ERR_NOEXIST
//  ERR_EINVAL
//  ERR_RDONLY
//  ERR_NESTING
//
type Error interface {
	String() string
	Errno() int
}

// error type... Weell, u naw no.
type error struct {
	ret int
	msg *string
}

func (e *error) String() string {
	if e == nil {
		return ""
	}
	return *e.msg
}
// TOPONDER: perhaps we'd like to make Error "nil-safe" (breaking convention? /hmmm)
func (e *error) Errno() int {
	if e == nil {
		return SUCCESS
	}
	return e.ret
}

func mkError(sts int, msg string) Error {
	// funcional overkill but we are guessing well be needing semi-complex
	// machinationarly of decoding internal tdb strings... or not as pkg/os
	// may provide errno out of the box, lazy mofo
	return &error{sts, &msg}
}

func init() {
	ns = make(map[string]*db)
}

// String returns string representation of db struct underlying DB.
func (file DB) String() string {
	var s = "db{pth: \"" + *file.db.pth + "\""
	if file.db.dbg {
		s += ", dbg: true"
	} else {
		s += ", dbg: false"
	}
	if file.db.cld {
		s += ", cld: true"
	} else {
		s += ", cld: false"
	}
	if file.db.ctx == nil {
		s += ", ctx: #f}"
	} else {
		s += ", ctx: #t}"
	}
	return s
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
func New(path string) (DB, Error) {
	return Open(path, 0, DEFAULT, O_RDWR|O_CREAT, USR_RW)
}

// Open is used by New with some reasonable default initial values apart from
// path name. Following is a signature of libtdb's original C tdb_open() function
// written in cgo convention:
//
// func tdb_open(name const *C.char, hash_size, tdb_flags, open_flags C.int, mode C.mode_t) *C.struct_tdb_context
//
//  /* tdb_flags */
//  DEFAULT           /* just a readability place holder */
//  CLEAR_IF_FIRST    /* beats me... */
//  INTERNAL          /* don't store on disk */
//  NOLOCK            /* don't do any locking */
//  NOMMAP            /* don't use mmap */
//  CONVERT           /* convert endian (internal use) */
//  BIGENDIAN         /* header is big-endian (internal use) */
//  NOSYNC            /* don't use synchronous transactions */
//  SEQNUM            /* maintain a sequence number */
//  VOLATILE          /* Activate the per-hashchain freelist, default 5 */
//  ALLOW_NESTING     /* Allow transactions to nest */
//  DISALLOW_NESTING  /* Disallow transactions to nest */
//  INCOMPATIBLE_HASH /* Better hashing: can't be opened by tdb < 1.2.6. */
//
//  /* open_flags */
//  /* 'man 2 open' on *nix, but what of pkg/os? TOPONDER */
//  O_RDONLY
//  /* O_WRONLY *//* is invalid */
//  O_RDWR
//  O_CREAT, O_TRUNC, O_APPEND
//  /* well, Ay dunno... */
//  O_CLOEXEC, O_EXCL
//  /* O_NOATIME *//* #define __USE_GNU */
//  O_NOFOLLOW, O_NONBLOCK = O_NDELAY
//
//  /* O_CREAT mode */
//  USR_RW = (S_IWUSR | S_IRUSR) /* helpful shortcut */
//  USR_RWX                      /* 00700 user (file owner) has read, write and execute permission */
//  USR_R                        /* 00400 user has read permission */
//  USR_W                        /* 00200 user has write permission */
//  USR_X                        /* 00100 user has execute permission */
//  GRP_RWX                      /* 00070 group has read, write and execute permission */
//  GRP_R                        /* 00040 group has read permission */
//  GRP_W                        /* 00020 group has write permission */
//  GRP_X                        /* 00010 group has execute permission */
//  OTH_RWX                      /* 00007 others have read, write and execute permission */
//  OTH_R                        /* 00004 others have read permission */
//  OTH_W                        /* 00002 others have write permission */
//  OTH_X                        /* 00001 others have execute permission */
//
func Open(path string, hash_size, tdb_flags, open_flags int, mode uint32) (DB, Error) {
	name := C.CString(path)
	defer C.free(unsafe.Pointer(name))
	var ctx tdb_CTX
	if old := ns[path]; old != nil { // now, what do we do?
		// if db is still "here" in the ns but closed we
		if old.cld {
			ctx = C.tdb_open(name, C.int(hash_size), C.int(tdb_flags), C.int(open_flags), C.mode_t(mode))
			if ctx == nil {
				return DB{old}, mkError(1, "tdb.Open() tdb_open old failed")
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
			return DB{fresh}, mkError(1, "tdb.Open() tdb_open fresh failed")
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
func (file DB) Close() Error {
	dbg := file.db.dbg
	if dbg {
		println("tdb.Close()", file.String())
	}
	if file.db.cld {
		if dbg {
			println("tdb.Close()", "db.ctx =", file.db.ctx)
		}
		return mkError(ERR_EINVAL, "tdb.Close() already closed")
	}
	var status = int(C.tdb_close(file.db.ctx))
	if dbg {
		println("tdb.Close()", "tdb_close() ->", status)
	}
	if status == SUCCESS {
		file.db.cld = true
		// for now, while testing let us hold on with that one
		// file.db.ctx = nil // argh! this does not stack up!
		return nil
	}
	// TODO: extract proper error string
	return mkError(status, "tdb.Close() SUCCESS not")
}

// Debug toggles debugging setting on/off. One must be careful not to
// become casualty of the schizophrenia of detoggling this setting via
// different variable instances of the same DB.
func (file DB) Debug() {
	if file.db.dbg {
		file.db.dbg = false
	} else {
		file.db.dbg = true
	}
}

// Local Variables:
// mode: Go
// End:
