package dmImpl

import (
	"testing"
)

func Test_postgresConnection_GetDatabase(t *testing.T) {

	_, err := NewPostgresConnection().GetDatabase()

	if err != nil {
		t.Errorf("got error %s", err.Error())
	}

}
