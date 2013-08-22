package zipcode

import (
	"bytes"
	"testing"
)

func TestZIPCodeDB(t *testing.T) {
	var testData = `"zip","city","state","latitude","longitude","timezone","dst"
"21230","Baltimore","MD","39.273107","-76.62613","-5","1"
"21231","Baltimore","MD","39.288607","-76.59116","-5","1"
"21233","Baltimore","MD","39.284707","-76.620489","-5","1"
"21234","Parkville","MD","39.385006","-76.54177","-5","1"
"21235","Baltimore","MD","39.284707","-76.620489","-5","1"
"21236","Nottingham","MD","39.389457","-76.48709","-5","1"
"21237","Rosedale","MD","39.33224","-76.50365","-5","1"
`
	zdb, err := NewDB(bytes.NewReader([]byte(testData)))
	if err != nil {
		t.Fatalf("creating ZIP Code db: %v", err)
	}
	if len(zdb) != 7 {
		t.Errorf("db: want length %d, got %d", 7, len(zdb))
	}
	zc, err := zdb.Lookup("21230")
	if err != nil {
		t.Fatalf("lookup: %v", err)
	}
	if zc.Code != "21230" {
		t.Errorf("lookup: want %s, got %s", "21230", zc.Code)
	}
	t.Logf("found ZIP Code data for %s: %+v", "21230", zc)
}
