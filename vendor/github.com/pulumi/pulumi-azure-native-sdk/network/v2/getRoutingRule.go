// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a network manager routing configuration routing rule.
// Azure REST API version: 2024-03-01.
func LookupRoutingRule(ctx *pulumi.Context, args *LookupRoutingRuleArgs, opts ...pulumi.InvokeOption) (*LookupRoutingRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRoutingRuleResult
	err := ctx.Invoke("azure-native:network:getRoutingRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoutingRuleArgs struct {
	// The name of the network manager Routing Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network manager routing Configuration rule collection.
	RuleCollectionName string `pulumi:"ruleCollectionName"`
	// The name of the rule.
	RuleName string `pulumi:"ruleName"`
}

// Network routing rule.
type LookupRoutingRuleResult struct {
	// A description for this rule.
	Description *string `pulumi:"description"`
	// Indicates the destination for this particular rule.
	Destination RoutingRuleRouteDestinationResponse `pulumi:"destination"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Indicates the next hop for this particular rule.
	NextHop RoutingRuleNextHopResponse `pulumi:"nextHop"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Unique identifier for this resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupRoutingRuleOutput(ctx *pulumi.Context, args LookupRoutingRuleOutputArgs, opts ...pulumi.InvokeOption) LookupRoutingRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoutingRuleResultOutput, error) {
			args := v.(LookupRoutingRuleArgs)
			opts = utilities.PkgInvokeDefaultOpts(opts)
			var rv LookupRoutingRuleResult
			secret, err := ctx.InvokePackageRaw("azure-native:network:getRoutingRule", args, &rv, "", opts...)
			if err != nil {
				return LookupRoutingRuleResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupRoutingRuleResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupRoutingRuleResultOutput), nil
			}
			return output, nil
		}).(LookupRoutingRuleResultOutput)
}

type LookupRoutingRuleOutputArgs struct {
	// The name of the network manager Routing Configuration.
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the network manager routing Configuration rule collection.
	RuleCollectionName pulumi.StringInput `pulumi:"ruleCollectionName"`
	// The name of the rule.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupRoutingRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutingRuleArgs)(nil)).Elem()
}

// Network routing rule.
type LookupRoutingRuleResultOutput struct{ *pulumi.OutputState }

func (LookupRoutingRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutingRuleResult)(nil)).Elem()
}

func (o LookupRoutingRuleResultOutput) ToLookupRoutingRuleResultOutput() LookupRoutingRuleResultOutput {
	return o
}

func (o LookupRoutingRuleResultOutput) ToLookupRoutingRuleResultOutputWithContext(ctx context.Context) LookupRoutingRuleResultOutput {
	return o
}

// A description for this rule.
func (o LookupRoutingRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoutingRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates the destination for this particular rule.
func (o LookupRoutingRuleResultOutput) Destination() RoutingRuleRouteDestinationResponseOutput {
	return o.ApplyT(func(v LookupRoutingRuleResult) RoutingRuleRouteDestinationResponse { return v.Destination }).(RoutingRuleRouteDestinationResponseOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupRoutingRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutingRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupRoutingRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutingRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupRoutingRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutingRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates the next hop for this particular rule.
func (o LookupRoutingRuleResultOutput) NextHop() RoutingRuleNextHopResponseOutput {
	return o.ApplyT(func(v LookupRoutingRuleResult) RoutingRuleNextHopResponse { return v.NextHop }).(RoutingRuleNextHopResponseOutput)
}

// The provisioning state of the resource.
func (o LookupRoutingRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutingRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Unique identifier for this resource.
func (o LookupRoutingRuleResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutingRuleResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o LookupRoutingRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRoutingRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupRoutingRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutingRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoutingRuleResultOutput{})
}