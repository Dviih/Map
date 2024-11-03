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

import "github.com/Dviih/Channel"

type Chan[K comparable, V interface{}] struct {
	m Map[K, V]

	channel *Channel.Channel[*KV[K, V]]
	closed  bool
}

func (_chan *Chan[K, V]) Send(key K, value V) {
	if _chan.closed {
		return
	}

	_chan.m.Store(key, value)

	_chan.channel.Sender() <- &KV[K, V]{
		Key:   key,
		Value: value,
	}
}

func (_chan *Chan[K, V]) Receive() <-chan *KV[K, V] {
	if _chan.closed {
		return nil
	}

	c := make(chan *KV[K, V])

	go func() {
		for {
			select {
			case data := <-_chan.sender:
				c <- data
			}
		}
	}()

	return c
}

func (_chan *Chan[K, V]) Close() {
	if _chan.closed {
		return
	}

	_chan.closed = true
}

func (_chan *Chan[K, V]) Load(key K) (V, error) {
	return _chan.m.Load(key)
}

func (_chan *Chan[K, V]) Delete(key K) {
	_chan.m.Delete(key)
}

func (_chan *Chan[K, V]) Map() map[K]V {
	return _chan.m.Map()
}

func (_chan *Chan[K, V]) Len() int {
	return _chan.m.Len()
}
