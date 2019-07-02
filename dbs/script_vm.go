package dbs

import (
	"errors"
	"github.com/robertkrimen/otto"
	"github.com/teawel/app/utils"
	"log"
)

var scriptVM *otto.Otto

func setupScriptVM() {
	scriptVM = otto.New()

	_, err := scriptVM.Run(`
function Query(records) {
	records.forEach(function (v) {
		v.time = new Date(v.time * 1000);
	});

	this.filter = function (fn) {
		if (fn != null && typeof(fn) == "function") {
			records = records.filter(function (v, k) {
				return fn.call(this, k, v);
			});
		}
		return this;
	};
	
	this.map = function (fn) {
		if (fn != null && typeof(fn) == "function") {
			records = records.map(function (v, k) {
				return fn.call(this, k, v);
			});
		}
		return this;
	};

	this.label = function (fn) {
		if (fn != null && typeof(fn) == "function") {
			records = records.map(function (v, k) {
				v["$label"] = fn.call(this, v);
				return v;
			});
		}
		return this;
	};

	this.result = function (name) {
		return records.map(function (v, k) {
			if (typeof(v["$label"]) == "undefined") {
				var hour = v.time.getHours().toString();
				var minutes = v.time.getMinutes().toString();
				if (hour.length == 1) {
					hour = "0" + hour;
				}
				if (minutes.length == 1) {
					minutes = "0" + minutes;
				}
				v["$label"] = hour + ":" + minutes;
			}
			if (name instanceof Array) {
				var resultValues = [];
				name.forEach(function (field) {
					resultValues.push(v.value[field]);
				});
				return {
					"value": resultValues,
					"label": v["$label"]
				};
			} else {
				return {
					"value": v.value[name],
					"label": v["$label"]
				};
			}
		});
	};
}
`)
	if err != nil {
		log.Println("[error]scriptVM: " + err.Error())
	}
}

func evalScript(records []*Record, script string) (interface{}, error) {
	if scriptVM == nil {
		return nil, errors.New("scriptVm should not be nil")
	}
	recordList := []map[string]interface{}{}
	for _, record := range records {
		recordList = append(recordList, map[string]interface{}{
			"time":  record.Time.Unix(),
			"key":   string(record.Key),
			"value": record.Value,
		})
	}
	v, err := scriptVM.Eval(`
	(function () {
		return new Query(` + string(utils.JSONEncode(recordList)) + `).` + script + `
	})()`)
	if err != nil {
		return nil, err
	}
	return v.Export()
}
