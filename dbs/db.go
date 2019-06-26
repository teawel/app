package dbs

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/teawel/app/utils"
	"log"
	"strconv"
	"sync"
	"time"
)

var lastId int64 = 0
var idLocker = sync.Mutex{}

type DB struct {
	db *leveldb.DB
}

func NewDB(path string) (*DB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}

// write value
// value,$time,$id -> json($value)
func (this *DB) Write(value map[string]string) error {
	if this.db == nil {
		return errors.New("db pointer should not be nil")
	}
	t := time.Now()
	key := "value," + utils.TimeFormat("YmdHis") + "," + this.genId(t)
	return this.db.Put([]byte(key), utils.JSONEncode(value), nil)
}

func (this *DB) Delete(key []byte) error {
	if this.db == nil {
		return errors.New("db pointer should not be nil")
	}
	return this.db.Delete(key, nil)
}

// query values
func (this *DB) Query(timePrefixes []string, script string) (result interface{}, err error) {
	if this.db == nil {
		return nil, errors.New("db pointer should not be nil")
	}

	snapshot, err := this.db.GetSnapshot()
	if err != nil {
		return nil, err
	}

	records := []*Record{}
	for _, timePrefix := range timePrefixes {
		it := snapshot.NewIterator(util.BytesPrefix([]byte("value,"+timePrefix)), nil)
		for it.Next() {
			record := new(Record)
			record.Key = it.Key()
			pieces := bytes.Split(it.Key(), []byte(","))
			nano, err := strconv.ParseInt(string(pieces[2]), 10, 64)
			if err != nil {
				log.Println("[error]" + err.Error())
				this.db.Delete(it.Key(), nil)
				continue
			}

			err = json.Unmarshal(it.Value(), &record.Value)
			if err != nil {
				log.Println("[error]" + err.Error())
				this.db.Delete(it.Key(), nil)
				continue
			}

			record.Time = time.Unix(nano/int64(time.Second), nano%int64(time.Second))
			records = append(records, record)
		}
		it.Release()
	}

	if len(records) > 0 {
		if len(script) > 0 {
			return evalScript(records, script)
		} else {
			recordList := []map[string]interface{}{}
			for _, record := range records {
				recordList = append(recordList, map[string]interface{}{
					"time":  record.Time.Unix(),
					"key":   string(record.Key),
					"value": record.Value,
				})
			}
			return recordList, nil
		}
	}

	return
}

func (this *DB) Close() error {
	if this.db == nil {
		return nil
	}
	return this.db.Close()
}

func (this *DB) genId(t time.Time) string {
	idLocker.Lock()
	id := t.UnixNano()
	if lastId == id {
		id++
	}
	lastId = id
	idLocker.Unlock()

	return strconv.FormatInt(id, 10)
}
