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
//
// ZIP Codes -- lookup city by ZIP, get ZIP's lng/lat
// Uses data from http://www.boutell.com/zipcodes/

package zipcode

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
)

var (
	ErrInvalidZIP = errors.New("invalid ZIP Code")
	ErrZIPExist   = errors.New("ZIP Code doesn't exist")
)

type ZIPCode struct {
	Code  string
	City  string
	State string
	Lat   float64
	Lng   float64
	TZ    string
	DST   bool
}

const (
	colZIP = iota
	colCity
	colState
	colLatitude
	colLongitude
	colTimezone
	colDST
)

type DB map[string]ZIPCode

func NewDB(r io.Reader) (DB, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = 7
	var (
		code string
		lat, lng float64
		err error
		row []string
	)
	zdb := make(DB)
	for row, err = cr.Read(); err == nil; row, err = cr.Read() {
		if row[colZIP] == "zip" { // first row
			continue
		}
		if lat, err = strconv.ParseFloat(row[colLatitude], 64); err != nil {
			return nil, err
		}
		if lng, err = strconv.ParseFloat(row[colLongitude], 64); err != nil {
			return nil, err
		}
		code = row[colZIP]
		zdb[code] = ZIPCode{
			Code:  code,
			City:  row[colCity],
			State: row[colState],
			Lat:   lat,
			Lng:   lng,
			TZ:    row[colTimezone],
			DST:   row[colDST] == "1",
		}
	}
	if err != io.EOF {
		return nil, err
	}
	return zdb, nil
}

func ValidateZIP(code string) bool {
	if len(code) != 5 {
		return false
	}
	for _, c := range code {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func (db DB) Lookup(code string) (*ZIPCode, error) {
	if !ValidateZIP(code) {
		return nil, ErrInvalidZIP
	}
	zip, ok := db[code]
	if !ok {
		return nil, ErrZIPExist
	}
	return &zip, nil
}
