package cloudformation

// AWSRDSOptionGroup_OptionConfiguration AWS CloudFormation Resource (AWS::RDS::OptionGroup.OptionConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html
type AWSRDSOptionGroup_OptionConfiguration struct {

	// DBSecurityGroupMemberships AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-dbsecuritygroupmemberships
	DBSecurityGroupMemberships []string `json:"DBSecurityGroupMemberships,omitempty"`

	// OptionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-optionname
	OptionName string `json:"OptionName,omitempty"`

	// OptionSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-optionsettings
	OptionSettings *AWSRDSOptionGroup_OptionSetting `json:"OptionSettings,omitempty"`

	// Port AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-port
	Port int `json:"Port,omitempty"`

	// VpcSecurityGroupMemberships AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-vpcsecuritygroupmemberships
	VpcSecurityGroupMemberships []string `json:"VpcSecurityGroupMemberships,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSOptionGroup_OptionConfiguration) AWSCloudFormationType() string {
	return "AWS::RDS::OptionGroup.OptionConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSOptionGroup_OptionConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
