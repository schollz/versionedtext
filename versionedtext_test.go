package versionedtext

import (
	"encoding/json"
	"testing"
)

func TestGeneral(t *testing.T) {
	d := NewVersionedText("A word")
	d.Update("A word and adding something at the end")
	d.Update("A (deleted) and adding something at the end, with another addition")
	bJSON, err := json.Marshal(d)
	if err != nil {
		t.Error(err)
	}

	var d2 VersionedText
	err = json.Unmarshal(bJSON, &d2)
	if err != nil {
		t.Error(err)
	}
	if d2.GetCurrent() != d.GetCurrent() {
		t.Errorf("Problem reloading from JSON")
	}

	snapshots := d.GetSnapshots()
	previousText, err := d.GetPreviousByTimestamp(snapshots[1])
	if err != nil {
		t.Error(err)
	}
	if previousText != "A word and adding something at the end" {
		t.Errorf("Did not reconstruct properly")
	}
}
