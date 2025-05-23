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

package service

import (
	"context"
	"encoding/json"
	"github.com/SENERGY-Platform/swagger-docs-provider/pkg/models"
)

type swaggerService interface {
	SwaggerGetDocs(ctx context.Context, userToken string, userRoles []string) ([]map[string]json.RawMessage, error)
	SwaggerListStorage(ctx context.Context) ([]models.SwaggerItem, error)
	SwaggerRefreshDocs(ctx context.Context) error
}

type asyncapiService interface {
	AsyncapiGetDocs(ctx context.Context) ([]json.RawMessage, error)
	AsyncapiPutDoc(ctx context.Context, id string, data []byte) error
	AsyncapiDeleteDoc(ctx context.Context, id string) error
	AsyncapiListStorage(ctx context.Context) ([]models.AsyncapiItem, error)
}
