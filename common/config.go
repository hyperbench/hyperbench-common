package common

/**
 *  Copyright (C) 2021 HyperBench.
 *  SPDX-License-Identifier: Apache-2.0
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 * @brief define config key
 * @file common.go
 * @author: wangxiaohui
 * @date 2022-01-18
 */

// config path
const (
	BenchmarkDirPath     = "__BenchmarkDirPath__"
	BenchmarkArchivePath = "__BenchmarkArchivePath__"
	BenchmarkConfigPath  = "__BenchmarkConfigPath__"
)

// worker
const (
	WorkerPath    = "worker"
	WorkerCapPath = "worker.cap"
)

// workers
const (
	WorkersPath = "workers"
)

// collector mode
const (
	CollectorSummaryMode = "summary"
	CollectorDetailsMode = "details"
)

// collector
const (
	CollectorModePath = "collector.mode"
)

// engine
const (
	EnginePath         = "engine"
	EngineRatePath     = "engine.rate"
	EngineDurationPath = "engine.duration"
	EngineCapPath      = "engine.cap"
	EngineURLsPath     = "engine.urls"
	EngineInstantPath  = "engine.instant"
	EngineWaitPath     = "engine.wait"
)

// client
const (
	ClientScriptPath       = "client.script"
	ClientTypePath         = "client.type"
	ClientConfigPath       = "client.config"
	ClientContractPath     = "client.contract"
	ClientContractArgsPath = "client.args"
	ClientOptionPath       = "client.options"
	ClientPluginPath       = "client.plugin"
)

// verify
const (
	VerifyEnablePath     = "verify.enable"
	VerifyPercentagePath = "verify.percentage"
)

// recorder
const (
	// csv
	RecorderCsvPath    = "recorder.csv"
	RecorderCsvDirPath = "recorder.csv.dir"
	// log
	LogLevelPath = "recorder.log.level"
	LogDirPath   = "recorder.log.dir"
	LogDumpPath  = "recorder.log.dump"
	// influxdb
	RecorderInflucDBPath = "recorder.influxdb"
	InfluxDBUrlPath      = "recorder.influxdb.url"
	InfluxDBDatabasePath = "recorder.influxdb.database"
	InfluxDBUsernamePath = "recorder.influxdb.username"
	InfluxDBPasswordPath = "recorder.influxdb.password"
)

// logger
const ()
