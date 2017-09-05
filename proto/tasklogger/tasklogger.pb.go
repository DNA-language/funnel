// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tasklogger.proto

/*
Package tasklogger is a generated protocol buffer package.

It is generated from these files:
	tasklogger.proto

It has these top-level messages:
	UpdateExecutorLogsRequest
	UpdateExecutorLogsResponse
	UpdateTaskLogsRequest
	UpdateTaskLogsResponse
	UpdateTaskStateRequest
	UpdateTaskStateResponse
*/
package tasklogger

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tes "github.com/ohsu-comp-bio/funnel/proto/tes"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateExecutorLogsRequest struct {
	Id       string           `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Step     int64            `protobuf:"varint,2,opt,name=step" json:"step,omitempty"`
	Log      *tes.ExecutorLog `protobuf:"bytes,4,opt,name=log" json:"log,omitempty"`
	WorkerId string           `protobuf:"bytes,5,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
}

func (m *UpdateExecutorLogsRequest) Reset()                    { *m = UpdateExecutorLogsRequest{} }
func (m *UpdateExecutorLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateExecutorLogsRequest) ProtoMessage()               {}
func (*UpdateExecutorLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpdateExecutorLogsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateExecutorLogsRequest) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *UpdateExecutorLogsRequest) GetLog() *tes.ExecutorLog {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *UpdateExecutorLogsRequest) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

type UpdateExecutorLogsResponse struct {
}

