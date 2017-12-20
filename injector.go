package chip

import "github.com/facebookgo/inject"

func Inject(result interface{}, dependencies ...interface{}) error {
	objects := []*inject.Object{
		{Value: result},
	}
	appendDependencyObjects(&objects, dependencies)

	var graph inject.Graph
	if err := graph.Provide(objects...); err != nil {
		return err
	}
	return graph.Populate()
}

func appendDependencyObjects(
	objects *[]*inject.Object, dependencies []interface{},
) {
	for _, dep := range dependencies {
		*objects = append(*objects, &inject.Object{
			Value:    dep,
			Complete: true,
		})
	}
}
