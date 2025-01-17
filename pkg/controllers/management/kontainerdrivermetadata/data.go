package kontainerdrivermetadata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/blang/semver"
	mVersion "github.com/mcuadros/go-version"
	"github.com/rancher/norman/types/convert"
	setting2 "github.com/rancher/rancher/pkg/api/store/setting"
	"github.com/rancher/rancher/pkg/namespace"
	"github.com/rancher/rancher/pkg/settings"
	"github.com/rancher/rke/util"
	v3 "github.com/rancher/types/apis/management.cattle.io/v3"
	"github.com/rancher/types/kdm"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OSType int

const (
	Linux OSType = iota
	Windows
)

const (
	APIVersion           = "management.cattle.io/v3"
	RancherVersionDev    = "2.5"
	DataJSONLocation     = "/var/lib/rancher-data/driver-metadata/data.json"
	sendRKELabel         = "io.cattle.rke_store"
	svcOptionLinuxKey    = "service-option-linux-key"
	svcOptionWindowsKey  = "service-option-windows-key"
	rkeSystemImageKind   = "RkeK8sSystemImage"
	rkeServiceOptionKind = "RkeK8sServiceOption"
	rkeAddonKind         = "RkeAddon"
)

var existLabel = map[string]string{sendRKELabel: "false"}

//settings corresponding to keys in setting2.MetadataSettings
var userUpdateSettingMap = map[string]settings.Setting{
	settings.KubernetesVersion.Name:            settings.KubernetesVersion,
	settings.KubernetesVersionsCurrent.Name:    settings.KubernetesVersionsCurrent,
	settings.KubernetesVersionsDeprecated.Name: settings.KubernetesVersionsDeprecated,
}

var rancherUpdateSettingMap = map[string]settings.Setting{
	settings.KubernetesVersion.Name:                 settings.KubernetesVersion,
	settings.KubernetesVersionsCurrent.Name:         settings.KubernetesVersionsCurrent,
	settings.KubernetesVersionsDeprecated.Name:      settings.KubernetesVersionsDeprecated,
	settings.UIKubernetesDefaultVersion.Name:        settings.UIKubernetesDefaultVersion,
	settings.UIKubernetesSupportedVersions.Name:     settings.UIKubernetesSupportedVersions,
	settings.KubernetesVersionToSystemImages.Name:   settings.KubernetesVersionToSystemImages,
	settings.KubernetesVersionToServiceOptions.Name: settings.KubernetesVersionToServiceOptions,
}

func (md *MetadataController) loadDataFromLocal() (kdm.Data, error) {
	if os.Getenv("CATTLE_DEV_MODE") != "" {
		return kdm.Data{}, nil
	}
	logrus.Infof("Retrieve data.json from local path %v", DataJSONLocation)
	data, err := ioutil.ReadFile(DataJSONLocation)
	if err != nil {
		return kdm.Data{}, err
	}
	return kdm.FromData(data)
}

func (md *MetadataController) createOrUpdateMetadata(data kdm.Data) error {
	localData, err := md.loadDataFromLocal()
	if err != nil {
		return err
	}
	if err := md.saveSystemImages(data.K8sVersionRKESystemImages, data.K8sVersionedTemplates,
		data.K8sVersionInfo, data.K8sVersionServiceOptions, data.K8sVersionWindowsServiceOptions, data.RancherDefaultK8sVersions); err != nil {
		return err
	}
	if err := md.saveAllServiceOptions(data.K8sVersionServiceOptions, data.K8sVersionWindowsServiceOptions,
		localData.K8sVersionServiceOptions, localData.K8sVersionWindowsServiceOptions); err != nil {
		return err
	}
	if err := md.saveAddons(data, localData.K8sVersionedTemplates); err != nil {
		return err
	}
	if err := md.saveCisConfigParams(data.CisConfigParams); err != nil {
		return fmt.Errorf("error saving cisDefaultConfigs: %v", err)
	}
	if err := md.saveCisBenchmarkVersions(data.CisBenchmarkVersionInfo); err != nil {
		return fmt.Errorf("error saving cisBechmarkVersions: %v", err)
	}
	return nil
}

