package logger

import (
	"go.uber.org/zap"
)

func NewLogger(service string, env string) (*zap.Logger, error) {
	var base *zap.Logger
	var err error

	switch env {
	case "prod", "production":
		base, err = zap.NewProduction()
	default:
		base, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, err
	}

	return base.With(
		zap.String("service", service),
		zap.String("env", env),
	), nil
}
