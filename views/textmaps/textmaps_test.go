package textmaps

import (
	"reflect"
	"testing"
)

func TestGetMap(t *testing.T) {
	Init()
	found := GetMap(Main)

	if !reflect.DeepEqual(found, MainMap) {
		t.Error("FAIL: Map found was not MainMap")
	}
}

func TestUpdateMap(t *testing.T) {
	var newTitle = "Hello test worked"
	result := UpdateMap(Main, "Title", newTitle)

	if result["Title"] != newTitle {
		t.Errorf("FAIL: UpdateMap failed, recieved, %s", result["Title"])
	}
}
