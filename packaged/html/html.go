// Code generated by go-bindata.
// sources:
// index.html
// registration.html
// DO NOT EDIT!

package html

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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x57\xfd\x6e\xd4\x46\x10\xff\x1f\x89\x77\x18\x4c\x84\xef\xc4\x9d\x9d\x84\xb4\x54\xe1\xee\x50\x94\xa0\x2a\x12\x2a\xa8\x89\x84\x50\x3f\xa2\x8d\x3d\xbe\x5d\x62\xef\x9a\xdd\xf5\x85\x13\x45\xe2\x41\x5a\x89\x67\xe1\x51\x78\x92\xce\x7a\xed\x9c\x9d\x4b\xda\xa3\xa8\xd2\x9d\xbc\x1f\xf3\xf9\x9b\x0f\x8f\x27\xf7\x8e\x5e\x1c\x9e\xbe\x7e\xf9\x0c\xb8\x2d\xf2\xd9\xdd\x3b\x13\xf7\x84\x9c\xc9\xf9\x34\x40\x19\xd4\x27\xc8\x52\x7a\x02\x4c\x0a\xb4\x8c\x08\x6d\x39\xc6\xb7\x95\x58\x4c\x83\x43\x25\x2d\x4a\x3b\x3e\x5d\x96\x18\x40\xe2\x77\xd3\xc0\xe2\x3b\x1b\x3b\x41\x4f\x20\xe1\x4c\x1b\xb4\xd3\xca\x66\xe3\x1f\x02\x88\xbd\x20\x2b\x6c\x8e\xb3\x63\x1b\x1a\x78\xad\x2a\x78\x21\x73\x21\x71\x12\xfb\xe3\x9a\x82\x0e\x2e\x40\x63\x3e\x0d\x8c\x5d\xe6\x68\x38\xa2\x0d\x80\x6b\xcc\xa6\x01\x33\x24\xd1\xc4\x89\x31\x71\x7d\x19\xd1\x2a\x88\xbf\x8a\x2f\x23\x53\xc7\xec\x12\x8d\x2a\x30\x2a\x84\x5c\x13\x51\x73\x84\xce\xd7\xfd\xb8\xa6\x36\xd1\x5c\xa9\x79\x8e\xac\x14\x26\x4a\x54\xe1\xc4\x3c\xcd\x58\x21\xf2\xe5\xf4\x15\xe6\x79\x96\x93\xa6\x3f\x0e\xf4\x42\xed\xef\x6d\x6f\x8f\x1e\xd3\x9f\x9e\xc2\xb2\x5c\x24\x6e\xe7\x57\x61\x6d\x5b\xb8\xb2\x2d\x04\x4b\xd8\x4d\xc3\x1a\x32\x12\x19\x3a\xc4\xe3\x06\xf2\xc9\xb9\x4a\x97\x6d\x08\x50\x83\x48\xa7\x81\x5f\x06\xde\xd4\x54\x2c\x20\xc9\xc9\xaf\x69\xd0\x80\xef\x2f\x9a\x2b\x47\x9f\xab\xb9\x0a\x66\x13\xd6\x60\x70\x3f\x98\xc1\xb1\x35\x84\x7a\xe4\x51\x87\x49\xcc\x66\x93\x98\xc8\xbb\x9c\x8d\x50\x23\xe6\x72\x7c\xa9\x59\xd9\x8a\xa5\xeb\x56\x92\xc6\xb9\x30\xd6\x99\x32\x11\x2d\x7d\xc6\x20\x63\xe3\xca\xa0\x1e\x97\x79\x65\xe8\x2a\x16\xb3\x13\x12\x02\x55\xe9\xf4\xac\x09\x21\xeb\x84\x5c\x97\x90\xab\xe4\xc2\x33\x3f\x57\x73\x10\x72\xc5\xbb\xb2\xb4\x5d\x79\xb4\x50\xd3\x92\x36\xf7\xc6\x63\x38\x79\x7e\x7c\xf4\xec\x04\x4e\x4e\x0f\x7e\x3e\x85\xf1\xd8\xd1\xb4\x68\x98\x5c\xa4\xb8\xf3\xef\xe8\xf1\xdd\x19\x41\xa4\x41\x79\x8c\x88\x49\x52\x72\x2e\x49\xd7\x6e\x87\x44\x48\x58\xd6\x54\x97\x12\x38\x93\xa9\xb9\xba\x5f\x19\xd7\x3c\x7b\x06\xec\xde\x6e\x00\xdc\x10\x85\xb7\x95\xb2\x68\xce\x1c\x09\x23\x6b\x74\x27\x18\xfc\x51\x9f\x28\x98\x35\x51\x2d\x35\x1a\x94\x09\x42\x81\x4c\x1a\x82\x90\x22\xc5\x12\x2b\x94\x8c\xe0\x50\x15\x25\x93\x02\x0d\x48\xc4\xd4\xbb\xd0\x7a\x08\x97\x5c\x24\xdc\x9d\x35\x22\x2c\xf1\x42\x2a\xb2\x0c\xb5\xdb\x5c\xb2\xa5\x19\x41\x8a\x25\xca\x54\xc8\x39\x01\x04\x96\x23\x95\xbe\x13\xb9\x1c\x41\x89\xda\xd0\x99\xd2\xbe\x1b\xbc\xb3\x11\x61\xf2\xe8\x1f\x82\x77\x13\x3e\xdf\x6d\x1a\x20\xa7\xb7\xa8\xa4\x48\x98\x45\x50\x0b\x2a\x91\x05\xd3\x42\x55\xc6\xf5\x1d\x29\x31\x37\x40\x61\x59\x8f\xda\xca\x9f\xc6\x4a\x03\xe7\x15\x79\xea\x1a\x12\xcb\x9d\x8f\x0e\x81\x0d\xa3\xf9\xfd\xff\x18\xcd\x2f\x1f\xff\x3c\xe5\xe8\x31\xfe\x51\x0b\x49\xa1\xb1\x5c\x55\x73\x6e\x41\x65\xe0\xba\x97\xe5\x2e\x0c\x74\xcd\x59\x2a\x43\x0b\xe7\x98\x29\x8d\xf7\xe0\x15\x67\xe4\x50\x06\x87\x5c\x53\x99\x16\x8c\xa2\x46\x44\x0d\x33\x45\x50\xa1\x71\xe4\x04\x20\x42\xa6\x55\x01\x0c\x8c\x25\xce\x68\x9d\x33\x8a\x22\x0a\x2b\x67\xa5\x5b\xf9\x7c\x62\x90\x0b\x4b\xcd\x1a\xce\x85\x85\xc2\x29\xfc\xf2\xf1\xaf\x4e\xa4\xc9\x19\x51\xcc\xc1\xe8\xe4\xaa\xed\xd2\x3e\x4e\xf5\xd8\x60\x65\x4c\x54\xca\x79\x40\x48\x53\x67\xa0\x1e\x80\x99\x6d\x7a\xaf\x47\x61\xaf\x45\x81\x55\x64\xaf\x3e\x93\xac\xc0\xb3\xb9\x66\xcb\x60\x76\xa4\x23\x38\x71\x12\x28\x26\x7c\x6f\xc5\x73\xae\xe3\x5b\x36\xeb\x98\x1e\xc8\x6e\xb6\x13\xb8\xab\x24\x72\x50\x0a\x2a\x0b\x65\x81\x9d\x2b\xca\x88\x3a\xb7\x7d\x2a\xb9\x9c\x76\xdb\x52\xab\x05\xf1\xeb\x51\x93\x2e\x35\xdd\x25\x57\x75\xcd\x30\x4d\x18\xb3\xfc\xc2\x09\xb2\xca\x95\x8e\xaf\xa7\x6f\x28\x86\x47\x1b\x15\xc3\x09\x77\xaa\xab\xd2\x69\x4d\x5d\x35\x08\x49\x79\x50\x30\x57\xf1\xd7\xf2\xdf\x65\x3a\xf5\xb5\x25\x30\x03\x45\x45\xc6\xb1\x3a\xdb\xa9\xb4\xa5\xdd\x30\xe5\xf7\x36\x7b\xff\xfc\x97\x8c\x3f\x52\x3e\x8b\x81\x65\x9a\x89\xd4\xf9\x63\xb8\xba\xf4\x5d\x2a\x63\xd4\xd1\xa8\xcf\x14\x75\xa2\x7a\x9f\x5d\x4c\x3a\xce\x7a\x57\x84\xe1\x9e\xb3\xa6\x50\x2e\xb9\x51\x2f\x44\xb2\x0a\x9f\x6b\x0c\x4e\x8d\xa9\x88\xc2\x77\x71\x92\xa4\x7b\xa2\x28\x15\x0c\xcb\x70\xb3\xa8\x75\xde\xb9\x35\x46\x8f\x6f\x82\xe2\x1a\x4a\x9b\xe2\xd4\xa7\xa3\x00\x5a\x25\xc7\x09\xba\xa6\x0e\x2e\xaf\x9a\x75\x97\xa1\xcf\x92\xaa\xa4\x7f\x59\xbf\x83\x2d\xd3\x73\x9a\xcd\x82\xb3\x73\x1a\xf8\x2e\x82\xd5\x88\xd0\xa7\x04\x38\x52\x49\x55\x90\x8a\x1a\x96\x6b\x62\x3a\xaf\xf5\x3e\x40\xb7\xd8\x41\xad\x27\xd9\xfd\x16\x63\x0e\x5e\x1e\x7f\x8d\x41\xeb\x26\xf5\xf7\xbd\xd9\xe7\xf6\x9a\xec\x0c\x16\xcf\x7e\x3a\x6a\xc6\x0a\x93\x68\x51\x5a\xdf\xe6\x9a\x59\x31\x51\x29\x46\x6f\xde\x56\xa8\x97\xf5\x98\xe8\x97\xe3\x9d\x68\x67\x37\xda\xae\x47\xcd\x37\xf5\x5c\xe4\x59\x57\x42\x68\xb5\x35\x48\x1b\xb7\x86\x91\xa6\x99\x66\x39\xc8\x2a\x59\xbf\xb3\x07\x43\x78\x7f\xf7\xce\xe7\x4f\xb0\x35\x08\xef\x4b\xb6\x00\x16\x0e\x23\x64\x09\x5f\xa7\xa0\x1f\xf5\xee\x01\x4d\x50\x35\x36\x51\xc9\x2c\x77\x0d\x94\x24\x96\x39\x95\xcf\x20\xfe\xfd\xd7\x38\x1e\x85\xe1\x10\xa6\x53\x2a\x1e\x9a\x66\x6f\x27\x69\x25\x3e\x78\x00\x57\x02\xb9\x32\xd6\x51\x5f\xb1\xb7\x07\x1d\x62\x7f\xce\x0c\x5f\x89\xbc\xef\x75\x76\xcc\xa4\x1f\xbd\xad\x61\xcb\x07\xfe\x38\x85\x29\xb9\x77\xc5\x39\x1c\xb5\x37\x07\xf4\xd6\xa3\xde\xeb\x6e\xc3\x5f\x9c\xa2\x69\x08\x0f\x3b\x2a\xa8\xda\x48\xc1\xce\x10\x1e\x86\xbf\x85\xc3\x27\x37\x8b\x77\xec\xad\xa2\x28\x47\x39\xb7\x1c\x9e\x76\x74\xef\xf7\xb5\xad\x91\x34\x46\xec\x53\x0b\xca\x0d\x76\x95\x7c\xfe\xe4\xf0\x6e\xc8\xfa\xee\xd5\x3f\x67\x84\xbf\x7c\x91\x65\xa6\x6b\x49\xa4\xea\x83\xc1\x30\xb2\xaa\x7c\xe2\xc6\xd7\xeb\xbc\x1e\x8f\x61\x94\x90\x8b\x17\xd7\x62\xdd\xcd\xf0\xad\x41\x50\xa7\x45\x2e\x80\x05\x2e\x79\x0a\x1a\x86\x0e\x5d\xc9\x0d\x02\x37\xf5\x2d\x30\x70\xc0\xf4\x59\xbc\x68\x96\xa6\x9e\x30\xf4\x84\x7d\x04\x57\x86\x84\xee\x7b\x6e\x04\xee\x8b\x84\x72\x8f\xc6\x47\xea\x91\x38\x78\x4f\xc9\xab\xf2\xfc\x54\x95\xfb\x3d\x27\x3f\x8c\x60\x67\x7b\x7b\xfb\x46\x59\x1a\x6d\xa5\xe5\x4d\x40\xd2\xef\x43\x8f\xe5\x43\xbb\xf6\x8b\xfa\xf2\xee\x9d\xfa\xd1\xad\xa0\xb8\xfd\x4e\x8a\x9b\xaf\xd8\xbf\x03\x00\x00\xff\xff\x6c\xaf\x8d\x26\xd7\x0e\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 3799, mode: os.FileMode(493), modTime: time.Unix(1455275674, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _registrationHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x6b\x6f\xdb\x36\x17\xfe\xfc\x16\xe8\x7f\xe0\xcb\x0d\x53\x03\x44\x92\x73\x71\xe2\xa6\x96\x8b\xa0\x17\xa0\xc0\x8a\x06\x5b\x8a\xa1\xd8\x8a\x82\x96\x28\x8b\x09\x25\xaa\x24\x65\xc7\xbb\xfc\xf7\x1d\x92\x92\x62\x2b\x72\x1c\x6f\x2d\xd0\x0f\x33\x10\x84\xa4\x9e\x73\x7b\xce\x11\x79\xa8\xf1\xff\x5f\xbe\x7b\x71\xf9\xe1\xe2\x15\xca\x74\xce\x27\x8f\x1f\x8d\xcd\x7f\xc4\x49\x31\x8b\x30\x2d\x30\xac\x20\x34\xce\x28\x49\x60\xf4\x3f\x04\xe3\x9c\x6a\x02\x60\x5d\xfa\xf4\x73\xc5\xe6\x11\x7e\x21\x0a\x4d\x0b\xed\x5f\x2e\x4b\x8a\x51\xec\x66\x11\xd6\xf4\x46\x87\x46\xd9\x33\x14\x67\x44\x2a\xaa\xa3\x4a\xa7\xfe\x08\xa3\xd0\x2a\xb5\xbf\xb1\x66\x9a\xd3\xc9\x1b\xed\x29\xf4\x41\x54\xe8\x5d\xc1\x59\x41\xc7\xa1\x5b\x7e\xfc\xe8\x16\x08\xeb\xd7\x48\x52\x1e\x79\x4a\x2f\x39\x55\x19\xa5\x1a\x95\x92\xa6\x54\xc7\x99\x87\x32\x18\x45\x9e\x71\x4b\x9d\x85\x21\xb9\x22\x37\xc1\x4c\x88\x19\xa7\xa4\x64\x2a\x88\x45\x6e\xd7\x42\xce\xa6\x2a\x84\xd8\x2a\x4e\xe4\xa7\x9c\x68\x2a\x19\xe1\xe1\x41\x30\x08\x06\xcd\xb2\xdf\x2c\x07\x39\x2b\x82\x58\x29\x6f\xd2\xe3\x05\xbe\xf5\x02\x3b\xe3\x98\x28\x88\x51\x85\x20\x11\xda\x87\x46\x16\x87\xbb\x0b\xa7\xc0\xa0\x4f\x16\x54\x89\x9c\x36\x3e\xdc\xd5\x73\x1b\x30\xc4\x6b\x44\x54\x37\x60\x10\x7b\x9e\x92\x9c\xf1\x65\xf4\x0b\xe5\x3c\xe5\x60\xee\xcf\x73\x39\x17\x67\xc7\x83\xc1\xfe\x29\xfc\xc1\x7f\xa6\x09\x67\xb1\x99\xb9\x91\xd7\xe5\xd8\x43\x1a\xf2\x1a\x79\x36\x9d\xeb\x6c\xa0\xb1\x85\xad\xa6\xc9\xfe\x7e\x2d\x66\xbf\x9d\xc5\x5c\x90\xeb\x8f\xfb\x66\xe2\x37\xe3\xa0\x19\xa3\x3f\xd6\x25\xcc\x2f\x61\xaa\xe4\x64\x79\x86\x0a\x51\xd0\x67\xeb\xcf\xff\x5a\xb1\x19\x36\x46\x61\x58\xd7\x25\x54\xed\x54\x24\x4b\x04\xea\x49\x59\x46\x58\xd2\x19\x53\x5a\x12\xcd\x44\x71\x5e\x96\xd8\x41\x0c\x98\x4a\xc4\x92\x08\xbb\x61\x5d\xdc\x09\x9b\xa3\x98\x43\x02\x22\x5c\x17\x2f\xae\x63\xb4\x8f\x0c\x9e\x8b\x99\xc0\x93\x31\xa9\x93\x95\x41\x66\xf0\x04\xbd\xd1\x0a\x6a\x36\x70\x35\x0b\xde\x90\xc9\x38\x04\x89\x55\xe1\x5a\xaf\x62\xb3\xc2\x5f\x48\x52\xe2\x96\xbd\x56\x19\xe8\x66\xf0\x9e\x8d\x59\x03\x4e\x09\x4a\x89\xcf\x45\x7c\x0d\xab\x21\x9b\xfc\x28\x66\x88\x15\x46\x7d\xad\xb8\x35\xd2\x8c\x1c\x11\x54\xda\x38\x61\xba\x62\x39\xad\x38\x57\xb1\xa4\xb4\x48\x85\xcc\xdb\xc0\xcc\x04\xb5\xd9\x80\x57\x3a\x13\x10\x66\x29\x14\x14\x24\x89\x0d\x6f\x0d\x8b\x40\x13\xec\x06\x4b\x51\x69\x43\x0f\xaf\xf2\x62\x8d\x1d\xa8\xaa\x9b\xf6\xb9\x14\x0b\x8c\xea\x99\x7f\xa3\xba\x02\x6b\x42\x11\x1e\x0e\x36\x69\x6e\xc1\x79\xe2\xb3\xa2\x04\x65\x26\x31\x04\x58\x96\x93\xbb\x85\x33\xe6\x64\x4a\xf9\xe4\xbd\xa2\xb2\x20\x39\x6c\x1e\x6e\xde\x03\xb4\xba\x4c\xd8\xb9\x48\xcc\x6b\xe8\xa8\x47\x46\xaa\x9d\xd8\x6a\xb7\x9b\x17\x30\x51\x69\x91\x8a\xb8\x52\x5d\xbf\xc2\xad\x8e\xad\xe5\x00\xc8\xf6\xa9\x94\x02\xa8\x04\xe3\x6a\xc1\x60\xcb\x72\xcb\xaf\xcc\xaa\x0a\x1a\x47\x9a\x84\x18\x54\x26\x16\x3d\x98\xbe\xb0\x8c\xa9\x56\xaf\xbf\xc8\xa8\x4d\x1e\x6c\xce\x92\x26\x78\xf2\x53\x3d\x0a\x56\x8b\xf3\x36\x12\xb7\xf6\x2f\x78\x7f\x95\x13\xc6\x77\x20\x9d\x1a\x7c\x43\x7a\x3d\x71\xa4\xbb\xc9\xd7\xe6\xba\x04\xdc\x42\xc8\x64\x0b\xdd\x2d\xec\xdb\x63\xfc\xa2\x76\x6d\x07\xd2\x57\x82\xb6\xbc\xdf\xce\x1d\xf5\x9b\xa2\xfd\x8f\xfd\x8d\xec\x9b\x4e\x27\x65\x32\xb7\xa7\xcc\x3f\x48\xc5\x1c\x8e\xdb\xc4\x0a\x77\x93\xb2\xfa\xe4\xdb\x4e\xcf\x8a\xa7\x5f\x3d\x51\x5d\xd4\xdd\xe3\x67\xc3\x21\xd2\x7b\x4c\xf5\xb9\xab\x4a\x52\x58\x24\x1c\xbb\x66\xbc\x29\xa4\x8e\xbd\xfa\x38\x3b\x19\xf4\x29\xb5\x42\x0f\x2a\xaf\x16\xed\x2a\xe9\xd0\x7f\x0d\x07\xb1\x90\xe6\x10\x02\xe2\x34\x8b\x2d\xcf\x50\x75\xc9\x3d\x47\x5c\xab\xa4\x5b\x75\x5a\xe8\xd2\x88\x36\xb5\xd6\xce\x37\xf9\xbc\xbd\xb2\x5a\xe8\x67\x69\x34\xa1\x39\x95\xca\x36\x0e\x27\x18\xd9\x22\x03\x49\x29\xa9\xed\x26\x7c\x4e\xe7\xc6\x8d\xb7\x18\x29\xf6\x3b\xd8\x3f\x1c\xc0\xe9\x0f\xa5\x43\x22\x0c\x9e\x98\x18\xa1\x8d\x35\x4e\x85\x75\x57\x25\x6c\x57\xf5\x5c\x51\x68\x5e\x74\x74\x70\x78\x74\xbc\xfa\xf7\x03\x53\xaa\xa2\x32\x5a\x6b\xc1\x9a\x5e\xfa\x3e\x85\x67\xbb\x69\x84\x4a\x70\xd1\xf5\xd5\x42\x4f\xd9\xda\xf5\x2d\x75\xb4\xb9\xb4\xd7\xfb\xbb\x4e\xa9\xdf\x53\xbd\x5b\x2d\x42\x2e\xa7\x95\xd6\x50\x3e\x6e\x3f\x51\xd5\x34\x67\xd0\xe0\xd4\xfb\x02\x3c\x96\x84\x29\x9a\x20\x18\x95\x92\xe5\x44\x2e\xf1\xe4\x67\x68\x5a\x51\x55\xda\x4a\x70\xd2\x3b\x98\x5d\xeb\x84\x43\xb3\x75\xd8\x3e\xb5\x89\x0c\xc6\xd0\x95\xb2\x52\x23\x25\xe3\x1d\xef\x6d\x57\x0a\x2e\x6c\xc7\xc1\xa8\x99\x07\x57\x70\x29\x01\x07\xac\xbe\xc9\x97\xd4\xec\x93\x82\x99\xdb\xa0\xbd\x85\x7d\x45\x2b\x70\xdb\xfc\xb2\x26\x1e\x74\xb5\xed\xb7\x56\xe7\xac\xe1\x16\xb6\x8f\x8a\xd3\x27\x5e\xe7\x4a\xe5\xc1\xc5\xce\x2b\x66\x6f\x6b\x7d\xde\xbe\x97\x8b\x42\x40\x0d\xc4\xb0\x9f\xbb\x37\xc6\xfb\xb8\x77\x5b\x0e\x81\x3d\x2c\x67\x4f\xd2\xaa\xb0\xfb\xc1\x93\xef\xf3\xe4\x32\xa3\xe0\xc6\xec\x42\x8a\x39\x83\xcb\xcb\x5e\xf7\x52\x78\x17\x12\x24\x34\x85\xb7\xf2\x82\x70\xaa\x35\x38\x35\xe5\x15\x65\x2a\x03\x67\x7a\xee\x93\xde\x70\xe0\x9d\x21\xef\xbb\xf4\x34\x9d\xa6\x89\xb7\xdf\x83\x38\x18\x38\xc8\x14\x2e\x4f\xb4\x1f\x72\x58\x43\x46\x34\x1e\xd1\xc3\x5e\xc8\x51\x0d\x19\x0e\x49\x92\x1c\xf5\x42\x8e\x6b\xc8\x51\x4c\x0e\xe2\x7e\x43\xc3\x06\x32\x78\x7a\x30\x9d\xf6\x42\x4e\x6a\xc8\x21\x39\xa5\xa4\xdf\xd0\x69\x03\x39\x3a\x99\x8e\x48\x2f\x64\x54\x43\x0e\x92\xe1\xe8\xb4\x3f\xa2\xa7\x0d\xe4\xf4\x18\x82\xea\x85\x9c\x37\xdc\xdd\x43\xef\xf9\xe1\x76\x7e\xcf\x1f\x40\xcd\xf9\x03\xa2\x32\x87\x94\x24\x4a\xbf\xa4\x29\xa9\xb8\x7e\x21\xb8\x90\x46\x86\xb3\x59\xa6\xef\x97\x20\xf2\xda\xc2\x95\xc1\x0f\x07\x08\x22\x43\xe0\x39\x82\xbc\x22\xf0\x0e\x99\x50\x91\x89\x05\x59\x67\x3b\x5f\x26\xf6\x9e\x6d\xad\x5a\x38\xbf\x73\xa8\xd6\xc4\xb9\xe6\xed\xdd\x75\x26\xa8\x37\xdd\x3b\x95\xbd\x82\xb5\x96\x36\xec\x0e\x58\x67\x4c\x26\x25\x91\x7a\x59\x7f\x48\x72\xef\xa0\x3f\xa3\x70\x6e\x13\x68\x23\xc2\xab\x66\x0d\xde\x7b\xbc\x69\x97\xd9\x45\xcf\xa7\xf7\x97\xaf\x47\xbb\x29\x6b\xb6\x21\xa7\xa0\x33\xed\xaa\x32\xc1\x9a\xaf\x3a\xee\xfb\x86\xfb\x38\xf9\x77\x00\x00\x00\xff\xff\x4b\x00\x7e\xd3\xae\x14\x00\x00")

func registrationHtmlBytes() ([]byte, error) {
	return bindataRead(
		_registrationHtml,
		"registration.html",
	)
}

func registrationHtml() (*asset, error) {
	bytes, err := registrationHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "registration.html", size: 5294, mode: os.FileMode(420), modTime: time.Unix(1455545263, 0)}
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
	"index.html": indexHtml,
	"registration.html": registrationHtml,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
	"registration.html": &bintree{registrationHtml, map[string]*bintree{}},
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

