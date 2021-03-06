// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// installer.kubedb.com_kubedbcatalogs.yaml
// installer.kubedb.com_kubedboperators.yaml
package crds

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _installerKubedbCom_kubedbcatalogsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x8f\x1b\x37\x0c\xbd\xfb\x57\x10\xe8\x21\x97\xda\xc6\xa2\x97\x62\x6e\xed\xa6\x05\x82\xf4\x0b\xd9\x34\x77\x8e\x44\x8f\xd9\xd5\x48\x0a\x49\xb9\xbb\xfd\xf5\x85\x34\x1e\xaf\xbd\xf6\x26\x5d\x03\x9d\x9b\x9e\xc8\x47\xf2\x91\xd2\x08\x33\x7f\x22\x51\x4e\xb1\x03\xcc\x4c\x0f\x46\xb1\xae\x74\x75\xff\xbd\xae\x38\xad\x77\x37\x3d\x19\xde\x2c\xee\x39\xfa\x0e\x6e\x8b\x5a\x1a\x3f\x90\xa6\x22\x8e\xde\xd2\x86\x23\x1b\xa7\xb8\x18\xc9\xd0\xa3\x61\xb7\x00\x70\x42\x58\xc1\x8f\x3c\x92\x1a\x8e\xb9\x83\x58\x42\x58\x00\x04\xec\x29\x68\xb5\x01\xc0\x9c\x3b\xb8\x2f\x3d\xf9\x7e\x01\x10\x71\xa4\x79\xe9\xd0\x30\xa4\x41\x57\x1c\xd5\x30\x04\x92\xd5\xb4\xb1\x72\x69\x5c\x68\x26\x57\x19\x06\x49\x25\x77\x70\xd1\x66\xe2\xdb\x07\x72\x68\x34\x24\xe1\x79\xbd\x7c\x8a\x5a\x17\x98\xb3\xba\xe4\xa9\x2d\xa7\x2a\xdf\x97\x9e\xde\xfe\x78\x3b\xa5\xd1\xf0\xc0\x6a\xef\xcf\xf7\x7e\x61\xb5\xb6\x9f\x43\x11\x0c\xcf\x0b\x68\x5b\xca\x71\x28\x01\xe5\xd9\xe6\x02\x20\x0b\x29\xc9\x8e\xfe\x8c\xf7\x31\xfd\x1d\x7f\x66\x0a\x5e\x3b\xd8\x60\xd0\x9a\x8d\xba\x94\xa9\x83\xdf\x6a\x25\x19\x1d\xf9\x05\xc0\x0e\x03\xfb\x26\xee\x54\x4b\xca\x14\x7f\xf8\xe3\xdd\xa7\xef\xee\xdc\x96\x46\x9c\xc0\xca\x9c\x32\x89\x1d\x4a\x9e\xf4\x3e\x74\xfa\x80\x01\x78\x52\x27\x9c\x1b\x23\xbc\xa9\x54\x93\x0d\xf8\xda\x5b\x52\xb0\x2d\xc1\x6e\xc2\xc8\x83\xb6\x30\x90\x36\x60\x5b\x56\x10\x6a\x35\x44\x6b\x29\x1d\xd1\x42\x35\xc1\x08\xa9\xff\x8b\x9c\xad\xe0\xae\xd6\x29\x0a\xba\x4d\x25\x78\x70\x29\xee\x48\x0c\x84\x5c\x1a\x22\xff\x73\x60\x56\xb0\xd4\x42\x06\x34\xda\x6b\x3b\x7f\x1c\x8d\x24\x62\xa8\x22\x14\xfa\x16\x30\x7a\x18\xf1\x11\x84\x6a\x0c\x28\xf1\x88\xad\x99\xe8\x0a\x7e\x4d\x42\xc0\x71\x93\x3a\xd8\x9a\x65\xed\xd6\xeb\x81\x6d\x9e\x6d\x97\xc6\xb1\x44\xb6\xc7\xb5\x4b\xd1\x84\xfb\x62\x49\x74\xed\x69\x47\x61\xad\x3c\x2c\x51\xdc\x96\x8d\x9c\x15\xa1\x35\x66\x5e\xb6\xc4\xa3\xb5\x03\x32\xfa\x6f\x64\x7f\x10\xf4\xcd\x51\xa6\xf6\x58\xdb\xa6\x26\x1c\x87\x03\xdc\x06\xeb\x45\xdd\xeb\x68\x01\x2b\xe0\xde\x6d\xca\xff\x49\xde\x0a\x55\x55\x3e\xfc\x74\xf7\x11\xe6\xa0\xad\x05\xa7\x9a\x37\xb5\x9f\xdc\xf4\x49\xf8\x2a\x14\xc7\x0d\xc9\xd4\xb8\x8d\xa4\xb1\x31\x52\xf4\x39\x71\xb4\xb6\x70\x81\x29\x9e\x8a\xae\xa5\x1f\xd9\x6a\xa7\x3f\x17\x52\xab\xfd\x59\xc1\x2d\xc6\x98\x0c\x7a\x82\x92\x3d\x1a\xf9\x15\xbc\x8b\x70\x8b\x23\x85\x5b\x54\xfa\xdf\x65\xaf\x0a\xeb\xb2\x4a\xfa\x75\xe1\x8f\x2f\xa6\x53\xc3\x49\xad\x03\x3c\xdf\x2b\x17\x3b\x74\x72\xea\xef\x32\xb9\xda\xad\x2a\x59\xf5\x82\x4d\x12\x10\xf2\xac\xf3\x49\x39\xa2\xb9\x74\x14\xf7\xb7\x52\xe5\x3a\x05\x5f\x36\xaf\x1f\x05\x54\x63\xa7\x54\xf5\x39\xdf\x9e\xeb\xea\x53\x0a\x84\xf1\xdc\xdd\x9c\x7f\xbd\xd7\x48\xa3\x43\xb7\xa5\x6b\x5c\x53\x1c\xd2\x15\x6e\x8f\xfa\x39\xbc\xde\x2d\x93\xb8\x14\xf1\xc1\x04\x7d\x7f\x85\xfb\xd0\xa7\x12\x1d\xc9\x15\xae\x49\x6d\x90\x4b\x1d\xfb\xaa\xa7\xa4\x87\xeb\xaa\x6d\xe3\xf6\x5a\xb7\x8b\x73\x5f\xbf\x4d\x09\xa1\xfe\x31\x7f\xdf\x91\x08\x7b\x7a\x4e\x7c\xf1\x64\xd5\x8f\x47\x1c\xce\xac\xbf\x34\xc3\x42\x03\xab\xc9\xe3\xcb\xa9\x5f\x88\x02\xed\xfa\x61\x39\x1f\xc2\xe5\x81\xf0\xbf\x96\x7a\x45\x99\x97\x82\x2f\xc1\x1d\x3d\x0f\x66\xac\xc9\xf1\xe5\x8b\xe6\x19\xb4\x9b\x1f\x5f\xbb\x1b\x0c\x79\x8b\x37\x4f\x58\x53\x6f\xb9\x7f\x19\x1d\x6d\x03\xb4\x57\x83\xef\xc0\xa4\x4c\xd1\xd4\x92\xd4\x46\x4c\xc8\xbf\x01\x00\x00\xff\xff\x5f\xc6\x97\x35\xd4\x09\x00\x00")

func installerKubedbCom_kubedbcatalogsYamlBytes() ([]byte, error) {
	return bindataRead(
		_installerKubedbCom_kubedbcatalogsYaml,
		"installer.kubedb.com_kubedbcatalogs.yaml",
	)
}

