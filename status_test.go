package apcupsd

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestStatus_parseKV(t *testing.T) {
	var tests = []struct {
		desc string
		kv   string
		s    *Status
		err  error
	}{
		{
			desc: "invalid format",
			kv:   "foo",
			err:  errInvalidKeyValuePair,
		},
		{
			desc: "invalid duration",
			kv:   "TIMELEFT : 1 ",
			err:  errInvalidDuration,
		},
		{
			desc: "unknown",
			kv:   "FOO : bar",
			s:    &Status{},
		},
		{
			desc: "OK string",
			kv:   "APC : 001,002,0003",
			s: &Status{
				APC: "001,002,0003",
			},
		},
		{
			desc: "OK float64",
			kv:   "LINEV : 120.0 Volts",
			s: &Status{
				LineVoltage: 120.0,
			},
		},
		{
			desc: "OK time.Time",
			kv:   "DATE     : 2020-04-27 10:00:00 -0000",
			s: &Status{
				Date: time.Date(2020, time.April, 27, 10, 0, 0, 0, time.UTC),
			},
		},
		{
			desc: "N/A time.Time",
			kv:   "XOFFBATT : N/A",
			s: &Status{
				XOnBattery: time.Time{},
			},
		},
		{
			desc: "OK time.Duration",
			kv:   "TIMELEFT: 10.5 Minutes",
			s: &Status{
				TimeLeft: 10*time.Minute + 30*time.Second,
			},
		},
		{
			desc: "OK NumberTransfers",
			kv:   "NUMXFERS: 1",
			s: &Status{
				NumberTransfers: 1,
			},
		},
		{
			desc: "OK NominalPower",
			kv:   "NOMPOWER: 865 Watts",
			s: &Status{
				NominalPower: 865,
			},
		},
		{
			desc: "OK Selftest",
			kv:   "SELFTEST: YES",
			s: &Status{
				Selftest: true,
			},
		},
		{
			desc: "No alarm ALARMDEL",
			kv:   "ALARMDEL: No alarm",
			s: &Status{
				AlarmDel: 0,
			},
		},
		{
			desc: "OK ITEMP",
			kv:   "ITEMP    : 35.1 C",
			s: &Status{
				InternalTemp: 35.1,
			},
		},
		{
			desc: "OK OUTPUTV",
			kv:   "OUTPUTV  : 230.4 Volts",
			s: &Status{
				OutputVoltage: 230.4,
			},
		},
		{
			desc: "OK LINEFREQ",
			kv:   "LINEFREQ : 50.0 Hz",
			s: &Status{
				LineFrequency: 50.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			s := new(Status)
			err := s.parseKV(tt.kv)

			// Simplify test table by nil'ing Status on errors
			if err != nil {
				s = nil
			}

			if want, got := tt.err, err; want != got {
				t.Fatalf("unexpected error:\n- want: %v\n-  got: %v", want, got)
			}

			if diff := cmp.Diff(tt.s, s); diff != "" {
				t.Fatalf("unexpected status (-want +got):\n%s", diff)
			}
		})
	}
}
