// firehose-nozzle-v2
//
// Copyright (c) 2018-Present Pivotal Software, Inc. All Rights Reserved.
//
// This program and the accompanying materials are made available under the terms of the under the Apache License,
// Version 2.0 (the "License”); you may not use this file except in compliance with the License. You may
// obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the
// License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing permissions and
// limitations under the License.

package nozzle

import "io"

type LogShipper interface {
	LogShip(string) error
}

type SampleShipper struct {
	writer io.Writer
}

func NewSampleShipper(writer io.Writer) LogShipper {
	return &SampleShipper{
		writer: writer,
	}
}

func (ss *SampleShipper) LogShip(log string) error {
	_, err := ss.writer.Write([]byte(log))
	if err != nil {
		return err
	}

	_, err = ss.writer.Write([]byte("\n"))
	return err
}
