package stats

import (
	"github.com/prometheus/client_golang/prometheus"
)

var _ Handler = (*Prometheus)(nil)

type (
	Prometheus struct {
		PrometheusOptions
		cacheName  string
		counterVec *prometheus.CounterVec
	}

	PrometheusOptions struct {
		namespace string
		subsystem string
		name      string
		help      string
	}

	PrometheusOption func(*PrometheusOptions)
)

func WithNamespace(namespace string) PrometheusOption {
	return func(o *PrometheusOptions) {
		o.namespace = namespace
	}
}

func WithSubsystem(subsystem string) PrometheusOption {
	return func(o *PrometheusOptions) {
		o.subsystem = subsystem
	}
}

func WithName(name string) PrometheusOption {
	return func(o *PrometheusOptions) {
		o.name = name
	}
}

func WithHelp(help string) PrometheusOption {
	return func(o *PrometheusOptions) {
		o.help = help
	}
}

// NewPrometheus is
func NewPrometheus(cacheName string, opts ...PrometheusOption) *Prometheus {
	o := PrometheusOptions{
		name: "cache_handle_total",
	}
	for _, opt := range opts {
		opt(&o)
	}

	vec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: o.namespace,
			Subsystem: o.subsystem,
			Name:      o.name,
			Help:      o.help,
		}, []string{"cache_name", "cache_type", "method", "err"},
	)
	prometheus.MustRegister(vec)

	return &Prometheus{
		PrometheusOptions: o,
		cacheName:         cacheName,
		counterVec:        vec,
	}
}

func (s *Prometheus) IncrHit() {
	s.counterVec.WithLabelValues(s.cacheName, "total", "hit", "").Inc()
}

func (s *Prometheus) IncrMiss() {
	s.counterVec.WithLabelValues(s.cacheName, "total", "miss", "").Inc()
}

func (s *Prometheus) IncrLocalHit() {
	s.counterVec.WithLabelValues(s.cacheName, "local", "hit", "").Inc()
}

func (s *Prometheus) IncrLocalMiss() {
	s.counterVec.WithLabelValues(s.cacheName, "local", "miss", "").Inc()
}

func (s *Prometheus) IncrRemoteHit() {
	s.counterVec.WithLabelValues(s.cacheName, "remote", "hit", "").Inc()
}

func (s *Prometheus) IncrRemoteMiss() {
	s.counterVec.WithLabelValues(s.cacheName, "remote", "miss", "").Inc()
}

func (s *Prometheus) IncrQuery() {
	s.counterVec.WithLabelValues(s.cacheName, "query", "query", "").Inc()
}

func (s *Prometheus) IncrQueryFail(err error) {
	s.counterVec.WithLabelValues(s.cacheName, "query", "queryFail", err.Error()).Inc()
}
