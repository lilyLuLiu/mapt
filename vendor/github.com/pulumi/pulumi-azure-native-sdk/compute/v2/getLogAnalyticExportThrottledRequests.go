// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Export logs that show total throttled Api requests for this subscription in the given time window.
//
// Uses Azure REST API version 2023-03-01.
//
// Other available API versions: 2020-12-01, 2021-03-01, 2021-04-01, 2021-07-01, 2021-11-01, 2022-03-01, 2022-08-01, 2022-11-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01, 2024-11-01.
func GetLogAnalyticExportThrottledRequests(ctx *pulumi.Context, args *GetLogAnalyticExportThrottledRequestsArgs, opts ...pulumi.InvokeOption) (*GetLogAnalyticExportThrottledRequestsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetLogAnalyticExportThrottledRequestsResult
	err := ctx.Invoke("azure-native:compute:getLogAnalyticExportThrottledRequests", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLogAnalyticExportThrottledRequestsArgs struct {
	// SAS Uri of the logging blob container to which LogAnalytics Api writes output logs to.
	BlobContainerSasUri string `pulumi:"blobContainerSasUri"`
	// From time of the query
	FromTime string `pulumi:"fromTime"`
	// Group query result by Client Application ID.
	GroupByClientApplicationId *bool `pulumi:"groupByClientApplicationId"`
	// Group query result by Operation Name.
	GroupByOperationName *bool `pulumi:"groupByOperationName"`
	// Group query result by Resource Name.
	GroupByResourceName *bool `pulumi:"groupByResourceName"`
	// Group query result by Throttle Policy applied.
	GroupByThrottlePolicy *bool `pulumi:"groupByThrottlePolicy"`
	// Group query result by User Agent.
	GroupByUserAgent *bool `pulumi:"groupByUserAgent"`
	// The location upon which virtual-machine-sizes is queried.
	Location string `pulumi:"location"`
	// To time of the query
	ToTime string `pulumi:"toTime"`
}

// LogAnalytics operation status response
type GetLogAnalyticExportThrottledRequestsResult struct {
	// LogAnalyticsOutput
	Properties LogAnalyticsOutputResponse `pulumi:"properties"`
}

func GetLogAnalyticExportThrottledRequestsOutput(ctx *pulumi.Context, args GetLogAnalyticExportThrottledRequestsOutputArgs, opts ...pulumi.InvokeOption) GetLogAnalyticExportThrottledRequestsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetLogAnalyticExportThrottledRequestsResultOutput, error) {
			args := v.(GetLogAnalyticExportThrottledRequestsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:compute:getLogAnalyticExportThrottledRequests", args, GetLogAnalyticExportThrottledRequestsResultOutput{}, options).(GetLogAnalyticExportThrottledRequestsResultOutput), nil
		}).(GetLogAnalyticExportThrottledRequestsResultOutput)
}

type GetLogAnalyticExportThrottledRequestsOutputArgs struct {
	// SAS Uri of the logging blob container to which LogAnalytics Api writes output logs to.
	BlobContainerSasUri pulumi.StringInput `pulumi:"blobContainerSasUri"`
	// From time of the query
	FromTime pulumi.StringInput `pulumi:"fromTime"`
	// Group query result by Client Application ID.
	GroupByClientApplicationId pulumi.BoolPtrInput `pulumi:"groupByClientApplicationId"`
	// Group query result by Operation Name.
	GroupByOperationName pulumi.BoolPtrInput `pulumi:"groupByOperationName"`
	// Group query result by Resource Name.
	GroupByResourceName pulumi.BoolPtrInput `pulumi:"groupByResourceName"`
	// Group query result by Throttle Policy applied.
	GroupByThrottlePolicy pulumi.BoolPtrInput `pulumi:"groupByThrottlePolicy"`
	// Group query result by User Agent.
	GroupByUserAgent pulumi.BoolPtrInput `pulumi:"groupByUserAgent"`
	// The location upon which virtual-machine-sizes is queried.
	Location pulumi.StringInput `pulumi:"location"`
	// To time of the query
	ToTime pulumi.StringInput `pulumi:"toTime"`
}

func (GetLogAnalyticExportThrottledRequestsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogAnalyticExportThrottledRequestsArgs)(nil)).Elem()
}

// LogAnalytics operation status response
type GetLogAnalyticExportThrottledRequestsResultOutput struct{ *pulumi.OutputState }

func (GetLogAnalyticExportThrottledRequestsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogAnalyticExportThrottledRequestsResult)(nil)).Elem()
}

func (o GetLogAnalyticExportThrottledRequestsResultOutput) ToGetLogAnalyticExportThrottledRequestsResultOutput() GetLogAnalyticExportThrottledRequestsResultOutput {
	return o
}

func (o GetLogAnalyticExportThrottledRequestsResultOutput) ToGetLogAnalyticExportThrottledRequestsResultOutputWithContext(ctx context.Context) GetLogAnalyticExportThrottledRequestsResultOutput {
	return o
}

// LogAnalyticsOutput
func (o GetLogAnalyticExportThrottledRequestsResultOutput) Properties() LogAnalyticsOutputResponseOutput {
	return o.ApplyT(func(v GetLogAnalyticExportThrottledRequestsResult) LogAnalyticsOutputResponse { return v.Properties }).(LogAnalyticsOutputResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogAnalyticExportThrottledRequestsResultOutput{})
}