func installerKubedbCom_kubedbcatalogsYaml() (*asset, error) {
	bytes, err := installerKubedbCom_kubedbcatalogsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "installer.kubedb.com_kubedbcatalogs.yaml", size: 2516, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _installerKubedbCom_kubedboperatorsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x6d\x73\xdc\x38\x72\xfe\xee\x5f\x81\x72\x52\x65\x29\xa5\x19\xad\x6f\x53\x57\x89\xee\xea\x52\x3a\xdb\xbb\xe5\x5a\xbf\xa8\x2c\xef\xee\x25\x97\xcb\x16\x86\xec\x99\xc1\x89\x04\xb8\x00\x28\x79\x36\x95\xff\x9e\x42\x03\xe0\x90\x43\x02\x04\x47\xd2\xd9\x7b\x47\x7e\xb1\x35\x24\x9b\xdd\x8d\x46\xa3\x5f\x1e\x12\xb4\x62\x3f\x80\x54\x4c\xf0\x0b\x42\x2b\x06\x9f\x34\x70\xf3\x97\x5a\xde\xfc\x9b\x5a\x32\x71\x7e\xfb\x7c\x05\x9a\x3e\x7f\x72\xc3\x78\x7e\x41\x5e\xd4\x4a\x8b\xf2\x03\x28\x51\xcb\x0c\x5e\xc2\x9a\x71\xa6\x99\xe0\x4f\x4a\xd0\x34\xa7\x9a\x5e\x3c\x21\x24\x93\x40\xcd\x8f\x1f\x59\x09\x4a\xd3\xb2\xba\x20\xbc\x2e\x8a\x27\x84\x14\x74\x05\x85\x32\xd7\x10\x42\xab\xea\x82\xdc\xd4\x2b\xc8\x57\x4f\x08\xe1\xb4\x04\xff\xa7\xa8\x40\x52\x2d\xa4\x5a\x32\xae\x34\x2d\x0a\x90\x4b\x7b\x66\x99\x89\xf2\x89\xaa\x20\x33\x24\x36\x52\xd4\xd5\x05\x19\xbc\xc6\x12\x74\x4f\xca\xa8\x86\x8d\x90\xcc\xff\xbd\xd8\x3f\xd6\xfc\x41\xab\x4a\x65\x22\x07\xfc\xd3\x8a\xf9\x5d\xbd\x82\x97\x7f\x7c\xef\xf8\xc0\x13\x05\x53\xfa\xbb\x81\x93\x6f\x98\xd2\x78\x41\x55\xd4\x92\x16\x3d\x19\xf0\x9c\x62\x7c\x53\x17\x54\x1e\x9e\x7d\x42\x48\x25\x41\x81\xbc\x85\xef\xf9\x0d\x17\x77\xfc\x1b\x06\x45\xae\x2e\xc8\x9a\x16\xca\x70\xa4\x32\x51\xc1\x05\x79\x67\xa4\xa9\x68\x06\xf9\x13\x42\x6e\x69\xc1\x72\xd4\xb0\x95\x47\x54\xc0\x2f\xaf\x5e\xff\xf0\xf5\x75\xb6\x85\x92\xda\x1f\x0d\x65\xf3\x18\xdd\x88\x6d\x95\xde\x0c\x77\xf3\x1b\x21\x39\xa8\x4c\xb2\x0a\x29\x92\x67\x86\x94\xbd\x86\xe4\x66\x80\x41\x11\xbd\x05\x72\x6b\x7f\x83\x9c\x28\x7c\x0c\x11\x6b\xa2\xb7\x4c\x11\x09\x28\x03\xd7\xc8\x52\x8b\x2c\x31\x97\x50\x4e\xc4\xea\xaf\x90\xe9\x25\xb9\x36\x72\x4a\x45\xd4\x56\xd4\x45\x4e\x32\xc1\x6f\x41\x6a\x22\x21\x13\x1b\xce\x7e\x69\x28\x2b\xa2\x05\x3e\xb2\xa0\x1a\x9c\x76\xfd\xc1\xb8\x06\xc9\x69\x61\x94\x50\xc3\x19\xa1\x3c\x27\x25\xdd\x11\x09\xe6\x19\xa4\xe6\x2d\x6a\x78\x89\x5a\x92\xb7\x42\x02\x61\x7c\x2d\x2e\xc8\x56\xeb\x4a\x5d\x9c\x9f\x6f\x98\xf6\x06\x9e\x89\xb2\xac\x39\xd3\xbb\xf3\x4c\x70\x2d\xd9\xaa\x36\xc3\x76\x9e\xc3\x2d\x14\xe7\x8a\x6d\x16\x54\x66\x5b\xa6\x21\xd3\xb5\x84\x73\x5a\xb1\x05\x32\xce\x35\xce\x92\x32\xff\x27\xe9\x66\x83\x7a\xd6\xe2\x54\xef\xcc\xb0\x29\x2d\x19\xdf\x34\x3f\xa3\x71\x05\xf5\x6e\xac\x8b\x30\x45\xa8\xbb\xcd\xf2\xbf\x57\xaf\xf9\xc9\x68\xe5\xc3\xab\xeb\x8f\xc4\x3f\x14\x87\xa0\xab\x73\xd4\xf6\xfe\x36\xb5\x57\xbc\x51\x14\xe3\x6b\x90\x76\xe0\xd6\x52\x94\x48\x11\x78\x5e\x09\xc6\x35\xfe\x91\x15\x0c\x78\x57\xe9\xaa\x5e\x95\x4c\x9b\x91\xfe\xb9\x06\xa5\xcd\xf8\x2c\xc9\x0b\xca\xb9\xd0\x64\x05\xa4\xae\x72\xaa\x21\x5f\x92\xd7\x9c\xbc\xa0\x25\x14\x2f\xa8\x82\x47\x57\xbb\xd1\xb0\x5a\x18\x95\x8e\x2b\xbe\xed\x9d\xfc\x31\x34\x3d\xcc\x81\x9e\xa8\xf3\x0b\x21\x25\xfd\xf4\x06\xf8\x46\x6f\x2f\xc8\x6f\xbf\x3e\x38\x57\x51\x6d\x4c\xf2\x82\xfc\xcf\x9f\xe9\xe2\x97\xbf\x9c\xfc\x79\x41\x17\xbf\x7c\xb5\xf8\xf7\xbf\xfc\xcb\x9f\xdd\x7f\x4e\xff\xe3\x9f\x0f\xee\x19\x64\xd2\xff\x6c\x07\xb0\xf9\xd9\xbb\xbb\x41\xa3\xe9\xfa\xa2\xeb\x0a\x32\x63\x41\x66\x18\xcd\x6d\x64\x2d\x24\x91\x90\x33\xe5\x67\x6f\x82\xfc\x34\xcf\xd1\xa5\xd3\xe2\x4a\xe4\xd7\x90\xd5\x92\xe9\xdd\x95\x28\x58\xd6\xbb\x94\x10\xa6\xa1\xec\xfd\x18\x94\x6f\x7f\x8a\x4a\x49\x77\xdd\xc7\xae\x71\x2d\xd9\x1d\x12\xeb\x88\xfb\x7a\x8d\x72\xb1\x35\x83\xfc\x0c\xc5\xac\x44\xfe\x4c\xa1\xdf\xc8\xeb\xc2\xcc\x90\x4c\x70\xa5\x25\x65\x5c\xab\xc3\x81\x0a\x08\x6c\x0e\x2e\x72\xb8\x0c\x70\xd0\xe3\xe2\x25\xfe\xb1\x02\x85\xb7\x35\x9c\xb7\xb9\x90\x75\x01\x0a\xd5\xef\x98\x5c\x0e\x10\x8d\x31\x64\xcf\xc3\x1a\xa4\x84\xfc\x65\x6d\x14\x79\xdd\x90\x7f\xbd\xe1\xa2\xf9\xf9\xd5\x27\xc8\x6a\x7d\xe0\xd1\x83\xbc\x7f\x34\xa6\x61\x09\x81\x24\x77\xac\x28\xdc\x63\x8c\xcf\xf5\x27\x0c\xc3\xe8\x84\x8d\x7c\x87\x6a\xdc\x1f\x7a\x4b\x35\x51\x54\x33\xb5\xde\xa1\x9c\x8d\x26\xe0\x93\x71\x3e\x18\x47\xec\x07\x8c\xac\x76\xce\xef\x98\x35\xee\x2c\x48\x76\x55\x6b\xc2\x34\x3a\xab\x6c\x2b\x84\x02\x42\xad\xa2\xf1\x79\xb7\x4c\xe0\xb2\x40\x04\x07\x22\x24\x29\x8d\x97\xc1\xa5\x08\x82\x14\x5b\xec\x2c\x51\x03\x7b\x72\x4c\x91\x52\x28\xbd\xd7\xb5\x9f\x3f\x86\xfc\x1d\xd3\xdb\x88\xf4\x40\x36\x26\xd2\x01\xa5\x89\xaa\x4b\xc3\xc4\x1d\xb0\xcd\x56\xab\x33\xc2\x96\xb0\xc4\xe1\x07\x9a\x6d\x5b\x8f\x2b\x01\x7a\x76\xb9\x3f\x68\x51\x38\x51\x3a\xb6\x04\x3f\xd7\x4c\x42\x69\x7c\x39\x39\x69\x1c\xbf\x73\xc6\x67\xfe\x7c\xcf\x4a\xc2\x8f\x19\x18\xa6\x33\x02\x3a\x5b\x9e\x9e\x91\x4c\x94\x55\xad\x8d\xce\x8d\x4c\xab\x9d\x99\xe2\x92\xba\xc5\x47\x8a\x7a\x13\xd7\x08\x14\x8e\x51\x1f\x1d\xe0\x60\xe3\x32\x6d\x1c\x0b\xdf\x90\xa7\x56\x49\x4f\xfd\x22\xaf\xea\x32\x48\x91\x59\x65\xa0\xfe\x4a\xaa\xb3\xad\x8b\x45\x32\x21\x25\xa8\x4a\x70\xa4\x88\x67\x5e\xed\x65\xf9\x5d\xd4\x18\x0c\xb1\x13\x75\x8a\x83\x8b\xc4\xb6\x6c\xb3\xf5\x63\x48\x25\xe0\x6f\x5d\x9b\x18\x9a\xbc\x24\xec\xfd\xfc\xd1\x99\x78\x97\x9c\x40\x59\xe9\x5d\xcb\xd2\x5a\x63\xac\x41\x96\x8d\x84\x14\x63\xe5\xd0\x61\x97\x07\x65\xf9\x67\x65\x65\x1c\xb3\x76\x96\x47\xbe\x22\x27\x68\x7a\x4c\x3f\x53\x38\x6d\x16\xa2\x3a\x5d\x92\x4b\x1f\x80\x87\x8e\x71\xa6\xb8\x68\x9e\xec\x1e\x61\x18\x55\x22\x42\xb4\x79\x7e\xf0\x9a\x31\x0f\xd8\x66\x0e\x78\xd6\x5b\x97\xbb\x47\x57\xdf\xd6\x6a\x14\x14\x90\x69\xe3\x87\x41\x96\x67\x84\x2a\x25\x32\x66\xa2\x95\x66\xfc\xa3\x24\xc9\x81\xa9\x59\x35\x87\x05\x4a\x17\x8a\x60\x58\xd1\x35\xdc\xb1\xeb\x7b\x22\x9a\xa4\xc4\xcc\xb4\xae\xa8\x6d\x87\x31\x4a\x91\x98\x39\x6e\xee\x7f\xa6\x5c\x7a\x16\x97\x8e\x8c\xdb\x7d\x90\xdd\x20\x9b\x2e\xec\x75\x67\x12\x08\xbb\xc5\xc7\x84\x8e\x94\x71\xe5\x42\xfd\x33\x42\xc9\x0d\xec\x6c\x56\x60\x12\x0f\x17\x17\xe1\xc5\x49\x54\x25\xd8\xc5\xc5\xf8\x80\x1b\xd8\x21\x21\x97\x46\x24\xdc\x9f\x3e\xf2\xf6\xb8\x81\xc1\x60\x63\xe8\xe8\x2d\xe2\x38\x56\xc8\x23\x6a\x02\x3d\xe9\x14\xfd\x11\x9b\x82\x17\x0c\x30\x9c\x4f\xbc\x27\x12\xd8\x0d\x1f\x7e\x08\x8e\x92\xf3\x43\x93\xc3\xd8\x81\x7d\xa6\xec\x00\x99\xb9\xb2\x65\x55\xb2\x9c\x5a\xa0\x75\xe1\x54\xf1\x49\xe1\x0f\x26\x89\x6e\xd8\x53\xe8\xf9\x5f\xf3\x70\x54\x72\x78\xbc\x13\xfa\x35\x3f\x23\xaf\x3e\x31\x65\x16\xfc\x97\x02\xd4\x3b\xa1\xf1\xcf\x25\xf9\x56\x5b\x1b\x7c\x33\xe2\x2a\x5a\x2c\x4e\x55\xac\x95\xe3\x28\xb5\x5e\x72\x1b\x7f\x1b\x75\xb4\x53\x4d\xb5\x34\x01\xf6\xb8\x4b\xdc\x1f\xcd\x04\x63\xca\x24\x7f\x42\x7a\xb5\x60\xc1\x00\x69\x0e\x84\xfa\xb1\xa3\xac\x15\xe6\x94\x5c\xf0\x05\xae\x97\x9e\xa7\xce\xb3\xac\xd6\xd3\xd9\x94\x9d\xf1\xe9\xb3\xe7\x1f\x9b\x4c\x31\xcc\xda\xb7\xda\x3c\xee\x4d\xe7\x21\xe9\x13\x72\xcf\xcc\x96\xde\x62\x10\xc6\xf8\xa6\x68\xc2\xaa\x33\x72\xb7\x65\xd9\x16\xe3\xf6\x64\xa2\x2b\xb0\x55\x93\x4a\x82\x59\xf7\xa8\x32\xae\xd1\xfc\xb2\x01\x69\xc2\x61\xe6\x95\xc0\xd2\x19\x95\x50\x15\x34\x83\x9c\xe4\x18\x74\xda\x9a\x05\xd5\xb0\x61\x19\x29\x41\x6e\xc0\xa4\xc5\xd9\x36\xd5\xfa\x93\x17\x14\x7b\x4c\x9e\x2c\xe1\xb4\x73\xf8\xf0\x21\x75\x0a\x4b\x0b\xe3\x99\x92\xae\x13\xed\x7a\x62\x0a\xbb\x07\x95\x80\xf8\xc5\x29\xb2\x61\xc0\xe1\x4a\x8c\x9f\x39\xd6\xc0\xbc\x60\x8e\x35\xe6\x58\x23\x78\xcc\xb1\x86\x3f\xe6\x58\x63\x8e\x35\xe6\x58\x63\x8e\x35\x7e\x45\xb1\x46\x22\x51\x5b\x4f\x99\x50\xd6\xf9\xd1\xd6\xb9\x0e\xeb\x38\x18\xd8\xf8\x06\x59\xa7\x64\x33\x22\x91\x09\x13\xae\xdd\x5a\xf6\x11\x4b\x44\x8c\x23\x11\x49\xf9\x06\xc8\xf3\xc5\xf3\xaf\xbe\x8a\x5b\xd6\x5a\xc8\x92\xea\x0b\x63\xe5\x5f\xff\x26\x41\x27\x6e\x36\x04\xaf\x1c\xb7\x87\x45\xab\x22\x16\xb9\xc8\xea\x36\x5c\xad\x1d\x1f\xa1\xb1\xc1\x0e\x55\x9e\xef\xd1\x9f\x70\x5e\xae\x29\x51\x77\x8a\xdf\xbd\x56\x42\x50\x38\x57\x75\x96\xc6\xb9\x6b\x52\x82\x26\x54\x77\x4a\x9b\xac\x84\xa6\x81\x64\xdb\x20\xb6\x99\x19\xa4\xe8\x7b\x23\x39\x11\xdc\x55\xae\x8d\xed\x2c\x13\x39\x0e\x77\x3b\xda\x4d\x11\x92\x01\x55\x60\x62\x88\x15\x34\x5c\x8b\xd2\x70\xc9\xb8\xf6\x0e\xd0\xb0\x0c\x5e\xab\x41\xc2\x27\xb0\xdc\x2c\x49\x5e\x23\x39\xca\x5d\x97\xf6\xd4\x4a\xad\x76\x4a\x43\x89\x3d\x16\x21\xf1\x1f\x23\xbe\x96\x3b\xa2\xc3\x15\x5d\xb8\x05\xae\x6b\x5a\x14\x3b\x02\xb7\x2c\xd3\x8d\xfe\xb0\x91\xcc\xb4\xed\x87\x85\x66\x4b\x4a\xc0\x7a\x38\x1b\xa3\x7e\xfa\x20\x7c\xb3\xa6\xb8\x0c\x66\x2a\xda\xd0\xc3\xf6\x4f\x7c\x92\x9a\xcb\xd0\x72\xde\x7f\x08\x57\xfe\x49\xda\x42\x72\x98\x93\xd4\x45\x61\xf4\x6d\x1b\x01\x7d\xf6\x7c\xb1\x7d\xd4\x67\xf9\x52\xbc\xed\x66\x75\x2c\xce\xf6\x8f\x6c\x27\xe3\xf2\xdd\x4b\xa3\x91\x31\x91\x09\xf9\x28\x2a\x51\x88\xcd\xae\xad\x7b\x9c\xfd\xd8\x60\x70\x94\x29\x51\xf5\xca\x45\xb6\xe3\x81\xdb\xbb\x83\xa1\x9c\x6b\xe6\x73\x1e\x3b\x74\xcc\x79\x6c\xef\x98\xf3\xd8\x44\x16\xe7\x3c\x16\x8f\x39\x8f\x9d\xf3\xd8\xd1\x63\xce\x63\x07\x2e\x9e\x6b\xe6\x73\xac\x11\x39\xe6\x58\xa3\x77\xcc\xb1\xc6\x1c\x6b\xcc\xb1\xc6\x1c\x6b\x44\x8f\x39\xd6\x18\xb8\xf8\xc1\x6a\xe6\xe3\xe4\xc6\xd4\xb3\xe8\x17\xda\xa2\x15\xe0\x20\x4b\xd1\xd3\x95\xc8\x8f\x80\xd4\x57\x22\x8f\x20\xea\x6d\x51\x33\x13\x8b\x42\x64\x54\x0f\xfb\x03\x2c\xa7\x1a\x32\xae\x92\xaf\x68\x69\x6b\xb5\x67\xe4\x17\xc1\xc1\x22\x9d\xcd\x34\xc3\xca\xaa\xd0\x5b\x90\xe6\xf2\x13\x75\x3a\x88\x54\x9d\x51\xfa\x83\xc7\x8c\xd2\x9f\x51\xfa\xee\x68\xa3\xf4\xb7\x54\x59\xbb\xb4\x0b\x61\x18\xb4\xdf\xf2\x0e\xc6\x01\xfd\x2e\xca\xef\x67\xc2\xec\x1b\x23\x74\xc6\x82\xaf\x32\xee\x07\xde\xca\x95\xbb\x76\x24\xe4\x57\x5d\x69\x22\xde\xdb\xe6\x70\xc8\x34\xcd\x73\xc8\x49\x05\x72\x61\x4d\x4f\x90\x35\xe3\xf9\x80\x2c\x5e\xfe\x20\xd9\x44\x1c\x7d\x97\xc9\x09\xad\x8b\x76\x77\xa5\xe3\xa0\x0f\x51\xf5\x23\x6b\x61\x33\x7e\x8f\x89\xaa\xc7\xcc\xcb\x2f\x6e\xd3\x53\x76\xcc\xdb\x7e\xae\x41\xee\x88\xb8\x05\xb9\xcf\x4c\x9a\xf7\x3c\x53\x92\x10\x5c\x7b\x98\x22\x19\x55\xd6\x51\x8f\x87\x5a\xd3\xb2\xd3\xe9\x7d\x90\x9e\xb0\x87\x24\x6c\x96\xef\x6b\x16\xa8\x88\xc4\xe8\x6d\xb0\xb4\x31\xd0\x9c\xa2\x32\x35\x84\xb7\xad\xab\xa4\x8b\x27\x05\xa7\x83\xa3\x1d\x28\x79\xa4\xa7\x05\xad\x36\xde\x48\xd9\x23\x9d\xe6\x41\x79\xe4\x9e\xa5\x0f\x72\x44\xf9\x83\x4c\x2b\x81\x90\x43\xf5\x1a\x2e\xdd\x3a\xdd\xaf\x86\x4c\x20\xda\xb2\xaf\xe9\x15\x11\x72\x5c\x3e\x32\xbd\x32\x42\x0e\xc5\x6f\x86\x4f\xf6\xca\x24\x93\x84\x6f\x97\x54\xc2\xa5\x92\x49\x24\x7b\x65\x95\x6e\xb9\x04\x6d\xab\x53\x31\x79\x6c\x65\x4f\xab\x96\x90\x43\x55\xbb\x5a\x01\xc3\xd4\xf9\xa0\x76\x32\x49\x31\xdd\x3a\x4b\xb0\x7e\x32\x89\x66\xa8\x98\xd1\xad\xa1\x4c\x26\xd9\xaf\xb7\xf4\xea\x28\x0f\xc3\xa6\x63\x71\x5f\x88\x98\x44\xd6\x7e\x20\xe2\x21\x8b\x11\x64\x7a\x41\x82\x1c\x6b\x97\x53\x0b\x13\x64\x62\x71\x82\x4c\x28\x50\x90\xa9\x45\x0a\x32\xb5\x50\x41\x26\xcb\x8b\x21\xc4\x9b\xd6\x57\x5d\xc6\x8e\xd6\xd7\x05\x26\xaf\x46\x93\x47\xb0\x1f\xed\x58\x56\x6d\xa0\x53\xd2\xca\x78\x89\xff\x35\x4b\x33\x1a\xfe\xff\xa5\xae\xa3\x94\x49\x65\x42\x61\x57\xfc\x6b\x51\xf0\x35\x87\xd6\xc3\x12\x89\x1a\x6e\x98\x22\xc6\x76\x6e\x69\x61\x02\x10\x0b\xdb\x72\xa9\x9a\xe1\xf4\x30\x5e\x4b\x9d\xdf\x77\x5b\x93\x9e\x9b\xc5\xd7\xa6\x79\x4c\x91\xa7\x37\xb0\x7b\x7a\xd6\xf3\x23\x4f\x5f\xf3\xa7\xa9\x54\xa9\x4b\x55\x3a\x3e\xa3\x89\x7c\x04\x2f\x76\xe4\x29\x9e\x7b\x9a\x3a\xb1\x87\xc2\xc5\x29\x81\xe0\x11\x45\xb9\xa4\x8b\xb9\xff\xf8\xce\xd4\x06\xe0\xfe\xc6\xa6\xbe\xe2\x13\xe3\xfd\xa9\x94\x6a\xa3\x8f\xa0\xae\xfb\x71\x10\x39\x69\x5e\x1b\xdf\x18\xcd\xeb\xd3\x70\x2a\xdd\x12\xa9\x83\x44\xc3\x90\xbf\x04\xca\x15\x79\xea\xab\x67\xcf\xd4\x9e\xc7\xa7\x0f\xd7\x71\x9c\x34\x87\xd3\x7d\x91\x76\x00\xb6\xef\x52\xc2\xd5\x83\x1c\xdf\x55\x0b\xdd\x57\x89\x56\xb0\x2f\x2f\xe6\xe4\xc4\x67\xba\xe1\xdc\x7b\x7f\x08\x89\x28\xca\xce\xed\x5c\xb3\x45\x43\x63\x9f\xff\x9a\x8c\x30\xd5\xbd\x7a\x58\x73\xd7\x02\x7c\x71\xb3\xa9\xdb\xed\x2d\x2a\x65\x06\xdf\x6d\x41\x76\x24\x65\xca\x7d\xed\x09\x3b\x10\xb2\xe6\xdc\x3c\x57\x70\x57\xd6\x4b\x22\x69\xdc\x8c\xfd\x68\x91\x2b\x93\xd8\xb0\x1f\xa5\xc6\xd8\x7f\x3f\x4a\x89\x50\x47\xe2\x0b\x98\xf8\x25\x29\x87\x99\x14\xdc\x4d\x22\xf3\x8b\xaf\xc4\xa1\x5e\x20\x4f\xd5\x2c\x6b\x64\x5c\x92\x57\x38\x09\xda\xcc\x31\x85\x23\x49\x8b\x42\xdc\xa5\x78\x9f\x64\xab\x4e\x8b\x0d\x16\x6d\x66\x1e\xa2\x65\x30\x19\x66\x7f\xf7\xc0\x30\xfb\x83\xd2\xd3\xaf\x04\x65\x9f\x58\xd4\x9b\xa1\xf6\x33\xd4\xbe\x05\xb5\xc7\x9b\xac\xe7\x1b\xc7\xdc\x87\x6d\x06\xb1\xf8\xa9\x98\x7b\xf2\xe3\x16\x70\x46\x45\x0a\x6c\x66\x88\xca\xba\xd0\xac\xda\x37\xac\x95\x65\xad\xb0\xe9\xa3\x05\x2a\xa9\x83\xea\x6c\xec\x8d\x00\x9a\x6d\x0f\xa7\x09\x3e\x07\x1b\xda\x0a\x3d\xb2\x6b\xb3\xd0\xa2\x70\xd8\x7a\x93\x57\x86\xc7\x08\x5c\xaf\x8a\x3d\x4c\x09\xff\xa5\xfb\x82\x61\x53\x34\xc1\xe6\xc4\x89\x59\x2c\x0b\x63\x0e\x66\xc9\xf2\x5e\x2d\xd6\x73\xed\xad\xbf\xb6\x2a\x73\x0b\xbe\x41\xb2\x61\xb7\xc0\xf7\x8b\xf0\x89\x3a\x3d\x1d\x83\x35\xe9\xc4\xd0\xa3\x1f\x58\x44\x88\x0e\x85\x1c\x67\x89\xcb\x7d\x84\x6c\x13\x08\x24\x2c\xf3\xbf\x6f\xad\x5e\x7f\x88\xd0\xdc\x37\x87\x82\x0b\x3c\xaa\xa7\x59\xe2\x9b\x01\x8c\x10\x65\xe3\xd2\xa4\xd5\x41\x27\xb4\x11\x8e\x68\x21\x10\x16\x76\x27\xf6\x98\xd2\x3e\xf8\x9b\xbd\x3e\x91\xd0\x32\x98\x02\x73\x1b\x6f\x17\xa4\xe6\x7f\xc7\x42\x1e\xa3\x0d\x80\x19\xf3\x18\x3d\xd2\x8b\xfd\x7f\x7f\xd0\xc7\x48\x71\xff\x0b\xc5\x40\x1e\x5d\xd4\xff\x5b\x42\x1f\x63\x85\xfc\x89\xdd\x2e\x32\x56\xc4\xbf\x27\x00\x70\x0c\x04\x99\x4c\x33\x50\xbc\x1f\x2e\xc8\x27\x53\x1d\x2a\xdc\x0f\x16\xe3\x93\x29\xce\x08\xc2\xd1\xeb\x3e\x37\x82\x70\x62\x41\xfe\xd8\x62\xfc\xa4\xd1\x99\x5a\x84\x77\xe5\xf5\x04\x36\x12\x0b\xf0\xfd\xd2\x7a\x8a\x88\xa3\xc5\xf7\xc3\xb2\x7a\x5a\xd1\x29\x56\x78\x1f\x2c\xa9\x27\x90\x1d\x2e\xba\xdf\x2b\x9c\x4a\xb6\xce\xc4\x0b\x53\x4b\xe8\xd3\xcb\xe7\x09\x58\x82\x09\xa5\x73\x5f\x18\x1f\xa1\xf8\x10\x65\xf3\x24\x8f\x98\x3c\xd3\xd2\x3c\x44\x72\x99\xfc\x31\x4a\xe4\x13\xcb\xe3\x29\x69\x39\x19\x4c\xcd\x63\xa5\x71\x9b\x09\x8f\x90\x4c\x2f\x8b\xb7\xb3\xe1\x31\xf1\x53\x4b\xe2\xed\x7c\x78\xac\x33\x95\x54\x0e\xef\x17\xbb\xd3\xbb\x29\x93\x4a\xe1\x49\xd6\x9a\x52\x79\x4d\x29\x7f\xdf\xbb\xa8\x3a\x0a\x5e\xe7\x9a\x1d\x0b\x60\x6f\xdb\x75\x00\xc5\x3e\xc8\x33\xbd\x15\x2c\x27\x55\xad\x1d\x94\x77\x32\x92\x7d\x90\xea\x3f\x14\xba\xbd\xa3\xfa\x28\xc4\x3d\x5e\xd2\x3e\x3b\x02\xe2\x1e\xa4\xe8\xa6\xe5\x11\x10\xf7\x30\x49\x07\x7d\x3f\x0a\xe2\x1e\xa4\x8a\xd0\xf7\xe3\x20\xee\xa3\x33\xfe\xd0\x84\xc2\x63\xe5\x71\xee\x41\x92\xe3\xf8\xf7\x08\xce\x3d\x5c\x21\x8f\xe2\xdf\x23\x38\xf7\xb0\x3a\x93\xf1\xef\x3d\x9c\x7b\xc4\xe4\x67\xfc\xfb\xc1\x31\xe3\xdf\x5b\xc7\x8c\x7f\x4f\x14\x76\xc6\xbf\xcf\xf8\xf7\xb1\x63\xc6\xbf\xcf\xf8\xf7\x19\xff\x3e\xe3\xdf\x67\xfc\xfb\x8c\x7f\x1f\x38\x66\xfc\xfb\x8c\x7f\x9f\xf1\xef\xad\x63\xc6\xbf\x8f\x88\x32\xe3\xdf\x67\xfc\xfb\x8c\x7f\x9f\xf1\xef\x33\xfe\x7d\xe0\x92\xcf\x82\x7f\xef\x14\xa1\x83\x20\xf8\x48\x39\x76\xff\xfd\x94\x89\x20\xf8\x20\xcd\x15\x8c\x83\xe0\x83\x6c\x07\xa9\x06\xbe\xf1\x93\x84\x84\x0f\x97\x5e\xdb\x08\xf9\x49\x48\xf8\x48\xd1\x7c\xe0\xab\xf4\xf7\xfc\xfa\x3c\x69\x21\xe4\x8f\x45\xc2\x87\x4d\x40\xcc\x48\xf8\x19\x09\x3f\x23\xe1\x67\x24\xfc\x8c\x84\xb7\xc7\x8c\x84\x9f\x91\xf0\x33\x12\x7e\x46\xc2\xcf\x48\xf8\xde\x31\x23\xe1\x07\xd9\x9d\x91\xf0\x33\x12\x7e\x46\xc2\xef\x8f\x19\x09\xdf\x3d\x66\x24\xfc\x8c\x84\x8f\x1c\x33\x12\xfe\x71\x90\xf0\xc1\x53\x94\x73\xa1\x6d\x70\x7f\xc8\x7f\xda\x62\x1a\x51\x51\xf8\xa1\x15\x53\x20\x6f\xa1\x97\xa9\xc4\x72\xbb\xd5\xae\xa2\x4a\x61\x06\x81\x08\xe1\x1f\x61\xb5\x15\xe2\xe6\x4f\x92\x0e\x4e\x7d\xfb\xf0\x95\x10\x05\xd0\x7e\x65\x22\xa3\xe1\x7b\x02\xc3\x0d\x9c\xae\x0a\x78\x5b\xeb\xf6\xd3\xa7\x3f\xd9\x92\xe9\x89\x31\x9d\xd0\x46\x8a\xba\xba\x92\x4c\x48\xa6\x77\x6f\x19\x67\x65\x3d\x88\x85\x1d\x6b\x39\xc4\x1b\x0d\x5b\xa0\x85\xde\x66\x5b\xc8\x06\x59\x1c\x4b\xc6\xad\xb4\xc1\xa9\x11\x97\x70\xf4\xdd\x0e\x39\xd8\x0b\xba\x9f\xc0\xc6\x30\x19\xdf\xbc\x00\xa9\x07\x65\x1a\x93\x38\xa3\x2f\x86\xd9\x22\x29\xfe\x64\x03\xdc\xc4\x50\x70\xac\xc2\x2c\xff\x20\xef\xc3\x83\xa5\x10\x59\x51\x47\x28\xc4\xfc\xe1\xa2\x91\x70\xea\x68\xd7\x0a\xbe\xab\x57\xd0\xb8\x8e\x6f\x7e\xce\xf9\x37\x42\x5e\xde\x0c\x8e\x43\x5c\x4f\xb7\x20\x4d\xc8\xeb\x27\xcf\x43\x1b\x51\x48\x01\x0b\x92\x1d\x26\xa7\x8b\x61\xb7\x12\xb8\xaa\xe7\x35\x7a\xd7\x0d\x39\x85\xde\x45\xad\x39\xdd\x3b\x67\x26\x55\xef\xc7\xf6\x9c\xe8\x9d\x0c\x0f\x4c\xef\xd2\x03\xbd\xa7\x2e\x17\x99\x19\xc4\x69\x8b\x85\x84\x0d\x53\x5a\x46\x56\x86\x80\xf9\x4a\xa8\x84\x62\x5a\x1c\x71\xab\xa6\x9b\x89\xf7\x84\x0d\xc5\xf3\x3f\x70\xc2\xf3\xd7\x3b\xa5\x69\xf2\x02\x9c\x49\xa6\x59\x46\x8b\xcb\x3c\xef\x77\x59\xc3\x73\xc7\x5a\xe1\x25\xa7\xc5\x4e\xb3\xac\xa7\xf6\xd8\x8d\xb8\x31\x12\x53\x3d\xc7\x16\x1b\xc4\xc8\xe2\x11\x9f\xdf\xa1\x95\x21\xee\xf9\x3f\x87\xcd\x34\xfd\x93\xd1\x37\x13\x5f\xb8\x57\xa5\x3e\xf8\x3b\x1a\xeb\xb1\x6d\x60\x20\x8a\xe5\x90\x51\xe9\xb3\x67\x90\xc7\xbc\x28\x58\xb0\x92\x0d\x2f\x7c\x64\x72\x7d\x25\x21\x7a\xee\x88\xf8\xec\x0d\x3e\xdc\xfd\xb8\x72\x1d\x82\x92\x7e\x32\x5e\x8c\xd0\x52\xd4\xb6\x72\xe1\xde\x1a\x8b\x04\xe4\x5e\x45\x3e\xc8\x27\x6f\x05\xf6\x70\xd7\xe2\x82\x6c\xb5\xae\xd4\xc5\xf9\xf9\x4d\xbd\x02\xc9\x41\x83\x5a\x32\x71\x9e\x8b\x4c\x9d\x67\x82\x67\x50\x69\xfc\xcf\x9a\x6d\x6a\x89\x81\xf1\x79\x49\x39\xdd\xc0\xc2\x3d\x76\xd1\x90\x5f\x34\x9a\x3e\x7f\x16\x5d\x2a\x23\x31\xbd\x7b\xeb\xee\x73\x69\xfc\x83\x7b\xfc\xa1\xce\xed\xca\x71\x94\xce\x65\xf3\x92\xd5\xeb\x35\x69\xe8\x33\x45\x44\xc9\xb4\xc9\xe0\xd6\x42\x12\xba\xb7\xd2\x70\xf5\x9f\x69\x93\xaa\xd2\xba\xd0\x58\xdc\x70\xd6\x81\xaf\xef\xd9\x37\x2d\xe1\x53\x55\xb0\x8c\xe9\x62\xb7\xcf\x8d\xcf\xec\x0b\xb4\x77\x4c\x85\x99\xb5\xb5\xb0\x66\x4f\x74\x1c\xe5\x85\xcf\x8a\x31\xf5\xfd\x62\x2d\x26\x7a\x5a\x41\x56\x9b\x85\xf5\x85\xe0\x1a\x3e\x0d\xba\xc0\xce\xf0\x5f\xbb\xeb\x89\xc0\x1f\x54\x83\xc5\x70\x75\x11\x59\x73\x4c\xec\x8f\x71\x24\x38\xf5\xae\x24\xbb\x65\x05\x6c\xe0\x95\xca\xa8\xed\x21\x25\x61\x7a\x9e\x5d\x06\xee\x46\xb3\x91\xa2\x50\xe4\x6e\x0b\xb8\xe5\x17\x35\x9c\x64\xa0\xc2\xed\x89\x8c\x72\xb2\xa1\x8c\xdb\xdd\xab\x2a\x4f\x14\xcb\x12\x1c\x11\x27\x15\x95\xc0\xb5\x27\xe4\x1a\x0c\x66\x71\x09\xd2\xcc\x99\x84\xcc\xd8\x5d\xc3\x4f\xf3\x56\xe9\x4f\x1c\xee\x7e\x32\x4f\x51\x64\x5d\xd0\x8d\x45\x09\xad\x5c\xb7\x3f\xdc\x23\xb7\x50\x34\x67\x1d\x7b\x56\x82\x8a\x60\x8a\x68\x59\x03\xa1\xc5\x1d\xdd\x85\x85\xbf\x73\x70\x99\x16\x6d\xa6\x2e\xc8\xf3\x53\x1c\x5c\xaa\x48\x43\x3b\x27\xbf\x39\xc5\xf7\x61\x5f\x5c\x5e\xfd\x74\xfd\x9f\xd7\x3f\x5d\xbe\x7c\xfb\xfa\x5d\xdc\x4c\x63\x89\x48\x46\x2b\xba\x62\x05\x8b\x79\xac\xde\x6b\xaa\xed\x9b\x70\x9a\xe6\xf9\x79\x2e\x45\x65\xe5\xf0\xd5\xaa\x46\x96\x48\x55\xfd\x65\xcb\x73\x18\xf9\x9d\x27\xf1\x1d\xc9\xce\x83\x36\x92\x72\xdd\x2c\xa4\x61\x43\x6a\x54\x28\x6b\xae\x59\x19\x44\x29\xa5\xb4\xa8\x69\x1e\x6d\xd5\x74\xbb\xfa\xf8\x86\x6d\x9b\xe5\xc8\x9d\x09\x65\xd8\x6e\x60\xe1\xc9\xee\xf6\xad\x5f\x72\xf5\xfe\xfa\xf5\x9f\x0e\x46\x63\x57\xc5\x2b\x82\x89\xa5\xdd\x94\xc2\xae\x19\xf2\x64\xed\x7c\x80\x52\xdc\xfe\x23\xe9\x67\x34\xa8\x68\x7c\x5c\xd0\xc4\xba\x0a\xac\x79\xdb\x3d\xf0\xd6\xfd\xa4\x44\xb0\xe2\x95\x75\x47\xa0\x62\x28\x9f\xd6\x5d\xfb\x09\x8a\x3d\x1a\x73\x2b\xd7\xcc\x82\xfe\x3a\x6f\x86\x48\x21\x46\xbd\xe2\x56\x28\xbd\xec\xcc\xe7\x35\x2d\x54\x70\xf2\x8d\x7b\x26\xe3\x5c\xdf\x9a\xc0\x26\x49\x3b\xcd\xd5\x24\x07\x2e\x3c\x6e\xc5\x3c\x05\xc1\x5b\x52\x64\xc4\x46\x49\x5a\x98\x5c\x38\x28\xca\x1a\x21\x32\xd0\x76\x5e\xe8\xf2\xbc\x63\x62\xca\xcb\x78\xd5\x3c\x31\xfe\x55\x82\x5a\x35\x9f\x24\x38\x70\x4c\xfb\xb8\x69\x8d\x88\x0e\x9a\x63\x27\xad\xa2\x7a\xab\xa2\xaf\xef\x96\x54\xdd\x40\x6e\x2f\x74\xeb\xa0\x8b\xe7\xec\x93\x1a\xd6\x3e\x1a\xf9\xd7\x40\x75\x2d\x01\xd7\xb9\x58\xb0\xb5\x02\x9f\xcb\xc5\x07\x2d\x32\x37\x8c\x0c\xef\x79\xb1\xfb\x20\x84\xfe\x86\x15\x60\xb1\xa7\x49\x03\xf8\xa3\x8b\x14\x2c\xfe\xac\x51\x95\x59\xea\x28\xd2\x5d\xa0\x72\xd0\x14\xd7\x0d\xe9\xd1\x95\xc5\x0c\xd8\x3d\x0d\x51\xd6\xfc\x52\x7d\x2b\x45\x1d\x74\x76\xbd\x05\xf2\xdb\xd7\x2f\x71\xde\xd4\x76\x55\x07\xae\xe5\xce\x02\x7c\x5d\xa3\xa4\x11\x30\x32\x4f\x5d\x6c\xf1\xbd\xb1\x9f\x03\x8b\x31\x71\x4c\xcd\x15\xe8\x25\x79\x4b\x77\x84\x16\x4a\xf8\xe0\x25\x32\xf5\xaf\x44\x7e\xdd\x8d\x3d\x97\x88\x56\xb1\xb7\x91\x95\xd0\x5b\x72\x70\x01\xf6\x86\xfb\xf7\x85\xb3\x81\xa6\x8f\xdc\xea\x83\x31\xde\x23\xab\xe9\x0d\x28\x52\x49\xc8\x20\x07\x9e\x05\x47\xa7\x55\xe0\xfb\xed\xbf\x46\x47\x30\x86\xc1\xc7\x11\x7c\x27\xb8\x31\xcb\x34\xb4\x3a\xcf\x59\xe6\xd0\x6f\x0e\x4a\xb6\x37\x49\x84\xe1\xb8\xb8\x8c\x22\x18\xc7\x18\x65\x6c\xfe\x4b\x0b\xd4\x91\xb5\x83\xa5\x7f\x57\xaf\xa0\x00\x6d\x83\xce\x5b\x5b\x2e\xb4\x9f\x15\x61\x25\xdd\x00\xa1\xda\x0f\x78\x6c\xbe\x02\x57\xb5\xf4\x1f\xb4\xd1\x24\x17\x60\x7b\x66\x8e\xb5\xef\x5f\xbf\x24\x5f\x91\x13\xc3\xdb\x29\x0e\xe3\x9a\xb2\x22\xf6\x7d\x71\xa5\xa9\x3c\x94\x95\xad\x3d\x69\x14\x01\x6d\x8e\x08\x69\xa7\xd4\x19\xe1\x82\xa8\x3a\xe2\xfb\x9c\x6c\x26\x12\xf6\x01\x76\x05\xd2\x0c\x2a\xa6\xfb\x3d\xd3\x0d\x99\x68\x98\xe7\xe9\xa6\xbb\x37\xd1\x30\xd5\x87\x30\xdd\x44\xc7\xf2\xbd\xea\xd7\x4c\xfd\xd1\xf3\x2b\xdf\x3f\xa0\x5f\x69\x2f\xd5\xc6\x46\xbb\x52\x5b\x43\x2c\x41\xd3\x9c\x6a\xea\xfc\x8d\xbf\x20\xec\x75\xd3\x87\x34\x36\x74\xe1\x70\x7c\x6c\x48\xa3\x43\x17\x9e\x4c\x7f\x53\x6f\xa4\xe0\x0d\xe3\xf5\xa7\xf7\xd5\x60\x3f\xd7\x1f\xbd\xb1\xbf\x7e\x85\xb7\xe1\x10\xa3\x1d\xa2\x92\x2d\xac\x24\xf7\xf9\x53\xb4\xaa\x68\x8f\xd7\x9d\xa1\x3c\x0b\x84\x26\x38\x5d\x69\x61\xf1\x08\x66\x05\xa6\x3c\x17\xe1\xb7\x48\x0e\x99\x6b\x3e\x82\xb5\x67\x68\x8a\x71\xfc\x1a\xe7\x7b\x4a\x3a\x59\xc0\x2d\x14\xc9\x29\xd3\x1b\x73\xb5\x09\x60\xbc\x76\xf1\x76\x07\xe5\x40\xb7\xbf\x07\x15\xc5\x73\x9a\x34\xcb\x48\x45\x52\x88\x22\xd8\xf9\xec\xc9\xf0\x41\x14\x60\x21\x77\x5e\x08\x73\xfb\x67\x97\x01\x2f\x4a\x95\x01\xa3\xe8\x8e\x0c\x98\x57\x7c\x6e\x19\xea\xc8\xca\xd1\x93\xc1\x2c\x33\x5d\x19\xd0\xe7\x7f\x5e\x19\x46\x53\xe4\x3b\xc6\x73\x71\xa7\xa6\xba\xca\x1f\xed\x6d\x7e\x5e\x67\xc6\x6d\x68\xc6\x37\xaa\xed\x2e\x69\x51\x24\x95\xa8\x86\xfc\xa5\xaf\xc4\xe2\xdb\x70\x98\x71\xf5\xfc\x0e\x7a\xd0\x20\xd1\x15\x18\xfd\xdb\xea\xfb\x17\x1c\x7e\xa7\xf8\xb4\x4d\xa9\xe8\x0b\x69\xe8\x68\x46\x8b\xeb\x0a\xb2\x64\xa3\xfc\xf6\xed\xf5\x65\xf7\x56\x63\xa2\xf6\xa5\x31\x23\x89\x39\x4f\x68\x5e\x32\xc4\xbe\x46\xcd\xf2\xce\xb6\xda\xc9\x89\x6f\x03\x6c\x98\xde\xd6\xab\x65\x26\xca\x56\x47\x60\xa1\xd8\x46\x9d\x3b\xab\x5a\x18\xce\xe3\xe8\x41\xc6\x0b\x7c\x8d\xcf\x1b\xfd\xfe\xf3\x86\x8e\xb9\xac\xe1\x1e\x15\x8e\x90\xbf\x38\x24\xd7\xb5\x01\xfb\xa2\xbf\xa3\x25\x58\x48\xaf\x4b\xe9\x9b\xef\x67\xd0\xa2\xda\xd2\x05\x3a\xff\x28\x69\x63\x2c\xcc\xc1\x71\xb7\x02\x5f\xd5\x35\x8f\xb3\x0d\x7f\x97\xca\xd8\x0c\x1f\x59\x70\xb3\xc4\x70\x12\x25\xdb\xae\x1f\xdc\xdb\x69\xf5\xad\xc5\xc8\x7d\x0f\x8b\x41\xb5\xb9\x37\x80\x8c\xf6\xdb\xc3\x13\x15\xeb\x70\xe8\x6c\x18\x1c\xd1\xfd\xe8\x17\xd5\xbe\x74\xdd\x37\xf9\xc6\x24\x95\x63\xde\xe1\x6e\x32\xbe\xc4\x3b\xd7\xc1\x3c\x24\x2a\xcc\x61\x8e\x32\x9c\x8b\x98\x4b\xba\xf9\xc8\xc8\x14\x1d\xc9\x55\x12\xc3\xce\xe8\x43\x1e\xd8\x4b\x93\x87\xf7\xd4\xf6\x88\xda\xae\xc9\xe4\xc3\x26\x3a\xc2\xec\xa0\xf9\x7e\x68\x1b\xd4\x43\xda\xea\x7d\xda\xab\x0f\x09\xf0\x19\xc4\x59\x3d\x32\xea\x67\x5d\x17\x85\x71\x64\xef\x6f\x41\x4a\x96\xf7\x26\x6a\x50\x10\x9c\x06\x57\x75\x51\x5c\x89\x82\x65\x3d\x94\xcb\xf8\x7d\xd7\x90\x49\xe8\x83\x1c\x02\xad\x98\x51\x48\x71\xbf\x39\x52\x88\xcd\x9b\xa1\x0c\x28\x06\xe3\x0b\xe7\xd3\xa5\xe0\x46\xd9\x8c\xf7\xc6\x3b\x16\xb6\xd0\x0d\x0c\xb7\x14\x12\x60\xc5\xc7\x80\x9c\xa4\x28\x41\x6f\xa1\x3e\x0a\xa6\xda\xbc\x41\x70\x24\xc8\x73\x04\x87\x20\x6f\x59\x06\x6f\xad\x1a\x8f\x61\xaf\x88\xbe\xf0\xf5\xe0\xa0\x98\xe3\xbd\x42\x78\x86\xa3\x39\xf4\xe7\x7d\x33\x6a\xbd\x53\x5d\xad\xa5\x4e\xea\xa3\x26\x34\x17\x39\x84\xbe\x04\xf0\x78\xb0\xff\xd0\xfb\xc9\x5f\x1a\x90\x73\x06\xe5\xcd\xa0\xbc\x19\x94\x17\x38\x66\x50\xde\xc4\xd3\x33\x28\x6f\xe8\x98\x41\x79\x33\x28\x6f\x06\xe5\x05\x89\x7f\x79\xa0\xb3\x19\x94\x37\x83\xf2\xbc\xa8\x33\x28\x6f\x06\xe5\x91\x19\x94\x37\x83\xf2\x66\x50\x5e\xef\x98\x41\x79\x33\x28\x6f\x06\xe5\xcd\xa0\xbc\xe6\x98\x41\x79\x7f\x97\xf3\x7d\x06\xe5\x0d\xc9\x30\x83\xf2\x1e\x4d\x86\x19\x94\x37\x24\xe8\x0c\xca\x9b\x41\x79\x33\x28\xef\x8b\x05\x86\xcd\xa0\xbc\x19\x94\x37\x83\xf2\x66\x50\xde\xaf\x03\x94\xf7\xc8\xf8\xbb\x4a\xe4\x97\x9f\xe3\x73\x9b\x55\xcf\x56\x0f\xe9\x76\xfb\xbc\x03\xd1\xcd\x56\x14\x39\x7e\x32\xd7\xd9\x97\x6f\x6a\x13\xaa\xb5\x64\xab\x5a\x0f\x74\x77\x8c\x0d\x66\xa2\x2c\x45\xbb\x91\xe1\x43\xb3\x25\xb1\x51\x1e\x2d\x2e\x3a\xee\xc0\x7d\xe2\x9d\x5c\x03\x0c\x37\x6f\x5a\xac\x62\xda\xe9\x4b\xa4\xee\x13\xd3\x62\x6d\x13\x51\xbb\xb2\x1e\x36\x4a\x63\x01\xce\x3a\x5c\xf4\xed\xa8\xe7\xe9\xa5\x9d\xc2\x66\x21\xa9\x2b\x8f\x56\x28\xec\xa7\xf7\x0e\x63\xeb\x83\xb0\x73\xd0\xfe\x19\xb7\xdb\xc2\x2c\xc9\xb5\x28\x81\xdc\x8a\xa2\x2e\xad\xf0\x0e\x2b\xd3\x29\x22\x6a\x41\xb2\x2d\x6e\x37\x86\x91\xe9\x9d\x21\x1b\xda\x96\x41\x38\x54\x86\x27\x89\x3e\xd1\xdc\xd2\xc0\x93\x2a\x91\x5f\x90\xff\xe6\xe4\xb9\x6d\x7b\x88\x3b\x6c\xe5\x7e\xfb\xfa\x65\x38\x9e\x5d\xd9\x27\x7f\x73\x8d\xea\x22\xbf\xb1\x77\x2a\xd0\x1b\x96\x93\x95\x75\x3a\xc6\x7b\x9e\x70\xb8\xb3\xa5\x7b\xb3\xf6\xda\x6f\x0f\x0f\x07\x75\xe8\x1b\x2d\x8b\xbe\x6c\xd8\x30\xe9\x1e\x73\x4a\xbe\xb6\xcf\xa9\x40\xba\xf8\xd0\x3c\x2b\xbc\x15\xfb\xfb\x0f\xcf\xdc\x66\x6f\xf2\x6e\x21\xef\x16\x8b\xc5\xc2\xc8\xe9\x8b\x9a\x03\x85\x59\xdc\xf1\x4b\xe4\x6c\x1d\xee\x37\x37\xda\x46\xdb\xde\xb3\xa2\xfc\x66\x3f\x56\x8a\xe5\xd0\x77\xb2\xc7\x8a\x49\x23\xdf\xa6\x8b\x36\x25\xee\xd1\x90\x68\x16\xe5\x41\x81\x27\x37\x23\x0e\x17\xb4\x70\x79\xe7\xe1\x4b\x3b\x93\x16\x56\xd7\xf9\xdb\x6f\x6f\x33\x9c\xdd\x3e\xc0\xa8\x45\x1a\x11\x8f\xd2\x84\x78\xf8\x06\xc4\xf1\xcd\x07\xdb\x64\x08\x4e\xfa\xa9\x8d\x87\x56\x83\x21\x50\x3f\x48\x69\x3a\x84\xca\xd3\x21\xe7\x7c\x9c\x89\x8e\xc4\xb3\x47\xc6\x7e\xf1\x2e\x43\xb4\xc3\x70\x8f\xee\x42\xdc\x49\xdc\xa7\xb3\x60\xc6\x67\x90\x68\xe2\x98\x4d\x6a\x29\x3c\x42\x3b\xe1\xb3\xb8\x95\xf1\x8e\xc2\xf4\x6e\x42\x4a\x79\xec\x3e\xad\x04\xcf\xc1\x20\xe1\xa9\x6d\x84\xbf\xb3\x45\x66\x14\xfe\x1d\x6b\x23\x1c\xdd\x42\x48\x43\xd5\x1d\x8f\x0d\x89\xb4\x0d\x8e\x6e\x19\x3c\x32\xcf\xb1\x36\xc1\xd1\x2d\x82\x47\xe6\x39\xd6\x16\x38\xba\x25\xf0\xa8\x3c\xc7\xd1\xd2\xad\x94\x0a\xe3\xdd\x71\xff\x76\xd9\x6c\xd7\x88\x29\x98\x3a\x6c\x92\xae\x99\x54\x0d\x8e\x18\x97\xbb\x40\x1e\xd2\x75\x3d\xb8\xb9\xb3\x4f\xca\x7b\x0d\xd7\x67\x66\x9e\xb3\x92\xca\x9d\x09\xb6\xc3\x1e\xa8\xe3\x30\xb9\xf0\x2c\xfa\x48\x85\x22\x88\x14\xa1\xef\xbb\xb8\x62\x23\x08\xc9\xf1\x3e\xf5\x58\x97\x3a\x06\x6f\x54\x3b\x95\xe9\xe1\x77\xb1\xba\xb8\x75\x7b\x1d\x96\x0a\x5a\x5b\x68\x36\x2f\x9a\xe5\x9e\x12\xb6\x42\x0c\xd3\xc1\xb8\x10\x93\xe1\x2b\x91\x2b\x9b\xc1\xd5\xdc\x58\x85\x90\xba\x45\xe3\xc4\x25\xb0\xbd\xf5\x67\xb8\xc2\x5f\xe2\x3e\xe0\x3e\x2a\x2d\x68\xcd\x87\x77\x7a\x8b\x68\x79\x40\x58\xb7\x93\x8e\xdd\x28\x51\x72\x28\x48\x45\x25\x2d\x41\x83\x74\x2b\x6c\x28\xac\x1c\xef\xb4\xf0\x68\xd1\xb6\xc3\xcc\x3b\x57\x07\xa7\x9e\x2c\x6e\xb4\x1c\x7a\x34\x49\xab\xc9\xe1\x12\x98\xc8\xc0\x0f\x7e\xef\xdc\x07\xe4\x20\xbe\xab\xce\x02\xf5\x13\x38\x15\x5e\xbd\x93\x0a\x8d\xc3\xd3\x60\xbc\x37\xf9\x78\x7d\xc9\x58\x4f\xd2\xcc\x10\xac\x20\xb5\x5d\x53\x4a\xe8\xea\x5d\xd0\x83\x37\x26\x1f\xb0\xd4\x3d\x36\x4d\xd2\x9b\x91\x8f\xd1\x88\x7c\xf8\x26\xe4\x71\x0d\xc8\xf8\xb6\xf6\x0f\xdf\x7c\x7c\x84\xc6\x63\x4a\x33\x61\x7c\xf3\x9c\x49\xcd\xc6\xc7\x68\x34\x1e\xd5\x64\xf4\xba\x0c\x52\x9d\xa6\xe3\x87\xd1\x65\x52\xf3\xf0\xf8\xc6\x21\x11\x61\xd0\xd7\x51\x4d\xc3\xa6\xd4\x10\x24\xfb\x58\x0d\xc3\xe3\x3c\x67\x34\xc2\x3e\xc6\x7b\xa2\x8d\x85\xe7\xd7\x51\x4d\xc2\xf8\x96\xde\x0f\xd1\x20\x3c\x3e\x55\x08\x9e\x92\x50\x15\x2c\xa3\x2f\x86\xde\x25\x39\xee\x73\x10\xee\xc5\xfc\xcb\x2c\x1b\xa2\x19\xfd\x24\x44\xb8\xdd\x47\x26\xbd\x47\x7c\x9f\x94\xca\x76\x5f\xa6\x7f\x60\x22\x14\x81\x1e\xd5\x5a\xb5\x4c\xa4\x8e\xa1\x16\x05\xc8\x61\xbd\x75\x6b\xe6\x6b\x72\x10\x16\xd9\x1d\x4b\x5b\xf7\x1f\x5a\x5e\x20\xc2\xef\xb9\xb2\x4a\xe4\xf6\xe5\x94\x8f\x0d\x2d\x9c\x3f\x5a\xd3\x6c\xeb\x92\x4b\x7b\xc6\xc4\xff\x83\x1b\x5b\x1a\x0f\xa6\xad\xab\xde\xef\x84\x09\x44\x4b\x56\x15\x40\x7e\xdf\xec\xf1\x7b\x06\xeb\x35\x64\xfa\x0f\xa4\x56\x8c\x6f\xdc\xfb\xf5\xc1\xad\x42\x9b\x5d\x76\x7f\xef\xff\xf7\x87\xfe\xec\x8a\x07\x4e\xf6\x79\x09\x69\xce\x2b\xbc\x90\xb0\x56\x6b\x02\x9c\x58\x96\x86\x51\x03\xf2\x1a\xdf\x87\xd3\xee\x29\x8b\x17\x62\xc0\xdb\x26\xa1\x96\xe4\xc7\x2d\xf0\xf6\x40\xba\x2f\x04\xc4\xf7\x32\xa7\x12\xc8\x3b\x71\x6d\x06\xa3\x2e\xe0\x8c\x5c\x49\x58\x83\xdc\xff\x82\xee\xed\x9d\x78\xf5\x09\xb2\x5a\x07\x80\x14\x23\xd3\xea\x26\xb4\x07\x5e\x47\x49\x6e\x53\xd1\xbd\x6a\x70\x1b\x54\xdf\xc3\xd9\x9b\xa2\xaf\xaf\x84\x5c\xa1\x70\x3a\x0c\x68\xeb\x06\x76\xaa\xd9\x96\xfd\xc6\x3e\x13\xbb\xe5\x67\x63\xbb\x74\xfb\xdd\xdd\xed\xbe\xec\xbf\xf3\xef\x5c\x95\x2b\xc6\x2d\x63\xf6\x81\x7e\x28\xf1\x99\x7e\xc7\xe6\x20\x06\xca\x5c\x84\x2c\x1d\xa3\xd8\xd0\x97\x40\x06\xb4\xfb\xde\x9b\x7b\xf3\x9e\xa7\x4d\xb7\x77\xcf\x14\x91\x60\xdf\xbb\xc6\x96\xac\x0b\x0b\xec\x97\x02\x02\x4c\xe3\x06\x7a\xcd\xd3\xed\x7b\x90\x6e\xb7\x7a\x63\x2d\xaf\x7e\xae\x69\xd1\x8d\x34\xdc\x4f\xf6\xa2\x00\xd5\xce\x56\xdc\xe6\xa6\x3b\x56\xe4\x19\x95\xf6\xab\x0a\x76\x8a\x13\x25\x5c\x91\x0d\x3d\x4b\x46\x79\xe3\x3e\x22\x0a\xc6\x91\x57\x2e\xb3\xa6\x52\xb3\xac\x2e\xa8\x24\x66\x2e\x6e\x84\xdc\x1d\xa5\xfb\xbd\x41\x5e\x43\x26\x78\x9e\x52\xee\xf8\x78\x78\x4f\x7b\x34\xb4\x6d\xcb\x33\x91\x63\x8c\xcc\x4a\x88\x04\x37\xad\xe9\x70\x62\x37\xf9\xf5\xd6\x29\xd6\xde\xa7\x34\x93\xb6\xf5\xd9\x08\x34\xda\x00\xcd\x26\xb8\x61\x1b\x8c\x60\x4e\x5b\x9e\xb9\x99\x95\x4b\xf2\xc7\x9d\x6f\x61\x9f\xb9\xb0\x87\x07\xdf\xe3\x42\x88\x80\xe3\xcf\x4d\x0e\x4b\xb1\x35\xcd\xd7\x42\xc2\x2d\x48\x72\x92\x0b\xec\x85\xc2\x2d\xcb\xf4\x69\xc8\xf4\xfe\x0b\xa4\x40\x23\xe3\xb0\xa1\x9a\xdd\x36\x9b\xa2\xfb\x3c\x5c\x3b\x90\x04\x55\xe4\x2b\x72\x82\xc4\x08\x2b\x4b\xc8\x19\xd5\x50\x04\x37\xa3\xf6\xdf\x99\x89\xbc\x1f\x79\xff\xfa\x60\xa4\x18\x34\x50\x08\xea\x38\x43\x1b\xcb\x1e\x78\xc2\x66\x39\x14\x21\x75\x39\x3f\xd7\xde\x52\xde\xce\xc1\x4e\xbb\xae\xd9\xbe\xdb\x3b\xc2\x91\xaf\x8c\xfc\xd5\xd8\x1a\x25\x12\x36\x38\x8f\xec\x1c\x39\x62\x16\x8d\xc6\xa5\x87\x35\xa4\xa1\xc0\x68\xb1\xdf\x07\xb9\xf3\xab\xdb\xee\xb2\xf3\xdb\x7e\x03\xc5\xce\xcf\x07\x1f\x66\xeb\x9c\xdb\x7f\xc4\xac\xf3\xf3\xc0\x5a\xb1\xe8\xc4\xcc\x9d\x13\xdd\xd0\xf7\x49\x54\x03\x07\x3f\xb9\xcd\x3e\x2f\xc8\xed\x73\x4c\x3a\x9e\xef\x7f\x43\x97\x63\x4b\x78\x9d\xd3\x6e\xef\xd9\xfc\x02\xb1\x0b\xf6\x07\x2d\x24\xdd\x80\xfb\xe5\xff\x03\x00\x00\xff\xff\xa9\xfc\xaa\xd9\xe7\x0e\x01\x00")

func installerKubedbCom_kubedboperatorsYamlBytes() ([]byte, error) {
	return bindataRead(
		_installerKubedbCom_kubedboperatorsYaml,
		"installer.kubedb.com_kubedboperators.yaml",
	)
}

func installerKubedbCom_kubedboperatorsYaml() (*asset, error) {
	bytes, err := installerKubedbCom_kubedboperatorsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "installer.kubedb.com_kubedboperators.yaml", size: 69351, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"installer.kubedb.com_kubedbcatalogs.yaml":  installerKubedbCom_kubedbcatalogsYaml,
	"installer.kubedb.com_kubedboperators.yaml": installerKubedbCom_kubedboperatorsYaml,
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
	"installer.kubedb.com_kubedbcatalogs.yaml":  {installerKubedbCom_kubedbcatalogsYaml, map[string]*bintree{}},
	"installer.kubedb.com_kubedboperators.yaml": {installerKubedbCom_kubedboperatorsYaml, map[string]*bintree{}},
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
