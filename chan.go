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

type Chan[K comparable, V interface{}] struct {
	m Map[K, V]

	sender chan *KV[K, V]
	closed bool
}

func (_chan *Chan[K, V]) Send(key K, value V) {
	if _chan.closed {
		return
	}

	if _chan.sender == nil {
		_chan.sender = make(chan *KV[K, V])
	}

	_chan.m.Store(key, value)

	_chan.sender <- &KV[K, V]{
		Key:   key,
		Value: value,
	}
}

