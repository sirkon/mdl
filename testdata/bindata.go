// Code generated by go-bindata.
// sources:
// allfeatures.md
// bindata.go
// codearray.md
// codecomment.md
// codelist.md
// commentcode.md
// curjoblike.md
// easy_example.md
// generate.go
// keycheck.md
// level_normalization.md
// maps.md
// metric.md
// rawsection
// rawunmarshaler.md
// regression.md
// regression2.md
// regression3.md
// scalar_decoder.md
// statuses.md
// struct_easy.md
// struct_real.md
// test.md
// DO NOT EDIT!

package testdata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _allfeaturesMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xce\xcd\x4a\xc5\x30\x10\xc5\xf1\xfd\x3c\xc5\xe1\x66\x21\xb8\x53\x10\x57\x2e\x44\xae\xab\xeb\x67\x7d\x80\x4c\xcd\x98\x44\x9a\x89\x26\xd3\x4a\xdf\x5e\x8a\x05\xb7\x3f\x0e\x87\xbf\xf7\xbe\xf1\x0f\x7d\x71\x13\xb5\x47\x2e\x82\x1b\x1c\x9e\xc5\xa4\x1d\x76\xbc\x8d\x9b\x5d\x5d\x93\xf7\x9e\xc8\xe1\x4f\x89\x9c\x43\xaf\x8a\x0b\x2a\x1c\x04\x8c\x31\x1b\xea\x07\x86\x97\x13\xac\x06\x5e\xb7\x7d\xff\x9e\x68\x38\x9e\x8e\x77\x6f\x38\xc7\xfd\xeb\xd3\x03\x8c\xc7\x49\xf6\x2b\x87\xc0\x73\x4c\x26\x8d\x9c\x73\x88\x8d\x35\xf4\xaa\xf4\x39\x77\xc3\x58\x9b\x82\x35\xc0\x92\xac\x48\xbc\x88\x9e\x19\x22\x2f\x82\x94\x0b\x18\xba\xc5\xae\xf2\x5f\x72\x49\xb9\x83\x51\xe6\x9e\xdf\x33\x2b\xfd\x06\x00\x00\xff\xff\x61\xa9\x98\x4e\xdb\x00\x00\x00")

func allfeaturesMdBytes() ([]byte, error) {
	return bindataRead(
		_allfeaturesMd,
		"allfeatures.md",
	)
}

