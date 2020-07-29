package urlbyter_test

import (
	"github.com/thmhoag/urlbyter/pkg/urlbyter"
	"reflect"
	"testing"
)

func TestBatchRunner_ProcessFile(t *testing.T) {
	type fields struct {
		opts *urlbyter.BatchOpts
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*urlbyter.BatchResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			br := urlbyter.NewBatchRunner(tt.fields.opts)
			got, err := br.ProcessFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessFile() got = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("batch processor should parse file and return results", func(t *testing.T) {


	})
}