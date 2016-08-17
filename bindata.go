// Code generated by go-bindata.
// sources:
// admin_http_assets/app.css
// admin_http_assets/app.js
// admin_http_assets/index.html
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

var _admin_http_assetsAppCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8a\xce\x4b\x8f\xb1\x4a\xce\xc9\x4f\xcc\x8e\xd5\x51\x00\x72\x74\xe1\xec\x94\xc4\x92\x44\x5d\x64\x81\x0a\x64\x9e\x1e\x8c\x0d\x64\x22\x24\x20\xe2\x19\x99\x29\xa9\x0a\xd5\x5c\x0a\x40\x90\x92\x59\x5c\x90\x93\x58\x69\xa5\x90\x97\x9f\x97\xaa\xa0\x98\x99\x5b\x90\x5f\x54\x92\x98\x57\x62\xcd\x55\x0b\x08\x00\x00\xff\xff\x9c\x9e\x21\xb0\x7a\x00\x00\x00")

func admin_http_assetsAppCssBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsAppCss,
		"admin_http_assets/app.css",
	)
}

func admin_http_assetsAppCss() (*asset, error) {
	bytes, err := admin_http_assetsAppCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/app.css", size: 122, mode: os.FileMode(420), modTime: time.Unix(1418065129, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _admin_http_assetsAppJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x58\xdf\x6f\xa4\xb6\x13\x7f\xcf\x5f\x61\xa1\x48\x18\x85\xb0\xf9\x7e\x1f\x2a\x95\x28\xba\xa6\xd7\x9e\x7a\x0f\x91\xaa\xf4\xda\x3e\xa4\x7b\x92\x03\x0e\x41\xf1\xda\xc8\x36\x9b\xac\xf6\xf8\xdf\x3b\x36\x06\x8c\x77\x2f\x97\xdb\x5c\xdb\x95\x22\xc0\xcc\x8f\xcf\x7c\x3c\x33\x1e\xb2\x26\x12\x91\xa6\x41\x17\x88\xd3\x47\x44\x78\xd5\x32\x22\xb3\x95\x28\x5b\x46\x71\x54\x10\x79\x2b\xf8\xa9\xa4\x8c\x6c\x4e\x79\x15\xa5\xe8\x26\xe2\xd5\x35\x55\xa2\x95\x05\x85\xc7\xa8\xad\xb3\x5b\x21\xb4\xd2\x92\x34\xd1\x32\x39\x3f\x3a\x02\x6b\x59\x21\xb8\x96\x82\x31\x2a\x71\x74\x45\x6a\xfe\x56\x33\xab\x7b\xac\x0a\xd1\x58\xbd\x63\xe9\x19\x39\x06\x7f\xc4\x48\xdc\xb5\xbc\xd0\xb5\xe0\xb8\x17\x4c\xd1\x28\x06\xb7\x56\x28\xd9\x1e\x21\xd4\xbf\xcd\x08\xd8\xd7\x0a\xa0\xdf\x2c\xcf\x61\xd5\x84\xf2\x56\xf0\xbb\xba\x82\xa5\x51\x11\x47\x8b\xc2\x2e\x2e\xa2\x64\x90\xfa\x40\x6e\x19\x0d\x84\xb4\x59\xf3\x64\xae\xe9\xa3\xac\x35\x95\x81\x98\x74\xcb\x6a\x91\xd7\xbc\xa4\x4f\x93\xc2\x8f\x8c\x14\x0f\xac\x56\x3a\xd0\xb8\x1d\xd6\x77\x55\x2e\xab\x4a\xd2\x8a\x68\x11\x7a\x21\xe3\x8b\x5d\xa5\x6b\xd1\xea\x10\xbc\x34\x6b\x20\xfa\x40\x37\xc0\xe2\x16\x2e\x39\x8a\x7f\x80\x4b\xdc\xc1\x63\x37\xea\xfe\x44\x95\xae\x39\x31\x14\x7f\xde\xc2\xa2\x9c\xa4\x7c\xf7\x47\x13\xf3\x6b\xc2\xea\xf2\xb2\x2c\xc1\x80\xe1\x7f\xf1\xf1\xe6\x63\xbe\x3c\xf9\x2b\xbf\x39\x3b\xfd\x7e\x79\x82\x73\xfb\x98\xbc\x39\x5e\x9c\x07\x3a\x16\xfc\x87\x4d\x43\xad\x16\x56\x94\x97\xf8\x92\xb1\x4f\xef\x6a\xa9\x74\x72\x45\x74\x71\x9f\x7c\xc2\xb0\x61\x0a\x08\xa3\x5c\xff\x42\xd4\x7d\xcd\xab\x64\x11\xda\xa1\x15\x7d\x02\x1b\x78\xcc\x98\x04\x99\xcc\x30\x3f\x49\x75\x2b\xf9\xf8\x68\x7e\x10\x9a\xce\xa7\xec\x02\x13\x2d\x4d\x66\x12\xe6\x67\x18\xaa\xd5\x1f\xc6\x3e\x98\xd6\xb2\xa5\xe7\x81\x84\x96\x9b\x1d\x2d\x64\x2b\x07\x00\xfd\xfc\xd4\x38\xcb\xa1\x5a\x87\x0a\x13\x19\xde\xe3\x13\x79\x1e\xef\x08\x53\x3b\x2e\xbb\xe0\xd9\x45\xe7\xb4\x7c\xe9\x41\xb2\x33\x8b\x5d\x82\xcd\x96\x05\x3b\x56\x55\xef\x80\x83\x6f\xcd\x5b\x7d\x87\xfa\x57\xe8\xe2\x02\x45\xaa\x5d\x45\xfb\xe2\x1c\x1d\xec\x23\x36\x8c\x72\x6e\x92\xac\xab\xd7\x9b\x74\xa2\x3b\x24\x3f\x4b\x9b\xab\xe7\x91\x81\xba\x7c\x4a\x7a\x1c\xb6\x87\x64\x15\xd5\x13\x95\x25\xd1\x24\x19\x50\x3a\x03\xda\xb5\x1a\xf3\xae\xf7\xda\x57\x63\x17\x7a\xc1\x76\xb9\xef\x5f\x73\xb3\xc5\x5d\x35\x04\xef\xe4\x8b\xa1\xcb\xc1\xab\x1e\xb5\x6f\x0d\xd2\xd1\x6b\x5e\x7d\x72\xf6\x8f\xbd\x8f\xa1\x7b\x96\xa5\x27\xb6\x93\x0f\xfb\x7b\xec\x3e\x27\xd9\xb1\x22\x6b\x8a\x93\x4c\xdf\x53\x3e\xc1\x86\xde\xd0\x4c\x9b\xf6\x42\x6c\x9e\xe8\x44\x0a\x04\x98\x8e\x66\xa9\x94\x3b\x56\x27\x90\xdb\x95\xaa\x72\x04\x32\x99\x21\x3c\x83\x1b\x21\xbb\x65\x40\xfc\xa8\x27\xe9\x4a\xac\xe9\x3e\x16\xa6\x7d\x36\x99\x68\xf9\x96\x2b\x1c\x5f\x4a\x8a\x36\xa2\x45\xaa\x75\x37\x8f\x84\x6b\xa4\x05\x2a\x29\xa3\xd0\x95\x87\xd3\x01\x41\xeb\x82\x56\xc1\x45\x86\x62\x74\x82\x8c\xb5\x09\xf4\xc8\x5b\xaf\x84\xb7\xb1\x6d\xb1\x71\x0e\x62\xdd\xb3\x2c\x84\x79\x03\xfc\x41\x49\x3b\x22\xa7\xc3\x04\x6f\x4d\x87\xcd\xa1\x6e\x2a\x73\x64\xbf\xe7\xe0\x0c\x8a\x29\xff\xee\x2c\x45\x7f\x92\x1a\x2a\xfa\x7f\xff\x3f\xeb\x82\x5c\x98\x9d\x45\x07\x64\x03\xe8\x7f\x65\x22\xc0\xf9\xf1\x1a\xe8\xff\x45\xaa\xec\x27\xe9\xc0\x64\x99\x0e\xf9\x67\xd3\x65\xf2\xf9\x4d\x12\x66\x18\x1f\x6c\xed\x99\x7b\xbc\xfd\xb5\x86\x01\x05\x48\xb7\xad\x31\x45\xbf\x35\x42\xb0\xdc\xf6\xd4\x14\x05\xef\xdc\xee\x98\x73\x1b\x8e\x6d\x7b\x5e\x47\x61\x26\x0d\x1e\x0e\x69\x29\x46\xf5\x6b\xfb\xc9\x3f\x14\xd0\xbf\x92\x5f\x13\x71\x26\x66\x9f\x34\x3b\x90\x7d\x89\xb9\x9e\x2f\x4b\xd7\x16\x26\x3d\xab\x93\xce\x88\x9f\x47\x00\xe8\x7b\xb8\x41\x0c\x2f\x46\xbf\x0f\x79\x5f\x1a\xfe\xf8\xfb\xda\xca\x18\x47\xe6\x67\x0b\x63\xf4\xf8\xfa\xba\x70\x07\x41\x98\xb8\x30\x0c\x7f\x09\xfe\x9b\xd8\xef\xec\x76\x3b\x46\x34\x66\xfe\x36\x03\xf5\x21\x58\xe6\x93\xba\x8f\x28\x45\x2f\x60\x75\x06\xcb\x1f\xfb\x77\xd1\xa5\xe8\x60\xda\xe0\x8f\x87\xa4\x21\xbb\xe9\xce\xb7\x99\xa8\xed\x67\xdb\x7b\xae\x34\xe1\x85\xfd\x6e\xb1\x0b\x56\x17\x0f\x08\x35\x5d\x35\x8c\x68\xfa\xbb\x84\x3a\x8d\xdb\x06\x72\xae\xdf\x8d\x2b\x2b\x7b\xaf\x57\x2c\x76\x99\x8b\x00\xf2\xad\x20\xb2\x1c\x4a\xd8\x2d\x4f\x9f\x9c\xb9\x87\x65\xfc\x8e\x9c\xa1\x70\x95\xe2\x8f\x94\x03\xf9\x2e\x1a\x7b\x3d\x0f\xdf\x8a\x87\x59\xa0\xf3\x91\x74\xee\x22\x2b\x98\x50\x14\xfb\x66\xbd\xcf\x82\x6e\xc7\x74\x61\x94\xd8\xcb\xcd\x97\xb5\x5a\xd5\x4a\xe1\xb8\x57\x8c\xf7\x19\x1f\x8a\x1d\x59\x17\xb9\x73\x95\x8e\xe3\xbe\x12\x6c\x0d\xcb\x93\x17\x0b\x33\xff\x2c\x04\x37\x41\x0f\xff\x24\x00\x63\x9b\x59\x80\xea\x06\xb6\x7e\xe9\x23\x39\xf2\xaf\x6e\x4a\x45\xf3\x8c\x80\x7c\x57\x2d\xd3\xf3\x66\x8f\x70\xb0\x41\x5e\x93\xc4\x3e\x99\x53\x3b\xea\x8c\xe3\xbf\x03\x00\x00\xff\xff\x0c\x69\x41\x3c\xce\x10\x00\x00")

