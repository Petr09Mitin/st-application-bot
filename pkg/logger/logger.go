package logger

import "go.uber.org/zap"

var (
	DefaultOutputPaths = []string{
		"stderr",
	}
)

func NewZapLogger(outputPaths []string) (logger *zap.Logger, err error) {
	if outputPaths == nil {
		outputPaths = DefaultOutputPaths
	}

	cfg := zap.NewProductionConfig()

	cfg.Sampling = nil

	cfg.OutputPaths = outputPaths

	logger, err = cfg.Build()
	if err != nil {
		return
	}

	return
}
