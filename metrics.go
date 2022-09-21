package utils

import "github.com/prometheus/client_golang/prometheus"

var (
	MetricMsgsIn = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "input_messages_per_topic",
			Help: "Messages received per topic.",
		},
		[]string{"topic"},
	)

	MetricMsgsOut = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ouput_msgs_per_topic",
			Help: "Messages send per topic and type.",
		},
		[]string{"topic", "msg_type"},
	)
)

func init() {
	prometheus.MustRegister(MetricMsgsIn)
	prometheus.MustRegister(MetricMsgsOut)
}
