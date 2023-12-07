package tests

import (
	"github.com/Kosmosman/service/orderdb"
	"github.com/Kosmosman/service/types"
	"testing"
)

func TestAddGetDB(t *testing.T) {
	db := orderdb.OrderDB{}
	db.Connect()
	db.ClearDB()
	testcases := map[string]struct {
		uid  string
		data string
		want string
	}{
		"add1": {uid: "1", data: "SomeData", want: "SomeData"},
		"add2": {uid: "2", data: "AnotherData", want: "AnotherData"},
		"add3": {uid: "3", data: "CommonData", want: "CommonData"},
		"add4": {uid: "4", data: "UncommonData", want: "UncommonData"},
		"add5": {uid: "5", data: "MaybeCommonData", want: "MaybeCommonData"},
		"add6": {uid: "6", data: "WhatAData", want: "WhatAData"},
	}

	for testname, data := range testcases {
		db.Add(&data.uid, &data.data)
		if db.Get(&data.uid) != data.want {
			t.Errorf("Incorrect answer in test %s\n", testname)
		}
	}

	testcasesIncorrectData := map[string]struct {
		uid  string
		want string
	}{
		"incorrect_add1": {uid: "11", want: ""},
		"incorrect_add2": {uid: "165", want: ""},
		"incorrect_add3": {uid: "rgdfbdf", want: ""},
		"incorrect_add4": {uid: "53", want: ""},
		"incorrect_add5": {uid: "3qref4w3", want: ""},
		"incorrect_add6": {uid: "efadvxfbht", want: ""},
	}

	for testname, data := range testcasesIncorrectData {
		if db.Get(&data.uid) != data.want {
			t.Errorf("Incorrect answer in test %s\n", testname)
		}
	}
	db.ClearDB()
}

func TestRestoreCache(t *testing.T) {
	var cache types.Cache
	db := orderdb.OrderDB{}
	db.Connect()
	db.ClearDB()
	testcases := map[string]struct {
		uid  string
		data string
	}{
		"test1": {"1", "Data1"},
		"test2": {"2", "Data2"},
		"test3": {"3", "Data3"},
		"test4": {"4", "Data4"},
		"test5": {"5", "Data5"},
		"test6": {"6", "Data6"},
		"test7": {"7", "Data7"},
	}
	for _, order := range testcases {
		db.Add(&order.uid, &order.data)
	}
	db.RestoreCache(&cache)
	for testname, data := range testcases {
		if dataByte, ok := cache.Data[data.uid]; !ok {
			t.Errorf("Incorrect ansver in test %s\n", testname)
		} else {
			if string(dataByte[1:len(dataByte)-1]) != data.data {
				t.Errorf("Incorrect ansver in test %s: expected \"%s\", result is \"%s\"\n", testname, data.data, string(dataByte))
			}
		}
	}
	db.ClearDB()
}
