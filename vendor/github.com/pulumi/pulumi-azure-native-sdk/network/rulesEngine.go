// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A rules engine configuration containing a list of rules that will run to modify the runtime behavior of the request and response.
// API Version: 2020-05-01.
type RulesEngine struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource status.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// A list of rules that define a particular Rules Engine Configuration.
	Rules RulesEngineRuleResponseArrayOutput `pulumi:"rules"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRulesEngine registers a new resource with the given unique name, arguments, and options.
func NewRulesEngine(ctx *pulumi.Context,
	name string, args *RulesEngineArgs, opts ...pulumi.ResourceOption) (*RulesEngine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FrontDoorName == nil {
		return nil, errors.New("invalid value for required argument 'FrontDoorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20200101:RulesEngine"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:RulesEngine"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:RulesEngine"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210601:RulesEngine"),
		},
	})
	opts = append(opts, aliases)
	var resource RulesEngine
	err := ctx.RegisterResource("azure-native:network:RulesEngine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRulesEngine gets an existing RulesEngine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRulesEngine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RulesEngineState, opts ...pulumi.ResourceOption) (*RulesEngine, error) {
	var resource RulesEngine
	err := ctx.ReadResource("azure-native:network:RulesEngine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RulesEngine resources.
type rulesEngineState struct {
}

type RulesEngineState struct {
}

func (RulesEngineState) ElementType() reflect.Type {
	return reflect.TypeOf((*rulesEngineState)(nil)).Elem()
}

type rulesEngineArgs struct {
	// Name of the Front Door which is globally unique.
	FrontDoorName string `pulumi:"frontDoorName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A list of rules that define a particular Rules Engine Configuration.
	Rules []RulesEngineRule `pulumi:"rules"`
	// Name of the Rules Engine which is unique within the Front Door.
	RulesEngineName *string `pulumi:"rulesEngineName"`
}

// The set of arguments for constructing a RulesEngine resource.
type RulesEngineArgs struct {
	// Name of the Front Door which is globally unique.
	FrontDoorName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// A list of rules that define a particular Rules Engine Configuration.
	Rules RulesEngineRuleArrayInput
	// Name of the Rules Engine which is unique within the Front Door.
	RulesEngineName pulumi.StringPtrInput
}

func (RulesEngineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rulesEngineArgs)(nil)).Elem()
}

type RulesEngineInput interface {
	pulumi.Input

	ToRulesEngineOutput() RulesEngineOutput
	ToRulesEngineOutputWithContext(ctx context.Context) RulesEngineOutput
}

func (*RulesEngine) ElementType() reflect.Type {
	return reflect.TypeOf((**RulesEngine)(nil)).Elem()
}

func (i *RulesEngine) ToRulesEngineOutput() RulesEngineOutput {
	return i.ToRulesEngineOutputWithContext(context.Background())
}

func (i *RulesEngine) ToRulesEngineOutputWithContext(ctx context.Context) RulesEngineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineOutput)
}

type RulesEngineOutput struct{ *pulumi.OutputState }

func (RulesEngineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RulesEngine)(nil)).Elem()
}

func (o RulesEngineOutput) ToRulesEngineOutput() RulesEngineOutput {
	return o
}

func (o RulesEngineOutput) ToRulesEngineOutputWithContext(ctx context.Context) RulesEngineOutput {
	return o
}

// Resource name.
func (o RulesEngineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RulesEngine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource status.
func (o RulesEngineOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *RulesEngine) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// A list of rules that define a particular Rules Engine Configuration.
func (o RulesEngineOutput) Rules() RulesEngineRuleResponseArrayOutput {
	return o.ApplyT(func(v *RulesEngine) RulesEngineRuleResponseArrayOutput { return v.Rules }).(RulesEngineRuleResponseArrayOutput)
}

// Resource type.
func (o RulesEngineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RulesEngine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RulesEngineOutput{})
}