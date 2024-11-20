package k8s

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestListPersistentVolumes(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&v1.PersistentVolume{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pv1",
			},
		},
		&v1.PersistentVolume{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pv2",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	pvs, err := k8sHandler.ListPersistentVolumes()
	if err != nil {
		t.Fatalf("ListPersistentVolumes failed: %v", err)
	}

	if len(pvs.Items) != 2 {
		t.Errorf("expected 2 PersistentVolumes, got %d", len(pvs.Items))
	}
}

func TestListPersistentVolumeClaims(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&v1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "pvc1",
				Namespace: "default",
			},
		},
		&v1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "pvc2",
				Namespace: "other-namespace",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	pvcs, err := k8sHandler.ListPersistentVolumeClaims("default")
	if err != nil {
		t.Fatalf("ListPersistentVolumeClaims failed: %v", err)
	}

	if len(pvcs.Items) != 1 {
		t.Errorf("expected 1 PersistentVolumeClaim in 'default' namespace, got %d", len(pvcs.Items))
	}

	allPvcs, err := k8sHandler.ListPersistentVolumeClaims("")
	if err != nil {
		t.Fatalf("ListPersistentVolumeClaims for all namespaces failed: %v", err)
	}

	if len(allPvcs.Items) != 2 {
		t.Errorf("expected 2 PersistentVolumeClaims across all namespaces, got %d", len(allPvcs.Items))
	}
}

func TestListStorageClasses(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&storagev1.StorageClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: "sc1",
			},
		},
		&storagev1.StorageClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: "sc2",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	storageClasses, err := k8sHandler.ListStorageClasses()
	if err != nil {
		t.Fatalf("ListStorageClasses failed: %v", err)
	}

	if len(storageClasses.Items) != 2 {
		t.Errorf("expected 2 StorageClasses, got %d", len(storageClasses.Items))
	}
}

func TestGetPersistentVolume(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&v1.PersistentVolume{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pv1",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	pv, err := k8sHandler.GetPersistentVolume("pv1")
	if err != nil {
		t.Fatalf("GetPersistentVolume failed: %v", err)
	}

	if pv.Name != "pv1" {
		t.Errorf("expected PersistentVolume name 'pv1', got %s", pv.Name)
	}

	_, err = k8sHandler.GetPersistentVolume("")
	if err == nil || err.Error() != "persistent volume name must be provided" {
		t.Errorf("expected error 'persistent volume name must be provided', got %v", err)
	}
}

func TestGetPersistentVolumeClaim(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&v1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "pvc1",
				Namespace: "default",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	pvc, err := k8sHandler.GetPersistentVolumeClaim("pvc1", "default")
	if err != nil {
		t.Fatalf("GetPersistentVolumeClaim failed: %v", err)
	}

	if pvc.Name != "pvc1" {
		t.Errorf("expected PersistentVolumeClaim name 'pvc1', got %s", pvc.Name)
	}

	_, err = k8sHandler.GetPersistentVolumeClaim("", "default")
	if err == nil || err.Error() != "persistent volume claim name must be provided" {
		t.Errorf("expected error 'persistent volume claim name must be provided', got %v", err)
	}
}

func TestGetStorageClass(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&storagev1.StorageClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: "sc1",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	sc, err := k8sHandler.GetStorageClass("sc1")
	if err != nil {
		t.Fatalf("GetStorageClass failed: %v", err)
	}

	if sc.Name != "sc1" {
		t.Errorf("expected StorageClass name 'sc1', got %s", sc.Name)
	}

	_, err = k8sHandler.GetStorageClass("")
	if err == nil || err.Error() != "storage class name must be provided" {
		t.Errorf("expected error 'storage class name must be provided', got %v", err)
	}
}