func (md *MetadataController) createOrUpdateMetadataFromLocal() error {
	driverData, err := md.loadDataFromLocal()
	if err != nil {
		return err
	}
	if err := md.saveSystemImages(driverData.K8sVersionRKESystemImages, driverData.K8sVersionedTemplates,
		driverData.K8sVersionInfo, driverData.K8sVersionServiceOptions, driverData.K8sVersionWindowsServiceOptions, driverData.RancherDefaultK8sVersions); err != nil {
		return err
	}
	if err := md.saveAllServiceOptions(driverData.K8sVersionServiceOptions, driverData.K8sVersionWindowsServiceOptions, nil, nil); err != nil {
		return err
	}
	if err := md.saveAddons(driverData, nil); err != nil {
		return err
	}
	if err := md.saveCisConfigParams(driverData.CisConfigParams); err != nil {
		return fmt.Errorf("error saving cisDefaultConfigs: %v", err)
	}
	if err := md.saveCisBenchmarkVersions(driverData.CisBenchmarkVersionInfo); err != nil {
		return fmt.Errorf("error saving cisBechmarkVersions: %v", err)
	}
	return nil
}

func (md *MetadataController) saveSystemImages(K8sVersionRKESystemImages map[string]v3.RKESystemImages,
	AddonsData map[string]map[string]string,
	K8sVersionInfo map[string]v3.K8sVersionInfo,
	ServiceOptions map[string]v3.KubernetesServicesOptions,
	ServiceOptionsWindows map[string]v3.KubernetesServicesOptions,
	DefaultK8sVersions map[string]string) error {
	maxVersionForMajorK8sVersion := map[string]string{}
	deprecatedMap := map[string]bool{}
	rancherVersion := GetRancherVersion()
	var maxIgnore []string
	for k8sVersion, systemImages := range K8sVersionRKESystemImages {
		rancherVersionInfo, minorOk := K8sVersionInfo[k8sVersion]
		if minorOk && toIgnoreForAllK8s(rancherVersionInfo, rancherVersion) {
			deprecatedMap[k8sVersion] = true
			continue
		}
		majorVersion := util.GetTagMajorVersion(k8sVersion)
		majorVersionInfo, majorOk := K8sVersionInfo[majorVersion]
		if majorOk && toIgnoreForAllK8s(majorVersionInfo, rancherVersion) {
			deprecatedMap[k8sVersion] = true
			continue
		}
		labelsMap, err := getLabelMap(k8sVersion, AddonsData, ServiceOptions, ServiceOptionsWindows)
		if err != nil {
			return err
		}
		if err := md.createOrUpdateSystemImageCRD(k8sVersion, systemImages, labelsMap); err != nil {
			return err
		}
		if minorOk && toIgnoreForK8sCurrent(rancherVersionInfo, rancherVersion) {
			maxIgnore = append(maxIgnore, k8sVersion)
			continue
		}
		if majorOk && toIgnoreForK8sCurrent(majorVersionInfo, rancherVersion) {
			maxIgnore = append(maxIgnore, k8sVersion)
			continue
		}
		if curr, ok := maxVersionForMajorK8sVersion[majorVersion]; !ok || mVersion.Compare(k8sVersion, curr, ">") {
			maxVersionForMajorK8sVersion[majorVersion] = k8sVersion
		}
	}
	logrus.Debugf("driverMetadata deprecated %v max incompatible versions %v", deprecatedMap, maxIgnore)

	return md.updateSettings(maxVersionForMajorK8sVersion, rancherVersion, ServiceOptions, DefaultK8sVersions, deprecatedMap)
}

func toIgnoreForAllK8s(rancherVersionInfo v3.K8sVersionInfo, rancherVersion string) bool {
	if rancherVersionInfo.DeprecateRancherVersion != "" && mVersion.Compare(rancherVersion, rancherVersionInfo.DeprecateRancherVersion, ">=") {
		return true
	}
	if rancherVersionInfo.MinRancherVersion != "" && mVersion.Compare(rancherVersion, rancherVersionInfo.MinRancherVersion, "<") {
		// only respect min versions, even if max is present - we need to support upgraded clusters
		return true
	}
	return false
}

