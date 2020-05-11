/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apis

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestConditionStatus(t *testing.T) {
	tests := []struct {
		name            string
		condition       *Condition
		expectedTrue    bool
		expectedFalse   bool
		expectedUnknown bool
	}{
		{
			name:            "true",
			condition:       &Condition{Status: corev1.ConditionTrue},
			expectedTrue:    true,
			expectedFalse:   false,
			expectedUnknown: false,
		},
		{
			name:            "false",
			condition:       &Condition{Status: corev1.ConditionFalse},
			expectedTrue:    false,
			expectedFalse:   true,
			expectedUnknown: false,
		},
		{
			name:            "unknown",
			condition:       &Condition{Status: corev1.ConditionUnknown},
			expectedTrue:    false,
			expectedFalse:   false,
			expectedUnknown: true,
		},
		{
			name:            "unset",
			condition:       &Condition{},
			expectedTrue:    false,
			expectedFalse:   false,
			expectedUnknown: false,
		},
		{
			name:            "nil",
			condition:       nil,
			expectedTrue:    false,
			expectedFalse:   false,
			expectedUnknown: true,
		},
	}
	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if expected, actual := c.expectedTrue, c.condition.IsTrue(); expected != actual {
				t.Errorf("%s: IsTrue() actually = %v, expected %v", c.name, actual, expected)
			}
			if expected, actual := c.expectedFalse, c.condition.IsFalse(); expected != actual {
				t.Errorf("%s: IsFalse() actually = %v, expected %v", c.name, actual, expected)
			}
			if expected, actual := c.expectedUnknown, c.condition.IsUnknown(); expected != actual {
				t.Errorf("%s: IsUnknown() actually = %v, expected %v", c.name, actual, expected)
			}
		})
	}
}
