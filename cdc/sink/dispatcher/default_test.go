// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package dispatcher

import (
	"github.com/pingcap/check"
	"github.com/pingcap/ticdc/cdc/model"
)

type DefaultDispatcherSuite struct{}

var _ = check.Suite(&DefaultDispatcherSuite{})

func (s DefaultDispatcherSuite) TestDefaultDispatcher(c *check.C) {
	testCases := []struct {
		row             *model.RowChangedEvent
		exceptPartition int32
	}{
		{row: &model.RowChangedEvent{
			Table: &model.TableName{
				Schema: "test",
				Table:  "t1",
			},
			IndieMarkCol: "id",
			Columns: map[string]*model.Column{
				"id": {
					Value: 1,
				},
			},
		}, exceptPartition: 7},
		{row: &model.RowChangedEvent{
			Table: &model.TableName{
				Schema: "test",
				Table:  "t1",
			},
			IndieMarkCol: "id",
			Columns: map[string]*model.Column{
				"id": {
					Value: 2,
				},
			},
		}, exceptPartition: 13},
		{row: &model.RowChangedEvent{
			Table: &model.TableName{
				Schema: "test",
				Table:  "t1",
			},
			IndieMarkCol: "id",
			Columns: map[string]*model.Column{
				"id": {
					Value: 3,
				},
			},
		}, exceptPartition: 11},
		{row: &model.RowChangedEvent{
			Table: &model.TableName{
				Schema: "test",
				Table:  "t2",
			},
			IndieMarkCol: "id",
			Columns: map[string]*model.Column{
				"id": {
					Value: 1,
				},
			},
		}, exceptPartition: 7},
		{row: &model.RowChangedEvent{
			Table: &model.TableName{
				Schema: "test",
				Table:  "t2",
			},
			IndieMarkCol: "id",
			Columns: map[string]*model.Column{
				"id": {
					Value: 2,
				},
			},
		}, exceptPartition: 13},
		{row: &model.RowChangedEvent{
			Table: &model.TableName{
				Schema: "test",
				Table:  "t2",
			},
			IndieMarkCol: "id",
			Columns: map[string]*model.Column{
				"id": {
					Value: 3,
				},
			},
		}, exceptPartition: 11},
		{row: &model.RowChangedEvent{
			Table: &model.TableName{
				Schema: "test",
				Table:  "t3",
			},
			Columns: map[string]*model.Column{
				"id": {
					Value: 1,
				},
			},
		}, exceptPartition: 3},
		{row: &model.RowChangedEvent{
			Table: &model.TableName{
				Schema: "test",
				Table:  "t3",
			},
			Columns: map[string]*model.Column{
				"id": {
					Value: 2,
				},
			},
		}, exceptPartition: 3},
		{row: &model.RowChangedEvent{
			Table: &model.TableName{
				Schema: "test",
				Table:  "t3",
			},
			Columns: map[string]*model.Column{
				"id": {
					Value: 3,
				},
			},
		}, exceptPartition: 3},
	}
	p := &defaultDispatcher{partitionNum: 16}
	for _, tc := range testCases {
		c.Assert(p.Dispatch(tc.row), check.Equals, tc.exceptPartition)
	}
}
