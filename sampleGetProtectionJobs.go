package main

import (
	"github.com/cohesity/management-sdk-go/models"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
)

func main() {
	Username := "clusterUsername"
	Password := "clusterPassword"
	ClusterVip := "prod-cluster.eng.cohesity.com"
	var Domain string
	client := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)

	var err error
	protectionJobs := client.ProtectionJobs()
	var result []*models.ProtectionJob
	var includeLastRunAndStats bool
	var policyIds []string
	var isActive bool = true
	var isDeleted bool = false
	var onlyReturnBasicSummary bool = true
	Environments := []models.EnvironmentsEnum{}
	var tenantIds []string
	var allUnderHierarchy bool
	var ids []int64
	var names []string

	result, err = protectionJobs.GetProtectionJobs(&includeLastRunAndStats, policyIds, &isActive, &isDeleted, &onlyReturnBasicSummary, Environments, tenantIds, &allUnderHierarchy, ids, names)
	if err != nil {
		// Handle Error.
	} else {
		if len(result) == 0 {
			println("There are no Protection Jobs in the cluster.")
		} else {
			println("Protection Job Name: ", result[0].Name)
		}
	}
}
