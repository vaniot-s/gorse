// Copyright 2020 gorse Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"encoding/json"
	"github.com/zhenghaoz/gorse/base"
	"go.uber.org/zap"
	"reflect"
)

/* ParamName */

// ParamName is the type of hyper-parameter names.
// 超参数是在开始学习过程之前设置值的参数，而不是通过训练得到的参数数据。
type ParamName string

// Predefined hyper-parameter names
// 预定义的 hyper-parameter名称
const (
	Lr          ParamName = "Lr"          // learning rate
	Reg         ParamName = "Reg"         // regularization strength
	NEpochs     ParamName = "NEpochs"     // number of epochs
	NFactors    ParamName = "NFactors"    // number of factors
	RandomState ParamName = "RandomState" // random state (seed)
	InitMean    ParamName = "InitMean"    // mean of gaussian initial parameter 均值高斯的初始参数
	InitStdDev  ParamName = "InitStdDev"  // standard deviation of gaussian initial parameter 高斯初始参数的标准偏差
	Alpha       ParamName = "Alpha"       // weight for negative samples in ALS
	Similarity  ParamName = "Similarity"
	UseFeature  ParamName = "UseFeature"
)

const (
	SimilarityCosine = "Cosine"
	SimilarityDot    = "Dot"
)

// Params stores hyper-parameters for an model. It is a map between strings
// (names) and interface{}s (values). For example, hyper-parameters for SVD
// is given by:
//  base.Params{
//		base.Lr:       0.007,
//		base.NEpochs:  100,
//		base.NFactors: 80,
//		base.Reg:      0.1,
//	}
type Params map[ParamName]interface{}

// Copy hyper-parameters.
func (parameters Params) Copy() Params {
	newParams := make(Params)
	for k, v := range parameters {
		newParams[k] = v
	}
	return newParams
}

// GetBool gets a boolean parameter by name. Returns _default if not exists or type doesn't match.
func (parameters Params) GetBool(name ParamName, _default bool) bool {
	if val, exist := parameters[name]; exist {
		switch val := val.(type) {
		case bool:
			return val
		default:
			base.Logger().Error("type mismatch",
				zap.String("param_name", string(name)),
				zap.String("actual_type", reflect.TypeOf(name).Name()))
		}
	}
	return _default
}

// GetInt gets a integer parameter by name. Returns _default if not exists or type doesn't match.
func (parameters Params) GetInt(name ParamName, _default int) int {
	if val, exist := parameters[name]; exist {
		switch val := val.(type) {
		case int:
			return val
		default:
			base.Logger().Error("type mismatch",
				zap.String("param_name", string(name)),
				zap.String("actual_type", reflect.TypeOf(name).Name()))
		}
	}
	return _default
}

// GetInt64 gets a int64 parameter by name. Returns _default if not exists or type doesn't match. The
// type will be converted if given int.
func (parameters Params) GetInt64(name ParamName, _default int64) int64 {
	if val, exist := parameters[name]; exist {
		switch val := val.(type) {
		case int64:
			return val
		case int:
			return int64(val)
		default:
			base.Logger().Error("type mismatch",
				zap.String("param_name", string(name)),
				zap.String("actual_type", reflect.TypeOf(name).Name()))
		}
	}
	return _default
}

func (parameters Params) GetFloat32(name ParamName, _default float32) float32 {
	if val, exist := parameters[name]; exist {
		switch val := val.(type) {
		case float32:
			return val
		case float64:
			return float32(val)
		case int:
			return float32(val)
		default:
			base.Logger().Error("type mismatch",
				zap.String("param_name", string(name)),
				zap.String("actual_type", reflect.TypeOf(name).Name()))
		}
	}
	return _default
}

// GetString gets a string parameter
func (parameters Params) GetString(name ParamName, _default string) string {
	if val, exist := parameters[name]; exist {
		return val.(string)
	}
	return _default
}

func (parameters Params) Overwrite(params Params) Params {
	merged := make(Params)
	for k, v := range parameters {
		merged[k] = v
	}
	for k, v := range params {
		merged[k] = v
	}
	return merged
}

func (parameters Params) ToString() string {
	b, err := json.Marshal(parameters)
	if err != nil {
		base.Logger().Fatal("failed to convert to string", zap.Error(err))
	}
	return string(b)
}

// ParamsGrid contains candidate for grid search.
type ParamsGrid map[ParamName][]interface{}

func (grid ParamsGrid) Len() int {
	return len(grid)
}

func (grid ParamsGrid) NumCombinations() int {
	count := 1
	for _, values := range grid {
		count *= len(values)
	}
	return count
}

func (grid ParamsGrid) Fill(_default ParamsGrid) {
	for param, values := range _default {
		if _, exist := grid[param]; !exist {
			grid[param] = values
		}
	}
}
