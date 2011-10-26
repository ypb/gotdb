WTF is it?
==========

[TDB](http://tdb.samba.org/) [golang](http://golang.org) bindings.

## goInstall

    (or (eq?
         (should!
          hopefully
          (work (goinstall github.com/ypb/gotdb)))
         0)
        (sh (&& (map make '(clean exp))
                 ./exp)))

## Done (see godocs)

    tdb_open, _close, _store, _fetch()

TODO (still)
====

    tdb_exists
       _first/nextkey
       _delete
       _append?
       _get/add/remove_flags
       _summary!
       _name _hash_size _map_size _get/enable_seqnum...
       _reopen/_all? _fd?
       _open_ex(...)
       _wipe_all/_repack

* definitely improve on errors/(debug/log)ing;
* traversals;
* locking and transactions... who needs that? moi, non
