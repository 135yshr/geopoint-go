package geopoint

import "testing"

func TestGeoPoint_String(t *testing.T) {
	type args struct {
		point [2]float64
	}
	tests := map[string]struct {
		args args
		want string
		err  error
	}{
		"東京駅のポイントを指定してPOINTを表す文字列を取得できること": {
			args: args{point: [2]float64{139.767125, 35.681236}},
			want: "POINT(139.767125, 35.681236)",
		},
		"仙台駅のポイントを指定してPOINTを表す文字列を取得できること": {
			args: args{point: [2]float64{140.882438, 38.260132}},
			want: "POINT(140.882438, 38.260132)",
		},
	}
	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := GeoPoint(arg.args.point)
			if actual.String() != arg.want {
				t.Fatalf("Error actual: %s, expected: %s", actual.String(), arg.want)
			}
		})
	}
}
