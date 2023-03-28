// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"github.com/protocol-laboratory/kafka-codec-go/codec"
	"github.com/protocol-laboratory/kafka-codec-go/knet"
)

func main() {
	cli, err := knet.NewKafkaNetClient(knet.KafkaNetClientConfig{
		Host: "localhost",
		Port: 9092,
	})
	if err != nil {
		panic(err)
	}
	apiVersions, err := cli.ApiVersions(&codec.ApiReq{
		BaseReq: codec.BaseReq{
			ApiVersion:    0,
			CorrelationId: 0,
			ClientId:      "",
		},
		ClientSoftwareName:    "kafka-codec-go",
		ClientSoftwareVersion: "0.0.1",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", apiVersions)
}
