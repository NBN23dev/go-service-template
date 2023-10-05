package tracer

import (
	"cloud.google.com/go/compute/metadata"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/trace"
)

func Init() error {
	projectId, err := metadata.ProjectID()

	if err != nil {
		return nil
	}

	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: projectId,
	})

	if err != nil {
		return err
	}

	trace.RegisterExporter(exporter)

	// By default, traces will be sampled relatively rarely. To change the
	// sampling frequency for your entire program, call ApplyConfig. Use a
	// ProbabilitySampler to sample a subset of traces, or use AlwaysSample to
	// collect a trace on every run.
	//
	// Be careful about using trace.AlwaysSample in a production application
	// with significant traffic: a new trace will be started and exported for
	// every request.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	return nil
}
