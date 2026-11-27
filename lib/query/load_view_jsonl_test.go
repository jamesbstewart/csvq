package query

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mithrandie/csvq/lib/file"
	"github.com/mithrandie/csvq/lib/option"
	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"
)

func TestLoadJSONLinesBeyondInitialCapacity(t *testing.T) {
	const count = fileLoadingPreparedRecordSetCap * 4
	var input strings.Builder
	for index := 0; index < count; index++ {
		fmt.Fprintf(&input, "{\"id\":\"%d\",\"payload\":\"data\"}\n", index)
	}
	reader, err := file.NewReader(strings.NewReader(input.String()), 2048)
	if err != nil {
		t.Fatal(err)
	}
	info := &FileInfo{Path: "records.jsonl"}
	view, err := loadViewFromJsonLinesFile(context.Background(), &option.Flags{CPU: 2}, reader, info, parser.Identifier{Literal: info.Path})
	if err != nil {
		t.Fatal(err)
	}
	if want := NewHeader("records", []string{"id", "payload"}); !reflect.DeepEqual(view.Header, want) {
		t.Fatalf("header = %v, want %v", view.Header, want)
	}
	if len(view.RecordSet) != count {
		t.Fatalf("read %d records, want %d", len(view.RecordSet), count)
	}
	for index, record := range view.RecordSet {
		want := NewRecord([]value.Primary{value.NewString(fmt.Sprint(index)), value.NewString("data")})
		if !reflect.DeepEqual(record, want) {
			t.Fatalf("record %d = %v, want %v", index, record, want)
		}
	}
}
