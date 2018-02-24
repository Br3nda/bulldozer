// Copyright 2017 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package middleware

import (
	"github.com/google/go-github/github"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func CtxMiddleware(logLevel logrus.Level) func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			deliveryID := github.DeliveryID(c.Request())

			logger := logrus.New()
			logger.SetLevel(logLevel)
			logger.Formatter = &logrus.JSONFormatter{}

			if deliveryID != "" {
				c.Set("log", logger.WithField("deliveryID", deliveryID))
			} else {
				c.Set("log", logger.WithFields(nil))
			}

			return next(c)
		}
	}
}
