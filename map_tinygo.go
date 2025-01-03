/*
 *     Enhanced `sync.Map` implementation for Go.
 *     Copyright (C) 2024  Dviih
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */

//go:build tinygo

package Map

import "reflect"

func (_map *Map[K, V]) Swap(key K, value V) (V, error) {
	p, err := _map.Load(key)
	if err != nil {
		p = _map.zero()
	}

	_map.Store(key, value)
	return p, nil
}

func (_map *Map[K, V]) CompareAndSwap(key K, old, new V) bool {
	c, err := _map.Load(key)
	if err != nil {
		return false
	}

	if reflect.DeepEqual(old, c) {
		return false
	}

	_map.Store(key, new)
	return true
}

func (_map *Map[K, V]) CompareAndDelete(key K, value V) bool {
	c, err := _map.Load(key)
	if err != nil {
		return false
	}

	if reflect.DeepEqual(value, c) {
		return false
	}

	_map.Delete(key)
	return true
}
