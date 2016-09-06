// Code generated by go-bindata.
// sources:
// internal/creds/ca_certificates.pem
// DO NOT EDIT!

package creds

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

var _ca_certificatesPem = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x79\x49\xd3\xab\xb0\x92\xe5\x9e\x5f\x71\xf7\x8e\x6a\x63\xc0\x36\x54\x44\x2d\x24\x66\x8c\xc0\xcc\xe0\x1d\x93\xc1\x0c\x06\x03\x66\xfa\xf5\x1d\xf6\x37\xf4\x9d\x5e\xbd\x57\xd5\xd1\x1d\x75\x97\xa9\xbc\x48\xd6\x97\x79\xce\xc9\x23\xb6\x79\xde\x87\x6e\xf9\xf7\x1f\x8e\x85\xe9\x5d\x16\xde\x6f\x6b\x38\xdc\x9a\xfb\xbf\xff\xe0\x1f\xcf\xdb\x35\x9c\x31\xb6\xa9\xeb\xe6\xfe\x43\x0b\xeb\xf4\x3b\xf8\xc3\x4a\xe3\x67\x97\xfe\x60\x01\xf6\x6f\xaf\x7f\x90\x17\x65\xed\x07\xcb\x9b\xb6\x2c\xc8\x2c\xb0\xf9\x77\x14\x43\xb2\xcc\xc9\x1c\xcb\x82\xa6\xce\xc0\x24\x43\x90\xc9\xbc\x96\x1c\xf1\x75\x05\x1a\xcc\xca\x47\x5e\xde\x44\x66\xc2\x21\x30\x1c\x01\x70\x50\x47\x46\x3f\xb1\x46\xc0\xb9\x86\x21\xf2\x93\xe2\x62\xce\xca\x1b\x08\x50\x22\xd8\x39\x3c\x9b\x23\xc9\xf4\x85\x5d\xe8\x5d\xf2\x94\xe3\x07\xc4\xf6\x1f\xf1\x19\x95\xdf\x71\x16\xda\x17\x4f\xdb\xc5\xb5\x93\x19\x84\xbb\x60\x89\x58\xd5\xa1\xa7\xe5\x89\xe8\x64\x86\xef\xe2\xa1\xc8\x2c\xa1\x6f\xee\x11\xa4\x7c\xce\x2e\x29\xc4\x65\x0b\x2a\x78\x42\xe3\xf8\x1d\x72\x1b\x9f\xb3\xf9\x5f\x62\x18\x72\x9b\xc9\x2e\x78\x15\x81\xf2\xbd\x1b\xcc\x11\xeb\xba\x68\xe6\x39\xa0\xc3\x4c\x73\x21\x68\x6c\x88\xbb\x73\xe2\x55\x75\xe0\x67\xb3\x6a\x83\xee\x23\xde\xdb\x0a\xef\xce\xd8\xe7\x42\xe6\x10\x6e\x91\xf8\x4a\x25\xf3\x5a\x15\xdf\xcd\xf6\x52\x57\x45\xe0\x9b\x95\xcc\x0b\xbb\x44\xcc\xc7\xb8\xae\xf0\xd4\x66\xe1\xfd\xb7\xbb\xc1\x80\xc1\x0b\x00\xe8\xb0\x30\xc0\x94\x05\x25\x9b\x05\x3c\x98\x5c\x62\xf6\xc4\x58\x0e\x9e\x87\xac\xbe\xe1\x57\x56\x24\x4c\x41\xbc\x05\x6c\x7e\xdc\x10\x99\x39\xf2\x94\x79\x93\xe3\xb3\x79\x45\x87\x2b\x06\x53\x96\x02\x57\xa8\x6b\xfa\x7a\xbb\xb5\x0e\x7f\x39\xad\x73\xb8\xd3\xae\x30\x3a\xab\x17\x8a\xdd\x1a\xd9\x49\xdf\x0e\x38\x64\xd3\x35\x07\xd0\x3c\x6f\xcf\xe3\xc4\x69\x3b\xee\x59\xf5\x1d\x65\x6e\x00\x16\x2b\xa5\xeb\xee\x91\x47\x1b\x1b\x3f\xec\xae\x6c\xc8\xa2\x78\xe5\x77\x17\x74\x9a\x4d\xa9\x78\x8e\x27\x26\x7a\x06\xb8\x7b\x9c\x93\xea\xea\x68\x6a\xe1\x00\xfa\x70\xd3\xd3\xad\x70\x26\xb3\xf9\x58\xb2\x18\x98\x78\x00\x42\x9d\x05\x46\x39\x65\x19\x2f\x20\x09\xbc\x2e\x33\x91\x26\xb3\x45\x62\x3c\x5d\x42\x58\x34\xa2\x58\xfa\x2b\x4c\x7e\xaf\x81\x57\x09\x60\xff\x37\x35\xf0\x2a\x01\xec\xaf\x35\x60\xe0\x13\x3b\xbd\x77\xe2\xf8\xc9\xe4\x9c\x72\x9a\x11\x6c\xde\x27\xe3\x81\x61\x23\x28\x42\x6e\x95\x27\x64\x67\x13\xa6\x73\xf2\x82\xec\x00\x47\xb6\x33\x7b\x05\x50\x5f\x7f\x63\xc9\xa0\x79\x08\x10\x0b\x8c\x60\x92\x5e\x1f\x32\xf1\x02\xc2\x6c\x12\x1a\xe0\x58\xfa\xa5\xd9\x58\xa3\xd5\xb7\xbe\x6f\x32\x59\x81\xc9\x10\x9e\xd1\xfe\x66\xdc\x19\x63\x92\x8c\x77\xb2\x0e\x61\xc0\x0b\x7c\x51\x87\xe7\xb2\xc3\xbb\x93\xbb\xc3\xaf\x81\xbc\x00\xc3\x5e\xf5\xa0\x54\xb6\x0e\x02\xd3\xc7\x69\x26\x43\xc0\x10\x40\x10\x5c\xe9\x49\xcc\x02\xe5\xd4\x5c\xe4\x7c\xbc\x30\x10\xe7\x01\x04\xf8\xc4\xce\xbd\xe0\x16\xe8\x89\x44\xc4\x81\x2c\x02\x08\xe0\x22\x6b\x3d\x44\x4b\x8e\x48\xce\xe0\x21\x34\x1c\x80\x01\x4a\x84\x40\x28\xf4\x53\xda\xd1\x0c\x73\xd8\xad\xd9\x69\x2f\x1c\x3d\x01\x8f\xee\x05\xa5\xf8\x48\xb1\x79\x0d\x9c\x42\x2b\xba\x6f\x88\xb2\xd6\x53\x47\xf1\xcd\xba\xde\x96\x7c\xb2\x2f\x72\xef\x10\x60\xc7\x47\xb1\xf5\xfa\xc2\x76\x23\xa5\x8e\xdd\x6b\x3a\xb1\x52\x77\xb6\x1e\x77\x19\x2f\x21\x94\x2f\x6c\xba\x5d\x9f\xd7\x83\xec\x39\x9d\x7b\xbf\x30\x1a\x20\xd6\xbe\xf6\x54\xb9\x49\x56\xe2\x29\x48\x49\x8e\xed\xc6\xe6\x71\xb9\xa5\x19\x77\x7d\xdc\xe3\xdd\xfa\x88\xcf\xa2\x23\x7b\x2e\xef\x6f\x3b\xfa\xb8\x54\xcd\x23\x3c\x49\x69\xca\xec\x8f\xf8\xa6\x87\x64\x4c\x7d\x80\x0b\xaf\x71\x7f\x42\x0b\xf6\x8f\x81\x4b\x4c\x1b\xbb\x7b\xf6\xc3\x0f\xf9\x1e\xff\xaf\x5f\xe1\xeb\x7b\x49\xac\x9a\x28\xac\xfe\x05\xfc\xe2\x3d\x9b\x65\x01\x2e\x7e\xe1\x17\x07\x8a\xb6\xf8\xdb\xf5\x22\x5e\x9e\xd9\x15\x28\x1f\xcd\x1f\xd8\xa0\x72\x6d\x0c\x99\xc1\x24\x80\x77\x89\x9d\xf8\x69\x27\x5d\x3c\xc6\x89\xef\xee\x9a\xb0\x50\x89\x6a\xf4\x44\x66\x3f\x89\xc6\x67\x09\xce\xca\xcf\xeb\x52\x24\x32\x37\x2c\xf0\xa6\xcc\xc0\xf9\x49\xca\x63\x0d\xd9\xf2\xa4\x73\x32\x89\x0a\x80\x23\xce\x98\xbc\x57\xac\x90\x27\xcd\x96\xa7\xef\x58\x01\xf9\x9f\x3b\x08\xfb\x6c\x21\x0f\x41\xe3\xab\x83\x34\x93\x70\x47\x57\x52\x76\x31\x69\x64\x96\xb7\x2f\xd4\x82\x4f\x10\xfc\xe8\x24\x30\x23\xe7\xe7\x75\xcc\xd9\x69\x48\xe6\x35\x28\xb3\x78\x66\xae\xf2\xab\x7f\x6f\xbf\xff\x7e\x60\x38\x00\x50\x32\xe4\x26\xf0\x5a\x3f\x81\x46\x86\xc0\x60\xf7\xca\xb6\x3a\x63\xc4\x39\x24\x05\x7b\x73\x5e\xe3\xa3\x57\x98\x73\xb7\xf5\xb7\xa1\xcb\x0a\x3a\x60\x8a\x13\x2e\x29\x96\x10\x15\x4a\x56\x0d\x41\x1a\xd8\x5b\x45\x52\x46\xba\xae\xb6\xa3\x12\x5d\xea\x7b\xc7\x9d\x1f\xf7\x33\xe6\xb0\xb2\xcd\x35\xc1\x85\xd8\xe4\xca\x91\x1a\xeb\x5d\x79\x5d\xb6\xbe\x26\xb0\xa7\xc3\xb9\x13\x0e\xc4\x46\xd9\xd3\xcc\xcc\x6d\xcb\x52\xab\x8f\xf3\xea\x1c\x1f\x82\x78\x83\xa2\x62\xf9\xd5\x41\x89\xf7\x2a\x16\x8e\x9c\x2f\x49\x41\x78\xb2\x57\x65\x7f\xc6\xd3\x3c\x59\x33\xe4\x39\x82\x39\x87\x3d\x9b\x25\xaa\x7a\x87\x69\x38\x85\xf7\x70\x6d\xfb\x6e\x72\xac\x59\x95\x4d\x25\x09\x36\xd5\x72\x9f\x32\x62\xf6\x31\x49\xbb\x55\x47\x7a\xed\xb7\x49\x60\xd1\xf6\x36\x32\x54\xeb\xc9\xcd\x85\x3d\x87\x0c\x28\x2b\x5c\xf2\xcb\xe3\x26\xc8\x63\xf2\x96\xd6\x6a\xc2\x86\xb7\x63\x94\x9d\xf6\xc4\xe4\x7a\xab\x67\xc4\x03\x19\x60\xb6\x25\x39\x06\xab\xc5\xc5\xe6\x00\x90\x19\x76\xa1\xe0\xe3\x5c\x31\x38\x07\xc3\xd4\x68\x34\xeb\x99\x7b\x6c\xa3\x9d\xd2\xda\xdd\xa1\x16\xea\x1d\xcb\xb8\xd2\x96\x02\x83\x77\x56\xf2\x73\x4c\xd1\x3a\x16\xcd\x4d\x41\xc7\x77\x99\x48\x36\xf4\x51\x50\x7d\x90\x21\x08\x80\x58\x64\x99\xe0\x20\x59\x86\x0e\x07\xae\x6f\x48\xb2\x10\x2f\x72\xc0\xcb\xa0\x0d\xd2\x36\x6f\x8a\xe0\x7e\x7c\x4c\x6e\x89\x71\x50\x60\x1e\xf7\x5d\xf5\x44\x1d\xb2\x0b\x90\x7c\xe0\x17\xc5\x0b\x99\xe1\xf0\x4e\x67\xc4\xeb\xdd\xf5\x88\x52\xf6\xd5\x86\x19\x09\x2b\x7c\xc8\x7d\xec\x46\xd3\xc4\x67\x6f\xa8\xb2\x31\x60\x48\x5b\x08\xb2\x09\x66\x3c\xdc\xbe\x7a\xe1\x93\x00\x25\x83\x7e\xc1\xd2\x37\x10\xea\x1f\xff\xe1\x0a\x39\x34\x21\x1b\x8c\x0d\x4b\x64\x27\x2a\x68\xb1\x50\x32\xf1\x98\x6b\x46\x95\xd0\x96\x88\xdd\xdf\x2f\x1e\x83\x7f\xd4\xfa\xbe\x88\x08\x7c\x0c\x48\xa5\x8f\x17\xe6\x9e\x88\x49\x1f\x11\x4a\x1e\xb1\xfb\x22\xae\xa7\x49\x03\x81\x7c\x9a\x02\x88\x41\xc3\x91\x80\xc1\xf3\x27\x0e\xd4\x88\x35\x44\xf6\x05\xfe\x86\x23\xc0\x15\xc0\x3c\xcf\x9b\x44\x32\x27\xfd\x46\x8f\x11\xa1\xad\xf1\x9f\x1b\x4c\x98\x0d\xde\x47\xcb\x20\xef\x4c\xc6\x0a\x21\xcc\xda\x2c\x2f\x33\x98\x8f\xb9\x00\x8c\x98\x40\xaf\x23\x1b\x1f\xbb\xbd\x37\xcb\x78\x4f\xf1\x72\x3c\x91\xc0\x41\x5d\x18\x12\x4b\xc8\xf8\x79\x21\xdc\x31\xf9\x68\x8e\x67\x40\x30\x83\x4a\x2a\x55\x4c\x30\xbb\xb8\xd6\xaa\x78\x61\x8a\x58\x42\xd3\xe9\xe3\x0e\x4c\xc8\xa2\x49\x0e\x8d\x2b\x82\xf8\x2c\xae\xe0\x82\x7d\x80\x03\xb2\xf9\xea\x52\xc5\x75\x65\x87\x5e\xf2\xb4\x5d\xa8\x5a\x16\xbe\xa8\xb6\xbc\xd3\xb8\xbf\x08\x23\xf0\x22\x4e\x1e\x9c\x75\x87\xc1\xbc\xfc\x79\xd3\x96\xae\x30\x7b\x9a\xa8\xf2\x8c\x4e\xb7\x22\x3f\xbb\xa9\x98\xe0\x6c\xa2\x5d\x81\x45\x4b\x59\xb0\x59\x4e\x25\xd9\xe6\x6a\x2a\x4b\xb5\x1d\x44\x45\x69\x30\x2c\x75\xbc\xc7\x8d\x16\x6d\xb1\xc7\xf5\x36\xa7\x97\xf4\x82\xc7\x5a\x6f\x3c\x3c\xab\xd2\x61\xc2\x71\x28\x50\xe2\xb2\xab\xdc\x73\xb6\x0f\x4b\x74\x75\xae\x9b\xeb\x8e\x9f\x4d\xe1\x48\x9e\x72\x6a\xd7\xb4\xc6\xc2\xd0\xf7\xe7\xa4\x46\x22\x56\x3f\xd6\xb4\xb6\x84\x87\x7c\x00\xd4\x45\x3f\x14\xb3\x9c\xaf\xa8\x18\x8c\xb5\xdb\x0c\x38\xe9\xa4\xed\x68\xb7\x1b\x47\xe9\x02\x55\x4d\xcc\xeb\x2e\x71\xa7\x42\x57\x5d\xae\xe6\x0b\xd9\xe3\xa9\x6e\xc1\xaa\xd3\x29\x32\x0f\x37\xf1\xca\xd4\xc1\xfe\x7a\xad\x92\xbb\x59\x12\x8a\x48\xe7\x30\x68\x08\xd6\xe5\x91\x74\x60\x89\xd3\x32\xef\x11\x57\x0a\xde\x1a\x0d\x37\xe3\x0e\x6e\x0d\xe4\x1b\xef\x80\x82\x3c\xc0\x4c\xd2\x2e\x9e\x9a\x52\xb6\x3d\x5a\x3c\x8b\x6a\x2d\xdc\x9f\x3d\xaa\x52\xd4\xe6\x14\xce\x79\xe6\x9a\x1a\x78\x5e\x00\xef\x86\x5c\xb1\x67\xc6\xaa\x06\xd3\xec\x8a\x7b\x62\x0b\x78\xf6\x49\xf3\xd9\x1d\xb3\x75\x16\xf8\x37\x62\x7f\xca\xdd\xc3\x28\x46\x94\x66\xfc\xc7\x7f\xfc\xbf\xa3\x1a\xcb\x52\x7f\xb0\xe0\xc7\xbf\xfd\x10\x89\x7f\x2a\x96\xaf\x2f\xb2\x79\x86\xdf\x64\xc3\x77\x63\xfd\x77\xb2\xa1\xfe\x46\x36\x60\xe2\xb2\x2f\xb2\x49\x84\xd8\x77\xdb\x4b\x2d\x50\xc8\xc2\xa7\xd3\x87\xce\x51\xf9\xc5\xfc\x8e\xcb\x82\x56\x05\xa4\xbb\x5c\x2c\xc8\x5d\x7c\x05\xc7\x42\xef\xd2\x06\x84\x80\x5f\x2c\x08\x13\xdf\x6c\x22\x52\x69\x13\xa9\xfc\x20\x1f\xee\x4d\x34\xf3\x8b\x64\x10\x07\x3e\xc8\xc7\xce\x5e\x84\xf4\x1d\xc3\xbc\x02\xb2\x7f\xd1\x6f\xff\x8c\x7c\x22\x04\x3f\xc4\x35\x06\x66\x64\xfd\x9c\x60\x12\xf3\x18\xd4\x42\xff\x22\xa0\x17\xcc\xc9\xc5\xef\xfd\xc1\x7f\xf6\x07\x0d\x5e\xeb\x18\x9b\x9d\xde\xcd\x42\xf4\x53\x10\xac\x1c\xc3\xc0\xb8\x10\xab\xcb\xc6\x63\x68\x3a\xe2\x8a\x32\x8e\x12\xaa\x4c\x2c\xba\x49\x72\xb4\x39\xe5\xdc\x90\x9d\x5b\xdb\xe2\x25\x56\x2e\x42\x8f\x65\x6a\x4c\xb7\x6a\x06\xfa\x37\xf5\x6e\x17\x0d\x8c\x92\xc7\xf5\x2e\x96\xfb\xde\xcc\xda\x8e\x1b\x33\xdd\x52\x4e\x60\x93\x2a\x49\x34\x64\x5b\x7d\x68\x5b\x49\xaa\x51\xc5\x8a\x9c\xe3\xdc\x43\x22\x30\x5b\xf9\x89\xd9\x74\x37\xe7\xf8\x19\x0a\xad\xeb\xab\x2e\x37\xde\x2c\x02\xa4\x55\x3a\xd0\xcf\xfd\x35\x64\x64\x50\x44\xa5\xb3\x81\x86\xab\x25\x77\x60\x3e\xb4\x63\xdc\xdf\xcc\x91\xae\xdc\x13\x4d\x1a\xd5\x7a\x88\x31\xa5\xb6\x11\x49\x1f\x38\xd1\x97\x4e\xf6\x33\x72\x76\xfe\xb3\x15\xe3\x9d\x4b\xf6\x45\x8f\x57\x14\xe5\x6c\xdc\xd8\xa6\xa6\x61\x5b\x81\x42\x1b\xe7\x7a\xdf\x3f\xf5\x96\x2b\x2f\x40\x4d\x5d\x50\xd4\x26\xc6\x4e\xc7\x8d\xce\x1e\x4d\xc9\xf0\x42\xa6\xc4\x37\xd1\x44\x4b\x52\x48\xf7\x52\xc3\x64\x7a\xaa\x1e\xb4\x0a\xd9\x7a\x62\xa6\xca\x6d\x8c\xce\x61\xe6\x8c\xb6\xda\x89\x00\x35\x4e\x66\xce\xfb\xb0\x5f\xb1\x73\xca\x53\xcf\x29\x26\x72\xf1\x14\xa7\x69\xe3\xa1\xb3\x79\x9d\xd8\xb1\x89\xbd\xb1\xdc\x18\x32\x07\x0c\x00\x1b\x4a\x9a\x90\x2c\x0d\x08\xd2\x6f\x89\x2a\x4f\x46\x80\x60\x08\x5e\x72\x16\xfb\x87\x7a\x16\xe2\xef\x64\x2e\x33\x3c\xf8\x1b\x43\xfd\x4c\x50\x98\x5d\x80\xf3\x9b\x5c\x4c\xf4\x41\x2e\x36\xe0\x5e\x04\xf4\x39\x07\x24\xdc\xf4\x22\x21\x83\x7f\xf5\x09\x2c\xc0\xe1\x23\x97\xe6\xd1\x0a\x66\xc4\x6e\x32\x4c\x0d\x41\x97\xdf\xaa\x6f\x4e\x78\x91\x8b\x5a\x27\x55\x44\x9a\x4b\xe2\x6b\xb8\x5a\x6b\x63\x64\x31\x45\x5c\xcf\xab\x4a\x7e\xf6\x81\xa7\xe5\x6a\xad\x2d\x11\x07\x75\xec\x83\x4c\x01\x6f\xae\x50\x40\x3c\x12\xa1\xe0\x24\x32\x00\x93\xbe\x82\x3d\xcc\xb2\x0e\x66\xbc\x00\x8d\x98\x05\x66\x30\xbc\x99\x4d\x42\x6f\x8e\xf8\xa2\x08\xec\x3f\xe1\x88\xe5\xe2\xc3\x31\x26\x2a\x3c\x22\x95\xfd\xdf\xda\xfb\xa5\xd4\x31\x20\x45\x39\x5f\xef\x75\x6b\x0e\xac\xbc\x00\x62\xdf\xf0\xf2\xba\x05\xf2\x4c\x27\xf3\xb5\x8e\x9e\xd3\x93\x74\xf4\x79\xbb\xa5\xcf\x9c\x6c\x0f\x17\x4e\x57\xd9\x3d\x92\xf0\x00\x17\x3c\xae\xa9\x3b\x15\xd3\x72\x31\x3e\xf0\x79\xdd\x10\xbb\xed\x13\x9e\x1d\x73\x7b\x50\xbd\x6a\x5e\xb7\xa7\xe3\x45\x5c\xe5\x8b\x7e\x7a\xfa\x1a\xb4\x1e\xd5\xa0\x76\xcd\x3c\x39\x2c\x5f\x13\xcf\x8d\x67\x1e\x29\x44\x1c\xe6\x9d\x87\x45\x74\x17\x8e\x92\x56\x94\xba\xb9\x4d\x57\xea\xb6\xac\xb8\x74\x74\x69\x2a\x51\xd6\x02\xec\xa0\xde\x84\x9b\xe0\x58\x4b\x4b\xce\xd1\xd6\xff\x0f\xfc\x24\xff\x29\x7e\x6e\x8b\x5f\xf0\xd3\x10\xc2\xf9\x59\xc1\x7a\x49\x9d\x01\x32\xb7\xb4\x9d\xe6\xec\x2c\xfd\x61\x3e\xf4\x02\xe0\x58\x88\xd5\xdc\x9f\x63\xbe\x50\x00\xe7\x6b\xcc\xe7\x9c\xa4\x8a\x76\x1f\xf5\x23\xf3\xd5\x33\x58\xa8\x59\xb7\x01\xf9\x35\xee\x63\x88\xcd\x8b\x93\x05\x16\xc4\x01\x4a\xe6\xff\x48\xce\x54\x0b\x8a\x11\x29\x67\xc1\xf7\x10\xda\x56\x17\x16\xee\x62\xc2\xc9\x22\x62\xdf\x63\xa9\xcd\x13\x88\x33\xbe\x84\xf8\xf0\x33\x14\x3a\x92\xd2\x46\x9e\xb0\xa4\x9f\x68\xfd\x05\xd6\xa1\xc7\x3c\xbf\x3c\x04\xec\x6d\x22\x58\x60\x90\xf9\x78\xfd\xb0\x38\x00\xf5\x46\xe7\x02\xbc\x11\x1a\x09\x8d\xcf\xd9\x88\x7c\x8d\x12\xc8\x96\x57\xcd\x2e\x77\xba\xdb\x4c\xd9\x25\x7b\xf3\x09\xf6\x33\xa1\xfc\xd3\xe1\xc5\x2e\x27\xed\x8b\x57\x56\xd0\x04\x4b\x99\x61\xaf\x9d\x74\x16\xfe\x91\xfc\x1e\x1f\x6a\x66\x91\xc5\x6f\xbb\xe3\x70\xf1\x8c\x2c\xf1\xb5\x4a\x16\x99\x67\x24\x95\xb3\x56\x00\xfc\x5b\x22\xa9\x3f\xdf\xb6\x00\x97\xd0\xdb\xe5\xf1\xbd\xfc\x1e\xe9\x3f\x26\xfa\x6a\x8c\x6e\xbf\x90\x55\x86\xa9\x16\x94\xd0\xca\xb2\xc0\x92\x27\xce\xf8\x9a\x95\x63\x0d\xbc\x9b\x0b\x70\x59\xc6\x9f\x01\xc7\xb2\xc0\x68\xd8\x2c\xe3\x21\xd0\xd6\x9b\x5f\x2b\x81\x64\x6b\xbe\x6e\xcb\x2b\xb6\x79\x8e\x6a\x4e\x2d\xf7\x1d\xdf\x25\xb0\x29\x1e\x17\x99\x9a\xeb\x93\x43\x95\xf0\x10\xac\xcb\xbe\x38\x6d\xa1\x38\xf2\xd6\x72\x0b\x25\x70\x02\xb3\x12\xb3\xa2\x7b\x27\x6c\xd0\xb6\xc8\x02\xb5\x53\x63\x79\x1f\x56\xb7\x2b\x77\xd8\x51\x56\x16\x9f\x98\xb3\xd8\xc6\x5b\x58\xda\xee\x92\x0e\x0b\x2f\x91\x25\xb2\x8a\xa3\x24\x4a\xf5\x09\x24\x7c\xbc\x97\x6f\x61\x18\x73\x37\x71\x49\x02\x3a\xb7\x88\x36\xbb\x63\xfb\x29\x47\x31\x77\xc0\x17\x53\x85\xb3\x97\x72\xbe\x7d\x5e\xc1\x2c\xf5\xe1\x00\x6d\x6a\x10\x0f\x5a\xcd\x3a\x99\x3a\xe4\x01\x31\x47\x02\x79\xbc\x1a\x8a\x91\x3e\x26\x92\x95\xad\x7c\xba\x9d\xb7\x1e\xa6\xd4\x73\x1f\x00\xa3\xb2\x2b\x77\x73\x4d\x37\xdb\x8a\x2f\xd2\x61\x26\x93\x58\xc6\x05\x9f\xba\x55\xf5\x56\x65\x8f\xcf\xce\x34\x78\x61\x08\x8a\x2c\x71\xb3\x48\x00\x78\x62\xca\xf0\x4e\xa7\x33\x50\x31\xae\x3e\x3d\x93\xca\xdb\xfa\x64\xba\x39\x97\x25\x74\x56\x22\x50\x0c\x8d\x50\x84\x26\x19\xb4\xa7\x72\xb8\xdf\xab\xa1\x43\xc7\xf3\xb1\x45\x27\x5e\xd8\xc2\xc7\xfc\x28\x7a\xc9\x60\x32\x27\xb9\xa6\x17\x16\xcb\x9f\x7a\xb5\x73\x62\xf6\xc3\xfa\xd1\x58\xc4\x83\x89\xfb\xf0\x48\xec\x8f\x51\xc2\x99\xc0\x0b\xc5\xd7\xbf\x8f\x11\xd8\x2f\x1e\x09\x32\xf7\x4b\x73\xc8\xed\x0c\x25\x92\x36\x77\xc4\x2a\x44\x15\x47\x6d\x91\x44\x0f\xbf\xe3\x27\x3b\x7d\xcc\xa2\x18\x30\xc0\x30\x9b\xe7\xb3\xdb\xc0\x63\x7a\xbf\x31\xf7\x03\xd5\xd7\xe9\x75\x24\x86\x8d\xe3\x67\x55\xdb\x6e\x92\x67\x28\x2f\x4c\xdc\xed\xa5\x87\x71\xf0\xf9\x2e\x3f\xd1\x9e\x6d\xeb\x09\x5d\x69\x9a\x0d\xb1\xd5\x39\x40\x1a\xd0\xfc\xcc\x5a\xab\xa6\x44\x62\xfb\x68\x26\x92\xc8\xf3\x98\xb9\xee\x8b\xc6\x53\x8e\xd3\x3e\xad\xac\xfc\x74\xba\xa5\x67\x5e\xa6\x9e\x57\x39\xe2\x41\x7b\x0c\x39\x29\xa9\xb8\xd2\xd0\xb0\x72\x24\x99\x7e\x0e\x88\x4d\xce\x6b\x52\x30\xe9\x90\xaa\x1e\x27\x37\x22\xe3\xd1\x4e\x84\xcb\x4c\x6a\xde\xc5\x7f\xcc\x9a\x4d\xc8\x47\x68\x20\xdf\xe7\xe7\x4b\x18\xf7\x29\x19\x1a\x12\x9f\x76\x22\x87\x01\x2f\x67\x0a\x47\xcc\x2b\x58\x40\xc5\x5d\x69\xfa\x7c\xe0\x40\x93\xd0\x9c\x41\x9e\xd5\x2c\x8f\xad\x12\x68\xe7\xe7\x02\x83\x34\x28\x09\xba\xcb\xb8\x1b\x2e\xf5\xc5\xde\x23\x65\x72\x67\x04\x8e\x84\x59\x4a\x8f\x58\x7a\x50\xce\x24\xd9\x0f\x5b\x52\x2d\xbc\x54\x11\x1f\xe3\xf0\x9c\x0f\x05\x00\x99\x2c\x2c\x0f\xd6\xe7\x84\xc4\x6c\x9a\x81\xa3\xc2\x28\xd1\x2a\x61\xc3\x98\xa0\xf7\x1f\x8f\x90\x25\x44\xac\x6f\xcb\x1b\x15\xf3\xdd\xbc\x5f\x29\x7a\xb7\x69\xb2\x5c\xed\x44\x93\x1f\xfe\x7b\x70\xcf\xdd\xb2\x1b\x9b\x76\x6f\xb8\xff\x15\xed\xbf\x57\x3e\x8d\x99\xd7\x81\xfe\x15\x77\xb9\x7b\x35\x7a\x9b\x7e\x03\x3e\xcb\x8d\x99\xdb\x42\xd6\xec\xc4\x3c\xf1\x3a\xc5\xbb\x48\x92\xf5\x87\x68\x7c\xbb\xcd\x39\xf6\x17\xa9\xea\x22\x88\xbe\xa4\x2a\x32\xc5\xea\x1e\x3a\x5a\x15\xdf\x3f\xa4\x2a\x32\xcb\x49\xf8\x02\xbc\x19\x7e\xce\x83\xaf\x24\xef\x9d\xf4\xe6\x7a\x64\x81\x49\xca\xbe\x2c\x9d\x84\x0f\xbd\xa4\x7d\x63\xd6\xa7\xa5\x13\x78\x53\xe6\xd4\xcc\x98\xb0\x90\xc3\x0c\x1b\xa4\xc2\x84\x4f\x5a\xc1\xcf\xc8\xfe\xc4\x68\x0e\x86\xc2\x84\xaf\xc8\xfe\x35\x86\x44\xfe\x6f\x9a\xdf\x99\xf8\xe9\x0b\xa3\xe7\x9f\x77\x7b\xc1\xee\x2c\xda\xc0\xff\xe2\x27\x5e\x4a\xc8\x64\xd9\x97\xaf\x9c\xe0\x9d\xb3\x2f\xb0\x88\xc0\x67\x99\x03\xe9\x17\xec\x0a\xb8\xd9\x5e\x88\xea\x4d\x31\x32\xff\x35\x89\x43\x2b\x22\x18\xfc\x6f\x2a\x1c\xfb\x9b\x0c\xff\x52\xe1\x54\x31\xe6\xbc\xaf\xa6\x8f\x93\x6d\x37\xbb\xf4\xe1\x9c\x4e\x67\x96\x4c\x8d\x25\x3c\x55\xc7\x5c\xd5\xab\xaa\x87\x18\x6b\x71\x08\x5c\xf4\xbb\x5d\xb0\xa4\xb3\x4d\xb8\x59\x2c\x81\xbb\x27\x6f\x85\xa5\x26\xf9\x74\x01\x40\xe6\x95\xb5\xa7\xa2\xec\xb8\xbd\xae\xf6\x30\x9b\x4f\xd5\xbb\xf4\xb1\xd0\x93\xc1\x5d\x68\x98\x23\x76\xcf\x0f\xee\x35\x3d\x90\xd6\x09\xc9\xc4\x10\x8e\x69\x36\xed\x61\xed\x6e\xad\x0a\xbf\x8e\xf0\x4a\x3d\x8e\xc7\xe7\x49\x4b\xf0\x2b\xd9\x52\xb5\x5b\x0b\xa1\xb8\x8f\xe5\x55\x51\x47\xfc\x08\x0e\x42\x3b\x60\x14\xc9\x6e\x93\x99\xdd\x6e\x81\x44\xe4\x49\xdd\x98\x10\x06\xe8\x51\xed\x44\xcd\x37\x9b\x6e\x2f\x51\xb7\xe4\xc1\x28\xcd\xba\xe1\x4b\x39\x90\x47\xc7\x3f\x1a\x87\x5c\xdd\xe4\x8f\xb2\x45\x57\xfb\x78\xc6\xec\x1d\xd3\x27\xd5\x21\xb3\xd6\xd4\xbc\x0f\xd3\x6d\x5f\x93\xba\x00\x1f\x7a\xd8\x8f\x9b\x35\x42\xce\x05\x5e\x25\x6f\xa9\x53\xd4\x6d\x97\xe3\xd8\xd9\x2c\xae\x3a\x8f\x63\x02\xd1\xd0\xa0\x9d\xbe\xa5\xb0\x2c\xf1\x8e\x85\x9b\x6d\x07\x73\x6c\x2c\xeb\x76\x8b\xb5\x66\x86\x1a\x49\xf6\x79\xb4\xd8\xa0\xd5\xe1\xa1\x18\xac\x62\x97\x0e\xfe\xa6\x28\x91\x3e\x2a\xd3\xa7\x0a\x27\xd0\x14\xd8\xe0\x53\xae\xfe\x0c\xa1\x41\xf0\x8f\xe0\xf6\x17\xeb\x07\x30\x7b\x43\x73\x23\x13\xb3\xd5\xa1\xa6\x4f\xe7\x9b\x38\x8f\x5c\x75\x94\x19\xdc\x75\xfe\xf0\xb4\xbf\x92\xff\x96\x8b\xfd\x46\xb5\xc2\x27\xd5\x42\x80\x9e\xb1\x76\x68\x65\x7e\x96\x4f\x9b\x61\xc7\xdf\x79\xc6\xea\xcf\xf6\xb5\xcb\xec\x5d\xea\x97\x72\xb3\x18\xc1\x96\xef\x3b\x2c\x47\x60\x78\x26\xbe\xb4\x1d\x6d\x28\xed\x0a\xf5\x29\x12\x71\x7a\xb7\xef\x35\x5b\x77\x7c\xe4\x17\xf1\x89\xcd\x57\x67\x91\xeb\x8b\x8e\x4a\x9f\xbb\x3d\x26\x3a\x1e\x5b\xbd\xdd\x12\x67\x77\x0f\x92\x0c\xc3\x0f\xfa\xf6\xee\xf6\x0a\x9d\x78\x3a\xb5\x3b\xe3\x45\x7d\x3e\x9c\x0f\xd7\x68\x10\xa3\x6b\x50\x47\x1e\xee\xed\x61\x71\x95\x87\x21\x6d\x49\xab\xdd\x24\x9e\x2e\x77\x5e\x0c\x81\xbc\xc1\x87\x93\xac\x08\xd8\xf9\x5e\x39\xe5\x2d\x0c\x28\x19\xca\x0f\xee\x3a\xd2\xda\x65\x1f\xc0\x28\xed\xf4\x4c\x5f\xbd\x43\x6f\xc2\x98\x52\xf1\x7b\x48\x39\xce\xe6\xd4\x95\x84\x43\xd3\x07\x07\x44\xa4\xfa\x2c\x78\x17\xaf\x7a\x2c\xb0\xf8\x60\x67\x58\x43\xca\x4d\xbd\xde\xc0\xae\xdd\x3c\x47\xc1\xb4\x5b\x42\xbe\xc3\xa7\x9d\xf7\x54\x2b\xf4\xb7\x91\x29\x9f\x7e\x5c\xb9\x2b\x07\xc4\xc5\x2a\xa8\x64\x6d\x49\x3c\xa1\x87\xc8\x28\x31\x16\x38\xd3\x91\x25\x18\xf6\xc8\x08\xe3\x8e\xdd\x3f\xae\xe7\xae\x06\xbc\xd5\xc5\x37\x79\x6e\x33\xdc\xa7\xf0\xd3\x19\x45\xed\xee\xe2\xb9\x51\x42\xfd\x37\x5d\x8b\x7f\x01\x86\x2d\x09\x10\x5f\x8f\x7c\x56\xda\x8d\x69\xf7\xaf\x78\xe5\xd5\x4b\x56\x91\xcb\x37\x1a\x83\x2b\x51\x1c\x88\xe3\x29\x89\x6f\xb2\x41\x0d\x8b\x45\x6f\xe8\xd2\xfe\x9b\xfc\xfe\x9f\x84\xc6\x33\x5a\xc1\x8a\xb8\x6c\xfe\x50\xcd\x6f\x34\x5e\x7e\x8f\x21\x1e\xff\xaf\xa3\xb1\xb2\x82\xea\x0b\x68\xa5\xf2\x27\xa0\x15\x34\xd9\xb0\xe5\x0c\xfb\x7e\x24\x14\x5e\x3f\xe1\x52\xc5\x37\xc8\x19\xf6\xbf\xae\x65\x31\xa0\x2d\x4f\x4f\x81\xda\x14\x1b\x93\x70\x01\x3b\x8f\xa0\xe8\x2c\xf7\x77\xaa\xb0\x30\x14\x33\x6e\x63\xe7\x7c\xb8\xb0\x1e\xd8\xe9\x54\xd0\x94\xe4\x74\x19\xc0\x29\x26\x28\xb3\xe6\x02\xff\x72\xa2\x49\xec\x7e\x25\x0f\x46\x60\x8d\xf3\x61\x83\xb6\x79\xbb\xda\x31\xbd\x56\x7b\xf6\x35\x4a\xdb\xd9\xf2\xdc\xb7\x77\x57\x56\xcd\x9d\xa7\x91\x63\x88\x6c\x39\xdc\x1d\x96\x0e\x8e\xd6\xc3\x77\x9e\xa4\x89\x47\x09\x76\x6a\xcf\x5c\xc9\xee\xf7\x99\xcc\x8d\xfc\x64\x3e\x04\xee\xb9\xab\xf7\xa7\xcd\x94\x25\x95\x3d\xae\xe1\xf6\xcc\x1c\xba\x61\x8e\xaf\x95\x33\x73\x7a\xb6\x87\x07\xdb\x1f\x6f\x5b\x9b\x25\x3a\xab\x4f\x98\x2b\xb6\xad\x12\xdc\x59\xfb\x5d\xa6\x11\xcf\xa2\xb4\x82\x7e\x4f\xeb\x38\xd3\x65\xbb\xad\xd9\x9d\xc2\x81\x6f\xf1\x21\xc8\x45\xc2\xb2\x28\x89\x23\xee\xba\xca\xb7\x89\x5c\x02\x53\x48\xcc\x2e\xd1\x56\xd1\xc7\xca\x67\xa1\xb9\x00\x3f\xee\x11\xbf\xd5\x5d\xea\xf9\x3c\x6b\xf1\x35\x67\xf5\xbc\xe4\x41\xe1\xb8\xb5\x79\x64\xf3\x4b\x7c\xc8\x1e\xb7\x52\xb1\x47\xdd\x3f\x6c\xb2\xe7\x63\x62\x96\x76\x05\xfa\xa6\xbf\xe2\xd8\xd6\x34\xc9\xe9\x60\x46\x27\xe1\xca\xf6\xdb\x9a\xdd\x46\x89\xe0\x29\x3d\xfb\xfd\x8c\xf9\x9a\x98\x32\xc1\x43\x50\xfe\x7c\x9e\x7b\xbb\x11\x32\x02\xc1\x0b\x80\xd9\x97\xf4\xfc\x70\xe2\x4c\xfc\xfc\x81\xbe\x06\x07\x32\x51\x7c\xcd\x76\xdf\xae\xf5\xc4\x43\xc8\x66\x93\x52\x80\xf2\x27\x43\x61\x02\x41\x10\xfc\xe4\x95\x33\x45\x4c\x82\x3f\xcb\x5a\xea\xbf\x1e\x50\x71\x24\xc9\x93\x76\x00\xbb\x86\xd3\x45\xf4\xe9\x56\x63\xea\xc2\x14\x71\x3d\xad\x6a\xfd\xae\xb2\xe2\x55\x65\x9f\x9e\xc7\xff\xa9\xca\x2f\x7a\x17\x94\x31\x22\x4d\xce\xb0\x3e\x0d\xf7\x03\xd8\x61\x3f\x7f\xed\xf3\x63\xf8\x7f\xe5\x63\xd8\xd7\xd7\xce\xc6\xa7\xe9\xce\x05\x93\xc6\x81\x05\x66\xa6\x2b\x59\x00\x20\xb6\x99\x4e\xe0\x57\xb3\x5d\x12\x5f\x1b\xc2\x55\xbf\xd1\x23\x96\x90\x09\xf9\x97\x1d\x39\x47\x40\xbf\x3e\xa0\x82\x0d\x08\xcc\x85\x45\x9e\xe4\xaa\x4b\x71\x2f\x9c\x80\x1a\xd8\x35\xc7\xe6\xe1\x7e\xfb\xc3\x9e\x02\xe7\xd4\xe1\x5c\x0f\x77\x96\xe3\x65\x64\x0b\x2a\xef\xa3\x69\x9f\x2e\xe7\x44\x70\x7f\x9f\x24\xb0\xaf\x51\x02\x18\xa0\x38\x0f\x8c\x8a\x17\x02\xdb\x46\x97\x8d\x51\x4d\xa1\x89\xe6\x16\xf7\x6e\xb8\xef\x8c\x19\x64\x85\xde\xda\x28\xc3\xaa\x4a\x59\x45\x6d\x6a\x67\xba\x6b\x8f\x5b\x5b\x61\x7b\xbb\x3a\x4b\x8d\x5e\x45\xd5\x12\x34\x37\xa3\xde\x8f\xcf\xfc\x78\x39\x4b\x6a\xa6\x8a\xb6\xf3\xd8\xf6\xbc\x7a\x4d\xb5\xc7\xfa\x38\x57\xc3\x76\x11\x05\x67\xbd\x64\xb6\x14\xe9\x47\xae\x88\x77\x95\x08\x30\x1a\xf9\xde\x3e\x31\x35\x85\xb0\xba\x9a\x8e\x37\xf1\x75\x90\xab\x63\xb6\x46\x71\x69\xc3\xcd\xc1\x6b\xf2\x3e\x10\xae\x97\xd8\xe6\xb9\xa1\xa7\xd5\x7e\x4b\x4a\x90\xc2\xaf\xdb\x9d\x5a\x82\x81\x4b\x58\x8c\xb8\x71\xca\xa1\x3e\x9c\x8e\xb9\x21\x76\x77\xe2\xe6\x5d\x6e\xf2\x03\x0e\xa3\x7a\xb5\x97\xc5\x34\xaf\x4a\x4f\xf7\x85\x7f\x1c\x34\x9a\x6d\x77\x76\xbd\xcf\x3a\xfa\xc2\xe9\x0d\xde\x4d\x20\x0f\xcf\xb7\x01\x8b\x37\xaa\x82\x86\x86\x52\x8c\xc1\xc5\xf7\x4d\x42\x8b\x37\xf1\x68\xed\xa1\xa6\x33\x74\xeb\x82\x64\x5c\xbb\x3d\x4e\xf3\x32\xa7\x47\x83\xd4\xb4\x81\x92\x5a\x54\x72\xc0\x87\x68\x74\x2d\x32\x32\x71\xac\x38\x0c\x8a\xda\xe2\xc7\x72\x35\x1a\x89\x2c\xf4\x4a\xef\xa4\x31\x39\x2b\x91\xb9\xa6\x3e\xa7\xae\xff\x09\x85\xfd\xef\x00\x00\x00\xff\xff\xb0\x06\x44\x31\x9c\x22\x00\x00")

func ca_certificatesPemBytes() ([]byte, error) {
	return bindataRead(
		_ca_certificatesPem,
		"ca_certificates.pem",
	)
}

func ca_certificatesPem() (*asset, error) {
	bytes, err := ca_certificatesPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ca_certificates.pem", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"ca_certificates.pem": ca_certificatesPem,
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
	"ca_certificates.pem": &bintree{ca_certificatesPem, map[string]*bintree{}},
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
