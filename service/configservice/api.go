// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package configservice provides a client for AWS Config.
package configservice

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
)

const opDeleteDeliveryChannel = "DeleteDeliveryChannel"

// DeleteDeliveryChannelRequest generates a request for the DeleteDeliveryChannel operation.
func (c *ConfigService) DeleteDeliveryChannelRequest(input *DeleteDeliveryChannelInput) (req *aws.Request, output *DeleteDeliveryChannelOutput) {
	op := &aws.Operation{
		Name:       opDeleteDeliveryChannel,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDeliveryChannelInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeleteDeliveryChannelOutput{}
	req.Data = output
	return
}

// Deletes the specified delivery channel.
//
// The delivery channel cannot be deleted if it is the only delivery channel
// and the configuration recorder is still running. To delete the delivery channel,
// stop the running configuration recorder using the StopConfigurationRecorder
// action.
func (c *ConfigService) DeleteDeliveryChannel(input *DeleteDeliveryChannelInput) (*DeleteDeliveryChannelOutput, error) {
	req, out := c.DeleteDeliveryChannelRequest(input)
	err := req.Send()
	return out, err
}

const opDeliverConfigSnapshot = "DeliverConfigSnapshot"

// DeliverConfigSnapshotRequest generates a request for the DeliverConfigSnapshot operation.
func (c *ConfigService) DeliverConfigSnapshotRequest(input *DeliverConfigSnapshotInput) (req *aws.Request, output *DeliverConfigSnapshotOutput) {
	op := &aws.Operation{
		Name:       opDeliverConfigSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeliverConfigSnapshotInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeliverConfigSnapshotOutput{}
	req.Data = output
	return
}

// Schedules delivery of a configuration snapshot to the Amazon S3 bucket in
// the specified delivery channel. After the delivery has started, AWS Config
// sends following notifications using an Amazon SNS topic that you have specified.
//
//  Notification of starting the delivery. Notification of delivery completed,
// if the delivery was successfully completed. Notification of delivery failure,
// if the delivery failed to complete.
func (c *ConfigService) DeliverConfigSnapshot(input *DeliverConfigSnapshotInput) (*DeliverConfigSnapshotOutput, error) {
	req, out := c.DeliverConfigSnapshotRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeConfigurationRecorderStatus = "DescribeConfigurationRecorderStatus"

// DescribeConfigurationRecorderStatusRequest generates a request for the DescribeConfigurationRecorderStatus operation.
func (c *ConfigService) DescribeConfigurationRecorderStatusRequest(input *DescribeConfigurationRecorderStatusInput) (req *aws.Request, output *DescribeConfigurationRecorderStatusOutput) {
	op := &aws.Operation{
		Name:       opDescribeConfigurationRecorderStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConfigurationRecorderStatusInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeConfigurationRecorderStatusOutput{}
	req.Data = output
	return
}

// Returns the current status of the specified configuration recorder. If a
// configuration recorder is not specified, this action returns the status of
// all configuration recorder associated with the account.
//
// Currently, you can specify only one configuration recorder per account.
func (c *ConfigService) DescribeConfigurationRecorderStatus(input *DescribeConfigurationRecorderStatusInput) (*DescribeConfigurationRecorderStatusOutput, error) {
	req, out := c.DescribeConfigurationRecorderStatusRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeConfigurationRecorders = "DescribeConfigurationRecorders"

// DescribeConfigurationRecordersRequest generates a request for the DescribeConfigurationRecorders operation.
func (c *ConfigService) DescribeConfigurationRecordersRequest(input *DescribeConfigurationRecordersInput) (req *aws.Request, output *DescribeConfigurationRecordersOutput) {
	op := &aws.Operation{
		Name:       opDescribeConfigurationRecorders,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConfigurationRecordersInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeConfigurationRecordersOutput{}
	req.Data = output
	return
}

// Returns the name of one or more specified configuration recorders. If the
// recorder name is not specified, this action returns the names of all the
// configuration recorders associated with the account.
//
//  Currently, you can specify only one configuration recorder per account.
func (c *ConfigService) DescribeConfigurationRecorders(input *DescribeConfigurationRecordersInput) (*DescribeConfigurationRecordersOutput, error) {
	req, out := c.DescribeConfigurationRecordersRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeDeliveryChannelStatus = "DescribeDeliveryChannelStatus"

// DescribeDeliveryChannelStatusRequest generates a request for the DescribeDeliveryChannelStatus operation.
func (c *ConfigService) DescribeDeliveryChannelStatusRequest(input *DescribeDeliveryChannelStatusInput) (req *aws.Request, output *DescribeDeliveryChannelStatusOutput) {
	op := &aws.Operation{
		Name:       opDescribeDeliveryChannelStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDeliveryChannelStatusInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeDeliveryChannelStatusOutput{}
	req.Data = output
	return
}

// Returns the current status of the specified delivery channel. If a delivery
// channel is not specified, this action returns the current status of all delivery
// channels associated with the account.
//
// Currently, you can specify only one delivery channel per account.
func (c *ConfigService) DescribeDeliveryChannelStatus(input *DescribeDeliveryChannelStatusInput) (*DescribeDeliveryChannelStatusOutput, error) {
	req, out := c.DescribeDeliveryChannelStatusRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeDeliveryChannels = "DescribeDeliveryChannels"

// DescribeDeliveryChannelsRequest generates a request for the DescribeDeliveryChannels operation.
func (c *ConfigService) DescribeDeliveryChannelsRequest(input *DescribeDeliveryChannelsInput) (req *aws.Request, output *DescribeDeliveryChannelsOutput) {
	op := &aws.Operation{
		Name:       opDescribeDeliveryChannels,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDeliveryChannelsInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeDeliveryChannelsOutput{}
	req.Data = output
	return
}

// Returns details about the specified delivery channel. If a delivery channel
// is not specified, this action returns the details of all delivery channels
// associated with the account.
//
//  Currently, you can specify only one delivery channel per account.
func (c *ConfigService) DescribeDeliveryChannels(input *DescribeDeliveryChannelsInput) (*DescribeDeliveryChannelsOutput, error) {
	req, out := c.DescribeDeliveryChannelsRequest(input)
	err := req.Send()
	return out, err
}

const opGetResourceConfigHistory = "GetResourceConfigHistory"

// GetResourceConfigHistoryRequest generates a request for the GetResourceConfigHistory operation.
func (c *ConfigService) GetResourceConfigHistoryRequest(input *GetResourceConfigHistoryInput) (req *aws.Request, output *GetResourceConfigHistoryOutput) {
	op := &aws.Operation{
		Name:       opGetResourceConfigHistory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetResourceConfigHistoryInput{}
	}

	req = c.newRequest(op, input, output)
	output = &GetResourceConfigHistoryOutput{}
	req.Data = output
	return
}

// Returns a list of configuration items for the specified resource. The list
// contains details about each state of the resource during the specified time
// interval. You can specify a limit on the number of results returned on the
// page. If a limit is specified, a nextToken is returned as part of the result
// that you can use to continue this request.
//
//  Each call to the API is limited to span a duration of seven days. It is
// likely that the number of records returned is smaller than the specified
// limit. In such cases, you can make another call, using the nextToken .
func (c *ConfigService) GetResourceConfigHistory(input *GetResourceConfigHistoryInput) (*GetResourceConfigHistoryOutput, error) {
	req, out := c.GetResourceConfigHistoryRequest(input)
	err := req.Send()
	return out, err
}

func (c *ConfigService) GetResourceConfigHistoryPages(input *GetResourceConfigHistoryInput, fn func(p *GetResourceConfigHistoryOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.GetResourceConfigHistoryRequest(input)
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*GetResourceConfigHistoryOutput), lastPage)
	})
}

const opPutConfigurationRecorder = "PutConfigurationRecorder"

// PutConfigurationRecorderRequest generates a request for the PutConfigurationRecorder operation.
func (c *ConfigService) PutConfigurationRecorderRequest(input *PutConfigurationRecorderInput) (req *aws.Request, output *PutConfigurationRecorderOutput) {
	op := &aws.Operation{
		Name:       opPutConfigurationRecorder,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutConfigurationRecorderInput{}
	}

	req = c.newRequest(op, input, output)
	output = &PutConfigurationRecorderOutput{}
	req.Data = output
	return
}

// Creates a new configuration recorder to record the selected resource configurations.
//
// You can use this action to change the role roleARN and/or the recordingGroup
// of an existing recorder. To change the role, call the action on the existing
// configuration recorder and specify a role.
//
//  Currently, you can specify only one configuration recorder per account.
//
// If ConfigurationRecorder does not have the recordingGroup parameter specified,
// the default is to record all supported resource types.
func (c *ConfigService) PutConfigurationRecorder(input *PutConfigurationRecorderInput) (*PutConfigurationRecorderOutput, error) {
	req, out := c.PutConfigurationRecorderRequest(input)
	err := req.Send()
	return out, err
}

const opPutDeliveryChannel = "PutDeliveryChannel"

// PutDeliveryChannelRequest generates a request for the PutDeliveryChannel operation.
func (c *ConfigService) PutDeliveryChannelRequest(input *PutDeliveryChannelInput) (req *aws.Request, output *PutDeliveryChannelOutput) {
	op := &aws.Operation{
		Name:       opPutDeliveryChannel,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutDeliveryChannelInput{}
	}

	req = c.newRequest(op, input, output)
	output = &PutDeliveryChannelOutput{}
	req.Data = output
	return
}

// Creates a new delivery channel object to deliver the configuration information
// to an Amazon S3 bucket, and to an Amazon SNS topic.
//
// You can use this action to change the Amazon S3 bucket or an Amazon SNS
// topic of the existing delivery channel. To change the Amazon S3 bucket or
// an Amazon SNS topic, call this action and specify the changed values for
// the S3 bucket and the SNS topic. If you specify a different value for either
// the S3 bucket or the SNS topic, this action will keep the existing value
// for the parameter that is not changed.
//
//  Currently, you can specify only one delivery channel per account.
func (c *ConfigService) PutDeliveryChannel(input *PutDeliveryChannelInput) (*PutDeliveryChannelOutput, error) {
	req, out := c.PutDeliveryChannelRequest(input)
	err := req.Send()
	return out, err
}

const opStartConfigurationRecorder = "StartConfigurationRecorder"

// StartConfigurationRecorderRequest generates a request for the StartConfigurationRecorder operation.
func (c *ConfigService) StartConfigurationRecorderRequest(input *StartConfigurationRecorderInput) (req *aws.Request, output *StartConfigurationRecorderOutput) {
	op := &aws.Operation{
		Name:       opStartConfigurationRecorder,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartConfigurationRecorderInput{}
	}

	req = c.newRequest(op, input, output)
	output = &StartConfigurationRecorderOutput{}
	req.Data = output
	return
}

// Starts recording configurations of the AWS resources you have selected to
// record in your AWS account.
//
// You must have created at least one delivery channel to successfully start
// the configuration recorder.
func (c *ConfigService) StartConfigurationRecorder(input *StartConfigurationRecorderInput) (*StartConfigurationRecorderOutput, error) {
	req, out := c.StartConfigurationRecorderRequest(input)
	err := req.Send()
	return out, err
}

const opStopConfigurationRecorder = "StopConfigurationRecorder"

// StopConfigurationRecorderRequest generates a request for the StopConfigurationRecorder operation.
func (c *ConfigService) StopConfigurationRecorderRequest(input *StopConfigurationRecorderInput) (req *aws.Request, output *StopConfigurationRecorderOutput) {
	op := &aws.Operation{
		Name:       opStopConfigurationRecorder,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopConfigurationRecorderInput{}
	}

	req = c.newRequest(op, input, output)
	output = &StopConfigurationRecorderOutput{}
	req.Data = output
	return
}

// Stops recording configurations of the AWS resources you have selected to
// record in your AWS account.
func (c *ConfigService) StopConfigurationRecorder(input *StopConfigurationRecorderInput) (*StopConfigurationRecorderOutput, error) {
	req, out := c.StopConfigurationRecorderRequest(input)
	err := req.Send()
	return out, err
}

// A list that contains the status of the delivery of either the snapshot or
// the configuration history to the specified Amazon S3 bucket.
type ConfigExportDeliveryInfo struct {
	// The time of the last attempted delivery.
	LastAttemptTime *time.Time `locationName:"lastAttemptTime" type:"timestamp" timestampFormat:"unix"`

	// The error code from the last attempted delivery.
	LastErrorCode *string `locationName:"lastErrorCode" type:"string"`

	// The error message from the last attempted delivery.
	LastErrorMessage *string `locationName:"lastErrorMessage" type:"string"`

	// Status of the last attempted delivery.
	LastStatus *string `locationName:"lastStatus" type:"string"`

	// The time of the last successful delivery.
	LastSuccessfulTime *time.Time `locationName:"lastSuccessfulTime" type:"timestamp" timestampFormat:"unix"`

	metadataConfigExportDeliveryInfo `json:"-" xml:"-"`
}

type metadataConfigExportDeliveryInfo struct {
	SDKShapeTraits bool `type:"structure"`
}

// A list that contains the status of the delivery of the configuration stream
// notification to the Amazon SNS topic.
type ConfigStreamDeliveryInfo struct {
	// The error code from the last attempted delivery.
	LastErrorCode *string `locationName:"lastErrorCode" type:"string"`

	// The error message from the last attempted delivery.
	LastErrorMessage *string `locationName:"lastErrorMessage" type:"string"`

	// Status of the last attempted delivery.
	//
	// Note Providing an SNS topic on a DeliveryChannel (http://docs.aws.amazon.com/config/latest/APIReference/API_DeliveryChannel.html)
	// for AWS Config is optional. If the SNS delivery is turned off, the last status
	// will be Not_Applicable.
	LastStatus *string `locationName:"lastStatus" type:"string"`

	// The time from the last status change.
	LastStatusChangeTime *time.Time `locationName:"lastStatusChangeTime" type:"timestamp" timestampFormat:"unix"`

	metadataConfigStreamDeliveryInfo `json:"-" xml:"-"`
}

type metadataConfigStreamDeliveryInfo struct {
	SDKShapeTraits bool `type:"structure"`
}

// A list that contains detailed configurations of a specified resource.
//
//  Currently, the list does not contain information about non-AWS components
// (for example, applications on your Amazon EC2 instances).
type ConfigurationItem struct {
	// The Amazon Resource Name (ARN) of the resource.
	ARN *string `locationName:"arn" type:"string"`

	// The 12 digit AWS account ID associated with the resource.
	AccountID *string `locationName:"accountId" type:"string"`

	// The Availability Zone associated with the resource.
	AvailabilityZone *string `locationName:"availabilityZone" type:"string"`

	// The description of the resource configuration.
	Configuration *string `locationName:"configuration" type:"string"`

	// The time when the configuration recording was initiated.
	ConfigurationItemCaptureTime *time.Time `locationName:"configurationItemCaptureTime" type:"timestamp" timestampFormat:"unix"`

	// Unique MD5 hash that represents the configuration item's state.
	//
	// You can use MD5 hash to compare the states of two or more configuration
	// items that are associated with the same resource.
	ConfigurationItemMD5Hash *string `locationName:"configurationItemMD5Hash" type:"string"`

	// The configuration item status.
	ConfigurationItemStatus *string `locationName:"configurationItemStatus" type:"string"`

	// An identifier that indicates the ordering of the configuration items of a
	// resource.
	ConfigurationStateID *string `locationName:"configurationStateId" type:"string"`

	// A list of CloudTrail event IDs.
	//
	// A populated field indicates that the current configuration was initiated
	// by the events recorded in the CloudTrail log. For more information about
	// CloudTrail, see What is AWS CloudTrail? (http://docs.aws.amazon.com/awscloudtrail/latest/userguide/what_is_cloud_trail_top_level.html).
	//
	// An empty field indicates that the current configuration was not initiated
	// by any event.
	RelatedEvents []*string `locationName:"relatedEvents" type:"list"`

	// A list of related AWS resources.
	Relationships []*Relationship `locationName:"relationships" type:"list"`

	// The time stamp when the resource was created.
	ResourceCreationTime *time.Time `locationName:"resourceCreationTime" type:"timestamp" timestampFormat:"unix"`

	// The ID of the resource (for example., sg-xxxxxx).
	ResourceID *string `locationName:"resourceId" type:"string"`

	// The type of AWS resource.
	ResourceType *string `locationName:"resourceType" type:"string"`

	// A mapping of key value tags associated with the resource.
	Tags map[string]*string `locationName:"tags" type:"map"`

	// The version number of the resource configuration.
	Version *string `locationName:"version" type:"string"`

	metadataConfigurationItem `json:"-" xml:"-"`
}

type metadataConfigurationItem struct {
	SDKShapeTraits bool `type:"structure"`
}

// An object that represents the recording of configuration changes of an AWS
// resource.
type ConfigurationRecorder struct {
	// The name of the recorder. By default, AWS Config automatically assigns the
	// name "default" when creating the configuration recorder. You cannot change
	// the assigned name.
	Name *string `locationName:"name" type:"string"`

	// The recording group specifies either to record configurations for all supported
	// resources or to provide a list of resource types to record. The list of resource
	// types must be a subset of supported resource types.
	RecordingGroup *RecordingGroup `locationName:"recordingGroup" type:"structure"`

	// Amazon Resource Name (ARN) of the IAM role used to describe the AWS resources
	// associated with the account.
	RoleARN *string `locationName:"roleARN" type:"string"`

	metadataConfigurationRecorder `json:"-" xml:"-"`
}

type metadataConfigurationRecorder struct {
	SDKShapeTraits bool `type:"structure"`
}

// The current status of the configuration recorder.
type ConfigurationRecorderStatus struct {
	// The error code indicating that the recording failed.
	LastErrorCode *string `locationName:"lastErrorCode" type:"string"`

	// The message indicating that the recording failed due to an error.
	LastErrorMessage *string `locationName:"lastErrorMessage" type:"string"`

	// The time the recorder was last started.
	LastStartTime *time.Time `locationName:"lastStartTime" type:"timestamp" timestampFormat:"unix"`

	// The last (previous) status of the recorder.
	LastStatus *string `locationName:"lastStatus" type:"string"`

	// The time when the status was last changed.
	LastStatusChangeTime *time.Time `locationName:"lastStatusChangeTime" type:"timestamp" timestampFormat:"unix"`

	// The time the recorder was last stopped.
	LastStopTime *time.Time `locationName:"lastStopTime" type:"timestamp" timestampFormat:"unix"`

	// The name of the configuration recorder.
	Name *string `locationName:"name" type:"string"`

	// Specifies whether the recorder is currently recording or not.
	Recording *bool `locationName:"recording" type:"boolean"`

	metadataConfigurationRecorderStatus `json:"-" xml:"-"`
}

type metadataConfigurationRecorderStatus struct {
	SDKShapeTraits bool `type:"structure"`
}

// The input for the DeleteDeliveryChannel action. The action accepts the following
// data in JSON format.
type DeleteDeliveryChannelInput struct {
	// The name of the delivery channel to delete.
	DeliveryChannelName *string `type:"string" required:"true"`

	metadataDeleteDeliveryChannelInput `json:"-" xml:"-"`
}

type metadataDeleteDeliveryChannelInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteDeliveryChannelOutput struct {
	metadataDeleteDeliveryChannelOutput `json:"-" xml:"-"`
}

type metadataDeleteDeliveryChannelOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The input for the DeliverConfigSnapshot action.
type DeliverConfigSnapshotInput struct {
	// The name of the delivery channel through which the snapshot is delivered.
	DeliveryChannelName *string `locationName:"deliveryChannelName" type:"string" required:"true"`

	metadataDeliverConfigSnapshotInput `json:"-" xml:"-"`
}

type metadataDeliverConfigSnapshotInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The output for the DeliverConfigSnapshot action in JSON format.
type DeliverConfigSnapshotOutput struct {
	// The ID of the snapshot that is being created.
	ConfigSnapshotID *string `locationName:"configSnapshotId" type:"string"`

	metadataDeliverConfigSnapshotOutput `json:"-" xml:"-"`
}

type metadataDeliverConfigSnapshotOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// A logical container used for storing the configuration changes of an AWS
// resource.
type DeliveryChannel struct {
	// The name of the delivery channel. By default, AWS Config automatically assigns
	// the name "default" when creating the delivery channel. You cannot change
	// the assigned name.
	Name *string `locationName:"name" type:"string"`

	// The name of the Amazon S3 bucket used to store configuration history for
	// the delivery channel.
	S3BucketName *string `locationName:"s3BucketName" type:"string"`

	// The prefix for the specified Amazon S3 bucket.
	S3KeyPrefix *string `locationName:"s3KeyPrefix" type:"string"`

	// The Amazon Resource Name (ARN) of the IAM role used for accessing the Amazon
	// S3 bucket and the Amazon SNS topic.
	SNSTopicARN *string `locationName:"snsTopicARN" type:"string"`

	metadataDeliveryChannel `json:"-" xml:"-"`
}

type metadataDeliveryChannel struct {
	SDKShapeTraits bool `type:"structure"`
}

// The status of a specified delivery channel.
//
// Valid values: Success | Failure
type DeliveryChannelStatus struct {
	// A list that contains the status of the delivery of the configuration history
	// to the specified Amazon S3 bucket.
	ConfigHistoryDeliveryInfo *ConfigExportDeliveryInfo `locationName:"configHistoryDeliveryInfo" type:"structure"`

	// A list containing the status of the delivery of the snapshot to the specified
	// Amazon S3 bucket.
	ConfigSnapshotDeliveryInfo *ConfigExportDeliveryInfo `locationName:"configSnapshotDeliveryInfo" type:"structure"`

	// A list containing the status of the delivery of the configuration stream
	// notification to the specified Amazon SNS topic.
	ConfigStreamDeliveryInfo *ConfigStreamDeliveryInfo `locationName:"configStreamDeliveryInfo" type:"structure"`

	// The name of the delivery channel.
	Name *string `locationName:"name" type:"string"`

	metadataDeliveryChannelStatus `json:"-" xml:"-"`
}

type metadataDeliveryChannelStatus struct {
	SDKShapeTraits bool `type:"structure"`
}

// The input for the DescribeConfigurationRecorderStatus action.
type DescribeConfigurationRecorderStatusInput struct {
	// The name(s) of the configuration recorder. If the name is not specified,
	// the action returns the current status of all the configuration recorders
	// associated with the account.
	ConfigurationRecorderNames []*string `type:"list"`

	metadataDescribeConfigurationRecorderStatusInput `json:"-" xml:"-"`
}

type metadataDescribeConfigurationRecorderStatusInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The output for the DescribeConfigurationRecorderStatus action in JSON format.
type DescribeConfigurationRecorderStatusOutput struct {
	// A list that contains status of the specified recorders.
	ConfigurationRecordersStatus []*ConfigurationRecorderStatus `type:"list"`

	metadataDescribeConfigurationRecorderStatusOutput `json:"-" xml:"-"`
}

type metadataDescribeConfigurationRecorderStatusOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The input for the DescribeConfigurationRecorders action.
type DescribeConfigurationRecordersInput struct {
	// A list of configuration recorder names.
	ConfigurationRecorderNames []*string `type:"list"`

	metadataDescribeConfigurationRecordersInput `json:"-" xml:"-"`
}

type metadataDescribeConfigurationRecordersInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The output for the DescribeConfigurationRecorders action.
type DescribeConfigurationRecordersOutput struct {
	// A list that contains the descriptions of the specified configuration recorders.
	ConfigurationRecorders []*ConfigurationRecorder `type:"list"`

	metadataDescribeConfigurationRecordersOutput `json:"-" xml:"-"`
}

type metadataDescribeConfigurationRecordersOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The input for the DeliveryChannelStatus action.
type DescribeDeliveryChannelStatusInput struct {
	// A list of delivery channel names.
	DeliveryChannelNames []*string `type:"list"`

	metadataDescribeDeliveryChannelStatusInput `json:"-" xml:"-"`
}

type metadataDescribeDeliveryChannelStatusInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The output for the DescribeDeliveryChannelStatus action.
type DescribeDeliveryChannelStatusOutput struct {
	// A list that contains the status of a specified delivery channel.
	DeliveryChannelsStatus []*DeliveryChannelStatus `type:"list"`

	metadataDescribeDeliveryChannelStatusOutput `json:"-" xml:"-"`
}

type metadataDescribeDeliveryChannelStatusOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The input for the DescribeDeliveryChannels action.
type DescribeDeliveryChannelsInput struct {
	// A list of delivery channel names.
	DeliveryChannelNames []*string `type:"list"`

	metadataDescribeDeliveryChannelsInput `json:"-" xml:"-"`
}

type metadataDescribeDeliveryChannelsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The output for the DescribeDeliveryChannels action.
type DescribeDeliveryChannelsOutput struct {
	// A list that contains the descriptions of the specified delivery channel.
	DeliveryChannels []*DeliveryChannel `type:"list"`

	metadataDescribeDeliveryChannelsOutput `json:"-" xml:"-"`
}

type metadataDescribeDeliveryChannelsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The input for the GetResourceConfigHistory action.
type GetResourceConfigHistoryInput struct {
	// The chronological order for configuration items listed. By default the results
	// are listed in reverse chronological order.
	ChronologicalOrder *string `locationName:"chronologicalOrder" type:"string"`

	// The time stamp that indicates an earlier time. If not specified, the action
	// returns paginated results that contain configuration items that start from
	// when the first configuration item was recorded.
	EarlierTime *time.Time `locationName:"earlierTime" type:"timestamp" timestampFormat:"unix"`

	// The time stamp that indicates a later time. If not specified, current time
	// is taken.
	LaterTime *time.Time `locationName:"laterTime" type:"timestamp" timestampFormat:"unix"`

	// The maximum number of configuration items returned in each page. The default
	// is 10. You cannot specify a limit greater than 100.
	Limit *int64 `locationName:"limit" type:"integer"`

	// An optional parameter used for pagination of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The ID of the resource (for example., sg-xxxxxx).
	ResourceID *string `locationName:"resourceId" type:"string" required:"true"`

	// The resource type.
	ResourceType *string `locationName:"resourceType" type:"string" required:"true"`

	metadataGetResourceConfigHistoryInput `json:"-" xml:"-"`
}

type metadataGetResourceConfigHistoryInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The output for the GetResourceConfigHistory action.
type GetResourceConfigHistoryOutput struct {
	// A list that contains the configuration history of one or more resources.
	ConfigurationItems []*ConfigurationItem `locationName:"configurationItems" type:"list"`

	// A token used for pagination of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataGetResourceConfigHistoryOutput `json:"-" xml:"-"`
}

type metadataGetResourceConfigHistoryOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The input for the PutConfigurationRecorder action.
type PutConfigurationRecorderInput struct {
	// The configuration recorder object that records each configuration change
	// made to the resources.
	ConfigurationRecorder *ConfigurationRecorder `type:"structure" required:"true"`

	metadataPutConfigurationRecorderInput `json:"-" xml:"-"`
}

type metadataPutConfigurationRecorderInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutConfigurationRecorderOutput struct {
	metadataPutConfigurationRecorderOutput `json:"-" xml:"-"`
}

type metadataPutConfigurationRecorderOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The input for the PutDeliveryChannel action.
type PutDeliveryChannelInput struct {
	// The configuration delivery channel object that delivers the configuration
	// information to an Amazon S3 bucket, and to an Amazon SNS topic.
	DeliveryChannel *DeliveryChannel `type:"structure" required:"true"`

	metadataPutDeliveryChannelInput `json:"-" xml:"-"`
}

type metadataPutDeliveryChannelInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutDeliveryChannelOutput struct {
	metadataPutDeliveryChannelOutput `json:"-" xml:"-"`
}

type metadataPutDeliveryChannelOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The group of AWS resource types that AWS Config records when starting the
// configuration recorder.
//
// recordingGroup can have one and only one parameter. Choose either allSupported
// or resourceTypes.
type RecordingGroup struct {
	// Records all supported resource types in the recording group. For a list of
	// supported resource types, see Supported resource types (http://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources).
	// If you specify allSupported, you cannot enumerate a list of resourceTypes.
	AllSupported *bool `locationName:"allSupported" type:"boolean"`

	// A comma-separated list of strings representing valid AWS resource types (e.g.,
	// AWS::EC2::Instance or AWS::CloudTrail::Trail). resourceTypes is only valid
	// if you have chosen not to select allSupported. For a list of valid resourceTypes
	// values, see the resourceType Value column in the following topic: Supported
	// AWS Resource Types (http://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources).
	ResourceTypes []*string `locationName:"resourceTypes" type:"list"`

	metadataRecordingGroup `json:"-" xml:"-"`
}

type metadataRecordingGroup struct {
	SDKShapeTraits bool `type:"structure"`
}

// The relationship of the related resource to the main resource.
type Relationship struct {
	// The name of the related resource.
	RelationshipName *string `locationName:"relationshipName" type:"string"`

	// The resource ID of the related resource (for example, sg-xxxxxx).
	ResourceID *string `locationName:"resourceId" type:"string"`

	// The resource type of the related resource.
	ResourceType *string `locationName:"resourceType" type:"string"`

	metadataRelationship `json:"-" xml:"-"`
}

type metadataRelationship struct {
	SDKShapeTraits bool `type:"structure"`
}

// The input for the StartConfigurationRecorder action.
type StartConfigurationRecorderInput struct {
	// The name of the recorder object that records each configuration change made
	// to the resources.
	ConfigurationRecorderName *string `type:"string" required:"true"`

	metadataStartConfigurationRecorderInput `json:"-" xml:"-"`
}

type metadataStartConfigurationRecorderInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type StartConfigurationRecorderOutput struct {
	metadataStartConfigurationRecorderOutput `json:"-" xml:"-"`
}

type metadataStartConfigurationRecorderOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The input for the StopConfigurationRecorder action.
type StopConfigurationRecorderInput struct {
	// The name of the recorder object that records each configuration change made
	// to the resources.
	ConfigurationRecorderName *string `type:"string" required:"true"`

	metadataStopConfigurationRecorderInput `json:"-" xml:"-"`
}

type metadataStopConfigurationRecorderInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type StopConfigurationRecorderOutput struct {
	metadataStopConfigurationRecorderOutput `json:"-" xml:"-"`
}

type metadataStopConfigurationRecorderOutput struct {
	SDKShapeTraits bool `type:"structure"`
}
