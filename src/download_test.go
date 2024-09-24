package src

import (
	"image"
	"testing"
)

func TestDownloadFiles(t *testing.T) {

	rect := image.Rect(0,0,100,100)
	img := image.NewNRGBA(rect)

	type args struct {
		img      image.Image
		art      []string
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Valid Inputs",args{img,[]string{},"lion.jpeg"},"lion",false},
		{"Invalid Inputs",args{img,[]string{},""},"",true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DownloadFiles(tt.args.img, tt.args.art, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DownloadFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
