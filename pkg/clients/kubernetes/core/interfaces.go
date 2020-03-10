package kubernetes_core

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//go:generate mockgen -source ./interfaces.go -destination mocks/mock_clients.go

type ServiceClient interface {
	Get(ctx context.Context, name, namespace string) (*corev1.Service, error)
	// include client.InNamespace(ns) in the options varargs list to specify a namespace
	// always returns a non-nil ServiceList, but its Items field may be empty
	List(ctx context.Context, options ...client.ListOption) (*corev1.ServiceList, error)
}

type SecretsClient interface {
	Create(ctx context.Context, secret *corev1.Secret, opts ...client.CreateOption) error
	Update(ctx context.Context, secret *corev1.Secret, opts ...client.UpdateOption) error
	Get(ctx context.Context, name, namespace string) (*corev1.Secret, error)
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.SecretList, error)
}

type ServiceAccountClient interface {
	// create the service account in the namespace on the resource's ObjectMeta
	Create(ctx context.Context, serviceAccount *corev1.ServiceAccount) error

	Get(ctx context.Context, name, namespace string) (*corev1.ServiceAccount, error)

	// update the service account in the namespace on the resource's ObjectMeta
	Update(ctx context.Context, serviceAccount *corev1.ServiceAccount) error
}

type ConfigMapClient interface {
	Create(ctx context.Context, configMap *corev1.ConfigMap) error
	Get(ctx context.Context, objKey client.ObjectKey) (*corev1.ConfigMap, error)
	Update(ctx context.Context, configMap *corev1.ConfigMap) error
}
