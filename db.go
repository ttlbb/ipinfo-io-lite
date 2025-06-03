package ipinfo_io_lite

import (
	"github.com/oschwald/maxminddb-golang/v2"
	"net/netip"
	"os"
)

type LiteDb struct {
	reader *maxminddb.Reader
}

func (db *LiteDb) loadFile(dst string) error {
	all, err := os.ReadFile(dst)
	if err != nil {
		return err
	}

	db.reader, err = maxminddb.FromBytes(all)
	if err != nil {
		return err
	}

	return nil
}

func (db *LiteDb) Query(addr netip.Addr) (row DbRecord, err error) {
	result := db.reader.Lookup(addr)
	row.Prefix = result.Prefix()
	err = result.Decode(&row)
	return
}

func NewDb(dst string) (*LiteDb, error) {
	db := &LiteDb{}
	if e := db.loadFile(dst); e != nil {
		return nil, e
	} else {
		return db, nil
	}
}
