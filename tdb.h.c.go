// godefs -g tdb tdb.h.c

// MACHINE GENERATED - DO NOT EDIT.

package tdb

// Constants
const (
	DEFAULT           = 0
	CLEAR_IF_FIRST    = 0x1
	INTERNAL          = 0x2
	NOLOCK            = 0x4
	NOMMAP            = 0x8
	CONVERT           = 0x10
	BIGENDIAN         = 0x20
	NOSYNC            = 0x40
	SEQNUM            = 0x80
	VOLATILE          = 0x100
	ALLOW_NESTING     = 0x200
	DISALLOW_NESTING  = 0x400
	INCOMPATIBLE_HASH = 0x800
	REPLACE           = 0x1
	INSERT            = 0x2
	MODIFY            = 0x3
	SUCCESS           = 0
	ERR_CORRUPT       = 0x1
	ERR_IO            = 0x2
	ERR_LOCK          = 0x3
	ERR_OOM           = 0x4
	ERR_EXISTS        = 0x5
	ERR_NOLOCK        = 0x6
	ERR_LOCK_TIMEOUT  = 0x7
	ERR_NOEXIST       = 0x8
	ERR_EINVAL        = 0x9
	ERR_RDONLY        = 0xa
	ERR_NESTING       = 0xb
	DEBUG_FATAL       = 0
	DEBUG_ERROR       = 0x1
	DEBUG_WARNING     = 0x2
	DEBUG_TRACE       = 0x3
	O_RDONLY          = 0
	O_RDWR            = 0x2
	O_CREAT           = 0x40
	O_TRUNC           = 0x200
	O_APPEND          = 0x400
	O_CLOEXEC         = 0x80000
	O_EXCL            = 0x80
	O_NOFOLLOW        = 0x20000
	O_NONBLOCK        = 0x800
	O_NDELAY          = 0x800
	USR_RW            = 0x180
	USR_RWX           = 0x1c0
	USR_R             = 0x100
	USR_W             = 0x80
	USR_X             = 0x40
	GRP_RWX           = 0x38
	GRP_R             = 0x20
	GRP_W             = 0x10
	GRP_X             = 0x8
	OTH_RWX           = 0x7
	OTH_R             = 0x4
	OTH_W             = 0x2
	OTH_X             = 0x1
)

// Types

type DATA struct {
	Dptr  *uint8
	Dsize uint32
}
