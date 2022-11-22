package zippopotam_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/johnnypivot/ziplookup/zippopotam"
)

func TestClientLookup(t *testing.T) {
	tests := map[string]struct {
		responseBody string
		responseCode int
		want         *zippopotam.Place
		err          error
	}{
		"90210": {
			responseCode: http.StatusOK,
			responseBody: `{"post code": "90210","country": "United States","country abbreviation": "US","places": [{"place name": "Beverly Hills","longitude": "-118.4065","state": "California","state abbreviation": "CA","latitude": "34.0901"}]}`,
			want: &zippopotam.Place{
				PlaceName:         "Beverly Hills",
				Longitude:         "-118.4065",
				State:             "California",
				StateAbbreviation: "CA",
				Latitude:          "34.0901",
			},
		},
		"00000": {
			responseCode: http.StatusNotFound,
			responseBody: `{}`,
			want:         nil,
			err:          zippopotam.ErrNoResults{Zip: "00000"},
		},
	}

	for zip, test := range tests {
		t.Run(zip, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.responseCode)
				w.Write([]byte(test.responseBody))
			}))
			defer ts.Close()

			hc := http.Client{Timeout: 10 * time.Second}
			zc := zippopotam.NewClient(ts.URL+"/", &hc)
			place, err := zc.Lookup(zip)
			if test.err != nil {
				if _, ok := err.(zippopotam.ErrNoResults); !ok {
					t.Errorf("expected error: %v, got: %v", test.err, err)
				}
			}
			if test.err == nil && !reflect.DeepEqual(place, test.want) {
				t.Errorf("got %v, want %v", place, test.want)
			}
		})
	}
}
