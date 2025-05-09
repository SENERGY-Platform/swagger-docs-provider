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

package api

import (
	"errors"
	"github.com/SENERGY-Platform/swagger-docs-provider/pkg/models"
	"net/http"
)

func GetStatusCode(err error) int {
	var nfe *models.NotFoundError
	if errors.As(err, &nfe) {
		return http.StatusNotFound
	}
	var iie *models.InvalidInputError
	if errors.As(err, &iie) {
		return http.StatusBadRequest
	}
	var ie *models.InternalError
	if errors.As(err, &ie) {
		return http.StatusInternalServerError
	}
	var rbe *models.ResourceBusyError
	if errors.As(err, &rbe) {
		return http.StatusConflict
	}
	return 0
}
