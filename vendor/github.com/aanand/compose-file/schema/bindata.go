// Code generated by go-bindata.
// sources:
// data/config_schema_v3.json
// DO NOT EDIT!

package schema

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

var _dataConfig_schema_v3Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\xcd\x8e\xdb\x36\x10\xbe\xfb\x29\x16\x4a\x6e\xf1\xae\x03\x34\x28\xd0\xdc\x7a\xec\xa9\x3d\xd7\x50\x04\x5a\xa2\x6d\x26\x92\xc8\x90\x94\x13\x27\xf0\xbb\x97\x14\x25\x99\xa4\xf8\x67\x5b\x69\x7a\xe8\x1e\x16\xbb\xe2\xcc\x70\x7e\x3e\xce\x0c\x47\xfa\xbe\x7a\x7a\xca\x5e\xb3\xf2\x08\x1b\x90\xbd\x7f\xca\x8e\x9c\x93\xf7\x9b\xcd\x47\x86\xdb\x67\xf5\xf4\x05\xd3\xc3\xa6\xa2\x60\xcf\x9f\xdf\xbe\xdb\xa8\x67\xaf\xb2\xb5\xe4\x43\x95\x64\x29\x71\xbb\x47\x87\x42\xad\x14\xa7\x5f\x5e\x24\xb3\x22\xe0\x67\x02\x25\x09\xde\x7d\x84\x25\x57\xcf\x28\xfc\xdc\x21\x0a\x25\xeb\x36\x3b\x41\xca\x90\xa0\xce\xd7\x2b\xb9\x46\x28\x26\x90\x72\x04\x99\x58\x95\xaa\x89\x67\x23\xc9\xf8\x40\x13\xcb\x38\x45\xed\x21\xeb\x1f\x5f\x7a\x09\x62\x91\x41\x7a\x42\xa5\x26\x61\x52\xf4\xd5\xe6\x2a\x7f\x33\x91\xad\x6d\xa9\x9a\xb2\xfd\x73\x02\x38\x87\xb4\xfd\x6b\xae\x5b\xbf\xfc\x61\x0b\x9e\xbf\xfd\xfe\xfc\xf7\xdb\xe7\xdf\x5e\x8a\xe7\xfc\xcd\x6b\x63\x59\x7a\x97\xc2\xbd\xda\xbe\x82\x7b\xd4\x22\x2e\xac\x99\xf6\xcf\x26\xca\xcb\xf0\xd7\x65\xda\x18\x54\x55\x4f\x0c\x6a\x63\xef\x3d\xa8\x19\x34\x6d\x6e\x21\xff\x82\xe9\xa7\x98\xcd\x13\xd9\x4f\xb2\x79\xd8\xdf\x61\xb3\x69\xce\x09\xd7\x5d\x13\x8d\xe0\x48\xf5\x93\x8c\x51\xdb\x3f\x16\xbf\xd5\x68\x74\x90\x56\x51\x68\x7b\xf7\x0a\x1a\x68\x77\xb9\xca\x85\x36\xbf\xaf\x26\x67\x79\xbc\x54\x41\x52\xe3\xb3\x7c\xe6\xf1\x87\x22\x68\x60\xcb\xb3\xc9\x05\x82\x6f\xd7\xa1\xba\xb2\x3d\x8a\x5b\xf8\xa7\x14\xb1\xd5\x1e\x3e\x09\xc9\xd6\xc1\xd6\xe4\xf4\xeb\xc6\x7f\xfe\x80\x4f\xeb\x1e\x5b\xa6\x75\x91\xb9\x38\xfc\xca\x7b\xa3\xc2\x5b\x2b\x17\xe0\xf2\x13\xa4\x7b\x54\xc3\x54\x0e\x40\x0f\x2c\xe0\xb2\x1a\x31\x5e\x60\x5a\x54\x48\x68\x7f\xb1\xd8\x67\xf2\xe2\x78\x9a\x58\xb5\xff\xf2\x95\x43\x60\x56\x02\x52\x08\x71\x86\x1d\x80\x52\x70\xce\xd6\x02\x40\x1c\x36\xcc\x6d\xe2\x53\xd6\xb5\xe8\x73\x07\xff\x18\x48\x38\xed\xa0\x2d\xb7\x12\xca\x2d\x2f\xf8\x40\x71\x47\x0a\x02\xa8\x04\x58\xd8\xfd\x22\xae\x4d\x03\xda\xa5\x50\x77\x8b\x1d\x09\x9e\x17\x98\x03\xa8\x85\xb4\x68\x41\x13\x03\x92\x3c\x75\xb0\xad\x58\xa1\xea\x5f\x10\x46\xfb\x42\xf1\x33\x4b\xc0\x54\x0c\x17\x8d\x47\xd5\x86\x80\xad\xc4\x48\x68\x4b\xdd\x32\x8b\xb1\x60\x10\xd0\xf2\x78\x27\x3f\x6e\x84\xfb\x52\x7c\x27\x80\x42\xcf\x04\x23\x85\x97\xff\x1c\x10\x60\x7b\x2a\xa6\x5c\x72\xb3\x1b\x04\x37\xa2\xb8\x6d\xc6\xd3\x90\x92\x60\xa6\x24\x2f\xf9\xbf\x12\xcc\xa0\xed\x18\xcb\x40\x7d\x69\x32\xd5\xf0\xc9\xc8\xb1\x1d\x0d\x17\x4e\x69\xbb\x66\x07\xa9\x6c\xe9\x0c\xca\x3d\xa6\x0d\x90\xca\x8e\x7b\x6b\xcb\x86\xa7\x1d\xc8\xd3\x1d\xa8\xdb\x20\xcb\x3a\xa8\x85\x77\xda\x4f\xcb\x43\x5c\x88\xa7\xa0\x38\x62\xc6\xd3\x73\xb8\xc6\x2e\x19\x53\x60\x8a\x1a\x70\x88\x13\x91\x32\x46\x52\x83\x1d\xac\xef\xd2\x74\x51\xf7\x69\x62\xf1\xe1\x20\x49\x7d\x98\x99\xf5\x1e\xc3\x72\xac\x6a\x57\x14\x89\x3b\x41\x6a\x09\xc6\xe4\xda\x32\xd9\x8b\xf1\x16\x42\x29\x14\xec\x1f\x0d\xd2\x0f\x2f\xaa\x7d\x0c\x9c\x8b\xfe\xaf\xba\xce\x72\xbb\xe0\xcb\x9f\xf9\x33\xf3\x89\x65\x61\x5a\x4b\x60\x44\xa5\x01\xa5\xac\xfc\x14\x32\x4f\x5c\xaf\xa4\x43\xbb\x5e\x34\xb8\xf2\x01\x74\x46\x6c\xfb\xc6\x9b\x6b\x6f\x2e\x65\x3d\xdb\xcd\x1d\x60\x52\xe8\xa2\x57\x80\x88\x35\x3e\xf5\x52\xd5\xbc\xaa\x1b\x87\x58\x4f\x07\x6a\x04\x18\x8c\x1f\x76\xaf\x23\x0d\x69\x88\x9c\xde\x25\x62\xc2\xc5\xfb\x6b\x90\xd7\xc3\xea\x95\x99\xde\xe5\x46\x44\x5d\x55\xe9\x8f\x9b\x4b\x91\x3c\x72\xda\x7e\x70\x13\x4e\x50\xe5\xcf\x15\x7d\x86\xd0\x0f\x18\xc1\x94\xcf\x4e\xd7\xbf\x53\xb0\xd5\xd6\x0f\xd7\x6b\x22\x12\xb7\x68\x78\x0e\xd0\xbc\x77\xec\x30\xae\x21\x68\x8d\xd4\x43\x21\xa8\x44\xd3\x5b\x9f\x13\x28\x19\x07\x34\x7a\x25\x60\xb0\xec\x28\xe2\xe7\x42\xd4\x83\xc5\x3b\x05\x76\x6c\x0a\x86\xbe\x41\x33\x9a\xd7\x7c\x3f\x08\xca\x0d\x1e\x5e\xa1\x56\x68\x03\xdb\xa8\x89\x8c\x63\x22\xe4\x1f\x04\xe6\xa2\x66\x4a\xd2\x03\x05\x25\x2c\x04\x36\x11\xae\x5c\x0c\x6b\x3d\xb6\x55\x47\x81\xc4\xb3\x21\x86\x37\x64\x7f\x67\x7f\xcf\x79\x3c\x66\x5d\x8d\x1a\xe4\x07\xb3\x23\x4b\x26\x24\x72\x95\xc4\xdd\xb9\x3b\x90\xb7\xaf\x9a\x8a\x8b\x82\xc0\x26\x75\xa5\xbb\x40\xeb\x10\xee\x1c\x12\x5a\x86\x23\xa0\x66\x94\x02\x7a\xf4\x0c\x0c\xef\xb9\x9b\xc1\xd5\x50\x38\xf5\x32\x66\xb0\xbd\xbc\xf5\xa0\x48\xee\xa4\xbf\x29\x27\xdb\x6a\xe4\xde\xb4\x78\x71\xa6\xc5\x8e\x45\xbb\x3b\x7d\x42\xb8\xe8\x49\x96\x2d\x8c\x44\x76\x85\xdc\x2a\xac\x2c\x75\x6f\x98\xd1\x6a\xf3\xb1\xc8\xb0\x4e\xa3\xb4\xe7\x75\xdb\x09\x6e\x63\x8d\xb8\x4e\x39\x3d\x83\xbb\x40\xf7\xa6\xe1\x81\xd4\xa8\x04\x2c\x86\xc3\x07\x2e\x19\x1d\xa9\x00\x87\x85\x7a\x59\x70\xd3\xc9\x0f\x1c\x79\x02\x28\xa8\x6b\x28\x36\x6d\x52\x8e\x90\x88\x41\x0d\xce\x77\xa5\xc4\x9e\x7d\x0f\x50\xdd\x51\x58\x80\x92\x0f\x6f\x24\x22\xad\x92\x70\xbe\x70\x0c\x76\x62\x29\x6d\xcb\x06\x7c\x2d\xc6\x6d\x7b\x12\x43\xd4\x50\x62\x2e\xde\xd2\x9c\x7a\x3f\xd0\x90\xc0\x70\x47\xcb\x99\xb3\xef\x0e\xd1\x35\xd5\x7b\x10\x33\xee\x38\x33\x5d\x2c\x40\x7a\x02\xd3\xf5\x2d\xca\x1f\xcd\x2c\x43\xaf\x50\x10\x2c\xd0\x7e\x5e\xca\x42\x01\x69\xe5\xe4\x14\x40\x3c\x88\x40\x09\x07\x59\x09\x1b\xc2\xa3\x87\xb5\x67\xf8\x82\xda\x0a\x7f\xb9\x61\xc3\xe5\xa0\x44\x6a\xd1\x86\x58\xf9\xee\x51\x47\x0b\xdd\x81\x30\xf5\xe6\xc4\xff\xa8\x59\x0f\xe4\xfd\x09\x9f\x91\xac\x3f\xd1\xc5\xdf\x67\x79\x32\x7d\x49\xba\xe8\x9d\xbe\x81\x0d\xa6\x4e\x00\x06\x6c\x4c\x7c\xfd\x18\xb3\x70\x24\x5b\xa0\xa8\x25\xcd\x80\x06\x2a\xd9\xf2\x2f\xde\x6b\xc6\xe7\x3c\x79\x3c\x1f\x21\x02\x9a\xa5\x0e\x47\xf2\x54\x2c\x73\x96\x60\x63\xef\xf9\x6d\x52\xa9\xeb\xbc\x51\xc6\xb4\x8e\xeb\x3e\x50\xb0\x6e\x27\x10\x12\x82\xe6\xf5\xc7\xf9\xb2\x2d\xbd\x49\xbd\xf8\x5b\xd2\xc7\x72\xde\x38\x92\xf6\x44\x75\x3b\x5d\x87\xd6\x93\xaf\xf2\xe4\x10\x7b\xa7\xc9\xcb\xe9\x7f\x63\x7f\xf7\x40\x5a\x1c\x5e\x9f\x47\x52\xc6\x40\xf5\x7f\xc6\x18\xa4\xfc\x7c\x7c\x05\x6a\xe2\x9d\x97\x83\x1b\x40\x63\xcd\x1d\x34\xf0\xcc\x2f\xf8\xa1\x38\x27\x4f\x4d\x07\x8e\xdc\x54\xc3\x26\x7b\x3f\xff\x34\xc9\x4c\xa1\xa1\x2b\xe9\x48\xe2\x99\xa2\x59\x9b\x0e\xce\x0b\x5b\xbe\x20\x6c\x5f\xde\x04\x0a\x45\xe8\xed\xc6\x0f\xca\xb0\x0b\x5c\xf7\xdd\x31\xb5\x9a\xcb\xd1\xbb\xf3\xef\x6b\x3c\x99\x4a\xe3\x9f\x7d\x6d\x23\xed\x6c\xcf\xb3\x01\xd4\x77\x73\x0e\xa3\xbe\x94\xc9\x0d\xff\x58\x24\xea\x5d\xa1\x96\x27\x72\xbd\xdf\xf6\x85\xd1\xf9\x0d\x8e\x3d\x05\x1a\xbf\x85\xc9\xc3\x87\x7d\x35\xfe\xbe\xac\x2e\xab\x7f\x02\x00\x00\xff\xff\xe9\x77\x87\x68\x3d\x28\x00\x00")

func dataConfig_schema_v3JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v3Json,
		"data/config_schema_v3.json",
	)
}

func dataConfig_schema_v3Json() (*asset, error) {
	bytes, err := dataConfig_schema_v3JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/config_schema_v3.json", size: 10301, mode: os.FileMode(436), modTime: time.Unix(1478109572, 0)}
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
	"data/config_schema_v3.json": dataConfig_schema_v3Json,
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
	"data": {nil, map[string]*bintree{
		"config_schema_v3.json": {dataConfig_schema_v3Json, map[string]*bintree{}},
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
