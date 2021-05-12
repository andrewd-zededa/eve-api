// Copyright(c) 2017-2018 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.12.2
// source: config/netconfig.proto

package config

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NetworkConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type NetworkType `protobuf:"varint,5,opt,name=type,proto3,enum=org.lfedge.eve.config.NetworkType" json:"type,omitempty"`
	// network ip specification
	Ip  *Ipspec               `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip,omitempty"`
	Dns []*ZnetStaticDNSEntry `protobuf:"bytes,7,rep,name=dns,proto3" json:"dns,omitempty"`
	// enterprise proxy
	EntProxy *ProxyConfig `protobuf:"bytes,8,opt,name=entProxy,proto3" json:"entProxy,omitempty"`
	// wireless specification
	Wireless *WirelessConfig `protobuf:"bytes,10,opt,name=wireless,proto3" json:"wireless,omitempty"`
}

func (x *NetworkConfig) Reset() {
	*x = NetworkConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkConfig) ProtoMessage() {}

func (x *NetworkConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_netconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkConfig.ProtoReflect.Descriptor instead.
func (*NetworkConfig) Descriptor() ([]byte, []int) {
	return file_config_netconfig_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NetworkConfig) GetType() NetworkType {
	if x != nil {
		return x.Type
	}
	return NetworkType_NETWORKTYPENOOP
}

func (x *NetworkConfig) GetIp() *Ipspec {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *NetworkConfig) GetDns() []*ZnetStaticDNSEntry {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *NetworkConfig) GetEntProxy() *ProxyConfig {
	if x != nil {
		return x.EntProxy
	}
	return nil
}

func (x *NetworkConfig) GetWireless() *WirelessConfig {
	if x != nil {
		return x.Wireless
	}
	return nil
}

type NetworkAdapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`           // name which we report in metrics and status
	NetworkId string `protobuf:"bytes,3,opt,name=networkId,proto3" json:"networkId,omitempty"` // UUID of NetworkInstance object
	Addr      string `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`           // Static IP address; could be IPv4 EID
	Hostname  string `protobuf:"bytes,5,opt,name=hostname,proto3" json:"hostname,omitempty"`   // Not currently used
	// more configuration for getting addr/EID
	CryptoEid     string `protobuf:"bytes,10,opt,name=cryptoEid,proto3" json:"cryptoEid,omitempty"`
	Lispsignature string `protobuf:"bytes,6,opt,name=lispsignature,proto3" json:"lispsignature,omitempty"`
	Pemcert       []byte `protobuf:"bytes,7,opt,name=pemcert,proto3" json:"pemcert,omitempty"`
	Pemprivatekey []byte `protobuf:"bytes,8,opt,name=pemprivatekey,proto3" json:"pemprivatekey,omitempty"`
	// Used in case of P2V, where we want to specify a macAddress
	// to vif, that is simulated towards app
	MacAddress string `protobuf:"bytes,9,opt,name=macAddress,proto3" json:"macAddress,omitempty"`
	// firewall
	Acls []*ACE `protobuf:"bytes,40,rep,name=acls,proto3" json:"acls,omitempty"`
	// access port vlan id
	// app interface with access vlan id of zero will be treated as trunk port
	// valid vlan id range: 2 - 4093
	// vlan id 1 is implicitly used by linux bridges
	AccessVlanId uint32 `protobuf:"varint,41,opt,name=access_vlan_id,json=accessVlanId,proto3" json:"access_vlan_id,omitempty"`
}

func (x *NetworkAdapter) Reset() {
	*x = NetworkAdapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkAdapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkAdapter) ProtoMessage() {}

func (x *NetworkAdapter) ProtoReflect() protoreflect.Message {
	mi := &file_config_netconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkAdapter.ProtoReflect.Descriptor instead.
func (*NetworkAdapter) Descriptor() ([]byte, []int) {
	return file_config_netconfig_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkAdapter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NetworkAdapter) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *NetworkAdapter) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *NetworkAdapter) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *NetworkAdapter) GetCryptoEid() string {
	if x != nil {
		return x.CryptoEid
	}
	return ""
}

func (x *NetworkAdapter) GetLispsignature() string {
	if x != nil {
		return x.Lispsignature
	}
	return ""
}

func (x *NetworkAdapter) GetPemcert() []byte {
	if x != nil {
		return x.Pemcert
	}
	return nil
}

func (x *NetworkAdapter) GetPemprivatekey() []byte {
	if x != nil {
		return x.Pemprivatekey
	}
	return nil
}

func (x *NetworkAdapter) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *NetworkAdapter) GetAcls() []*ACE {
	if x != nil {
		return x.Acls
	}
	return nil
}

func (x *NetworkAdapter) GetAccessVlanId() uint32 {
	if x != nil {
		return x.AccessVlanId
	}
	return 0
}

type WirelessConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        WirelessType      `protobuf:"varint,1,opt,name=type,proto3,enum=org.lfedge.eve.config.WirelessType" json:"type,omitempty"` // either LTE or Wifi
	CellularCfg []*CellularConfig `protobuf:"bytes,5,rep,name=cellularCfg,proto3" json:"cellularCfg,omitempty"`                            // Cellular config
	WifiCfg     []*WifiConfig     `protobuf:"bytes,10,rep,name=wifiCfg,proto3" json:"wifiCfg,omitempty"`                                   // Wifi, can be multiple APs on a single wlan, e.g. one for 2.5Ghz, other 5Ghz SSIDs
}

func (x *WirelessConfig) Reset() {
	*x = WirelessConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WirelessConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WirelessConfig) ProtoMessage() {}

func (x *WirelessConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_netconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WirelessConfig.ProtoReflect.Descriptor instead.
func (*WirelessConfig) Descriptor() ([]byte, []int) {
	return file_config_netconfig_proto_rawDescGZIP(), []int{2}
}

func (x *WirelessConfig) GetType() WirelessType {
	if x != nil {
		return x.Type
	}
	return WirelessType_TypeNOOP
}

func (x *WirelessConfig) GetCellularCfg() []*CellularConfig {
	if x != nil {
		return x.CellularCfg
	}
	return nil
}

func (x *WirelessConfig) GetWifiCfg() []*WifiConfig {
	if x != nil {
		return x.WifiCfg
	}
	return nil
}

type CellularConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	APN string `protobuf:"bytes,1,opt,name=APN,proto3" json:"APN,omitempty"` // APN string
}

func (x *CellularConfig) Reset() {
	*x = CellularConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netconfig_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CellularConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CellularConfig) ProtoMessage() {}

func (x *CellularConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_netconfig_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CellularConfig.ProtoReflect.Descriptor instead.
func (*CellularConfig) Descriptor() ([]byte, []int) {
	return file_config_netconfig_proto_rawDescGZIP(), []int{3}
}

func (x *CellularConfig) GetAPN() string {
	if x != nil {
		return x.APN
	}
	return ""
}

type WifiConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WifiSSID  string        `protobuf:"bytes,1,opt,name=wifiSSID,proto3" json:"wifiSSID,omitempty"`                                             // SSID for WIFI
	KeyScheme WiFiKeyScheme `protobuf:"varint,2,opt,name=keyScheme,proto3,enum=org.lfedge.eve.config.WiFiKeyScheme" json:"keyScheme,omitempty"` // key management scheme, WPA-PSK, WPS-EPA, etc
	// to be deprecated, use cipherData instead
	Identity string `protobuf:"bytes,5,opt,name=identity,proto3" json:"identity,omitempty"` // WPA2 Enterprise user identity/username
	// to be deprecated, use cipherData instead
	Password   string                 `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`     // a string of hashed psk or password
	Crypto     *WifiConfigCryptoblock `protobuf:"bytes,20,opt,name=crypto,proto3" json:"crypto,omitempty"`         // encrypted block
	Priority   int32                  `protobuf:"varint,25,opt,name=priority,proto3" json:"priority,omitempty"`    // priority of connection, default is 0
	CipherData *CipherBlock           `protobuf:"bytes,30,opt,name=cipherData,proto3" json:"cipherData,omitempty"` // contains encrypted credential information
}

func (x *WifiConfig) Reset() {
	*x = WifiConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netconfig_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiConfig) ProtoMessage() {}

func (x *WifiConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_netconfig_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiConfig.ProtoReflect.Descriptor instead.
func (*WifiConfig) Descriptor() ([]byte, []int) {
	return file_config_netconfig_proto_rawDescGZIP(), []int{4}
}

func (x *WifiConfig) GetWifiSSID() string {
	if x != nil {
		return x.WifiSSID
	}
	return ""
}

func (x *WifiConfig) GetKeyScheme() WiFiKeyScheme {
	if x != nil {
		return x.KeyScheme
	}
	return WiFiKeyScheme_SchemeNOOP
}

func (x *WifiConfig) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *WifiConfig) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *WifiConfig) GetCrypto() *WifiConfigCryptoblock {
	if x != nil {
		return x.Crypto
	}
	return nil
}

func (x *WifiConfig) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *WifiConfig) GetCipherData() *CipherBlock {
	if x != nil {
		return x.CipherData
	}
	return nil
}

type WifiConfigCryptoblock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,11,opt,name=identity,proto3" json:"identity,omitempty"` // encrypted username if not empty
	Password string `protobuf:"bytes,12,opt,name=password,proto3" json:"password,omitempty"` // encrypted Password if not empty
}

