// Copyright (c) 2018 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"go.uber.org/zap"
	
	"github.com/northwesternmutual/grammes"
	"github.com/northwesternmutual/grammes/examples/exampleutil"
)

func main() {
	logger := exampleutil.SetupLogger()
	defer logger.Sync()

	// Create a new Grammes client with a standard websocket.
	client, err := grammes.DialWithWebSocket("ws://127.0.0.1:8182")
	if err != nil {
		logger.Fatal("Couldn't create client", zap.Error(err))
	}

	// Drop all vertices on the graph currently.
	client.DropAll()

	// Drop the testing vertices when finished.
	defer client.DropAll()

	// Add testing vertices to the graph.
	client.AddVertex("person")
	client.AddVertex("car")

	// Create a graph traversal.
	g := grammes.Traversal()

	// Create a query using the created traversal to drop the vertices with the person label.
	err = client.DropVerticesByQuery(g.V().HasLabel("person").Drop())
	if err != nil {
		logger.Fatal("Error dropping vertices", zap.Error(err))
	}

	// Count the vertices after dropping.
	count, err := client.VertexCount()
	if err != nil {
		logger.Fatal("Unable to count vertices", zap.Error(err))
	}

	// Print out the vertices that are left over after dropping.
	// This should be 1.
	logger.Info("vertices left over", zap.Int64("count", count))
}