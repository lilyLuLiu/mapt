// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Access Review History Definition.
// Azure REST API version: 2021-12-01-preview. Prior API version in Azure Native 1.x: 2021-11-16-preview.
type AccessReviewHistoryDefinitionById struct {
	pulumi.CustomResourceState

	// Date time when history definition was created
	CreatedDateTime pulumi.StringOutput `pulumi:"createdDateTime"`
	// Collection of review decisions which the history data should be filtered on. For example if Approve and Deny are supplied the data will only contain review results in which the decision maker approved or denied a review request.
	Decisions pulumi.StringArrayOutput `pulumi:"decisions"`
	// The display name for the history definition.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The DateTime when the review is scheduled to end. Required if type is endDate
	EndDate pulumi.StringPtrOutput `pulumi:"endDate"`
	// Set of access review history instances for this history definition.
	Instances AccessReviewHistoryInstanceResponseArrayOutput `pulumi:"instances"`
	// The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// The access review history definition unique id.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of times to repeat the access review. Required and must be positive if type is numbered.
	NumberOfOccurrences pulumi.IntPtrOutput `pulumi:"numberOfOccurrences"`
	// The identity id
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The identity display name
	PrincipalName pulumi.StringOutput `pulumi:"principalName"`
	// The identity type : user/servicePrincipal
	PrincipalType pulumi.StringOutput `pulumi:"principalType"`
	// Date time used when selecting review data, all reviews included in data end on or before this date. For use only with one-time/non-recurring reports.
	ReviewHistoryPeriodEndDateTime pulumi.StringOutput `pulumi:"reviewHistoryPeriodEndDateTime"`
	// Date time used when selecting review data, all reviews included in data start on or after this date. For use only with one-time/non-recurring reports.
	ReviewHistoryPeriodStartDateTime pulumi.StringOutput `pulumi:"reviewHistoryPeriodStartDateTime"`
	// A collection of scopes used when selecting review history data
	Scopes AccessReviewScopeResponseArrayOutput `pulumi:"scopes"`
	// The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
	StartDate pulumi.StringPtrOutput `pulumi:"startDate"`
	// This read-only field specifies the of the requested review history data. This is either requested, in-progress, done or error.
	Status pulumi.StringOutput `pulumi:"status"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The user principal name(if valid)
	UserPrincipalName pulumi.StringOutput `pulumi:"userPrincipalName"`
}

// NewAccessReviewHistoryDefinitionById registers a new resource with the given unique name, arguments, and options.
func NewAccessReviewHistoryDefinitionById(ctx *pulumi.Context,
	name string, args *AccessReviewHistoryDefinitionByIdArgs, opts ...pulumi.ResourceOption) (*AccessReviewHistoryDefinitionById, error) {
	if args == nil {
		args = &AccessReviewHistoryDefinitionByIdArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20211116preview:AccessReviewHistoryDefinitionById"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20211201preview:AccessReviewHistoryDefinitionById"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AccessReviewHistoryDefinitionById
	err := ctx.RegisterResource("azure-native:authorization:AccessReviewHistoryDefinitionById", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessReviewHistoryDefinitionById gets an existing AccessReviewHistoryDefinitionById resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessReviewHistoryDefinitionById(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessReviewHistoryDefinitionByIdState, opts ...pulumi.ResourceOption) (*AccessReviewHistoryDefinitionById, error) {
	var resource AccessReviewHistoryDefinitionById
	err := ctx.ReadResource("azure-native:authorization:AccessReviewHistoryDefinitionById", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessReviewHistoryDefinitionById resources.
type accessReviewHistoryDefinitionByIdState struct {
}

type AccessReviewHistoryDefinitionByIdState struct {
}

func (AccessReviewHistoryDefinitionByIdState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessReviewHistoryDefinitionByIdState)(nil)).Elem()
}

type accessReviewHistoryDefinitionByIdArgs struct {
	// Collection of review decisions which the history data should be filtered on. For example if Approve and Deny are supplied the data will only contain review results in which the decision maker approved or denied a review request.
	Decisions []string `pulumi:"decisions"`
	// The display name for the history definition.
	DisplayName *string `pulumi:"displayName"`
	// The DateTime when the review is scheduled to end. Required if type is endDate
	EndDate *string `pulumi:"endDate"`
	// The id of the access review history definition.
	HistoryDefinitionId *string `pulumi:"historyDefinitionId"`
	// Set of access review history instances for this history definition.
	Instances []AccessReviewHistoryInstance `pulumi:"instances"`
	// The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
	Interval *int `pulumi:"interval"`
	// The number of times to repeat the access review. Required and must be positive if type is numbered.
	NumberOfOccurrences *int `pulumi:"numberOfOccurrences"`
	// A collection of scopes used when selecting review history data
	Scopes []AccessReviewScope `pulumi:"scopes"`
	// The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
	StartDate *string `pulumi:"startDate"`
	// The recurrence range type. The possible values are: endDate, noEnd, numbered.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a AccessReviewHistoryDefinitionById resource.
type AccessReviewHistoryDefinitionByIdArgs struct {
	// Collection of review decisions which the history data should be filtered on. For example if Approve and Deny are supplied the data will only contain review results in which the decision maker approved or denied a review request.
	Decisions pulumi.StringArrayInput
	// The display name for the history definition.
	DisplayName pulumi.StringPtrInput
	// The DateTime when the review is scheduled to end. Required if type is endDate
	EndDate pulumi.StringPtrInput
	// The id of the access review history definition.
	HistoryDefinitionId pulumi.StringPtrInput
	// Set of access review history instances for this history definition.
	Instances AccessReviewHistoryInstanceArrayInput
	// The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
	Interval pulumi.IntPtrInput
	// The number of times to repeat the access review. Required and must be positive if type is numbered.
	NumberOfOccurrences pulumi.IntPtrInput
	// A collection of scopes used when selecting review history data
	Scopes AccessReviewScopeArrayInput
	// The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
	StartDate pulumi.StringPtrInput
	// The recurrence range type. The possible values are: endDate, noEnd, numbered.
	Type pulumi.StringPtrInput
}

func (AccessReviewHistoryDefinitionByIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessReviewHistoryDefinitionByIdArgs)(nil)).Elem()
}

type AccessReviewHistoryDefinitionByIdInput interface {
	pulumi.Input

	ToAccessReviewHistoryDefinitionByIdOutput() AccessReviewHistoryDefinitionByIdOutput
	ToAccessReviewHistoryDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewHistoryDefinitionByIdOutput
}

func (*AccessReviewHistoryDefinitionById) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessReviewHistoryDefinitionById)(nil)).Elem()
}

func (i *AccessReviewHistoryDefinitionById) ToAccessReviewHistoryDefinitionByIdOutput() AccessReviewHistoryDefinitionByIdOutput {
	return i.ToAccessReviewHistoryDefinitionByIdOutputWithContext(context.Background())
}

func (i *AccessReviewHistoryDefinitionById) ToAccessReviewHistoryDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewHistoryDefinitionByIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewHistoryDefinitionByIdOutput)
}

type AccessReviewHistoryDefinitionByIdOutput struct{ *pulumi.OutputState }

func (AccessReviewHistoryDefinitionByIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessReviewHistoryDefinitionById)(nil)).Elem()
}

func (o AccessReviewHistoryDefinitionByIdOutput) ToAccessReviewHistoryDefinitionByIdOutput() AccessReviewHistoryDefinitionByIdOutput {
	return o
}

func (o AccessReviewHistoryDefinitionByIdOutput) ToAccessReviewHistoryDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewHistoryDefinitionByIdOutput {
	return o
}

// Date time when history definition was created
func (o AccessReviewHistoryDefinitionByIdOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.CreatedDateTime }).(pulumi.StringOutput)
}

// Collection of review decisions which the history data should be filtered on. For example if Approve and Deny are supplied the data will only contain review results in which the decision maker approved or denied a review request.
func (o AccessReviewHistoryDefinitionByIdOutput) Decisions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringArrayOutput { return v.Decisions }).(pulumi.StringArrayOutput)
}

// The display name for the history definition.
func (o AccessReviewHistoryDefinitionByIdOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The DateTime when the review is scheduled to end. Required if type is endDate
func (o AccessReviewHistoryDefinitionByIdOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringPtrOutput { return v.EndDate }).(pulumi.StringPtrOutput)
}

// Set of access review history instances for this history definition.
func (o AccessReviewHistoryDefinitionByIdOutput) Instances() AccessReviewHistoryInstanceResponseArrayOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) AccessReviewHistoryInstanceResponseArrayOutput {
		return v.Instances
	}).(AccessReviewHistoryInstanceResponseArrayOutput)
}

// The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
func (o AccessReviewHistoryDefinitionByIdOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

// The access review history definition unique id.
func (o AccessReviewHistoryDefinitionByIdOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of times to repeat the access review. Required and must be positive if type is numbered.
func (o AccessReviewHistoryDefinitionByIdOutput) NumberOfOccurrences() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.IntPtrOutput { return v.NumberOfOccurrences }).(pulumi.IntPtrOutput)
}

// The identity id
func (o AccessReviewHistoryDefinitionByIdOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

// The identity display name
func (o AccessReviewHistoryDefinitionByIdOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

// The identity type : user/servicePrincipal
func (o AccessReviewHistoryDefinitionByIdOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

// Date time used when selecting review data, all reviews included in data end on or before this date. For use only with one-time/non-recurring reports.
func (o AccessReviewHistoryDefinitionByIdOutput) ReviewHistoryPeriodEndDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput {
		return v.ReviewHistoryPeriodEndDateTime
	}).(pulumi.StringOutput)
}

// Date time used when selecting review data, all reviews included in data start on or after this date. For use only with one-time/non-recurring reports.
func (o AccessReviewHistoryDefinitionByIdOutput) ReviewHistoryPeriodStartDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput {
		return v.ReviewHistoryPeriodStartDateTime
	}).(pulumi.StringOutput)
}

// A collection of scopes used when selecting review history data
func (o AccessReviewHistoryDefinitionByIdOutput) Scopes() AccessReviewScopeResponseArrayOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) AccessReviewScopeResponseArrayOutput { return v.Scopes }).(AccessReviewScopeResponseArrayOutput)
}

// The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
func (o AccessReviewHistoryDefinitionByIdOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringPtrOutput { return v.StartDate }).(pulumi.StringPtrOutput)
}

// This read-only field specifies the of the requested review history data. This is either requested, in-progress, done or error.
func (o AccessReviewHistoryDefinitionByIdOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The resource type.
func (o AccessReviewHistoryDefinitionByIdOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The user principal name(if valid)
func (o AccessReviewHistoryDefinitionByIdOutput) UserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewHistoryDefinitionById) pulumi.StringOutput { return v.UserPrincipalName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessReviewHistoryDefinitionByIdOutput{})
}