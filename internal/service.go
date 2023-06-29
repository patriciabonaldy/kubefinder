package internal

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus" //nolint:depguard

	v1 "k8s.io/api/core/v1"                       //nolint:depguard
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" //nolint:depguard
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd" //nolint:depguard
	"k8s.io/client-go/util/homedir"
)

type service struct {
	client *kubernetes.Clientset
	log    *logrus.Logger
}

type Service interface {
	FindInConfigMaps(word string) (interface{}, error)
}

var (
	configMap = make(map[string]ConfigMap, 0)
)

func NewService(log *logrus.Logger) (Service, error) {
	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeConfig file")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "absolute path to the kubeConfig file")
	}
	flag.Parse()

	// use the current context in kubeConfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		return nil, err
	}

	// create the client
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	s := &service{client: client, log: log}
	s.initializeMap()

	return s, nil
}

func (s *service) FindInConfigMaps(word string) (interface{}, error) {
	result := make([]ConfigMap, 0)
	var validID = regexp.MustCompile(fmt.Sprintf(`%s.*?\\n`, word))
	for _, config := range configMap {
		matches := validID.FindAllString(config.Value, -1)
		if len(matches) > 0 {
			for _, m := range matches {
				match := ConfigMap{
					Name:      config.Name,
					Value:     normalizeString(m),
					NameSpace: config.NameSpace,
				}

				result = append(result, match)
			}
		}
	}

	return result, nil
}

func normalizeString(m string) string {
	return strings.ReplaceAll(strings.ReplaceAll(m, "\\n", ""), "\\\"", "")
}

func (s *service) initializeMap() {
	ctx := context.Background()
	nameSpaces, err := s.client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, ns := range nameSpaces.Items {
		var configMaps *v1.ConfigMapList

		configMaps, err = s.client.CoreV1().ConfigMaps(ns.Name).List(ctx, metav1.ListOptions{})
		if err != nil {
			s.log.WithError(err).WithField("namespace", ns.Name).Warn("failed getting configmaps")
		}

		fillConfigMaps(ns.Name, configMaps)
	}
}

func fillConfigMaps(nameSpace string, configMaps *v1.ConfigMapList) {
	for _, cm := range configMaps.Items {
		data, _ := cm.Marshal()
		configMap[cm.Name] = ConfigMap{
			Name:      cm.Name,
			Value:     string(data),
			NameSpace: nameSpace,
		}
	}
}
