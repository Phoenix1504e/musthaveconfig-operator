package controllers

import (
    "context"
    "fmt"

    corev1 "k8s.io/api/core/v1"
    apierrors "k8s.io/apimachinery/pkg/api/errors"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/types"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/log"

    v1alpha1 "github.com/Phoenix1504e/musthaveconfig-operator/api/v1alpha1"
)

// MustHaveConfigReconciler reconciles a MustHaveConfig object
type MustHaveConfigReconciler struct {
    client.Client
    Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=ops.aditya.dev,resources=musthaveconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ops.aditya.dev,resources=musthaveconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ops.aditya.dev,resources=musthaveconfigs/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// Reconcile is called whenever a MustHaveConfig or owned ConfigMap changes.
func (r *MustHaveConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    logger := log.FromContext(ctx)

    // 1. Load the MustHaveConfig object
    var mhc v1alpha1.MustHaveConfig
    if err := r.Get(ctx, req.NamespacedName, &mhc); err != nil {
        if apierrors.IsNotFound(err) {
            // Deleted – nothing to do
            return ctrl.Result{}, nil
        }
        return ctrl.Result{}, err
    }

    logger.Info("Reconciling MustHaveConfig", "name", mhc.Name, "namespace", mhc.Namespace)

    // 2. Figure out where the ConfigMap should live
    targetNamespace := mhc.Spec.Namespace
    if targetNamespace == "" {
        // default to the same namespace as the CR if not set
        targetNamespace = mhc.Namespace
    }
    cmName := mhc.Name

    desiredKey := mhc.Spec.Key
    desiredValue := mhc.Spec.Value

    if desiredKey == "" {
        mhc.Status.Synced = false
        mhc.Status.Message = "spec.key is empty"
        _ = r.Status().Update(ctx, &mhc)
        return ctrl.Result{}, nil
    }

    // 3. Get or create the ConfigMap
    var cm corev1.ConfigMap
    cmNN := types.NamespacedName{Name: cmName, Namespace: targetNamespace}

    err := r.Get(ctx, cmNN, &cm)
    if apierrors.IsNotFound(err) {
        // ConfigMap doesn't exist -> create it
        cm = corev1.ConfigMap{
            ObjectMeta: metav1.ObjectMeta{
                Name:      cmName,
                Namespace: targetNamespace,
            },
            Data: map[string]string{
                desiredKey: desiredValue,
            },
        }

        // make operator the owner so ConfigMap is cleaned up with the CR
        if err := ctrl.SetControllerReference(&mhc, &cm, r.Scheme); err != nil {
            return ctrl.Result{}, err
        }

        if err := r.Create(ctx, &cm); err != nil {
            logger.Error(err, "failed to create ConfigMap")
            mhc.Status.Synced = false
            mhc.Status.Message = fmt.Sprintf("failed to create ConfigMap: %v", err)
            _ = r.Status().Update(ctx, &mhc)
            return ctrl.Result{}, err
        }

        logger.Info("Created ConfigMap", "namespace", targetNamespace, "name", cmName)
    } else if err != nil {
        // Some other error fetching ConfigMap
        return ctrl.Result{}, err
    } else {
        // ConfigMap exists -> ensure key/value matches desired state
        if cm.Data == nil {
            cm.Data = map[string]string{}
        }
        current, ok := cm.Data[desiredKey]
        if !ok || current != desiredValue {
            cm.Data[desiredKey] = desiredValue
            if err := r.Update(ctx, &cm); err != nil {
                logger.Error(err, "failed to update ConfigMap")
                mhc.Status.Synced = false
                mhc.Status.Message = fmt.Sprintf("failed to update ConfigMap: %v", err)
                _ = r.Status().Update(ctx, &mhc)
                return ctrl.Result{}, err
            }
            logger.Info("Updated ConfigMap", "namespace", targetNamespace, "name", cmName)
        }
    }

    // 4. Mark status as synced
    mhc.Status.Synced = true
    mhc.Status.Message = "ConfigMap in desired state"
    if err := r.Status().Update(ctx, &mhc); err != nil {
        logger.Error(err, "failed to update MustHaveConfig status")
        // non-fatal
    }

    return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MustHaveConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&v1alpha1.MustHaveConfig{}).
        Owns(&corev1.ConfigMap{}).
        Complete(r)
}