func allfeaturesMd() (*asset, error) {
	bytes, err := allfeaturesMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "allfeatures.md", size: 219, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(436), modTime: time.Unix(1533911326, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _codearrayMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x48\x48\x28\x2e\xcc\xe1\x0a\x76\xf5\x71\x75\x0e\x51\xd0\x52\x70\x0b\xf2\xf7\x55\x28\x49\x4c\xca\x49\xe5\x4a\x48\x48\xe0\xe2\xc2\x2d\x6f\x44\x48\x81\x31\x5c\x41\x56\x71\x7e\x1e\x57\x35\x97\x82\x82\x52\xa2\x92\x95\x82\xa1\x11\x57\x2d\x48\x14\x10\x00\x00\xff\xff\x3f\xf6\x74\x52\x7b\x00\x00\x00")

func codearrayMdBytes() ([]byte, error) {
	return bindataRead(
		_codearrayMd,
		"codearray.md",
	)
}

func codearrayMd() (*asset, error) {
	bytes, err := codearrayMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "codearray.md", size: 123, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _codecommentMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x48\x48\x28\x2e\xcc\xe1\x0a\x76\xf5\x71\x75\x0e\x51\xd0\x52\x70\x0b\xf2\xf7\x55\x28\x49\x4c\xca\x49\xe5\x4a\x48\x48\xe0\x0a\xc9\xc8\x2c\x56\x28\x4f\x2c\x56\xc8\x2a\x2d\x2e\x51\x48\x54\x28\x4a\x2d\x2c\x4d\x2d\x2e\x01\x04\x00\x00\xff\xff\x1e\x5b\x8f\x65\x36\x00\x00\x00")

func codecommentMdBytes() ([]byte, error) {
	return bindataRead(
		_codecommentMd,
		"codecomment.md",
	)
}

func codecommentMd() (*asset, error) {
	bytes, err := codecommentMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "codecomment.md", size: 54, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _codelistMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x28\x4e\x2d\x51\x28\x29\xca\x2c\xcb\x4c\xcc\xe1\xe2\x4a\x48\x48\x28\xa8\x2c\xc9\xc8\xcf\xe3\x4a\x54\xd0\xb6\x55\x30\x04\x09\x70\x71\x41\x14\x15\xa5\xe2\x56\x91\x90\x90\x50\x5c\x98\xc3\x15\xec\xea\xe3\xea\x1c\xa2\xa0\xa5\xe0\x16\xe4\xef\xab\x50\x92\x98\x94\x93\xaa\x80\x6c\x44\x6a\x51\x51\x7e\x11\x20\x00\x00\xff\xff\x2d\x18\x55\x32\x73\x00\x00\x00")

func codelistMdBytes() ([]byte, error) {
	return bindataRead(
		_codelistMd,
		"codelist.md",
	)
}

func codelistMd() (*asset, error) {
	bytes, err := codelistMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "codelist.md", size: 115, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commentcodeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\xc9\xc8\x2c\x56\x28\xcf\xcc\xc9\x51\x48\x4a\x55\x48\x54\x28\x4a\x2d\x2c\x4d\x2d\x2e\xe1\x4a\x48\x48\x28\x2e\xcc\xe1\x0a\x76\xf5\x71\x75\x0e\x51\xd0\x52\x70\x0b\xf2\xf7\x55\x28\x49\x4c\xca\x49\x05\x49\x01\x02\x00\x00\xff\xff\xfc\xe4\xb5\xd4\x35\x00\x00\x00")

func commentcodeMdBytes() ([]byte, error) {
	return bindataRead(
		_commentcodeMd,
		"commentcode.md",
	)
}

func commentcodeMd() (*asset, error) {
	bytes, err := commentcodeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "commentcode.md", size: 53, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _curjoblikeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xc1\x4e\xc2\x40\x10\x86\xef\xfb\x14\x7f\xda\x8b\x10\x62\xe2\x95\x1b\x28\x07\x03\xda\x04\x8d\x57\x77\x69\xa7\x76\x0d\xec\xc2\xcc\x14\x25\xc6\xc4\xa7\xf1\xc1\x7c\x12\xd3\x6d\xf0\x60\xe0\xfa\x4f\xbe\x6f\xbf\xcd\xf1\x40\xbc\xf7\x25\x5d\xde\x91\x36\xb1\x32\x26\xcf\x73\x30\xed\x5a\x12\x15\x33\x44\xed\x59\x14\x31\x10\x62\x8d\xbd\x63\xef\x82\xe2\xcd\x6b\x03\xeb\x2c\x6a\x4f\xeb\xca\x58\x6b\x5f\x25\x06\xf3\x61\x00\x20\x73\xd9\x18\x57\xe6\xb3\x9b\x8d\x19\x42\xa8\x8c\xa1\x3a\xad\x58\x9d\x51\xac\xb2\x31\xb2\x74\xc9\x8e\xa2\xbe\x4b\xb6\x31\x08\x5d\x14\xf3\xc1\x7f\x44\xd4\x69\x2b\x1d\x57\xcc\xb3\x51\xbf\x25\xc3\x75\x0c\xd2\x6e\xa8\xea\x4e\xee\xb4\x6e\xb6\x5c\x16\xcb\x11\xee\x8b\xc7\xe7\xc9\xd3\xe4\x76\x31\x99\x2e\x66\x03\x53\x84\xf5\x21\x65\x8b\xdb\x6c\xd7\x84\x3a\x32\xfa\x57\x60\x13\x61\xe1\x05\x2f\x7e\x4f\x01\x4e\x10\xb5\x21\x16\x34\x4e\xa0\x4d\x82\x08\x52\x36\xb4\xa1\xf3\xa9\x49\x73\xac\x65\x72\x12\x43\x37\xdf\x4c\x41\xcc\x91\xc7\xf0\xa1\x8c\xcc\x54\x2a\x76\x2d\xf1\x01\x72\x08\xea\xde\xe1\x14\x3f\x5f\xdf\x7f\x9f\xf9\x0d\x00\x00\xff\xff\x39\x69\xb6\x4f\xc6\x01\x00\x00")

func curjoblikeMdBytes() ([]byte, error) {
	return bindataRead(
		_curjoblikeMd,
		"curjoblike.md",
	)
}

func curjoblikeMd() (*asset, error) {
	bytes, err := curjoblikeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "curjoblike.md", size: 454, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _easy_exampleMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcb\x41\x0a\x02\x31\x0c\x46\xe1\xfd\x7f\x8a\xd8\x9e\xc0\xee\xbd\x82\x67\x48\x98\x46\x67\xa0\x4d\x86\x10\x18\xbc\xbd\x14\x74\xf7\x78\xf0\x31\x73\xc8\x05\x79\x2b\x3d\xa8\x35\x98\xcc\x55\xe5\x29\x53\x0b\x98\x19\xa8\xf4\x3a\x74\xf4\x3b\x6a\xdd\xbc\xeb\x9a\xe7\x27\x77\x37\x9c\x71\x58\x52\xd9\x75\x0c\xa7\xcb\x63\xf4\xdb\xdf\xfc\x50\x43\x88\x75\x9f\xb4\xb9\xa5\x5a\xe2\x1b\x00\x00\xff\xff\xf4\x4d\x02\x2b\x70\x00\x00\x00")

func easy_exampleMdBytes() ([]byte, error) {
	return bindataRead(
		_easy_exampleMd,
		"easy_example.md",
	)
}

func easy_exampleMd() (*asset, error) {
	bytes, err := easy_exampleMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "easy_example.md", size: 112, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _generateGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x49\x2d\x2e\x49\x49\x2c\x49\xe4\xe2\xd2\xd7\x4f\xcf\xb7\x4a\x4f\xcd\x4b\x2d\x4a\x2c\x49\x55\x48\xcf\xd7\x4d\xca\xcc\x03\xc9\x28\xe8\x16\x64\xa7\xc3\xd5\x29\xe8\x71\x01\x02\x00\x00\xff\xff\x2b\xe0\x95\x6b\x3b\x00\x00\x00")

func generateGoBytes() ([]byte, error) {
	return bindataRead(
		_generateGo,
		"generate.go",
	)
}

func generateGo() (*asset, error) {
	bytes, err := generateGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "generate.go", size: 59, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _keycheckMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x48\x48\x28\x4a\x2c\xe7\xca\x4e\xad\x34\x54\xb0\x55\x50\x2a\x4b\xcc\x29\x4d\x35\x54\x02\xf1\x8d\xe0\x7c\x23\x30\xdf\x18\xce\x37\x56\xe2\x4a\x48\x48\x00\x04\x00\x00\xff\xff\x2f\x5d\x1e\x45\x3a\x00\x00\x00")

func keycheckMdBytes() ([]byte, error) {
	return bindataRead(
		_keycheckMd,
		"keycheck.md",
	)
}

func keycheckMd() (*asset, error) {
	bytes, err := keycheckMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "keycheck.md", size: 58, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _level_normalizationMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x56\x56\xc8\x30\xe4\x52\x56\x06\x33\x8c\xb8\x20\xb4\x31\x20\x00\x00\xff\xff\x28\xc3\x13\x4e\x17\x00\x00\x00")

func level_normalizationMdBytes() ([]byte, error) {
	return bindataRead(
		_level_normalizationMd,
		"level_normalization.md",
	)
}

func level_normalizationMd() (*asset, error) {
	bytes, err := level_normalizationMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "level_normalization.md", size: 23, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mapsMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xc8\x4e\xad\x34\xe4\x4a\x48\x48\xa8\x4c\xcc\xcd\xe1\x4a\xb4\x52\x00\x73\xb8\xb8\xc0\x12\x46\xc8\x12\x46\x10\x89\x8b\xdd\x17\x76\xe9\x82\x08\x85\x0b\x7b\x2f\x6c\xbd\xb0\xeb\x62\xf3\x85\xed\x17\xfb\x2f\x6c\xba\xb0\xef\xc2\x56\x40\x00\x00\x00\xff\xff\x6a\xfc\x7f\xb3\x4e\x00\x00\x00")

func mapsMdBytes() ([]byte, error) {
	return bindataRead(
		_mapsMd,
		"maps.md",
	)
}

func mapsMd() (*asset, error) {
	bytes, err := mapsMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "maps.md", size: 78, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\xdf\x4a\xf3\x40\x10\xc5\xef\xf7\x29\x0e\xfd\x3e\x68\x2b\xc5\x3e\x41\x2f\xd4\xc6\x3f\xa0\x54\x5a\xa5\x78\x65\xd6\x66\x4a\x16\x36\xd9\x76\x67\xd2\x98\x3b\x9f\xc6\x07\xf3\x49\x64\x93\xb4\x14\x11\xa9\x48\x6f\x72\x71\x66\xe6\x9c\xdf\x0c\x59\x49\x0d\x63\x69\x2c\xc1\x30\xd8\x65\x54\xa6\xe4\x09\x26\x47\xfc\xf1\xf6\x3e\xcc\x48\xbc\x59\xf0\x90\xc5\x79\x43\x3c\x5c\x15\x2f\xd6\x70\x4a\xfe\xb9\xc8\xcd\x9a\x4f\xb3\x24\x56\x2a\x8e\x63\xaf\x4b\x25\xd5\x8a\x30\x42\x67\xe1\x8a\x5c\x7a\xf5\xd7\x57\x03\x6c\xc8\xb3\x71\x79\xbf\x13\xfa\x94\xfa\x87\x75\x41\xc1\x4b\x29\x49\x89\x09\xda\xd3\x56\xc2\xd2\x79\x04\xe3\x82\x50\x30\x79\x46\x99\x3a\x6c\x33\x13\xb4\x14\x81\x4e\x23\xd1\x95\x52\x27\x58\x91\x87\xb6\x16\x4d\x5e\xa8\xea\x3c\xc1\x86\x42\x26\x87\x48\x5e\x5b\x35\x8b\x6e\xa3\x8b\x07\x05\x00\xff\xc7\x5a\x08\x67\x33\x24\x5a\x68\x50\x4b\x21\x31\x7a\xd5\x0b\xe9\x85\xd0\x9b\x71\x3f\x94\x6b\xbf\xa6\xde\x9d\xcf\xbb\x3b\xc9\x57\xad\xa8\xad\xad\xd5\x76\x3f\x75\x39\x9d\xdc\x6d\x09\xd5\xfc\x3a\x9a\x46\x75\x04\x46\x4d\x64\xb3\x7d\xc3\xdb\x8e\xa0\x34\x92\x86\x5d\xf6\xf1\x8f\x88\x7c\x28\xe9\xd5\x74\xf2\x78\x8f\xf3\xa7\xdd\xc0\x1e\x7a\x6b\x88\xdc\x21\xd3\x22\xe4\x21\x29\xed\x37\xfe\x9d\xfd\xc7\x33\xe3\x70\xfa\xd6\xe7\x3b\xfa\xe6\x17\x39\x16\xf3\xaf\xef\xfc\xf5\xad\xd4\xc8\x9f\x01\x00\x00\xff\xff\xb8\x94\x2d\xb2\x9a\x03\x00\x00")

func metricMdBytes() ([]byte, error) {
	return bindataRead(
		_metricMd,
		"metric.md",
	)
}

func metricMd() (*asset, error) {
	bytes, err := metricMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metric.md", size: 922, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rawsection = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x54\xb0\x55\x30\xe4\x4a\x52\xb0\x55\x50\x4a\x4a\x55\xe2\x52\x48\x56\xb0\x55\x28\x29\x2a\x4d\xe5\x4a\x51\x50\xb0\x55\x50\x30\xcc\x4e\xe2\x4a\xb5\x35\xd2\x33\x00\x04\x00\x00\xff\xff\xc2\xa6\x04\x08\x28\x00\x00\x00")

func rawsectionBytes() ([]byte, error) {
	return bindataRead(
		_rawsection,
		"rawsection",
	)
}

func rawsection() (*asset, error) {
	bytes, err := rawsectionBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rawsection", size: 40, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rawunmarshalerMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x48\x48\x28\x4a\x2c\xe7\x4a\xb4\x35\xe4\x4a\xb2\xf5\x74\xe1\x4a\xb6\x55\x32\x54\xe2\x4a\xb1\x2d\x29\x2a\x4d\xe5\x4a\xb5\x35\x34\xb2\xc8\x4e\xe2\x4a\x48\x48\x00\x04\x00\x00\xff\xff\xb0\x40\xbe\x8e\x28\x00\x00\x00")

func rawunmarshalerMdBytes() ([]byte, error) {
	return bindataRead(
		_rawunmarshalerMd,
		"rawunmarshaler.md",
	)
}

func rawunmarshalerMd() (*asset, error) {
	bytes, err := rawunmarshalerMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rawunmarshaler.md", size: 40, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regressionMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x48\x48\x28\x4a\x2c\xe7\xca\x4b\x54\xb0\x55\xd0\x35\xe4\xca\x4b\x52\xb0\x55\x30\x34\xd2\x33\xe0\x4a\x48\x48\x00\x04\x00\x00\xff\xff\x1f\x8b\x88\x74\x1c\x00\x00\x00")

func regressionMdBytes() ([]byte, error) {
	return bindataRead(
		_regressionMd,
		"regression.md",
	)
}

func regressionMd() (*asset, error) {
	bytes, err := regressionMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "regression.md", size: 28, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regression2Md = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x4e\xcb\x4c\xcd\x49\x31\xe4\x52\x56\x86\xb0\x8c\xb8\x12\x12\x12\x8a\x0b\x73\xb8\x82\x5d\x7d\x5c\x9d\x43\x40\x3c\x40\x00\x00\x00\xff\xff\x54\x31\x4c\xf9\x23\x00\x00\x00")

func regression2MdBytes() ([]byte, error) {
	return bindataRead(
		_regression2Md,
		"regression2.md",
	)
}

func regression2Md() (*asset, error) {
	bytes, err := regression2MdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "regression2.md", size: 35, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regression3Md = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x4e\xcb\x4c\xcd\x49\xe1\x52\x56\x4e\xce\x4f\x49\xe5\x4a\x48\x48\x28\xa8\x2c\xc9\xc8\xcf\xe3\xaa\xe6\xaa\x05\xf1\xb8\x14\x14\x14\x14\xb8\x00\x01\x00\x00\xff\xff\x4c\x6c\x9d\x65\x25\x00\x00\x00")

func regression3MdBytes() ([]byte, error) {
	return bindataRead(
		_regression3Md,
		"regression3.md",
	)
}

func regression3Md() (*asset, error) {
	bytes, err := regression3MdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "regression3.md", size: 37, mode: os.FileMode(436), modTime: time.Unix(1533911318, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scalar_decoderMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\x8e\xc1\xca\x82\x40\x14\x85\xd7\xff\x7d\x8a\xc3\xac\xf4\x87\x82\x71\xef\x2a\x6c\x65\x08\xd5\x03\xdc\xeb\x68\x36\xa9\xa3\x39\x23\x05\xd1\xbb\x87\xb3\x39\x9c\xc3\xc7\x81\x8f\x99\x17\x79\x91\x20\x87\xa6\x1a\x39\x76\x9a\xcc\x36\xf6\x3a\xa3\x06\x39\x94\xce\x14\x31\x33\x6d\xe1\x9f\x03\x5d\x8a\xb2\x38\x5c\xf1\x8f\xe3\xb9\x3a\x21\x48\x3d\xb4\x91\x33\x73\x37\xd1\x2c\xa6\x97\xae\xc5\x28\xd6\xd1\x6d\x75\x26\xb6\x24\xc5\x87\xfe\x66\x71\xd6\x24\xaa\xac\x4a\x95\xd2\x37\x9e\xc2\xdd\x7a\x58\x8f\xc7\xea\x03\x04\x8b\xb8\x66\x1a\x11\xda\x77\xd8\x70\x98\xc6\x21\xaa\x29\xdd\xd7\xd1\xe2\x17\x00\x00\xff\xff\xee\xcd\x47\x88\xaf\x00\x00\x00")

func scalar_decoderMdBytes() ([]byte, error) {
	return bindataRead(
		_scalar_decoderMd,
		"scalar_decoder.md",
	)
}

func scalar_decoderMd() (*asset, error) {
	bytes, err := scalar_decoderMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scalar_decoder.md", size: 175, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _statusesMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcc\xb1\x09\x02\x41\x10\x46\xe1\x7c\xaa\x18\xb0\x81\x45\xce\xd0\x62\xac\xe1\xce\x5c\x3d\x14\x03\x41\xec\x64\x05\x07\x17\x65\xd9\x16\xde\xdf\x91\x6c\x60\x72\xc9\x8b\x3e\xde\x6a\x9c\x76\xd3\x7e\xdc\xae\x53\x32\x3e\xae\x59\x47\x1a\xa1\xb3\x66\xa7\xe9\x40\xd1\x55\x0f\xbe\x66\x7f\x39\xa4\xc1\xa8\x94\x4e\x74\x77\x2a\xd1\x93\x97\x6e\xd3\x8f\xc1\x93\xac\x93\x6e\x4e\xa3\xf0\xe6\x45\xe8\xf2\x0b\x00\x00\xff\xff\xc4\x59\x52\xc9\x75\x00\x00\x00")

func statusesMdBytes() ([]byte, error) {
	return bindataRead(
		_statusesMd,
		"statuses.md",
	)
}

func statusesMd() (*asset, error) {
	bytes, err := statusesMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "statuses.md", size: 117, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _struct_easyMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x48\xe4\x4a\x48\x48\x48\xcf\xe7\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x4d\xcc\xcc\xe3\xe2\x4a\x2b\xcd\x4b\x06\x33\x35\x34\x15\xaa\xb9\x38\x0b\x12\xf3\x32\x93\x35\x94\x52\x8b\x8a\xf2\x8b\x94\x34\xb9\x6a\x41\x7a\xb8\xb8\x94\x15\x92\xb8\xb2\x4a\x8b\x4b\x14\x12\x15\x4a\x52\x2b\x4a\x00\x01\x00\x00\xff\xff\x2a\x52\x0d\xd3\x4c\x00\x00\x00")

func struct_easyMdBytes() ([]byte, error) {
	return bindataRead(
		_struct_easyMd,
		"struct_easy.md",
	)
}

func struct_easyMd() (*asset, error) {
	bytes, err := struct_easyMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "struct_easy.md", size: 76, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _struct_realMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcd\x31\x6e\x83\x30\x18\x05\xe0\xfd\x3f\xc5\x2f\xbc\x31\x54\x6a\xab\x8e\x0c\x14\xb9\x13\x55\x25\xe0\x00\x36\x95\x37\x06\x70\x5b\x55\xd9\x48\x86\x2c\x19\x72\x15\x07\x82\x84\x14\x48\xae\xf0\x7c\xa3\x08\x86\x6c\xef\x49\xdf\xd3\x53\x4a\x59\xfd\x4f\x3a\x7a\xa6\x32\x0a\x5e\x02\xfa\x8e\x5e\x9f\xde\x48\x29\x45\xa2\xf9\x33\x76\x43\x42\x88\xda\x9a\x5a\x5b\x43\x21\xa3\xf3\x2d\x06\x4c\x18\x30\x63\x86\xf3\x47\xf6\x3b\x38\x9c\x70\xc1\xe8\xf7\x70\xec\xb7\x8c\xdb\x8a\xce\x70\xe8\xd7\xd4\x63\xf4\x2d\xae\xe8\xe0\x96\x99\x3f\x60\xc2\xc8\x2b\x78\xd4\xe5\xf2\xa7\xa9\x28\xc9\x64\x5c\x48\x2e\xe2\xf7\x54\xb2\xe6\x38\xe7\x5c\xa6\x32\x29\x38\xe4\x8f\xec\xeb\x93\x7f\x75\x59\x99\x05\xdf\x03\x00\x00\xff\xff\x5a\xf8\xcc\x0a\xba\x00\x00\x00")

func struct_realMdBytes() ([]byte, error) {
	return bindataRead(
		_struct_realMd,
		"struct_real.md",
	)
}

func struct_realMd() (*asset, error) {
	bytes, err := struct_realMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "struct_real.md", size: 186, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x48\x48\x28\x4a\x2c\xe7\x4a\x54\xb0\x55\x50\x32\x54\xe2\x4a\x48\x48\xe0\xe2\x52\x2e\x4a\x2d\x2c\x4d\x2d\x2e\xe1\x82\x4a\x26\x29\xd8\x2a\x18\x82\x38\x80\x00\x00\x00\xff\xff\x7e\x1e\x9d\xd6\x2d\x00\x00\x00")

func testMdBytes() ([]byte, error) {
	return bindataRead(
		_testMd,
		"test.md",
	)
}

func testMd() (*asset, error) {
	bytes, err := testMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test.md", size: 45, mode: os.FileMode(436), modTime: time.Unix(1511864412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"allfeatures.md": allfeaturesMd,
	"bindata.go": bindataGo,
	"codearray.md": codearrayMd,
	"codecomment.md": codecommentMd,
	"codelist.md": codelistMd,
	"commentcode.md": commentcodeMd,
	"curjoblike.md": curjoblikeMd,
	"easy_example.md": easy_exampleMd,
	"generate.go": generateGo,
	"keycheck.md": keycheckMd,
	"level_normalization.md": level_normalizationMd,
	"maps.md": mapsMd,
	"metric.md": metricMd,
	"rawsection": rawsection,
	"rawunmarshaler.md": rawunmarshalerMd,
	"regression.md": regressionMd,
	"regression2.md": regression2Md,
	"regression3.md": regression3Md,
	"scalar_decoder.md": scalar_decoderMd,
	"statuses.md": statusesMd,
	"struct_easy.md": struct_easyMd,
	"struct_real.md": struct_realMd,
	"test.md": testMd,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"allfeatures.md": &bintree{allfeaturesMd, map[string]*bintree{}},
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
	"codearray.md": &bintree{codearrayMd, map[string]*bintree{}},
	"codecomment.md": &bintree{codecommentMd, map[string]*bintree{}},
	"codelist.md": &bintree{codelistMd, map[string]*bintree{}},
	"commentcode.md": &bintree{commentcodeMd, map[string]*bintree{}},
	"curjoblike.md": &bintree{curjoblikeMd, map[string]*bintree{}},
	"easy_example.md": &bintree{easy_exampleMd, map[string]*bintree{}},
	"generate.go": &bintree{generateGo, map[string]*bintree{}},
	"keycheck.md": &bintree{keycheckMd, map[string]*bintree{}},
	"level_normalization.md": &bintree{level_normalizationMd, map[string]*bintree{}},
	"maps.md": &bintree{mapsMd, map[string]*bintree{}},
	"metric.md": &bintree{metricMd, map[string]*bintree{}},
	"rawsection": &bintree{rawsection, map[string]*bintree{}},
	"rawunmarshaler.md": &bintree{rawunmarshalerMd, map[string]*bintree{}},
	"regression.md": &bintree{regressionMd, map[string]*bintree{}},
	"regression2.md": &bintree{regression2Md, map[string]*bintree{}},
	"regression3.md": &bintree{regression3Md, map[string]*bintree{}},
	"scalar_decoder.md": &bintree{scalar_decoderMd, map[string]*bintree{}},
	"statuses.md": &bintree{statusesMd, map[string]*bintree{}},
	"struct_easy.md": &bintree{struct_easyMd, map[string]*bintree{}},
	"struct_real.md": &bintree{struct_realMd, map[string]*bintree{}},
	"test.md": &bintree{testMd, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

