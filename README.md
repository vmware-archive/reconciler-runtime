# reconciler-runtime

`reconciler-runtime` builds on top of the [Kubernetes `controller-runtime`](https://github.com/kubernetes-sigs/controller-runtime) project. `controller-runtime` provides infrastructure for creating and operating controllers, but provides little support for the business logic of implementing a reconciler within a controller. The [`Reconciler` interface](https://godoc.org/sigs.k8s.io/controller-runtime/pkg/reconcile#Reconciler) provided by `controller-runtime` is the handoff point with `reconciler-runtime`.

## Reconcilers

### ParentReconciler

A [`ParentReconciler`](https://godoc.org/github.com/projectriff/reconciler-runtime/reconcilers#ParentReconciler) is responsible for orchestrating the reconciliation of a single resource. The reconciler delegates the manipulation of other resources to SubReconcilers.

The parent is responsible for:
- fetching the resource being reconciled
- creating a stash to pass state between sub reconcilers
- passing the resource to each sub reconciler in turn
- reflects the observed generation on the status
- updates the resource status if it was modified
- logging the reconcilers activities
- records events for mutations and errors

The implementor is responsible for:
- defining the set of sub reconcilers

### SubReconciler

The [`SubReconciler`](https://godoc.org/github.com/projectriff/reconciler-runtime/reconcilers#SubReconciler) interface defines the contract between the parent and sub reconcilers.

There are two types of sub reconcilers provided by `reconciler-runtime`:

#### ChildReconciler

The [`ChildReconciler`](https://godoc.org/github.com/projectriff/reconciler-runtime/reconcilers#ChildReconciler) is a sub reconciler that is responsible for managing a single controlled resource. A developer defines their desired state for the child resource (if any), and the reconciler creates/updates/deletes the resource to match the desired state. The child resource is also used to update the parent's status. Mutations and errors are recorded for the parent.

The ChildReconciler is responsible for:
- looking up an existing child
- creating/updating/deleting the child resource based on the desired state
- setting the owner reference on the child resource
- logging the reconcilers activities
- recording child mutations and errors for the parent resource

The implementor is responsible for:
- defining the desired resource
- indicating if two resources are semantically equal
- merging the actual resource with the desired state (often as simple as copying the spec and labels)
- updating the parent's status from the child

#### SyncReconciler

The [`SyncReconciler`](https://godoc.org/github.com/projectriff/reconciler-runtime/reconcilers#SyncReconciler) is the minimal type-aware sub reconciler. It is used to manage a portion of the parent's reconciliation that is custom, or whose behavior is not covered by another sub reconciler type. Common uses include looking up reference data for the reconciliation, or controlling resources that are not kubernetes resources.

### Utilities

#### Config

The [`Config`](https://godoc.org/github.com/projectriff/reconciler-runtime/reconcilers#Config) is a single object that contains the key APIs needed by a reconciler. The config object is provided to the reconciler when initialized and is preconfigured for the reconciler.

#### Stash

The stash allows passing arbitrary state between sub reconcilers within the scope of a single reconciler request. Values are stored on the context by [`StashValue`](https://godoc.org/github.com/projectriff/reconciler-runtime/reconcilers#StashValue) and accessed via [`RetrieveValue`](https://godoc.org/github.com/projectriff/reconciler-runtime/reconcilers#RetrieveValue).

#### Tracker

The [`Tracker`](https://godoc.org/github.com/projectriff/reconciler-runtime/tracker#Tracker) provides a means for one resource to watch another resource for mutations, triggering the reconciliation of the resource defining the reference.

## Testing

While `controller-runtime` focuses its testing efforts on integration testing by spinning up a new API Server and etcd, `reconciler-runtime` focuses on unit testing reconcilers. The state for each test case is pure, preventing side effects from one test case impacting the next.

The table test pattern is used to declare each test case in a test suite with the resource being reconciled, other given resources in the cluster, and all expected resource mutations (create, update, delete).

There are two test suites, one for reconcilers and an optimized harness for testing sub reconcilers.

### ReconcilerTestSuite

[`ReconcilerTestCase`](https://godoc.org/github.com/projectriff/reconciler-runtime/testing#ReconcilerTestCase)

### SubReconcilerTestSuite

[`SubReconcilerTestCase`](https://godoc.org/github.com/projectriff/reconciler-runtime/testing#SubReconcilerTestCase)

## Code of Conduct

Please refer to the [Contributor Code of Conduct](CODE_OF_CONDUCT.adoc).