func toIgnoreForK8sCurrent(majorVersionInfo v3.K8sVersionInfo, rancherVersion string) bool {
	if majorVersionInfo.MaxRancherVersion != "" && mVersion.Compare(rancherVersion, majorVersionInfo.MaxRancherVersion, ">") {
		// include in K8sVersionCurrent only if less then max version
		return true
	}
	return false
}

func (md *MetadataController) saveAllServiceOptions(linuxSvcOptions map[string]v3.KubernetesServicesOptions,
	windowsSvcOptions map[string]v3.KubernetesServicesOptions, localLinuxSvcOptions map[string]v3.KubernetesServicesOptions,
	localWindowsSvcOptions map[string]v3.KubernetesServicesOptions) error {
	// save linux options
	if err := md.saveServiceOptions(linuxSvcOptions, localLinuxSvcOptions, Linux); err != nil {
		return err
	}
	// save windows options
	if err := md.saveServiceOptions(windowsSvcOptions, localWindowsSvcOptions, Windows); err != nil {
		return err
	}
	return nil
}

func (md *MetadataController) saveServiceOptions(k8sVersionServiceOptions map[string]v3.KubernetesServicesOptions,
	localK8sVersionServiceOptions map[string]v3.KubernetesServicesOptions, osType OSType) error {
	rkeDataKeys := getRKEVendorOptions(localK8sVersionServiceOptions)
	for k8sVersion, serviceOptions := range k8sVersionServiceOptions {
		if err := md.createOrUpdateServiceOptionCRD(k8sVersion, serviceOptions, rkeDataKeys, osType); err != nil {
			return err
		}
	}
	return nil
}

func (md *MetadataController) saveAddons(data kdm.Data, localTemplates map[string]map[string]string) error {
	k8sVersionedTemplates := data.K8sVersionedTemplates
	rkeAddonKeys := getRKEVendorData(localTemplates)
	for addon, template := range k8sVersionedTemplates[kdm.TemplateKeys] {
		if err := md.createOrUpdateAddonCRD(addon, template, rkeAddonKeys); err != nil {
			return err
		}
	}
	return nil
}

func (md *MetadataController) saveCisConfigParams(cisConfigParams map[string]v3.CisConfigParams) error {
	for k8sVersion, cisConfigParam := range cisConfigParams {
		logrus.Debugf("saveCisConfigParams k8sversion: %v cisConfigParam: %+v", k8sVersion, cisConfigParam)
		if err := md.createOrUpdateCisConfigCRD(k8sVersion, cisConfigParam); err != nil {
			return fmt.Errorf("error saving cisConfig: %v", err)
		}
	}
	return nil
}

func (md *MetadataController) saveCisBenchmarkVersions(cisBenchmarkVersionInfo map[string]v3.CisBenchmarkVersionInfo) error {
	for benchmarkVersion, info := range cisBenchmarkVersionInfo {
		logrus.Debugf("saveCisBenchmarkVersions bechmarkVersion: %v info: %+v", benchmarkVersion, info)
		if err := md.createOrUpdateCisBenchmarkVersionCRD(benchmarkVersion, info); err != nil {
			return fmt.Errorf("error saving cisBenchmarkVersions: %v", err)
		}
	}
	return nil
}

