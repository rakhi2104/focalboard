// Code generated by go-bindata. DO NOT EDIT.
// sources:
// postgres_files/000001_init.down.sql
// postgres_files/000001_init.up.sql
// postgres_files/000002_system_settings_table.down.sql
// postgres_files/000002_system_settings_table.up.sql
// postgres_files/000003_blocks_rootid.down.sql
// postgres_files/000003_blocks_rootid.up.sql
// postgres_files/000004_auth_table.down.sql
// postgres_files/000004_auth_table.up.sql
// postgres_files/000005_blocks_modifiedby.down.sql
// postgres_files/000005_blocks_modifiedby.up.sql
// postgres_files/000006_sharing_table.down.sql
// postgres_files/000006_sharing_table.up.sql
// postgres_files/000007_workspaces_table.down.sql
// postgres_files/000007_workspaces_table.up.sql
// postgres_files/000008_teams.down.sql
// postgres_files/000008_teams.up.sql
package postgres

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

var __000001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x2e\xb6\xe6\x02\x04\x00\x00\xff\xff\x45\xbe\x01\x0f\x13\x00\x00\x00")

func _000001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initDownSql,
		"000001_init.down.sql",
	)
}

func _000001_initDownSql() (*asset, error) {
	bytes, err := _000001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.down.sql", size: 19, mode: os.FileMode(420), modTime: time.Unix(1603074564, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x31\x4f\xc3\x30\x10\x85\xe7\xf8\x57\xbc\x31\x91\xb2\x21\xb1\x30\xb9\xe5\x0a\x86\xc4\xa9\x9c\x2b\xb4\x2c\x55\x88\x0f\x61\x11\x4a\x14\x9b\x81\x7f\x8f\xda\x21\x43\xba\xdd\x7d\xba\xef\x9e\xde\xda\x91\x66\x02\xeb\x55\x45\x30\x1b\xd8\x86\x41\x7b\xd3\x72\x8b\xf7\xe1\xa7\xff\x8a\xc8\x55\x16\x3c\x5e\xb4\x5b\x3f\x6a\x97\xdf\xdc\x16\xa5\xca\xc2\x29\xca\x94\x8e\x5d\x02\x9b\x9a\x5a\xd6\xf5\x96\xdf\x2e\xae\xdd\x55\x15\xee\x69\xa3\x77\x15\xc3\x36\xaf\xf9\xf9\x7c\xec\x26\x39\xa5\xe3\xd5\x9b\xd8\x7f\xca\x77\x87\x95\x79\x30\x96\x4b\x95\xa5\xbf\x51\xc0\xb4\xbf\xcc\x21\x0d\xf3\xf2\x11\x64\xf0\x11\x4f\x6d\x63\x4b\x95\xf5\x93\x74\x49\xce\xe9\xb3\xf9\x3b\xfa\x25\xf2\x32\xc8\x02\x6d\x9d\xa9\xb5\x3b\xe0\x99\x0e\xc8\x83\x2f\x31\xf7\x28\x54\x71\xa7\xfe\x03\x00\x00\xff\xff\xa3\xc9\xa2\x70\x0c\x01\x00\x00")

func _000001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initUpSql,
		"000001_init.up.sql",
	)
}

func _000001_initUpSql() (*asset, error) {
	bytes, err := _000001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.up.sql", size: 268, mode: os.FileMode(420), modTime: time.Unix(1607029670, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_system_settings_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\xae\x2c\x2e\x49\xcd\x8d\x2f\x4e\x2d\x29\xc9\xcc\x4b\x2f\xb6\xe6\x02\x04\x00\x00\xff\xff\x8b\x60\xbf\x1e\x1c\x00\x00\x00")

func _000002_system_settings_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_system_settings_tableDownSql,
		"000002_system_settings_table.down.sql",
	)
}

