// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package resource

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/aelsabbahy/goss/system"
	"github.com/aelsabbahy/goss/util"
)

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type AddrMap map[string]*Addr

func (r AddrMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Addr, error) {
	sysres := sys.NewAddr(sr, sys, config)
	res, err := NewAddr(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r AddrMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Addr, system.Addr, bool) {
	sysres := sys.NewAddr(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewAddr(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *AddrMap) UnmarshalJSON(data []byte) error {
	resEmpty := Addr{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Addr
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *AddrMap) UnmarshalYAML(data []byte) error {
func (r *AddrMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := Addr{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Addr
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type CommandMap map[string]*Command

func (r CommandMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Command, error) {
	sysres := sys.NewCommand(sr, sys, config)
	res, err := NewCommand(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r CommandMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Command, system.Command, bool) {
	sysres := sys.NewCommand(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewCommand(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *CommandMap) UnmarshalJSON(data []byte) error {
	resEmpty := Command{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Command
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *CommandMap) UnmarshalYAML(data []byte) error {
func (r *CommandMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := Command{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Command
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type DNSMap map[string]*DNS

func (r DNSMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*DNS, error) {
	sysres := sys.NewDNS(sr, sys, config)
	res, err := NewDNS(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r DNSMap) AppendSysResourceIfExists(sr string, sys *system.System) (*DNS, system.DNS, bool) {
	sysres := sys.NewDNS(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewDNS(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *DNSMap) UnmarshalJSON(data []byte) error {
	resEmpty := DNS{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*DNS
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *DNSMap) UnmarshalYAML(data []byte) error {
func (r *DNSMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := DNS{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*DNS
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type FileMap map[string]*File

func (r FileMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*File, error) {
	sysres := sys.NewFile(sr, sys, config)
	res, err := NewFile(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r FileMap) AppendSysResourceIfExists(sr string, sys *system.System) (*File, system.File, bool) {
	sysres := sys.NewFile(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewFile(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *FileMap) UnmarshalJSON(data []byte) error {
	resEmpty := File{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*File
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *FileMap) UnmarshalYAML(data []byte) error {
func (r *FileMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := File{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*File
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type GossfileMap map[string]*Gossfile

func (r GossfileMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Gossfile, error) {
	sysres := sys.NewGossfile(sr, sys, config)
	res, err := NewGossfile(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r GossfileMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Gossfile, system.Gossfile, bool) {
	sysres := sys.NewGossfile(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewGossfile(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *GossfileMap) UnmarshalJSON(data []byte) error {
	resEmpty := Gossfile{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Gossfile
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *GossfileMap) UnmarshalYAML(data []byte) error {
func (r *GossfileMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := Gossfile{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Gossfile
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type GroupMap map[string]*Group

func (r GroupMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Group, error) {
	sysres := sys.NewGroup(sr, sys, config)
	res, err := NewGroup(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r GroupMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Group, system.Group, bool) {
	sysres := sys.NewGroup(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewGroup(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *GroupMap) UnmarshalJSON(data []byte) error {
	resEmpty := Group{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Group
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *GroupMap) UnmarshalYAML(data []byte) error {
func (r *GroupMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := Group{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Group
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type PackageMap map[string]*Package

func (r PackageMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Package, error) {
	sysres := sys.NewPackage(sr, sys, config)
	res, err := NewPackage(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r PackageMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Package, system.Package, bool) {
	sysres := sys.NewPackage(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewPackage(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *PackageMap) UnmarshalJSON(data []byte) error {
	resEmpty := Package{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Package
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *PackageMap) UnmarshalYAML(data []byte) error {
func (r *PackageMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := Package{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Package
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type PortMap map[string]*Port

func (r PortMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Port, error) {
	sysres := sys.NewPort(sr, sys, config)
	res, err := NewPort(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r PortMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Port, system.Port, bool) {
	sysres := sys.NewPort(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewPort(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *PortMap) UnmarshalJSON(data []byte) error {
	resEmpty := Port{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Port
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *PortMap) UnmarshalYAML(data []byte) error {
func (r *PortMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := Port{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Port
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type ProcessMap map[string]*Process

func (r ProcessMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Process, error) {
	sysres := sys.NewProcess(sr, sys, config)
	res, err := NewProcess(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r ProcessMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Process, system.Process, bool) {
	sysres := sys.NewProcess(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewProcess(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *ProcessMap) UnmarshalJSON(data []byte) error {
	resEmpty := Process{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Process
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *ProcessMap) UnmarshalYAML(data []byte) error {
func (r *ProcessMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := Process{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Process
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type ServiceMap map[string]*Service

func (r ServiceMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Service, error) {
	sysres := sys.NewService(sr, sys, config)
	res, err := NewService(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r ServiceMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Service, system.Service, bool) {
	sysres := sys.NewService(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewService(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *ServiceMap) UnmarshalJSON(data []byte) error {
	resEmpty := Service{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Service
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *ServiceMap) UnmarshalYAML(data []byte) error {
func (r *ServiceMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := Service{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Service
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type UserMap map[string]*User

func (r UserMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*User, error) {
	sysres := sys.NewUser(sr, sys, config)
	res, err := NewUser(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r UserMap) AppendSysResourceIfExists(sr string, sys *system.System) (*User, system.User, bool) {
	sysres := sys.NewUser(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewUser(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *UserMap) UnmarshalJSON(data []byte) error {
	resEmpty := User{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*User
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *UserMap) UnmarshalYAML(data []byte) error {
func (r *UserMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := User{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*User
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type KernelParamMap map[string]*KernelParam

func (r KernelParamMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*KernelParam, error) {
	sysres := sys.NewKernelParam(sr, sys, config)
	res, err := NewKernelParam(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r KernelParamMap) AppendSysResourceIfExists(sr string, sys *system.System) (*KernelParam, system.KernelParam, bool) {
	sysres := sys.NewKernelParam(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewKernelParam(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *KernelParamMap) UnmarshalJSON(data []byte) error {
	resEmpty := KernelParam{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*KernelParam
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *KernelParamMap) UnmarshalYAML(data []byte) error {
func (r *KernelParamMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := KernelParam{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*KernelParam
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type MountMap map[string]*Mount

func (r MountMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Mount, error) {
	sysres := sys.NewMount(sr, sys, config)
	res, err := NewMount(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r MountMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Mount, system.Mount, bool) {
	sysres := sys.NewMount(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewMount(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *MountMap) UnmarshalJSON(data []byte) error {
	resEmpty := Mount{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Mount
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *MountMap) UnmarshalYAML(data []byte) error {
func (r *MountMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := Mount{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Mount
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type InterfaceMap map[string]*Interface

func (r InterfaceMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*Interface, error) {
	sysres := sys.NewInterface(sr, sys, config)
	res, err := NewInterface(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r InterfaceMap) AppendSysResourceIfExists(sr string, sys *system.System) (*Interface, system.Interface, bool) {
	sysres := sys.NewInterface(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewInterface(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *InterfaceMap) UnmarshalJSON(data []byte) error {
	resEmpty := Interface{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Interface
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *InterfaceMap) UnmarshalYAML(data []byte) error {
func (r *InterfaceMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := Interface{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*Interface
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//go:generate sed -i -e "/^\\/\\/ +build genny/d" resource_list.go
//go:generate goimports -w resource_list.go resource_list.go

type HTTPMap map[string]*HTTP

func (r HTTPMap) AppendSysResource(sr string, sys *system.System, config util.Config) (*HTTP, error) {
	sysres := sys.NewHTTP(sr, sys, config)
	res, err := NewHTTP(sysres, config)
	if err != nil {
		return nil, err
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, nil
}

func (r HTTPMap) AppendSysResourceIfExists(sr string, sys *system.System) (*HTTP, system.HTTP, bool) {
	sysres := sys.NewHTTP(sr, sys, util.Config{})
	// FIXME: Do we want to be silent about errors?
	res, _ := NewHTTP(sysres, util.Config{})
	if e, _ := sysres.Exists(); e != true {
		return res, sysres, false
	}
	if old_res, ok := r[res.ID()]; ok {
		res.Title = old_res.Title
		res.Meta = old_res.Meta
	}
	r[res.ID()] = res
	return res, sysres, true
}

func (r *HTTPMap) UnmarshalJSON(data []byte) error {
	resEmpty := HTTP{}
	validAttrs, err := validAttrs(resEmpty, "json")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := json.Unmarshal(data, &validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*HTTP
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}

//func (r *HTTPMap) UnmarshalYAML(data []byte) error {
func (r *HTTPMap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	resEmpty := HTTP{}
	validAttrs, err := validAttrs(resEmpty, "yaml")
	if err != nil {
		return err
	}
	var validate map[string]map[string]interface{}
	if err := unmarshal(&validate); err != nil {
		return err
	}

	typ := reflect.TypeOf(resEmpty)
	typs := strings.Split(typ.String(), ".")[1]
	for id, v := range validate {
		for k, _ := range v {
			if !validAttrs[k] {
				return fmt.Errorf("Invalid Attribute for %s:%s: %s", typs, id, k)
			}
		}
	}

	var tmp map[string]*HTTP
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	for id, res := range tmp {
		if res == nil {
			return fmt.Errorf("Could not parse resource %s:%s", typs, id)
		}
		res.SetID(id)
	}

	*r = tmp

	return nil
}
