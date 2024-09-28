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
	"testing"
)

var _chan = Chan[string, uint64]{}

func TestChan(t *testing.T) {
	go func() {
		select {
		case data := <-_chan.Receive():
			t.Logf("data: %v", data)
		}
	}()

	_chan.Send("one", 64)
}

func TestChanLoad(t *testing.T) {
	one, err := _chan.Load("one")
	if err != nil {
		t.Errorf("failed to load one: %v", err)
	}

	t.Logf("loaded one: %v", one)
}

func TestChanMap(t *testing.T) {
	t.Logf("map: %v", _chan.Map())
}

func TestChanLen(t *testing.T) {
	t.Logf("length: %d", _chan.Len())
}

func TestChanClose(t *testing.T) {
	_chan.Close()
}