func (md *MetadataController) createOrUpdateSystemImageCRD(k8sVersion string, systemImages v3.RKESystemImages, pluginsMap map[string]string) error {
	sysImage, err := md.getRKESystemImage(k8sVersion)
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		sysImage = &v3.RKEK8sSystemImage{
			ObjectMeta: metav1.ObjectMeta{
				Name:      k8sVersion,
				Namespace: namespace.GlobalNamespace,
				Labels:    pluginsMap,
			},
			SystemImages: systemImages,
			TypeMeta: metav1.TypeMeta{
				Kind:       rkeSystemImageKind,
				APIVersion: APIVersion,
			},
		}
		if _, err := md.SystemImages.Create(sysImage); err != nil && !errors.IsAlreadyExists(err) {
			return err
		}
		return nil
	}
	dataEqual := reflect.DeepEqual(sysImage.SystemImages, systemImages)
	labelsEqual := reflect.DeepEqual(sysImage.Labels, pluginsMap)
	if dataEqual && labelsEqual {
		return nil
	}
	sysImageCopy := sysImage.DeepCopy()
	if !dataEqual {
		logrus.Debugf("systemImage changed %s", k8sVersion)
		sysImageCopy.SystemImages = systemImages
	}
	if !labelsEqual {
		logrus.Debugf("systemImage labels changed %s old: %v new: %v", k8sVersion, sysImageCopy.Labels, pluginsMap)
		for k, v := range pluginsMap {
			sysImageCopy.Labels[k] = v
		}
	}
	if _, err := md.SystemImages.Update(sysImageCopy); err != nil {
		return err
	}
	return nil
}

func (md *MetadataController) createOrUpdateServiceOptionCRD(k8sVersion string, serviceOptions v3.KubernetesServicesOptions, rkeDataKeys map[string]bool, osType OSType) error {
	svcOption, err := md.getRKEServiceOption(k8sVersion, osType)
	_, exists := rkeDataKeys[k8sVersion]
	name := getVersionNameWithOsType(k8sVersion, osType)
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		svcOption = &v3.RKEK8sServiceOption{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace.GlobalNamespace,
			},
			ServiceOptions: serviceOptions,
			TypeMeta: metav1.TypeMeta{
				Kind:       rkeServiceOptionKind,
				APIVersion: APIVersion,
			},
		}
		if exists {
			svcOption.Labels = existLabel
		}
		if _, err := md.ServiceOptions.Create(svcOption); err != nil && !errors.IsAlreadyExists(err) {
			return err
		}
		return nil
	}
	var svcOptionCopy *v3.RKEK8sServiceOption
	dataEqual := reflect.DeepEqual(svcOption.ServiceOptions, serviceOptions)
	labelsEqual := labelEqual(svcOption.Labels, exists)
	if dataEqual && labelsEqual {
		return nil
	}
	svcOptionCopy = svcOption.DeepCopy()
	if !dataEqual {
		logrus.Debugf("serviceOptions changed %s", name)
		svcOptionCopy.ServiceOptions = serviceOptions
	}
	if !labelsEqual {
		logrus.Debugf("serviceOptions labels changed %s old: %v new: %v", name, svcOptionCopy.Labels, exists)
		updateLabel(svcOptionCopy.Labels, exists)
	}
	if svcOptionCopy != nil {
		if _, err := md.ServiceOptions.Update(svcOptionCopy); err != nil {
			return err
		}
	}
	return nil
}

func (md *MetadataController) createOrUpdateAddonCRD(addonName, template string, rkeAddonKeys map[string]bool) error {
	_, exists := rkeAddonKeys[addonName]
	addon, err := md.getRKEAddon(addonName)
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		addon = &v3.RKEAddon{
			ObjectMeta: metav1.ObjectMeta{
				Name:      addonName,
				Namespace: namespace.GlobalNamespace,
			},
			Template: template,
			TypeMeta: metav1.TypeMeta{
				Kind:       rkeAddonKind,
				APIVersion: APIVersion,
			},
		}
		if exists {
			addon.Labels = existLabel
		}
		if _, err := md.Addons.Create(addon); err != nil && !errors.IsAlreadyExists(err) {
			return err
		}
		return nil
	}
	var addonCopy *v3.RKEAddon
	dataEqual := reflect.DeepEqual(addon.Template, template)
	labelsEqual := labelEqual(addon.Labels, exists)
	if dataEqual && labelsEqual {
		return nil
	}
	addonCopy = addon.DeepCopy()
	if !dataEqual {
		logrus.Debugf("addonTemplate changed %s", addonName)
		addonCopy.Template = template
	}
	if !labelsEqual {
		logrus.Debugf("addonTemplate labels changed %s old: %v new: %v", addonName, addonCopy.Labels, exists)
		updateLabel(addonCopy.Labels, exists)
	}
	if addonCopy != nil {
		if _, err := md.Addons.Update(addonCopy); err != nil {
			return err
		}
	}
	return nil
}

