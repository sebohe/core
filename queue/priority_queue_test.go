// priority_queue_test.go - Tests for priority queue.
// Copyright (C) 2017  David Anthony Stainton, Yawning Angel
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package queue

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPriorityQueue(t *testing.T) {
	require := require.New(t)

	testEntries := []Entry{
		{
			Value:    []byte("That books do not take the place of experience,"),
			Priority: 0,
		},

		{
			Value:    []byte("and that learning is no substitute for genius,"),
			Priority: 1,
		},
		{
			Value:    []byte("are two kindred phenomena;"),
			Priority: 2,
		},
		{
			// Test duplicate priorities.
			//
			// Note: Duplicate priorities aren't guaranteed to be dequeued in
			// insertion order, though in this specific instance it will be.
			Value:    []byte("their common ground is that the abstract can never take the place of the perceptive."),
			Priority: 2,
		},
		{
			Value:    []byte(" -- Arthur_Schopenhauer"),
			Priority: 3,
		},
	}

	q := New()
	for _, v := range testEntries {
		q.Enqueue(v.Priority, v.Value)
	}
	require.Equal(len(testEntries), q.Len(), "Queue length (full)")

	for i, expected := range testEntries {
		require.Equal(len(testEntries)-i, q.Len(), "Queue length")

		// Peek
		ent := q.Peek()
		require.Equal(expected.Value, ent.Value, "Peek(): Value")
		require.Equal(expected.Priority, ent.Priority, "Peek(): Priority")

		// Pop
		ent = q.Pop()
		require.Equal(expected.Value, ent.Value, "Pop(): Value")
		require.Equal(expected.Priority, ent.Priority, "Pop(): Priority")

		t.Logf("ent[%d]: %d %s", i, ent.Priority, string(ent.Value))
	}

	require.Equal(0, q.Len(), "Queue length (empty)")
	require.Nil(q.Peek(), "Peek() (empty)")
	require.Nil(q.Pop(), "Pop() (empty)")

	// Refill the queue.
	for _, v := range testEntries {
		q.Enqueue(v.Priority, v.Value)
	}
	require.Equal(len(testEntries), q.Len(), "Queue length (full), pre-rand test")

	r := rand.New(rand.NewSource(23)) // Don't do this in production.
	for i := 0; i < len(testEntries); i++ {
		ent := q.DequeueRandom(r)
		t.Logf("random ent[%d]: %d %s", i, ent.Priority, string(ent.Value))
	}
	require.Equal(0, q.Len(), "Queue length (empty), post-rand test")
}
