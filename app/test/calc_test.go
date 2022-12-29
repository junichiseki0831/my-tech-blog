package calc

import "testing"

func TestSummarize(t *testing.T) {
    var nums []int = []int{1, 2, 3, 4, 5}
    if !(Summarize(nums) == 15) {
        t.Error(`miss`)
    }
}