func (md *MetadataController) createOrUpdateCisConfigCRD(
	k8sVersion string,
	cisConfigParams v3.CisConfigParams,
) error {
	cisConfig, err := md.CisConfigLister.Get(namespace.GlobalNamespace, k8sVersion)
	if err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("error finding cisConfig for k8sVersion: %v", k8sVersion)
		}
		cisConfig = &v3.CisConfig{
			TypeMeta: metav1.TypeMeta{
				Kind:       v3.CisConfigGroupVersionKind.Kind,
				APIVersion: v3.SchemeGroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      k8sVersion,
				Namespace: namespace.GlobalNamespace,
			},
			Params: cisConfigParams,
		}
		logrus.Debugf("driverMetadata: creating cisConfig CRD for k8sVersion: %v",
			k8sVersion)
		if _, err := md.CisConfig.Create(cisConfig); err != nil && !errors.IsAlreadyExists(err) {
			return err
		}
		return nil
	}
	if reflect.DeepEqual(cisConfig.Params, cisConfigParams) {
		return nil
	}
	logrus.Debugf("driverMetadata: cisConfigParams changed for k8sVersion: %v old: %v new: %v",
		k8sVersion, cisConfig.Params, cisConfigParams)
	cisConfigCopy := cisConfig.DeepCopy()
	cisConfigCopy.Params = cisConfigParams
	if _, err := md.CisConfig.Update(cisConfigCopy); err != nil {
		return err
	}
	return nil
}

func (md *MetadataController) createOrUpdateCisBenchmarkVersionCRD(
	benchmarkVersion string,
	cisBenchmarkVersionInfo v3.CisBenchmarkVersionInfo,
) error {
	cisBenchmarkVersion, err := md.CisBenchmarkVersionLister.Get(namespace.GlobalNamespace, benchmarkVersion)
	if err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("error finding cisBenchmarkVerionInfo for benchmarkVersion: %v", benchmarkVersion)
		}
		cisBenchmarkVersion = &v3.CisBenchmarkVersion{
			TypeMeta: metav1.TypeMeta{
				Kind:       v3.CisBenchmarkVersionGroupVersionKind.Kind,
				APIVersion: v3.SchemeGroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      benchmarkVersion,
				Namespace: namespace.GlobalNamespace,
			},
			Info: cisBenchmarkVersionInfo,
		}
		logrus.Debugf("driverMetadata: creating cisBenchmarkVerion CRD for benchmarkVersion: %v",
			benchmarkVersion)
		if _, err := md.CisBenchmarkVersion.Create(cisBenchmarkVersion); err != nil && !errors.IsAlreadyExists(err) {
			return err
		}
		return nil
	}
	if reflect.DeepEqual(cisBenchmarkVersion.Info, cisBenchmarkVersionInfo) {
		return nil
	}
	logrus.Debugf("driverMetadata: cisBenchmarkVerionInfochanged for benchmarkVersion: %v old: %v new: %v",
		benchmarkVersion, cisBenchmarkVersion.Info, cisBenchmarkVersionInfo)
	cisBenchmarkVersionCopy := cisBenchmarkVersion.DeepCopy()
	cisBenchmarkVersionCopy.Info = cisBenchmarkVersionInfo
	if _, err := md.CisBenchmarkVersion.Update(cisBenchmarkVersionCopy); err != nil {
		return err
	}
	return nil
}

