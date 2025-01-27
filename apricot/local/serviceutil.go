/*
 * === This file is part of ALICE O² ===
 *
 * Copyright 2019-2020 CERN and copyright holders of ALICE O².
 * Author: Teo Mrnjavac <teo.mrnjavac@cern.ch>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * In applying this license CERN does not waive the privileges and
 * immunities granted to it by virtue of its status as an
 * Intergovernmental Organization or submit itself to any jurisdiction.
 */

package local

import (
	"fmt"

	"github.com/AliceO2Group/Control/configuration/cfgbackend"
	"github.com/AliceO2Group/Control/configuration/componentcfg"
)

func (s *Service) queryToAbsPath(query *componentcfg.Query) (absolutePath string, err error) {
	if query == nil {
		return
	}

	var timestamp string

	if len(query.Timestamp) == 0 {
		keyPrefix := query.AbsoluteWithoutTimestamp()
		if s.src.IsDir(keyPrefix) {
			var keys []string
			keys, err = s.src.GetKeysByPrefix(keyPrefix)
			if err != nil {
				return
			}
			timestamp, err = componentcfg.GetLatestTimestamp(keys, query)
			if err != nil {
				timestamp = "" // no timestamp keys found, we'll try to fall back to timestampless
			}
		}
	} else {
		timestamp = query.Timestamp
	}
	absolutePath = query.AbsoluteWithoutTimestamp() + componentcfg.SEPARATOR + timestamp
	if exists, _ := s.src.Exists(absolutePath); exists && len(timestamp) > 0 {
		err = nil
		return
	}

	// falling back to timestampless configuration
	absolutePath = query.AbsoluteWithoutTimestamp()
	if exists, _ := s.src.Exists(absolutePath); exists {
		err = nil
		return
	}

	err = fmt.Errorf("no payload at configuration path %s", absolutePath)

	return
}

func (s *Service) getStringMap(path string) map[string]string {
	tree, err := s.src.GetRecursive(path)
	if err != nil {
		return nil
	}
	if tree.Type() == cfgbackend.IT_Map {
		responseMap := tree.Map()
		theMap := make(map[string]string, len(responseMap))
		for k, v := range responseMap {
			if v.Type() != cfgbackend.IT_Value {
				continue
			}
			theMap[k] = v.Value()
		}
		return theMap
	}
	return nil
}

func (s *Service) resolveComponentQuery(query *componentcfg.Query) (resolved *componentcfg.Query, err error) {
	resolved = &componentcfg.Query{}
	*resolved = *query
	if _, err = s.queryToAbsPath(resolved); err == nil {
		// requested path exists, return it
		return
	}

	resolved = query.WithFallbackRunType()
	if _, err = s.queryToAbsPath(resolved); err == nil {
		// path with run type ANY exists, return it
		return
	}

	resolved = query.WithFallbackRoleName()
	if _, err = s.queryToAbsPath(resolved); err == nil {
		// path with role name "any" exists, return it
		return
	}

	resolved = resolved.WithFallbackRunType()
	if _, err = s.queryToAbsPath(resolved); err == nil {
		// path with run type ANY and role name "any" exists, return it
		return
	}

	return nil, fmt.Errorf("could not resolve configuration path %s", query.AbsoluteRaw())
}
