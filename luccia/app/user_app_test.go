package app

import "testing"

func TestCreateUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)
}
