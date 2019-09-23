// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package block

import (
	"fmt"

	"github.com/m3db/m3/src/query/models"
)

// Metadata is metadata for a block, describing size and common tags accross
// constituent series.
type Metadata struct {
	// Bounds represents the time bounds for all series in the block.
	Bounds models.Bounds
	// Tags contains any tags common across all series in the block.
	Tags models.Tags
	// ResultMetadata contains metadata from any database access operations during
	// fetching block details.
	ResultMetadata ResultMetadata
}

// Equals returns a boolean reporting whether the compared metadata has equal
// fields.
func (m Metadata) Equals(other Metadata) bool {
	return m.Tags.Equals(other.Tags) && m.Bounds.Equals(other.Bounds)
}

// String returns a string representation of metadata.
func (m Metadata) String() string {
	return fmt.Sprintf("Bounds: %v, Tags: %v", m.Bounds, m.Tags)
}

// ResultMetadata describes metadata common to each type of query results,
// indicating any additional information about the result.
type ResultMetadata struct {
	// LocalOnly indicates that this query was executed only on the local store.
	LocalOnly bool
	// Exhaustive indicates whether the underlying data set presents a full
	// collection of retrieved data.
	Exhaustive bool
	// Warnings is a list of warnings that indicate potetitally partial or
	// incomplete results.
	Warnings []Warning
}

func NewResultMetadata() ResultMetadata {
	return ResultMetadata{
		LocalOnly:  true,
		Exhaustive: true,
	}
}

func (m ResultMetadata) CombineMetadata(other ResultMetadata) ResultMetadata {
	combinedWarnings := make([]Warning, 0, len(m.Warnings)+len(other.Warnings))
	for _, w := range m.Warnings {
		combinedWarnings = append(combinedWarnings, w)
	}

	meta := ResultMetadata{
		LocalOnly:  m.LocalOnly && other.LocalOnly,
		Exhaustive: m.Exhaustive && other.Exhaustive,
		Warnings:   combinedWarnings,
	}

	for _, w := range other.Warnings {
		meta.AddWarning(w.Name, w.Message)
	}

	return meta
}

func (m ResultMetadata) IsDefault() bool {
	return m.Exhaustive && m.LocalOnly && len(m.Warnings) == 0
}

// AddWarning adds a warning to the result metadata.
// NB: warnings are expected to be small in general, so it's better to iterate
// over the array rather than introduce a map.
func (m *ResultMetadata) AddWarning(name string, message string) {
	for _, warning := range m.Warnings {
		// Dedupe warnings.
		if warning.Name == name && warning.Message == message {
			return
		}
	}

	m.Warnings = append(m.Warnings, Warning{
		Name:    name,
		Message: message,
	})
}

// Warning is a message that indicates potential partial or incomplete results.
type Warning struct {
	// Name is the name of the store originating the warning.
	Name string
	// Message is the content of the warning message.
	Message string
}

// Header formats the warning into a format to send in a response header.
func (w Warning) Header() string {
	return fmt.Sprintf("%s_%s", w.Name, w.Message)
}