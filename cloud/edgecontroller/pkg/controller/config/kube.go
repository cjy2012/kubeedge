package config

import (
	"time"

	"github.com/kubeedge/kubeedge/cloud/edgecontroller/pkg/controller/constants"
	"github.com/kubeedge/kubeedge/common/beehive/pkg/common/config"
	"github.com/kubeedge/kubeedge/common/beehive/pkg/common/log"
)

// KubeMaster is the url of edge master(kube api server)
var KubeMaster string

// KubeConfig is the config used connect to edge master
var KubeConfig string

// KubeNamespace is the namespace to watch(default is NamespaceAll)
var KubeNamespace string

// KubeContentType is the content type communicate with edge master(default is "application/vnd.kubernetes.protobuf")
var KubeContentType string

// KubeQPS is the QPS communicate with edge master(default is 1024)
var KubeQPS float32

// KubeBurst default is 10
var KubeBurst int

// KubeUpdateNodeFrequency is the time duration for update node status(default is 20s)
var KubeUpdateNodeFrequency time.Duration

func init() {
	if km, err := config.CONFIG.GetValue("controller.kube.master").ToString(); err != nil {
		log.LOGGER.Errorf("kube master not set")
	} else {
		KubeMaster = km
	}
	log.LOGGER.Infof("kube master: %s", KubeMaster)

	if kc, err := config.CONFIG.GetValue("controller.kube.kubeconfig").ToString(); err != nil {
		log.LOGGER.Errorf("kube config not set")
	} else {
		KubeConfig = kc
	}
	log.LOGGER.Infof("kube config: %s", KubeConfig)

	if kn, err := config.CONFIG.GetValue("controller.kube.namespace").ToString(); err != nil {
		KubeNamespace = constants.DefaultKubeNamespace
	} else {
		KubeNamespace = kn
	}
	log.LOGGER.Infof("kube namespace: %s", KubeNamespace)

	if kct, err := config.CONFIG.GetValue("controller.kube.content_type").ToString(); err != nil {
		KubeContentType = constants.DefaultKubeContentType
	} else {
		KubeContentType = kct
	}
	log.LOGGER.Infof("kube content type: %s", KubeContentType)

	if kqps, err := config.CONFIG.GetValue("controller.kube.qps").ToFloat64(); err != nil {
		KubeQPS = constants.DefaultKubeQPS
	} else {
		KubeQPS = float32(kqps)
	}
	log.LOGGER.Infof("kube QPS: %f", KubeQPS)

	if kb, err := config.CONFIG.GetValue("controller.kube.burst").ToInt(); err != nil {
		KubeBurst = constants.DefaultKubeBurst
	} else {
		KubeBurst = kb
	}
	log.LOGGER.Infof("kube burst: %d", KubeBurst)
	
	if kuf, err := config.CONFIG.GetValue("controller.kube.node_update_frequency").ToInt64(); err != nil {
		KubeUpdateNodeFrequency = constants.DefaultKubeUpdateNodeFrequency * time.Second
	} else {
		KubeUpdateNodeFrequency = time.Duration(kuf) * time.Second
	}
	log.LOGGER.Infof("kube update frequency: %v", KubeUpdateNodeFrequency)	
}
