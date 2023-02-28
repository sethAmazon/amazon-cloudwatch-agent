// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package host

import (
	"log"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"

	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/translate/otel/common"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/translate/otel/exporter/awscloudwatch"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/translate/otel/processor/cumulativetodeltaprocessor"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/translate/otel/processor/ec2taggerprocessor"
)

type translator struct {
	name      string
	receivers common.TranslatorMap[component.Config]
}

var _ common.Translator[*common.ComponentTranslators] = (*translator)(nil)

// NewTranslator creates a new host pipeline translator. The receiver types
// passed in are converted to config.ComponentIDs, sorted, and used directly
// in the translated pipeline.
func NewTranslator(
	name string,
	receivers common.TranslatorMap[component.Config],
) common.Translator[*common.ComponentTranslators] {
	return &translator{name, receivers}
}

func (t translator) ID() component.ID {
	return component.NewIDWithName(component.DataTypeMetrics, t.name)
}

// Translate creates a pipeline if metrics section exists.
func (t translator) Translate(conf *confmap.Conf) (*common.ComponentTranslators, error) {
	if conf == nil || !conf.IsSet(common.MetricsKey) {
		return nil, &common.MissingKeyError{ID: t.ID(), JsonKey: common.MetricsKey}
	} else if len(t.receivers) == 0 {
		log.Printf("D! pipeline %s has no receivers", t.name)
		return nil, nil
	}

	translators := common.ComponentTranslators{
		Receivers:  t.receivers,
		Processors: common.NewTranslatorMap[component.Config](),
		Exporters:  common.NewTranslatorMap(awscloudwatch.NewTranslator()),
	}

	// we need to add delta processor because (only) diskio and net input plugins report delta metric
	if common.PipelineNameHostDeltaMetrics == t.name {
		log.Printf("D! delta processor required because metrics with diskio or net are required")
		translators.Processors.Add(cumulativetodeltaprocessor.NewTranslatorWithName(t.name))
	}

	if conf.IsSet(common.ConfigKey(common.MetricsKey, "append_dimensions")) {
		log.Printf("D! ec2tagger processor required because append_dimensions is set")
		translators.Processors.Add(ec2taggerprocessor.NewTranslator())
	}
	return &translators, nil
}
