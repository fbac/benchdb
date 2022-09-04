package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestGetCSVData(t *testing.T) {
	fileData, _ := os.Open("test/data.csv")
	fileNoData, _ := os.Open("test/nodata.csv")
	defer fileData.Close()
	defer fileNoData.Close()

	output := [][]string{
		{"2017-01-01 00:00:00", "host_000000", "92.77"},
	}

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			"Check test/nodata.csv",
			args{filename: "test/nodata.csv"},
			nil,
			true,
		},
		{
			"Check test/data.csv",
			args{filename: "test/data.csv"},
			output,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCSVData(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCSVData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCSVData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeReader(t *testing.T) {
	fileData, _ := os.Open("test/data.csv")
	fileNoData, _ := os.Open("test/nodata.csv")
	defer fileData.Close()
	defer fileNoData.Close()

	type args struct {
		f *os.File
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Check file with data",
			args{f: fileData},
			"*csv.Reader",
			false,
		},
		{
			"File with no data - returns error EOF",
			args{f: fileNoData},
			"*csv.Reader",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := initializeReader(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitializeReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			result := fmt.Sprintf("%T", got)
			if result != "*csv.Reader" {
				t.Errorf("TestInitializeReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileHasData(t *testing.T) {
	fileData, _ := os.Open("test/data.csv")
	fileNoData, _ := os.Open("test/nodata.csv")
	defer fileData.Close()
	defer fileNoData.Close()

	type args struct {
		file *os.File
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Check file with data",
			args{file: fileData},
			true,
		},
		{
			"Check file with no data",
			args{file: fileNoData},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileHasData(tt.args.file); got != tt.want {
				t.Errorf("fileHasData() = %v, want %v", got, tt.want)
			}
		})
	}
}
