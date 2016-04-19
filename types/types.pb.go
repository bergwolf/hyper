// Code generated by protoc-gen-go.
// source: types/types.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	types/types.proto

It has these top-level messages:
	ContainerPort
	EnvironmentVar
	VolumeMount
	WaitingStatus
	RunningStatus
	TermStatus
	ContainerStatus
	ContainerInfo
	Container
	RBDVolumeSource
	PodVolume
	PodSpec
	PodStatus
	PodInfo
	PodInfoRequest
	PodInfoResponse
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Types definitions for HyperContainer
type ContainerPort struct {
	Name          string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	HostPort      int32  `protobuf:"varint,2,opt,name=hostPort" json:"hostPort,omitempty"`
	ContainerPort int32  `protobuf:"varint,3,opt,name=containerPort" json:"containerPort,omitempty"`
	Protocol      string `protobuf:"bytes,4,opt,name=protocol" json:"protocol,omitempty"`
	HostIP        string `protobuf:"bytes,5,opt,name=hostIP" json:"hostIP,omitempty"`
}

func (m *ContainerPort) Reset()                    { *m = ContainerPort{} }
func (m *ContainerPort) String() string            { return proto.CompactTextString(m) }
func (*ContainerPort) ProtoMessage()               {}
func (*ContainerPort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EnvironmentVar struct {
	Env   string `protobuf:"bytes,1,opt,name=env" json:"env,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *EnvironmentVar) Reset()                    { *m = EnvironmentVar{} }
func (m *EnvironmentVar) String() string            { return proto.CompactTextString(m) }
func (*EnvironmentVar) ProtoMessage()               {}
func (*EnvironmentVar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type VolumeMount struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	MountPath string `protobuf:"bytes,2,opt,name=mountPath" json:"mountPath,omitempty"`
	ReadOnly  bool   `protobuf:"varint,3,opt,name=readOnly" json:"readOnly,omitempty"`
}

func (m *VolumeMount) Reset()                    { *m = VolumeMount{} }
func (m *VolumeMount) String() string            { return proto.CompactTextString(m) }
func (*VolumeMount) ProtoMessage()               {}
func (*VolumeMount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type WaitingStatus struct {
	Reason string `protobuf:"bytes,1,opt,name=reason" json:"reason,omitempty"`
}

func (m *WaitingStatus) Reset()                    { *m = WaitingStatus{} }
func (m *WaitingStatus) String() string            { return proto.CompactTextString(m) }
func (*WaitingStatus) ProtoMessage()               {}
func (*WaitingStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RunningStatus struct {
	StartedAt string `protobuf:"bytes,1,opt,name=startedAt" json:"startedAt,omitempty"`
}

func (m *RunningStatus) Reset()                    { *m = RunningStatus{} }
func (m *RunningStatus) String() string            { return proto.CompactTextString(m) }
func (*RunningStatus) ProtoMessage()               {}
func (*RunningStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type TermStatus struct {
	ExitCode   int32  `protobuf:"varint,1,opt,name=exitCode" json:"exitCode,omitempty"`
	Reason     string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	Message    string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	StartedAt  string `protobuf:"bytes,4,opt,name=startedAt" json:"startedAt,omitempty"`
	FinishedAt string `protobuf:"bytes,5,opt,name=finishedAt" json:"finishedAt,omitempty"`
}

func (m *TermStatus) Reset()                    { *m = TermStatus{} }
func (m *TermStatus) String() string            { return proto.CompactTextString(m) }
func (*TermStatus) ProtoMessage()               {}
func (*TermStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ContainerStatus struct {
	Name        string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ContainerID string         `protobuf:"bytes,2,opt,name=containerID" json:"containerID,omitempty"`
	Phase       string         `protobuf:"bytes,3,opt,name=phase" json:"phase,omitempty"`
	Waiting     *WaitingStatus `protobuf:"bytes,4,opt,name=waiting" json:"waiting,omitempty"`
	Running     *RunningStatus `protobuf:"bytes,5,opt,name=running" json:"running,omitempty"`
	Terminated  *TermStatus    `protobuf:"bytes,6,opt,name=terminated" json:"terminated,omitempty"`
}

func (m *ContainerStatus) Reset()                    { *m = ContainerStatus{} }
func (m *ContainerStatus) String() string            { return proto.CompactTextString(m) }
func (*ContainerStatus) ProtoMessage()               {}
func (*ContainerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ContainerStatus) GetWaiting() *WaitingStatus {
	if m != nil {
		return m.Waiting
	}
	return nil
}

func (m *ContainerStatus) GetRunning() *RunningStatus {
	if m != nil {
		return m.Running
	}
	return nil
}

func (m *ContainerStatus) GetTerminated() *TermStatus {
	if m != nil {
		return m.Terminated
	}
	return nil
}

// TODO: embed Container message to avoid repetition
type ContainerInfo struct {
	Name            string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ContainerID     string            `protobuf:"bytes,2,opt,name=containerID" json:"containerID,omitempty"`
	Image           string            `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	ImageID         string            `protobuf:"bytes,4,opt,name=imageID" json:"imageID,omitempty"`
	Commands        []string          `protobuf:"bytes,5,rep,name=commands" json:"commands,omitempty"`
	Args            []string          `protobuf:"bytes,6,rep,name=args" json:"args,omitempty"`
	WorkingDir      string            `protobuf:"bytes,7,opt,name=workingDir" json:"workingDir,omitempty"`
	Ports           []*ContainerPort  `protobuf:"bytes,8,rep,name=ports" json:"ports,omitempty"`
	Env             []*EnvironmentVar `protobuf:"bytes,9,rep,name=env" json:"env,omitempty"`
	VolumeMounts    []*VolumeMount    `protobuf:"bytes,10,rep,name=volumeMounts" json:"volumeMounts,omitempty"`
	Tty             bool              `protobuf:"varint,11,opt,name=tty" json:"tty,omitempty"`
	ImagePullPolicy string            `protobuf:"bytes,12,opt,name=imagePullPolicy" json:"imagePullPolicy,omitempty"`
	PodID           string            `protobuf:"bytes,13,opt,name=podID" json:"podID,omitempty"`
	Status          *ContainerStatus  `protobuf:"bytes,14,opt,name=status" json:"status,omitempty"`
}

func (m *ContainerInfo) Reset()                    { *m = ContainerInfo{} }
func (m *ContainerInfo) String() string            { return proto.CompactTextString(m) }
func (*ContainerInfo) ProtoMessage()               {}
func (*ContainerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ContainerInfo) GetPorts() []*ContainerPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ContainerInfo) GetEnv() []*EnvironmentVar {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *ContainerInfo) GetVolumeMounts() []*VolumeMount {
	if m != nil {
		return m.VolumeMounts
	}
	return nil
}

