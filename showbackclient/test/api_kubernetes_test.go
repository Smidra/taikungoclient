/*
Taikun - WebApi

Testing KubernetesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package taikunshowback

import (
	"context"
	openapiclient "github.com/Smidra/taikungoclient/showbackclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_taikunshowback_KubernetesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KubernetesAPIService KubernetesAddK8sAlert", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		httpRes, err := apiClient.KubernetesAPI.KubernetesAddK8sAlert(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesAddK8sEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		httpRes, err := apiClient.KubernetesAPI.KubernetesAddK8sEvents(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesAlertList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesAlertList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesCli", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesCli(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesConfigMapList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesConfigMapList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesCrdList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesCrdList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesCronJobList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesCronJobList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDaemonSetList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDaemonSetList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDashboardList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDashboardList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDeploymentList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDeploymentList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeConfigMap", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeConfigMap(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeCrd", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeCrd(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeCronjob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeCronjob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeDaemonSet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeDaemonSet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeDeployment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeDeployment(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeIngress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeIngress(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeNetworkPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeNetworkPolicy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeNode(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribePdb", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribePdb(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribePod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribePod(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribePvc", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribePvc(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeSecret", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeSecret(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeService", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeService(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeStorageClass", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeStorageClass(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeSts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeSts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDownload", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDownload(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesExport", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesExport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesGetSupportedList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesGetSupportedList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesHelmReleaseList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesHelmReleaseList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesIngressList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesIngressList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesInteractiveShell", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesInteractiveShell(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesJobsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesJobsList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesKillPod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32
		var metadataName string
		var namespace string

		httpRes, err := apiClient.KubernetesAPI.KubernetesKillPod(context.Background(), projectId, metadataName, namespace).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesKubeConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesKubeConfig(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesNamespaceList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesNamespaceList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesNetworkPolicyList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesNetworkPolicyList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesNodeList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesNodeList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesOverview", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesOverview(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPatchCrd", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesPatchCrd(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPatchCronJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesPatchCronJob(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPatchIngress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesPatchIngress(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPatchJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesPatchJob(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPatchNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesPatchNode(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPatchPdb", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesPatchPdb(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPatchPod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesPatchPod(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPatchPvc", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesPatchPvc(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPatchSecret", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesPatchSecret(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPatchSts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesPatchSts(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPdbList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesPdbList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPodList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesPodList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPodLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesPodLogs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPvcList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesPvcList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesQuota", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesQuota(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesRemovealerts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesRemovealerts(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesRestartDaemonSet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesRestartDaemonSet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesRestartDeployment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesRestartDeployment(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesRestartSts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesRestartSts(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesSecretList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesSecretList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesServiceList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesServiceList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesSilenceManager", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.KubernetesAPI.KubernetesSilenceManager(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesStorageClassList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesStorageClassList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesStsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesStsList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesUpdateAlert", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var alertId int32

		httpRes, err := apiClient.KubernetesAPI.KubernetesUpdateAlert(context.Background(), alertId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