func (x *WifiConfigCryptoblock) Reset() {
	*x = WifiConfigCryptoblock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netconfig_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiConfigCryptoblock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiConfigCryptoblock) ProtoMessage() {}

func (x *WifiConfigCryptoblock) ProtoReflect() protoreflect.Message {
	mi := &file_config_netconfig_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiConfigCryptoblock.ProtoReflect.Descriptor instead.
func (*WifiConfigCryptoblock) Descriptor() ([]byte, []int) {
	return file_config_netconfig_proto_rawDescGZIP(), []int{4, 0}
}

func (x *WifiConfigCryptoblock) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *WifiConfigCryptoblock) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_config_netconfig_proto protoreflect.FileDescriptor

var file_config_netconfig_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6e, 0x65, 0x74, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a,
	0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x66, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x6e, 0x65, 0x74, 0x63, 0x6d, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc6, 0x02, 0x0a, 0x0d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64,
	0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x69, 0x70,
	0x73, 0x70, 0x65, 0x63, 0x52, 0x02, 0x69, 0x70, 0x12, 0x3b, 0x0a, 0x03, 0x64, 0x6e, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64,
	0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x5a, 0x6e,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x44, 0x4e, 0x53, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x64, 0x6e, 0x73, 0x12, 0x3e, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x41, 0x0a, 0x08, 0x77, 0x69, 0x72, 0x65, 0x6c, 0x65, 0x73,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x57, 0x69, 0x72, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08,
	0x77, 0x69, 0x72, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x22, 0xec, 0x02, 0x0a, 0x0e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x45, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x45, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6c,
	0x69, 0x73, 0x70, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x73, 0x70, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6d, 0x63, 0x65, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x70, 0x65, 0x6d, 0x63, 0x65, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70,
	0x65, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0d, 0x70, 0x65, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x6b, 0x65,
	0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x2e, 0x0a, 0x04, 0x61, 0x63, 0x6c, 0x73, 0x18, 0x28, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x43, 0x45, 0x52, 0x04, 0x61, 0x63, 0x6c,
	0x73, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x76, 0x6c, 0x61, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x56, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0xcf, 0x01, 0x0a, 0x0e, 0x57, 0x69, 0x72, 0x65,
	0x6c, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x57, 0x69, 0x72, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x43,
	0x66, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0b, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x43, 0x66, 0x67, 0x12, 0x3b, 0x0a, 0x07,
	0x77, 0x69, 0x66, 0x69, 0x43, 0x66, 0x67, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x07, 0x77, 0x69, 0x66, 0x69, 0x43, 0x66, 0x67, 0x22, 0x22, 0x0a, 0x0e, 0x43, 0x65, 0x6c,
	0x6c, 0x75, 0x6c, 0x61, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x41,
	0x50, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x50, 0x4e, 0x22, 0x92, 0x03,
	0x0a, 0x0a, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x77, 0x69, 0x66, 0x69, 0x53, 0x53, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x77, 0x69, 0x66, 0x69, 0x53, 0x53, 0x49, 0x44, 0x12, 0x42, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x57, 0x69, 0x46, 0x69, 0x4b, 0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x45, 0x0a, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x57, 0x69, 0x66,
	0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x19, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x45, 0x0a, 0x0b, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x42, 0x3d, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f,
	0x65, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_netconfig_proto_rawDescOnce sync.Once
	file_config_netconfig_proto_rawDescData = file_config_netconfig_proto_rawDesc
)

