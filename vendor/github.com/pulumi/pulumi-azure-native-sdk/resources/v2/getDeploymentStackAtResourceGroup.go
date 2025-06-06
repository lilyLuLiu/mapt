// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Deployment Stack with a given name.
//
// Uses Azure REST API version 2022-08-01-preview.
//
// Other available API versions: 2024-03-01.
func LookupDeploymentStackAtResourceGroup(ctx *pulumi.Context, args *LookupDeploymentStackAtResourceGroupArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentStackAtResourceGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDeploymentStackAtResourceGroupResult
	err := ctx.Invoke("azure-native:resources:getDeploymentStackAtResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentStackAtResourceGroupArgs struct {
	// Name of the deployment stack.
	DeploymentStackName string `pulumi:"deploymentStackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Deployment stack object.
type LookupDeploymentStackAtResourceGroupResult struct {
	// Defines the behavior of resources that are not managed immediately after the stack is updated.
	ActionOnUnmanage DeploymentStackPropertiesResponseActionOnUnmanage `pulumi:"actionOnUnmanage"`
	// The debug setting of the deployment.
	DebugSetting *DeploymentStacksDebugSettingResponse `pulumi:"debugSetting"`
	// An array of resources that were deleted during the most recent update.
	DeletedResources []ResourceReferenceResponse `pulumi:"deletedResources"`
	// Defines how resources deployed by the stack are locked.
	DenySettings DenySettingsResponse `pulumi:"denySettings"`
	// The resourceId of the deployment resource created by the deployment stack.
	DeploymentId string `pulumi:"deploymentId"`
	// The scope at which the initial deployment should be created. If a scope is not specified, it will default to the scope of the deployment stack. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroupId}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}').
	DeploymentScope *string `pulumi:"deploymentScope"`
	// Deployment stack description.
	Description *string `pulumi:"description"`
	// An array of resources that were detached during the most recent update.
	DetachedResources []ResourceReferenceResponse `pulumi:"detachedResources"`
	// The duration of the deployment stack update.
	Duration string `pulumi:"duration"`
	// Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows the OData error response format.).
	Error *ErrorResponseResponse `pulumi:"error"`
	// An array of resources that failed to reach goal state during the most recent update.
	FailedResources []ResourceReferenceExtendedResponse `pulumi:"failedResources"`
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// The location of the deployment stack. It cannot be changed after creation. It must be one of the supported Azure locations.
	Location *string `pulumi:"location"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// The outputs of the underlying deployment.
	Outputs interface{} `pulumi:"outputs"`
	// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
	Parameters interface{} `pulumi:"parameters"`
	// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
	ParametersLink *DeploymentStacksParametersLinkResponse `pulumi:"parametersLink"`
	// State of the deployment stack.
	ProvisioningState string `pulumi:"provisioningState"`
	// An array of resources currently managed by the deployment stack.
	Resources []ManagedResourceReferenceResponse `pulumi:"resources"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Deployment stack resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Type of this resource.
	Type string `pulumi:"type"`
}

func LookupDeploymentStackAtResourceGroupOutput(ctx *pulumi.Context, args LookupDeploymentStackAtResourceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentStackAtResourceGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDeploymentStackAtResourceGroupResultOutput, error) {
			args := v.(LookupDeploymentStackAtResourceGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:resources:getDeploymentStackAtResourceGroup", args, LookupDeploymentStackAtResourceGroupResultOutput{}, options).(LookupDeploymentStackAtResourceGroupResultOutput), nil
		}).(LookupDeploymentStackAtResourceGroupResultOutput)
}

type LookupDeploymentStackAtResourceGroupOutputArgs struct {
	// Name of the deployment stack.
	DeploymentStackName pulumi.StringInput `pulumi:"deploymentStackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDeploymentStackAtResourceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentStackAtResourceGroupArgs)(nil)).Elem()
}

// Deployment stack object.
type LookupDeploymentStackAtResourceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentStackAtResourceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentStackAtResourceGroupResult)(nil)).Elem()
}

func (o LookupDeploymentStackAtResourceGroupResultOutput) ToLookupDeploymentStackAtResourceGroupResultOutput() LookupDeploymentStackAtResourceGroupResultOutput {
	return o
}

func (o LookupDeploymentStackAtResourceGroupResultOutput) ToLookupDeploymentStackAtResourceGroupResultOutputWithContext(ctx context.Context) LookupDeploymentStackAtResourceGroupResultOutput {
	return o
}