func _000002_system_settings_tableDownSql() (*asset, error) {
	bytes, err := _000002_system_settings_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_system_settings_table.down.sql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1603229117, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_system_settings_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x28\xae\x2c\x2e\x49\xcd\x8d\x2f\x4e\x2d\x29\xc9\xcc\x4b\x2f\x56\xd0\xe0\xe2\xcc\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x34\x30\xd0\xd4\xe1\xe2\x2c\x4b\xcc\x29\x4d\x55\x08\x71\x8d\x08\xd1\xe1\xe2\x0c\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\x54\xd0\xc8\x4c\xd1\xe4\xd2\xb4\xe6\x02\x04\x00\x00\xff\xff\x17\x95\xca\x5b\x61\x00\x00\x00")

func _000002_system_settings_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_system_settings_tableUpSql,
		"000002_system_settings_table.up.sql",
	)
}

func _000002_system_settings_tableUpSql() (*asset, error) {
	bytes, err := _000002_system_settings_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_system_settings_table.up.sql", size: 97, mode: os.FileMode(420), modTime: time.Unix(1603229117, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000003_blocks_rootidDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x2e\xe6\x72\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xca\xcf\x2f\x89\xcf\x4c\xb1\xe6\x02\x04\x00\x00\xff\xff\x94\x1c\x55\xb9\x28\x00\x00\x00")

func _000003_blocks_rootidDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_blocks_rootidDownSql,
		"000003_blocks_rootid.down.sql",
	)
}

func _000003_blocks_rootidDownSql() (*asset, error) {
	bytes, err := _000003_blocks_rootidDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_blocks_rootid.down.sql", size: 40, mode: os.FileMode(420), modTime: time.Unix(1610349080, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000003_blocks_rootidUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x2e\xe6\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xca\xcf\x2f\x89\xcf\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\xce\x60\x70\x4e\x33\x00\x00\x00")

func _000003_blocks_rootidUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_blocks_rootidUpSql,
		"000003_blocks_rootid.up.sql",
	)
}

func _000003_blocks_rootidUpSql() (*asset, error) {
	bytes, err := _000003_blocks_rootidUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_blocks_rootid.up.sql", size: 51, mode: os.FileMode(420), modTime: time.Unix(1610349080, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000004_auth_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\xb6\xe6\x42\x12\x29\x4e\x2d\x2e\xce\xcc\xcf\x2b\xb6\xe6\x02\x04\x00\x00\xff\xff\xa5\xe0\x77\xaa\x27\x00\x00\x00")

func _000004_auth_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_auth_tableDownSql,
		"000004_auth_table.down.sql",
	)
}

func _000004_auth_tableDownSql() (*asset, error) {
	bytes, err := _000004_auth_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_auth_table.down.sql", size: 39, mode: os.FileMode(420), modTime: time.Unix(1611791102, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000004_auth_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x4f\x4b\x03\x31\x10\xc5\xcf\x9b\x4f\x31\xc7\x5d\xe8\xa1\x16\x7a\xf2\x94\x96\xa8\xf1\xcf\x56\xb2\x41\xec\x69\x19\x36\x23\x06\xbb\x7f\xc8\xa4\xfa\xf5\x25\x08\x16\x9a\xd5\x4b\xe7\xf8\x7b\xc3\xbc\x37\x6f\x6b\x94\xb4\x0a\xac\xdc\x3c\x2a\xd0\x37\x50\xef\x2c\xa8\x57\xdd\xd8\x06\x8e\x4c\x81\xa1\x14\x85\x77\xf0\x22\xcd\xf6\x4e\x9a\xf2\x6a\xb9\xac\x16\xa2\x48\xd2\x80\x3d\x9d\x73\xea\xd1\x1f\x7e\xe1\x6a\xbd\x4e\x70\x42\xe6\xaf\x31\x64\x47\xfa\x37\x6c\x99\xba\x40\xf1\x5c\xc1\x63\x7c\x6f\x99\xc2\xa7\xef\x4e\x16\xab\x93\xe4\x30\x62\xe6\x12\xc6\x89\xe1\x67\xee\x9b\x5d\xbd\x10\x45\x17\x08\x23\xb5\x18\x13\xdb\xe8\x5b\x5d\xdb\x94\x7d\x72\x33\xd4\xd1\x81\x72\xfa\x6c\xf4\x93\x34\x7b\x78\x50\x7b\x28\xbd\xab\x44\x75\x2d\xc4\x3f\x95\x31\x31\xfb\x71\xf8\xa3\xb5\x38\x7e\xd0\x30\x57\x65\x9b\xef\x5e\xf8\xce\x5c\xf0\xef\x00\x00\x00\xff\xff\x16\x8f\x58\x39\xeb\x01\x00\x00")