func file_config_netconfig_proto_rawDescGZIP() []byte {
	file_config_netconfig_proto_rawDescOnce.Do(func() {
		file_config_netconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_netconfig_proto_rawDescData)
	})
	return file_config_netconfig_proto_rawDescData
}

var file_config_netconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_config_netconfig_proto_goTypes = []interface{}{
	(*NetworkConfig)(nil),         // 0: org.lfedge.eve.config.NetworkConfig
	(*NetworkAdapter)(nil),        // 1: org.lfedge.eve.config.NetworkAdapter
	(*WirelessConfig)(nil),        // 2: org.lfedge.eve.config.WirelessConfig
	(*CellularConfig)(nil),        // 3: org.lfedge.eve.config.CellularConfig
	(*WifiConfig)(nil),            // 4: org.lfedge.eve.config.WifiConfig
	(*WifiConfigCryptoblock)(nil), // 5: org.lfedge.eve.config.WifiConfig.cryptoblock
	(NetworkType)(0),              // 6: org.lfedge.eve.config.NetworkType
	(*Ipspec)(nil),                // 7: org.lfedge.eve.config.ipspec
	(*ZnetStaticDNSEntry)(nil),    // 8: org.lfedge.eve.config.ZnetStaticDNSEntry
	(*ProxyConfig)(nil),           // 9: org.lfedge.eve.config.ProxyConfig
	(*ACE)(nil),                   // 10: org.lfedge.eve.config.ACE
	(WirelessType)(0),             // 11: org.lfedge.eve.config.WirelessType
	(WiFiKeyScheme)(0),            // 12: org.lfedge.eve.config.WiFiKeyScheme
	(*CipherBlock)(nil),           // 13: org.lfedge.eve.config.CipherBlock
}
var file_config_netconfig_proto_depIdxs = []int32{
	6,  // 0: org.lfedge.eve.config.NetworkConfig.type:type_name -> org.lfedge.eve.config.NetworkType
	7,  // 1: org.lfedge.eve.config.NetworkConfig.ip:type_name -> org.lfedge.eve.config.ipspec
	8,  // 2: org.lfedge.eve.config.NetworkConfig.dns:type_name -> org.lfedge.eve.config.ZnetStaticDNSEntry
	9,  // 3: org.lfedge.eve.config.NetworkConfig.entProxy:type_name -> org.lfedge.eve.config.ProxyConfig
	2,  // 4: org.lfedge.eve.config.NetworkConfig.wireless:type_name -> org.lfedge.eve.config.WirelessConfig
	10, // 5: org.lfedge.eve.config.NetworkAdapter.acls:type_name -> org.lfedge.eve.config.ACE
	11, // 6: org.lfedge.eve.config.WirelessConfig.type:type_name -> org.lfedge.eve.config.WirelessType
	3,  // 7: org.lfedge.eve.config.WirelessConfig.cellularCfg:type_name -> org.lfedge.eve.config.CellularConfig
	4,  // 8: org.lfedge.eve.config.WirelessConfig.wifiCfg:type_name -> org.lfedge.eve.config.WifiConfig
	12, // 9: org.lfedge.eve.config.WifiConfig.keyScheme:type_name -> org.lfedge.eve.config.WiFiKeyScheme
	5,  // 10: org.lfedge.eve.config.WifiConfig.crypto:type_name -> org.lfedge.eve.config.WifiConfig.cryptoblock
	13, // 11: org.lfedge.eve.config.WifiConfig.cipherData:type_name -> org.lfedge.eve.config.CipherBlock
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_config_netconfig_proto_init() }
func file_config_netconfig_proto_init() {
	if File_config_netconfig_proto != nil {
		return
	}
	file_config_acipherinfo_proto_init()
	file_config_fw_proto_init()
	file_config_netcmn_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_netconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_netconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkAdapter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_netconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WirelessConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_netconfig_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CellularConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_netconfig_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_netconfig_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiConfigCryptoblock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_netconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_netconfig_proto_goTypes,
		DependencyIndexes: file_config_netconfig_proto_depIdxs,
		MessageInfos:      file_config_netconfig_proto_msgTypes,
	}.Build()
	File_config_netconfig_proto = out.File
	file_config_netconfig_proto_rawDesc = nil
	file_config_netconfig_proto_goTypes = nil
	file_config_netconfig_proto_depIdxs = nil
}
