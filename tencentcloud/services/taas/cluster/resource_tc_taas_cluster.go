package cluster

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	//taas "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/taas/v20230927"
	taas "github.com/tencentcloudstack/terraform-provider-tencentcloud/taas/v20230927"
	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
	"log"
)

func ResourceTencentCloudTaasCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudTaasClusterCreate,
		Read:   resourceTencentCloudTaasClusterRead,
		Update: resourceTencentCloudTaasClusterUpdate,
		Delete: resourceTencentCloudTaasClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "clusterId 集群 ID",
			},

			"cluster_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "clusterName 集群名",
			},

			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "description 集群描述",
			},

			"public_access": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "publicAccess 是否公网访问",
			},

			"tags": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "tags 集群标签信息",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tag_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "标签键",
						},
						"tag_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "标签值",
						},
					},
				},
			},

			"status": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "status 集群状态 \\n- CLUSTER_PENDING  创建中\\n\\n- CLUSTER_RUNNING 运行中\\n\\n-\n CLUSTER_TERMINATING \\n销毁中",
			},

			"associated_rh_instance": {
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "associatedRHInstance 关联的 ResultHouse 实例信息",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "地域",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ResultHouse 实例 Id",
						},
					},
				},
			},

			"authorized_tcr_registry_infos": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "authorizedTCRRegistryInfos 关联的tcr实例信息",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_account_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "服务凭证id",
						},
						"region": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "地域",
						},
						"registry_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "实例id",
						},
						"namespaces": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional:    true,
							Description: "命名空间",
						},
					},
				},
			},

			"third_party_service_accounts": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "thirdPartyServiceAccounts 关联的第三方凭证信息",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_account_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "业务id",
						},
						"domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "域名",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "用户名",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "密码",
						},
					},
				},
			},

			"vpc_config": {
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "vpcConfig 私有网络参数配置",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "私有网络 的 Id",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "子网的 Id",
						},
					},
				},
			},

			"updated_at": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "updatedAt 更新时间",
			},

			"created_at": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "createdAt 创建时间",
			},
		},
	}
}

func resourceTencentCloudTaasClusterCreate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_taas_cluster.create")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	var (
		request   = taas.NewCreateClusterRequest()
		response  = taas.NewCreateClusterResponse()
		clusterId string
	)
	if v, ok := d.GetOk("cluster_name"); ok {
		request.ClusterName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("public_access"); ok {
		request.PublicAccess = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("tags"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			tag := taas.Tag{}
			if v, ok := dMap["tag_key"]; ok {
				tag.TagKey = helper.String(v.(string))
			}
			if v, ok := dMap["tag_value"]; ok {
				tag.TagValue = helper.String(v.(string))
			}
			request.Tags = append(request.Tags, &tag)
		}
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "associated_rh_instance"); ok {
		associatedRHInstance := taas.AssociatedRHInstance{}
		if v, ok := dMap["region"]; ok {
			associatedRHInstance.Region = helper.String(v.(string))
		}
		if v, ok := dMap["instance_id"]; ok {
			associatedRHInstance.InstanceId = helper.String(v.(string))
		}
		request.AssociatedRHInstance = &associatedRHInstance
	}

	if v, ok := d.GetOk("authorized_tcr_registry_infos"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			authorizedTCRRegistryInfo := taas.AuthorizedTCRRegistryInfo{}
			if v, ok := dMap["service_account_id"]; ok {
				authorizedTCRRegistryInfo.ServiceAccountId = helper.String(v.(string))
			}
			if v, ok := dMap["region"]; ok {
				authorizedTCRRegistryInfo.Region = helper.String(v.(string))
			}
			if v, ok := dMap["registry_id"]; ok {
				authorizedTCRRegistryInfo.RegistryId = helper.String(v.(string))
			}
			if v, ok := dMap["namespaces"]; ok {
				namespacesSet := v.(*schema.Set).List()
				for i := range namespacesSet {
					namespaces := namespacesSet[i].(string)
					authorizedTCRRegistryInfo.Namespaces = append(authorizedTCRRegistryInfo.Namespaces, &namespaces)
				}
			}
			request.AuthorizedTCRRegistryInfos = append(request.AuthorizedTCRRegistryInfos, &authorizedTCRRegistryInfo)
		}
	}

	if v, ok := d.GetOk("third_party_service_accounts"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			thirdPartyServiceAccount := taas.ThirdPartyServiceAccount{}
			if v, ok := dMap["service_account_id"]; ok {
				thirdPartyServiceAccount.ServiceAccountId = helper.String(v.(string))
			}
			if v, ok := dMap["domain"]; ok {
				thirdPartyServiceAccount.Domain = helper.String(v.(string))
			}
			if v, ok := dMap["username"]; ok {
				thirdPartyServiceAccount.Username = helper.String(v.(string))
			}
			if v, ok := dMap["password"]; ok {
				thirdPartyServiceAccount.Password = helper.String(v.(string))
			}
			request.ThirdPartyServiceAccounts = append(request.ThirdPartyServiceAccounts, &thirdPartyServiceAccount)
		}
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "vpc_config"); ok {
		vpcConfig := taas.VpcConfig{}
		if v, ok := dMap["vpc_id"]; ok {
			vpcConfig.VpcId = helper.String(v.(string))
		}
		if v, ok := dMap["subnet_id"]; ok {
			vpcConfig.SubnetId = helper.String(v.(string))
		}
		request.VpcConfig = &vpcConfig
	}

	err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseTaasClient().CreateCluster(request)
		if e != nil {
			return tccommon.RetryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create taas Cluster failed, reason:%+v", logId, err)
		return err
	}
	clusterId = *response.Response.ClusterId
	d.SetId(clusterId)

	return resourceTencentCloudTaasClusterRead(d, meta)
}

func resourceTencentCloudTaasClusterRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_taas_cluster.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := context.WithValue(context.TODO(), tccommon.LogIdKey, logId)

	service := TaasService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	clusterId := d.Id()

	Cluster, err := service.DescribeTaasClusterById(ctx, clusterId)
	if err != nil {
		return err
	}

	if Cluster == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TaasCluster` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("cluster_id", clusterId)
	return nil
}

func resourceTencentCloudTaasClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_taas_cluster.update")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	request := taas.NewModifyClusterAttributesRequest()

	clusterId := d.Id()

	request.ClusterId = &clusterId

	immutableArgs := []string{"cluster_id", "status", "updated_at", "created_at"}
	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	needChange := false
	mutableArgs := []string{"cluster_name", "description", "public_access", "tags", "associated_rh_instance", "authorized_tcr_registry_infos", "third_party_service_accounts", "vpc_config"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		if v, ok := d.GetOk("cluster_name"); ok {
			request.ClusterName = helper.String(v.(string))
		}
		if v, ok := d.GetOk("description"); ok {
			request.Description = helper.String(v.(string))
		}
		if v, ok := d.GetOkExists("public_access"); ok {
			request.PublicAccess = helper.Bool(v.(bool))
		}
		if v, ok := d.GetOk("tags"); ok {
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				tag := taas.Tag{}
				if v, ok := dMap["tag_key"]; ok {
					tag.TagKey = helper.String(v.(string))
				}
				if v, ok := dMap["tag_value"]; ok {
					tag.TagValue = helper.String(v.(string))
				}
				request.Tags = append(request.Tags, &tag)
			}
		}
		if dMap, ok := helper.InterfacesHeadMap(d, "associated_rh_instance"); ok {
			associatedRHInstance := taas.AssociatedRHInstance{}
			if v, ok := dMap["region"]; ok {
				associatedRHInstance.Region = helper.String(v.(string))
			}
			if v, ok := dMap["instance_id"]; ok {
				associatedRHInstance.InstanceId = helper.String(v.(string))
			}
			request.AssociatedRHInstance = &associatedRHInstance
		}
		if v, ok := d.GetOk("authorized_tcr_registry_infos"); ok {
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				authorizedTCRRegistryInfo := taas.AuthorizedTCRRegistryInfo{}
				if v, ok := dMap["service_account_id"]; ok {
					authorizedTCRRegistryInfo.ServiceAccountId = helper.String(v.(string))
				}
				if v, ok := dMap["region"]; ok {
					authorizedTCRRegistryInfo.Region = helper.String(v.(string))
				}
				if v, ok := dMap["registry_id"]; ok {
					authorizedTCRRegistryInfo.RegistryId = helper.String(v.(string))
				}
				if v, ok := dMap["namespaces"]; ok {
					namespacesSet := v.(*schema.Set).List()
					for i := range namespacesSet {
						namespaces := namespacesSet[i].(string)
						authorizedTCRRegistryInfo.Namespaces = append(authorizedTCRRegistryInfo.Namespaces, &namespaces)
					}
				}
				request.AuthorizedTCRRegistryInfos = append(request.AuthorizedTCRRegistryInfos, &authorizedTCRRegistryInfo)
			}
		}
		if v, ok := d.GetOk("third_party_service_accounts"); ok {
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				thirdPartyServiceAccount := taas.ThirdPartyServiceAccount{}
				if v, ok := dMap["service_account_id"]; ok {
					thirdPartyServiceAccount.ServiceAccountId = helper.String(v.(string))
				}
				if v, ok := dMap["domain"]; ok {
					thirdPartyServiceAccount.Domain = helper.String(v.(string))
				}
				if v, ok := dMap["username"]; ok {
					thirdPartyServiceAccount.Username = helper.String(v.(string))
				}
				if v, ok := dMap["password"]; ok {
					thirdPartyServiceAccount.Password = helper.String(v.(string))
				}
				request.ThirdPartyServiceAccounts = append(request.ThirdPartyServiceAccounts, &thirdPartyServiceAccount)
			}
		}
		//if dMap, ok := helper.InterfacesHeadMap(d, "vpc_config"); ok {
		//	vpcConfig := taas.VpcConfig{}
		//	if v, ok := dMap["vpc_id"]; ok {
		//		vpcConfig.VpcId = helper.String(v.(string))
		//	}
		//	if v, ok := dMap["subnet_id"]; ok {
		//		vpcConfig.SubnetId = helper.String(v.(string))
		//	}
		//	request.VpcConfig = &vpcConfig
		//}

		err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
			result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseTaasClient().ModifyClusterAttributes(request)
			if e != nil {
				return tccommon.RetryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update taas Cluster failed, reason:%+v", logId, err)
			return err
		}

	}
	return resourceTencentCloudTaasClusterRead(d, meta)
}

func resourceTencentCloudTaasClusterDelete(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_taas_cluster.delete")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)
	ctx := context.WithValue(context.TODO(), tccommon.LogIdKey, logId)

	service := TaasService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}
	clusterId := d.Id()

	if err := service.DeleteTaasClusterById(ctx, clusterId); err != nil {
		return err
	}

	return nil
}
