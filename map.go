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

package Map

import (
	"errors"
	"sync"
)

type KV[K comparable, V interface{}] struct {
	Key   K
	Value V
}

type Map[K comparable, V interface{}] struct {
	m sync.Map
}

var (
	KeyNotFound = errors.New("key not found")
)

func (_map *Map[K, V]) zero() V {
	var value V
	return value
}

func (_map *Map[K, V]) Load(key K) (V, error) {
	v, ok := _map.m.Load(key)
	if !ok {
		return _map.zero(), KeyNotFound
	}

	return v.(V), nil
}

func (_map *Map[K, V]) Store(key K, value V) {
	_map.m.Store(key, value)
}

func (_map *Map[K, V]) Delete(key K) {
	_map.m.Delete(key)
}

func (_map *Map[K, V]) Clear() {
	_map.m.Clear()
}

func (_map *Map[K, V]) LoadOrStore(key K, value V) (V, error) {
	v, ok := _map.m.LoadOrStore(key, value)
	if !ok {
		return _map.zero(), KeyNotFound
	}

	return v.(V), nil
}

func (_map *Map[K, V]) LoadAndDelete(key K) (V, error) {
	v, ok := _map.m.LoadAndDelete(key)
	if !ok {
		return _map.zero(), KeyNotFound
	}

	return v.(V), nil
}

func (_map *Map[K, V]) Range(fn func(K, V) bool) {
	_map.m.Range(func(k, v any) bool {
		return fn(k.(K), v.(V))
	})
}

func (_map *Map[K, V]) Map() map[K]V {
	m := make(map[K]V)

	_map.Range(func(k K, v V) bool {
		m[k] = v
		return true
	})

	return m
}

func (_map *Map[K, V]) Len() int {
	i := 0

	_map.Range(func(_ K, _ V) bool {
		i++
		return true
	})

	return i
}

func New[K comparable, V interface{}]() *Map[K, V] {
	return &Map[K, V]{}
}
