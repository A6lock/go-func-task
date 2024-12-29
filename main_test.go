package main

import "testing"

func TestMaskingLinks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No links",
			input:    "Hello, this is a test string.",
			expected: "Hello, this is a test string.",
		},
		{
			name:     "Single link",
			input:    "Visit my page: http://example.com",
			expected: "Visit my page: http://***********",
		},
		{
			name:     "Multiple links",
			input:    "Here are two links: http://first.com and http://second.com",
			expected: "Here are two links: http://********* and http://**********",
		},
		{
			name:     "Link at the end",
			input:    "Check this out: http://example.com",
			expected: "Check this out: http://***********",
		},
		{
			name:     "Link at the beginning",
			input:    "http://example.com is a great site",
			expected: "http://*********** is a great site",
		},
		{
			name:     "Link with spaces",
			input:    "This is a link: http://example.com/page with spaces",
			expected: "This is a link: http://**************** with spaces",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maskingLinks(tt.input)
			if result != tt.expected {
				t.Errorf("maskingLinks() = %v, want %v", result, tt.expected)
			}
		})
	}
}
