// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package watchers

import (
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"

	"github.com/cilium/cilium/pkg/egressgateway"
	"github.com/cilium/cilium/pkg/k8s"
	cilium_v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	cilium_v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	"github.com/cilium/cilium/pkg/k8s/client"
	"github.com/cilium/cilium/pkg/k8s/informer"
	"github.com/cilium/cilium/pkg/k8s/watchers/resources"
	"github.com/cilium/cilium/pkg/logging/logfields"
)

func (k *K8sWatcher) ciliumEgressGatewayPolicyInit(ciliumNPClient client.Clientset) {
	apiGroup := k8sAPIGroupCiliumEgressGatewayPolicyV2
	_, egpController := informer.NewInformer(
		cache.NewListWatchFromClient(ciliumNPClient.CiliumV2().RESTClient(),
			"ciliumegressgatewaypolicies", v1.NamespaceAll, fields.Everything()),
		&cilium_v2.CiliumEgressGatewayPolicy{},
		0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				var valid, equal bool
				defer func() { k.K8sEventReceived(apiGroup, metricCEGP, resources.MetricCreate, valid, equal) }()
				if cegp := k8s.ObjToCEGP(obj); cegp != nil {
					valid = true
					err := k.addCiliumEgressGatewayPolicy(cegp)
					k.K8sEventProcessed(metricCEGP, resources.MetricCreate, err == nil)
				}
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				var valid, equal bool
				defer func() { k.K8sEventReceived(apiGroup, metricCEGP, resources.MetricUpdate, valid, equal) }()

				newCegp := k8s.ObjToCEGP(newObj)
				if newCegp == nil {
					return
				}
				valid = true
				addErr := k.addCiliumEgressGatewayPolicy(newCegp)
				k.K8sEventProcessed(metricCEGP, resources.MetricUpdate, addErr == nil)
			},
			DeleteFunc: func(obj interface{}) {
				var valid, equal bool
				defer func() { k.K8sEventReceived(apiGroup, metricCEGP, resources.MetricDelete, valid, equal) }()
				cegp := k8s.ObjToCEGP(obj)
				if cegp == nil {
					return
				}
				valid = true
				k.deleteCiliumEgressGatewayPolicy(cegp)
				k.K8sEventProcessed(metricCEGP, resources.MetricDelete, true)
			},
		},
		k8s.ConvertToCiliumEgressGatewayPolicy,
	)

	k.blockWaitGroupToSyncResources(
		k.stop,
		nil,
		egpController.HasSynced,
		k8sAPIGroupCiliumEgressGatewayPolicyV2,
	)

	go egpController.Run(k.stop)
	k.k8sAPIGroups.AddAPI(k8sAPIGroupCiliumEgressGatewayPolicyV2)
}

func (k *K8sWatcher) addCiliumEgressGatewayPolicy(cegp *cilium_v2.CiliumEgressGatewayPolicy) error {
	scopedLog := log.WithFields(logrus.Fields{
		logfields.CiliumEgressGatewayPolicyName: cegp.ObjectMeta.Name,
		logfields.K8sUID:                        cegp.ObjectMeta.UID,
		logfields.K8sAPIVersion:                 cegp.TypeMeta.APIVersion,
	})

	ep, err := egressgateway.ParseCEGP(cegp)
	if err != nil {
		scopedLog.WithError(err).Warn("Failed to add CiliumEgressGatewayPolicy: malformed policy config.")
		return err
	}
	k.egressGatewayManager.OnAddEgressPolicy(*ep)

	return err
}

func (k *K8sWatcher) deleteCiliumEgressGatewayPolicy(cegp *cilium_v2.CiliumEgressGatewayPolicy) {
	epID := egressgateway.ParseCEGPConfigID(cegp)
	k.egressGatewayManager.OnDeleteEgressPolicy(epID)
}

func (k *K8sWatcher) ciliumEgressNATPolicyInit(ciliumNPClient client.Clientset) {
	apiGroup := k8sAPIGroupCiliumEgressNATPolicyV2
	_, egpController := informer.NewInformer(
		cache.NewListWatchFromClient(ciliumNPClient.CiliumV2alpha1().RESTClient(),
			"ciliumegressnatpolicies", v1.NamespaceAll, fields.Everything()),
		&cilium_v2alpha1.CiliumEgressNATPolicy{},
		0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				var valid, equal bool
				defer func() { k.K8sEventReceived(apiGroup, metricCENP, resources.MetricCreate, valid, equal) }()
				if cenp := k8s.ObjToCENP(obj); cenp != nil {
					valid = true
					err := k.addCiliumEgressNATPolicy(cenp)
					k.K8sEventProcessed(metricCENP, resources.MetricCreate, err == nil)
				}
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				var valid, equal bool
				defer func() { k.K8sEventReceived(apiGroup, metricCENP, resources.MetricUpdate, valid, equal) }()

				newCenp := k8s.ObjToCENP(newObj)
				if newCenp == nil {
					return
				}
				valid = true
				addErr := k.addCiliumEgressNATPolicy(newCenp)
				k.K8sEventProcessed(metricCENP, resources.MetricUpdate, addErr == nil)
			},
			DeleteFunc: func(obj interface{}) {
				var valid, equal bool
				defer func() { k.K8sEventReceived(apiGroup, metricCENP, resources.MetricDelete, valid, equal) }()
				cenp := k8s.ObjToCENP(obj)
				if cenp == nil {
					return
				}
				valid = true
				k.deleteCiliumEgressNATPolicy(cenp)
				k.K8sEventProcessed(metricCENP, resources.MetricDelete, true)
			},
		},
		k8s.ConvertToCiliumEgressNATPolicy,
	)

	k.blockWaitGroupToSyncResources(
		k.stop,
		nil,
		egpController.HasSynced,
		k8sAPIGroupCiliumEgressNATPolicyV2,
	)

	go egpController.Run(k.stop)
	k.k8sAPIGroups.AddAPI(k8sAPIGroupCiliumEgressNATPolicyV2)
}

func (k *K8sWatcher) addCiliumEgressNATPolicy(cenp *cilium_v2alpha1.CiliumEgressNATPolicy) error {
	scopedLog := log.WithFields(logrus.Fields{
		logfields.CiliumEgressNATPolicyName: cenp.ObjectMeta.Name,
		logfields.K8sUID:                    cenp.ObjectMeta.UID,
		logfields.K8sAPIVersion:             cenp.TypeMeta.APIVersion,
	})

	ep, err := egressgateway.ParseCENP(cenp)
	if err != nil {
		scopedLog.WithError(err).Warn("Failed to add CiliumEgressNATPolicy: malformed policy config.")
		return err
	}
	k.egressGatewayManager.OnAddEgressPolicy(*ep)

	return err
}

func (k *K8sWatcher) deleteCiliumEgressNATPolicy(cenp *cilium_v2alpha1.CiliumEgressNATPolicy) {
	epID := egressgateway.ParseCENPConfigID(cenp)
	k.egressGatewayManager.OnDeleteEgressPolicy(epID)
}