func getLabelMap(k8sVersion string, data map[string]map[string]string,
	svcOption map[string]v3.KubernetesServicesOptions, svcOptionWindows map[string]v3.KubernetesServicesOptions) (map[string]string, error) {
	toMatch, err := semver.Make(k8sVersion[1:])
	if err != nil {
		return nil, fmt.Errorf("k8sVersion not sem-ver %s %v", k8sVersion, err)
	}
	labelMap := map[string]string{"cattle.io/creator": "norman"}
	for addon, addonData := range data {
		if addon == kdm.TemplateKeys {
			continue
		}
		found := false
		for k8sRange, key := range addonData {
			testRange, err := semver.ParseRange(k8sRange)
			if err != nil {
				logrus.Errorf("getPluginData: range for %s not sem-ver %v %v", addon, testRange, err)
				continue
			}
			if testRange(toMatch) {
				labelMap[addon] = key
				found = true
				break
			}
		}
		if !found {
			logrus.Debugf("getPluginData: no template found for k8sVersion %s plugin %s", k8sVersion, addon)
		}
	}
	// store service options
	majorKey := util.GetTagMajorVersion(k8sVersion)
	if _, ok := svcOption[k8sVersion]; ok {
		labelMap[svcOptionLinuxKey] = getVersionNameWithOsType(k8sVersion, Linux)
	} else if _, ok := svcOption[majorKey]; ok {
		labelMap[svcOptionLinuxKey] = getVersionNameWithOsType(majorKey, Linux)
	}

	if _, ok := svcOptionWindows[k8sVersion]; ok {
		labelMap[svcOptionWindowsKey] = getVersionNameWithOsType(k8sVersion, Windows)
	} else if _, ok := svcOptionWindows[majorKey]; ok {
		labelMap[svcOptionWindowsKey] = getVersionNameWithOsType(majorKey, Windows)
	}

	return labelMap, nil
}

func getRKEVendorData(templates map[string]map[string]string) map[string]bool {
	keys := map[string]bool{}
	if templates == nil {
		return keys
	}
	templateData, ok := templates[kdm.TemplateKeys]
	if !ok {
		return keys
	}
	for templateKey := range templateData {
		keys[templateKey] = true
	}
	return keys
}

func getRKEVendorOptions(options map[string]v3.KubernetesServicesOptions) map[string]bool {
	keys := map[string]bool{}
	for k8sVersion := range options {
		keys[k8sVersion] = true
	}
	return keys
}

func (md *MetadataController) getRKEAddon(name string) (*v3.RKEAddon, error) {
	return md.AddonsLister.Get(namespace.GlobalNamespace, name)
}

func (md *MetadataController) getRKEServiceOption(k8sVersion string, osType OSType) (*v3.RKEK8sServiceOption, error) {
	return md.ServiceOptionsLister.Get(namespace.GlobalNamespace, getVersionNameWithOsType(k8sVersion, osType))
}

func (md *MetadataController) getRKESystemImage(k8sVersion string) (*v3.RKEK8sSystemImage, error) {
	return md.SystemImagesLister.Get(namespace.GlobalNamespace, k8sVersion)
}

func getVersionNameWithOsType(str string, osType OSType) string {
	if osType == Windows {
		return getWindowsName(str)
	}
	return str
}

func getWindowsName(str string) string {
	return fmt.Sprintf("w%s", str)
}

func (md *MetadataController) updateSettings(maxVersionForMajorK8sVersion map[string]string, rancherVersion string,
	K8sVersionServiceOptions map[string]v3.KubernetesServicesOptions, DefaultK8sVersions map[string]string,
	deprecated map[string]bool) error {

	userSettings, userUpdated, err := md.getUserSettings()
	if err != nil {
		return err
	}

	updateSettings, err := toUpdate(maxVersionForMajorK8sVersion, deprecated, DefaultK8sVersions, rancherVersion, K8sVersionServiceOptions)
	if err != nil {
		return err
	}

	if !userUpdated {
		if err := md.updateSettingFromFields(updateSettings, map[string]string{}); err != nil {
			return err
		}
	} else {
		userMaxVersionForMajorK8sVersion, userDeprecated, err := getUserSettings(userSettings, DefaultK8sVersions)
		if err != nil {
			return err
		}

		if len(userMaxVersionForMajorK8sVersion) == 0 {
			userMaxVersionForMajorK8sVersion = maxVersionForMajorK8sVersion
		}

		if len(userDeprecated) == 0 {
			userDeprecated = deprecated
		}

		userUpdateSettings, err := toUpdate(userMaxVersionForMajorK8sVersion, userDeprecated, DefaultK8sVersions, rancherVersion, K8sVersionServiceOptions)
		if err != nil {
			return err
		}

		if err := md.updateSettingFromFields(userUpdateSettings, userSettings); err != nil {
			return err
		}
	}

	return nil
}

