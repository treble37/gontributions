// Code generated by go-bindata.
// sources:
// templates/default.html
// templates/detailed.html
// DO NOT EDIT!

package main

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

var _templatesDefaultHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\xc1\x6e\xc3\x20\x0c\x86\xcf\xdd\x53\x78\x7d\x80\xa0\x5e\x2b\xc6\x25\xd7\x69\xda\x2b\x10\x70\x03\x55\x82\x23\x30\x9a\xaa\xaa\xef\x3e\x13\x36\xa9\xdb\xc9\xd8\xfe\xfd\xfd\xe2\xd7\xaf\x9e\x1c\xdf\x36\x84\xc0\xeb\x62\x5e\x74\x2f\x07\x1d\xd0\xfa\x56\xd5\xef\x63\x22\x7f\x93\x2a\x9b\x93\x79\x8f\x85\x81\x2e\x40\x1b\x26\x28\x54\xb3\x43\x70\x94\x38\xc7\xa9\x72\xa4\x54\xe4\xec\xb4\x8b\x6b\x83\x1d\xee\x77\xc8\x36\xcd\x08\xc3\xf8\xac\x82\xc7\xa3\x69\x96\x68\x44\x30\x7c\x66\xba\xa2\xe3\xe1\xc3\xae\x28\x9b\x33\xb4\xe1\x48\x35\xb1\x74\xff\xf1\x72\xd3\xb9\x98\xfc\x0f\x46\x75\x2f\x1d\xcd\x98\xd1\x32\x7a\xa0\xd4\x18\x5b\x65\x2f\x6d\x83\x7c\x45\x0e\xa0\x2d\x84\x8c\x97\xb7\x63\x60\xde\xca\x59\xa9\x59\xa6\x75\x1a\x1c\xad\xea\x5a\x27\xbb\x04\x35\x3f\x9b\x1d\xcd\xfc\xd7\xdb\x1a\xad\xe2\x9e\x4d\xcf\x44\x3e\xbb\x87\xf6\x1d\x00\x00\xff\xff\x73\xf7\x4b\x23\x4c\x01\x00\x00")

func templatesDefaultHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDefaultHtml,
		"templates/default.html",
	)
}

func templatesDefaultHtml() (*asset, error) {
	bytes, err := templatesDefaultHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/default.html", size: 332, mode: os.FileMode(420), modTime: time.Unix(1469203179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDetailedHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x92\xdf\x6a\xeb\x30\x0c\xc6\xaf\xcf\x79\x0a\xad\x0f\x10\xd3\xdb\xe1\xf9\xa6\xbb\x2c\x63\x0c\xf6\x00\xfe\xa3\xc6\x2e\xad\x1d\x6c\x99\x51\x4a\xdf\x7d\x76\xb2\x04\x27\xf4\x4a\xca\xf7\xfd\x24\x19\x29\xfc\xc5\x04\x4d\xb7\x01\xc1\xd2\xf5\x22\xfe\xf3\x29\xfc\xe3\x16\xa5\xa9\x91\xcd\x89\x0a\xe6\x56\x62\x71\xf6\xe2\xe8\x12\x41\x38\x41\x18\xd0\x43\x0a\x39\x6a\x04\x1d\x3c\x45\xa7\x32\xb9\xe0\x53\x29\xdb\x57\xf8\x7e\x87\x28\x7d\x8f\xd0\x1d\x5a\x1b\x1e\x8f\xc9\x74\x27\xe8\x3e\x63\x38\xa3\xa6\xee\xfb\xeb\x38\xe9\x5c\x09\x2e\xc1\x46\x3c\xbd\xed\x0a\xb3\x01\x76\xa2\xd5\x3e\xe4\x15\x8b\xc8\x99\x14\x9c\xa9\xbf\x91\x78\x49\xb8\xb4\x7a\x4a\x2f\xa4\x37\x13\x58\xd1\xc8\xc4\x98\x6d\xde\xf5\x8e\x49\x47\x37\xd4\x87\xcf\x6c\xdb\xf3\x89\xbd\x6a\xd5\x8c\xa8\x65\x87\x90\x3d\x15\x61\xbd\xaf\x4d\xd9\x9c\xb5\xe5\xdc\x89\x43\x44\x49\x68\xa0\x4c\x2a\xce\x90\xc9\x94\xcf\xda\xeb\xc7\x91\x85\x65\x67\x96\x68\x48\xaf\x8c\xf5\x45\xcd\xaa\xd3\xe1\xca\xce\x59\xc9\x8b\x65\x7d\x3b\x73\x27\xfa\xf5\xc9\xea\x0a\xdd\x78\xf3\xe9\xd6\xe5\x88\xe3\xcf\xf0\x1b\x00\x00\xff\xff\x6a\x3f\x76\x91\x24\x02\x00\x00")

func templatesDetailedHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDetailedHtml,
		"templates/detailed.html",
	)
}

func templatesDetailedHtml() (*asset, error) {
	bytes, err := templatesDetailedHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/detailed.html", size: 548, mode: os.FileMode(420), modTime: time.Unix(1469203306, 0)}
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
	"templates/default.html": templatesDefaultHtml,
	"templates/detailed.html": templatesDetailedHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"default.html": &bintree{templatesDefaultHtml, map[string]*bintree{}},
		"detailed.html": &bintree{templatesDetailedHtml, map[string]*bintree{}},
	}},
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