func admin_http_assetsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsAppJs,
		"admin_http_assets/app.js",
	)
}

func admin_http_assetsAppJs() (*asset, error) {
	bytes, err := admin_http_assetsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/app.js", size: 4302, mode: os.FileMode(420), modTime: time.Unix(1471380932, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _admin_http_assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x3b\x7b\x8f\xdb\x36\xf2\xff\xe7\x53\x30\x42\x81\x24\xf8\x45\xd2\xee\xa6\x45\xdb\x85\xbd\x40\x9a\x26\xbf\xf6\xae\x69\x82\x4d\x7b\x0f\x14\xbd\x82\x96\x68\x8b\x59\x4a\x54\x49\xca\xbb\xbe\xa0\xdf\xfd\x66\x28\xd1\x7a\x58\x96\xe4\xec\x2b\x69\x93\x58\xa4\x86\xc3\x79\xcf\x90\x14\x67\x0f\xbf\x7f\xf3\xe2\x97\x7f\xbf\x7d\x49\x12\x93\x8a\xb3\x07\x33\xfc\x21\x82\x66\xab\xb9\xc7\x32\x8f\x64\x2b\x9f\xe6\xf9\xdc\x8b\xa8\x5a\xc8\xcc\x57\x4c\xd0\x8d\x9f\xad\x3c\x84\x64\x34\x3e\x7b\x40\xc8\x2c\x65\x86\x92\x28\xa1\x4a\x33\x33\xf7\x0a\xb3\xf4\xbf\xf1\xea\x17\x89\x31\xb9\xcf\xfe\x2c\xf8\x7a\xee\xfd\xcb\xff\xf5\xb9\xff\x42\xa6\x39\x35\x7c\x21\x98\x47\x22\x99\x19\x96\xc1\xa8\x1f\x5f\xce\x59\xbc\x62\x8d\x71\x19\x4d\xd9\xdc\x5b\x73\x76\x99\x4b\x65\x1a\xa0\x97\x3c\x36\xc9\x3c\x66\x6b\x1e\x31\xdf\x36\x9e\x12\x9e\x71\xc3\xa9\xf0\x75\x44\x05\x9b\x1f\xef\xa0\x89\x99\x8e\x14\xcf\x0d\x97\x59\x03\xd3\x0e\x18\x2d\x4c\x22\xd5\x0e\x84\xe0\xd9\x05\x01\xd6\xe7\x1e\x8f\x10\x41\xa2\xd8\x72\xee\x05\x41\x08\x7f\x96\x74\x8d\x9d\x01\xfc\x03\xc0\x08\x6d\xb8\x11\xec\xec\x45\x29\xb0\x73\x2b\xb0\x9f\xff\x9f\xd0\x38\xe5\xd9\x2c\x2c\x5f\x5a\xb8\x87\xbe\x4f\x7e\xa2\x86\x69\x03\xf3\xa5\x39\x17\x2c\x26\x34\x8b\x09\xc0\xf1\x25\x87\xc6\x8b\x77\xef\x88\xef\x77\x28\xd0\x66\x23\x98\x4e\x18\x33\x8e\x8e\x30\x4c\xe9\x55\x14\x67\xc1\x42\x4a\xa3\x8d\xa2\x39\x36\x00\x65\xb8\xed\x08\x9f\x05\x27\xc1\x51\x18\x69\x5d\xf7\x05\x30\x4f\x00\x3d\x5e\x4d\xcd\x1b\x2b\x20\x2a\x88\x49\x58\xca\x6e\x71\x6e\xdf\x4e\xd0\xa1\x60\x70\x1e\x30\xc2\x0e\xb1\x3f\xfc\xf2\xfa\xa7\xaf\x88\x4e\x78\x6a\xa5\x76\xce\x74\x2e\xb3\x38\x78\xaf\xc9\x8f\x2f\xbf\x21\xba\xc8\xd1\x6c\x88\x5c\x56\x80\x4c\xc0\x8c\x99\xd1\xa5\x88\x59\xcc\x29\xf9\xb3\x60\x8a\x33\xed\xf8\x04\xa4\xbf\xf1\x25\x11\x06\x10\x90\x6f\x7f\xb7\x7d\xa5\xd5\x10\xad\xa2\xb9\x87\x86\xac\x4f\xc3\x50\x6a\x1d\x54\x5c\x23\xa3\xe8\x30\x5f\x01\x19\x6b\x60\xf4\xeb\xe0\xa4\x6e\x5b\xf6\xde\x03\xc9\xb3\xb0\x44\x33\x15\xa3\x2a\x59\x09\x8f\x83\x2f\x01\x5f\xd5\xea\xc7\xf6\xf0\x37\x96\xc5\x7c\xf9\x3b\xb2\x30\x0b\x4b\x8f\x7c\x30\x5b\xc8\x78\x43\x94\x14\x68\xf8\x32\x2a\x90\x6f\x8f\xd4\x92\x7b\xc5\xaf\xc0\xba\x32\xba\x5e\x50\xe5\x98\x8f\xf9\x9a\x44\x82\x6a\x3d\xf7\xaa\x17\xe5\x8f\xcf\xb3\x35\x03\xc7\xf6\x2a\x7c\xd0\xcb\x57\xd4\xfa\x11\x8e\x6b\x8f\x44\xb7\xa1\x3c\x63\xaa\x7a\xd7\x87\xd7\x47\x22\x1b\x10\x00\x43\x3b\x10\x0b\x05\x3a\xda\x6a\xde\xeb\xba\xd2\x2c\xa4\x5b\xf4\x21\xe0\x1f\x98\x2b\x92\x42\xd0\x5c\x33\xe2\x1e\x9a\xd3\x16\xa2\x01\xed\xd8\x85\x9f\x06\x8c\xb5\x4a\x07\x45\x23\xc3\xd7\xac\xf5\xd6\x12\xbf\xa5\xf3\x07\x99\xb2\x06\x71\x25\x81\x82\x77\xd0\xed\x1b\xbf\xa0\xf1\x6b\x66\x14\x8f\x74\x78\xf2\x65\x02\xaa\x46\x11\x7f\x47\xd1\x58\x6d\xef\x20\xe6\x59\x58\x88\x7e\xa1\x80\xc2\xc3\x00\xb8\xaa\x65\x01\x1a\xaf\x61\xaa\x87\x07\xfb\x14\x59\xa9\x3d\x85\x96\x4d\x06\xf8\x06\xba\x04\x53\x73\xef\x35\x74\xbe\x30\xa2\x61\x08\x7d\xaa\x50\xf2\xb2\x1c\x29\x24\xbd\x68\x6a\x1d\x70\x18\x7c\xa1\x58\xce\x28\xc4\xda\xb2\x83\x67\xc4\x3e\x80\xa5\x7f\xf8\x60\x9f\x82\x54\xaf\xfe\xfa\x0b\xd8\xc7\x46\x03\x41\x8b\x5e\xe1\xa7\xb1\x7f\x7c\xd2\xd6\x5d\x72\x72\xf6\x0f\x2a\x78\x6c\xed\x15\xdc\xe3\xa4\x23\x7b\x43\x21\x09\x39\x1c\x65\xc3\xfe\x8b\x5c\xc6\x2c\xd3\x2c\xee\x68\x1b\xc7\xb8\xb4\xd7\xed\x57\xbb\x9d\x16\xbc\x45\x02\x34\x87\xa1\x18\x91\x0a\xdc\xa3\x1f\x12\x7a\x77\x66\x41\xc8\x1e\x8a\x66\x06\x83\xc0\x01\x74\xc6\x20\x6e\x60\x7b\xc9\x57\xc1\x4f\x6c\x45\xa3\xcd\x1f\xa5\xdd\xfd\xb1\xde\x92\x8f\x4a\x30\x3d\xbc\x77\x86\x3b\x4e\xfe\xb0\x9c\xec\x1b\xb4\x87\x97\x2e\xd5\xd0\x85\x1a\x69\x5a\x7a\xc3\xce\xa6\x58\x41\x69\x07\xdf\x09\x1a\x5d\x08\xae\xcd\xbd\x99\xc1\x6b\x6a\xa2\x84\xe4\xe0\xec\xfc\x6a\xc8\x10\x4a\x38\x5d\x2c\x20\x53\xf2\x6c\x35\x0e\xaa\xd8\x8a\x0d\x62\x7c\x1e\xa1\xf6\xf4\x0d\xd9\x54\xd3\x65\x17\xe8\xae\x56\x56\xc1\xc2\x09\x78\x47\x56\x7b\xa5\xe2\xec\x66\x11\x94\x52\xd9\x6f\x5f\x35\xe4\x56\x2e\x53\x80\xad\x64\x46\x00\xb7\x7a\x67\x57\xc6\x8f\x20\x53\x62\x6e\x82\xa8\x6c\x23\x16\x8f\x2e\x20\x80\xb1\x54\xae\xd9\xd6\x80\x1e\x7f\xc1\xc1\x2a\xae\x9e\x00\xd4\x36\x35\xac\xc4\x26\x4f\xb0\x0c\x24\xdb\x27\xbf\x1c\xe6\x47\x5c\x45\x50\xe9\x86\x67\x18\xbf\xaf\xe5\x0a\x3d\xce\x60\x3b\xc1\x9e\xcf\xd9\xa5\xe2\x40\xb9\xee\xb1\xee\xa5\x54\x69\x55\xde\xaa\x0a\xec\x15\x74\x79\x8e\x76\x7c\xef\xd3\x38\x26\xf6\x61\xa5\x64\x91\x93\x84\x6a\x7f\xc9\x58\xbc\x00\x9e\x5d\x0a\x58\xda\x41\x20\x16\x50\x41\xca\x31\x5e\xc7\xb1\x9b\xf7\xf1\x13\x78\x23\xab\x48\xd1\xa5\xef\x6e\xdc\xeb\x8d\x88\x87\x7c\xe0\x67\x76\x39\xec\x49\xf7\xe4\x41\xaa\xf6\x20\xa7\x1d\x7d\xb8\x07\xa9\x40\x8a\x78\xdc\x23\x54\x90\xb1\xcb\x29\x60\x50\x8c\xde\x90\xdb\x6c\x2d\xe4\xfe\xbc\xa6\x16\xfc\x61\x72\x6d\x39\x48\x8f\x5f\xf4\x0f\xb4\x83\x79\x96\x17\xb6\xb6\x49\x65\x8c\x0b\x1a\x10\xbb\x13\x44\x00\x86\xea\x55\xfe\x28\xf1\xb1\x39\x4b\x55\x5d\x79\x24\x87\x68\xc3\x12\x78\x8f\x75\x96\x05\x53\xb8\x84\x56\x6c\x8f\x46\xec\xac\x98\x0a\xd1\x3f\x13\x79\xd9\xf6\x75\xb4\x8e\x00\x14\x60\x1d\x74\x80\x6e\x8b\x45\xe7\x34\x1b\x42\xc3\x94\x92\x2a\x80\x15\x3c\xf4\x42\x85\xfa\xf2\x2a\x67\x91\xc1\xf5\x04\x68\x8f\xa5\xb9\xd9\x10\x97\xb8\x10\xd3\x00\xb9\xed\x44\xde\x7e\x35\xc1\xf2\x6e\x58\x2d\x10\x20\x9c\x5a\x32\x7c\x1c\x57\x0b\x82\xdd\x07\xfd\x66\x93\x23\x91\x45\xba\xc0\xfa\xbc\x9f\x1b\x88\x67\x8e\x9b\x14\x1f\x7b\xb9\xb9\x96\x49\x01\xda\x9b\x30\x29\x8b\x66\x9f\x49\x41\x38\xb1\x5c\x12\x23\x09\x64\x0a\x09\xd5\x31\xc0\x03\xdd\x56\x0f\xe5\x72\xfe\x31\x25\xb9\xd4\x1c\x97\x66\x15\xf4\x53\x28\xa1\x89\x7f\x8c\xf9\x0c\x8c\x92\x08\x0e\xd9\xea\xc9\x6d\x59\xe3\x9e\x17\x8b\xc2\x18\x88\x6b\x95\xd8\x17\x26\x23\xf0\xd7\xd7\xa9\xfd\xc9\x15\x4f\xa9\xda\xd8\xe7\x85\x90\x98\x62\x4b\x9d\x96\x99\xd5\xea\x34\xe6\x1a\x93\x42\xdc\x11\x57\x2d\xf1\xe7\x31\xe4\xbb\x72\x9a\x43\xc8\xee\x8b\x9a\x07\x55\x1b\x21\xda\xd0\x6e\x05\xf2\x7c\xb5\x82\x6a\x8b\x1a\x39\x56\x83\xd0\xd5\xea\xc6\xca\x8f\x7a\xd2\x4f\xa0\x00\x79\x55\x64\xd1\xd8\x22\xef\x7c\xac\x54\x7f\x53\x18\x74\x71\x64\x96\x9a\x21\xc0\x1f\x31\xe3\x02\xbf\xe4\xb1\x66\xd1\x93\x21\xc8\x7f\x52\x6e\xc6\xa1\x6e\xaf\xc2\xa1\x75\x85\x43\x6b\x23\x39\xbc\xc6\xa1\xc1\xb2\x18\x58\x82\xd6\x60\x53\xaa\x7e\x0b\x08\xc2\x7e\x95\x9a\x29\x90\x4e\xda\x53\x60\x51\xde\x37\x54\x3c\x35\xec\xfb\xda\xe5\xd3\x75\x22\x41\x49\xf6\x7d\x57\x50\x20\x8d\x00\x9c\xcc\xe5\xb5\x25\x3e\x8e\x67\x69\xe7\x96\x75\xba\x43\xac\x55\xae\x99\x7b\x36\x64\x00\x66\x04\x1b\xa2\xa7\x95\x08\xab\x18\x86\x06\xf9\x91\x39\xb0\x85\x61\x6f\xfa\xab\xd4\x0f\xd4\x93\x65\xc5\xc6\x29\xd1\x45\x8a\x09\x8e\xae\x3f\x93\x0a\x0b\xd5\x66\xe3\x9e\xb7\x5d\x85\xda\xc6\xb8\xea\x2a\xc0\x1d\x75\x95\xd8\x0e\x55\x96\xc5\x76\x4d\x75\x55\x38\xf6\x29\xcc\x62\xc6\xfd\x98\x42\x50\x45\xd8\x55\xae\x98\xd6\x36\x25\x7c\x2e\x8a\x2a\x43\xe2\x76\x7d\x52\x98\x25\xb6\x26\x2c\x51\x0a\x33\x56\x4f\xde\x5b\x51\x8c\x7c\xb9\x00\xee\x38\xe3\xdb\x76\x2f\x6f\x4d\x93\x0b\xff\xf3\xdb\x91\xff\xed\xef\xff\xf7\x45\x78\xb0\xc9\xb9\x59\xae\x69\x75\x35\x9a\xd1\x42\xf9\x31\xe4\x5a\x48\xf5\x50\xd7\xe8\x27\x8d\xaa\xf9\xcf\x82\x66\x86\xff\x17\x56\x65\xc4\x21\xfb\x64\x6d\x72\x44\x93\x98\x5e\x9d\x16\x2f\xed\xf3\x6d\x6a\x10\x67\xb8\xa6\xf6\x4a\x14\x1f\xab\x39\x5c\xef\x20\x86\xcf\x77\xfd\xe2\xe4\xf0\xc9\x2e\x5d\xce\x21\x78\xd9\xfd\x8a\xc1\xad\x53\x00\x62\x37\xb7\x6f\x8a\xd8\xee\x79\xcd\x52\xd6\xfd\xb6\xe0\x2f\x8f\x85\x2f\xd9\x23\x21\x08\xba\x20\x96\x92\x9a\x24\x4c\xb9\x6f\x00\xfa\x46\x5a\x1e\xc8\x05\xdb\x0c\x2e\x78\x2c\x10\x5a\xc7\xf8\x39\xc6\xbd\x1c\x8d\xc4\x31\xe6\xe8\x21\x90\x77\xb9\x94\x62\x08\xe0\x2d\x54\xec\x62\x88\x3f\x3c\x7d\x46\xef\x9d\x9f\x90\x3b\xda\x49\x46\xa9\x1f\xb0\xc4\xc2\x38\x50\x46\xaf\x81\xa5\x05\xe4\xfb\x8d\x47\xa8\xe2\xd4\x4f\x78\x0c\x36\x08\x76\xa9\x0a\x66\x3f\x48\xc0\xd0\x34\x70\x36\xe8\xd0\xf2\x6c\x29\x3d\xbb\xc1\x0c\x66\x33\x78\x9a\xb8\x3b\x02\x6d\xe8\xc0\x21\x29\xda\x00\x53\xa3\xa7\x4b\x83\x83\x27\x1c\x38\x0d\x8e\x1f\x59\x8d\x76\xc7\x6e\x6d\xc5\x7b\xe6\x4d\x16\x69\x3d\xe8\xa4\x7f\x13\xde\x86\x1b\x2b\xf5\x5b\xde\x80\xb7\x26\xd6\xb4\xca\x18\xad\x52\x05\x31\xd3\x10\x63\x9b\xdf\x8d\xf4\xf0\x74\x86\x61\x68\xd4\x0e\x41\xaa\x6b\x85\x84\xf2\x55\x62\x86\x0c\x12\x42\xd7\x70\xea\xbb\xc6\xdb\x52\xc4\x96\xc6\x0f\x8f\x62\x9a\xad\x98\x7a\x74\x4a\x1e\xc6\x81\xcc\x04\xcf\xd8\x53\xf2\x08\x15\x03\x5d\xae\xe7\x2f\xb4\x8a\x78\xb2\x49\xde\xc8\x24\x53\xcf\x4a\xaf\x3b\xcf\xa4\x63\xd6\x8f\x9c\x83\x96\x01\xfa\xe6\xb1\xf7\xee\xc1\xf4\xe3\xb7\x06\xb8\xad\xeb\xe2\x40\x63\x46\xf0\x86\x6c\x34\x89\x63\x70\x9d\x3e\xcf\x1d\xf0\xe8\x06\x0f\xc4\x31\x41\x6e\x8f\x8b\xdc\xa6\xad\x41\x36\xe4\x85\xaf\xf9\x2a\xbb\x2b\x56\x86\x83\xd8\xf7\x75\x08\x29\x43\xd9\xd3\xfb\x3c\x52\xfc\x88\xef\x6b\xee\x22\xc3\x1e\xb6\xea\xea\x3d\x10\xc3\x5c\x11\xfc\x9d\x6d\xdc\x22\xeb\x17\x48\xbc\x23\x07\x48\x9d\x6d\x1b\x57\x18\xf6\xce\x7a\x87\x4c\x94\x94\x8f\x72\xd1\x22\x5e\xb3\x2c\x7e\x2e\x84\x2d\x1f\x87\xb6\x0c\xed\x0c\x16\xe9\x1e\xaa\xda\x27\x67\x6e\xf5\x60\x69\x1a\x5f\x50\x76\x4f\xcc\x3a\xc3\xf7\x2d\x26\x9b\xc4\x3f\xb5\xad\x57\x5c\x69\x53\xb5\xa5\xc2\xef\x9a\x35\xd7\xf8\x69\xf3\x0f\x54\x27\x23\x87\xb4\x7b\x97\x94\x77\xa8\xc2\xb7\x36\x4f\x3a\x25\xba\xd6\xb8\x1a\x2b\xc8\x7b\xa6\xfe\x9d\x4b\xc0\x8e\x81\x46\xc7\x38\x0f\x35\xf0\x3d\xb3\xd1\xda\xbd\x3d\x9f\xba\x7b\x7b\x7e\xf0\xee\xed\x1d\xb2\x54\xad\xfb\x1c\x53\xdb\xe6\x01\x61\xae\x46\xb1\x73\x9c\x50\xbd\xb9\x4b\x1e\xcb\x0d\x18\x28\xc7\xa2\x8b\x85\xbc\xea\x7e\x16\x50\x5a\x63\x59\xb6\x54\x96\xd8\xaa\x61\x5a\xfc\x7e\x62\x64\xbf\xad\xea\x94\x2a\x04\xb4\xab\x96\x89\x84\xef\xc9\xc8\xb7\xbe\xf7\x55\x07\xee\xc9\xbb\x5f\x37\xff\x89\x60\x77\xdf\xab\xfd\x7d\x78\xdd\x68\x7c\x0a\xbe\xfd\x3a\x1c\xf7\x85\xc2\xed\xd7\xe0\x76\x37\xc8\x75\x7f\xe7\x2e\x73\x40\x56\x51\x8c\xfc\x8d\xae\xe9\x3b\x7b\x31\xc1\x22\x9b\x1f\xfc\x5f\xe3\x1a\x06\x79\x8b\x8e\x16\x13\x6a\xf0\x26\x0a\x81\x3c\x86\xf7\x38\xf0\xd1\x5d\x67\x20\x5a\xda\x76\x4e\x57\x4c\x13\x21\x69\x4c\x96\x14\x12\xdb\xf6\x3e\x43\xdf\x35\x0b\xfa\x9e\x5e\x05\x2b\x29\x57\x82\xd1\x9c\x6b\x7b\xd7\x02\xfb\x42\xc1\x17\x3a\x7c\x8f\xd7\x41\x36\xe1\x71\x70\x0c\x7f\xaa\xd6\xf8\x15\x8e\xe9\x17\x60\xde\x77\xef\xde\xb4\xf1\xee\x23\x3a\x02\x77\x08\xa0\x82\xc6\x23\xa5\xf7\x3a\x90\x6a\x05\x24\x9e\x04\xc7\x47\x61\xd5\x39\xfd\x9a\xc9\x28\x2a\x28\x95\xb5\x2c\x54\xc4\xa6\xf0\x0d\x7c\x02\x92\x48\xc8\x22\x5e\xc2\x58\xd6\x11\xa7\x43\x59\x70\xbf\x96\xc4\x11\x0a\xf7\x28\x6c\xf6\xf9\x26\x17\x7a\x44\x16\xf6\xf6\xcf\x3e\x72\x4a\xf7\xc3\x05\x50\x08\xbe\x67\x58\x0a\x51\xda\x40\x84\xe0\xe0\x7f\x45\x8e\x5b\xab\x36\x8e\xbc\x96\x31\x15\x01\x5e\xce\xe9\xb9\xb8\x92\xe2\xcb\xee\xcd\x94\x59\xf2\xac\xfd\xde\xde\xdd\xf2\xce\x7e\xb5\x48\x89\xf5\xed\x53\xf2\xe1\x83\x7d\x70\x3b\x5a\xc9\xb3\x96\x33\xd9\xc7\xc6\x36\xf2\xf2\x1a\x3b\xc8\x3b\x5b\xc5\xbb\x1c\x60\x3c\x68\x5e\x71\x69\x00\xd4\x73\x74\xee\xb7\xd0\x05\x13\x48\xc1\xdc\xc3\x2c\xe6\x9d\xbd\x2d\x73\xd9\x2c\xb4\x6f\x5a\xb0\xdd\x74\x5a\x72\x8e\x03\x5c\x78\xb6\x28\xa6\xd4\x67\x55\xed\x3a\xa5\x3c\x68\x97\xd2\x88\xd4\x96\xbe\xfb\xaa\xe8\x4e\xed\xdc\x80\x1f\x3e\xb6\x1d\x3a\xae\xed\x5e\x3a\xd8\x7f\x07\x61\x82\x98\x71\x4b\xc3\xab\x77\x9e\xa7\x8a\x19\x07\x38\x31\x5b\x14\x13\xc4\x7c\x7c\xf2\x75\x70\x04\xff\x1f\x9f\x9e\x1c\x1d\x7d\x39\xf8\x09\x44\x4f\xcd\xd2\x23\x78\x9c\xf8\x10\xc1\x97\xf0\x7b\x04\x7f\x4a\x1e\x25\x52\x9b\x53\xbc\xa8\xf7\xe8\x30\xa1\xef\xbd\xf8\x55\xfa\xc1\x12\xa2\x4b\xfb\x8e\x59\x4f\x52\xc7\x14\x1e\xb3\x25\x2d\x84\xf1\x1a\xbb\x0d\x11\xcd\x22\x26\x1e\x3f\xf1\xce\x5e\x08\xa9\xd9\x6e\xae\xde\x53\x20\x54\x95\x41\xa7\x02\x58\xb6\x92\x7f\x63\x1a\x79\x81\x53\x94\xb1\xa4\x3b\x47\x2b\x29\xbb\xe4\x5d\xc7\x3e\x00\xb7\x79\x7f\x16\x96\x77\x76\xff\x17\x00\x00\xff\xff\x4f\xc5\x38\xa5\xc4\x3b\x00\x00")

func admin_http_assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsIndexHtml,
		"admin_http_assets/index.html",
	)
}

func admin_http_assetsIndexHtml() (*asset, error) {
	bytes, err := admin_http_assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/index.html", size: 15300, mode: os.FileMode(420), modTime: time.Unix(1471423144, 0)}
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
	"admin_http_assets/app.css": admin_http_assetsAppCss,
	"admin_http_assets/app.js": admin_http_assetsAppJs,
	"admin_http_assets/index.html": admin_http_assetsIndexHtml,
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
	"admin_http_assets": &bintree{nil, map[string]*bintree{
		"app.css": &bintree{admin_http_assetsAppCss, map[string]*bintree{}},
		"app.js": &bintree{admin_http_assetsAppJs, map[string]*bintree{}},
		"index.html": &bintree{admin_http_assetsIndexHtml, map[string]*bintree{}},
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