func (md *MetadataController) getUserSettings() (map[string]string, bool, error) {
	userSettings := map[string]string{}
	get := func(key string) string {
		if setting, ok := userUpdateSettingMap[key]; ok {
			return setting.Get()
		}
		return ""
	}
	for key := range userUpdateSettingMap {
		setting, err := md.SettingLister.Get("", key)
		if err != nil {
			if !errors.IsNotFound(err) {
				return nil, false, fmt.Errorf("driverMetadata: error getting setting %s: %v", key, err)
			}
			setting, err = md.Settings.Get(key, metav1.GetOptions{})
			if err != nil {
				return nil, false, fmt.Errorf("driverMetadata: error getting setting %s: %v", key, err)
			}
		}
		if val, ok := setting.Labels[setting2.UserUpdateLabel]; ok && convert.ToString(val) == "true" {
			userSettings[key] = get(key)
		}
	}
	logrus.Debugf("driverMetadata: userSettings %v", userSettings)
	if len(userSettings) > 0 {
		return userSettings, true, nil
	}
	return userSettings, false, nil
}

func toUpdate(maxVersionForMajorK8sVersion map[string]string, deprecated map[string]bool,
	defaultK8sVersions map[string]string, rancherVersion string, k8sVersionServiceOptions map[string]v3.KubernetesServicesOptions) (map[string]string, error) {

	var k8sVersionsCurrent []string
	var maxVersions []string
	for k, v := range maxVersionForMajorK8sVersion {
		if !deprecated[k] {
			k8sVersionsCurrent = append(k8sVersionsCurrent, v)
			maxVersions = append(maxVersions, k)
		}
	}
	if len(maxVersions) == 0 {
		return nil, fmt.Errorf("driverMetadata: no max version %v", maxVersionForMajorK8sVersion)
	}
	sort.Strings(k8sVersionsCurrent)
	sort.Strings(maxVersions)

	defaultK8sVersion, err := getDefaultK8sVersion(defaultK8sVersions, k8sVersionsCurrent, rancherVersion)
	if err != nil {
		return nil, err
	}

	k8sVersionRKESystemImages := map[string]interface{}{}
	k8sVersionSvcOptions := map[string]v3.KubernetesServicesOptions{}

	for majorVersion, k8sVersion := range maxVersionForMajorK8sVersion {
		if !deprecated[k8sVersion] {
			k8sVersionRKESystemImages[k8sVersion] = nil
			k8sVersionSvcOptions[k8sVersion] = k8sVersionServiceOptions[majorVersion]
		}
	}

	k8sCurrRKEdata, err := marshal(k8sVersionRKESystemImages)
	if err != nil {
		return nil, err
	}

	k8sSvcOptionData, err := marshal(k8sVersionSvcOptions)
	if err != nil {
		return nil, err
	}

	deprecatedData, err := marshal(deprecated)
	if err != nil {
		return nil, err
	}

	minVersion := maxVersions[0]
	maxVersion := util.GetTagMajorVersion(defaultK8sVersion)
	uiSupported := fmt.Sprintf(">=%s.x <=%s.x", minVersion, maxVersion)
	uiDefaultRange := fmt.Sprintf("<=%s.x", maxVersion)

	return map[string]string{
		settings.KubernetesVersionsCurrent.Name:         strings.Join(k8sVersionsCurrent, ","),
		settings.KubernetesVersion.Name:                 defaultK8sVersion,
		settings.KubernetesVersionsDeprecated.Name:      deprecatedData,
		settings.UIKubernetesDefaultVersion.Name:        uiDefaultRange,
		settings.UIKubernetesSupportedVersions.Name:     uiSupported,
		settings.KubernetesVersionToSystemImages.Name:   k8sCurrRKEdata,
		settings.KubernetesVersionToServiceOptions.Name: k8sSvcOptionData,
	}, nil
}

