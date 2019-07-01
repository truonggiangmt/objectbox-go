/*
 * Copyright 2019 ObjectBox Ltd. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import "time"

//go:generate objectbox-gogen

// Tests all available GO & ObjectBox types
// TODO rename; e.g. TestEntity
type Entity struct {
	// base types
	Id           uint64
	Int          int
	Int8         int8
	Int16        int16
	Int32        int32
	Int64        int64
	Uint         uint
	Uint8        uint8
	Uint16       uint16
	Uint32       uint32
	Uint64       uint64
	Bool         bool
	String       string
	StringVector []string
	Byte         byte
	ByteVector   []byte
	Rune         rune
	Float32      float32
	Float64      float64

	// converters
	Date       time.Time  `date type:"int64" converter:"timeInt64"`
	Complex128 complex128 `type:"[]byte" converter:"complex128Bytes"`

	// one-to-many relations
	Related     TestEntityRelated  `link`
	RelatedPtr  *TestEntityRelated `link`
	RelatedPtr2 *TestEntityRelated `link`

	// many-to-many relations
	RelatedSlice    []EntityByValue
	RelatedPtrSlice []*TestEntityRelated

	IntPtr          *int
	Int8Ptr         *int8
	Int16Ptr        *int16
	Int32Ptr        *int32
	Int64Ptr        *int64
	UintPtr         *uint
	Uint8Ptr        *uint8
	Uint16Ptr       *uint16
	Uint32Ptr       *uint32
	Uint64Ptr       *uint64
	BoolPtr         *bool
	StringPtr       *string
	StringVectorPtr *[]string
	BytePtr         *byte
	ByteVectorPtr   *[]byte
	RunePtr         *rune
	Float32Ptr      *float32
	Float64Ptr      *float64
}

type TestStringIdEntity struct {
	Id string `id`
}

type TestEntityInline struct {
	BaseWithDate   `inline`
	*BaseWithValue `inline`

	Id uint64
}

type TestEntityRelated struct {
	Id   uint64
	Name string

	// have another level of relations
	Next      *EntityByValue `link`
	NextSlice []EntityByValue
}
