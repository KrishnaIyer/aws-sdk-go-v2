// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AMITypes string

// Enum values for AMITypes
const (
	AMITypesAl2X8664    AMITypes = "AL2_x86_64"
	AMITypesAl2X8664Gpu AMITypes = "AL2_x86_64_GPU"
	AMITypesAl2Arm64    AMITypes = "AL2_ARM_64"
)

// Values returns all known values for AMITypes. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AMITypes) Values() []AMITypes {
	return []AMITypes{
		"AL2_x86_64",
		"AL2_x86_64_GPU",
		"AL2_ARM_64",
	}
}

type ClusterStatus string

// Enum values for ClusterStatus
const (
	ClusterStatusCreating ClusterStatus = "CREATING"
	ClusterStatusActive   ClusterStatus = "ACTIVE"
	ClusterStatusDeleting ClusterStatus = "DELETING"
	ClusterStatusFailed   ClusterStatus = "FAILED"
	ClusterStatusUpdating ClusterStatus = "UPDATING"
)

// Values returns all known values for ClusterStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ClusterStatus) Values() []ClusterStatus {
	return []ClusterStatus{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeSubnetNotFound            ErrorCode = "SubnetNotFound"
	ErrorCodeSecurityGroupNotFound     ErrorCode = "SecurityGroupNotFound"
	ErrorCodeEniLimitReached           ErrorCode = "EniLimitReached"
	ErrorCodeIpNotAvailable            ErrorCode = "IpNotAvailable"
	ErrorCodeAccessDenied              ErrorCode = "AccessDenied"
	ErrorCodeOperationNotPermitted     ErrorCode = "OperationNotPermitted"
	ErrorCodeVpcIdNotFound             ErrorCode = "VpcIdNotFound"
	ErrorCodeUnknown                   ErrorCode = "Unknown"
	ErrorCodeNodeCreationFailure       ErrorCode = "NodeCreationFailure"
	ErrorCodePodEvictionFailure        ErrorCode = "PodEvictionFailure"
	ErrorCodeInsufficientFreeAddresses ErrorCode = "InsufficientFreeAddresses"
	ErrorCodeClusterUnreachable        ErrorCode = "ClusterUnreachable"
)

// Values returns all known values for ErrorCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"SubnetNotFound",
		"SecurityGroupNotFound",
		"EniLimitReached",
		"IpNotAvailable",
		"AccessDenied",
		"OperationNotPermitted",
		"VpcIdNotFound",
		"Unknown",
		"NodeCreationFailure",
		"PodEvictionFailure",
		"InsufficientFreeAddresses",
		"ClusterUnreachable",
	}
}

type FargateProfileStatus string

// Enum values for FargateProfileStatus
const (
	FargateProfileStatusCreating     FargateProfileStatus = "CREATING"
	FargateProfileStatusActive       FargateProfileStatus = "ACTIVE"
	FargateProfileStatusDeleting     FargateProfileStatus = "DELETING"
	FargateProfileStatusCreateFailed FargateProfileStatus = "CREATE_FAILED"
	FargateProfileStatusDeleteFailed FargateProfileStatus = "DELETE_FAILED"
)

// Values returns all known values for FargateProfileStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FargateProfileStatus) Values() []FargateProfileStatus {
	return []FargateProfileStatus{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"CREATE_FAILED",
		"DELETE_FAILED",
	}
}

type LogType string

// Enum values for LogType
const (
	LogTypeApi               LogType = "api"
	LogTypeAudit             LogType = "audit"
	LogTypeAuthenticator     LogType = "authenticator"
	LogTypeControllerManager LogType = "controllerManager"
	LogTypeScheduler         LogType = "scheduler"
)

// Values returns all known values for LogType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LogType) Values() []LogType {
	return []LogType{
		"api",
		"audit",
		"authenticator",
		"controllerManager",
		"scheduler",
	}
}

type NodegroupIssueCode string

// Enum values for NodegroupIssueCode
const (
	NodegroupIssueCodeAutoScalingGroupNotFound             NodegroupIssueCode = "AutoScalingGroupNotFound"
	NodegroupIssueCodeAutoScalingGroupInvalidConfiguration NodegroupIssueCode = "AutoScalingGroupInvalidConfiguration"
	NodegroupIssueCodeEc2SecurityGroupNotFound             NodegroupIssueCode = "Ec2SecurityGroupNotFound"
	NodegroupIssueCodeEc2SecurityGroupDeletionFailure      NodegroupIssueCode = "Ec2SecurityGroupDeletionFailure"
	NodegroupIssueCodeEc2LaunchTemplateNotFound            NodegroupIssueCode = "Ec2LaunchTemplateNotFound"
	NodegroupIssueCodeEc2LaunchTemplateVersionMismatch     NodegroupIssueCode = "Ec2LaunchTemplateVersionMismatch"
	NodegroupIssueCodeEc2SubnetNotFound                    NodegroupIssueCode = "Ec2SubnetNotFound"
	NodegroupIssueCodeEc2SubnetInvalidConfiguration        NodegroupIssueCode = "Ec2SubnetInvalidConfiguration"
	NodegroupIssueCodeIamInstanceProfileNotFound           NodegroupIssueCode = "IamInstanceProfileNotFound"
	NodegroupIssueCodeIamLimitExceeded                     NodegroupIssueCode = "IamLimitExceeded"
	NodegroupIssueCodeIamNodeRoleNotFound                  NodegroupIssueCode = "IamNodeRoleNotFound"
	NodegroupIssueCodeNodeCreationFailure                  NodegroupIssueCode = "NodeCreationFailure"
	NodegroupIssueCodeAsgInstanceLaunchFailures            NodegroupIssueCode = "AsgInstanceLaunchFailures"
	NodegroupIssueCodeInstanceLimitExceeded                NodegroupIssueCode = "InstanceLimitExceeded"
	NodegroupIssueCodeInsufficientFreeAddresses            NodegroupIssueCode = "InsufficientFreeAddresses"
	NodegroupIssueCodeAccessDenied                         NodegroupIssueCode = "AccessDenied"
	NodegroupIssueCodeInternalFailure                      NodegroupIssueCode = "InternalFailure"
	NodegroupIssueCodeClusterUnreachable                   NodegroupIssueCode = "ClusterUnreachable"
)

