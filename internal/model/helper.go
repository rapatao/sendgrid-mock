package model

func MergeCategories(values ...Categories) Categories {
	var categories []string
	for _, value := range values {
		categories = append(categories, value...)
	}

	return categories
}

func MergeCustomArgs(values ...CustomArgs) CustomArgs {
	args := map[string]string{}

	for _, value := range values {
		if value == nil {
			continue
		}

		for k, v := range value {
			args[k] = v
		}
	}

	return args
}
