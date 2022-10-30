package server

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestQueryToRange(t *testing.T) {
	v := url.Values{}
	v.Add("begin", "A22102312094410.jpg")
	v.Add("end", "A22102313094310.jpg")
	tests := []struct {
		name url.Values
		want []string
	}{
		{
			name: v,
			want: []string{"A22102312094410.jpg", "A22102313094310.jpg"},
		},
	}

	for _, tt := range tests {
		t.Run("Convert", func(t *testing.T) {
			got, err := queryToRange(tt.name)
			if err != nil {
				t.Errorf("queryToRange() error = %v, wantErr %v", err, false)
				return
			}
			if got[0] != tt.want[0] {
				t.Errorf("queryToRange() = %v, want %v", got, tt.want)
			}
			if got[1] != tt.want[1] {
				t.Errorf("queryToRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileNameToInt(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "A22102312094410.jpg",
			want: 22102312094410,
		},
		{
			name: "A22102313094310.jpg",
			want: 22102313094310,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fileNameToInt(tt.name)
			if err != nil {
				t.Errorf("fileNameToInt() error = %v, wantErr %v", err, false)
				return
			}
			if got != tt.want {
				t.Errorf("fileNameToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHealthCheckHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestEnableCors(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test EnableCors",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EnableCors(http.HandlerFunc(HealthCheckHandler))
		})
	}
}
