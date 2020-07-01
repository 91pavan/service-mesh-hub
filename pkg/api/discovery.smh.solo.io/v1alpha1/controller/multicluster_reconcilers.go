// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	discovery_smh_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the MeshService Resource across clusters.
// implemented by the user
type MulticlusterMeshServiceReconciler interface {
	ReconcileMeshService(clusterName string, obj *discovery_smh_solo_io_v1alpha1.MeshService) (reconcile.Result, error)
}

// Reconcile deletion events for the MeshService Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterMeshServiceDeletionReconciler interface {
	ReconcileMeshServiceDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterMeshServiceReconcilerFuncs struct {
	OnReconcileMeshService         func(clusterName string, obj *discovery_smh_solo_io_v1alpha1.MeshService) (reconcile.Result, error)
	OnReconcileMeshServiceDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterMeshServiceReconcilerFuncs) ReconcileMeshService(clusterName string, obj *discovery_smh_solo_io_v1alpha1.MeshService) (reconcile.Result, error) {
	if f.OnReconcileMeshService == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileMeshService(clusterName, obj)
}

func (f *MulticlusterMeshServiceReconcilerFuncs) ReconcileMeshServiceDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileMeshServiceDeletion == nil {
		return nil
	}
	return f.OnReconcileMeshServiceDeletion(clusterName, req)
}

type MulticlusterMeshServiceReconcileLoop interface {
	// AddMulticlusterMeshServiceReconciler adds a MulticlusterMeshServiceReconciler to the MulticlusterMeshServiceReconcileLoop.
	AddMulticlusterMeshServiceReconciler(ctx context.Context, rec MulticlusterMeshServiceReconciler, predicates ...predicate.Predicate)
}

type multiclusterMeshServiceReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterMeshServiceReconcileLoop) AddMulticlusterMeshServiceReconciler(ctx context.Context, rec MulticlusterMeshServiceReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericMeshServiceMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterMeshServiceReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterMeshServiceReconcileLoop {
	return &multiclusterMeshServiceReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &discovery_smh_solo_io_v1alpha1.MeshService{})}
}

type genericMeshServiceMulticlusterReconciler struct {
	reconciler MulticlusterMeshServiceReconciler
}

func (g genericMeshServiceMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterMeshServiceDeletionReconciler); ok {
		return deletionReconciler.ReconcileMeshServiceDeletion(cluster, req)
	}
	return nil
}

func (g genericMeshServiceMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*discovery_smh_solo_io_v1alpha1.MeshService)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: MeshService handler received event for %T", object)
	}
	return g.reconciler.ReconcileMeshService(cluster, obj)
}

// Reconcile Upsert events for the MeshWorkload Resource across clusters.
// implemented by the user
type MulticlusterMeshWorkloadReconciler interface {
	ReconcileMeshWorkload(clusterName string, obj *discovery_smh_solo_io_v1alpha1.MeshWorkload) (reconcile.Result, error)
}

// Reconcile deletion events for the MeshWorkload Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterMeshWorkloadDeletionReconciler interface {
	ReconcileMeshWorkloadDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterMeshWorkloadReconcilerFuncs struct {
	OnReconcileMeshWorkload         func(clusterName string, obj *discovery_smh_solo_io_v1alpha1.MeshWorkload) (reconcile.Result, error)
	OnReconcileMeshWorkloadDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterMeshWorkloadReconcilerFuncs) ReconcileMeshWorkload(clusterName string, obj *discovery_smh_solo_io_v1alpha1.MeshWorkload) (reconcile.Result, error) {
	if f.OnReconcileMeshWorkload == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileMeshWorkload(clusterName, obj)
}

func (f *MulticlusterMeshWorkloadReconcilerFuncs) ReconcileMeshWorkloadDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileMeshWorkloadDeletion == nil {
		return nil
	}
	return f.OnReconcileMeshWorkloadDeletion(clusterName, req)
}

type MulticlusterMeshWorkloadReconcileLoop interface {
	// AddMulticlusterMeshWorkloadReconciler adds a MulticlusterMeshWorkloadReconciler to the MulticlusterMeshWorkloadReconcileLoop.
	AddMulticlusterMeshWorkloadReconciler(ctx context.Context, rec MulticlusterMeshWorkloadReconciler, predicates ...predicate.Predicate)
}

type multiclusterMeshWorkloadReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterMeshWorkloadReconcileLoop) AddMulticlusterMeshWorkloadReconciler(ctx context.Context, rec MulticlusterMeshWorkloadReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericMeshWorkloadMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterMeshWorkloadReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterMeshWorkloadReconcileLoop {
	return &multiclusterMeshWorkloadReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &discovery_smh_solo_io_v1alpha1.MeshWorkload{})}
}

type genericMeshWorkloadMulticlusterReconciler struct {
	reconciler MulticlusterMeshWorkloadReconciler
}

func (g genericMeshWorkloadMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterMeshWorkloadDeletionReconciler); ok {
		return deletionReconciler.ReconcileMeshWorkloadDeletion(cluster, req)
	}
	return nil
}

func (g genericMeshWorkloadMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*discovery_smh_solo_io_v1alpha1.MeshWorkload)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: MeshWorkload handler received event for %T", object)
	}
	return g.reconciler.ReconcileMeshWorkload(cluster, obj)
}

// Reconcile Upsert events for the Mesh Resource across clusters.
// implemented by the user
type MulticlusterMeshReconciler interface {
	ReconcileMesh(clusterName string, obj *discovery_smh_solo_io_v1alpha1.Mesh) (reconcile.Result, error)
}

// Reconcile deletion events for the Mesh Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterMeshDeletionReconciler interface {
	ReconcileMeshDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterMeshReconcilerFuncs struct {
	OnReconcileMesh         func(clusterName string, obj *discovery_smh_solo_io_v1alpha1.Mesh) (reconcile.Result, error)
	OnReconcileMeshDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterMeshReconcilerFuncs) ReconcileMesh(clusterName string, obj *discovery_smh_solo_io_v1alpha1.Mesh) (reconcile.Result, error) {
	if f.OnReconcileMesh == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileMesh(clusterName, obj)
}

func (f *MulticlusterMeshReconcilerFuncs) ReconcileMeshDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileMeshDeletion == nil {
		return nil
	}
	return f.OnReconcileMeshDeletion(clusterName, req)
}

type MulticlusterMeshReconcileLoop interface {
	// AddMulticlusterMeshReconciler adds a MulticlusterMeshReconciler to the MulticlusterMeshReconcileLoop.
	AddMulticlusterMeshReconciler(ctx context.Context, rec MulticlusterMeshReconciler, predicates ...predicate.Predicate)
}

type multiclusterMeshReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterMeshReconcileLoop) AddMulticlusterMeshReconciler(ctx context.Context, rec MulticlusterMeshReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericMeshMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterMeshReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterMeshReconcileLoop {
	return &multiclusterMeshReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &discovery_smh_solo_io_v1alpha1.Mesh{})}
}

type genericMeshMulticlusterReconciler struct {
	reconciler MulticlusterMeshReconciler
}

func (g genericMeshMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterMeshDeletionReconciler); ok {
		return deletionReconciler.ReconcileMeshDeletion(cluster, req)
	}
	return nil
}

func (g genericMeshMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*discovery_smh_solo_io_v1alpha1.Mesh)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Mesh handler received event for %T", object)
	}
	return g.reconciler.ReconcileMesh(cluster, obj)
}
