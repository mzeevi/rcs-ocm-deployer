package e2e_tests

import (
	mock "github.com/dana-team/rcs-ocm-deployer/test/e2e_tests/mocks"
	"github.com/dana-team/rcs-ocm-deployer/test/e2e_tests/testconsts"
	utilst "github.com/dana-team/rcs-ocm-deployer/test/e2e_tests/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = PDescribe("Validate the placement controller", func() {

	It("Should update the site in status and an annotation", func() {
		baseCapp := mock.CreateBaseCapp()
		desiredCapp := utilst.CreateCapp(k8sClient, baseCapp)

		By("Checks unique creation of Capp")
		assertionCapp := utilst.GetCapp(k8sClient, desiredCapp.Name, desiredCapp.Namespace)
		Expect(assertionCapp.Name).ShouldNot(Equal(baseCapp.Name))

		By("Checks if Capp got the site in status and in annotation")
		Eventually(func() string {
			assertionCapp = utilst.GetCapp(k8sClient, assertionCapp.Name, assertionCapp.Namespace)
			return assertionCapp.Status.ApplicationLinks.Site
		}, testconsts.TimeoutCapp, testconsts.CappCreationInterval).ShouldNot(Equal(""), "Should fetch capp.")
		Eventually(func() string {
			assertionCapp = utilst.GetCapp(k8sClient, assertionCapp.Name, assertionCapp.Namespace)
			return assertionCapp.Annotations["dana.io/has-placement"]
		}, testconsts.TimeoutCapp, testconsts.CappCreationInterval).ShouldNot(Equal(""), "Should fetch capp.")

	})

	It("Should update a site from placement in status and an annotation", func() {
		baseCapp := mock.CreateBaseCapp()
		baseCapp.Spec.Site = testconsts.CappPlacement
		desiredCapp := utilst.CreateCapp(k8sClient, baseCapp)

		By("Checks unique creation of Capp")
		assertionCapp := utilst.GetCapp(k8sClient, desiredCapp.Name, desiredCapp.Namespace)
		Expect(assertionCapp.Name).ShouldNot(Equal(baseCapp.Name))

		By("Checks if got the site in status and in annotation")
		Eventually(func() bool {
			assertionCapp = utilst.GetCapp(k8sClient, assertionCapp.Name, assertionCapp.Namespace)
			status, err := utilst.IsSiteInPlacement(k8sClient, assertionCapp.Status.ApplicationLinks.Site, assertionCapp.Spec.Site, "test")
			Expect(err).Should(BeNil())
			return status
		}, testconsts.TimeoutCapp, testconsts.CappCreationInterval).Should(BeTrue(), "Should fetch capp.")
		Eventually(func() bool {
			assertionCapp = utilst.GetCapp(k8sClient, assertionCapp.Name, assertionCapp.Namespace)
			status, err := utilst.IsSiteInPlacement(k8sClient, assertionCapp.Annotations["dana.io/has-placement"], assertionCapp.Spec.Site, "test")
			Expect(err).Should(BeNil())
			return status
		}, testconsts.TimeoutCapp, testconsts.CappCreationInterval).Should(BeTrue(), "Should fetch capp.")
	})

	It("Should update the selected cluster in status and an annotation ", func() {
		baseCapp := mock.CreateBaseCapp()
		baseCapp.Spec.Site = testconsts.CappCluster
		desiredCapp := utilst.CreateCapp(k8sClient, baseCapp)

		By("Checks unique creation of Capp")
		assertionCapp := utilst.GetCapp(k8sClient, desiredCapp.Name, desiredCapp.Namespace)
		Expect(assertionCapp.Name).ShouldNot(Equal(baseCapp.Name))

		By("Checks if got the cluster in status and in annotation")
		Eventually(func() bool {
			assertionCapp = utilst.GetCapp(k8sClient, assertionCapp.Name, assertionCapp.Namespace)
			return assertionCapp.Status.ApplicationLinks.Site == assertionCapp.Spec.Site
		}, testconsts.TimeoutCapp, testconsts.CappCreationInterval).Should(BeTrue(), "Should fetch capp.")
		Eventually(func() bool {
			assertionCapp = utilst.GetCapp(k8sClient, assertionCapp.Name, assertionCapp.Namespace)
			return assertionCapp.Annotations["dana.io/has-placement"] == assertionCapp.Spec.Site
		}, testconsts.TimeoutCapp, testconsts.CappCreationInterval).Should(BeTrue(), "Should fetch capp.")
	})
})
