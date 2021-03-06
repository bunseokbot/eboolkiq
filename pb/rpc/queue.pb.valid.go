// Copyright 2020 Danggeun Market Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

func (x *ListReq) CheckValid() error {
	if x == nil {
		return ErrNilRequest
	}
	return nil
}

func (x *GetReq) CheckValid() error {
	if x == nil {
		return ErrNilRequest
	}

	if x.Name == "" {
		return ErrEmptyQueueName
	}

	return nil
}

func (x *CreateReq) CheckValid() error {
	if x == nil {
		return ErrNilRequest
	}

	if x.Queue == nil {
		return ErrNilQueue
	}

	if x.Queue.Name == "" {
		return ErrEmptyQueueName
	}

	return nil
}

func (x *DeleteReq) CheckValid() error {
	if x == nil {
		return ErrNilRequest
	}

	if x.Name == "" {
		return ErrEmptyQueueName
	}

	return nil
}

func (x *UpdateReq) CheckValid() error {
	if x == nil {
		return ErrNilRequest
	}

	if x.Queue == nil {
		return ErrNilQueue
	}

	if x.Queue.Name == "" {
		return ErrEmptyQueueName
	}

	return nil
}

func (x *FlushReq) CheckValid() error {
	if x == nil {
		return ErrNilRequest
	}

	if x.Name == "" {
		return ErrNilQueue
	}

	return nil
}

func (x *CountJobReq) CheckValid() error {
	if x == nil {
		return ErrNilRequest
	}

	if x.Name == "" {
		return ErrNilQueue
	}

	return nil
}
