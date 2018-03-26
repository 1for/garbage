package sort

func ShellSort( data []int) []int{
	var i, j, tmp, increament int

	var count = len(data)


	for increament = (count / 2); increament > 0; increament-- {

		for i = increament; i < count; i++ {
			tmp = data[i]
			for j = i; j > 0 && data[j-1] > tmp; j-- {
				data[j] = data[j-1]
			}
			data[j] = tmp
		}
	}

	return data;
}