func _000004_auth_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_auth_tableUpSql,
		"000004_auth_table.up.sql",
	)
}

func _000004_auth_tableUpSql() (*asset, error) {
	bytes, err := _000004_auth_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_auth_table.up.sql", size: 491, mode: os.FileMode(420), modTime: time.Unix(1611791102, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000005_blocks_modifiedbyDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x2e\xe6\x72\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\xcd\x4f\xc9\x4c\xcb\x4c\x4d\x89\x4f\xaa\xb4\xe6\x02\x04\x00\x00\xff\xff\xe4\x42\x8b\x2e\x2c\x00\x00\x00")

func _000005_blocks_modifiedbyDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000005_blocks_modifiedbyDownSql,
		"000005_blocks_modifiedby.down.sql",
	)
}

func _000005_blocks_modifiedbyDownSql() (*asset, error) {
	bytes, err := _000005_blocks_modifiedbyDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000005_blocks_modifiedby.down.sql", size: 44, mode: os.FileMode(420), modTime: time.Unix(1611791102, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000005_blocks_modifiedbyUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x2e\xe6\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\xcd\x4f\xc9\x4c\xcb\x4c\x4d\x89\x4f\xaa\x54\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\xea\xb0\x5a\x65\x37\x00\x00\x00")

func _000005_blocks_modifiedbyUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000005_blocks_modifiedbyUpSql,
		"000005_blocks_modifiedby.up.sql",
	)
}

func _000005_blocks_modifiedbyUpSql() (*asset, error) {
	bytes, err := _000005_blocks_modifiedbyUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000005_blocks_modifiedby.up.sql", size: 55, mode: os.FileMode(420), modTime: time.Unix(1611791102, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000006_sharing_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\xce\x48\x2c\xca\xcc\x4b\xb7\xe6\x02\x04\x00\x00\xff\xff\xdd\x4c\x62\xe8\x14\x00\x00\x00")

func _000006_sharing_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000006_sharing_tableDownSql,
		"000006_sharing_table.down.sql",
	)
}

func _000006_sharing_tableDownSql() (*asset, error) {
	bytes, err := _000006_sharing_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000006_sharing_table.down.sql", size: 20, mode: os.FileMode(420), modTime: time.Unix(1611791102, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000006_sharing_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcc\xb1\xca\xc2\x30\x10\x00\xe0\x39\xf7\x14\x37\xb6\xd0\xa1\x3f\x3f\xb8\x38\x5d\xcb\xa9\xc1\xda\x4a\x1a\xc4\x4e\x25\xe5\xa2\x06\xb5\x15\x8d\x83\x6f\x2f\x2e\x0e\xee\x1f\x5f\x69\x98\x2c\xa3\xa5\xa2\x62\xd4\x0b\xac\x1b\x8b\xbc\xd7\xad\x6d\xf1\x71\x72\xf7\x30\x1e\x31\x01\x15\x04\x77\x64\xca\x15\x99\xe4\x7f\x96\x66\xa0\xfc\xe8\x86\x8b\x17\x2c\x9a\xa6\x62\xaa\x33\x50\x71\x3a\xfb\xf1\xab\xfe\xf2\xfc\xc3\xae\x93\x84\x43\xf0\xd2\x0f\xaf\x9f\xe0\x79\x13\x17\x7d\xef\x22\x16\x7a\xa9\x6b\x9b\x81\xda\x1a\xbd\x21\xd3\xe1\x9a\x3b\x4c\x82\xa4\x90\xce\xe1\x1d\x00\x00\xff\xff\x6c\x91\x98\xb6\x9f\x00\x00\x00")

