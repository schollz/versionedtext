package versionedtext

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGeneral(t *testing.T) {
	d := NewVersionedText("A word")
	d.Update("A word and adding something at the end")
	d.Update("A (deleted) and adding something at the end, with another addition")

	// Test getting a snapshot
	snapshots := d.GetSnapshots()
	previousText, err := d.GetPreviousByTimestamp(snapshots[1])
	if err != nil {
		t.Error(err)
	}
	if previousText != "A word and adding something at the end" {
		t.Errorf("Did not reconstruct properly")
	}

	// Test putting them in a struct and marshaling
	type SomeStruct struct {
		SomeThing string
		SomeDiffs VersionedText
	}
	var d1 SomeStruct
	d1.SomeDiffs = d
	d1.SomeThing = "Some thing"
	bJSON, err := json.Marshal(d1)
	fmt.Println(string(bJSON))
	if err != nil {
		t.Error(err)
	}
	var d2 SomeStruct
	err = json.Unmarshal(bJSON, &d2)
	if err != nil {
		t.Error(err)
	}
	if d2.SomeDiffs.GetCurrent() != d.GetCurrent() {
		t.Errorf("Problem reloading from JSON")
	}

}
