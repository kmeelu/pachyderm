// Code generated by protoc-gen-go.
// source: server/pps/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	server/pps/persist/persist.proto

It has these top-level messages:
	JobInfo
	JobInfos
	JobOutput
	JobState
	PipelineInfo
	PipelineInfoChange
	PipelineInfos
	SubscribePipelineInfosRequest
	ListPipelineInfosRequest
	Shard
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"
import google_protobuf1 "go.pedge.io/pb/go/google/protobuf"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"
import pachyderm_pps "github.com/pachyderm/pachyderm/src/client/pps"

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

type JobInfo struct {
	JobID                      string                      `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	Transform                  *pachyderm_pps.Transform    `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	PipelineName               string                      `protobuf:"bytes,3,opt,name=pipeline_name,json=pipelineName" json:"pipeline_name,omitempty"`
	Parallelism                uint64                      `protobuf:"varint,4,opt,name=parallelism" json:"parallelism,omitempty"`
	Inputs                     []*pachyderm_pps.JobInput   `protobuf:"bytes,5,rep,name=inputs" json:"inputs,omitempty"`
	ParentJob                  *pachyderm_pps.Job          `protobuf:"bytes,6,opt,name=parent_job,json=parentJob" json:"parent_job,omitempty"`
	CreatedAt                  *google_protobuf1.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	OutputCommit               *pfs.Commit                 `protobuf:"bytes,8,opt,name=output_commit,json=outputCommit" json:"output_commit,omitempty"`
	State                      pachyderm_pps.JobState      `protobuf:"varint,9,opt,name=state,enum=pachyderm.pps.JobState" json:"state,omitempty"`
	CommitIndex                string                      `protobuf:"bytes,10,opt,name=commit_index,json=commitIndex" json:"commit_index,omitempty"`
	PodsStarted                uint64                      `protobuf:"varint,11,opt,name=pods_started,json=podsStarted" json:"pods_started,omitempty"`
	PodsSucceeded              uint64                      `protobuf:"varint,12,opt,name=pods_succeeded,json=podsSucceeded" json:"pods_succeeded,omitempty"`
	PodsFailed                 uint64                      `protobuf:"varint,13,opt,name=pods_failed,json=podsFailed" json:"pods_failed,omitempty"`
	NonEmptyFilterShardNumbers []uint64                    `protobuf:"varint,14,rep,name=non_empty_filter_shard_numbers,json=nonEmptyFilterShardNumbers" json:"non_empty_filter_shard_numbers,omitempty"`
	ShardModulus               uint64                      `protobuf:"varint,15,opt,name=shard_modulus,json=shardModulus" json:"shard_modulus,omitempty"`
}

func (m *JobInfo) Reset()                    { *m = JobInfo{} }
func (m *JobInfo) String() string            { return proto.CompactTextString(m) }
func (*JobInfo) ProtoMessage()               {}
func (*JobInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *JobInfo) GetTransform() *pachyderm_pps.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *JobInfo) GetInputs() []*pachyderm_pps.JobInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *JobInfo) GetParentJob() *pachyderm_pps.Job {
	if m != nil {
		return m.ParentJob
	}
	return nil
}

func (m *JobInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *JobInfo) GetOutputCommit() *pfs.Commit {
	if m != nil {
		return m.OutputCommit
	}
	return nil
}

type JobInfos struct {
	JobInfo []*JobInfo `protobuf:"bytes,1,rep,name=job_info,json=jobInfo" json:"job_info,omitempty"`
}

func (m *JobInfos) Reset()                    { *m = JobInfos{} }
func (m *JobInfos) String() string            { return proto.CompactTextString(m) }
func (*JobInfos) ProtoMessage()               {}
func (*JobInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JobInfos) GetJobInfo() []*JobInfo {
	if m != nil {
		return m.JobInfo
	}
	return nil
}

