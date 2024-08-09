/*
 * Copyright 2020 The AKACoder Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package enum

import (
	"reflect"
	"strconv"
)

const (
	elementDescriptionTag = "desc"
	elementNameTag        = "str"
)

type IElement interface {
	Int() int
	IntStr() string
	Name() string
	String() string
	Description() string
}

type Element struct {
	value int
	vStr  string
	name  string
	str   string
	desc  string
}

func (t *Element) Int() int {
	return t.value
}

func (t *Element) IntStr() string {
	return t.vStr
}

func (t *Element) Name() string {
	return t.name
}

func (t *Element) String() string {
	return t.str
}

func (t *Element) Description() string {
	return t.desc
}

func Build(enums interface{}, begin int, prefix string, strKey string) {
	buildEnum(enums, begin, false, func(value int, name string, tag reflect.StructTag) reflect.Value {
		var taggedName string
		if strKey != "" {
			taggedName = tag.Get(strKey)
		} else {
			taggedName = tag.Get(elementNameTag)
		}

		enumStr := prefix + name
		if taggedName != "" {
			enumStr = taggedName
		}

		vInt := value + begin
		vStr := strconv.Itoa(vInt)

		return reflect.ValueOf(Element{
			value: vInt,
			vStr:  vStr,
			name:  name,
			str:   enumStr,
			desc:  tag.Get(elementDescriptionTag),
		})
	})
	return
}

func SimpleBuild(enums interface{}) {
	Build(enums, 0, "", "")
}

func SimpleBuildWithNameKey(enums interface{}, strKey string) {
	Build(enums, 0, "", strKey)
}
