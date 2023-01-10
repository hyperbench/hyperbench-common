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
 * @brief Define the structure related to statistics
 * @file result.go
 * @author: wangxiaohui
 * @date 2022-01-18
 */

// Report contains need report data about tx been executed info.
type Report struct {
	Cur *Data
	Sum *Data
}

// DataType data type.
type DataType string

const (
	// Cur current data type.
	Cur DataType = "Cur"
	// Sum sum data type.
	Sum DataType = "Sum"
)

// Data define the data which will be reported.
type Data struct {
	// Type data type.
	Type DataType
	// Results the slice of AggData.
	Results []AggData
}

// AggData aggregation data.
type AggData struct {
	// key
	Label    string
	Time     int64
	Duration int64

	// request
	Num      int
	Statuses map[Status]int

	// latency
	Send    Latency
	Confirm Latency
	Write   Latency
}

// Latency is the percent of latency in increase order
type Latency struct {
	Avg  int64 // Avg is the average latency
	P0   int64 // P0  is the minimal latency
	P50  int64 // P50 is the median of latency
	P90  int64 // P90 is the latency exactly larger than 90%
	P95  int64 // P90 is the latency exactly larger than 95%
	P99  int64 // P99 is the latency exactly larger than 95%
	P100 int64 // P100 is the maximal latency
}

// RemoteStatistic remote statistic data.
type RemoteStatistic struct {
	// time info
	Start int64
	End   int64
	// block info
	BlockNum int
	// transaction info of blockchain
	TxNum int
	// tps of blockchain
	CTps float64
	// bps of blockchain
	Bps float64
	// sent transaction info
	SentTx int64
	// missed transaction info
	MissedTx int64
	// tps of system
	Tps float64
}
