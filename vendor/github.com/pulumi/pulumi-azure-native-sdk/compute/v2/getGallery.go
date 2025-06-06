// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about a Shared Image Gallery.
//
// Uses Azure REST API version 2022-03-03.
//
// Other available API versions: 2022-08-03, 2023-07-03, 2024-03-03.
func LookupGallery(ctx *pulumi.Context, args *LookupGalleryArgs, opts ...pulumi.InvokeOption) (*LookupGalleryResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGalleryResult
	err := ctx.Invoke("azure-native:compute:getGallery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryArgs struct {
	// The expand query option to apply on the operation.
	Expand *string `pulumi:"expand"`
	// The name of the Shared Image Gallery.
	GalleryName string `pulumi:"galleryName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The select expression to apply on the operation.
	Select *string `pulumi:"select"`
}

// Specifies information about the Shared Image Gallery that you want to create or update.
type LookupGalleryResult struct {
	// The description of this Shared Image Gallery resource. This property is updatable.
	Description *string `pulumi:"description"`
	// Resource Id
	Id string `pulumi:"id"`
	// Describes the gallery unique name.
	Identifier *GalleryIdentifierResponse `pulumi:"identifier"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// The provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// Profile for gallery sharing to subscription or tenant
	SharingProfile *SharingProfileResponse `pulumi:"sharingProfile"`
	// Sharing status of current gallery.
	SharingStatus SharingStatusResponse `pulumi:"sharingStatus"`
	// Contains information about the soft deletion policy of the gallery.
	SoftDeletePolicy *SoftDeletePolicyResponse `pulumi:"softDeletePolicy"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupGalleryOutput(ctx *pulumi.Context, args LookupGalleryOutputArgs, opts ...pulumi.InvokeOption) LookupGalleryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGalleryResultOutput, error) {
			args := v.(LookupGalleryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:compute:getGallery", args, LookupGalleryResultOutput{}, options).(LookupGalleryResultOutput), nil
		}).(LookupGalleryResultOutput)
}

type LookupGalleryOutputArgs struct {
	// The expand query option to apply on the operation.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the Shared Image Gallery.
	GalleryName pulumi.StringInput `pulumi:"galleryName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The select expression to apply on the operation.
	Select pulumi.StringPtrInput `pulumi:"select"`
}

func (LookupGalleryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryArgs)(nil)).Elem()
}

// Specifies information about the Shared Image Gallery that you want to create or update.
type LookupGalleryResultOutput struct{ *pulumi.OutputState }

func (LookupGalleryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryResult)(nil)).Elem()
}

func (o LookupGalleryResultOutput) ToLookupGalleryResultOutput() LookupGalleryResultOutput {
	return o
}

func (o LookupGalleryResultOutput) ToLookupGalleryResultOutputWithContext(ctx context.Context) LookupGalleryResultOutput {
	return o
}

// The description of this Shared Image Gallery resource. This property is updatable.
func (o LookupGalleryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupGalleryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Id }).(pulumi.StringOutput)
}

// Describes the gallery unique name.
func (o LookupGalleryResultOutput) Identifier() GalleryIdentifierResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryResult) *GalleryIdentifierResponse { return v.Identifier }).(GalleryIdentifierResponsePtrOutput)
}

// Resource location
func (o LookupGalleryResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupGalleryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state, which only appears in the response.
func (o LookupGalleryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Profile for gallery sharing to subscription or tenant
func (o LookupGalleryResultOutput) SharingProfile() SharingProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryResult) *SharingProfileResponse { return v.SharingProfile }).(SharingProfileResponsePtrOutput)
}

// Sharing status of current gallery.
func (o LookupGalleryResultOutput) SharingStatus() SharingStatusResponseOutput {
	return o.ApplyT(func(v LookupGalleryResult) SharingStatusResponse { return v.SharingStatus }).(SharingStatusResponseOutput)
}

// Contains information about the soft deletion policy of the gallery.
func (o LookupGalleryResultOutput) SoftDeletePolicy() SoftDeletePolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryResult) *SoftDeletePolicyResponse { return v.SoftDeletePolicy }).(SoftDeletePolicyResponsePtrOutput)
}

// Resource tags
func (o LookupGalleryResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGalleryResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupGalleryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryResultOutput{})
}
