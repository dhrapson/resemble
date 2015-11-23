//Copyright (C) 2015 dhrapson

//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
//along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"errors"
	. "github.com/dhrapson/resemble/configure"
	"gopkg.in/yaml.v2"
)

type ResembleConfig struct {
	typeName string `yaml: "type"`
	matchers []HttpMatcher
}

func (c *ResembleConfig) Parse(data []byte) error {

	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}
	if c.typeName == "" {
		return errors.New("Resemble config: invalid `type`")
	}

	return nil
}