// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package efsiface provides an interface to enable mocking the Amazon Elastic File System service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package efsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/efs"
)

// EFSAPI provides an interface to enable mocking the
// efs.EFS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Elastic File System.
//    func myFunc(svc efsiface.EFSAPI) bool {
//        // Make svc.CreateAccessPoint request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := efs.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockEFSClient struct {
//        efsiface.EFSAPI
//    }
//    func (m *mockEFSClient) CreateAccessPoint(input *efs.CreateAccessPointInput) (*efs.CreateAccessPointOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockEFSClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type EFSAPI interface {
	CreateAccessPoint(*efs.CreateAccessPointInput) (*efs.CreateAccessPointOutput, error)
	CreateAccessPointWithContext(aws.Context, *efs.CreateAccessPointInput, ...request.Option) (*efs.CreateAccessPointOutput, error)
	CreateAccessPointRequest(*efs.CreateAccessPointInput) (*request.Request, *efs.CreateAccessPointOutput)

	CreateFileSystem(*efs.CreateFileSystemInput) (*efs.FileSystemDescription, error)
	CreateFileSystemWithContext(aws.Context, *efs.CreateFileSystemInput, ...request.Option) (*efs.FileSystemDescription, error)
	CreateFileSystemRequest(*efs.CreateFileSystemInput) (*request.Request, *efs.FileSystemDescription)

	CreateMountTarget(*efs.CreateMountTargetInput) (*efs.MountTargetDescription, error)
	CreateMountTargetWithContext(aws.Context, *efs.CreateMountTargetInput, ...request.Option) (*efs.MountTargetDescription, error)
	CreateMountTargetRequest(*efs.CreateMountTargetInput) (*request.Request, *efs.MountTargetDescription)

	CreateTags(*efs.CreateTagsInput) (*efs.CreateTagsOutput, error)
	CreateTagsWithContext(aws.Context, *efs.CreateTagsInput, ...request.Option) (*efs.CreateTagsOutput, error)
	CreateTagsRequest(*efs.CreateTagsInput) (*request.Request, *efs.CreateTagsOutput)

	DeleteAccessPoint(*efs.DeleteAccessPointInput) (*efs.DeleteAccessPointOutput, error)
	DeleteAccessPointWithContext(aws.Context, *efs.DeleteAccessPointInput, ...request.Option) (*efs.DeleteAccessPointOutput, error)
	DeleteAccessPointRequest(*efs.DeleteAccessPointInput) (*request.Request, *efs.DeleteAccessPointOutput)

	DeleteFileSystem(*efs.DeleteFileSystemInput) (*efs.DeleteFileSystemOutput, error)
	DeleteFileSystemWithContext(aws.Context, *efs.DeleteFileSystemInput, ...request.Option) (*efs.DeleteFileSystemOutput, error)
	DeleteFileSystemRequest(*efs.DeleteFileSystemInput) (*request.Request, *efs.DeleteFileSystemOutput)

	DeleteFileSystemPolicy(*efs.DeleteFileSystemPolicyInput) (*efs.DeleteFileSystemPolicyOutput, error)
	DeleteFileSystemPolicyWithContext(aws.Context, *efs.DeleteFileSystemPolicyInput, ...request.Option) (*efs.DeleteFileSystemPolicyOutput, error)
	DeleteFileSystemPolicyRequest(*efs.DeleteFileSystemPolicyInput) (*request.Request, *efs.DeleteFileSystemPolicyOutput)

	DeleteMountTarget(*efs.DeleteMountTargetInput) (*efs.DeleteMountTargetOutput, error)
	DeleteMountTargetWithContext(aws.Context, *efs.DeleteMountTargetInput, ...request.Option) (*efs.DeleteMountTargetOutput, error)
	DeleteMountTargetRequest(*efs.DeleteMountTargetInput) (*request.Request, *efs.DeleteMountTargetOutput)

	DeleteTags(*efs.DeleteTagsInput) (*efs.DeleteTagsOutput, error)
	DeleteTagsWithContext(aws.Context, *efs.DeleteTagsInput, ...request.Option) (*efs.DeleteTagsOutput, error)
	DeleteTagsRequest(*efs.DeleteTagsInput) (*request.Request, *efs.DeleteTagsOutput)

	DescribeAccessPoints(*efs.DescribeAccessPointsInput) (*efs.DescribeAccessPointsOutput, error)
	DescribeAccessPointsWithContext(aws.Context, *efs.DescribeAccessPointsInput, ...request.Option) (*efs.DescribeAccessPointsOutput, error)
	DescribeAccessPointsRequest(*efs.DescribeAccessPointsInput) (*request.Request, *efs.DescribeAccessPointsOutput)

	DescribeAccessPointsPages(*efs.DescribeAccessPointsInput, func(*efs.DescribeAccessPointsOutput, bool) bool) error
	DescribeAccessPointsPagesWithContext(aws.Context, *efs.DescribeAccessPointsInput, func(*efs.DescribeAccessPointsOutput, bool) bool, ...request.Option) error

	DescribeFileSystemPolicy(*efs.DescribeFileSystemPolicyInput) (*efs.DescribeFileSystemPolicyOutput, error)
	DescribeFileSystemPolicyWithContext(aws.Context, *efs.DescribeFileSystemPolicyInput, ...request.Option) (*efs.DescribeFileSystemPolicyOutput, error)
	DescribeFileSystemPolicyRequest(*efs.DescribeFileSystemPolicyInput) (*request.Request, *efs.DescribeFileSystemPolicyOutput)

	DescribeFileSystems(*efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error)
	DescribeFileSystemsWithContext(aws.Context, *efs.DescribeFileSystemsInput, ...request.Option) (*efs.DescribeFileSystemsOutput, error)
	DescribeFileSystemsRequest(*efs.DescribeFileSystemsInput) (*request.Request, *efs.DescribeFileSystemsOutput)

	DescribeFileSystemsPages(*efs.DescribeFileSystemsInput, func(*efs.DescribeFileSystemsOutput, bool) bool) error
	DescribeFileSystemsPagesWithContext(aws.Context, *efs.DescribeFileSystemsInput, func(*efs.DescribeFileSystemsOutput, bool) bool, ...request.Option) error

	DescribeLifecycleConfiguration(*efs.DescribeLifecycleConfigurationInput) (*efs.DescribeLifecycleConfigurationOutput, error)
	DescribeLifecycleConfigurationWithContext(aws.Context, *efs.DescribeLifecycleConfigurationInput, ...request.Option) (*efs.DescribeLifecycleConfigurationOutput, error)
	DescribeLifecycleConfigurationRequest(*efs.DescribeLifecycleConfigurationInput) (*request.Request, *efs.DescribeLifecycleConfigurationOutput)

	DescribeMountTargetSecurityGroups(*efs.DescribeMountTargetSecurityGroupsInput) (*efs.DescribeMountTargetSecurityGroupsOutput, error)
	DescribeMountTargetSecurityGroupsWithContext(aws.Context, *efs.DescribeMountTargetSecurityGroupsInput, ...request.Option) (*efs.DescribeMountTargetSecurityGroupsOutput, error)
	DescribeMountTargetSecurityGroupsRequest(*efs.DescribeMountTargetSecurityGroupsInput) (*request.Request, *efs.DescribeMountTargetSecurityGroupsOutput)

	DescribeMountTargets(*efs.DescribeMountTargetsInput) (*efs.DescribeMountTargetsOutput, error)
	DescribeMountTargetsWithContext(aws.Context, *efs.DescribeMountTargetsInput, ...request.Option) (*efs.DescribeMountTargetsOutput, error)
	DescribeMountTargetsRequest(*efs.DescribeMountTargetsInput) (*request.Request, *efs.DescribeMountTargetsOutput)

	DescribeTags(*efs.DescribeTagsInput) (*efs.DescribeTagsOutput, error)
	DescribeTagsWithContext(aws.Context, *efs.DescribeTagsInput, ...request.Option) (*efs.DescribeTagsOutput, error)
	DescribeTagsRequest(*efs.DescribeTagsInput) (*request.Request, *efs.DescribeTagsOutput)

	DescribeTagsPages(*efs.DescribeTagsInput, func(*efs.DescribeTagsOutput, bool) bool) error
	DescribeTagsPagesWithContext(aws.Context, *efs.DescribeTagsInput, func(*efs.DescribeTagsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*efs.ListTagsForResourceInput) (*efs.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *efs.ListTagsForResourceInput, ...request.Option) (*efs.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*efs.ListTagsForResourceInput) (*request.Request, *efs.ListTagsForResourceOutput)

	ListTagsForResourcePages(*efs.ListTagsForResourceInput, func(*efs.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *efs.ListTagsForResourceInput, func(*efs.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	ModifyMountTargetSecurityGroups(*efs.ModifyMountTargetSecurityGroupsInput) (*efs.ModifyMountTargetSecurityGroupsOutput, error)
	ModifyMountTargetSecurityGroupsWithContext(aws.Context, *efs.ModifyMountTargetSecurityGroupsInput, ...request.Option) (*efs.ModifyMountTargetSecurityGroupsOutput, error)
	ModifyMountTargetSecurityGroupsRequest(*efs.ModifyMountTargetSecurityGroupsInput) (*request.Request, *efs.ModifyMountTargetSecurityGroupsOutput)

	PutFileSystemPolicy(*efs.PutFileSystemPolicyInput) (*efs.PutFileSystemPolicyOutput, error)
	PutFileSystemPolicyWithContext(aws.Context, *efs.PutFileSystemPolicyInput, ...request.Option) (*efs.PutFileSystemPolicyOutput, error)
	PutFileSystemPolicyRequest(*efs.PutFileSystemPolicyInput) (*request.Request, *efs.PutFileSystemPolicyOutput)

	PutLifecycleConfiguration(*efs.PutLifecycleConfigurationInput) (*efs.PutLifecycleConfigurationOutput, error)
	PutLifecycleConfigurationWithContext(aws.Context, *efs.PutLifecycleConfigurationInput, ...request.Option) (*efs.PutLifecycleConfigurationOutput, error)
	PutLifecycleConfigurationRequest(*efs.PutLifecycleConfigurationInput) (*request.Request, *efs.PutLifecycleConfigurationOutput)

	TagResource(*efs.TagResourceInput) (*efs.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *efs.TagResourceInput, ...request.Option) (*efs.TagResourceOutput, error)
	TagResourceRequest(*efs.TagResourceInput) (*request.Request, *efs.TagResourceOutput)

	UntagResource(*efs.UntagResourceInput) (*efs.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *efs.UntagResourceInput, ...request.Option) (*efs.UntagResourceOutput, error)
	UntagResourceRequest(*efs.UntagResourceInput) (*request.Request, *efs.UntagResourceOutput)

	UpdateFileSystem(*efs.UpdateFileSystemInput) (*efs.UpdateFileSystemOutput, error)
	UpdateFileSystemWithContext(aws.Context, *efs.UpdateFileSystemInput, ...request.Option) (*efs.UpdateFileSystemOutput, error)
	UpdateFileSystemRequest(*efs.UpdateFileSystemInput) (*request.Request, *efs.UpdateFileSystemOutput)
}

var _ EFSAPI = (*efs.EFS)(nil)
