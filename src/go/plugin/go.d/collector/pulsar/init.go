// SPDX-License-Identifier: GPL-3.0-or-later

package pulsar

import (
	"errors"

	"github.com/netdata/netdata/go/plugins/pkg/matcher"
	"github.com/netdata/netdata/go/plugins/plugin/go.d/pkg/prometheus"
	"github.com/netdata/netdata/go/plugins/plugin/go.d/pkg/web"
)

func (c *Collector) validateConfig() error {
	if c.URL == "" {
		return errors.New("url not set")
	}
	return nil
}

func (c *Collector) initPrometheusClient() (prometheus.Prometheus, error) {
	client, err := web.NewHTTPClient(c.ClientConfig)
	if err != nil {
		return nil, err
	}

	return prometheus.New(client, c.RequestConfig), nil
}

func (c *Collector) initTopicFilerMatcher() (matcher.Matcher, error) {
	if c.TopicFilter.Empty() {
		return matcher.FALSE(), nil
	}
	return c.TopicFilter.Parse()
}