type JobOutput struct {
	JobID        string      `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	OutputCommit *pfs.Commit `protobuf:"bytes,2,opt,name=output_commit,json=outputCommit" json:"output_commit,omitempty"`
}

func (m *JobOutput) Reset()                    { *m = JobOutput{} }
func (m *JobOutput) String() string            { return proto.CompactTextString(m) }
func (*JobOutput) ProtoMessage()               {}
func (*JobOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *JobOutput) GetOutputCommit() *pfs.Commit {
	if m != nil {
		return m.OutputCommit
	}
	return nil
}

type JobState struct {
	JobID string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	State pachyderm_pps.JobState `protobuf:"varint,2,opt,name=state,enum=pachyderm.pps.JobState" json:"state,omitempty"`
}

func (m *JobState) Reset()                    { *m = JobState{} }
func (m *JobState) String() string            { return proto.CompactTextString(m) }
func (*JobState) ProtoMessage()               {}
func (*JobState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PipelineInfo struct {
	PipelineName string                         `protobuf:"bytes,1,opt,name=pipeline_name,json=pipelineName" json:"pipeline_name,omitempty"`
	Transform    *pachyderm_pps.Transform       `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Parallelism  uint64                         `protobuf:"varint,3,opt,name=parallelism" json:"parallelism,omitempty"`
	Inputs       []*pachyderm_pps.PipelineInput `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	OutputRepo   *pfs.Repo                      `protobuf:"bytes,5,opt,name=output_repo,json=outputRepo" json:"output_repo,omitempty"`
	CreatedAt    *google_protobuf1.Timestamp    `protobuf:"bytes,6,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Shard        uint64                         `protobuf:"varint,7,opt,name=shard" json:"shard,omitempty"`
}

func (m *PipelineInfo) Reset()                    { *m = PipelineInfo{} }
func (m *PipelineInfo) String() string            { return proto.CompactTextString(m) }
func (*PipelineInfo) ProtoMessage()               {}
func (*PipelineInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PipelineInfo) GetTransform() *pachyderm_pps.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *PipelineInfo) GetInputs() []*pachyderm_pps.PipelineInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PipelineInfo) GetOutputRepo() *pfs.Repo {
	if m != nil {
		return m.OutputRepo
	}
	return nil
}

