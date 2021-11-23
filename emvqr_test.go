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

func TestGetIn(t *testing.T) {
	type args struct {
		data string
		ids  []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "not found",
			args: args{
				data: "0103333",
				ids:  []string{"02"},
			},
			want: "",
		},
		{
			name: "invalid data",
			args: args{
				data: "010233",
				ids:  []string{"01", "02"},
			},
			want: "",
		},
		{
			name: "simple",
			args: args{
				data: "0103333",
				ids:  []string{"01"},
			},
			want: "333",
		},
		{
			name: "next",
			args: args{
				data: "000201010211",
				ids:  []string{"01"},
			},
			want: "11",
		},
		{
			name: "sub",
			args: args{
				data: "00020101021126280010A000000775011001064159995204829953037045802VN5913123VIETNAMESE6005HANOI610610000062290313G7AUTO03 SAPO0708G7AUTO04630458BA",
				ids:  []string{"62", "07"},
			},
			want: "G7AUTO04",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := emvqr.GetIn(tt.args.data, tt.args.ids...); got != tt.want {
				t.Errorf("GetIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_GetIn(b *testing.B) {
	data := "00020101021126280010A000000775011001064159995204829953037045802VN5913123VIETNAMESE6005HANOI610610000062290313G7AUTO03 SAPO0708G7AUTO04630458BA"
	want := "G7AUTO04"

	for n := 0; n < b.N; n++ {
		if emvqr.GetIn(data, "62", "07") != want {
			b.Errorf("GetIn() = %v, want %v", emvqr.GetIn(data, "62", "07"), want)
		}
	}
}
