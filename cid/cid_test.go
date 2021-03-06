// Copyright (C) 2014 Jakob Borg and other contributors. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package cid

import "testing"

func TestGet(t *testing.T) {
	m := NewMap()

	if i := m.Get("foo"); i != 1 {
		t.Errorf("Unexpected id %d != 1", i)
	}
	if i := m.Get("bar"); i != 2 {
		t.Errorf("Unexpected id %d != 2", i)
	}
	if i := m.Get("foo"); i != 1 {
		t.Errorf("Unexpected id %d != 1", i)
	}
	if i := m.Get("bar"); i != 2 {
		t.Errorf("Unexpected id %d != 2", i)
	}

	if LocalID != 0 {
		t.Error("LocalID should be 0")
	}
	if i := m.Get(LocalName); i != LocalID {
		t.Errorf("Unexpected id %d != %c", i, LocalID)
	}
}
