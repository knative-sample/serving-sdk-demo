package manager

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/serving/pkg/apis/serving/v1"
)

func (m *Manager) CreateKsvc() error {
	logger := m.logger
	logger.Infof("create KSVC Manager is run")
	ksvc := &v1.Service{

		ObjectMeta: metav1.ObjectMeta{
			//Name:      npName,
			GenerateName: "manager-create-ksvc-",
			Labels: map[string]string{
				"foo": "bar",
			},
			Annotations: map[string]string{
				"foo": "bar",
			},
		},
		Spec: v1.ServiceSpec{
			ConfigurationSpec: v1.ConfigurationSpec{
				Template: v1.RevisionTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"revision-foo": "bar",
						},
						Annotations: map[string]string{
							"revision-foo": "bar",
						},
					},
					Spec: v1.RevisionSpec{
						PodSpec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									ImagePullPolicy: corev1.PullAlways,
									Image:           "registry.cn-hangzhou.aliyuncs.com/knative-sample/helloworld-go:160e4dc8",
									Env: []corev1.EnvVar{{
										Name:  "TARGET",
										Value: "serving sdk demo, manager <.|.>",
									}},
								},
							},
						},
					},
				},
			},

			RouteSpec: v1.RouteSpec{
				// set traffic
				Traffic: []v1.TrafficTarget{},
			},
		},
	}
	_, err := m.servingclient.ServingV1().Services("default").Create(ksvc)
	if err != nil {
		logger.Errorf("create ksvc error:%s", err.Error())
		return err
	}

	logger.Infof("create ksvc success!!")
	return nil
}
