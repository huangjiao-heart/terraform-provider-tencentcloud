package taas

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

func (me *TaasService) DescribeTaasRHInstanceById(ctx context.Context, instanceId string) (RHInstance *taas.ReadOnlyRHInstance, errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := taas.NewDescribeRHInstancesRequest()
	request.InstanceIdList = []*string{&instanceId}

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
	instances := make([]*taas.ReadOnlyRHInstance, 0)
	for {
		request.PageNumber = &offset
		request.PageSize = &limit
		response, err := me.client.UseTaasClient().DescribeRHInstances(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.InstanceSet) < 1 {
			return
		}
		instances = append(instances, response.Response.InstanceSet...)
		if len(response.Response.InstanceSet) < int(limit) {
			break
		}

		offset += limit
	}

	RHInstance = instances[0]
	return
}
func (me *TaasService) DeleteTaasRHInstanceById(ctx context.Context, instanceId string) (errRet error) {
	logId := tccommon.GetLogId(ctx)

	request := taas.NewDeleteRHInstanceRequest()
	request.InstanceId = &instanceId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTaasClient().DeleteRHInstance(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}
