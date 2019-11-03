// Package swagger Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// apidocs.swagger.json
package swagger

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
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

// Mode return file modify time
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

var _apidocsSwaggerJson = []byte(`{
  "swagger": "2.0",
  "info": {
    "title": "openboard",
    "description": "auth semantics:\n\n Add{T}(s)    (Add{T}(s)Req)    returns {T}(s)Resp // POST\n Ovr{T}(s)    (Ovr{T}(s)Req)    returns {T}(s)Resp // PUT\n Mod{T}(s)    (Mod{T}(s)Req)    returns {T}(s)Resp // PATCH\n Get{T}       (Get{T}Req)       returns {T}Resp    // GET\n Fnd{T}s      (Fnd{T}sReq)      returns {T}sResp   // GET\n Rmv{T}(s)    (Rmv{T}(s)Req)    returns EmptyResp  // DELETE\n Unr{T}       (Unr{T}Req)       returns {T}Resp    // PATCH",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/auth": {
      "post": {
        "operationId": "AddAuth",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbAuthResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbAddAuthReq"
            }
          }
        ],
        "tags": [
          "Auth"
        ]
      }
    },
    "/auth/{token}": {
      "delete": {
        "operationId": "RmvAuth",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbRmvAuthResp"
            }
          }
        },
        "parameters": [
          {
            "name": "token",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Auth"
        ]
      }
    },
    "/roles": {
      "get": {
        "operationId": "FndRoles",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbRolesResp"
            }
          }
        },
        "parameters": [
          {
            "name": "roleIds",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "integer",
              "format": "int64"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "roleNames",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "limit",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "lapse",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          }
        ],
        "tags": [
          "UserSvc"
        ]
      }
    },
    "/user": {
      "post": {
        "operationId": "AddUser",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbUserResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbAddUserReq"
            }
          }
        ],
        "tags": [
          "UserSvc"
        ]
      }
    },
    "/user/{id}": {
      "delete": {
        "operationId": "RmvUser",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbRmvUserResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int64"
          }
        ],
        "tags": [
          "UserSvc"
        ]
      },
      "put": {
        "operationId": "OvrUser",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbUserResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbOvrUserReq"
            }
          }
        ],
        "tags": [
          "UserSvc"
        ]
      }
    },
    "/users": {
      "get": {
        "operationId": "FndUsers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbUsersResp"
            }
          }
        },
        "parameters": [
          {
            "name": "roleIds",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "integer",
              "format": "int64"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "email",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "emailHold",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          },
          {
            "name": "altmail",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "altmailHold",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          },
          {
            "name": "createdFirst",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "date-time"
          },
          {
            "name": "createdFinal",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "date-time"
          },
          {
            "name": "createdDesc",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          }
        ],
        "tags": [
          "UserSvc"
        ]
      }
    },
    "/voucher": {
      "post": {
        "operationId": "AddVoucher",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbAddVoucherResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbAddVoucherReq"
            }
          }
        ],
        "tags": [
          "Auth"
        ]
      }
    }
  },
  "definitions": {
    "pbAddAuthReq": {
      "type": "object",
      "properties": {
        "username": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "pbAddUserReq": {
      "type": "object",
      "properties": {
        "username": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "emailHold": {
          "type": "boolean",
          "format": "boolean"
        },
        "altmail": {
          "type": "string"
        },
        "altmailHold": {
          "type": "boolean",
          "format": "boolean"
        },
        "fullName": {
          "type": "string"
        },
        "avatar": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "roleIds": {
          "type": "array",
          "items": {
            "type": "integer",
            "format": "int64"
          }
        }
      }
    },
    "pbAddVoucherReq": {
      "type": "object",
      "properties": {
        "notify": {
          "type": "boolean",
          "format": "boolean"
        },
        "email": {
          "type": "string"
        },
        "password": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "pbAddVoucherResp": {
      "type": "object"
    },
    "pbAuthResp": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "pbOvrUserReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "req": {
          "$ref": "#/definitions/pbAddUserReq"
        }
      }
    },
    "pbRmvAuthResp": {
      "type": "object"
    },
    "pbRmvUserResp": {
      "type": "object"
    },
    "pbRoleResp": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "pbRolesResp": {
      "type": "object",
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbRoleResp"
          }
        },
        "total": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "pbUser": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "username": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "emailHold": {
          "type": "boolean",
          "format": "boolean"
        },
        "altmail": {
          "type": "string"
        },
        "altmailHold": {
          "type": "boolean",
          "format": "boolean"
        },
        "fullName": {
          "type": "string"
        },
        "avatar": {
          "type": "string"
        },
        "roles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbRoleResp"
          }
        },
        "lastLogin": {
          "type": "string",
          "format": "date-time"
        },
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "updated": {
          "type": "string",
          "format": "date-time"
        },
        "deleted": {
          "type": "string",
          "format": "date-time"
        },
        "blocked": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "pbUserResp": {
      "type": "object",
      "properties": {
        "item": {
          "$ref": "#/definitions/pbUser"
        }
      }
    },
    "pbUsersResp": {
      "type": "object",
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbUser"
          }
        },
        "total": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  }
}
`)

func apidocsSwaggerJsonBytes() ([]byte, error) {
	return _apidocsSwaggerJson, nil
}

func apidocsSwaggerJson() (*asset, error) {
	bytes, err := apidocsSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "apidocs.swagger.json", size: 11369, mode: os.FileMode(436), modTime: time.Unix(1572816172, 0)}
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
	"apidocs.swagger.json": apidocsSwaggerJson,
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
	"apidocs.swagger.json": &bintree{apidocsSwaggerJson, map[string]*bintree{}},
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
