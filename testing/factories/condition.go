/*
Copyright 2020 the original author or authors.

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

package factories

import (
	"fmt"

	"github.com/projectriff/reconciler-runtime/apis"
	corev1 "k8s.io/api/core/v1"
)

type ConditionFactory interface {
	Create() apis.Condition

	Type(t apis.ConditionType) ConditionFactory
	Unknown() ConditionFactory
	True() ConditionFactory
	False() ConditionFactory
	Reason(reason, message string) ConditionFactory
	Info() ConditionFactory
	Warning() ConditionFactory
	Error() ConditionFactory
}

type conditionImpl struct {
	target *apis.Condition
}

func Condition(seed ...apis.Condition) ConditionFactory {
	var target *apis.Condition
	switch len(seed) {
	case 0:
		target = &apis.Condition{}
	case 1:
		target = &seed[0]
	default:
		panic(fmt.Errorf("expected exactly zero or one seed, got %v", seed))
	}
	return &conditionImpl{
		target: target,
	}
}

func (f *conditionImpl) deepCopy() *conditionImpl {
	return &conditionImpl{
		target: f.target.DeepCopy(),
	}
}

func (f *conditionImpl) Create() apis.Condition {
	t := f.deepCopy().target
	return *t
}

func (f *conditionImpl) mutation(m func(*apis.Condition)) ConditionFactory {
	f = f.deepCopy()
	m(f.target)
	return f
}

func (f *conditionImpl) Type(t apis.ConditionType) ConditionFactory {
	return f.mutation(func(c *apis.Condition) {
		c.Type = t
	})
}

func (f *conditionImpl) Unknown() ConditionFactory {
	return f.mutation(func(c *apis.Condition) {
		c.Status = corev1.ConditionUnknown
	})
}

func (f *conditionImpl) True() ConditionFactory {
	return f.mutation(func(c *apis.Condition) {
		c.Status = corev1.ConditionTrue
		c.Reason = ""
		c.Message = ""
	})
}

func (f *conditionImpl) False() ConditionFactory {
	return f.mutation(func(c *apis.Condition) {
		c.Status = corev1.ConditionFalse
	})
}

func (f *conditionImpl) Reason(reason, message string) ConditionFactory {
	return f.mutation(func(c *apis.Condition) {
		c.Reason = reason
		c.Message = message
	})
}

func (f *conditionImpl) Info() ConditionFactory {
	return f.mutation(func(c *apis.Condition) {
		c.Severity = apis.ConditionSeverityInfo
	})
}

func (f *conditionImpl) Warning() ConditionFactory {
	return f.mutation(func(c *apis.Condition) {
		c.Severity = apis.ConditionSeverityWarning
	})
}

func (f *conditionImpl) Error() ConditionFactory {
	return f.mutation(func(c *apis.Condition) {
		c.Severity = apis.ConditionSeverityError
	})
}
