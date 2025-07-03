package util

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	d1 := "2025-07-01"
	d2 := "2025-07-03"
	d1Time, _ := StrToDate(d1)
	d2Time, _ := StrToDate(d2)
	diff := d2Time.Sub(d1Time)
	fmt.Println(diff)

	days, err := BetweenDays(d1, d2)
	if err != nil {
		t.Errorf("BetweenDays error: %v", err)
		return
	}
	fmt.Println("days: ", days)
	if days != int64(diff.Hours()/24) {
		t.Errorf("BetweenDays expected %d, got %d", int64(diff.Hours()/24), days)
	} else {
		fmt.Println("BetweenDays passed")
	}

}
func TestBetweenDays(t *testing.T) {
	tests := []struct {
		name    string
		start   string
		end     string
		want    int64
		wantErr bool
	}{
		{
			name:    "normal case",
			start:   "2025-07-01",
			end:     "2025-07-03",
			want:    2,
			wantErr: false,
		},
		{
			name:    "same day",
			start:   "2025-07-01",
			end:     "2025-07-01",
			want:    0,
			wantErr: false,
		},
		{
			name:    "invalid date format",
			start:   "2025/07/01",
			end:     "2025-07-03",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BetweenDays(tt.start, tt.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("BetweenDays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("BetweenDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBetweenDays2(t *testing.T) {
	tests := []struct {
		name    string
		start   string
		end     string
		want    int64
		wantErr bool
	}{
		{
			name:    "normal case",
			start:   "2025-07-01",
			end:     "2025-07-03",
			want:    3,
			wantErr: false,
		},
		{
			name:    "future date case",
			start:   "2025-07-01",
			end:     "2025-12-31",
			want:    3,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BetweenDays2(tt.start, tt.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("BetweenDays2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("BetweenDays2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToDate(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    time.Time
		wantErr bool
	}{
		{
			name:    "date only",
			str:     "2025-07-03",
			want:    time.Date(2025, 7, 3, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "date time",
			str:     "2025-07-03 15:04:05",
			want:    time.Date(2025, 7, 3, 15, 4, 5, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "invalid format",
			str:     "2025/07/03",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StrToDate(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !got.Equal(tt.want) {
				t.Errorf("StrToDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
