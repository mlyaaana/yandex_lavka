package slice

func Map[T any, R any](
	collection []T,
	mapper func(item T) (R, error),
) ([]R, error) {
	result := make([]R, len(collection))

	for i, item := range collection {
		mapped, err := mapper(item)
		if err != nil {
			return nil, err
		}
		result[i] = mapped
	}

	return result, nil
}
