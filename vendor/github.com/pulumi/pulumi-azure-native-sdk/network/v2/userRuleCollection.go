// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines the user rule collection.
// Azure REST API version: 2022-04-01-preview. Prior API version in Azure Native 1.x: 2021-02-01-preview
type UserRuleCollection struct {
	pulumi.CustomResourceState

	// Groups for configuration
	AppliesToGroups NetworkManagerSecurityGroupItemResponseArrayOutput `pulumi:"appliesToGroups"`
	// A description of the user rule collection.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewUserRuleCollection registers a new resource with the given unique name, arguments, and options.
func NewUserRuleCollection(ctx *pulumi.Context,
	name string, args *UserRuleCollectionArgs, opts ...pulumi.ResourceOption) (*UserRuleCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppliesToGroups == nil {
		return nil, errors.New("invalid value for required argument 'AppliesToGroups'")
	}
	if args.ConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationName'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210201preview:UserRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:UserRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:UserRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:UserRuleCollection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource UserRuleCollection
	err := ctx.RegisterResource("azure-native:network:UserRuleCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserRuleCollection gets an existing UserRuleCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserRuleCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserRuleCollectionState, opts ...pulumi.ResourceOption) (*UserRuleCollection, error) {
	var resource UserRuleCollection
	err := ctx.ReadResource("azure-native:network:UserRuleCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserRuleCollection resources.
type userRuleCollectionState struct {
}

type UserRuleCollectionState struct {
}

func (UserRuleCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*userRuleCollectionState)(nil)).Elem()
}

type userRuleCollectionArgs struct {
	// Groups for configuration
	AppliesToGroups []NetworkManagerSecurityGroupItem `pulumi:"appliesToGroups"`
	// The name of the network manager Security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// A description of the user rule collection.
	Description *string `pulumi:"description"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName *string `pulumi:"ruleCollectionName"`
}

// The set of arguments for constructing a UserRuleCollection resource.
type UserRuleCollectionArgs struct {
	// Groups for configuration
	AppliesToGroups NetworkManagerSecurityGroupItemArrayInput
	// The name of the network manager Security Configuration.
	ConfigurationName pulumi.StringInput
	// A description of the user rule collection.
	Description pulumi.StringPtrInput
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName pulumi.StringPtrInput
}

func (UserRuleCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userRuleCollectionArgs)(nil)).Elem()
}

type UserRuleCollectionInput interface {
	pulumi.Input

	ToUserRuleCollectionOutput() UserRuleCollectionOutput
	ToUserRuleCollectionOutputWithContext(ctx context.Context) UserRuleCollectionOutput
}

func (*UserRuleCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRuleCollection)(nil)).Elem()
}

func (i *UserRuleCollection) ToUserRuleCollectionOutput() UserRuleCollectionOutput {
	return i.ToUserRuleCollectionOutputWithContext(context.Background())
}

func (i *UserRuleCollection) ToUserRuleCollectionOutputWithContext(ctx context.Context) UserRuleCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRuleCollectionOutput)
}

type UserRuleCollectionOutput struct{ *pulumi.OutputState }

func (UserRuleCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRuleCollection)(nil)).Elem()
}

func (o UserRuleCollectionOutput) ToUserRuleCollectionOutput() UserRuleCollectionOutput {
	return o
}

func (o UserRuleCollectionOutput) ToUserRuleCollectionOutputWithContext(ctx context.Context) UserRuleCollectionOutput {
	return o
}

// Groups for configuration
func (o UserRuleCollectionOutput) AppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v *UserRuleCollection) NetworkManagerSecurityGroupItemResponseArrayOutput {
		return v.AppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

// A description of the user rule collection.
func (o UserRuleCollectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserRuleCollection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o UserRuleCollectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRuleCollection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource name.
func (o UserRuleCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRuleCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o UserRuleCollectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRuleCollection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o UserRuleCollectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *UserRuleCollection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o UserRuleCollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRuleCollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UserRuleCollectionOutput{})
}