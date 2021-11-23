package emvqr_test

import (
	"testing"

	"giautm.dev/emvqr"
)

func Test_BuildPayload(t *testing.T) {
	type args struct {
		pairs []emvqr.Pair
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
				pairs: []emvqr.Pair{
					emvqr.PayloadFormatIndicator(),
					emvqr.PointOfInitiationMethod(true),
					emvqr.AdditionalData(emvqr.String("08", "tien le")),
					emvqr.List("38",
						emvqr.String("00", "A000000727"),
						emvqr.List("01",
							emvqr.String("00", "970415"),
							emvqr.String("01", "113366668888"),
						),
						emvqr.String("02", "QRIBFTTA"),
					),
					emvqr.TransactionCurrency("704"),
					emvqr.TransactionAmountUint(6000),
					emvqr.CountryCode("VN"),
				},
			},
			want: "00020101021262110807tien le38560010A0000007270126000697041501121133666688880208QRIBFTTA5303704540460005802VN63046893",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := emvqr.BuildPayload(tt.args.pairs...)
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
	pairs := []emvqr.Pair{
		emvqr.PayloadFormatIndicator(),
		emvqr.PointOfInitiationMethod(false),
		emvqr.List("38",
			emvqr.String("00", "A000000727"),
			emvqr.List("01",
				emvqr.String("00", "970415"),
				emvqr.String("01", "113366668888"),
			),
			emvqr.String("02", "QRIBFTTA"),
		),
		emvqr.TransactionCurrency("704"),
		emvqr.CountryCode("VN"),
	}

	for n := 0; n < b.N; n++ {
		emvqr.BuildPayload(pairs...)
	}
}
