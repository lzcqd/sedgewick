package quick_find

import (
	"bufio"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	cases := []struct {
		in   int
		want *Sites
	}{
		{1, &Sites{[]int{0}, 1}},
		{10, &Sites{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10}},
	}

	for _, c := range cases {
		got := Init(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Init(%d) == %+v, want %+v", c.in, got, c.want)
		}
	}
}

func TestFind(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{0, 0},
		{3, 3},
		{5, 11},
	}

	s := Init(10)
	s.id[5] = 11

	for _, c := range cases {
		got := s.Find(c.in)
		if got != c.want {
			t.Errorf("Find(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestUnion(t *testing.T) {
	cases := []struct {
		inputFile string
		want      *Sites
	}{
		{"../test_data/tinyUF.txt", &Sites{[]int{1, 1, 1, 8, 8, 1, 1, 1, 8, 8}, 2}},
	}

	for _, c := range cases {
		file, err := os.Open(c.inputFile)
		checkErr(t, err)
		defer file.Close()

		scanner := bufio.NewScanner(file)

		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		checkErr(t, scanner.Err())

		count, _ := strconv.Atoi(lines[0])
		lines = lines[1:]

		s := Init(count)

		for _, l := range lines {
			unionString := strings.Split(l, " ")
			union := make([]int, len(unionString))
			for i, v := range unionString {
				union[i], _ = strconv.Atoi(v)
			}
			s.Union(union[0], union[1])
		}

		if !reflect.DeepEqual(s, c.want) {
			t.Errorf("Union result is %+v, want %+v", s, c.want)
		}

	}
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error encountered: %q", err.Error())
	}
}

func TestConnected(t *testing.T) {
	cases := []struct {
		in   [2]int
		want bool
	}{
		{[2]int{1, 2}, false},
		{[2]int{3, 4}, true},
	}

	s := Init(5)
	s.id[3] = 4

	for _, c := range cases {
		got := s.Connected(c.in[0], c.in[1])
		if got != c.want {
			t.Errorf("Connected(%d, %d) == %t, want %t", c.in[0], c.in[1], got, c.want)
		}
	}
}

func TestCount(t *testing.T) {
	cases := []struct {
		count  int
		unions [][2]int
		want   int
	}{
		{5, [][2]int{[2]int{0, 1}}, 4},
		{5, [][2]int{{1, 2}, {2, 3}, {0, 4}}, 2},
	}
	for _, c := range cases {
		s := Init(c.count)
		for _, u := range c.unions {
			s.Union(u[0], u[1])
		}
		got := s.Count()
		if got != c.want {
			t.Errorf("Count returned %d, want %d, sites: %+v", got, c.want, s)
		}
	}

}
