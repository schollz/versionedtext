package versionedtext

import (
	"encoding/json"
	"strconv"
	"testing"
)

func BenchmarkUpdate(b *testing.B) {
	d := NewVersionedText("A word")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Update(strconv.Itoa(i))
	}
}

func BenchmarkRebuild500thOf1000(b *testing.B) {
	d := NewVersionedText("A word")
	for i := 0; i < 1001; i++ {
		d.Update(strconv.Itoa(i))
	}
	snapshots := d.GetSnapshots()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.GetPreviousByTimestamp(snapshots[500])
	}
}

func BenchmarkRebuild1000thOf1000(b *testing.B) {
	d := NewVersionedText("A word")
	for i := 0; i < 1001; i++ {
		d.Update(strconv.Itoa(i))
	}
	snapshots := d.GetSnapshots()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.GetPreviousByTimestamp(snapshots[1000])
	}
}

func BenchmarkRebuild10000thof10000(b *testing.B) {
	d := NewVersionedText("A word")
	for i := 0; i < 10001; i++ {
		d.Update(strconv.Itoa(i))
	}
	snapshots := d.GetSnapshots()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.GetPreviousByTimestamp(snapshots[10000])
	}
}

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