// Values returns all known values for NodegroupIssueCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NodegroupIssueCode) Values() []NodegroupIssueCode {
	return []NodegroupIssueCode{
		"AutoScalingGroupNotFound",
		"AutoScalingGroupInvalidConfiguration",
		"Ec2SecurityGroupNotFound",
		"Ec2SecurityGroupDeletionFailure",
		"Ec2LaunchTemplateNotFound",
		"Ec2LaunchTemplateVersionMismatch",
		"Ec2SubnetNotFound",
		"Ec2SubnetInvalidConfiguration",
		"IamInstanceProfileNotFound",
		"IamLimitExceeded",
		"IamNodeRoleNotFound",
		"NodeCreationFailure",
		"AsgInstanceLaunchFailures",
		"InstanceLimitExceeded",
		"InsufficientFreeAddresses",
		"AccessDenied",
		"InternalFailure",
		"ClusterUnreachable",
	}
}

type NodegroupStatus string

// Enum values for NodegroupStatus
const (
	NodegroupStatusCreating     NodegroupStatus = "CREATING"
	NodegroupStatusActive       NodegroupStatus = "ACTIVE"
	NodegroupStatusUpdating     NodegroupStatus = "UPDATING"
	NodegroupStatusDeleting     NodegroupStatus = "DELETING"
	NodegroupStatusCreateFailed NodegroupStatus = "CREATE_FAILED"
	NodegroupStatusDeleteFailed NodegroupStatus = "DELETE_FAILED"
	NodegroupStatusDegraded     NodegroupStatus = "DEGRADED"
)

// Values returns all known values for NodegroupStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NodegroupStatus) Values() []NodegroupStatus {
	return []NodegroupStatus{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"CREATE_FAILED",
		"DELETE_FAILED",
		"DEGRADED",
	}
}

type UpdateParamType string

// Enum values for UpdateParamType
const (
	UpdateParamTypeVersion               UpdateParamType = "Version"
	UpdateParamTypePlatformVersion       UpdateParamType = "PlatformVersion"
	UpdateParamTypeEndpointPrivateAccess UpdateParamType = "EndpointPrivateAccess"
	UpdateParamTypeEndpointPublicAccess  UpdateParamType = "EndpointPublicAccess"
	UpdateParamTypeClusterLogging        UpdateParamType = "ClusterLogging"
	UpdateParamTypeDesiredSize           UpdateParamType = "DesiredSize"
	UpdateParamTypeLabelsToAdd           UpdateParamType = "LabelsToAdd"
	UpdateParamTypeLabelsToRemove        UpdateParamType = "LabelsToRemove"
	UpdateParamTypeMaxSize               UpdateParamType = "MaxSize"
	UpdateParamTypeMinSize               UpdateParamType = "MinSize"
	UpdateParamTypeReleaseVersion        UpdateParamType = "ReleaseVersion"
	UpdateParamTypePublicAccessCidrs     UpdateParamType = "PublicAccessCidrs"
)

// Values returns all known values for UpdateParamType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UpdateParamType) Values() []UpdateParamType {
	return []UpdateParamType{
		"Version",
		"PlatformVersion",
		"EndpointPrivateAccess",
		"EndpointPublicAccess",
		"ClusterLogging",
		"DesiredSize",
		"LabelsToAdd",
		"LabelsToRemove",
		"MaxSize",
		"MinSize",
		"ReleaseVersion",
		"PublicAccessCidrs",
	}
}

type UpdateStatus string

// Enum values for UpdateStatus
const (
	UpdateStatusInProgress UpdateStatus = "InProgress"
	UpdateStatusFailed     UpdateStatus = "Failed"
	UpdateStatusCancelled  UpdateStatus = "Cancelled"
	UpdateStatusSuccessful UpdateStatus = "Successful"
)

// Values returns all known values for UpdateStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UpdateStatus) Values() []UpdateStatus {
	return []UpdateStatus{
		"InProgress",
		"Failed",
		"Cancelled",
		"Successful",
	}
}

type UpdateType string

// Enum values for UpdateType
const (
	UpdateTypeVersionUpdate        UpdateType = "VersionUpdate"
	UpdateTypeEndpointAccessUpdate UpdateType = "EndpointAccessUpdate"
	UpdateTypeLoggingUpdate        UpdateType = "LoggingUpdate"
	UpdateTypeConfigUpdate         UpdateType = "ConfigUpdate"
)

// Values returns all known values for UpdateType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UpdateType) Values() []UpdateType {
	return []UpdateType{
		"VersionUpdate",
		"EndpointAccessUpdate",
		"LoggingUpdate",
		"ConfigUpdate",
	}
}
