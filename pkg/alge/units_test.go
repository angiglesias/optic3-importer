package alge

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testDateStr      string = "2014-02-03"
	testStartTimeStr string = "10:35"
)

func TestDistanceMarshal(t *testing.T) {
	tests := []struct {
		value Distance
		enc   string
		err   bool
	}{
		// test cases
		{
			value: Steeplechase,
			enc:   "Steeplechase",
		},
		{
			value: Hurdles,
			enc:   "Hurdles",
		},
		{
			value: Regular,
			enc:   "Regular",
		},
		{
			value: Relay,
			enc:   "Relay",
		},
		{
			value: 33,
			enc:   "error",
			err:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.enc, func(t *testing.T) {
			enc, err := test.value.MarshalText()
			if test.err {
				assert.Error(t, err, "Unsupported encoding value should fail")
				return
			}
			require.NoError(t, err, "Unexpected error encoding value")
			assert.EqualValues(t, test.enc, string(enc), "Mismatched encoding value")
		})
	}
}

func TestDistanceUnmarshal(t *testing.T) {
	tests := []struct {
		dec Distance
		enc string
		err bool
	}{
		// test cases
		{
			dec: Steeplechase,
			enc: "Steeplechase",
		},
		{
			dec: Hurdles,
			enc: "Hurdles",
		},
		{
			dec: Regular,
			enc: "Regular",
		},
		{
			dec: Relay,
			enc: "Relay",
		},
		{
			dec: 33,
			enc: "33",
			err: true,
		},
	}
	for _, test := range tests {
		t.Run(test.enc, func(t *testing.T) {
			var dec Distance
			err := dec.UnmarshalText([]byte(test.enc))
			if test.err {
				assert.Error(t, err, "Unsupported encoded value should fail")
				return
			}
			require.NoError(t, err, "Unexpected error decoding value")
			assert.EqualValues(t, test.dec, dec, "Mismatched decoded value")
		})
	}
}

func TestWindMeasMarshal(t *testing.T) {
	tests := []struct {
		value WindMeas
		enc   string
		err   bool
	}{
		// test cases
		{
			value: TenSecWithDelay,
			enc:   "10SecondsWith10SecondsDelay",
		},
		{
			value: ThirteenSec,
			enc:   "13Seconds",
		},
		{
			value: TenSec,
			enc:   "10Seconds",
		},
		{
			value: FiveSec,
			enc:   "5Seconds",
		},
		{
			value: None,
			enc:   "None",
		},
		{
			value: 33,
			enc:   "error",
			err:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.enc, func(t *testing.T) {
			enc, err := test.value.MarshalText()
			if test.err {
				assert.Error(t, err, "Unsupported encoding value should fail")
				return
			}
			require.NoError(t, err, "Unexpected error encoding value")
			assert.EqualValues(t, test.enc, string(enc), "Mismatched encoding value")
		})
	}
}

func TestWindMeasUnmarshal(t *testing.T) {
	tests := []struct {
		dec WindMeas
		enc string
		err bool
	}{
		// test cases
		{
			dec: TenSecWithDelay,
			enc: "10SecondsWith10SecondsDelay",
		},
		{
			dec: ThirteenSec,
			enc: "13Seconds",
		},
		{
			dec: TenSec,
			enc: "10Seconds",
		},
		{
			dec: FiveSec,
			enc: "5Seconds",
		},
		{
			dec: None,
			enc: "None",
		},
		{
			dec: 33,
			enc: "error",
			err: true,
		},
	}
	for _, test := range tests {
		t.Run(test.enc, func(t *testing.T) {
			var dec WindMeas
			err := dec.UnmarshalText([]byte(test.enc))
			if test.err {
				assert.Error(t, err, "Unsupported encoded value should fail")
				return
			}
			require.NoError(t, err, "Unexpected error decoding value")
			assert.EqualValues(t, test.dec, dec, "Mismatched decoded value")
		})
	}
}

func TestDateUnmarshal(t *testing.T) {
	testDateTime := time.Date(2014, 02, 03, 0, 0, 0, 0, time.UTC)
	var d Date

	err := d.UnmarshalText([]byte(testDateStr))
	require.NoError(t, err, "Error decoding date")
	assert.EqualValues(t, testDateTime, d, "Mismatched dates")
}

func TestDateMarshal(t *testing.T) {
	d := Date(time.Date(2014, 02, 03, 0, 0, 0, 0, time.UTC))

	enc, err := d.MarshalText()
	require.NoError(t, err, "Error enconding date")
	assert.EqualValues(t, testDateStr, string(enc), "Mismatched dates")
}

func TestStartTimeUnmarshal(t *testing.T) {
	testStartTime := time.Date(0, time.January, 1, 10, 35, 0, 0, time.UTC)
	var s StartTime

	err := s.UnmarshalText([]byte(testStartTimeStr))
	require.NoError(t, err, "Error decoding start time")
	assert.EqualValues(t, testStartTime, s, "Mismatched start time")
}

func TestStartTimeMarshal(t *testing.T) {
	s := StartTime(time.Date(0, time.January, 1, 10, 35, 0, 0, time.UTC))

	enc, err := s.MarshalText()
	require.NoError(t, err, "Error encoding start time")
	assert.EqualValues(t, testStartTimeStr, string(enc), "Mismatched start times")
}
