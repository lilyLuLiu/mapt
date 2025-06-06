// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified custom IP prefix in a specified resource group.
//
// Uses Azure REST API version 2023-02-01.
//
// Other available API versions: 2021-03-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
func LookupCustomIPPrefix(ctx *pulumi.Context, args *LookupCustomIPPrefixArgs, opts ...pulumi.InvokeOption) (*LookupCustomIPPrefixResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomIPPrefixResult
	err := ctx.Invoke("azure-native:network:getCustomIPPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomIPPrefixArgs struct {
	// The name of the custom IP prefix.
	CustomIpPrefixName string `pulumi:"customIpPrefixName"`
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Custom IP prefix resource.
type LookupCustomIPPrefixResult struct {
	// The ASN for CIDR advertising. Should be an integer as string.
	Asn *string `pulumi:"asn"`
	// Authorization message for WAN validation.
	AuthorizationMessage *string `pulumi:"authorizationMessage"`
	// The list of all Children for IPv6 /48 CustomIpPrefix.
	ChildCustomIpPrefixes []SubResourceResponse `pulumi:"childCustomIpPrefixes"`
	// The prefix range in CIDR notation. Should include the start address and the prefix length.
	Cidr *string `pulumi:"cidr"`
	// The commissioned state of the Custom IP Prefix.
	CommissionedState *string `pulumi:"commissionedState"`
	// The Parent CustomIpPrefix for IPv6 /64 CustomIpPrefix.
	CustomIpPrefixParent *SubResourceResponse `pulumi:"customIpPrefixParent"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Whether to do express route advertise.
	ExpressRouteAdvertise *bool `pulumi:"expressRouteAdvertise"`
	// The extended location of the custom IP prefix.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// The reason why resource is in failed state.
	FailedReason string `pulumi:"failedReason"`
	// The Geo for CIDR advertising. Should be an Geo code.
	Geo *string `pulumi:"geo"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Whether to Advertise the range to Internet.
	NoInternetAdvertise *bool `pulumi:"noInternetAdvertise"`
	// Type of custom IP prefix. Should be Singular, Parent, or Child.
	PrefixType *string `pulumi:"prefixType"`
	// The provisioning state of the custom IP prefix resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The list of all referenced PublicIpPrefixes.
	PublicIpPrefixes []SubResourceResponse `pulumi:"publicIpPrefixes"`
	// The resource GUID property of the custom IP prefix resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Signed message for WAN validation.
	SignedMessage *string `pulumi:"signedMessage"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

func LookupCustomIPPrefixOutput(ctx *pulumi.Context, args LookupCustomIPPrefixOutputArgs, opts ...pulumi.InvokeOption) LookupCustomIPPrefixResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCustomIPPrefixResultOutput, error) {
			args := v.(LookupCustomIPPrefixArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getCustomIPPrefix", args, LookupCustomIPPrefixResultOutput{}, options).(LookupCustomIPPrefixResultOutput), nil
		}).(LookupCustomIPPrefixResultOutput)
}

type LookupCustomIPPrefixOutputArgs struct {
	// The name of the custom IP prefix.
	CustomIpPrefixName pulumi.StringInput `pulumi:"customIpPrefixName"`
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCustomIPPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomIPPrefixArgs)(nil)).Elem()
}

// Custom IP prefix resource.
type LookupCustomIPPrefixResultOutput struct{ *pulumi.OutputState }

func (LookupCustomIPPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomIPPrefixResult)(nil)).Elem()
}

func (o LookupCustomIPPrefixResultOutput) ToLookupCustomIPPrefixResultOutput() LookupCustomIPPrefixResultOutput {
	return o
}

func (o LookupCustomIPPrefixResultOutput) ToLookupCustomIPPrefixResultOutputWithContext(ctx context.Context) LookupCustomIPPrefixResultOutput {
	return o
}

// The ASN for CIDR advertising. Should be an integer as string.
func (o LookupCustomIPPrefixResultOutput) Asn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Asn }).(pulumi.StringPtrOutput)
}

// Authorization message for WAN validation.
func (o LookupCustomIPPrefixResultOutput) AuthorizationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.AuthorizationMessage }).(pulumi.StringPtrOutput)
}

// The list of all Children for IPv6 /48 CustomIpPrefix.
func (o LookupCustomIPPrefixResultOutput) ChildCustomIpPrefixes() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) []SubResourceResponse { return v.ChildCustomIpPrefixes }).(SubResourceResponseArrayOutput)
}

// The prefix range in CIDR notation. Should include the start address and the prefix length.
func (o LookupCustomIPPrefixResultOutput) Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Cidr }).(pulumi.StringPtrOutput)
}

// The commissioned state of the Custom IP Prefix.
func (o LookupCustomIPPrefixResultOutput) CommissionedState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.CommissionedState }).(pulumi.StringPtrOutput)
}

// The Parent CustomIpPrefix for IPv6 /64 CustomIpPrefix.
func (o LookupCustomIPPrefixResultOutput) CustomIpPrefixParent() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *SubResourceResponse { return v.CustomIpPrefixParent }).(SubResourceResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupCustomIPPrefixResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Whether to do express route advertise.
func (o LookupCustomIPPrefixResultOutput) ExpressRouteAdvertise() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *bool { return v.ExpressRouteAdvertise }).(pulumi.BoolPtrOutput)
}

// The extended location of the custom IP prefix.
func (o LookupCustomIPPrefixResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The reason why resource is in failed state.
func (o LookupCustomIPPrefixResultOutput) FailedReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.FailedReason }).(pulumi.StringOutput)
}

// The Geo for CIDR advertising. Should be an Geo code.
func (o LookupCustomIPPrefixResultOutput) Geo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Geo }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupCustomIPPrefixResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupCustomIPPrefixResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupCustomIPPrefixResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether to Advertise the range to Internet.
func (o LookupCustomIPPrefixResultOutput) NoInternetAdvertise() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *bool { return v.NoInternetAdvertise }).(pulumi.BoolPtrOutput)
}

// Type of custom IP prefix. Should be Singular, Parent, or Child.
func (o LookupCustomIPPrefixResultOutput) PrefixType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.PrefixType }).(pulumi.StringPtrOutput)
}

// The provisioning state of the custom IP prefix resource.
func (o LookupCustomIPPrefixResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of all referenced PublicIpPrefixes.
func (o LookupCustomIPPrefixResultOutput) PublicIpPrefixes() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) []SubResourceResponse { return v.PublicIpPrefixes }).(SubResourceResponseArrayOutput)
}

// The resource GUID property of the custom IP prefix resource.
func (o LookupCustomIPPrefixResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Signed message for WAN validation.
func (o LookupCustomIPPrefixResultOutput) SignedMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.SignedMessage }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupCustomIPPrefixResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupCustomIPPrefixResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of availability zones denoting the IP allocated for the resource needs to come from.
func (o LookupCustomIPPrefixResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomIPPrefixResultOutput{})
}
