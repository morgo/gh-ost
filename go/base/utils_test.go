/*
   Copyright 2016 GitHub Inc.
	 See https://github.com/github/gh-ost/blob/master/LICENSE
*/

package base

import (
	"testing"
	"time"

	"github.com/openark/golib/log"
	test "github.com/openark/golib/tests"
)

func init() {
	log.SetLevel(log.ERROR)
}

func TestStringContainsAll(t *testing.T) {
	s := `insert,delete,update`

	test.S(t).ExpectFalse(StringContainsAll(s))
	test.S(t).ExpectFalse(StringContainsAll(s, ""))
	test.S(t).ExpectFalse(StringContainsAll(s, "drop"))
	test.S(t).ExpectTrue(StringContainsAll(s, "insert"))
	test.S(t).ExpectFalse(StringContainsAll(s, "insert", "drop"))
	test.S(t).ExpectTrue(StringContainsAll(s, "insert", ""))
	test.S(t).ExpectTrue(StringContainsAll(s, "insert", "update", "delete"))
}

func TestFindP90(t *testing.T) {
	times := []time.Duration{
		1 * time.Second,
		2 * time.Second,
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
		1 * time.Second,
		1 * time.Second,
		1 * time.Second,
		1 * time.Second,
		1 * time.Second,
	}
	test.S(t).ExpectEquals(findP90(times), 3*time.Second)
}