func (m *PipelineInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type PipelineInfoChange struct {
	Pipeline *PipelineInfo `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	Removed  bool          `protobuf:"varint,2,opt,name=removed" json:"removed,omitempty"`
}

func (m *PipelineInfoChange) Reset()                    { *m = PipelineInfoChange{} }
func (m *PipelineInfoChange) String() string            { return proto.CompactTextString(m) }
func (*PipelineInfoChange) ProtoMessage()               {}
func (*PipelineInfoChange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PipelineInfoChange) GetPipeline() *PipelineInfo {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type PipelineInfos struct {
	PipelineInfo []*PipelineInfo `protobuf:"bytes,1,rep,name=pipeline_info,json=pipelineInfo" json:"pipeline_info,omitempty"`
}

func (m *PipelineInfos) Reset()                    { *m = PipelineInfos{} }
func (m *PipelineInfos) String() string            { return proto.CompactTextString(m) }
func (*PipelineInfos) ProtoMessage()               {}
func (*PipelineInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PipelineInfos) GetPipelineInfo() []*PipelineInfo {
	if m != nil {
		return m.PipelineInfo
	}
	return nil
}

type SubscribePipelineInfosRequest struct {
	IncludeInitial bool   `protobuf:"varint,1,opt,name=include_initial,json=includeInitial" json:"include_initial,omitempty"`
	Shard          *Shard `protobuf:"bytes,2,opt,name=shard" json:"shard,omitempty"`
}

func (m *SubscribePipelineInfosRequest) Reset()                    { *m = SubscribePipelineInfosRequest{} }
func (m *SubscribePipelineInfosRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribePipelineInfosRequest) ProtoMessage()               {}
func (*SubscribePipelineInfosRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SubscribePipelineInfosRequest) GetShard() *Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

type ListPipelineInfosRequest struct {
	Shard *Shard `protobuf:"bytes,1,opt,name=shard" json:"shard,omitempty"`
}

func (m *ListPipelineInfosRequest) Reset()                    { *m = ListPipelineInfosRequest{} }
func (m *ListPipelineInfosRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPipelineInfosRequest) ProtoMessage()               {}
func (*ListPipelineInfosRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListPipelineInfosRequest) GetShard() *Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

// As in, sharding
type Shard struct {
	Number uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *Shard) Reset()                    { *m = Shard{} }
func (m *Shard) String() string            { return proto.CompactTextString(m) }
func (*Shard) ProtoMessage()               {}
func (*Shard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*JobInfo)(nil), "pachyderm.pps.persist.JobInfo")
	proto.RegisterType((*JobInfos)(nil), "pachyderm.pps.persist.JobInfos")
	proto.RegisterType((*JobOutput)(nil), "pachyderm.pps.persist.JobOutput")
	proto.RegisterType((*JobState)(nil), "pachyderm.pps.persist.JobState")
	proto.RegisterType((*PipelineInfo)(nil), "pachyderm.pps.persist.PipelineInfo")
	proto.RegisterType((*PipelineInfoChange)(nil), "pachyderm.pps.persist.PipelineInfoChange")
	proto.RegisterType((*PipelineInfos)(nil), "pachyderm.pps.persist.PipelineInfos")
	proto.RegisterType((*SubscribePipelineInfosRequest)(nil), "pachyderm.pps.persist.SubscribePipelineInfosRequest")
	proto.RegisterType((*ListPipelineInfosRequest)(nil), "pachyderm.pps.persist.ListPipelineInfosRequest")
	proto.RegisterType((*Shard)(nil), "pachyderm.pps.persist.Shard")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for API service

type APIClient interface {
	// Job rpcs
	// job_id cannot be set
	// timestamp cannot be set
	CreateJobInfo(ctx context.Context, in *JobInfo, opts ...grpc.CallOption) (*JobInfo, error)
	InspectJob(ctx context.Context, in *pachyderm_pps.InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error)
	// ordered by time, latest to earliest
	ListJobInfos(ctx context.Context, in *pachyderm_pps.ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error)
	// should only be called when rolling back if a Job does not start!
	DeleteJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// JobOutput rpcs
	CreateJobOutput(ctx context.Context, in *JobOutput, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// JobState rpcs
	CreateJobState(ctx context.Context, in *JobState, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// Pipeline rpcs
	CreatePipelineInfo(ctx context.Context, in *PipelineInfo, opts ...grpc.CallOption) (*PipelineInfo, error)
	GetPipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*PipelineInfo, error)
	// ordered by time, latest to earliest
	ListPipelineInfos(ctx context.Context, in *ListPipelineInfosRequest, opts ...grpc.CallOption) (*PipelineInfos, error)
	DeletePipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	SubscribePipelineInfos(ctx context.Context, in *SubscribePipelineInfosRequest, opts ...grpc.CallOption) (API_SubscribePipelineInfosClient, error)
	// Shard rpcs
	// Returns the new job info
	StartPod(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error)
	SucceedPod(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error)
	FailPod(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CreateJobInfo(ctx context.Context, in *JobInfo, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) InspectJob(ctx context.Context, in *pachyderm_pps.InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/InspectJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListJobInfos(ctx context.Context, in *pachyderm_pps.ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error) {
	out := new(JobInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/ListJobInfos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/DeleteJobInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateJobOutput(ctx context.Context, in *JobOutput, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobOutput", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateJobState(ctx context.Context, in *JobState, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreatePipelineInfo(ctx context.Context, in *PipelineInfo, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreatePipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetPipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/GetPipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListPipelineInfos(ctx context.Context, in *ListPipelineInfosRequest, opts ...grpc.CallOption) (*PipelineInfos, error) {
	out := new(PipelineInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/ListPipelineInfos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeletePipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/DeletePipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) SubscribePipelineInfos(ctx context.Context, in *SubscribePipelineInfosRequest, opts ...grpc.CallOption) (API_SubscribePipelineInfosClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_API_serviceDesc.Streams[0], c.cc, "/pachyderm.pps.persist.API/SubscribePipelineInfos", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPISubscribePipelineInfosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_SubscribePipelineInfosClient interface {
	Recv() (*PipelineInfoChange, error)
	grpc.ClientStream
}

type aPISubscribePipelineInfosClient struct {
	grpc.ClientStream
}

func (x *aPISubscribePipelineInfosClient) Recv() (*PipelineInfoChange, error) {
	m := new(PipelineInfoChange)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) StartPod(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/StartPod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) SucceedPod(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/SucceedPod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) FailPod(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/FailPod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	// Job rpcs
	// job_id cannot be set
	// timestamp cannot be set
	CreateJobInfo(context.Context, *JobInfo) (*JobInfo, error)
	InspectJob(context.Context, *pachyderm_pps.InspectJobRequest) (*JobInfo, error)
	// ordered by time, latest to earliest
	ListJobInfos(context.Context, *pachyderm_pps.ListJobRequest) (*JobInfos, error)
	// should only be called when rolling back if a Job does not start!
	DeleteJobInfo(context.Context, *pachyderm_pps.Job) (*google_protobuf.Empty, error)
	// JobOutput rpcs
	CreateJobOutput(context.Context, *JobOutput) (*google_protobuf.Empty, error)
	// JobState rpcs
	CreateJobState(context.Context, *JobState) (*google_protobuf.Empty, error)
	// Pipeline rpcs
	CreatePipelineInfo(context.Context, *PipelineInfo) (*PipelineInfo, error)
	GetPipelineInfo(context.Context, *pachyderm_pps.Pipeline) (*PipelineInfo, error)
	// ordered by time, latest to earliest
	ListPipelineInfos(context.Context, *ListPipelineInfosRequest) (*PipelineInfos, error)
	DeletePipelineInfo(context.Context, *pachyderm_pps.Pipeline) (*google_protobuf.Empty, error)
	SubscribePipelineInfos(*SubscribePipelineInfosRequest, API_SubscribePipelineInfosServer) error
	// Shard rpcs
	// Returns the new job info
	StartPod(context.Context, *pachyderm_pps.Job) (*JobInfo, error)
	SucceedPod(context.Context, *pachyderm_pps.Job) (*JobInfo, error)
	FailPod(context.Context, *pachyderm_pps.Job) (*JobInfo, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateJobInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/CreateJobInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateJobInfo(ctx, req.(*JobInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_InspectJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pachyderm_pps.InspectJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).InspectJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/InspectJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).InspectJob(ctx, req.(*pachyderm_pps.InspectJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListJobInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pachyderm_pps.ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListJobInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/ListJobInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListJobInfos(ctx, req.(*pachyderm_pps.ListJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DeleteJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeleteJobInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/DeleteJobInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeleteJobInfo(ctx, req.(*pachyderm_pps.Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateJobOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobOutput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateJobOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/CreateJobOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateJobOutput(ctx, req.(*JobOutput))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateJobState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateJobState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/CreateJobState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateJobState(ctx, req.(*JobState))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreatePipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreatePipelineInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/CreatePipelineInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreatePipelineInfo(ctx, req.(*PipelineInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetPipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pachyderm_pps.Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetPipelineInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/GetPipelineInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetPipelineInfo(ctx, req.(*pachyderm_pps.Pipeline))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListPipelineInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelineInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListPipelineInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/ListPipelineInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListPipelineInfos(ctx, req.(*ListPipelineInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DeletePipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pachyderm_pps.Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeletePipelineInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/DeletePipelineInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeletePipelineInfo(ctx, req.(*pachyderm_pps.Pipeline))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_SubscribePipelineInfos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribePipelineInfosRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).SubscribePipelineInfos(m, &aPISubscribePipelineInfosServer{stream})
}

type API_SubscribePipelineInfosServer interface {
	Send(*PipelineInfoChange) error
	grpc.ServerStream
}

type aPISubscribePipelineInfosServer struct {
	grpc.ServerStream
}

func (x *aPISubscribePipelineInfosServer) Send(m *PipelineInfoChange) error {
	return x.ServerStream.SendMsg(m)
}

func _API_StartPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).StartPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/StartPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).StartPod(ctx, req.(*pachyderm_pps.Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_SucceedPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).SucceedPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/SucceedPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).SucceedPod(ctx, req.(*pachyderm_pps.Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_FailPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).FailPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.persist.API/FailPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).FailPod(ctx, req.(*pachyderm_pps.Job))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.persist.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJobInfo",
			Handler:    _API_CreateJobInfo_Handler,
		},
		{
			MethodName: "InspectJob",
			Handler:    _API_InspectJob_Handler,
		},
		{
			MethodName: "ListJobInfos",
			Handler:    _API_ListJobInfos_Handler,
		},
		{
			MethodName: "DeleteJobInfo",
			Handler:    _API_DeleteJobInfo_Handler,
		},
		{
			MethodName: "CreateJobOutput",
			Handler:    _API_CreateJobOutput_Handler,
		},
		{
			MethodName: "CreateJobState",
			Handler:    _API_CreateJobState_Handler,
		},
		{
			MethodName: "CreatePipelineInfo",
			Handler:    _API_CreatePipelineInfo_Handler,
		},
		{
			MethodName: "GetPipelineInfo",
			Handler:    _API_GetPipelineInfo_Handler,
		},
		{
			MethodName: "ListPipelineInfos",
			Handler:    _API_ListPipelineInfos_Handler,
		},
		{
			MethodName: "DeletePipelineInfo",
			Handler:    _API_DeletePipelineInfo_Handler,
		},
		{
			MethodName: "StartPod",
			Handler:    _API_StartPod_Handler,
		},
		{
			MethodName: "SucceedPod",
			Handler:    _API_SucceedPod_Handler,
		},
		{
			MethodName: "FailPod",
			Handler:    _API_FailPod_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribePipelineInfos",
			Handler:       _API_SubscribePipelineInfos_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 999 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x72, 0x1b, 0x45,
	0x10, 0x8d, 0x6c, 0x5d, 0x5b, 0x17, 0x17, 0x53, 0xc1, 0xd9, 0x12, 0x49, 0x2c, 0x36, 0x50, 0x5c,
	0xaa, 0x58, 0x05, 0x93, 0xa2, 0x8a, 0x07, 0x2a, 0xc4, 0x26, 0x09, 0x02, 0x62, 0x94, 0xb1, 0x5f,
	0xe0, 0x65, 0x59, 0xed, 0x8e, 0xec, 0x4d, 0xed, 0x8d, 0xbd, 0xa4, 0x70, 0x15, 0x7c, 0x00, 0x3f,
	0xc0, 0xf7, 0xf0, 0x69, 0xf4, 0xf4, 0xec, 0xca, 0xb2, 0xa4, 0x95, 0x37, 0x7e, 0x70, 0x59, 0x73,
	0xe6, 0x74, 0xf7, 0xcc, 0xe9, 0x3e, 0x23, 0xc1, 0x28, 0x11, 0xf1, 0x5b, 0x11, 0x8f, 0xa3, 0x28,
	0x19, 0x47, 0x22, 0x4e, 0xdc, 0x24, 0x2d, 0xfe, 0x1b, 0x51, 0x1c, 0xa6, 0x21, 0x7b, 0x3f, 0xb2,
	0xec, 0x8b, 0x4b, 0x47, 0xc4, 0xbe, 0x81, 0x24, 0x23, 0xdf, 0x1c, 0x7e, 0x70, 0x1e, 0x86, 0xe7,
	0x9e, 0x18, 0x13, 0x69, 0x96, 0xcd, 0xc7, 0xc2, 0x8f, 0xd2, 0x4b, 0x15, 0x33, 0x3c, 0x58, 0xdd,
	0x4c, 0x5d, 0x5f, 0x24, 0xa9, 0xe5, 0x47, 0x39, 0xe1, 0xae, 0xed, 0xb9, 0x22, 0xc0, 0x52, 0xf3,
	0x44, 0xfe, 0xad, 0xa2, 0xf2, 0x30, 0x51, 0x8e, 0xea, 0xff, 0x36, 0xa0, 0xf5, 0x63, 0x38, 0x9b,
	0x04, 0x73, 0x3c, 0x0c, 0x34, 0xdf, 0x84, 0x33, 0xd3, 0x75, 0xb4, 0xda, 0xa8, 0xf6, 0x69, 0x87,
	0x37, 0x70, 0x35, 0x71, 0xd8, 0xd7, 0xd0, 0x49, 0x63, 0x2b, 0x48, 0xe6, 0x61, 0xec, 0x6b, 0x3b,
	0xb8, 0xd3, 0x3d, 0xd4, 0x8c, 0xeb, 0xe7, 0x3e, 0x2b, 0xf6, 0xf9, 0x15, 0x95, 0x3d, 0x82, 0x7e,
	0xe4, 0x46, 0xc2, 0x73, 0x03, 0x61, 0x06, 0x96, 0x2f, 0xb4, 0x5d, 0xca, 0xda, 0x2b, 0xc0, 0x13,
	0xc4, 0xd8, 0x08, 0xba, 0x91, 0x15, 0x5b, 0x9e, 0x87, 0x50, 0xe2, 0x6b, 0x75, 0xa4, 0xd4, 0xf9,
	0x32, 0xc4, 0xc6, 0xd0, 0x74, 0x83, 0x28, 0x4b, 0x13, 0xad, 0x31, 0xda, 0xc5, 0xda, 0xf7, 0x56,
	0x6a, 0xd3, 0xe9, 0x71, 0x9f, 0xe7, 0x34, 0xf6, 0x25, 0x00, 0xc6, 0xe3, 0x55, 0x4d, 0x3c, 0xbf,
	0xd6, 0xa4, 0x03, 0xb3, 0xf5, 0x20, 0xde, 0x51, 0x2c, 0xfc, 0xc8, 0xbe, 0x01, 0xb0, 0x63, 0x61,
	0xa5, 0xc2, 0x31, 0xad, 0x54, 0x6b, 0x51, 0xc8, 0xd0, 0x50, 0x3a, 0x1b, 0x85, 0xce, 0xc6, 0x59,
	0xa1, 0x33, 0xef, 0xe4, 0xec, 0x67, 0x29, 0x7b, 0x0c, 0xfd, 0x30, 0x4b, 0xb1, 0xb0, 0x69, 0x87,
	0xbe, 0xef, 0xa6, 0x5a, 0x9b, 0xa2, 0xbb, 0x86, 0x54, 0xfe, 0x98, 0x20, 0xde, 0x53, 0x0c, 0xb5,
	0x62, 0x5f, 0x40, 0x03, 0xb3, 0xa4, 0x42, 0xeb, 0x20, 0x73, 0xb0, 0xe9, 0x3e, 0xa7, 0x72, 0x9b,
	0x2b, 0x16, 0xfb, 0x10, 0x7a, 0x2a, 0xb3, 0xe9, 0x06, 0x8e, 0xf8, 0x53, 0x03, 0x52, 0xb1, 0xab,
	0xb0, 0x89, 0x84, 0x24, 0x25, 0x0a, 0x9d, 0xc4, 0xc4, 0x80, 0x18, 0x4f, 0xa5, 0x75, 0x73, 0x15,
	0x11, 0x3b, 0x55, 0x10, 0xfb, 0x18, 0x06, 0x8a, 0x92, 0xd9, 0xb6, 0x10, 0x0e, 0x92, 0x7a, 0x44,
	0xea, 0x13, 0xa9, 0x00, 0xd9, 0x01, 0x50, 0x94, 0x39, 0xb7, 0x5c, 0x0f, 0x39, 0x7d, 0xe2, 0x80,
	0x84, 0x5e, 0x10, 0xc2, 0x8e, 0xe0, 0x61, 0x10, 0x06, 0x26, 0xcd, 0xa3, 0x39, 0x77, 0xbd, 0x54,
	0xc4, 0x66, 0x72, 0x61, 0xc5, 0x8e, 0x19, 0x64, 0xfe, 0x0c, 0xc7, 0x57, 0x1b, 0x60, 0x97, 0xea,
	0x7c, 0x88, 0xac, 0xe7, 0x92, 0xf4, 0x82, 0x38, 0xa7, 0x92, 0x72, 0xa2, 0x18, 0x72, 0x30, 0x54,
	0x88, 0x1f, 0x3a, 0x99, 0x97, 0x25, 0xda, 0x1e, 0x95, 0xe9, 0x11, 0xf8, 0x4a, 0x61, 0xfa, 0x73,
	0x68, 0xe7, 0x73, 0x99, 0x60, 0x7b, 0xda, 0x34, 0x98, 0xb8, 0xc0, 0xd1, 0x94, 0x43, 0xf0, 0xd0,
	0xd8, 0x68, 0x1c, 0x23, 0x0f, 0xe1, 0xad, 0x37, 0xea, 0x83, 0x7e, 0x06, 0x1d, 0xc4, 0x7e, 0x21,
	0xfd, 0xcb, 0x06, 0x7c, 0xad, 0x85, 0x3b, 0x37, 0xb4, 0x50, 0x9f, 0xd2, 0xe1, 0xa8, 0x4d, 0x65,
	0x49, 0x17, 0x5d, 0xde, 0xa9, 0xd2, 0x65, 0xfd, 0xbf, 0x1d, 0xe8, 0x4d, 0x73, 0x63, 0x90, 0x19,
	0xd7, 0xdc, 0x53, 0xdb, 0xe0, 0x9e, 0xdb, 0x5a, 0x73, 0xc5, 0x75, 0xbb, 0xeb, 0xae, 0x7b, 0xb2,
	0x70, 0x5d, 0x9d, 0x04, 0xbf, 0xbf, 0x92, 0xf6, 0xea, 0xac, 0xcb, 0xd6, 0xfb, 0x1c, 0xba, 0xb9,
	0x92, 0xb1, 0x88, 0x42, 0x34, 0xac, 0x3c, 0x51, 0x87, 0x74, 0xe4, 0x08, 0x70, 0x50, 0xbb, 0xf2,
	0xf3, 0x8a, 0xe7, 0x9a, 0xef, 0xe2, 0xb9, 0xbb, 0xa8, 0xad, 0x9c, 0x15, 0x72, 0x6a, 0x9d, 0xab,
	0x85, 0x1e, 0x02, 0x5b, 0x56, 0xf0, 0xf8, 0xc2, 0x0a, 0xce, 0x05, 0x7b, 0x0a, 0xed, 0x42, 0x32,
	0x92, 0xb0, 0x7b, 0xf8, 0xa8, 0x64, 0x76, 0x96, 0x83, 0xf9, 0x22, 0x88, 0x69, 0xd0, 0x8a, 0x85,
	0x1f, 0xbe, 0x45, 0x3b, 0x48, 0x85, 0xdb, 0xbc, 0x58, 0xea, 0xbf, 0x42, 0x7f, 0x39, 0x26, 0x61,
	0x3f, 0x2c, 0xf5, 0x6c, 0x69, 0x58, 0x2b, 0x15, 0x5c, 0x34, 0x96, 0xc6, 0xf6, 0x2f, 0x78, 0x70,
	0x9a, 0xcd, 0x12, 0x3b, 0x76, 0x67, 0xe2, 0x5a, 0x0d, 0x2e, 0xfe, 0xc8, 0x50, 0x10, 0xf6, 0x09,
	0xec, 0xb9, 0x81, 0xed, 0x65, 0x8e, 0xac, 0xe4, 0xa6, 0xae, 0xe5, 0xd1, 0xed, 0xda, 0x7c, 0x90,
	0xc3, 0x13, 0x85, 0xb2, 0xc3, 0x42, 0x2b, 0x35, 0x1e, 0xf7, 0x4b, 0xce, 0x42, 0x06, 0x2d, 0x94,
	0x3c, 0x01, 0xed, 0x67, 0x04, 0x37, 0x16, 0x5e, 0xe4, 0xab, 0x55, 0xcf, 0x77, 0x00, 0x0d, 0x5a,
	0xb3, 0x7d, 0x68, 0xaa, 0x67, 0x82, 0xa2, 0xeb, 0x3c, 0x5f, 0x1d, 0xfe, 0xd3, 0x81, 0xdd, 0x67,
	0xd3, 0x09, 0x7b, 0x0d, 0xfd, 0x63, 0xea, 0x72, 0xf1, 0x95, 0x74, 0x83, 0xcf, 0x87, 0x37, 0xec,
	0xeb, 0x77, 0xd8, 0x14, 0x60, 0x12, 0x24, 0x91, 0xb0, 0xe9, 0xa1, 0x1f, 0xad, 0xf0, 0xaf, 0xb6,
	0xf2, 0xfb, 0x55, 0xca, 0xd8, 0x93, 0xea, 0x2c, 0x5e, 0xa7, 0x07, 0x2b, 0x11, 0xf9, 0x66, 0x91,
	0xf0, 0x60, 0x7b, 0xc2, 0x04, 0x33, 0x7e, 0x0b, 0xfd, 0xef, 0x85, 0x27, 0xae, 0xae, 0xbd, 0xe1,
	0xeb, 0x6a, 0xb8, 0xbf, 0xe6, 0x0d, 0x7a, 0x5f, 0x31, 0xfc, 0x15, 0xec, 0x2d, 0x54, 0xcb, 0x5f,
	0xba, 0x51, 0x79, 0x51, 0xc5, 0xd8, 0x92, 0xee, 0x27, 0x18, 0x2c, 0xd2, 0xa9, 0x27, 0x6e, 0xcb,
	0x15, 0x88, 0xb0, 0x25, 0xd9, 0xef, 0xc0, 0x54, 0xb2, 0xeb, 0x8f, 0x5b, 0x05, 0x47, 0x0c, 0xab,
	0x90, 0xb0, 0xc2, 0x6b, 0xd8, 0x7b, 0x29, 0xae, 0xcd, 0x2a, 0xbb, 0x57, 0xf2, 0x58, 0x55, 0x4d,
	0xe9, 0xc1, 0x7b, 0x6b, 0xf3, 0xcf, 0xc6, 0x25, 0xb1, 0x65, 0x4e, 0x19, 0x7e, 0x54, 0xa1, 0x98,
	0xec, 0xfe, 0x4b, 0x60, 0xaa, 0xfb, 0xd5, 0xee, 0x50, 0xae, 0xf5, 0xdf, 0xb0, 0xbf, 0xf9, 0xd1,
	0x60, 0x4f, 0xca, 0x5c, 0xba, 0xed, 0x8d, 0x19, 0x7e, 0x56, 0xe1, 0x02, 0xea, 0x95, 0xd5, 0xef,
	0x3c, 0xae, 0xb1, 0xef, 0xa0, 0x4d, 0xbf, 0x36, 0xa6, 0xa1, 0xb3, 0x71, 0x80, 0x6f, 0x76, 0xd6,
	0x11, 0x40, 0xfe, 0x53, 0xe4, 0xf6, 0x39, 0x9e, 0x42, 0x4b, 0xfe, 0x54, 0xb9, 0x75, 0x82, 0xa3,
	0xce, 0x6f, 0xad, 0x1c, 0x9c, 0x35, 0x49, 0xe2, 0xaf, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xbe,
	0xe2, 0x92, 0xaa, 0xc8, 0x0b, 0x00, 0x00,
}