func (md *MetadataController) updateSettingFromFields(updateField map[string]string, skip map[string]string) error {
	for key, setting := range rancherUpdateSettingMap {
		if _, ok := skip[key]; ok {
			continue
		}
		if _, ok := updateField[key]; !ok {
			return fmt.Errorf("driverMetadata: updated value not present for setting %s", key)
		}
		oldVal := setting.Get()
		newVal := updateField[key]
		if oldVal != newVal {
			if err := setting.Set(newVal); err != nil {
				return err
			}
		}
	}
	return nil
}

func getUserSettings(userSettings map[string]string, defaultK8sVersions map[string]string) (map[string]string, map[string]bool, error) {
	userMaxVersionForMajorK8sVersion := map[string]string{}
	if val, ok := userSettings[settings.KubernetesVersionsCurrent.Name]; ok {
		versions := strings.Split(val, ",")
		for _, version := range versions {
			userMaxVersionForMajorK8sVersion[util.GetTagMajorVersion(version)] = version
		}
	}

	userDeprecated := map[string]bool{}
	if val, ok := userSettings[settings.KubernetesVersionsDeprecated.Name]; ok {
		deprecatedVersions := make(map[string]bool)
		if val != "" {
			if err := json.Unmarshal([]byte(val), &deprecatedVersions); err != nil {
				return nil, nil, err
			}
		}
		for key, val := range deprecatedVersions {
			userDeprecated[key] = val
		}
	}

	if val, ok := userSettings[settings.KubernetesVersion.Name]; ok {
		defaultK8sVersions["user"] = val
	}

	return userMaxVersionForMajorK8sVersion, userDeprecated, nil
}

func getDefaultK8sVersion(rancherDefaultK8sVersions map[string]string, k8sCurrVersions []string, rancherVersion string) (string, error) {
	defaultK8sVersion, ok := rancherDefaultK8sVersions["user"]
	if defaultK8sVersion != "" {
		found := false
		for _, k8sVersion := range k8sCurrVersions {
			if k8sVersion == defaultK8sVersion {
				found = true
				break
			}
		}
		if !found {
			return "", fmt.Errorf("driverMetadata: unable to find default k8s version in current k8s %s %v", defaultK8sVersion, k8sCurrVersions)
		}
		return defaultK8sVersion, nil
	}

	defaultK8sVersionRange, ok := rancherDefaultK8sVersions[rancherVersion]
	if !ok || defaultK8sVersionRange == "" {
		defaultK8sVersionRange = rancherDefaultK8sVersions["default"]
	}
	// get matching default k8s from k8s curr
	toMatch := util.GetTagMajorVersion(defaultK8sVersionRange)

	for _, k8sCurr := range k8sCurrVersions {
		toTest := util.GetTagMajorVersion(k8sCurr)
		if toTest == toMatch {
			defaultK8sVersion = k8sCurr
			break
		}
	}
	if defaultK8sVersion == "" {
		return "", fmt.Errorf("driverMetadata: unable to find default k8s version in current k8s %s %v", defaultK8sVersionRange, k8sCurrVersions)
	}
	return defaultK8sVersion, nil
}

func marshal(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func labelEqual(labels map[string]string, exists bool) bool {
	toSendValue := "true"
	if exists {
		toSendValue = "false"
	}
	if _, ok := labels[sendRKELabel]; !ok {
		return toSendValue == "true"
	}
	return toSendValue == labels[sendRKELabel]
}

func updateLabel(labels map[string]string, exists bool) {
	if exists {
		labels[sendRKELabel] = "false"
	} else {
		delete(labels, sendRKELabel)
	}

}
