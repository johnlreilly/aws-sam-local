package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSStepFunctionsActivity AWS CloudFormation Resource (AWS::StepFunctions::Activity)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html
type AWSStepFunctionsActivity struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html#cfn-stepfunctions-activity-name
	Name string `json:"Name,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSStepFunctionsActivity) AWSCloudFormationType() string {
	return "AWS::StepFunctions::Activity"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSStepFunctionsActivity) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSStepFunctionsActivity) MarshalJSON() ([]byte, error) {
	type Properties AWSStepFunctionsActivity
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSStepFunctionsActivity) UnmarshalJSON(b []byte) error {
	type Properties AWSStepFunctionsActivity
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSStepFunctionsActivity(*res.Properties)
	}

	return nil
}

// GetAllAWSStepFunctionsActivityResources retrieves all AWSStepFunctionsActivity items from an AWS CloudFormation template
func (t *Template) GetAllAWSStepFunctionsActivityResources() map[string]AWSStepFunctionsActivity {
	results := map[string]AWSStepFunctionsActivity{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSStepFunctionsActivity:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::StepFunctions::Activity" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSStepFunctionsActivity
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSStepFunctionsActivityWithName retrieves all AWSStepFunctionsActivity items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSStepFunctionsActivityWithName(name string) (AWSStepFunctionsActivity, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSStepFunctionsActivity:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::StepFunctions::Activity" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSStepFunctionsActivity
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSStepFunctionsActivity{}, errors.New("resource not found")
}
