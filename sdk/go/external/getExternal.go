// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package external

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-external/sdk/go/external/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetExternal(ctx *pulumi.Context, args *GetExternalArgs, opts ...pulumi.InvokeOption) (*GetExternalResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetExternalResult
	err := ctx.Invoke("external:index/getExternal:getExternal", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getExternal.
type GetExternalArgs struct {
	Programs []string `pulumi:"programs"`
	// A map of string values to pass to the external program as the query arguments. If not supplied, the program will receive an empty object as its input.
	Query map[string]string `pulumi:"query"`
	// Working directory of the program. If not supplied, the program will run in the current directory.
	WorkingDir *string `pulumi:"workingDir"`
}

// A collection of values returned by getExternal.
type GetExternalResult struct {
	// The id of the data source. This will always be set to `-`
	Id       string   `pulumi:"id"`
	Programs []string `pulumi:"programs"`
	// A map of string values to pass to the external program as the query arguments. If not supplied, the program will receive an empty object as its input.
	Query map[string]string `pulumi:"query"`
	// A map of string values returned from the external program.
	Result map[string]string `pulumi:"result"`
	// Working directory of the program. If not supplied, the program will run in the current directory.
	WorkingDir *string `pulumi:"workingDir"`
}

func GetExternalOutput(ctx *pulumi.Context, args GetExternalOutputArgs, opts ...pulumi.InvokeOption) GetExternalResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetExternalResultOutput, error) {
			args := v.(GetExternalArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("external:index/getExternal:getExternal", args, GetExternalResultOutput{}, options).(GetExternalResultOutput), nil
		}).(GetExternalResultOutput)
}

// A collection of arguments for invoking getExternal.
type GetExternalOutputArgs struct {
	Programs pulumi.StringArrayInput `pulumi:"programs"`
	// A map of string values to pass to the external program as the query arguments. If not supplied, the program will receive an empty object as its input.
	Query pulumi.StringMapInput `pulumi:"query"`
	// Working directory of the program. If not supplied, the program will run in the current directory.
	WorkingDir pulumi.StringPtrInput `pulumi:"workingDir"`
}

func (GetExternalOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExternalArgs)(nil)).Elem()
}

// A collection of values returned by getExternal.
type GetExternalResultOutput struct{ *pulumi.OutputState }

func (GetExternalResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExternalResult)(nil)).Elem()
}

func (o GetExternalResultOutput) ToGetExternalResultOutput() GetExternalResultOutput {
	return o
}

func (o GetExternalResultOutput) ToGetExternalResultOutputWithContext(ctx context.Context) GetExternalResultOutput {
	return o
}

// The id of the data source. This will always be set to `-`
func (o GetExternalResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetExternalResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetExternalResultOutput) Programs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetExternalResult) []string { return v.Programs }).(pulumi.StringArrayOutput)
}

// A map of string values to pass to the external program as the query arguments. If not supplied, the program will receive an empty object as its input.
func (o GetExternalResultOutput) Query() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetExternalResult) map[string]string { return v.Query }).(pulumi.StringMapOutput)
}

// A map of string values returned from the external program.
func (o GetExternalResultOutput) Result() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetExternalResult) map[string]string { return v.Result }).(pulumi.StringMapOutput)
}

// Working directory of the program. If not supplied, the program will run in the current directory.
func (o GetExternalResultOutput) WorkingDir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetExternalResult) *string { return v.WorkingDir }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetExternalResultOutput{})
}