func (m *ContainerInfo) GetStatus() *ContainerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type Container struct {
	Name            string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ContainerID     string            `protobuf:"bytes,2,opt,name=containerID" json:"containerID,omitempty"`
	Image           string            `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	ImageID         string            `protobuf:"bytes,4,opt,name=imageID" json:"imageID,omitempty"`
	Commands        []string          `protobuf:"bytes,5,rep,name=commands" json:"commands,omitempty"`
	Args            []string          `protobuf:"bytes,6,rep,name=args" json:"args,omitempty"`
	WorkingDir      string            `protobuf:"bytes,7,opt,name=workingDir" json:"workingDir,omitempty"`
	Ports           []*ContainerPort  `protobuf:"bytes,8,rep,name=ports" json:"ports,omitempty"`
	Env             []*EnvironmentVar `protobuf:"bytes,9,rep,name=env" json:"env,omitempty"`
	VolumeMounts    []*VolumeMount    `protobuf:"bytes,10,rep,name=volumeMounts" json:"volumeMounts,omitempty"`
	Tty             bool              `protobuf:"varint,11,opt,name=tty" json:"tty,omitempty"`
	ImagePullPolicy string            `protobuf:"bytes,12,opt,name=imagePullPolicy" json:"imagePullPolicy,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Container) GetPorts() []*ContainerPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Container) GetEnv() []*EnvironmentVar {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Container) GetVolumeMounts() []*VolumeMount {
	if m != nil {
		return m.VolumeMounts
	}
	return nil
}

