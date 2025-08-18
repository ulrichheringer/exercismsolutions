package cards

func FavoriteCards() []int {
	return []int{2, 6, 9}
}
func GetItem(slice []int, index int) int {
	if index >= len(slice) || index < 0 {
		return -1
	}
	return slice[index]
}
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
		return append(slice, value)
	}
	return append(slice[:index], append([]int{value}, slice[index+1:]...)...)
}
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}
func RemoveItem(slice []int, index int) []int {
	if index >= len(slice) || index < 0 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
