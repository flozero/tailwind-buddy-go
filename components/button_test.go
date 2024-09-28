package components

import (
	"strings"
	"testing"
)

func TestButton(t *testing.T) {
	button := Button()

	testCases := []struct {
		name     string
		props    ButtonProps
		expected []string
	}{
		{
			name:     "Default props",
			props:    ButtonProps{},
			expected: []string{"px-4", "py-2", "rounded", "bg-blue-500", "text-white", "text-sm"},
		},
		{
			name: "Custom color only",
			props: ButtonProps{
				Color: "secondary",
			},
			expected: []string{"px-4", "py-2", "rounded", "bg-gray-200", "text-gray-800", "text-sm"},
		},
		{
			name: "Custom size only",
			props: ButtonProps{
				Size: "large",
			},
			
			expected: []string{"px-4", "py-2", "rounded", "bg-blue-500", "text-white", "text-lg", "font-bold"},
		},
		{
			name: "Custom color and size",
			props: ButtonProps{
				Color: "secondary",
				Size:  "large",
			},
			expected: []string{"px-4", "py-2", "rounded", "bg-gray-200", "text-gray-800", "text-lg" },
		},
		{
			name: "Compound variant",
			props: ButtonProps{
				Color: "primary",
				Size:  "large",
			},
			expected: []string{"px-4", "py-2", "rounded", "bg-blue-500", "text-white", "text-lg", "font-bold"},
		},
		{
			name: "Custom class only",
			props: ButtonProps{
				Class: "custom-class",
			},
			expected: []string{"px-4", "py-2", "rounded", "bg-blue-500", "text-white", "text-sm", "custom-class"},
		},
		{
			name: "All props",
			props: ButtonProps{
				Color: "secondary",
				Size:  "large",
				Class: "custom-class",
			},
			expected: []string{"px-4", "py-2", "rounded", "bg-gray-200", "text-gray-800", "text-lg", "custom-class"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := button.Root(tc.props)
			classes := strings.Split(result, " ")

			for _, expectedClass := range tc.expected {
				if !contains(classes, expectedClass) {
					t.Errorf("Expected class '%s' not found in result: %s", expectedClass, result)
				}
			}

			if len(classes) != len(tc.expected) {
				t.Errorf("Expected %d classes, but got %d. Result: %s", len(tc.expected), len(classes), result)
			}
		})
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}