type RBDVolumeSource struct {
	Monitors []string `protobuf:"bytes,1,rep,name=monitors" json:"monitors,omitempty"`
	Image    string   `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	FsType   string   `protobuf:"bytes,3,opt,name=fsType" json:"fsType,omitempty"`
	Pool     string   `protobuf:"bytes,4,opt,name=pool" json:"pool,omitempty"`
	User     string   `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
	Keyring  string   `protobuf:"bytes,6,opt,name=keyring" json:"keyring,omitempty"`
	ReadOnly bool     `protobuf:"varint,7,opt,name=readOnly" json:"readOnly,omitempty"`
}

func (m *RBDVolumeSource) Reset()                    { *m = RBDVolumeSource{} }
func (m *RBDVolumeSource) String() string            { return proto.CompactTextString(m) }
func (*RBDVolumeSource) ProtoMessage()               {}
func (*RBDVolumeSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type PodVolume struct {
	Name   string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Source string           `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	Driver string           `protobuf:"bytes,3,opt,name=driver" json:"driver,omitempty"`
	Rbd    *RBDVolumeSource `protobuf:"bytes,4,opt,name=rbd" json:"rbd,omitempty"`
}

func (m *PodVolume) Reset()                    { *m = PodVolume{} }
func (m *PodVolume) String() string            { return proto.CompactTextString(m) }
func (*PodVolume) ProtoMessage()               {}
func (*PodVolume) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PodVolume) GetRbd() *RBDVolumeSource {
	if m != nil {
		return m.Rbd
	}
	return nil
}

type PodSpec struct {
	Volumes    []*PodVolume      `protobuf:"bytes,1,rep,name=volumes" json:"volumes,omitempty"`
	Containers []*Container      `protobuf:"bytes,2,rep,name=containers" json:"containers,omitempty"`
	Labels     map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Vcpu       int32             `protobuf:"varint,4,opt,name=vcpu" json:"vcpu,omitempty"`
	Memory     int32             `protobuf:"varint,5,opt,name=memory" json:"memory,omitempty"`
}

func (m *PodSpec) Reset()                    { *m = PodSpec{} }
func (m *PodSpec) String() string            { return proto.CompactTextString(m) }
func (*PodSpec) ProtoMessage()               {}
func (*PodSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PodSpec) GetVolumes() []*PodVolume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *PodSpec) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *PodSpec) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type PodStatus struct {
	Phase           string             `protobuf:"bytes,1,opt,name=phase" json:"phase,omitempty"`
	Message         string             `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Reason          string             `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
	HostIP          string             `protobuf:"bytes,4,opt,name=hostIP" json:"hostIP,omitempty"`
	PodIP           []string           `protobuf:"bytes,5,rep,name=podIP" json:"podIP,omitempty"`
	StartTime       string             `protobuf:"bytes,6,opt,name=startTime" json:"startTime,omitempty"`
	ContainerStatus []*ContainerStatus `protobuf:"bytes,7,rep,name=containerStatus" json:"containerStatus,omitempty"`
	FinishTime      string             `protobuf:"bytes,8,opt,name=finishTime" json:"finishTime,omitempty"`
}

func (m *PodStatus) Reset()                    { *m = PodStatus{} }
func (m *PodStatus) String() string            { return proto.CompactTextString(m) }
func (*PodStatus) ProtoMessage()               {}
func (*PodStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PodStatus) GetContainerStatus() []*ContainerStatus {
	if m != nil {
		return m.ContainerStatus
	}
	return nil
}

type PodInfo struct {
	Kind       string     `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	ApiVersion string     `protobuf:"bytes,2,opt,name=apiVersion" json:"apiVersion,omitempty"`
	Vm         string     `protobuf:"bytes,3,opt,name=vm" json:"vm,omitempty"`
	Spec       *PodSpec   `protobuf:"bytes,4,opt,name=spec" json:"spec,omitempty"`
	Status     *PodStatus `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

func (m *PodInfo) Reset()                    { *m = PodInfo{} }
func (m *PodInfo) String() string            { return proto.CompactTextString(m) }
func (*PodInfo) ProtoMessage()               {}
func (*PodInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *PodInfo) GetSpec() *PodSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *PodInfo) GetStatus() *PodStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type PodInfoRequest struct {
	PodID string `protobuf:"bytes,1,opt,name=podID" json:"podID,omitempty"`
}

func (m *PodInfoRequest) Reset()                    { *m = PodInfoRequest{} }
func (m *PodInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*PodInfoRequest) ProtoMessage()               {}
func (*PodInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type PodInfoResponse struct {
	PodInfo *PodInfo `protobuf:"bytes,1,opt,name=podInfo" json:"podInfo,omitempty"`
}

func (m *PodInfoResponse) Reset()                    { *m = PodInfoResponse{} }
func (m *PodInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*PodInfoResponse) ProtoMessage()               {}
func (*PodInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *PodInfoResponse) GetPodInfo() *PodInfo {
	if m != nil {
		return m.PodInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerPort)(nil), "types.ContainerPort")
	proto.RegisterType((*EnvironmentVar)(nil), "types.EnvironmentVar")
	proto.RegisterType((*VolumeMount)(nil), "types.VolumeMount")
	proto.RegisterType((*WaitingStatus)(nil), "types.WaitingStatus")
	proto.RegisterType((*RunningStatus)(nil), "types.RunningStatus")
	proto.RegisterType((*TermStatus)(nil), "types.TermStatus")
	proto.RegisterType((*ContainerStatus)(nil), "types.ContainerStatus")
	proto.RegisterType((*ContainerInfo)(nil), "types.ContainerInfo")
	proto.RegisterType((*Container)(nil), "types.Container")
	proto.RegisterType((*RBDVolumeSource)(nil), "types.RBDVolumeSource")
	proto.RegisterType((*PodVolume)(nil), "types.PodVolume")
	proto.RegisterType((*PodSpec)(nil), "types.PodSpec")
	proto.RegisterType((*PodStatus)(nil), "types.PodStatus")
	proto.RegisterType((*PodInfo)(nil), "types.PodInfo")
	proto.RegisterType((*PodInfoRequest)(nil), "types.PodInfoRequest")
	proto.RegisterType((*PodInfoResponse)(nil), "types.PodInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for PublicAPI service

type PublicAPIClient interface {
	PodInfo(ctx context.Context, in *PodInfoRequest, opts ...grpc.CallOption) (*PodInfoResponse, error)
}

type publicAPIClient struct {
	cc *grpc.ClientConn
}

func NewPublicAPIClient(cc *grpc.ClientConn) PublicAPIClient {
	return &publicAPIClient{cc}
}

func (c *publicAPIClient) PodInfo(ctx context.Context, in *PodInfoRequest, opts ...grpc.CallOption) (*PodInfoResponse, error) {
	out := new(PodInfoResponse)
	err := grpc.Invoke(ctx, "/types.PublicAPI/PodInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PublicAPI service

type PublicAPIServer interface {
	PodInfo(context.Context, *PodInfoRequest) (*PodInfoResponse, error)
}

func RegisterPublicAPIServer(s *grpc.Server, srv PublicAPIServer) {
	s.RegisterService(&_PublicAPI_serviceDesc, srv)
}

func _PublicAPI_PodInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PodInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PublicAPIServer).PodInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _PublicAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.PublicAPI",
	HandlerType: (*PublicAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PodInfo",
			Handler:    _PublicAPI_PodInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 860 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xec, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x26, 0x71, 0x7e, 0x9a, 0x93, 0x26, 0xd9, 0x1a, 0xba, 0x58, 0x15, 0x52, 0x17, 0xc3, 0xa2,
	0x15, 0x82, 0xae, 0x14, 0x6e, 0xd0, 0x5e, 0xb1, 0xb4, 0x15, 0x8a, 0x04, 0xc2, 0x6a, 0xab, 0x72,
	0x87, 0x34, 0xb1, 0x27, 0xcd, 0xa8, 0xf6, 0x8c, 0x99, 0x19, 0xa7, 0xe4, 0x92, 0xe7, 0x40, 0xe2,
	0x39, 0x78, 0x07, 0xee, 0x79, 0x1e, 0x8e, 0x67, 0x26, 0xfe, 0x03, 0x21, 0x1e, 0x80, 0x9b, 0x48,
	0x73, 0xf2, 0xcd, 0x7c, 0xdf, 0xf9, 0xce, 0x8f, 0xe1, 0x44, 0xef, 0x73, 0xaa, 0x5e, 0x9b, 0xdf,
	0x8b, 0x5c, 0x0a, 0x2d, 0xfc, 0xa1, 0x39, 0x84, 0x5b, 0x98, 0x5d, 0x0a, 0xae, 0x09, 0xe3, 0x54,
	0x46, 0x42, 0x6a, 0xff, 0x18, 0x06, 0x9c, 0x64, 0x34, 0xe8, 0xbd, 0xe8, 0xbd, 0x9a, 0xf8, 0xcf,
	0xe0, 0x68, 0x2b, 0x94, 0x2e, 0xff, 0x09, 0xfa, 0x18, 0x19, 0xfa, 0xa7, 0x30, 0x8b, 0x9b, 0x17,
	0x02, 0xcf, 0x84, 0x11, 0x68, 0xde, 0x8d, 0x45, 0x1a, 0x0c, 0xcc, 0xd5, 0x39, 0x8c, 0xca, 0xab,
	0xab, 0x28, 0x18, 0x96, 0xe7, 0xf0, 0x33, 0x98, 0x5f, 0xf3, 0x1d, 0x93, 0x82, 0x67, 0x94, 0xeb,
	0x7b, 0x22, 0xfd, 0x29, 0x78, 0x94, 0xef, 0x1c, 0xd3, 0x0c, 0x86, 0x3b, 0x92, 0x16, 0xd4, 0xd0,
	0x4c, 0xc2, 0xaf, 0x60, 0x7a, 0x2f, 0xd2, 0x22, 0xa3, 0xdf, 0x89, 0x82, 0x77, 0x55, 0x9d, 0xc0,
	0x24, 0x2b, 0xc3, 0x11, 0xd1, 0x5b, 0x8b, 0x2f, 0xf9, 0x25, 0x25, 0xc9, 0xf7, 0x3c, 0xdd, 0x1b,
	0x45, 0x47, 0xe1, 0x39, 0xcc, 0x7e, 0x20, 0x4c, 0x33, 0xfe, 0x70, 0xab, 0x89, 0x2e, 0x54, 0x29,
	0x08, 0x21, 0x4a, 0x70, 0xfb, 0x4a, 0x18, 0xc2, 0xec, 0xa6, 0xe0, 0xbc, 0x06, 0xe0, 0xb3, 0x4a,
	0x13, 0xa9, 0x69, 0xf2, 0x56, 0x3b, 0xcc, 0x06, 0xe0, 0x8e, 0xca, 0xcc, 0x01, 0x90, 0x84, 0xfe,
	0xcc, 0xf4, 0xa5, 0x48, 0xac, 0x92, 0x61, 0xe3, 0x4d, 0x2b, 0x63, 0x01, 0xe3, 0x8c, 0x2a, 0x45,
	0x1e, 0xa8, 0x51, 0x31, 0x69, 0xbf, 0x69, 0x8d, 0xf1, 0x01, 0x36, 0x8c, 0x33, 0xb5, 0x35, 0x31,
	0x6b, 0xce, 0xef, 0x3d, 0x58, 0x54, 0x75, 0x70, 0x6c, 0xed, 0x9c, 0xdf, 0x85, 0x69, 0xe5, 0xfb,
	0xea, 0xca, 0xd1, 0xa1, 0x69, 0xf9, 0x96, 0xa8, 0x03, 0xd9, 0x4b, 0x18, 0x3f, 0xd9, 0x94, 0x0d,
	0xd5, 0x74, 0xf9, 0xde, 0x85, 0x2d, 0x79, 0xdb, 0x08, 0x84, 0x49, 0x9b, 0xb8, 0x61, 0xaf, 0x61,
	0x6d, 0x3b, 0x5e, 0x02, 0x68, 0xcc, 0x9d, 0x71, 0x82, 0xea, 0x83, 0x91, 0x41, 0x9e, 0x38, 0x64,
	0x6d, 0x4a, 0xf8, 0x67, 0xbf, 0xd1, 0x42, 0x2b, 0xbe, 0x11, 0xff, 0x51, 0x38, 0xcb, 0x6a, 0x97,
	0xd0, 0x36, 0x73, 0xc4, 0xff, 0x07, 0x87, 0x72, 0xc6, 0x22, 0xcb, 0x08, 0x4f, 0x14, 0x6a, 0xf4,
	0x30, 0x82, 0x8f, 0x12, 0xf9, 0xa0, 0x50, 0x87, 0x67, 0x3d, 0x7c, 0x12, 0xf2, 0x11, 0xc5, 0x5e,
	0x31, 0x19, 0x8c, 0xcd, 0x9d, 0x8f, 0xd0, 0x0c, 0x6c, 0x48, 0x15, 0x1c, 0x21, 0xa4, 0x4e, 0xaa,
	0xdd, 0xde, 0xa1, 0xed, 0xb9, 0x89, 0x81, 0x9c, 0x3a, 0x48, 0xa7, 0x2f, 0x5f, 0xc1, 0xf1, 0xae,
	0xee, 0x3d, 0x15, 0x80, 0x01, 0xfb, 0x0e, 0xdc, 0x6c, 0x4b, 0xec, 0x60, 0xad, 0xf7, 0xc1, 0xb4,
	0x6c, 0x38, 0xff, 0x7d, 0x58, 0x98, 0x24, 0xa2, 0x22, 0x4d, 0x23, 0x91, 0xb2, 0x78, 0x1f, 0x1c,
	0x57, 0x55, 0x12, 0x09, 0xe6, 0x36, 0x33, 0xc7, 0x4f, 0x60, 0xa4, 0x8c, 0x75, 0xc1, 0xdc, 0x78,
	0xfa, 0xbc, 0x2b, 0xd4, 0x19, 0xfb, 0x6b, 0x1f, 0x26, 0x55, 0xec, 0x7f, 0x53, 0xdb, 0xa6, 0x86,
	0xbf, 0xe0, 0xc4, 0xdc, 0x7c, 0x7d, 0x65, 0x2f, 0xde, 0x8a, 0x42, 0xc6, 0xb4, 0x4c, 0x30, 0x13,
	0x9c, 0x69, 0x21, 0x15, 0xfa, 0xe4, 0x35, 0x2d, 0xe9, 0x1f, 0x76, 0xd2, 0x46, 0xdd, 0x21, 0xa3,
	0xb3, 0x08, 0xf3, 0xcf, 0x45, 0xb5, 0xb1, 0xf0, 0x54, 0x28, 0x2a, 0xed, 0x48, 0x96, 0xf6, 0x3d,
	0xd2, 0xbd, 0x2c, 0xa7, 0x64, 0xf4, 0xb7, 0x15, 0x33, 0x36, 0x2b, 0xe6, 0x47, 0x98, 0x44, 0x22,
	0xb1, 0x12, 0x3a, 0x05, 0x42, 0x26, 0x65, 0x44, 0xd5, 0xcc, 0x89, 0x64, 0x3b, 0x7c, 0xdd, 0x73,
	0xbe, 0x7a, 0x72, 0x9d, 0xb8, 0x31, 0x3d, 0x74, 0x40, 0x27, 0x9f, 0xf0, 0x8f, 0x1e, 0x8c, 0x91,
	0xe0, 0x36, 0xa7, 0xb1, 0xff, 0x21, 0x8c, 0xad, 0x7f, 0x36, 0xb5, 0xe9, 0xf2, 0x99, 0xbb, 0x54,
	0x2b, 0xf8, 0x18, 0xa0, 0x6a, 0x0a, 0x85, 0xbc, 0x4d, 0x54, 0xdd, 0x48, 0x9f, 0xc2, 0x28, 0x25,
	0x6b, 0x9a, 0x2a, 0x54, 0x52, 0x22, 0xce, 0xea, 0x77, 0x4a, 0xa2, 0x8b, 0x6f, 0xcd, 0x9f, 0xd7,
	0x5c, 0xcb, 0x7d, 0x99, 0xd3, 0x2e, 0xce, 0x0b, 0x23, 0xd3, 0x2c, 0xbb, 0x8c, 0x66, 0x42, 0xee,
	0x8d, 0x43, 0xc3, 0xb3, 0xcf, 0x61, 0xda, 0x04, 0x63, 0xdd, 0xd0, 0xb0, 0x7f, 0x5c, 0xe7, 0x6f,
	0xfa, 0x5f, 0xf6, 0xc2, 0xdf, 0x7a, 0xc6, 0x2e, 0xb7, 0x5d, 0xaa, 0xd5, 0xd5, 0xeb, 0x2e, 0xce,
	0xca, 0x30, 0xb7, 0x59, 0xbd, 0xce, 0xe7, 0x64, 0xd0, 0x1c, 0xaa, 0xc8, 0x75, 0xf2, 0x61, 0xcf,
	0xde, 0x31, 0x2c, 0x81, 0xad, 0xd7, 0x6b, 0x58, 0xc4, 0xed, 0x91, 0xc2, 0xb2, 0x79, 0xff, 0x32,
	0x70, 0x4f, 0xc6, 0xed, 0xc3, 0x0a, 0xc3, 0x39, 0x48, 0x9c, 0x38, 0x1c, 0x0c, 0x92, 0xb3, 0x7b,
	0x74, 0x95, 0x55, 0x9b, 0x1e, 0xa0, 0xbf, 0xcb, 0x9c, 0xb6, 0x0f, 0x60, 0xa0, 0xd0, 0x3a, 0x57,
	0xcd, 0x79, 0xdb, 0x50, 0xff, 0x45, 0x35, 0xef, 0x76, 0xdb, 0x36, 0x0a, 0xe7, 0x88, 0xcf, 0x61,
	0xee, 0x88, 0x6f, 0xe8, 0x4f, 0x05, 0x55, 0xba, 0x5e, 0x19, 0xf6, 0x33, 0xb4, 0x84, 0x45, 0x05,
	0x50, 0xb9, 0xe0, 0x8a, 0xfa, 0xe7, 0x30, 0xce, 0x6d, 0xc8, 0x60, 0x5a, 0xb4, 0x65, 0x74, 0xf9,
	0x0d, 0xba, 0x5d, 0xac, 0x71, 0x62, 0xde, 0x46, 0x2b, 0xff, 0x4d, 0x9d, 0xda, 0x69, 0x1b, 0xe7,
	0x18, 0xcf, 0x9e, 0x77, 0xc3, 0x96, 0x27, 0x7c, 0x67, 0x3d, 0x32, 0x1f, 0xf6, 0x2f, 0xfe, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0xd4, 0x5b, 0x52, 0xb3, 0x45, 0x08, 0x00, 0x00,
}
