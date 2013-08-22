// Copyright (c) 2013 Public Good Software, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package zipcode

import (
	"bytes"
	"testing"
)

func TestZIPCodeDB(t *testing.T) {
	var fixture = `"zip","city","state","latitude","longitude","timezone","dst"
"21230","Baltimore","MD","39.273107","-76.62613","-5","1"
"21231","Baltimore","MD","39.288607","-76.59116","-5","1"
"21233","Baltimore","MD","39.284707","-76.620489","-5","1"
"21234","Parkville","MD","39.385006","-76.54177","-5","1"
"21235","Baltimore","MD","39.284707","-76.620489","-5","1"
"21236","Nottingham","MD","39.389457","-76.48709","-5","1"
"21237","Rosedale","MD","39.33224","-76.50365","-5","1"
`
	zdb, err := NewDB(bytes.NewReader([]byte(fixture)))
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
