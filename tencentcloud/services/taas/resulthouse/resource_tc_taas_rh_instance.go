package taas

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

func ResourceTencentCloudTaasRHInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudTaasRHInstanceCreate,
		Read:   resourceTencentCloudTaasRHInstanceRead,
		Update: resourceTencentCloudTaasRHInstanceUpdate,
		Delete: resourceTencentCloudTaasRHInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "instanceId",
			},

			"instance_type": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "instanceType",
			},

			"instance_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "instanceName",
			},

			"volume": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "volume",
			},

			"volume_used": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "volumeUsed",
			},

			"renew_flag": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "renewFlag",
			},

			"charge_type": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "chargeType",
			},

			"instance_status": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "instanceStatus",
			},

			"created_time": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "createdTime 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ",
			},

			"expired_time": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "expiredTime 按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ",
			},

			"tags": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "tags",
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

			"period": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "period",
			},

			//"count": {
			//	Optional:    true,
			//	Type:        schema.TypeInt,
			//	Description: "count",
			//},

			"auto_voucher": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "autoVoucher",
			},

			"deal_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "dealName",
			},

			"root_password": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "rootPassword",
			},
		},
	}
}

func resourceTencentCloudTaasRHInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_taas_rh_instance.create")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	var (
		request    = taas.NewCreateRHInstanceRequest()
		response   = taas.NewCreateRHInstanceResponse()
		instanceId string
	)
	if v, ok := d.GetOk("instance_type"); ok {
		request.InstanceType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_name"); ok {
		request.InstanceName = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("volume"); ok {
		request.Volume = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("renew_flag"); ok {
		request.RenewFlag = helper.String(v.(string))
	}

	if v, ok := d.GetOk("charge_type"); ok {
		request.ChargeType = helper.String(v.(string))
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

	if v, ok := d.GetOkExists("period"); ok {
		request.Period = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("count"); ok {
		request.Count = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("auto_voucher"); ok {
		request.AutoVoucher = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("deal_name"); ok {
		request.DealName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("root_password"); ok {
		request.RootPassword = helper.String(v.(string))
	}
	count := int64(1)
	request.Count = &count

	err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseTaasClient().CreateRHInstance(request)
		if e != nil {
			return tccommon.RetryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create taas RHInstance failed, reason:%+v", logId, err)
		return err
	}
	instanceId = *response.Response.InstanceId
	d.SetId(instanceId)

	return resourceTencentCloudTaasRHInstanceRead(d, meta)
}

func resourceTencentCloudTaasRHInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_taas_rh_instance.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := context.WithValue(context.TODO(), tccommon.LogIdKey, logId)

	service := TaasService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	instanceId := d.Id()

	RHInstance, err := service.DescribeTaasRHInstanceById(ctx, instanceId)
	if err != nil {
		return err
	}

	if RHInstance == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TaasRHInstance` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("instance_id", instanceId)
	return nil
}

func resourceTencentCloudTaasRHInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_taas_rh_instance.update")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	request := taas.NewModifyRHInstanceAttributesRequest()

	instanceId := d.Id()

	request.InstanceId = &instanceId

	immutableArgs := []string{"instance_id", "volume_used", "instance_status", "created_time", "expired_time"}
	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	needChange := false
	mutableArgs := []string{"instance_type", "instance_name", "volume", "renew_flag", "charge_type", "tags", "period", "count", "auto_voucher", "deal_name", "root_password"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		if v, ok := d.GetOk("instance_type"); ok {
			request.InstanceType = helper.String(v.(string))
		}
		if v, ok := d.GetOk("instance_name"); ok {
			request.InstanceName = helper.String(v.(string))
		}
		if v, ok := d.GetOkExists("volume"); ok {
			request.Volume = helper.IntInt64(v.(int))
		}
		if v, ok := d.GetOk("renew_flag"); ok {
			request.RenewFlag = helper.String(v.(string))
		}
		//if v, ok := d.GetOk("charge_type"); ok {
		//	request.ChargeType = helper.String(v.(string))
		//}
		//if v, ok := d.GetOk("tags"); ok {
		//	for _, item := range v.([]interface{}) {
		//		dMap := item.(map[string]interface{})
		//		tag := taas.Tag{}
		//		if v, ok := dMap["tag_key"]; ok {
		//			tag.TagKey = helper.String(v.(string))
		//		}
		//		if v, ok := dMap["tag_value"]; ok {
		//			tag.TagValue = helper.String(v.(string))
		//		}
		//		request.Tags = append(request.Tags, &tag)
		//	}
		//}
		if v, ok := d.GetOkExists("period"); ok {
			request.Period = helper.IntInt64(v.(int))
		}
		//if v, ok := d.GetOkExists("count"); ok {
		//	request.Count = helper.IntInt64(v.(int))
		//}
		if v, ok := d.GetOkExists("auto_voucher"); ok {
			request.AutoVoucher = helper.Bool(v.(bool))
		}
		//if v, ok := d.GetOk("deal_name"); ok {
		//	request.DealName = helper.String(v.(string))
		//}
		//if v, ok := d.GetOk("root_password"); ok {
		//	request.RootPassword = helper.String(v.(string))
		//}

		err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
			result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseTaasClient().ModifyRHInstanceAttributes(request)
			if e != nil {
				return tccommon.RetryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update taas RHInstance failed, reason:%+v", logId, err)
			return err
		}

	}
	return resourceTencentCloudTaasRHInstanceRead(d, meta)
}

func resourceTencentCloudTaasRHInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_taas_rh_instance.delete")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)
	ctx := context.WithValue(context.TODO(), tccommon.LogIdKey, logId)

	service := TaasService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}
	instanceId := d.Id()

	if err := service.DeleteTaasRHInstanceById(ctx, instanceId); err != nil {
		return err
	}

	return nil
}
