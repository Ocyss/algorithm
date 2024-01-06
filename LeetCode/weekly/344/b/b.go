package main

// https://github.com/Ocyss
type FrequencyTracker struct {
	v map[int]int
	k map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{map[int]int{}, map[int]int{}}
}

func (f FrequencyTracker) Add(number int) {
	f.k[f.v[number]]--
	f.v[number]++
	f.k[f.v[number]]++
}

func (f FrequencyTracker) DeleteOne(number int) {
	f.k[f.v[number]]--
	f.v[number]--
	f.k[f.v[number]]++
	if f.v[number] <= 0 {
		delete(f.v, number)
	}
}

func (f FrequencyTracker) HasFrequency(frequency int) (ans bool) {
	if f.k[frequency] > 0 {
		return true
	}
	return false
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
