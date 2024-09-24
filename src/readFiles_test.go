package src

import (
	"image"
	"os"
	"reflect"
	"testing"
)

func TestGetArt(t *testing.T) {
	// rect := image.Rect(0, 0, 10, 10)
	// img := image.NewRGBA(rect)

	EmptyTempFile, _ := os.CreateTemp("", "test_file_*.txt")
	defer os.Remove(EmptyTempFile.Name())


	tempFile, _ := os.CreateTemp("", "test_file_*.txt")
	tempFile.Write([]byte(".....\n.....\n"))
	tempFile.Seek(0,0)
	defer os.Remove(tempFile.Name())
	type args struct {
		file *os.File
	}
	tests := []struct {
		name    string
		args    args
		want    image.Image
		want1   []string
		wantErr bool
	}{
		{"Valid input",args{tempFile},nil,nil,true},
		{"Empty file input", args{EmptyTempFile}, nil, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GetArt(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAsciiArt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAsciiArt() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetAsciiArt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
