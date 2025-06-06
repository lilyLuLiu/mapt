// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about a dedicated host.
//
// Uses Azure REST API version 2023-03-01.
//
// Other available API versions: 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01, 2024-11-01.
func LookupDedicatedHost(ctx *pulumi.Context, args *LookupDedicatedHostArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHostResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDedicatedHostResult
	err := ctx.Invoke("azure-native:compute:getDedicatedHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedHostArgs struct {
	// The expand expression to apply on the operation. 'InstanceView' will retrieve the list of instance views of the dedicated host. 'UserData' is not supported for dedicated host.
	Expand *string `pulumi:"expand"`
	// The name of the dedicated host group.
	HostGroupName string `pulumi:"hostGroupName"`
	// The name of the dedicated host.
	HostName string `pulumi:"hostName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Specifies information about the Dedicated host.
type LookupDedicatedHostResult struct {
	// Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
	AutoReplaceOnFailure *bool `pulumi:"autoReplaceOnFailure"`
	// A unique id generated and assigned to the dedicated host by the platform. Does not change throughout the lifetime of the host.
	HostId string `pulumi:"hostId"`
	// Resource Id
	Id string `pulumi:"id"`
	// The dedicated host instance view.
	InstanceView DedicatedHostInstanceViewResponse `pulumi:"instanceView"`
	// Specifies the software license type that will be applied to the VMs deployed on the dedicated host. Possible values are: **None,** **Windows_Server_Hybrid,** **Windows_Server_Perpetual.** The default value is: **None.**
	LicenseType *string `pulumi:"licenseType"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Fault domain of the dedicated host within a dedicated host group.
	PlatformFaultDomain *int `pulumi:"platformFaultDomain"`
	// The provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// The date when the host was first provisioned.
	ProvisioningTime string `pulumi:"provisioningTime"`
	// SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
	Sku SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Specifies the time at which the Dedicated Host resource was created. Minimum api-version: 2021-11-01.
	TimeCreated string `pulumi:"timeCreated"`
	// Resource type
	Type string `pulumi:"type"`
	// A list of references to all virtual machines in the Dedicated Host.
	VirtualMachines []SubResourceReadOnlyResponse `pulumi:"virtualMachines"`
}

func LookupDedicatedHostOutput(ctx *pulumi.Context, args LookupDedicatedHostOutputArgs, opts ...pulumi.InvokeOption) LookupDedicatedHostResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDedicatedHostResultOutput, error) {
			args := v.(LookupDedicatedHostArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:compute:getDedicatedHost", args, LookupDedicatedHostResultOutput{}, options).(LookupDedicatedHostResultOutput), nil
		}).(LookupDedicatedHostResultOutput)
}

type LookupDedicatedHostOutputArgs struct {
	// The expand expression to apply on the operation. 'InstanceView' will retrieve the list of instance views of the dedicated host. 'UserData' is not supported for dedicated host.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the dedicated host group.
	HostGroupName pulumi.StringInput `pulumi:"hostGroupName"`
	// The name of the dedicated host.
	HostName pulumi.StringInput `pulumi:"hostName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDedicatedHostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHostArgs)(nil)).Elem()
}

// Specifies information about the Dedicated host.
type LookupDedicatedHostResultOutput struct{ *pulumi.OutputState }

func (LookupDedicatedHostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHostResult)(nil)).Elem()
}

func (o LookupDedicatedHostResultOutput) ToLookupDedicatedHostResultOutput() LookupDedicatedHostResultOutput {
	return o
}

func (o LookupDedicatedHostResultOutput) ToLookupDedicatedHostResultOutputWithContext(ctx context.Context) LookupDedicatedHostResultOutput {
	return o
}

// Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
func (o LookupDedicatedHostResultOutput) AutoReplaceOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) *bool { return v.AutoReplaceOnFailure }).(pulumi.BoolPtrOutput)
}

// A unique id generated and assigned to the dedicated host by the platform. Does not change throughout the lifetime of the host.
func (o LookupDedicatedHostResultOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.HostId }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupDedicatedHostResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.Id }).(pulumi.StringOutput)
}

// The dedicated host instance view.
func (o LookupDedicatedHostResultOutput) InstanceView() DedicatedHostInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) DedicatedHostInstanceViewResponse { return v.InstanceView }).(DedicatedHostInstanceViewResponseOutput)
}

// Specifies the software license type that will be applied to the VMs deployed on the dedicated host. Possible values are: **None,** **Windows_Server_Hybrid,** **Windows_Server_Perpetual.** The default value is: **None.**
func (o LookupDedicatedHostResultOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Resource location
func (o LookupDedicatedHostResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupDedicatedHostResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.Name }).(pulumi.StringOutput)
}

// Fault domain of the dedicated host within a dedicated host group.
func (o LookupDedicatedHostResultOutput) PlatformFaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) *int { return v.PlatformFaultDomain }).(pulumi.IntPtrOutput)
}

// The provisioning state, which only appears in the response.
func (o LookupDedicatedHostResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The date when the host was first provisioned.
func (o LookupDedicatedHostResultOutput) ProvisioningTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.ProvisioningTime }).(pulumi.StringOutput)
}

// SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
func (o LookupDedicatedHostResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// Resource tags
func (o LookupDedicatedHostResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the time at which the Dedicated Host resource was created. Minimum api-version: 2021-11-01.
func (o LookupDedicatedHostResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

// Resource type
func (o LookupDedicatedHostResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of references to all virtual machines in the Dedicated Host.
func (o LookupDedicatedHostResultOutput) VirtualMachines() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) []SubResourceReadOnlyResponse { return v.VirtualMachines }).(SubResourceReadOnlyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDedicatedHostResultOutput{})
}
