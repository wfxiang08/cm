// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabletserver

import "encoding/binary"

var pack = binary.BigEndian

const (
	RC_DELETED = 1

	// MAX_KEY_LEN is a value less than memcache's limit of 250.
	MAX_KEY_LEN = 200

	// MAX_DATA_LEN prevents large rows from being inserted in rowcache.
	MAX_DATA_LEN = 8000
)
