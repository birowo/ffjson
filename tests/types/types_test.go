/**
 *  Copyright 2014 Paul Querna
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package types

import (
	"encoding/json"
	ff "github.com/pquerna/ffjson/tests/types/ff"
	"reflect"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	var record ff.Everything
	var recordTripped ff.Everything
	ff.NewEverything(&record)

	buf1, err := json.Marshal(&record)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	err = json.Unmarshal(buf1, &recordTripped)
	if err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	good := reflect.DeepEqual(record.FooStruct, recordTripped.FooStruct)
	if !good {
		t.Fatalf("Expected: %v\n Got: %v", *record.FooStruct, *recordTripped.FooStruct)
	}

	record.FooStruct = nil
	recordTripped.FooStruct = nil

	good = reflect.DeepEqual(record, recordTripped)
	if !good {
		t.Fatalf("Expected: %v\n Got: %v", record, recordTripped)
	}
}
