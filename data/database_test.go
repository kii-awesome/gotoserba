package data_test

import (
	"testing"

	"github.com/kii-awesome/gotoserba/data"
	"github.com/stretchr/testify/assert"
)

var db = data.ConnectDb()
func TestConnectDb(t *testing.T) {
	assert.NotNil(t, db)
}
