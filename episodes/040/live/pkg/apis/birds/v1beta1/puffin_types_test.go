package v1beta1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/heptio/tgik/episodes/040/live/pkg/apis/birds/v1beta1"
	. "github.com/heptio/tgik/episodes/040/live/pkg/client/clientset/versioned/typed/birds/v1beta1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the Puffin resource tests

var _ = Describe("Puffin", func() {
	var instance Puffin
	var expected Puffin
	var client PuffinInterface

	BeforeEach(func() {
		instance = Puffin{}
		instance.Name = "instance-1"

		expected = instance
	})

	AfterEach(func() {
		client.Delete(instance.Name, &metav1.DeleteOptions{})
	})

	// INSERT YOUR CODE HERE - add more "Describe" tests

	// Automatically created storage tests
	Describe("when sending a storage request", func() {
		Context("for a valid config", func() {
			It("should provide CRUD access to the object", func() {
				client = cs.BirdsV1beta1().Puffins("default")

				By("returning success from the create request")
				actual, err := client.Create(&instance)
				Expect(err).ShouldNot(HaveOccurred())

				By("defaulting the expected fields")
				Expect(actual.Spec).To(Equal(expected.Spec))

				By("returning the item for list requests")
				result, err := client.List(metav1.ListOptions{})
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result.Items).To(HaveLen(1))
				Expect(result.Items[0].Spec).To(Equal(expected.Spec))

				By("returning the item for get requests")
				actual, err = client.Get(instance.Name, metav1.GetOptions{})
				Expect(err).ShouldNot(HaveOccurred())
				Expect(actual.Spec).To(Equal(expected.Spec))

				By("deleting the item for delete requests")
				err = client.Delete(instance.Name, &metav1.DeleteOptions{})
				Expect(err).ShouldNot(HaveOccurred())
				result, err = client.List(metav1.ListOptions{})
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result.Items).To(HaveLen(0))
			})
		})
	})
})
