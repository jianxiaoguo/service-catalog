/*
Copyright 2017 The Kubernetes Authors.

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

package controller

import (
	"encoding/json"

	osb "sigs.k8s.io/go-open-service-broker-client/v2"
	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
)

const (
	originatingIdentityPlatform = "kubernetes"
)

func buildOriginatingIdentity(userInfo *v1beta1.UserInfo) (*osb.OriginatingIdentity, error) {
	if userInfo == nil {
		return nil, nil
	}
	oiValue, err := json.Marshal(userInfo)
	if err != nil {
		return nil, err
	}
	oi := &osb.OriginatingIdentity{
		Platform: originatingIdentityPlatform,
		Value:    string(oiValue),
	}
	return oi, nil
}