func (m *UpdateExecutorLogsResponse) Reset()                    { *m = UpdateExecutorLogsResponse{} }
func (m *UpdateExecutorLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateExecutorLogsResponse) ProtoMessage()               {}
func (*UpdateExecutorLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type UpdateTaskLogsRequest struct {
	Id      string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	TaskLog *tes.TaskLog `protobuf:"bytes,2,opt,name=task_log,json=taskLog" json:"task_log,omitempty"`
}

func (m *UpdateTaskLogsRequest) Reset()                    { *m = UpdateTaskLogsRequest{} }
func (m *UpdateTaskLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateTaskLogsRequest) ProtoMessage()               {}
func (*UpdateTaskLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateTaskLogsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTaskLogsRequest) GetTaskLog() *tes.TaskLog {
	if m != nil {
		return m.TaskLog
	}
	return nil
}

type UpdateTaskLogsResponse struct {
}

func (m *UpdateTaskLogsResponse) Reset()                    { *m = UpdateTaskLogsResponse{} }
func (m *UpdateTaskLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateTaskLogsResponse) ProtoMessage()               {}
func (*UpdateTaskLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type UpdateTaskStateRequest struct {
	Id    string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	State tes.State `protobuf:"varint,2,opt,name=state,enum=tes.State" json:"state,omitempty"`
}

func (m *UpdateTaskStateRequest) Reset()                    { *m = UpdateTaskStateRequest{} }
func (m *UpdateTaskStateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateTaskStateRequest) ProtoMessage()               {}
func (*UpdateTaskStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateTaskStateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTaskStateRequest) GetState() tes.State {
	if m != nil {
		return m.State
	}
	return tes.State_UNKNOWN
}

type UpdateTaskStateResponse struct {
}

func (m *UpdateTaskStateResponse) Reset()                    { *m = UpdateTaskStateResponse{} }
func (m *UpdateTaskStateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateTaskStateResponse) ProtoMessage()               {}
func (*UpdateTaskStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*UpdateExecutorLogsRequest)(nil), "tasklogger.UpdateExecutorLogsRequest")
	proto.RegisterType((*UpdateExecutorLogsResponse)(nil), "tasklogger.UpdateExecutorLogsResponse")
	proto.RegisterType((*UpdateTaskLogsRequest)(nil), "tasklogger.UpdateTaskLogsRequest")
	proto.RegisterType((*UpdateTaskLogsResponse)(nil), "tasklogger.UpdateTaskLogsResponse")
	proto.RegisterType((*UpdateTaskStateRequest)(nil), "tasklogger.UpdateTaskStateRequest")
	proto.RegisterType((*UpdateTaskStateResponse)(nil), "tasklogger.UpdateTaskStateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskLoggerService service

type TaskLoggerServiceClient interface {
	UpdateExecutorLogs(ctx context.Context, in *UpdateExecutorLogsRequest, opts ...grpc.CallOption) (*UpdateExecutorLogsResponse, error)
	UpdateTaskLogs(ctx context.Context, in *UpdateTaskLogsRequest, opts ...grpc.CallOption) (*UpdateTaskLogsResponse, error)
	UpdateTaskState(ctx context.Context, in *UpdateTaskStateRequest, opts ...grpc.CallOption) (*UpdateTaskStateResponse, error)
}

type taskLoggerServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskLoggerServiceClient(cc *grpc.ClientConn) TaskLoggerServiceClient {
	return &taskLoggerServiceClient{cc}
}

func (c *taskLoggerServiceClient) UpdateExecutorLogs(ctx context.Context, in *UpdateExecutorLogsRequest, opts ...grpc.CallOption) (*UpdateExecutorLogsResponse, error) {
	out := new(UpdateExecutorLogsResponse)
	err := grpc.Invoke(ctx, "/tasklogger.TaskLoggerService/UpdateExecutorLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskLoggerServiceClient) UpdateTaskLogs(ctx context.Context, in *UpdateTaskLogsRequest, opts ...grpc.CallOption) (*UpdateTaskLogsResponse, error) {
	out := new(UpdateTaskLogsResponse)
	err := grpc.Invoke(ctx, "/tasklogger.TaskLoggerService/UpdateTaskLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskLoggerServiceClient) UpdateTaskState(ctx context.Context, in *UpdateTaskStateRequest, opts ...grpc.CallOption) (*UpdateTaskStateResponse, error) {
	out := new(UpdateTaskStateResponse)
	err := grpc.Invoke(ctx, "/tasklogger.TaskLoggerService/UpdateTaskState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskLoggerService service

type TaskLoggerServiceServer interface {
	UpdateExecutorLogs(context.Context, *UpdateExecutorLogsRequest) (*UpdateExecutorLogsResponse, error)
	UpdateTaskLogs(context.Context, *UpdateTaskLogsRequest) (*UpdateTaskLogsResponse, error)
	UpdateTaskState(context.Context, *UpdateTaskStateRequest) (*UpdateTaskStateResponse, error)
}

func RegisterTaskLoggerServiceServer(s *grpc.Server, srv TaskLoggerServiceServer) {
	s.RegisterService(&_TaskLoggerService_serviceDesc, srv)
}

func _TaskLoggerService_UpdateExecutorLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExecutorLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskLoggerServiceServer).UpdateExecutorLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasklogger.TaskLoggerService/UpdateExecutorLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskLoggerServiceServer).UpdateExecutorLogs(ctx, req.(*UpdateExecutorLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskLoggerService_UpdateTaskLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskLoggerServiceServer).UpdateTaskLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasklogger.TaskLoggerService/UpdateTaskLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskLoggerServiceServer).UpdateTaskLogs(ctx, req.(*UpdateTaskLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskLoggerService_UpdateTaskState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskLoggerServiceServer).UpdateTaskState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasklogger.TaskLoggerService/UpdateTaskState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskLoggerServiceServer).UpdateTaskState(ctx, req.(*UpdateTaskStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskLoggerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tasklogger.TaskLoggerService",
	HandlerType: (*TaskLoggerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateExecutorLogs",
			Handler:    _TaskLoggerService_UpdateExecutorLogs_Handler,
		},
		{
			MethodName: "UpdateTaskLogs",
			Handler:    _TaskLoggerService_UpdateTaskLogs_Handler,
		},
		{
			MethodName: "UpdateTaskState",
			Handler:    _TaskLoggerService_UpdateTaskState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tasklogger.proto",
}

func init() { proto.RegisterFile("tasklogger.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0xb5, 0x05, 0x14, 0x46, 0x83, 0xb8, 0x89, 0x5a, 0x2a, 0x87, 0xba, 0x46, 0xe5, 0x04, 0x09,
	0x7e, 0x83, 0x07, 0x8d, 0x07, 0x53, 0xf4, 0x60, 0x62, 0x42, 0x56, 0x3a, 0xd9, 0x6c, 0x20, 0xdd,
	0xba, 0xbb, 0xa8, 0x47, 0xff, 0xc9, 0x1f, 0x34, 0xdd, 0xad, 0x5a, 0x41, 0xea, 0x6d, 0x33, 0x6f,
	0xe6, 0xbd, 0x37, 0x6f, 0x16, 0x3a, 0x86, 0xe9, 0xd9, 0x5c, 0x72, 0x8e, 0x6a, 0x90, 0x29, 0x69,
	0x24, 0x81, 0x9f, 0x4a, 0xd8, 0x32, 0xa8, 0x5d, 0x39, 0xec, 0x71, 0x29, 0xf9, 0x1c, 0x87, 0x2c,
	0x13, 0x43, 0x96, 0xa6, 0xd2, 0x30, 0x23, 0x64, 0x5a, 0xa0, 0xf4, 0xdd, 0x83, 0xee, 0x7d, 0x96,
	0x30, 0x83, 0x97, 0x6f, 0x38, 0x5d, 0x18, 0xa9, 0x6e, 0x24, 0xd7, 0x31, 0x3e, 0x2f, 0x50, 0x1b,
	0xd2, 0x06, 0x5f, 0x24, 0x81, 0x17, 0x79, 0xfd, 0x56, 0xec, 0x8b, 0x84, 0x10, 0xa8, 0x6b, 0x83,
	0x59, 0xe0, 0x47, 0x5e, 0xbf, 0x16, 0xdb, 0x37, 0xa1, 0x50, 0x9b, 0x4b, 0x1e, 0xd4, 0x23, 0xaf,
	0xbf, 0x3d, 0xea, 0x0c, 0x72, 0xe1, 0x12, 0x55, 0x9c, 0x83, 0xe4, 0x08, 0x5a, 0xaf, 0x52, 0xcd,
	0x50, 0x4d, 0x44, 0x12, 0x34, 0x2c, 0x5d, 0xd3, 0x15, 0xae, 0x12, 0xda, 0x83, 0xf0, 0x2f, 0x07,
	0x3a, 0x93, 0xa9, 0x46, 0x7a, 0x0b, 0xfb, 0x0e, 0xbd, 0x63, 0x7a, 0x56, 0xe5, 0xed, 0x1c, 0x9a,
	0x79, 0x00, 0x93, 0xdc, 0x8c, 0x6f, 0xcd, 0xec, 0x58, 0x33, 0xc5, 0x5c, 0xbc, 0x65, 0xdc, 0x83,
	0x06, 0x70, 0xb0, 0xcc, 0x58, 0x68, 0x5d, 0x97, 0x91, 0xb1, 0x61, 0x06, 0xd7, 0x89, 0x45, 0xd0,
	0xd0, 0x39, 0x6e, 0x95, 0xda, 0x23, 0xb0, 0x4a, 0x6e, 0xc2, 0x01, 0xb4, 0x0b, 0x87, 0x2b, 0x5c,
	0x4e, 0x66, 0xf4, 0xe1, 0xc3, 0x5e, 0xa1, 0xcd, 0x51, 0x8d, 0x51, 0xbd, 0x88, 0x29, 0x12, 0x04,
	0xb2, 0x1a, 0x03, 0x39, 0x1d, 0x94, 0xee, 0xbc, 0xf6, 0x50, 0xe1, 0xd9, 0x7f, 0x6d, 0xc5, 0x86,
	0x1b, 0xe4, 0x01, 0xda, 0xbf, 0xb7, 0x27, 0xc7, 0xab, 0xb3, 0x4b, 0x59, 0x87, 0xb4, 0xaa, 0xe5,
	0x9b, 0xfa, 0x11, 0x76, 0x97, 0x56, 0x26, 0x6b, 0x06, 0xcb, 0xd9, 0x86, 0x27, 0x95, 0x3d, 0x5f,
	0xec, 0x4f, 0x9b, 0xf6, 0xc3, 0x5e, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x33, 0x9a, 0x5c,
	0xf9, 0x02, 0x00, 0x00,
}