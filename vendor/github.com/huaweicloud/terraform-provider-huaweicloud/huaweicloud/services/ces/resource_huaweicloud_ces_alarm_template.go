// ---------------------------------------------------------------
// *** AUTO GENERATED CODE ***
// @Product CES
// ---------------------------------------------------------------

package ces

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jmespath/go-jmespath"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func ResourceCesAlarmTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCesAlarmTemplateCreate,
		UpdateContext: resourceCesAlarmTemplateUpdate,
		ReadContext:   resourceCesAlarmTemplateRead,
		DeleteContext: resourceCesAlarmTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the name of the CES alarm template.`,
			},
			"policies": {
				Type:        schema.TypeList,
				Elem:        AlarmTemplatePolicySchema(),
				Required:    true,
				Description: `Specifies the policy list of the CES alarm template.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the description of the CES alarm template.`,
			},
			"delete_associate_alarm": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Specifies whether delete the alarm rule which the alarm template associated with`,
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Indicates the type of the CES alarm template.`,
			},
			"association_alarm_total": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Indicates the total num of the alarm that associated with the alarm template.`,
			},
		},
	}
}

func AlarmTemplatePolicySchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the namespace of the service.`,
			},
			"dimension_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the Resource dimension.`,
			},
			"metric_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the alarm metric name.`,
			},
			"period": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Specifies the judgment period of alarm condition.`,
			},
			"filter": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the data rollup methods.`,
			},
			"comparison_operator": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the comparison conditions for alarm threshold.`,
			},
			"value": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Specifies the alarm threshold.`,
			},
			"unit": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the unit string of the alarm threshold.`,
			},
			"count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Specifies the number of consecutive triggering of alarms.`,
			},
			"suppress_duration": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Specifies the alarm suppression cycle.`,
			},
			"alarm_level": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the alarm level.`,
			},
		},
	}
	return &sc
}

func resourceCesAlarmTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// createAlarmTemplate: create CES alarm template
	var (
		createAlarmTemplateHttpUrl = "v2/{project_id}/alarm-templates"
		createAlarmTemplateProduct = "ces"
	)
	createAlarmTemplateClient, err := cfg.NewServiceClient(createAlarmTemplateProduct, region)
	if err != nil {
		return diag.Errorf("error creating CES Client: %s", err)
	}

	createAlarmTemplatePath := createAlarmTemplateClient.Endpoint + createAlarmTemplateHttpUrl
	createAlarmTemplatePath = strings.ReplaceAll(createAlarmTemplatePath, "{project_id}",
		createAlarmTemplateClient.ProjectID)

	createAlarmTemplateOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			201,
		},
	}
	createAlarmTemplateOpt.JSONBody = utils.RemoveNil(buildAlarmTemplateBodyParams(d))
	createAlarmTemplateResp, err := createAlarmTemplateClient.Request("POST",
		createAlarmTemplatePath, &createAlarmTemplateOpt)
	if err != nil {
		return diag.Errorf("error creating CES alarm template: %s", err)
	}

	createAlarmTemplateRespBody, err := utils.FlattenResponse(createAlarmTemplateResp)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := jmespath.Search("template_id", createAlarmTemplateRespBody)
	if err != nil {
		return diag.Errorf("error creating CES alarm template: ID is not found in API response")
	}
	d.SetId(id.(string))

	return resourceCesAlarmTemplateRead(ctx, d, meta)
}

func resourceCesAlarmTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	updateAlarmTemplateHasChanges := []string{
		"name",
		"description",
		"policies",
	}

	if d.HasChanges(updateAlarmTemplateHasChanges...) {
		// updateAlarmTemplate: update CES alarm template
		var (
			updateAlarmTemplateHttpUrl = "v2/{project_id}/alarm-templates/{template_id}"
			updateAlarmTemplateProduct = "ces"
		)
		updateAlarmTemplateClient, err := cfg.NewServiceClient(updateAlarmTemplateProduct, region)
		if err != nil {
			return diag.Errorf("error creating CES Client: %s", err)
		}

		updateAlarmTemplatePath := updateAlarmTemplateClient.Endpoint + updateAlarmTemplateHttpUrl
		updateAlarmTemplatePath = strings.ReplaceAll(updateAlarmTemplatePath, "{project_id}",
			updateAlarmTemplateClient.ProjectID)
		updateAlarmTemplatePath = strings.ReplaceAll(updateAlarmTemplatePath, "{template_id}", d.Id())

		updateAlarmTemplateOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			OkCodes: []int{
				204,
			},
		}
		updateAlarmTemplateOpt.JSONBody = utils.RemoveNil(buildAlarmTemplateBodyParams(d))
		_, err = updateAlarmTemplateClient.Request("PUT", updateAlarmTemplatePath, &updateAlarmTemplateOpt)
		if err != nil {
			return diag.Errorf("error updating CES alarm template: %s", err)
		}
	}
	return resourceCesAlarmTemplateRead(ctx, d, meta)
}

func buildAlarmTemplateBodyParams(d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"template_name":        utils.ValueIngoreEmpty(d.Get("name")),
		"template_description": utils.ValueIngoreEmpty(d.Get("description")),
		"policies":             buildAlarmTemplatePoliciesChildBody(d),
	}
	return bodyParams
}

