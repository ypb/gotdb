/*
 Copyright 2011 AUTHORS. All rights reserved.
 Use of this source code is governed by a LGPL-style
 license that can be found in the LICENSE file.
*/

/* #include <sys/types.h> */
/* #include <sys/stat.h> */
#include <fcntl.h> /*HA!*/
#include <tdb.h>

enum DB_OPEN_F
{
  /* tdb_open() tdb_flags */
  $DEFAULT = TDB_DEFAULT, /* just a readability place holder */
  $CLEAR_IF_FIRST = TDB_CLEAR_IF_FIRST,
  $INTERNAL = TDB_INTERNAL, /* don't store on disk */
  $NOLOCK = TDB_NOLOCK, /* don't do any locking */
  $NOMMAP = TDB_NOMMAP, /* don't use mmap */
  $CONVERT = TDB_CONVERT, /* convert endian (internal use) */
  $BIGENDIAN = TDB_BIGENDIAN, /* header is big-endian (internal use) */
  $NOSYNC = TDB_NOSYNC, /* don't use synchronous transactions */
  $SEQNUM = TDB_SEQNUM, /* maintain a sequence number */
  $VOLATILE = TDB_VOLATILE, /* Activate the per-hashchain freelist, default 5 */
  $ALLOW_NESTING = TDB_ALLOW_NESTING, /* Allow transactions to nest */
  $DISALLOW_NESTING = TDB_DISALLOW_NESTING, /* Disallow transactions to nest */
  $INCOMPATIBLE_HASH = TDB_INCOMPATIBLE_HASH, /* Better hashing: can't be opened by tdb < 1.2.6. */
};

enum DB_STORE_F
{
  /* tdb_store() flags */
  $REPLACE = TDB_REPLACE,    /* Unused */
  $INSERT = TDB_INSERT,      /* Don't overwrite an existing entry */
  $MODIFY = TDB_MODIFY,      /* Don't create an existing entry    */
};

/* typedef enum TDB_ERROR $DB_ERROR; *//* write by hand? */

enum DB_ERROR
{
  /* error codes */
  $SUCCESS = TDB_SUCCESS,
  $ERR_CORRUPT = TDB_ERR_CORRUPT,
  $ERR_IO = TDB_ERR_IO,
  $ERR_LOCK = TDB_ERR_LOCK,
  $ERR_OOM = TDB_ERR_OOM,
  $ERR_EXISTS = TDB_ERR_EXISTS,
  $ERR_NOLOCK = TDB_ERR_NOLOCK,
  $ERR_LOCK_TIMEOUT = TDB_ERR_LOCK_TIMEOUT,
  $ERR_NOEXIST = TDB_ERR_NOEXIST,
  $ERR_EINVAL = TDB_ERR_EINVAL,
  $ERR_RDONLY = TDB_ERR_RDONLY,
  $ERR_NESTING = TDB_ERR_NESTING,
};

enum debug_level
{
  /* debugging uses one of the following levels */
  $DEBUG_FATAL = TDB_DEBUG_FATAL,
  $DEBUG_ERROR = TDB_DEBUG_ERROR,
  $DEBUG_WARNING = TDB_DEBUG_WARNING,
  $DEBUG_TRACE = TDB_DEBUG_TRACE,
};

typedef TDB_DATA $DATA;
/* oh dear... */
/* typedef TDB_CONTEXT* $CONTEXT; */

enum open_flag
{
  /* unix standard like in pkg/os but what of other plats? */
  $O_RDONLY = O_RDONLY,
  /* $O_WRONLY = O_WRONLY, *//* is invalid */
  $O_RDWR = O_RDWR,
  $O_CREAT = O_CREAT,
  $O_TRUNC = O_TRUNC,
  $O_APPEND = O_APPEND,
/* well, Ay dunno... */
  $O_CLOEXEC = O_CLOEXEC,
  $O_EXCL = O_EXCL,
  /* $O_NOATIME = O_NOATIME, *//* #define __USE_GNU */
  $O_NOFOLLOW = O_NOFOLLOW,
  $O_NONBLOCK = O_NONBLOCK,
  $O_NDELAY = O_NDELAY,
};

/* spurious as hell and a waste of time, butt...abstractly eedookayshanal */
enum mode
{
  /* mode_t constants */
  $USR_RW = (S_IWUSR | S_IRUSR), /* helpful shortcut */
  $USR_RWX = S_IRWXU, /* 00700 user (file owner) has read, write and execute permission */
  $USR_R = S_IRUSR,   /* 00400 user has read permission */
  $USR_W = S_IWUSR,   /* 00200 user has write permission */
  $USR_X = S_IXUSR,   /* 00100 user has execute permission */
  $GRP_RWX = S_IRWXG, /* 00070 group has read, write and execute permission */
  $GRP_R = S_IRGRP,   /* 00040 group has read permission */
  $GRP_W = S_IWGRP,   /* 00020 group has write permission */
  $GRP_X = S_IXGRP,   /* 00010 group has execute permission */
  $OTH_RWX = S_IRWXO, /* 00007 others have read, write and execute permission */
  $OTH_R = S_IROTH,   /* 00004 others have read permission */
  $OTH_W = S_IWOTH,   /* 00002 others have write permission */
  $OTH_X = S_IXOTH,   /* 00001 others have execute permission */
};
