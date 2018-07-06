package puffin

import (
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	"k8s.io/client-go/tools/record"

	birdsv1beta1 "github.com/heptio/tgik/episodes/040/live/pkg/apis/birds/v1beta1"
	birdsv1beta1client "github.com/heptio/tgik/episodes/040/live/pkg/client/clientset/versioned/typed/birds/v1beta1"
	birdsv1beta1informer "github.com/heptio/tgik/episodes/040/live/pkg/client/informers/externalversions/birds/v1beta1"
	birdsv1beta1lister "github.com/heptio/tgik/episodes/040/live/pkg/client/listers/birds/v1beta1"

	"github.com/heptio/tgik/episodes/040/live/pkg/inject/args"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for Puffin resources goes here.

func (bc *PuffinController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE

	if !bc.logged {
		log.Printf("Implement the Reconcile function on puffin.PuffinController to reconcile %s\n", k.Name)
		bc.logged = true

		puffin, err := bc.puffinclient.Puffins(k.Namespace).Get(k.Name, v1.GetOptions{})
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		puffin.Status.Message = fmt.Sprintf("We operated on this puffin: %s", puffin.Spec.Color)
		bc.puffinclient.Puffins(k.Namespace).Update(puffin)
	}


	return nil
}

// +kubebuilder:controller:group=birds,version=v1beta1,kind=Puffin,resource=puffins
type PuffinController struct {
	// INSERT ADDITIONAL FIELDS HERE
	puffinLister birdsv1beta1lister.PuffinLister
	puffinclient birdsv1beta1client.BirdsV1beta1Interface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	puffinrecorder record.EventRecorder

	// After we log something, we will flip this bit
	logged bool
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &PuffinController{
		puffinLister: arguments.ControllerManager.GetInformerProvider(&birdsv1beta1.Puffin{}).(birdsv1beta1informer.PuffinInformer).Lister(),

		puffinclient:   arguments.Clientset.BirdsV1beta1(),
		puffinrecorder: arguments.CreateRecorder("PuffinController"),
	}

	// Create a new controller that will call PuffinController.Reconcile on changes to Puffins
	gc := &controller.GenericController{
		Name:             "PuffinController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&birdsv1beta1.Puffin{}); err != nil {
		return gc, err
	}

	// IMPORTANT:
	// To watch additional resource types - such as those created by your controller - add gc.Watch* function calls here
	// Watch function calls will transform each object event into a Puffin Key to be reconciled by the controller.
	//
	// **********
	// For any new Watched types, you MUST add the appropriate // +kubebuilder:informer and // +kubebuilder:rbac
	// annotations to the PuffinController and run "kubebuilder generate.
	// This will generate the code to start the informers and create the RBAC rules needed for running in a cluster.
	// See:
	// https://godoc.org/github.com/kubernetes-sigs/kubebuilder/pkg/gen/controller#example-package
	// **********

	return gc, nil
}
