package inject

import (
	birdsv1beta1 "github.com/heptio/tgik/episodes/040/live/pkg/apis/birds/v1beta1"
	rscheme "github.com/heptio/tgik/episodes/040/live/pkg/client/clientset/versioned/scheme"
	"github.com/heptio/tgik/episodes/040/live/pkg/controller/puffin"
	"github.com/heptio/tgik/episodes/040/live/pkg/inject/args"
	"github.com/kubernetes-sigs/kubebuilder/pkg/inject/run"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	rscheme.AddToScheme(scheme.Scheme)

	// Inject Informers
	Inject = append(Inject, func(arguments args.InjectArgs) error {
		Injector.ControllerManager = arguments.ControllerManager

		if err := arguments.ControllerManager.AddInformerProvider(&birdsv1beta1.Puffin{}, arguments.Informers.Birds().V1beta1().Puffins()); err != nil {
			return err
		}

		// Add Kubernetes informers

		if c, err := puffin.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		return nil
	})

	// Inject CRDs
	Injector.CRDs = append(Injector.CRDs, &birdsv1beta1.PuffinCRD)
	// Inject PolicyRules
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{"birds.fabulous.af"},
		Resources: []string{"*"},
		Verbs:     []string{"*"},
	})
	// Inject GroupVersions
	Injector.GroupVersions = append(Injector.GroupVersions, schema.GroupVersion{
		Group:   "birds.fabulous.af",
		Version: "v1beta1",
	})
	Injector.RunFns = append(Injector.RunFns, func(arguments run.RunArguments) error {
		Injector.ControllerManager.RunInformersAndControllers(arguments)
		return nil
	})
}
