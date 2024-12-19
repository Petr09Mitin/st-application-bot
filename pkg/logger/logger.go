package logger

import "go.uber.org/zap"

func NewZapLogger(outputPaths []string) (logger *zap.Logger, err error) {
	cfg := zap.NewProductionConfig()

	cfg.Sampling = nil

	cfg.OutputPaths = outputPaths

	logger, err = cfg.Build()
	if err != nil {
		return
	}

	return
}
