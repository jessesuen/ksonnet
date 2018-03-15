// Copyright 2017 The ksonnet authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package kubecfg

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrintComponents(t *testing.T) {
	cases := []struct {
		name       string
		components []string
		expected   string
	}{
		{
			name:       "print",
			components: []string{"a", "b"},
			expected: `COMPONENT
=========
a
b
`,
		},
		{
			name:       "Check that components are displayed in alphabetical order",
			components: []string{"b", "a"},
			expected: `COMPONENT
=========
a
b
`,
		},
		{
			name:       "Check empty components scenario",
			components: []string{},
			expected: `COMPONENT
=========
`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := printComponents(&buf, tc.components)
			require.NoError(t, err)
			require.EqualValues(t, tc.expected, buf.String())
		})
	}
}
