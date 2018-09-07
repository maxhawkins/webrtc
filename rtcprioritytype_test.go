package webrtc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRTCPriorityType(t *testing.T) {
	testCases := []struct {
		priorityString   string
		priorityUint16   uint16
		expectedPriority RTCPriorityType
	}{
		{"unknown", 0, RTCPriorityType(Unknown)},
		{"very-low", 100, RTCPriorityTypeVeryLow},
		{"low", 200, RTCPriorityTypeLow},
		{"medium", 300, RTCPriorityTypeMedium},
		{"high", 1000, RTCPriorityTypeHigh},
	}

	for i, testCase := range testCases {
		assert.Equal(t,
			newRTCPriorityTypeFromString(testCase.priorityString),
			testCase.expectedPriority,
			"testCase: %d %v", i, testCase,
		)

		assert.Equal(t,
			newRTCPriorityTypeFromUint16(testCase.priorityUint16),
			testCase.expectedPriority,
			"testCase: %d %v", i, testCase,
		)
	}
}

func TestRTCPriorityType_String(t *testing.T) {
	testCases := []struct {
		priority       RTCPriorityType
		expectedString string
	}{
		{RTCPriorityType(Unknown), "unknown"},
		{RTCPriorityTypeVeryLow, "very-low"},
		{RTCPriorityTypeLow, "low"},
		{RTCPriorityTypeMedium, "medium"},
		{RTCPriorityTypeHigh, "high"},
	}

	for i, testCase := range testCases {
		assert.Equal(t,
			testCase.priority.String(),
			testCase.expectedString,
			"testCase: %d %v", i, testCase,
		)
	}
}
