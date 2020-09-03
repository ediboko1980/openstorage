package sdk

import (
	"context"

	"github.com/libopenstorage/openstorage/api"
	"github.com/libopenstorage/openstorage/api/errors"
	"github.com/libopenstorage/openstorage/cluster"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeDrainServer struct {
	server serverAccessor
}

func (nd *NodeDrainServer) cluster() cluster.Cluster {
	return nd.server.cluster()
}

func (nd *NodeDrainServer) DrainAttachments(
	ctx context.Context,
	req *api.SdkNodeDrainAttachmentsRequest,
) (*api.SdkNodeDrainResponse, error) {
	if nd.cluster() == nil {
		return nil, status.Error(codes.Unavailable, errors.ErrResourceNotInitialized.Error())
	}

	return nd.cluster().DrainAttachments(ctx, req)

}

func (nd *NodeDrainServer) UpdateNodeDrainJobState(
	ctx context.Context,
	req *api.SdkUpdateNodeDrainJobRequest,
) (*api.SdkUpdateNodeDrainJobResponse, error) {
	if nd.cluster() == nil {
		return nil, status.Error(codes.Unavailable, errors.ErrResourceNotInitialized.Error())
	}

	return nd.cluster().UpdateNodeDrainJobState(ctx, req)
}

func (nd *NodeDrainServer) GetDrainStatus(
	ctx context.Context,
	req *api.SdkGetNodeDrainJobStatusRequest,
) (*api.SdkGetNodeDrainJobStatusResponse, error) {
	if nd.cluster() == nil {
		return nil, status.Error(codes.Unavailable, errors.ErrResourceNotInitialized.Error())
	}

	return nd.cluster().GetDrainStatus(ctx, req)
}

func (nd *NodeDrainServer) EnumerateNodeDrainJobs(
	ctx context.Context,
	req *api.SdkEnumerateNodeDrainJobsRequest,
) (*api.SdkEnumerateNodeDrainJobsResponse, error) {
	if nd.cluster() == nil {
		return nil, status.Error(codes.Unavailable, errors.ErrResourceNotInitialized.Error())
	}

	return nd.cluster().EnumerateNodeDrainJobs(ctx, req)
}
