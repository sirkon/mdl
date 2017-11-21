// Code generated by go-bindata.
// sources:
// bindata.go
// codearray.md
// codecomment.md
// commentcode.md
// generate.go
// level_normalization.md
// maps.md
// rawsection
// regression.md
// scalar_decoder.md
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

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(436), modTime: time.Unix(1511287079, 0)}
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

	info := bindataFileInfo{name: "codearray.md", size: 123, mode: os.FileMode(436), modTime: time.Unix(1511114177, 0)}
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

	info := bindataFileInfo{name: "codecomment.md", size: 54, mode: os.FileMode(436), modTime: time.Unix(1511109818, 0)}
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

	info := bindataFileInfo{name: "commentcode.md", size: 53, mode: os.FileMode(436), modTime: time.Unix(1511110059, 0)}
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

	info := bindataFileInfo{name: "generate.go", size: 59, mode: os.FileMode(436), modTime: time.Unix(1510394881, 0)}
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

	info := bindataFileInfo{name: "level_normalization.md", size: 23, mode: os.FileMode(436), modTime: time.Unix(1510425266, 0)}
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

	info := bindataFileInfo{name: "maps.md", size: 78, mode: os.FileMode(436), modTime: time.Unix(1511287031, 0)}
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

	info := bindataFileInfo{name: "rawsection", size: 40, mode: os.FileMode(436), modTime: time.Unix(1510405192, 0)}
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

	info := bindataFileInfo{name: "regression.md", size: 28, mode: os.FileMode(436), modTime: time.Unix(1511040126, 0)}
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

	info := bindataFileInfo{name: "scalar_decoder.md", size: 175, mode: os.FileMode(436), modTime: time.Unix(1511033510, 0)}
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

	info := bindataFileInfo{name: "test.md", size: 45, mode: os.FileMode(436), modTime: time.Unix(1510423269, 0)}
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
	"bindata.go": bindataGo,
	"codearray.md": codearrayMd,
	"codecomment.md": codecommentMd,
	"commentcode.md": commentcodeMd,
	"generate.go": generateGo,
	"level_normalization.md": level_normalizationMd,
	"maps.md": mapsMd,
	"rawsection": rawsection,
	"regression.md": regressionMd,
	"scalar_decoder.md": scalar_decoderMd,
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
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
	"codearray.md": &bintree{codearrayMd, map[string]*bintree{}},
	"codecomment.md": &bintree{codecommentMd, map[string]*bintree{}},
	"commentcode.md": &bintree{commentcodeMd, map[string]*bintree{}},
	"generate.go": &bintree{generateGo, map[string]*bintree{}},
	"level_normalization.md": &bintree{level_normalizationMd, map[string]*bintree{}},
	"maps.md": &bintree{mapsMd, map[string]*bintree{}},
	"rawsection": &bintree{rawsection, map[string]*bintree{}},
	"regression.md": &bintree{regressionMd, map[string]*bintree{}},
	"scalar_decoder.md": &bintree{scalar_decoderMd, map[string]*bintree{}},
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

