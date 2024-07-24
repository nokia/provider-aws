// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package convert

import types "github.com/aws/aws-sdk-go-v2/service/s3/types"

type ConverterImpl struct{}

func (c *ConverterImpl) DeepCopyAWSLambdaFunctionConfiguration(source []types.LambdaFunctionConfiguration) []types.LambdaFunctionConfiguration {
	var typesLambdaFunctionConfigurationList []types.LambdaFunctionConfiguration
	if source != nil {
		typesLambdaFunctionConfigurationList = make([]types.LambdaFunctionConfiguration, len(source))
		for i := 0; i < len(source); i++ {
			typesLambdaFunctionConfigurationList[i] = c.typesLambdaFunctionConfigurationToTypesLambdaFunctionConfiguration(source[i])
		}
	}
	return typesLambdaFunctionConfigurationList
}
func (c *ConverterImpl) DeepCopyAWSQueueConfiguration(source []types.QueueConfiguration) []types.QueueConfiguration {
	var typesQueueConfigurationList []types.QueueConfiguration
	if source != nil {
		typesQueueConfigurationList = make([]types.QueueConfiguration, len(source))
		for i := 0; i < len(source); i++ {
			typesQueueConfigurationList[i] = c.typesQueueConfigurationToTypesQueueConfiguration(source[i])
		}
	}
	return typesQueueConfigurationList
}
func (c *ConverterImpl) DeepCopyAWSTopicConfiguration(source []types.TopicConfiguration) []types.TopicConfiguration {
	var typesTopicConfigurationList []types.TopicConfiguration
	if source != nil {
		typesTopicConfigurationList = make([]types.TopicConfiguration, len(source))
		for i := 0; i < len(source); i++ {
			typesTopicConfigurationList[i] = c.typesTopicConfigurationToTypesTopicConfiguration(source[i])
		}
	}
	return typesTopicConfigurationList
}
func (c *ConverterImpl) pTypesNotificationConfigurationFilterToPTypesNotificationConfigurationFilter(source *types.NotificationConfigurationFilter) *types.NotificationConfigurationFilter {
	var pTypesNotificationConfigurationFilter *types.NotificationConfigurationFilter
	if source != nil {
		var typesNotificationConfigurationFilter types.NotificationConfigurationFilter
		typesNotificationConfigurationFilter.Key = c.pTypesS3KeyFilterToPTypesS3KeyFilter((*source).Key)
		pTypesNotificationConfigurationFilter = &typesNotificationConfigurationFilter
	}
	return pTypesNotificationConfigurationFilter
}
func (c *ConverterImpl) pTypesS3KeyFilterToPTypesS3KeyFilter(source *types.S3KeyFilter) *types.S3KeyFilter {
	var pTypesS3KeyFilter *types.S3KeyFilter
	if source != nil {
		var typesS3KeyFilter types.S3KeyFilter
		var typesFilterRuleList []types.FilterRule
		if (*source).FilterRules != nil {
			typesFilterRuleList = make([]types.FilterRule, len((*source).FilterRules))
			for i := 0; i < len((*source).FilterRules); i++ {
				typesFilterRuleList[i] = c.typesFilterRuleToTypesFilterRule((*source).FilterRules[i])
			}
		}
		typesS3KeyFilter.FilterRules = typesFilterRuleList
		pTypesS3KeyFilter = &typesS3KeyFilter
	}
	return pTypesS3KeyFilter
}
func (c *ConverterImpl) typesFilterRuleToTypesFilterRule(source types.FilterRule) types.FilterRule {
	var typesFilterRule types.FilterRule
	typesFilterRule.Name = types.FilterRuleName(source.Name)
	var pString *string
	if source.Value != nil {
		xstring := *source.Value
		pString = &xstring
	}
	typesFilterRule.Value = pString
	return typesFilterRule
}
func (c *ConverterImpl) typesLambdaFunctionConfigurationToTypesLambdaFunctionConfiguration(source types.LambdaFunctionConfiguration) types.LambdaFunctionConfiguration {
	var typesLambdaFunctionConfiguration types.LambdaFunctionConfiguration
	var typesEventList []types.Event
	if source.Events != nil {
		typesEventList = make([]types.Event, len(source.Events))
		for i := 0; i < len(source.Events); i++ {
			typesEventList[i] = types.Event(source.Events[i])
		}
	}
	typesLambdaFunctionConfiguration.Events = typesEventList
	var pString *string
	if source.LambdaFunctionArn != nil {
		xstring := *source.LambdaFunctionArn
		pString = &xstring
	}
	typesLambdaFunctionConfiguration.LambdaFunctionArn = pString
	typesLambdaFunctionConfiguration.Filter = c.pTypesNotificationConfigurationFilterToPTypesNotificationConfigurationFilter(source.Filter)
	var pString2 *string
	if source.Id != nil {
		xstring2 := *source.Id
		pString2 = &xstring2
	}
	typesLambdaFunctionConfiguration.Id = pString2
	return typesLambdaFunctionConfiguration
}
func (c *ConverterImpl) typesQueueConfigurationToTypesQueueConfiguration(source types.QueueConfiguration) types.QueueConfiguration {
	var typesQueueConfiguration types.QueueConfiguration
	var typesEventList []types.Event
	if source.Events != nil {
		typesEventList = make([]types.Event, len(source.Events))
		for i := 0; i < len(source.Events); i++ {
			typesEventList[i] = types.Event(source.Events[i])
		}
	}
	typesQueueConfiguration.Events = typesEventList
	var pString *string
	if source.QueueArn != nil {
		xstring := *source.QueueArn
		pString = &xstring
	}
	typesQueueConfiguration.QueueArn = pString
	typesQueueConfiguration.Filter = c.pTypesNotificationConfigurationFilterToPTypesNotificationConfigurationFilter(source.Filter)
	var pString2 *string
	if source.Id != nil {
		xstring2 := *source.Id
		pString2 = &xstring2
	}
	typesQueueConfiguration.Id = pString2
	return typesQueueConfiguration
}
func (c *ConverterImpl) typesTopicConfigurationToTypesTopicConfiguration(source types.TopicConfiguration) types.TopicConfiguration {
	var typesTopicConfiguration types.TopicConfiguration
	var typesEventList []types.Event
	if source.Events != nil {
		typesEventList = make([]types.Event, len(source.Events))
		for i := 0; i < len(source.Events); i++ {
			typesEventList[i] = types.Event(source.Events[i])
		}
	}
	typesTopicConfiguration.Events = typesEventList
	var pString *string
	if source.TopicArn != nil {
		xstring := *source.TopicArn
		pString = &xstring
	}
	typesTopicConfiguration.TopicArn = pString
	typesTopicConfiguration.Filter = c.pTypesNotificationConfigurationFilterToPTypesNotificationConfigurationFilter(source.Filter)
	var pString2 *string
	if source.Id != nil {
		xstring2 := *source.Id
		pString2 = &xstring2
	}
	typesTopicConfiguration.Id = pString2
	return typesTopicConfiguration
}