func _000006_sharing_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000006_sharing_tableUpSql,
		"000006_sharing_table.up.sql",
	)
}

func _000006_sharing_tableUpSql() (*asset, error) {
	bytes, err := _000006_sharing_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000006_sharing_table.up.sql", size: 159, mode: os.FileMode(420), modTime: time.Unix(1611791102, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000007_workspaces_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\xcf\x2f\xca\x2e\x2e\x48\x4c\x4e\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\xc4\x05\x92\x8e\x17\x00\x00\x00")

func _000007_workspaces_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000007_workspaces_tableDownSql,
		"000007_workspaces_table.down.sql",
	)
}

func _000007_workspaces_tableDownSql() (*asset, error) {
	bytes, err := _000007_workspaces_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000007_workspaces_table.down.sql", size: 23, mode: os.FileMode(420), modTime: time.Unix(1611791102, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000007_workspaces_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcc\xc1\x6a\x83\x30\x00\x06\xe0\x73\xf2\x14\xff\xd1\x40\x0e\x8e\xc1\x2e\x3b\x45\xc9\xb6\x6c\x2e\x96\x98\x96\x7a\x12\xdb\xa4\x12\xa4\x2a\x4d\xa4\xf4\xed\x0b\x3d\xf4\xd0\xf3\x07\x5f\x69\xa4\xb0\x12\x56\x14\x95\x84\xfa\x82\xae\x2d\xe4\x5e\x35\xb6\xc1\x75\xbe\x8c\x71\xe9\x8f\x3e\x22\xa3\x24\x38\xec\x84\x29\x7f\x84\xc9\xde\x3f\x18\xa7\x24\x86\x61\x5a\x97\x2e\xcd\xa3\x9f\x9e\xf4\x96\xe7\xec\x71\xe8\x6d\x55\x71\x0a\x00\xd1\xa7\x14\xa6\x21\xe2\xb7\xa9\x35\xa7\xe4\x3c\xbb\x70\x0a\xde\x75\x87\xdb\xcb\xb8\x2e\xae\x4f\xbe\xeb\x13\x0a\xf5\xad\xb4\xe5\x94\x6c\x8c\xfa\x17\xa6\xc5\x9f\x6c\x91\x05\xc7\x28\xfb\xa4\xf7\x00\x00\x00\xff\xff\x3b\x70\x91\x2c\xb3\x00\x00\x00")

func _000007_workspaces_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000007_workspaces_tableUpSql,
		"000007_workspaces_table.up.sql",
	)
}

func _000007_workspaces_tableUpSql() (*asset, error) {
	bytes, err := _000007_workspaces_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000007_workspaces_table.up.sql", size: 179, mode: os.FileMode(420), modTime: time.Unix(1611791102, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000008_teamsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x2e\xe6\x72\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xcf\x2f\xca\x2e\x2e\x48\x4c\x4e\x8d\xcf\x4c\xb1\xe6\xe2\x42\x56\x5d\x9c\x91\x58\x94\x99\x97\x4e\xb4\xf2\xd4\xe2\xe2\xcc\xfc\x3c\x54\xe3\x13\x4b\x4b\x32\xe2\x8b\x53\x8b\xca\x32\x93\x53\xad\xb9\x00\x01\x00\x00\xff\xff\xdd\xff\x41\x9f\x8c\x00\x00\x00")

func _000008_teamsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000008_teamsDownSql,
		"000008_teams.down.sql",
	)
}

