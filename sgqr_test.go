package sgqr_test

import (
	"testing"

	"giautm.dev/sgqr"
)

func Test_BuildPayload(t *testing.T) {
	type args struct {
		pairs []sgqr.Pair
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "VietQR",
			args: args{
				pairs: []sgqr.Pair{
					sgqr.PayloadFormatIndicator(),
					sgqr.PointOfInitiationMethod(true),
					sgqr.AdditionalData(sgqr.String("08", "tien le")),
					sgqr.Array("38",
						sgqr.String("00", "A000000727"),
						sgqr.Array("01",
							sgqr.String("00", "970415"),
							sgqr.String("01", "113366668888"),
						),
						sgqr.String("02", "QRIBFTTA"),
					),
					sgqr.TransactionCurrency("704"),
					sgqr.TransactionAmountUint(6000),
					sgqr.CountryCode("VN"),
				},
			},
			want: "00020101021138560010A0000007270126000697041501121133666688880208QRIBFTTA53037045802VN6304F443",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sgqr.BuildPayload(tt.args.pairs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuildPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_BuildPayload(b *testing.B) {
	pairs := []sgqr.Pair{
		sgqr.PayloadFormatIndicator(),
		sgqr.PointOfInitiationMethod(false),
		sgqr.Array("38",
			sgqr.String("00", "A000000727"),
			sgqr.Array("01",
				sgqr.String("00", "970415"),
				sgqr.String("01", "113366668888"),
			),
			sgqr.String("02", "QRIBFTTA"),
		),
		sgqr.TransactionCurrency("704"),
		sgqr.CountryCode("VN"),
	}

	for n := 0; n < b.N; n++ {
		sgqr.BuildPayload(pairs...)
	}
}
