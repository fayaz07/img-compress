package main

func largeSize(size int) int {
	return int(float64(size) * .75)
}

func mediumSize(size int) int {
	return int(float64(size) * .50)
}

func smallSize(size int) int {
	return int(float64(size) * .25)
}

func thumbnail(width int, height int) (int, int) {
	if width > height {
		if width > 512 {
			height = height * 512 / width
			width = 512
		}
	} else {
		if height > 512 {
			width = width * 512 / height
			height = 512
		}
	}
	return width, height
}
