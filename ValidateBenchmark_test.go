package validator

import (
	"testing"
	"time"
)

// go clean -testcache
// go test -run Benchmark -bench=. -benchmem .
// go test -run BenchmarkValidate -bench=. -benchmem .

func BenchmarkValidate(b *testing.B) {
	type Article struct {
		Title   string
		Age     uint8
		Phones  [4]string
		Images  []string
		Slices  [][]string
		Maps    []map[int]string
		Pages   []int
		Options map[int]string
		Date    time.Time
	}

	type Table struct {
		O string
		F Filter
		A Article
	}

	now := time.Now()

	table := []Table{
		// each:range

		{
			"Validate(each:range:int)",
			Filter{{
				Field: "Pages",
				Check: Rule{"each:range", []uint8{10, 20}},
			}},
			Article{Pages: []int{20}},
		},
		{
			"Validate(each:range:string)",
			Filter{{
				Field: "Images",
				Check: Rule{"each:range", []uint8{4, 6}},
			}},
			Article{Images: []string{"img1"}},
		},
		{
			"Validate(each:range:slice)",
			Filter{{
				Field: "Slices",
				Check: Rule{"each:range", []uint8{1, 3}},
			}},
			Article{Slices: [][]string{{"img1", "img2", "img3"}}},
		},
		{
			"Validate(each:range:map)",
			Filter{{
				Field: "Maps",
				Check: Rule{"each:range", []uint8{1, 3}},
			}},
			Article{Maps: []map[int]string{{1: "img1", 2: "img2", 3: "img3"}}},
		},

		// each:match

		{
			"Validate(each:match:slice)",
			Filter{{
				Field: "Images",
				Check: Rule{"each:match", `img\d`},
			}},
			Article{Images: []string{"img1"}},
		},
		{
			"Validate(each:match:map)",
			Filter{{
				Field: "Options",
				Check: Rule{"each:match", `img\d`},
			}},
			Article{Options: map[int]string{1: "img1"}},
		},

		// each:min

		{
			"Validate(each:min:int)",
			Filter{{
				Field: "Pages",
				Check: Rule{"each:min", 10},
			}},
			Article{Pages: []int{20}},
		},
		{
			"Validate(each:min:string)",
			Filter{{
				Field: "Images",
				Check: Rule{"each:min", 4},
			}},
			Article{Images: []string{"img1"}},
		},
		{
			"Validate(each:min:slice)",
			Filter{{
				Field: "Slices",
				Check: Rule{"each:min", 3},
			}},
			Article{Slices: [][]string{{"img1", "img2", "img3"}}},
		},
		{
			"Validate(each:min:map)",
			Filter{{
				Field: "Maps",
				Check: Rule{"each:min", 3},
			}},
			Article{Maps: []map[int]string{{1: "img1", 2: "img2", 3: "img3"}}},
		},

		// range

		{
			"Validate(range:int)",
			Filter{{
				Field: "Age",
				Check: Range{1, 20},
			}},
			Article{Age: 20},
		},
		{
			"Validate(range:string)",
			Filter{{
				Field: "Title",
				Check: Range{1, 20},
			}},
			Article{Title: "Buonasera signorina"},
		},
		{
			"Validate(range:slice)",
			Filter{{
				Field: "Images",
				Check: Range{1, 3},
			}},
			Article{Images: []string{"1", "2", "3"}},
		},
		{
			"Validate(range:map)",
			Filter{{
				Field: "Options",
				Check: Range{1, 3},
			}},
			Article{Options: map[int]string{1: "1", 2: "2", 3: "3"}},
		},

		// match

		{
			"Validate(match:string)",
			Filter{{
				Field: "Title",
				Check: Range{"match", `Buonasera`},
			}},
			Article{Title: "Buonasera"},
		},

		// min

		{
			"Validate(min:int)",
			Filter{{
				Field: "Age",
				Check: Rule{"min", 10},
			}},
			Article{Age: 20},
		},
		{
			"Validate(min:string)",
			Filter{{
				Field: "Title",
				Check: Rule{"min", 10},
			}},
			Article{Title: "Buonasera signorina"},
		},
		{
			"Validate(min:slice)",
			Filter{{
				Field: "Images",
				Check: Rule{"min", 3},
			}},
			Article{Images: []string{"1", "2", "3"}},
		},
		{
			"Validate(min:map)",
			Filter{{
				Field: "Options",
				Check: Rule{"min", 3},
			}},
			Article{Options: map[int]string{1: "1", 2: "2", 3: "3"}},
		},

		// date:min

		{
			"Validate(date:min:int64)",
			Filter{{
				Field: "Date",
				Check: Rule{"date:min", now.Unix()},
			}},
			Article{Date: now},
		},
		{
			"Validate(date:min:string)",
			Filter{{
				Field: "Date",
				Check: Rule{"date:min", now.Format(time.RFC3339)},
			}},
			Article{Date: now},
		},
		{
			"Validate(date:min:time)",
			Filter{{
				Field: "Date",
				Check: Rule{"date:min", now},
			}},
			Article{Date: now},
		},

		// time:min

		{
			"Validate(time:min:int64)",
			Filter{{
				Field: "Date",
				Check: Rule{"time:min", now.UnixNano()},
			}},
			Article{Date: now},
		},
		{
			"Validate(time:min:string)",
			Filter{{
				Field: "Date",
				Check: Rule{"time:min", "1705427897842183962"},
			}},
			Article{Date: now},
		},
		{
			"Validate(time:min:time)",
			Filter{{
				Field: "Date",
				Check: Rule{"time:min", now},
			}},
			Article{Date: now},
		},
	}

	for _, f := range table {
		filter := f

		b.Run(filter.O, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				filter.F.Validate(filter.A)
			}
		})
	}
}
