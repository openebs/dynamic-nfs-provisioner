/*
Copyright 2021 The OpenEBS Authors.

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

package hook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPvc_hook_action(t *testing.T) {
	tests := []struct {
		name        string
		hook        *PVCHook
		obj         interface{}
		expectedObj interface{}
		actionType  HookActionType
	}{
		{
			name:        "when PVC hook is nil, object should not be modified",
			hook:        nil,
			obj:         generateFakePvcObj("ns1", "name1", map[string]string{"test.io/key": "val"}, []string{"test.io/finalizer"}),
			expectedObj: generateFakePvcObj("ns1", "name1", map[string]string{"test.io/key": "val"}, []string{"test.io/finalizer"}),
			actionType:  HookActionAdd,
		},
		{
			name:        "when PVC hook is configured to add metadata, object should be modified",
			hook:        buildPVCHook(map[string]string{"test.io/key": "val"}, []string{"test.io/finalizer"}),
			obj:         generateFakePvcObj("ns2", "name2", nil, nil),
			expectedObj: generateFakePvcObj("ns2", "name2", map[string]string{"test.io/key": "val"}, []string{"test.io/finalizer"}),
			actionType:  HookActionAdd,
		},
		{
			name:        "when PVC hook is configured to remove metadata, object should be modified",
			hook:        buildPVCHook(map[string]string{"test.io/key": "val"}, []string{"test.io/finalizer"}),
			obj:         generateFakePvcObj("ns3", "name3", map[string]string{"test.io/key": "val"}, []string{"test.io/finalizer"}),
			expectedObj: generateFakePvcObj("ns3", "name3", map[string]string{}, []string{}),
			actionType:  HookActionRemove,
		},
		{
			name:        "when PVC hook is configured to remove non-existing metadata, object should not be modified",
			hook:        buildPVCHook(map[string]string{"test.com/key": "val"}, []string{"test.com/finalizer"}),
			obj:         generateFakePvcObj("ns4", "name4", map[string]string{"test.io/key": "val"}, []string{"test.io/finalizer"}),
			expectedObj: generateFakePvcObj("ns4", "name4", map[string]string{"test.io/key": "val"}, []string{"test.io/finalizer"}),
			actionType:  HookActionRemove,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.NotNil(t, test.obj, "object should not be nil")
			err := pvc_hook_action(test.hook, test.actionType, test.obj)
			assert.Nil(t, err, "pvc_hook_action returned error")
			assert.Equal(t, test.expectedObj, test.obj, "object should match")
		})
	}
}