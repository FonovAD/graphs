package controller_task

import (
	"context"
	gograph "github.com/yourbasic/graph"
	"golang_graphs/internal/models"
)

type controller struct {
}

func NewController() Controller {
	return &controller{}
}

type Controller interface {
	TaskComponents(ctx context.Context, graph models.Graph) (int, error)
	TaskIsEulerUndirected(ctx context.Context, graph models.Graph) (bool, error)
	TaskIsBipartition(ctx context.Context, graph models.Graph) (bool, error)
}

func (c *controller) TaskComponents(ctx context.Context, graph models.Graph) (int, error) {
	g := createGograph(graph)

	components := len(gograph.Components(g))

	return components, nil
}

func (c *controller) TaskIsEulerUndirected(ctx context.Context, graph models.Graph) (bool, error) {
	g := createGograph(graph)

	components := len(gograph.Components(g))

	if components > 1 {
		return false, nil
	}

	_, isEuler := gograph.EulerUndirected(g)

	return isEuler, nil
}

func (c *controller) TaskIsBipartition(ctx context.Context, graph models.Graph) (bool, error) {
	g := createGograph(graph)

	components := len(gograph.Components(g))

	if components > 1 {
		return false, nil
	}

	_, isBipartition := gograph.Bipartition(g)

	return isBipartition, nil
}

func createGograph(graph models.Graph) *gograph.Mutable {
	if len(graph.Nodes) == 0 {
		return gograph.New(0)
	}

	m := make(map[int]int)

	for i := 0; i < len(graph.Nodes); i++ {
		m[graph.Nodes[i].Id] = i
	}

	g := gograph.New(len(graph.Nodes))

	for _, link := range graph.Links {
		g.AddBoth(m[link.Source.Id], m[link.Target.Id])
	}

	return g
}
