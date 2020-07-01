// Code generated by skv2. DO NOT EDIT.

// The Input Reconciler calls a simple func() error whenever a
// storage event is received for any of:
// * MeshServices
// * MeshWorkloads
// * Meshes
// * TrafficPolicies
// * AccessPolicies
// * VirtualMeshes
// for a given cluster or set of clusters.
//
// Input Reconcilers can be be constructed from either a single Manager (watch events in a single cluster)
// or a ClusterWatcher (watch events in multiple clusters).
package input

import (
	"context"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/reconcile"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	discovery_smh_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	discovery_smh_solo_io_v1alpha1_controllers "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/controller"

	networking_smh_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	networking_smh_solo_io_v1alpha1_controllers "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/controller"
)

// the multiClusterReconciler reconciles events for input resources across clusters
type multiClusterReconciler interface {
	discovery_smh_solo_io_v1alpha1_controllers.MulticlusterMeshServiceReconciler
	discovery_smh_solo_io_v1alpha1_controllers.MulticlusterMeshWorkloadReconciler
	discovery_smh_solo_io_v1alpha1_controllers.MulticlusterMeshReconciler

	networking_smh_solo_io_v1alpha1_controllers.MulticlusterTrafficPolicyReconciler
	networking_smh_solo_io_v1alpha1_controllers.MulticlusterAccessPolicyReconciler
	networking_smh_solo_io_v1alpha1_controllers.MulticlusterVirtualMeshReconciler
}

var _ multiClusterReconciler = &multiClusterReconcilerImpl{}

type multiClusterReconcilerImpl struct {
	ctx           context.Context
	reconcileFunc func(cluster string) error
}

// register the reconcile func with the cluster watcher
func RegisterMultiClusterReconciler(
	ctx context.Context,
	clusters multicluster.ClusterWatcher,
	reconcileFunc func(cluster string) error,
) {
	r := &multiClusterReconcilerImpl{
		ctx:           ctx,
		reconcileFunc: reconcileFunc,
	}

	// initialize reconcile loops

	discovery_smh_solo_io_v1alpha1_controllers.NewMulticlusterMeshServiceReconcileLoop("MeshService", clusters).AddMulticlusterMeshServiceReconciler(ctx, r)
	discovery_smh_solo_io_v1alpha1_controllers.NewMulticlusterMeshWorkloadReconcileLoop("MeshWorkload", clusters).AddMulticlusterMeshWorkloadReconciler(ctx, r)
	discovery_smh_solo_io_v1alpha1_controllers.NewMulticlusterMeshReconcileLoop("Mesh", clusters).AddMulticlusterMeshReconciler(ctx, r)

	networking_smh_solo_io_v1alpha1_controllers.NewMulticlusterTrafficPolicyReconcileLoop("TrafficPolicy", clusters).AddMulticlusterTrafficPolicyReconciler(ctx, r)
	networking_smh_solo_io_v1alpha1_controllers.NewMulticlusterAccessPolicyReconcileLoop("AccessPolicy", clusters).AddMulticlusterAccessPolicyReconciler(ctx, r)
	networking_smh_solo_io_v1alpha1_controllers.NewMulticlusterVirtualMeshReconcileLoop("VirtualMesh", clusters).AddMulticlusterVirtualMeshReconciler(ctx, r)
}

func (r *multiClusterReconcilerImpl) ReconcileMeshService(clusterName string, obj *discovery_smh_solo_io_v1alpha1.MeshService) (reconcile.Result, error) {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", clusterName, "obj", obj)
	return reconcile.Result{}, r.reconcileFunc(clusterName)
}

func (r *multiClusterReconcilerImpl) ReconcileMeshServiceDeletion(clusterName string, req reconcile.Request) error {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", clusterName, "obj", obj)
	return r.reconcileFunc(clusterName)
}

func (r *multiClusterReconcilerImpl) ReconcileMeshWorkload(clusterName string, obj *discovery_smh_solo_io_v1alpha1.MeshWorkload) (reconcile.Result, error) {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", clusterName, "obj", obj)
	return reconcile.Result{}, r.reconcileFunc(clusterName)
}

func (r *multiClusterReconcilerImpl) ReconcileMeshWorkloadDeletion(clusterName string, req reconcile.Request) error {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", clusterName, "obj", obj)
	return r.reconcileFunc(clusterName)
}

func (r *multiClusterReconcilerImpl) ReconcileMesh(clusterName string, obj *discovery_smh_solo_io_v1alpha1.Mesh) (reconcile.Result, error) {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", clusterName, "obj", obj)
	return reconcile.Result{}, r.reconcileFunc(clusterName)
}

func (r *multiClusterReconcilerImpl) ReconcileMeshDeletion(clusterName string, req reconcile.Request) error {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", clusterName, "obj", obj)
	return r.reconcileFunc(clusterName)
}

func (r *multiClusterReconcilerImpl) ReconcileTrafficPolicy(clusterName string, obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) (reconcile.Result, error) {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", clusterName, "obj", obj)
	return reconcile.Result{}, r.reconcileFunc(clusterName)
}

func (r *multiClusterReconcilerImpl) ReconcileTrafficPolicyDeletion(clusterName string, req reconcile.Request) error {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", clusterName, "obj", obj)
	return r.reconcileFunc(clusterName)
}

func (r *multiClusterReconcilerImpl) ReconcileAccessPolicy(clusterName string, obj *networking_smh_solo_io_v1alpha1.AccessPolicy) (reconcile.Result, error) {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", clusterName, "obj", obj)
	return reconcile.Result{}, r.reconcileFunc(clusterName)
}

func (r *multiClusterReconcilerImpl) ReconcileAccessPolicyDeletion(clusterName string, req reconcile.Request) error {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", clusterName, "obj", obj)
	return r.reconcileFunc(clusterName)
}

