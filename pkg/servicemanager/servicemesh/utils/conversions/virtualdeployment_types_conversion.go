/*
  Copyright (c) 2022, Oracle and/or its affiliates. All rights reserved.
  Licensed under the Universal Permissive License v 1.0 as shown at http://oss.oracle.com/licenses/upl.
*/

package conversions

import (
	sdk "github.com/oracle/oci-go-sdk/v65/servicemesh"

	api "github.com/oracle/oci-service-operator/api/v1beta1"
	"github.com/oracle/oci-service-operator/apis/servicemesh.oci/v1beta1"
)

// ConvertCrdVirtualDeploymentToSdkVirtualDeployment converts a CRD object to an object that can be sent to the API
func ConvertCrdVirtualDeploymentToSdkVirtualDeployment(crdObj *v1beta1.VirtualDeployment, sdkObj *sdk.VirtualDeployment, vsId *api.OCID) {
	sdkObj.Id = (*string)(&crdObj.Status.VirtualDeploymentId)
	sdkObj.CompartmentId = (*string)(&crdObj.Spec.CompartmentId)
	sdkObj.Name = GetSpecName(crdObj.Spec.Name, &crdObj.ObjectMeta)
	sdkObj.Description = (*string)(crdObj.Spec.Description)
	sdkObj.VirtualServiceId = (*string)(vsId)
	sdkObj.AccessLogging = convertCrdAccessLoggingToSdkAccessLogging(crdObj.Spec.AccessLogging)

	listener := ConvertCrdVirtualDeploymentListenerToSdkVirtualDeploymentListener(crdObj.Spec.Listener)
	sdkObj.Listeners = listener
	// TODO: check validation in webhook
	if crdObj.Spec.ServiceDiscovery.Type == v1beta1.ServiceDiscoveryTypeDns {
		sdkObj.ServiceDiscovery = sdk.DnsServiceDiscoveryConfiguration{
			Hostname: &crdObj.Spec.ServiceDiscovery.Hostname,
		}
	}

	// TODO: AccessLogging.  Waiting for PR review.
	if crdObj.Spec.FreeFormTags != nil {
		sdkObj.FreeformTags = crdObj.Spec.FreeFormTags
	}
	if crdObj.Spec.DefinedTags != nil {
		sdkObj.DefinedTags = map[string]map[string]interface{}{}
		ConvertCrdDefinedTagsToSdkDefinedTags(&crdObj.Spec.DefinedTags, &sdkObj.DefinedTags)
	}
}

// ConvertCrdVirtualDeploymentListenerToSdkVirtualDeploymentListener converts a listener from a CRD object to a listener for an SDK object
func ConvertCrdVirtualDeploymentListenerToSdkVirtualDeploymentListener(crdListener []v1beta1.Listener) (sdkListeners []sdk.VirtualDeploymentListener) {
	sdkListeners = make([]sdk.VirtualDeploymentListener, 0, len(crdListener))

	for _, l := range crdListener {
		i := int(l.Port)
		sdkListeners = append(sdkListeners, sdk.VirtualDeploymentListener{
			Protocol: sdk.VirtualDeploymentListenerProtocolEnum(l.Protocol),
			Port:     &i,
		})
	}

	return sdkListeners
}
