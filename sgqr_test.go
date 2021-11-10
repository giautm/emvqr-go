package sgqr_test

import (
	"testing"

	"giautm.dev/sgqr"
)

func Test_BuildPayload(t *testing.T) {
	type args struct {
		pairs sgqr.Array
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "VietQR",
			args: args{
				pairs: sgqr.Array{
					sgqr.PayloadFormatIndicator(),
					sgqr.PointOfInitiationMethod(false),
					sgqr.NewArray("38",
						sgqr.NewString("00", "A000000727"),
						sgqr.NewArray("01",
							sgqr.NewString("00", "970415"),
							sgqr.NewString("01", "113366668888"),
						),
						sgqr.NewString("02", "QRIBFTTA"),
					),
					sgqr.TransactionCurrency("704"),
					sgqr.CountryCode("VN"),
				},
			},
			want: "00020101021138560010A0000007270126000697041501121133666688880208QRIBFTTA53037045802VN6304F443",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := sgqr.BuildPayload(tt.args.pairs); got != tt.want {
				t.Errorf("BuildPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}