func _000008_teamsDownSql() (*asset, error) {
	bytes, err := _000008_teamsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000008_teams.down.sql", size: 140, mode: os.FileMode(420), modTime: time.Unix(1616709037, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000008_teamsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x2e\xe6\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xcf\x2f\xca\x2e\x2e\x48\x4c\x4e\x8d\xcf\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd3\xb4\xe6\xe2\x42\xd6\x58\x9c\x91\x58\x94\x99\x97\x4e\x8e\xce\xd4\xe2\xe2\xcc\xfc\x3c\x14\x4b\x13\x4b\x4b\x32\xe2\x8b\x53\x8b\xca\x32\x93\x53\xe1\x5a\x8d\x0c\x34\xad\xb9\x00\x01\x00\x00\xff\xff\xba\x55\x30\xd8\xad\x00\x00\x00")

func _000008_teamsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000008_teamsUpSql,
		"000008_teams.up.sql",
	)
}

func _000008_teamsUpSql() (*asset, error) {
	bytes, err := _000008_teamsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000008_teams.up.sql", size: 173, mode: os.FileMode(420), modTime: time.Unix(1616709035, 0)}
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
	"000001_init.down.sql": _000001_initDownSql,
	"000001_init.up.sql": _000001_initUpSql,
	"000002_system_settings_table.down.sql": _000002_system_settings_tableDownSql,
	"000002_system_settings_table.up.sql": _000002_system_settings_tableUpSql,
	"000003_blocks_rootid.down.sql": _000003_blocks_rootidDownSql,
	"000003_blocks_rootid.up.sql": _000003_blocks_rootidUpSql,
	"000004_auth_table.down.sql": _000004_auth_tableDownSql,
	"000004_auth_table.up.sql": _000004_auth_tableUpSql,
	"000005_blocks_modifiedby.down.sql": _000005_blocks_modifiedbyDownSql,
	"000005_blocks_modifiedby.up.sql": _000005_blocks_modifiedbyUpSql,
	"000006_sharing_table.down.sql": _000006_sharing_tableDownSql,
	"000006_sharing_table.up.sql": _000006_sharing_tableUpSql,
	"000007_workspaces_table.down.sql": _000007_workspaces_tableDownSql,
	"000007_workspaces_table.up.sql": _000007_workspaces_tableUpSql,
	"000008_teams.down.sql": _000008_teamsDownSql,
	"000008_teams.up.sql": _000008_teamsUpSql,
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
	"000001_init.down.sql": &bintree{_000001_initDownSql, map[string]*bintree{}},
	"000001_init.up.sql": &bintree{_000001_initUpSql, map[string]*bintree{}},
	"000002_system_settings_table.down.sql": &bintree{_000002_system_settings_tableDownSql, map[string]*bintree{}},
	"000002_system_settings_table.up.sql": &bintree{_000002_system_settings_tableUpSql, map[string]*bintree{}},
	"000003_blocks_rootid.down.sql": &bintree{_000003_blocks_rootidDownSql, map[string]*bintree{}},
	"000003_blocks_rootid.up.sql": &bintree{_000003_blocks_rootidUpSql, map[string]*bintree{}},
	"000004_auth_table.down.sql": &bintree{_000004_auth_tableDownSql, map[string]*bintree{}},
	"000004_auth_table.up.sql": &bintree{_000004_auth_tableUpSql, map[string]*bintree{}},
	"000005_blocks_modifiedby.down.sql": &bintree{_000005_blocks_modifiedbyDownSql, map[string]*bintree{}},
	"000005_blocks_modifiedby.up.sql": &bintree{_000005_blocks_modifiedbyUpSql, map[string]*bintree{}},
	"000006_sharing_table.down.sql": &bintree{_000006_sharing_tableDownSql, map[string]*bintree{}},
	"000006_sharing_table.up.sql": &bintree{_000006_sharing_tableUpSql, map[string]*bintree{}},
	"000007_workspaces_table.down.sql": &bintree{_000007_workspaces_tableDownSql, map[string]*bintree{}},
	"000007_workspaces_table.up.sql": &bintree{_000007_workspaces_tableUpSql, map[string]*bintree{}},
	"000008_teams.down.sql": &bintree{_000008_teamsDownSql, map[string]*bintree{}},
	"000008_teams.up.sql": &bintree{_000008_teamsUpSql, map[string]*bintree{}},
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