func (r *multiClusterReconcilerImpl) ReconcileVirtualMesh(clusterName string, obj *networking_smh_solo_io_v1alpha1.VirtualMesh) (reconcile.Result, error) {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", clusterName, "obj", obj)
	return reconcile.Result{}, r.reconcileFunc(clusterName)
}

func (r *multiClusterReconcilerImpl) ReconcileVirtualMeshDeletion(clusterName string, req reconcile.Request) error {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", clusterName, "obj", obj)
	return r.reconcileFunc(clusterName)
}

// the singleClusterReconciler reconciles events for input resources across clusters
type singleClusterReconciler interface {
	discovery_smh_solo_io_v1alpha1_controllers.MeshServiceReconciler
	discovery_smh_solo_io_v1alpha1_controllers.MeshWorkloadReconciler
	discovery_smh_solo_io_v1alpha1_controllers.MeshReconciler

	networking_smh_solo_io_v1alpha1_controllers.TrafficPolicyReconciler
	networking_smh_solo_io_v1alpha1_controllers.AccessPolicyReconciler
	networking_smh_solo_io_v1alpha1_controllers.VirtualMeshReconciler
}

var _ singleClusterReconciler = &singleClusterReconcilerImpl{}

type singleClusterReconcilerImpl struct {
	ctx           context.Context
	reconcileFunc func(metav1.Object) error
}

// register the reconcile func with the manager
func RegisterSingleClusterReconciler(
	ctx context.Context,
	mgr manager.Manager,
	reconcileFunc func() error,
) error {
	r := &singleClusterReconcilerImpl{
		ctx:           ctx,
		reconcileFunc: reconcileFunc,
	}

	// initialize reconcile loops

	if err := discovery_smh_solo_io_v1alpha1_controllers.NewMeshServiceReconcileLoop("MeshService", mgr, reconcile.Options{}).RunMeshServiceReconciler(ctx, r); err != nil {
		return err
	}
	if err := discovery_smh_solo_io_v1alpha1_controllers.NewMeshWorkloadReconcileLoop("MeshWorkload", mgr, reconcile.Options{}).RunMeshWorkloadReconciler(ctx, r); err != nil {
		return err
	}
	if err := discovery_smh_solo_io_v1alpha1_controllers.NewMeshReconcileLoop("Mesh", mgr, reconcile.Options{}).RunMeshReconciler(ctx, r); err != nil {
		return err
	}

	if err := networking_smh_solo_io_v1alpha1_controllers.NewTrafficPolicyReconcileLoop("TrafficPolicy", mgr, reconcile.Options{}).RunTrafficPolicyReconciler(ctx, r); err != nil {
		return err
	}
	if err := networking_smh_solo_io_v1alpha1_controllers.NewAccessPolicyReconcileLoop("AccessPolicy", mgr, reconcile.Options{}).RunAccessPolicyReconciler(ctx, r); err != nil {
		return err
	}
	if err := networking_smh_solo_io_v1alpha1_controllers.NewVirtualMeshReconcileLoop("VirtualMesh", mgr, reconcile.Options{}).RunVirtualMeshReconciler(ctx, r); err != nil {
		return err
	}

	return nil
}

func (r *singleClusterReconcilerImpl) ReconcileMeshService(obj *discovery_smh_solo_io_v1alpha1.MeshService) (reconcile.Result, error) {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", obj)
	return reconcile.Result{}, r.reconcileFunc()
}

func (r *singleClusterReconcilerImpl) ReconcileMeshServiceDeletion(obj reconcile.Request) error {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", obj)
	return r.reconcileFunc()
}

func (r *singleClusterReconcilerImpl) ReconcileMeshWorkload(obj *discovery_smh_solo_io_v1alpha1.MeshWorkload) (reconcile.Result, error) {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", obj)
	return reconcile.Result{}, r.reconcileFunc()
}

func (r *singleClusterReconcilerImpl) ReconcileMeshWorkloadDeletion(obj reconcile.Request) error {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", obj)
	return r.reconcileFunc()
}

func (r *singleClusterReconcilerImpl) ReconcileMesh(obj *discovery_smh_solo_io_v1alpha1.Mesh) (reconcile.Result, error) {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", obj)
	return reconcile.Result{}, r.reconcileFunc()
}

func (r *singleClusterReconcilerImpl) ReconcileMeshDeletion(obj reconcile.Request) error {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", obj)
	return r.reconcileFunc()
}

func (r *singleClusterReconcilerImpl) ReconcileTrafficPolicy(obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) (reconcile.Result, error) {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", obj)
	return reconcile.Result{}, r.reconcileFunc()
}

func (r *singleClusterReconcilerImpl) ReconcileTrafficPolicyDeletion(obj reconcile.Request) error {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", obj)
	return r.reconcileFunc()
}

func (r *singleClusterReconcilerImpl) ReconcileAccessPolicy(obj *networking_smh_solo_io_v1alpha1.AccessPolicy) (reconcile.Result, error) {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", obj)
	return reconcile.Result{}, r.reconcileFunc()
}

func (r *singleClusterReconcilerImpl) ReconcileAccessPolicyDeletion(obj reconcile.Request) error {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", obj)
	return r.reconcileFunc()
}

func (r *singleClusterReconcilerImpl) ReconcileVirtualMesh(obj *networking_smh_solo_io_v1alpha1.VirtualMesh) (reconcile.Result, error) {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", obj)
	return reconcile.Result{}, r.reconcileFunc()
}

func (r *singleClusterReconcilerImpl) ReconcileVirtualMeshDeletion(obj reconcile.Request) error {
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", obj)
	return r.reconcileFunc()
}