func buildAlarmTemplatePoliciesChildBody(d *schema.ResourceData) []map[string]interface{} {
	rawParams := d.Get("policies").([]interface{})
	if len(rawParams) == 0 {
		return nil
	}

	params := make([]map[string]interface{}, 0, len(rawParams))
	for _, rawParam := range rawParams {
		raw := rawParam.(map[string]interface{})
		param := map[string]interface{}{
			"namespace":           utils.ValueIngoreEmpty(raw["namespace"]),
			"dimension_name":      utils.ValueIngoreEmpty(raw["dimension_name"]),
			"metric_name":         utils.ValueIngoreEmpty(raw["metric_name"]),
			"period":              utils.ValueIngoreEmpty(raw["period"]),
			"filter":              utils.ValueIngoreEmpty(raw["filter"]),
			"comparison_operator": utils.ValueIngoreEmpty(raw["comparison_operator"]),
			"value":               raw["value"],
			"unit":                utils.ValueIngoreEmpty(raw["unit"]),
			"count":               utils.ValueIngoreEmpty(raw["count"]),
			"alarm_level":         utils.ValueIngoreEmpty(raw["alarm_level"]),
			"suppress_duration":   raw["suppress_duration"],
		}
		params = append(params, param)
	}
	return params
}

func resourceCesAlarmTemplateRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	var mErr *multierror.Error

	// getAlarmTemplate: Query CES alarm template
	var (
		getAlarmTemplateHttpUrl = "v2/{project_id}/alarm-templates/{template_id}"
		getAlarmTemplateProduct = "ces"
	)
	getAlarmTemplateClient, err := cfg.NewServiceClient(getAlarmTemplateProduct, region)
	if err != nil {
		return diag.Errorf("error creating CES Client: %s", err)
	}

	getAlarmTemplatePath := getAlarmTemplateClient.Endpoint + getAlarmTemplateHttpUrl
	getAlarmTemplatePath = strings.ReplaceAll(getAlarmTemplatePath, "{project_id}",
		getAlarmTemplateClient.ProjectID)
	getAlarmTemplatePath = strings.ReplaceAll(getAlarmTemplatePath, "{template_id}", d.Id())

	getAlarmTemplateOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	getAlarmTemplateResp, err := getAlarmTemplateClient.Request("GET", getAlarmTemplatePath, &getAlarmTemplateOpt)

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving CES alarm template")
	}

	getAlarmTemplateRespBody, err := utils.FlattenResponse(getAlarmTemplateResp)
	if err != nil {
		return diag.FromErr(err)
	}

	mErr = multierror.Append(
		mErr,
		d.Set("region", region),
		d.Set("name", utils.PathSearch("template_name", getAlarmTemplateRespBody, nil)),
		d.Set("type", utils.PathSearch("template_type", getAlarmTemplateRespBody, nil)),
		d.Set("description", utils.PathSearch("template_description",
			getAlarmTemplateRespBody, nil)),
		d.Set("association_alarm_total", utils.PathSearch("association_alarm_total",
			getAlarmTemplateRespBody, nil)),
		d.Set("policies", flattenGetAlarmTemplateResponseBodyPolicy(getAlarmTemplateRespBody)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenGetAlarmTemplateResponseBodyPolicy(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("policies", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"namespace":           utils.PathSearch("namespace", v, nil),
			"dimension_name":      utils.PathSearch("dimension_name", v, nil),
			"metric_name":         utils.PathSearch("metric_name", v, nil),
			"period":              utils.PathSearch("period", v, nil),
			"filter":              utils.PathSearch("filter", v, nil),
			"comparison_operator": utils.PathSearch("comparison_operator", v, nil),
			"value":               utils.PathSearch("value", v, nil),
			"unit":                utils.PathSearch("unit", v, nil),
			"count":               utils.PathSearch("count", v, nil),
			"alarm_level":         utils.PathSearch("alarm_level", v, nil),
			"suppress_duration":   utils.PathSearch("suppress_duration", v, nil),
		})
	}
	return rst
}

func resourceCesAlarmTemplateDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// deleteAlarmTemplate: Delete CES alarm template
	var (
		deleteAlarmTemplateHttpUrl = "v2/{project_id}/alarm-templates/batch-delete"
		deleteAlarmTemplateProduct = "ces"
	)
	deleteAlarmTemplateClient, err := cfg.NewServiceClient(deleteAlarmTemplateProduct, region)
	if err != nil {
		return diag.Errorf("error creating CES Client: %s", err)
	}

	deleteAlarmTemplatePath := deleteAlarmTemplateClient.Endpoint + deleteAlarmTemplateHttpUrl
	deleteAlarmTemplatePath = strings.ReplaceAll(deleteAlarmTemplatePath, "{project_id}",
		deleteAlarmTemplateClient.ProjectID)

	deleteAlarmTemplateOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	deleteAlarmTemplateOpt.JSONBody = utils.RemoveNil(buildDeleteAlarmTemplateBodyParams(d))
	fmt.Println("")
	_, err = deleteAlarmTemplateClient.Request("POST", deleteAlarmTemplatePath, &deleteAlarmTemplateOpt)
	if err != nil {
		return diag.Errorf("error deleting CES alarm template: %s", err)
	}

	return nil
}

func buildDeleteAlarmTemplateBodyParams(d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"template_ids":           []interface{}{d.Id()},
		"delete_associate_alarm": utils.ValueIngoreEmpty(d.Get("delete_associate_alarm")),
	}
	return bodyParams
}