// Defines the behavior of resources that are not managed immediately after the stack is updated.
func (o LookupDeploymentStackAtResourceGroupResultOutput) ActionOnUnmanage() DeploymentStackPropertiesResponseActionOnUnmanageOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) DeploymentStackPropertiesResponseActionOnUnmanage {
		return v.ActionOnUnmanage
	}).(DeploymentStackPropertiesResponseActionOnUnmanageOutput)
}

// The debug setting of the deployment.
func (o LookupDeploymentStackAtResourceGroupResultOutput) DebugSetting() DeploymentStacksDebugSettingResponsePtrOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) *DeploymentStacksDebugSettingResponse {
		return v.DebugSetting
	}).(DeploymentStacksDebugSettingResponsePtrOutput)
}

// An array of resources that were deleted during the most recent update.
func (o LookupDeploymentStackAtResourceGroupResultOutput) DeletedResources() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) []ResourceReferenceResponse {
		return v.DeletedResources
	}).(ResourceReferenceResponseArrayOutput)
}

// Defines how resources deployed by the stack are locked.
func (o LookupDeploymentStackAtResourceGroupResultOutput) DenySettings() DenySettingsResponseOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) DenySettingsResponse { return v.DenySettings }).(DenySettingsResponseOutput)
}

// The resourceId of the deployment resource created by the deployment stack.
func (o LookupDeploymentStackAtResourceGroupResultOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) string { return v.DeploymentId }).(pulumi.StringOutput)
}

// The scope at which the initial deployment should be created. If a scope is not specified, it will default to the scope of the deployment stack. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroupId}'), subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}').
func (o LookupDeploymentStackAtResourceGroupResultOutput) DeploymentScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) *string { return v.DeploymentScope }).(pulumi.StringPtrOutput)
}

// Deployment stack description.
func (o LookupDeploymentStackAtResourceGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// An array of resources that were detached during the most recent update.
func (o LookupDeploymentStackAtResourceGroupResultOutput) DetachedResources() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) []ResourceReferenceResponse {
		return v.DetachedResources
	}).(ResourceReferenceResponseArrayOutput)
}

// The duration of the deployment stack update.
func (o LookupDeploymentStackAtResourceGroupResultOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) string { return v.Duration }).(pulumi.StringOutput)
}

// Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows the OData error response format.).
func (o LookupDeploymentStackAtResourceGroupResultOutput) Error() ErrorResponseResponsePtrOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) *ErrorResponseResponse { return v.Error }).(ErrorResponseResponsePtrOutput)
}

// An array of resources that failed to reach goal state during the most recent update.
func (o LookupDeploymentStackAtResourceGroupResultOutput) FailedResources() ResourceReferenceExtendedResponseArrayOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) []ResourceReferenceExtendedResponse {
		return v.FailedResources
	}).(ResourceReferenceExtendedResponseArrayOutput)
}

// String Id used to locate any resource on Azure.
func (o LookupDeploymentStackAtResourceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the deployment stack. It cannot be changed after creation. It must be one of the supported Azure locations.
func (o LookupDeploymentStackAtResourceGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of this resource.
func (o LookupDeploymentStackAtResourceGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The outputs of the underlying deployment.
func (o LookupDeploymentStackAtResourceGroupResultOutput) Outputs() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) interface{} { return v.Outputs }).(pulumi.AnyOutput)
}

// Name and value pairs that define the deployment parameters for the template. Use this element when providing the parameter values directly in the request, rather than linking to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
func (o LookupDeploymentStackAtResourceGroupResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

// The URI of parameters file. Use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
func (o LookupDeploymentStackAtResourceGroupResultOutput) ParametersLink() DeploymentStacksParametersLinkResponsePtrOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) *DeploymentStacksParametersLinkResponse {
		return v.ParametersLink
	}).(DeploymentStacksParametersLinkResponsePtrOutput)
}

// State of the deployment stack.
func (o LookupDeploymentStackAtResourceGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// An array of resources currently managed by the deployment stack.
func (o LookupDeploymentStackAtResourceGroupResultOutput) Resources() ManagedResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) []ManagedResourceReferenceResponse {
		return v.Resources
	}).(ManagedResourceReferenceResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDeploymentStackAtResourceGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Deployment stack resource tags.
func (o LookupDeploymentStackAtResourceGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of this resource.
func (o LookupDeploymentStackAtResourceGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentStackAtResourceGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentStackAtResourceGroupResultOutput{})
}
