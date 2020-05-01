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

package reconcilers

import (
	"context"
)

const stashNonce string = "controller-stash-nonce"

type stashMap map[StashKey]interface{}

func WithStash(ctx context.Context) context.Context {
	return context.WithValue(ctx, stashNonce, stashMap{})
}

type StashKey string

func StashValue(ctx context.Context, key StashKey, value interface{}) {
	stash := ctx.Value(stashNonce).(stashMap)
	stash[key] = value
}

func RetrieveValue(ctx context.Context, key StashKey) interface{} {
	stash := ctx.Value(stashNonce).(stashMap)
	return stash[key]
}
