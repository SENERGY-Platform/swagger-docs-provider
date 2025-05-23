/*
 * Copyright 2025 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package discovery_hdl

import (
	"context"
	"fmt"
	"github.com/SENERGY-Platform/swagger-docs-provider/pkg/components/kong_clt"
	"github.com/SENERGY-Platform/swagger-docs-provider/pkg/models"
	"github.com/SENERGY-Platform/swagger-docs-provider/pkg/util/slog_attr"
	"time"
)

type Handler struct {
	kongClient    kong_clt.ClientItf
	timeout       time.Duration
	hostBlacklist map[string]struct{}
}

func New(kongClient kong_clt.ClientItf, timeout time.Duration, hostBlacklist []string) *Handler {
	blackList := make(map[string]struct{})
	for _, host := range hostBlacklist {
		blackList[host] = struct{}{}
	}
	return &Handler{
		kongClient:    kongClient,
		timeout:       timeout,
		hostBlacklist: blackList,
	}
}

func (h *Handler) GetServices(ctx context.Context) (map[string]models.Service, error) {
	ctxWt, cf := context.WithTimeout(ctx, h.timeout)
	defer cf()
	kRoutes, err := h.kongClient.GetRoutes(ctxWt)
	if err != nil {
		return nil, models.NewInternalError(err)
	}
	ctxWt2, cf2 := context.WithTimeout(ctx, h.timeout)
	defer cf2()
	kServices, err := h.kongClient.GetServices(ctxWt2)
	if err != nil {
		return nil, models.NewInternalError(err)
	}
	kSrvMap := getKongSrvMap(kServices)
	services := make(map[string]models.Service)
	for _, kRoute := range kRoutes {
		if len(kRoute.Paths) == 0 {
			continue
		}
		kService, ok := kSrvMap[kRoute.Service.ID]
		if !ok {
			continue
		}
		if _, ok = h.hostBlacklist[kService.Host]; ok {
			continue
		}
		id := fmt.Sprintf("%s%d", kService.Host, kService.Port)
		service, ok := services[id]
		if !ok {
			service.ID = id
			service.Host = kService.Host
			service.Port = kService.Port
			service.Protocol = kService.Protocol
		}
		service.ExtPaths = append(service.ExtPaths, kRoute.Paths...)
		services[id] = service
	}
	for _, service := range services {
		logger.Debug("found service", slog_attr.HostKey, service.Host, slog_attr.PortKey, service.Port, slog_attr.ExternalPathsKey, service.ExtPaths)
	}
	return services, nil
}

func getKongSrvMap(kServices []kong_clt.Service) map[string]kong_clt.Service {
	srvMap := make(map[string]kong_clt.Service)
	for _, kService := range kServices {
		srvMap[kService.ID] = kService
	}
	return srvMap
}
