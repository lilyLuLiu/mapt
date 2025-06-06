// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets load balancer backend address pool.
//
// Uses Azure REST API version 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
func LookupLoadBalancerBackendAddressPool(ctx *pulumi.Context, args *LookupLoadBalancerBackendAddressPoolArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerBackendAddressPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLoadBalancerBackendAddressPoolResult
	err := ctx.Invoke("azure-native:network:getLoadBalancerBackendAddressPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadBalancerBackendAddressPoolArgs struct {
	// The name of the backend address pool.
	BackendAddressPoolName string `pulumi:"backendAddressPoolName"`
	// The name of the load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Pool of backend IP addresses.
type LookupLoadBalancerBackendAddressPoolResult struct {
	// An array of references to IP addresses defined in network interfaces.
	BackendIPConfigurations []NetworkInterfaceIPConfigurationResponse `pulumi:"backendIPConfigurations"`
	// Amount of seconds Load Balancer waits for before sending RESET to client and backend address.
	DrainPeriodInSeconds *int `pulumi:"drainPeriodInSeconds"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// An array of references to inbound NAT rules that use this backend address pool.
	InboundNatRules []SubResourceResponse `pulumi:"inboundNatRules"`
	// An array of backend addresses.
	LoadBalancerBackendAddresses []LoadBalancerBackendAddressResponse `pulumi:"loadBalancerBackendAddresses"`
	// An array of references to load balancing rules that use this backend address pool.
	LoadBalancingRules []SubResourceResponse `pulumi:"loadBalancingRules"`
	// The location of the backend address pool.
	Location *string `pulumi:"location"`
	// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// A reference to an outbound rule that uses this backend address pool.
	OutboundRule SubResourceResponse `pulumi:"outboundRule"`
	// An array of references to outbound rules that use this backend address pool.
	OutboundRules []SubResourceResponse `pulumi:"outboundRules"`
	// The provisioning state of the backend address pool resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// An array of gateway load balancer tunnel interfaces.
	TunnelInterfaces []GatewayLoadBalancerTunnelInterfaceResponse `pulumi:"tunnelInterfaces"`
	// Type of the resource.
	Type string `pulumi:"type"`
	// A reference to a virtual network.
	VirtualNetwork *SubResourceResponse `pulumi:"virtualNetwork"`
}

func LookupLoadBalancerBackendAddressPoolOutput(ctx *pulumi.Context, args LookupLoadBalancerBackendAddressPoolOutputArgs, opts ...pulumi.InvokeOption) LookupLoadBalancerBackendAddressPoolResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLoadBalancerBackendAddressPoolResultOutput, error) {
			args := v.(LookupLoadBalancerBackendAddressPoolArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getLoadBalancerBackendAddressPool", args, LookupLoadBalancerBackendAddressPoolResultOutput{}, options).(LookupLoadBalancerBackendAddressPoolResultOutput), nil
		}).(LookupLoadBalancerBackendAddressPoolResultOutput)
}

type LookupLoadBalancerBackendAddressPoolOutputArgs struct {
	// The name of the backend address pool.
	BackendAddressPoolName pulumi.StringInput `pulumi:"backendAddressPoolName"`
	// The name of the load balancer.
	LoadBalancerName pulumi.StringInput `pulumi:"loadBalancerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLoadBalancerBackendAddressPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerBackendAddressPoolArgs)(nil)).Elem()
}

// Pool of backend IP addresses.
type LookupLoadBalancerBackendAddressPoolResultOutput struct{ *pulumi.OutputState }

func (LookupLoadBalancerBackendAddressPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerBackendAddressPoolResult)(nil)).Elem()
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) ToLookupLoadBalancerBackendAddressPoolResultOutput() LookupLoadBalancerBackendAddressPoolResultOutput {
	return o
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) ToLookupLoadBalancerBackendAddressPoolResultOutputWithContext(ctx context.Context) LookupLoadBalancerBackendAddressPoolResultOutput {
	return o
}

// An array of references to IP addresses defined in network interfaces.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) BackendIPConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) []NetworkInterfaceIPConfigurationResponse {
		return v.BackendIPConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

// Amount of seconds Load Balancer waits for before sending RESET to client and backend address.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) DrainPeriodInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) *int { return v.DrainPeriodInSeconds }).(pulumi.IntPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// An array of references to inbound NAT rules that use this backend address pool.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) InboundNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) []SubResourceResponse { return v.InboundNatRules }).(SubResourceResponseArrayOutput)
}

// An array of backend addresses.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) LoadBalancerBackendAddresses() LoadBalancerBackendAddressResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) []LoadBalancerBackendAddressResponse {
		return v.LoadBalancerBackendAddresses
	}).(LoadBalancerBackendAddressResponseArrayOutput)
}

// An array of references to load balancing rules that use this backend address pool.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) LoadBalancingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) []SubResourceResponse { return v.LoadBalancingRules }).(SubResourceResponseArrayOutput)
}

// The location of the backend address pool.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A reference to an outbound rule that uses this backend address pool.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) OutboundRule() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) SubResourceResponse { return v.OutboundRule }).(SubResourceResponseOutput)
}

// An array of references to outbound rules that use this backend address pool.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) OutboundRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) []SubResourceResponse { return v.OutboundRules }).(SubResourceResponseArrayOutput)
}

// The provisioning state of the backend address pool resource.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// An array of gateway load balancer tunnel interfaces.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) TunnelInterfaces() GatewayLoadBalancerTunnelInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) []GatewayLoadBalancerTunnelInterfaceResponse {
		return v.TunnelInterfaces
	}).(GatewayLoadBalancerTunnelInterfaceResponseArrayOutput)
}

// Type of the resource.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

// A reference to a virtual network.
func (o LookupLoadBalancerBackendAddressPoolResultOutput) VirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) *SubResourceResponse { return v.VirtualNetwork }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadBalancerBackendAddressPoolResultOutput{})
}
