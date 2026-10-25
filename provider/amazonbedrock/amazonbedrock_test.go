package amazonbedrock

import (
	"context"
	"errors"
	"testing"
)

type mockClient struct {
	invokeFunc func(ctx context.Context, modelID string, input []byte) ([]byte, error)
}

func (m *mockClient) InvokeModel(ctx context.Context, modelID string, input []byte) ([]byte, error) {
	return m.invokeFunc(ctx, modelID, input)
}

func TestProvider_CallModel(t *testing.T) {
	ctx := context.Background()
	input := []byte(`{"prompt":"hello"}`)

	tests := []struct {
		name      string
		modelID   string
		mockFn    func(ctx context.Context, modelID string, input []byte) ([]byte, error)
		want      []byte
		wantErr   error
	}{
		{
			name:    "success",
			modelID: "test-model",
			mockFn: func(ctx context.Context, modelID string, input []byte) ([]byte, error) {
				return []byte(`{"response":"ok"}`), nil
			},
			want:    []byte(`{"response":"ok"}`),
			wantErr: nil,
		},
		{
			name:    "model not found",
			modelID: "missing-model",
			mockFn: func(ctx context.Context, modelID string, input []byte) ([]byte, error) {
				return nil, ErrModelNotFound
			},
			want:    nil,
			wantErr: ErrModelNotFound,
		},
		{
			name:    "empty modelID",
			modelID: "",
			mockFn: func(ctx context.Context, modelID string, input []byte) ([]byte, error) {
				return nil, nil
			},
			want:    nil,
			wantErr: errors.New("modelID cannot be empty"),
		},
		{
			name:    "other error",
			modelID: "error-model",
			mockFn: func(ctx context.Context, modelID string, input []byte) ([]byte, error) {
				return nil, errors.New("some error")
			},
			want:    nil,
			wantErr: errors.New("failed to invoke model: some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewProvider(&mockClient{invokeFunc: tt.mockFn})
			got, err := p.CallModel(ctx, tt.modelID, input)
			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.wantErr)
				}
				if err.Error() != tt.wantErr.Error() && !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if string(got) != string(tt.want) {
				t.Errorf("expected output %s, got %s", string(tt.want), string(got))
			}
		})
	}
}
