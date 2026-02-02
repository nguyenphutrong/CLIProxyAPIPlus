package executor

import (
	"testing"
)

func TestGitHubCopilotExecutor_normalizeModel(t *testing.T) {
	executor := &GitHubCopilotExecutor{}

	tests := []struct {
		name     string
		modelID  string
		body     string
		expected string
	}{
		{
			name:     "strips copilot- prefix",
			modelID:  "copilot-gpt-5",
			body:     `{"model":"copilot-gpt-5","messages":[]}`,
			expected: `{"model":"gpt-5","messages":[]}`,
		},
		{
			name:     "strips copilot/ prefix",
			modelID:  "copilot/gpt-5",
			body:     `{"model":"copilot/gpt-5","messages":[]}`,
			expected: `{"model":"gpt-5","messages":[]}`,
		},
		{
			name:     "passes through unprefixed model",
			modelID:  "gpt-5",
			body:     `{"model":"gpt-5","messages":[]}`,
			expected: `{"model":"gpt-5","messages":[]}`,
		},
		{
			name:     "handles empty model",
			modelID:  "",
			body:     `{"messages":[]}`,
			expected: `{"messages":[]}`,
		},
		{
			name:     "strips only first prefix on double-prefixed",
			modelID:  "copilot-copilot-gpt-5",
			body:     `{"model":"copilot-copilot-gpt-5","messages":[]}`,
			expected: `{"model":"copilot-gpt-5","messages":[]}`,
		},
		{
			name:     "case-insensitive prefix matching - uppercase",
			modelID:  "COPILOT-gpt-5",
			body:     `{"model":"COPILOT-gpt-5","messages":[]}`,
			expected: `{"model":"gpt-5","messages":[]}`,
		},
		{
			name:     "case-insensitive prefix matching - mixed case",
			modelID:  "CoPiLoT-gpt-5",
			body:     `{"model":"CoPiLoT-gpt-5","messages":[]}`,
			expected: `{"model":"gpt-5","messages":[]}`,
		},
		{
			name:     "strips copilot- from real Gemini model",
			modelID:  "copilot-gemini-3-flash-preview",
			body:     `{"model":"copilot-gemini-3-flash-preview","messages":[]}`,
			expected: `{"model":"gemini-3-flash-preview","messages":[]}`,
		},
		{
			name:     "strips copilot- from real Claude model",
			modelID:  "copilot-claude-sonnet-4.5",
			body:     `{"model":"copilot-claude-sonnet-4.5","messages":[]}`,
			expected: `{"model":"claude-sonnet-4.5","messages":[]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := executor.normalizeModel(tt.modelID, []byte(tt.body))
			if string(result) != tt.expected {
				t.Errorf("normalizeModel() = %s, want %s", string(result), tt.expected)
			}
		})
	}
}
