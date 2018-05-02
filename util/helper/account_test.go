package helper

import (
	"testing"

	"github.com/irisnet/iris-sync-server/module/logger"
)

func setUp()  {
	InitClientPool()
}

func TestQueryAccountBalance(t *testing.T) {
	setUp()

	type args struct {
		address string
		delay   bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name:"the balance of address is not empty",
			args: struct {
				address string
				delay   bool
			}{
				address: "ADBC4AAB3A089BDC8A8155AB97E64CD2CF4A0E9F",
				delay: false},
		},
		{
			name:"the balance of address is empty",
			args: struct {
				address string
				delay   bool
			}{
				address: "ADBC4AAB3A089BDC8A8155AB97E64CD2CF4A0E9E",
				delay: false},
		},
	}
	for _, tt := range tests {
		account := QueryAccountBalance(tt.args.address, tt.args.delay)
		logger.Info.Printf("the balance of %s is %+v\n", tt.args.address, account)
	}
}