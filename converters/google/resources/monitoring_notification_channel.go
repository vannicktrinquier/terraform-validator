// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var sensitiveLabels = []string{"auth_token", "service_key", "password"}

func sensitiveLabelCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	for _, sl := range sensitiveLabels {
		mapLabel := diff.Get("labels." + sl).(string)
		authLabel := diff.Get("sensitive_labels.0." + sl).(string)
		if mapLabel != "" && authLabel != "" {
			return fmt.Errorf("Sensitive label [%s] cannot be set in both `labels` and the `sensitive_labels` block.", sl)
		}
	}
	return nil
}

const MonitoringNotificationChannelAssetType string = "monitoring.googleapis.com/NotificationChannel"

func resourceConverterMonitoringNotificationChannel() ResourceConverter {
	return ResourceConverter{
		AssetType: MonitoringNotificationChannelAssetType,
		Convert:   GetMonitoringNotificationChannelCaiObject,
	}
}

func GetMonitoringNotificationChannelCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//monitoring.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetMonitoringNotificationChannelApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: MonitoringNotificationChannelAssetType,
			Resource: &AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/monitoring/v3/rest",
				DiscoveryName:        "NotificationChannel",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetMonitoringNotificationChannelApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandMonitoringNotificationChannelLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	typeProp, err := expandMonitoringNotificationChannelType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	userLabelsProp, err := expandMonitoringNotificationChannelUserLabels(d.Get("user_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("user_labels"); !isEmptyValue(reflect.ValueOf(userLabelsProp)) && (ok || !reflect.DeepEqual(v, userLabelsProp)) {
		obj["userLabels"] = userLabelsProp
	}
	descriptionProp, err := expandMonitoringNotificationChannelDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandMonitoringNotificationChannelDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandMonitoringNotificationChannelEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); ok || !reflect.DeepEqual(v, enabledProp) {
		obj["enabled"] = enabledProp
	}

	return resourceMonitoringNotificationChannelEncoder(d, config, obj)
}

func resourceMonitoringNotificationChannelEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	labelmap, ok := obj["labels"]
	if !ok {
		labelmap = make(map[string]string)
	}

	var labels map[string]string
	labels = labelmap.(map[string]string)

	for _, sl := range sensitiveLabels {
		if auth, _ := d.GetOkExists("sensitive_labels.0." + sl); auth != "" {
			labels[sl] = auth.(string)
		}
	}

	obj["labels"] = labels

	return obj, nil
}

func expandMonitoringNotificationChannelLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMonitoringNotificationChannelType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringNotificationChannelUserLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMonitoringNotificationChannelDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringNotificationChannelDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringNotificationChannelEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}