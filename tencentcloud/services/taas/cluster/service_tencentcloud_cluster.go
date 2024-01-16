package cluster

import (
	"context"
	//taas "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/taas/v20230927"
	taas "github.com/tencentcloudstack/terraform-provider-tencentcloud/taas/v20230927"
	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/connectivity"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/ratelimit"
	"log"
)

func NewTaasService(client *connectivity.TencentCloudClient) TaasService {
	return TaasService{client: client}
}

type TaasService struct {
	client *connectivity.TencentCloudClient
}

func (me *TaasService) DescribeTaasClusterById(ctx context.Context, clusterId string) (Cluster *taas.ReadOnlyCluster, errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := taas.NewDescribeClustersRequest()
	request.ClusterIdList = []*string{&clusterId}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	var (
		offset int64 = 0
		limit  int64 = 20
	)
	instances := make([]*taas.ReadOnlyCluster, 0)
	for {
		request.PageNumber = &offset
		request.PageSize = &limit
		response, err := me.client.UseTaasClient().DescribeClusters(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.InstanceSet) < 1 {
			break
		}
		instances = append(instances, response.Response.InstanceSet...)
		if len(response.Response.InstanceSet) < int(limit) {
			break
		}

		offset += limit
	}

	Cluster = instances[0]
	return
}
func (me *TaasService) DeleteTaasClusterById(ctx context.Context, clusterId string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := taas.NewDeleteClusterRequest()
	request.ClusterId = &clusterId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTaasClient().DeleteCluster(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}
