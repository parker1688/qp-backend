// Copyright 2019 HenryYee.
//
// Licensed under the AGPL, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.gnu.org/licenses/agpl-3.0.en.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package global

import (
	"bootpkg/common/conf"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	JSON          = jsoniter.ConfigCompatibleWithStandardLibrary
	CONFIG        *conf.Config
	G_DB          *gorm.DB
	G_DB_SHARDING *gorm.DB
	G_REDIS       redis.UniversalClient
	G_LOG         *zap.SugaredLogger
	LANG          ut.Translator

	VALIDATE *validator.Validate
)
