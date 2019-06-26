package dbs

import (
	"github.com/teawel/app/utils"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestDb_genId(t *testing.T) {
	db := new(DB)

	wg := sync.WaitGroup{}
	count := 1000
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			id := db.genId(time.Now())
			if !strings.HasSuffix(id, "0") {
				t.Log(id)
			}
		}()
	}
	wg.Wait()
}

func TestDB_Write(t *testing.T) {
	db, err := NewDB(os.TempDir() + "test.db")
	if err != nil {
		t.Fatal(err)
	}

	err = db.Write(map[string]string{
		"hello": "world",
		"name":  "lu",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDB_Query(t *testing.T) {
	db, err := NewDB(os.TempDir() + "test.db")
	if err != nil {
		t.Fatal(err)
	}

	result, err := db.Query([]string{utils.TimeFormat("Ymd")}, `
map(function (k, v) {
	v.value.name = "Hello";
	return v;
}).
result("name")
`)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(utils.JSONEncodePretty(result)))
}

func TestOtto(t *testing.T) {
	if scriptVM == nil {
		t.Fatal("scriptVM should not be nil")
	}

	before := time.Now()
	records := []*Record{}
	for i := 0; i < 100; i++ {
		records = append(records, &Record{
			Time: time.Unix(1561432133, 0),
			Key:  []byte(strconv.Itoa(i)),
			Value: map[string]string{
				"hello": "world",
				"name":  "name" + strconv.Itoa(i),
				"age":   strconv.Itoa(i),
			},
		})
	}

	v, err := evalScript(records, `
			filter(function (k, v) {
				return k % 3 == 0;
			})
			.map(function (k, v) {
				v.value.name = "[" + v.value.name + "]";
				return v;
			})
			/**.label(function (v) {
				var date = v.time;
				return date.getFullYear() + "-" + (date.getMonth() + 1) + "-" + date.getDate() + " " + date.getHours() + ":" + date.getMinutes() + ":" + date.getSeconds();
			})**/	
			.result("age")`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(utils.JSONEncode(v)))

	t.Log(time.Since(before).Seconds()*1000, "ms")
}
