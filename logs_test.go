package gohan_logs_benchmark

import (
	"testing"

	"github.com/cloudwan/gohan/log"
	"github.com/cloudwan/gohan/util"
)

func BenchmarkNewLogger(b *testing.B) {
	for n := 0; n < b.N; n++ {
		log.NewLogger()
	}
}

func BenchmarkNewLoggerForModule(b *testing.B) {
	for n := 0; n < b.N; n++ {
		log.NewLoggerForModule("gohan.module.name")
	}
}

func BenchmarkLogNoFormat(b *testing.B) {
	cfg := util.GetConfig()
	if err := cfg.ReadConfig("devNullLog.yaml"); err != nil {
		panic(err)
	}
	log.SetUpLogging(cfg)
	logger := log.NewLoggerForModule("gohan.module.name")

	for n := 0; n < b.N; n++ {
		logger.Info("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")
	}
}

func BenchmarkLogWithFormat(b *testing.B) {
	cfg := util.GetConfig()
	if err := cfg.ReadConfig("devNullLog.yaml"); err != nil {
		panic(err)
	}
	log.SetUpLogging(cfg)
	logger := log.NewLoggerForModule("gohan.module.name")

	for n := 0; n < b.N; n++ {
		logger.Info("Number: %d, string: %s", 42, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")
	}
}

func BenchmarkEachLogCreatesNewLogger(b *testing.B) {
	cfg := util.GetConfig()
	if err := cfg.ReadConfig("devNullLog.yaml"); err != nil {
		panic(err)
	}
	log.SetUpLogging(cfg)

	for n := 0; n < b.N; n++ {
		logger := log.NewLogger()
		logger.Info("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")
	}
}

func BenchmarkEachLogCreatesNewLoggerForModule(b *testing.B) {
	cfg := util.GetConfig()
	if err := cfg.ReadConfig("devNullLog.yaml"); err != nil {
		panic(err)
	}
	log.SetUpLogging(cfg)

	for n := 0; n < b.N; n++ {
		logger := log.NewLoggerForModule("gohan.module.name")
		logger.Info("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")
	}
}
