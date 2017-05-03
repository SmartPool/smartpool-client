package storage

import (
	"encoding/gob"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

var SmartPoolDir = getSmartPoolDir()

func getSmartPoolDir() string {
	result, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(result, ".smartpool")
}

func getFile(id string) string {
	return filepath.Join(SmartPoolDir, id)
}

func getRigDataFile() string {
	return filepath.Join(SmartPoolDir, "rig_data")
}

type GobFileStorage struct {
	mu             sync.Mutex
	registeredType map[string]bool
}

func getName(value interface{}) string {
	// Default to printed representation for unnamed types
	rt := reflect.TypeOf(value)
	name := rt.String()

	// But for named types (or pointers to them), qualify with import path (but see inner comment).
	// Dereference one pointer looking for a named type.
	star := ""
	if rt.Name() == "" {
		if pt := rt; pt.Kind() == reflect.Ptr {
			star = "*"
			// NOTE: The following line should be rt = pt.Elem() to implement
			// what the comment above claims, but fixing it would break compatibility
			// with existing gobs.
			//
			// Given package p imported as "full/p" with these definitions:
			//     package p
			//     type T1 struct { ... }
			// this table shows the intended and actual strings used by gob to
			// name the types:
			//
			// Type      Correct string     Actual string
			//
			// T1        full/p.T1          full/p.T1
			// *T1       *full/p.T1         *p.T1
			//
			// The missing full path cannot be fixed without breaking existing gob decoders.
			rt = pt
		}
	}
	if rt.Name() != "" {
		if rt.PkgPath() == "" {
			name = star + rt.Name()
		} else {
			name = star + rt.PkgPath() + "." + rt.Name()
		}
	}
	return name
}

func (gfs *GobFileStorage) register(value interface{}) {
	gfs.mu.Lock()
	defer gfs.mu.Unlock()
	name := getName(value)
	if _, found := gfs.registeredType[name]; !found {
		gob.Register(value)
		gfs.registeredType[name] = true
	}
}

func (gfs *GobFileStorage) Persist(data interface{}, id string) error {
	err := os.MkdirAll(SmartPoolDir, 0766)
	if err != nil {
		return err
	}
	f, err := os.Create(getFile(id))
	if err != nil {
		return err
	}
	defer f.Close()
	gfs.register(data)
	enc := gob.NewEncoder(f)
	return enc.Encode(data)
}

func (gfs *GobFileStorage) Load(data interface{}, id string) (interface{}, error) {
	f, err := os.Open(getFile(id))
	if err != nil {
		return data, err
	}
	defer f.Close()
	gfs.register(data)
	dec := gob.NewDecoder(f)
	err = dec.Decode(data)
	return data, err
}

func NewGobFileStorage() *GobFileStorage {
	return &GobFileStorage{
		sync.Mutex{},
		map[string]bool{},
	}
}
