package collect

import (
	"fmt"
	"github.com/melodywen/supports/utils"
	"log"
	"reflect"
	"testing"
)

func TestTypeTransformOrFail(t *testing.T) {
	t.Run("MapSlice-nil", func(t *testing.T) {
		var data []int
		var data2 interface{}
		data2 = data
		got := TypeTransformOrFail[interface{}, []int](data2)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapSlice() = %v, want %v", got, want)
		}
	})
	t.Run("MapSlice-nil", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("run time panic: %v", err)
			}
		}()
		var data []int
		var data2 interface{}
		data2 = data
		got := TypeTransformOrFail[interface{}, []string](data2)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapSlice() = %v, want %v", got, want)
		}
	})
}

func TestMapSlice(t *testing.T) {
	t.Run("MapSlice-nil", func(t *testing.T) {
		var data []int
		got := MapSlice(data, func(key int, value int) string {
			return fmt.Sprintf("%d", value)
		})
		var want []string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapSlice() = %v, want %v", got, want)
		}
	})

	t.Run("MapSlice-score", func(t *testing.T) {
		data := []int{10, 20, 30}
		got := MapSlice(data, func(key int, value int) string {
			return fmt.Sprintf("get-score->%d", value)
		})
		want := []string{"get-score->10", "get-score->20", "get-score->30"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapSlice() = %v, want %v", got, want)
		}
	})
}

func TestMapMap(t *testing.T) {
	t.Run("MapMap-nil", func(t *testing.T) {
		var data map[string]int
		got := MapMap(data, func(key string, value int) string {
			return fmt.Sprintf("%s->%d", key, value)
		})
		var want map[string]string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapSlice() = %v, want %v", got, want)
		}
	})

	t.Run("MapMap-score", func(t *testing.T) {
		data := map[string]int{
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}
		got := MapMap(data, func(key string, value int) string {
			return fmt.Sprintf("%s->%d", key, value)
		})
		want := map[string]string{
			"english": "english->60", "language": "language->80", "mathematics": "mathematics->70",
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapSlice() = %v, want %v", got, want)
		}
	})
}

func TestMapSliceWithKeys(t *testing.T) {
	t.Run("MapSliceWithKeys-nil", func(t *testing.T) {
		var data []int
		got := MapSliceWithKeys(data, func(key int, value int) (string, string) {
			return fmt.Sprintf("key-%d", key), fmt.Sprintf("value-%d", value)
		})
		var want map[string]string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapSliceWithKeys() = %v, want %v", got, want)
		}
	})

	t.Run("MapSliceWithKeys-score", func(t *testing.T) {
		data := []int{10, 20, 30}
		got := MapSliceWithKeys(data, func(key int, value int) (string, string) {
			return fmt.Sprintf("key-%d", key), fmt.Sprintf("value-%d", value)
		})
		want := map[string]string{"key-0": "value-10", "key-1": "value-20", "key-2": "value-30"}
		if !reflect.DeepEqual(got, want) {
			utils.Dump(got)
			t.Errorf("MapSliceWithKeys() = %v, want %v", got, want)
		}
	})
}

func TestMapMapWithKeys(t *testing.T) {
	t.Run("MapMapWithKeys-nil", func(t *testing.T) {
		var data map[int]int
		got := MapMapWithKeys(data, func(key int, value int) (string, string) {
			return fmt.Sprintf("key-%d", value), fmt.Sprintf("value-%d", value)
		})
		var want map[string]string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapMapWithKeys() = %v, want %v", got, want)
		}
	})

	t.Run("MapMapWithKeys-score", func(t *testing.T) {
		data := map[int]int{11: 10, 12: 20, 13: 30}
		got := MapMapWithKeys(data, func(key int, value int) (string, string) {
			return fmt.Sprintf("key-%d", key), fmt.Sprintf("value-%d", value)
		})
		want := map[string]string{"key-11": "value-10", "key-12": "value-20", "key-13": "value-30"}
		if !reflect.DeepEqual(got, want) {
			utils.Dump(got)
			t.Errorf("MapMapWithKeys() = %v, want %v", got, want)
		}
	})
}

func TestFilterSlice(t *testing.T) {
	t.Run("FilterSlice-nil", func(t *testing.T) {
		var data []int
		got := FilterSlice(data, func(key int, value int) bool {
			return value > 30
		})
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FilterSlice() = %v, want %v", got, want)
		}
	})

	t.Run("FilterSlice-score", func(t *testing.T) {
		data := []int{10, 20, 30}
		got := FilterSlice(data, func(key int, value int) bool {
			return value > 11
		})
		want := []int{20, 30}
		if !reflect.DeepEqual(got, want) {
			utils.Dump(got)
			t.Errorf("FilterSlice() = %v, want %v", got, want)
		}
	})
}

func TestFilterMap(t *testing.T) {
	t.Run("FilterMap-nil", func(t *testing.T) {
		var data map[int]int
		got := FilterMap(data, func(key int, value int) bool {
			return value > 30
		})
		var want map[int]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FilterMap() = %v, want %v", got, want)
		}
	})

	t.Run("FilterMap-score", func(t *testing.T) {
		data := map[int]int{11: 10, 12: 20, 13: 30}
		got := FilterMap(data, func(key int, value int) bool {
			return value > 11
		})
		want := map[int]int{12: 20, 13: 30}
		if !reflect.DeepEqual(got, want) {
			utils.Dump(got)
			t.Errorf("FilterMap() = %v, want %v", got, want)
		}
	})
}

func TestEachSlice(t *testing.T) {
	t.Run("EachSlice-score", func(t *testing.T) {
		data := []int{10, 20, 30, 50}
		got := 0
		EachSlice(data, func(key int, value int) bool {
			got += value
			return value < 21
		})
		want := 60
		if !reflect.DeepEqual(got, want) {
			utils.Dump(got)
			t.Errorf("EachSlice() = %v, want %v", got, want)
		}
	})
}

func TestEachMap(t *testing.T) {
	t.Run("EachMap-score", func(t *testing.T) {
		data := map[int]int{11: 10, 12: 20, 13: 30}
		got := 0
		EachMap(data, func(key int, value int) bool {
			got += value + key
			return value < 21
		})
		want := 96
		if got > want {
			utils.Dump(got)
			t.Errorf("EachMap() = %v, want %v", got, want)
		}
	})
}

func TestEverySlice(t *testing.T) {
	t.Run("EverySlice-score", func(t *testing.T) {
		data := []int{10, 20, 30, 50}
		got := EverySlice(data, func(key int, value int) bool {
			return value < 21
		})
		want := false
		if !reflect.DeepEqual(got, want) {
			utils.Dump(got)
			t.Errorf("EverySlice() = %v, want %v", got, want)
		}
	})
	t.Run("EverySlice-score", func(t *testing.T) {
		data := []int{10, 20, 30, 50}
		got := EverySlice(data, func(key int, value int) bool {
			return value < 71
		})
		want := true
		if !reflect.DeepEqual(got, want) {
			utils.Dump(got)
			t.Errorf("EverySlice() = %v, want %v", got, want)
		}
	})
}

func TestEveryMap(t *testing.T) {
	t.Run("EveryMap-score", func(t *testing.T) {
		data := map[int]int{11: 10, 12: 20, 13: 30}
		got := EveryMap(data, func(key int, value int) bool {
			return value < 21
		})
		want := false
		if !reflect.DeepEqual(got, want) {
			utils.Dump(got)
			t.Errorf("EveryMap() = %v, want %v", got, want)
		}
	})
	t.Run("EveryMap-score", func(t *testing.T) {
		data := map[int]int{11: 10, 12: 20, 13: 30}
		got := EveryMap(data, func(key int, value int) bool {
			return value < 40
		})
		want := true
		if !reflect.DeepEqual(got, want) {
			utils.Dump(got)
			t.Errorf("EveryMap() = %v, want %v", got, want)
		}
	})
}

func TestContainsSlice(t *testing.T) {
	t.Run("ContainsSlice-nil", func(t *testing.T) {
		var data []int
		got := ContainsSlice(data, 1)
		var want bool
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ContainsSlice() = %v, want %v", got, want)
		}
	})
	t.Run("ContainsSlice-score-true", func(t *testing.T) {
		data := []int{10, 20, 30}
		got := ContainsSlice(data, 10)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ContainsSlice() = %v, want %v", got, want)
		}
	})
	t.Run("ContainsSlice-score-true", func(t *testing.T) {
		data := []int{10, 20, 30}
		got := ContainsSlice(data, 11)
		want := false
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ContainsSlice() = %v, want %v", got, want)
		}
	})
}

func TestContainsMap(t *testing.T) {
	t.Run("ContainsMap-nil", func(t *testing.T) {
		var data []map[int]int
		got := ContainsMap(data, 1, 2)
		var want bool
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ContainsMap() = %v, want %v", got, want)
		}
	})
	t.Run("ContainsMap-score-false", func(t *testing.T) {
		data := []map[int]int{{11: 10, 12: 20, 13: 30}}
		got := ContainsMap(data, 10, 11)
		want := false
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ContainsMap() = %v, want %v", got, want)
		}
	})
	t.Run("ContainsMap-score-false", func(t *testing.T) {
		data := []map[int]int{{11: 10, 12: 20, 13: 30}}
		got := ContainsMap(data, 12, 11)
		want := false
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ContainsMap() = %v, want %v", got, want)
		}
	})
	t.Run("ContainsMap-score-true", func(t *testing.T) {
		data := []map[int]int{{11: 10, 12: 20, 13: 30}}
		got := ContainsMap(data, 12, 20)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ContainsMap() = %v, want %v", got, want)
		}
	})

}

func TestDoesNotContainsSlice(t *testing.T) {
	t.Run("DoesNotContainsSlice-nil", func(t *testing.T) {
		var data []int
		got := DoesNotContainsSlice(data, 1)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoesNotContainsSlice() = %v, want %v", got, want)
		}
	})
	t.Run("DoesNotContainsSlice-score-true", func(t *testing.T) {
		data := []int{10, 20, 30}
		got := DoesNotContainsSlice(data, 10)
		want := false
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoesNotContainsSlice() = %v, want %v", got, want)
		}
	})
	t.Run("DoesNotContainsSlice-score-true", func(t *testing.T) {
		data := []int{10, 20, 30}
		got := DoesNotContainsSlice(data, 11)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoesNotContainsSlice() = %v, want %v", got, want)
		}
	})
}

func TestDoesNotContainsMap(t *testing.T) {
	t.Run("DoesNotContainsMap-nil", func(t *testing.T) {
		var data []map[int]int
		got := DoesNotContainsMap(data, 1, 2)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoesNotContainsMap() = %v, want %v", got, want)
		}
	})
	t.Run("DoesNotContainsMap-score-false", func(t *testing.T) {
		data := []map[int]int{{11: 10, 12: 20, 13: 30}}
		got := DoesNotContainsMap(data, 10, 11)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoesNotContainsMap() = %v, want %v", got, want)
		}
	})
	t.Run("DoesNotContainsMap-score-false", func(t *testing.T) {
		data := []map[int]int{{11: 10, 12: 20, 13: 30}}
		got := DoesNotContainsMap(data, 12, 11)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoesNotContainsMap() = %v, want %v", got, want)
		}
	})
	t.Run("DoesNotContainsMap-score-true", func(t *testing.T) {
		data := []map[int]int{{11: 10, 12: 20, 13: 30}}
		got := DoesNotContainsMap(data, 12, 20)
		want := false
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DoesNotContainsMap() = %v, want %v", got, want)
		}
	})
}

func TestSort(t *testing.T) {
	t.Run("Sort-nil", func(t *testing.T) {
		data := []int{}
		got := Sort(data)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Sort() = %v, want %v", got, want)
		}
	})
	t.Run("Sort-int", func(t *testing.T) {
		data := []int{5, 3, 2, 1, 4}
		got := Sort(data)
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Sort() = %v, want %v", got, want)
		}
	})
	t.Run("Sort-float", func(t *testing.T) {
		data := []float64{1.5, 1.3, 1.4, 1.1, 1.2}
		got := Sort(data)
		want := []float64{1.1, 1.2, 1.3, 1.4, 1.5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Sort() = %v, want %v", got, want)
		}
	})
	t.Run("Sort-string", func(t *testing.T) {
		data := []string{"g", "d", "f", "a"}
		got := Sort(data)
		want := []string{"a", "d", "f", "g"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Sort() = %v, want %v", got, want)
		}
	})
}

func TestSortDesc(t *testing.T) {
	t.Run("SortDesc-nil", func(t *testing.T) {
		data := []int{}
		got := SortDesc(data)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("SortDesc() = %v, want %v", got, want)
		}
	})
	t.Run("SortDesc-int", func(t *testing.T) {
		data := []int{5, 3, 2, 1, 4}
		got := SortDesc(data)
		want := []int{5, 4, 3, 2, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("SortDesc() = %v, want %v", got, want)
		}
	})
	t.Run("SortDesc-float", func(t *testing.T) {
		data := []float64{1.5, 1.3, 1.4, 1.1, 1.2}
		got := SortDesc(data)
		want := []float64{1.5, 1.4, 1.3, 1.2, 1.1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("SortDesc() = %v, want %v", got, want)
		}
	})
	t.Run("SortDesc-string", func(t *testing.T) {
		data := []string{"g", "d", "f", "a"}
		got := SortDesc(data)
		want := []string{"g", "f", "d", "a"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("SortDesc() = %v, want %v", got, want)
		}
	})
}

func TestSortBy(t *testing.T) {
	t.Run("SortBy-score", func(t *testing.T) {
		data := []map[string]int{{
			"username":    10001,
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10003,
			"english":     70,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10002,
			"english":     68,
			"mathematics": 70,
			"language":    80,
		}}
		got := SortByDesc(data, func(key int, value map[string]int) int {
			return value["english"]
		})
		want := []map[string]int{{
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10003,
			"english":     70,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10002,
			"english":     68,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10001,
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapSlice() = %v, want %v", got, want)
		}
	})
	t.Run("SortBy-score", func(t *testing.T) {
		data := []map[string]int{{
			"username":    10001,
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10003,
			"english":     70,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10002,
			"english":     68,
			"mathematics": 70,
			"language":    80,
		}}
		got := SortBy(data, func(key int, value map[string]int) int {
			return value["english"]
		})
		want := []map[string]int{{
			"username":    10001,
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10002,
			"english":     68,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10003,
			"english":     70,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    80,
		}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapSlice() = %v, want %v", got, want)
		}
	})
	t.Run("SortBy-score", func(t *testing.T) {
		var data []map[string]int
		got := SortBy(data, func(key int, value map[string]int) int {
			return value["english"]
		})
		var want []map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapSlice() = %v, want %v", got, want)
		}
	})
}

func TestKeys(t *testing.T) {
	t.Run("Keys-nil", func(t *testing.T) {
		var data map[string]int
		got := Keys(data)
		var want []string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Keys() = %v, want %v", got, want)
		}
	})
	t.Run("Keys-score-false", func(t *testing.T) {
		data := map[string]int{
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}
		got := Keys(data)
		want := []string{"english", "mathematics", "language"}
		if !reflect.DeepEqual(Sort(got), Sort(want)) {
			t.Errorf("Keys() = %v, want %v", got, want)
		}
	})
}

func TestValues(t *testing.T) {
	t.Run("Values-nil", func(t *testing.T) {
		var data map[string]int
		got := Values(data)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Values() = %v, want %v", got, want)
		}
	})
	t.Run("Values-score-false", func(t *testing.T) {
		data := map[string]int{
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}
		got := Values(data)
		want := []int{70, 80, 60}
		if !reflect.DeepEqual(Sort(got), Sort(want)) {
			t.Errorf("Values() = %v, want %v", got, want)
		}
	})
}

func TestExcept(t *testing.T) {
	t.Run("Except-nil", func(t *testing.T) {
		var data map[string]int
		got := Except(data, []string{})
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Except() = %v, want %v", got, want)
		}
	})
	t.Run("Except-score-false", func(t *testing.T) {
		data := map[string]int{
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}
		got := Except(data, []string{"language", "english"})
		want := map[string]int{
			"mathematics": 70,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Except() = %v, want %v", got, want)
		}
	})
}
func TestOnly(t *testing.T) {
	t.Run("Only-nil", func(t *testing.T) {
		var data map[string]int
		got := Only(data, []string{})
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Only() = %v, want %v", got, want)
		}
	})
	t.Run("Only-score-false", func(t *testing.T) {
		data := map[string]int{
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}
		got := Only(data, []string{"language", "english"})
		want := map[string]int{
			"english":  60,
			"language": 80,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Only() = %v, want %v", got, want)
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("sum-score-false", func(t *testing.T) {
		data := []int{70, 80, 60}
		got := Sum(data)
		want := 210
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Sum() = %v, want %v", got, want)
		}
	})
}

func TestSumSlice(t *testing.T) {
	t.Run("SumSlice-score", func(t *testing.T) {
		data := []map[string]int{{
			"username":    10001,
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10003,
			"english":     70,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10002,
			"english":     68,
			"mathematics": 70,
			"language":    80,
		}}
		got := SumSlice(data, func(_ int, v map[string]int) int {
			return v["english"]
		})
		want := 278
		if !reflect.DeepEqual(got, want) {
			t.Errorf("SumSlice() = %v, want %v", got, want)
		}
	})
}

func TestAverage(t *testing.T) {
	t.Run("Average-score-nil", func(t *testing.T) {
		data := []int{}
		got := Average(data)
		want := 0
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Average() = %v, want %v", got, want)
		}
	})
	t.Run("Average-score", func(t *testing.T) {
		data := []int{70, 80, 60}
		got := Average(data)
		want := 70
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Average() = %v, want %v", got, want)
		}
	})

}

func TestAverageSlice(t *testing.T) {
	t.Run("AverageSlice-score-nil", func(t *testing.T) {
		data := []map[string]int{}
		got := AverageSlice(data, func(_ int, v map[string]int) int {
			return v["english"]
		})
		want := 0
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AverageSlice() = %v, want %v", got, want)
		}
	})
	t.Run("AverageSlice-score", func(t *testing.T) {
		data := []map[string]int{{
			"username":    10001,
			"english":     260,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10003,
			"english":     70,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10002,
			"english":     68,
			"mathematics": 70,
			"language":    80,
		}}
		got := AverageSlice(data, func(_ int, v map[string]int) int {
			return v["english"]
		})
		want := 119
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AverageSlice() = %v, want %v", got, want)
		}
	})
}

func TestChunk(t *testing.T) {
	t.Run("Chunk-score-nil", func(t *testing.T) {
		var data []int
		got := Chunk(data, 4)
		var want [][]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Chunk() = %v, want %v", got, want)
		}
	})
	t.Run("Chunk-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := Chunk(data, 3)
		want := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Chunk() = %v, want %v", got, want)
		}
	})
	t.Run("Chunk-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := Chunk(data, 4)
		want := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Chunk() = %v, want %v", got, want)
		}
	})
}

func TestCollapse(t *testing.T) {
	t.Run("Collapse-score-nil", func(t *testing.T) {
		var data [][]int
		got := Collapse(data)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Collapse() = %v, want %v", got, want)
		}
	})
	t.Run("Collapse-score", func(t *testing.T) {
		data := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9}}
		got := Collapse(data)
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Collapse() = %v, want %v", got, want)
		}
	})
}

func TestFlatten(t *testing.T) {
	t.Run("Flatten-score-nil", func(t *testing.T) {
		var data map[string][]int
		got := Flatten(data)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})
	t.Run("Flatten-score", func(t *testing.T) {
		data := map[string][]int{
			"english":     {60, 61},
			"mathematics": {71, 72},
			"language":    {83, 84},
		}
		got := Flatten(data)
		want := []int{60, 61, 71, 72, 83, 84}
		if !reflect.DeepEqual(Sort(got), want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})
}

func TestFlip(t *testing.T) {
	t.Run("Flip-score-nil", func(t *testing.T) {
		var data map[string]int
		got := Flip(data)
		var want map[int]string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flip() = %v, want %v", got, want)
		}
	})
	t.Run("Flip-score-nil", func(t *testing.T) {
		data := map[string]int{"english": 60, "mathematics": 70, "language": 80}
		got := Flip(data)
		want := map[int]string{60: "english", 70: "mathematics", 80: "language"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flip() = %v, want %v", got, want)
		}
	})
}

func TestFlatMap(t *testing.T) {
	t.Run("FlatMap-score-nil", func(t *testing.T) {
		var data map[int]string
		got := FlatMap(data, func(k int, v string) string {
			return v
		})
		var want map[string]string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FlatMap() = %v, want %v", got, want)
		}
	})
	t.Run("FlatMap-score-nil", func(t *testing.T) {
		data := map[int]string{60: "english", 70: "mathematics", 80: "language"}
		got := FlatMap(data, func(k int, v string) string {
			return v
		})
		want := map[string]string{"english": "english", "language": "language", "mathematics": "mathematics"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FlatMap() = %v, want %v", got, want)
		}
	})
}

func TestForget(t *testing.T) {
	t.Run("Forget-score-nil", func(t *testing.T) {
		var data map[string]int
		got := Forget(data, []string{"english", "english2"})
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Forget() = %v, want %v", got, want)
		}
	})
	t.Run("Forget-score-nil", func(t *testing.T) {
		data := map[string]int{"english": 60, "mathematics": 70, "language": 80}
		got := Forget(data, []string{"english", "english2"})
		want := map[string]int{"mathematics": 70, "language": 80}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Forget() = %v, want %v", got, want)
		}
	})
}

func TestCombine(t *testing.T) {
	t.Run("Combine-score-nil", func(t *testing.T) {
		var data []string
		var data2 []int
		got := Combine(data, data2)
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Combine() = %v, want %v", got, want)
		}
	})
	t.Run("Combine-score", func(t *testing.T) {
		data := []string{"english", "mathematics", "language"}
		data2 := []int{60, 70, 80}
		got := Combine(data, data2)
		want := map[string]int{
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Combine() = %v, want %v", got, want)
		}
	})
	t.Run("Combine-score", func(t *testing.T) {
		data := []string{"english", "mathematics"}
		data2 := []int{60, 70, 80}
		got := Combine(data, data2)
		want := map[string]int{
			"english":     60,
			"mathematics": 70,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Combine() = %v, want %v", got, want)
		}
	})
	t.Run("Combine-score", func(t *testing.T) {
		data := []string{"english", "mathematics", "language"}
		data2 := []int{60, 70}
		got := Combine(data, data2)
		want := map[string]int{
			"english":     60,
			"mathematics": 70,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Combine() = %v, want %v", got, want)
		}
	})
}
func TestGet(t *testing.T) {
	t.Run("Get-score-nil", func(t *testing.T) {
		var data map[string]int
		got := Get(data, "english")
		var want int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Get() = %v, want %v", got, want)
		}
	})
	t.Run("Get-score-true", func(t *testing.T) {
		data := map[string]int{"english": 60, "mathematics": 70, "language": 80}
		got := Get(data, "english")
		want := 60
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Get() = %v, want %v", got, want)
		}
	})
	t.Run("Get-score-false", func(t *testing.T) {
		data := map[string]int{"english": 60, "mathematics": 70, "language": 80}
		got := Get(data, "english2")
		want := 0
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Get() = %v, want %v", got, want)
		}
	})
}

func TestGetOrDefault(t *testing.T) {
	t.Run("GetOrDefault-score-nil", func(t *testing.T) {
		var data map[string]int
		got := GetOrDefault(data, "english", -1)
		want := -1
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetOrDefault() = %v, want %v", got, want)
		}
	})
	t.Run("GetOrDefault-score-true", func(t *testing.T) {
		data := map[string]int{"english": 60, "mathematics": 70, "language": 80}
		got := GetOrDefault(data, "english", -1)
		want := 60
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetOrDefault() = %v, want %v", got, want)
		}
	})
	t.Run("GetOrDefault-score-false", func(t *testing.T) {
		data := map[string]int{"english": 60, "mathematics": 70, "language": 80}
		got := GetOrDefault(data, "english2", -1)
		want := -1
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetOrDefault() = %v, want %v", got, want)
		}
	})
}

func TestHas(t *testing.T) {
	t.Run("Has-score-nil", func(t *testing.T) {
		var data map[string]int
		got := Has(data, []string{"english"})
		want := false
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Has() = %v, want %v", got, want)
		}
	})
	t.Run("Has-score-true", func(t *testing.T) {
		data := map[string]int{"english": 60, "mathematics": 70, "language": 80}
		got := Has(data, []string{"english1"})
		want := false
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Has() = %v, want %v", got, want)
		}
	})
	t.Run("Has-score-false", func(t *testing.T) {
		data := map[string]int{"english": 60, "mathematics": 70, "language": 80}
		got := Has(data, []string{"english", "mathematics"})
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Has() = %v, want %v", got, want)
		}
	})
}

func TestFirst(t *testing.T) {
	t.Run("First-score-nil", func(t *testing.T) {
		var data []map[string]int
		got := First(data, func(key int, item map[string]int) bool {
			return key > 1
		})
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("First() = %v, want %v", got, want)
		}
	})
	t.Run("First-score-true", func(t *testing.T) {
		data := []map[string]int{{
			"username":    10001,
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10003,
			"english":     70,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10002,
			"english":     68,
			"mathematics": 70,
			"language":    80,
		}}
		got := First(data, func(key int, item map[string]int) bool {
			return item["english"] > 65
		})
		want := map[string]int{
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    80,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("First() = %v, want %v", got, want)
		}
	})
	t.Run("First-score-none", func(t *testing.T) {
		data := []map[string]int{}
		got := First(data, func(key int, item map[string]int) bool {
			return item["english"] > 65
		})
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("First() = %v, want %v", got, want)
		}
	})

}

func TestLast(t *testing.T) {
	t.Run("Last-score-nil", func(t *testing.T) {
		var data []map[string]int
		got := Last(data, func(key int, item map[string]int) bool {
			return key > 1
		})
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Last() = %v, want %v", got, want)
		}
	})
	t.Run("Last-score-true", func(t *testing.T) {
		data := []map[string]int{{
			"username":    10001,
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10004,
			"english":     68,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10003,
			"english":     70,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10002,
			"english":     268,
			"mathematics": 70,
			"language":    80,
		}}
		got := Last(data, func(key int, item map[string]int) bool {
			return item["english"] < 85
		})
		want := map[string]int{
			"username":    10003,
			"english":     70,
			"mathematics": 70,
			"language":    80,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Last() = %v, want %v", got, want)
		}
	})
	t.Run("Last-score-none", func(t *testing.T) {
		data := []map[string]int{}
		got := Last(data, func(key int, item map[string]int) bool {
			return item["english"] < 85
		})
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Last() = %v, want %v", got, want)
		}
	})
}
func TestCountBy(t *testing.T) {
	t.Run("CountBy-score-nil", func(t *testing.T) {
		var data []string
		got := CountBy(data, func(i int, v string) string {
			return v
		})
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("CountBy() = %v, want %v", got, want)
		}
	})
	t.Run("CountBy-score", func(t *testing.T) {
		data := []string{"english", "english", "language", "language", "language"}
		got := CountBy(data, func(i int, v string) string {
			return v
		})
		want := map[string]int{"english": 2, "language": 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("CountBy() = %v, want %v", got, want)
		}
	})
}

func TestCrossJoin(t *testing.T) {
	t.Run("CrossJoin-score-nil", func(t *testing.T) {
		var data [][]string
		var data2 [][]string
		got := CrossJoin(data, data2)
		var want [][]string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("CrossJoin() = %v, want %v", got, want)
		}
	})
	t.Run("CrossJoin-score", func(t *testing.T) {
		data := [][]string{{"tom"}, {"jerry"}}
		data2 := [][]string{{"tom2", "tom21"}, {"jerry2", "jerry21"}}
		got := CrossJoin(data, data2)
		want := [][]string{{"tom", "tom2", "tom21"}, {"tom", "jerry2", "jerry21"}, {"jerry", "tom2", "tom21"}, {"jerry", "jerry2", "jerry21"}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("CrossJoin() = %v, want %v", got, want)
		}
	})
}

func TestDiff(t *testing.T) {
	t.Run("Diff-score-nil", func(t *testing.T) {
		var data []string
		var data2 []string
		got := Diff(data, data2)
		var want []string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Diff() = %v, want %v", got, want)
		}
	})
	t.Run("Diff-score", func(t *testing.T) {
		data := []string{"one", "two", "three", "four"}
		data2 := []string{"one", "two2", "three3", "four"}
		got := Diff(data, data2)
		want := []string{"two", "three"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Diff() = %v, want %v", got, want)
		}
	})
}

func TestDiffAssoc(t *testing.T) {
	t.Run("DiffAssoc-score-nil", func(t *testing.T) {
		var data map[string]int
		var data2 map[string]int
		got := DiffAssoc(data, data2)
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DiffAssoc() = %v, want %v", got, want)
		}
	})
	t.Run("DiffAssoc-score-nil", func(t *testing.T) {
		data := map[string]int{"one": 10, "two": 20, "three": 30, "four": 40, "five": 50}
		data2 := map[string]int{"one": 10, "two": 20, "three": 31, "four": 41, "five": 50}
		got := DiffAssoc(data, data2)
		want := map[string]int{"three": 30, "four": 40}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DiffAssoc() = %v, want %v", got, want)
		}
	})
}

func TestDiffKeys(t *testing.T) {
	t.Run("DiffKeys-score-nil", func(t *testing.T) {
		var data map[string]int
		var data2 map[string]int
		got := DiffKeys(data, data2)
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DiffKeys() = %v, want %v", got, want)
		}
	})
	t.Run("DiffKeys-score", func(t *testing.T) {
		data := map[string]int{"one": 10, "two": 20, "three": 30, "four": 40, "five": 50}
		data2 := map[string]int{"two": 1, "four": 2, "six": 3, "eight": 4}
		got := DiffKeys(data, data2)
		want := map[string]int{"five": 50, "one": 10, "three": 30}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DiffKeys() = %v, want %v", got, want)
		}
	})
}

func TestIntersect(t *testing.T) {
	t.Run("Intersect-score-nil", func(t *testing.T) {
		var data []string
		var data2 []string
		got := Intersect(data, data2)
		var want []string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Intersect() = %v, want %v", got, want)
		}
	})
	t.Run("Intersect-score", func(t *testing.T) {
		data := []string{"one", "two", "three", "four"}
		data2 := []string{"one", "two2", "three3", "four"}
		got := Intersect(data, data2)
		want := []string{"one", "four"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Intersect() = %v, want %v", got, want)
		}
	})
}

func TestIntersectByKeys(t *testing.T) {
	t.Run("IntersectByKeys-score-nil", func(t *testing.T) {
		var data map[string]int
		var data2 map[string]int
		got := IntersectByKeys(data, data2)
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IntersectByKeys() = %v, want %v", got, want)
		}
	})
	t.Run("IntersectByKeys-score", func(t *testing.T) {
		data := map[string]int{"one": 10, "two": 20, "three": 30, "four": 40, "five": 50}
		data2 := map[string]int{"two": 1, "four": 2, "six": 3, "eight": 4}
		got := IntersectByKeys(data, data2)
		want := map[string]int{"two": 20, "four": 40}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IntersectByKeys() = %v, want %v", got, want)
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("Min-score-false", func(t *testing.T) {
		data := []int{70, 80, 60}
		got := Min(data)
		want := 60
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Min() = %v, want %v", got, want)
		}
	})
}
func TestMinSlice(t *testing.T) {
	t.Run("MinSlice-score", func(t *testing.T) {
		data := []map[string]int{{
			"username":    10001,
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10003,
			"english":     70,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10002,
			"english":     68,
			"mathematics": 70,
			"language":    80,
		}}
		got := MinSlice(data, func(_ int, v map[string]int) int {
			return v["english"]
		})
		want := 60
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MinSlice() = %v, want %v", got, want)
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("Max-score-false", func(t *testing.T) {
		data := []int{70, 80, 60}
		got := Max(data)
		want := 80
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Max() = %v, want %v", got, want)
		}
	})
}
func TestMaxSlice(t *testing.T) {
	t.Run("MaxSlice-score", func(t *testing.T) {
		data := []map[string]int{{
			"username":    10001,
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10003,
			"english":     70,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10002,
			"english":     68,
			"mathematics": 70,
			"language":    80,
		}}
		got := MaxSlice(data, func(_ int, v map[string]int) int {
			return v["english"]
		})
		want := 80
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MaxSlice() = %v, want %v", got, want)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("Pop-score-nil", func(t *testing.T) {
		var data []int
		got, _ := Pop(data)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Pop() = %v, want %v", got, want)
		}
	})
	t.Run("Pop-score", func(t *testing.T) {
		data := []int{70, 80, 60}
		got, got1 := Pop(data)
		want := []int{70, 80}
		want1 := 60
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(got1, want1) {
			t.Errorf("Pop() = %v, want %v", got, want)
		}
	})
	t.Run("Pop-score", func(t *testing.T) {
		data := []int{}
		got, got1 := Pop(data)
		want := []int{}
		want1 := 0
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(got1, want1) {
			t.Errorf("Pop() = %v, want %v", got, want)
		}
	})
	t.Run("Pop-score", func(t *testing.T) {
		data := []int{70}
		got, got1 := Pop(data)
		want := []int{}
		want1 := 70
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(got1, want1) {
			t.Errorf("Pop() = %v, want %v", got, want)
		}
	})
}

func TestPull(t *testing.T) {
	t.Run("Pull-score-nil", func(t *testing.T) {
		var data map[string]int
		got, _ := Pull(data, "english")
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Pull() = %v, want %v", got, want)
		}
	})
	t.Run("Pull-score-nil", func(t *testing.T) {
		data := map[string]int{
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}
		got, score := Pull(data, "english")
		want := map[string]int{
			"mathematics": 70,
			"language":    80,
		}
		wantScore := 60
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(score, wantScore) {
			t.Errorf("Pull() = %v, want %v", got, want)
		}
	})
}

func TestPrepend(t *testing.T) {
	t.Run("Prepend-score-nil", func(t *testing.T) {
		var data []int
		got := Prepend(data, 50, 30)
		want := []int{50, 30}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Prepend() = %v, want %v", got, want)
		}
	})
	t.Run("Prepend-score", func(t *testing.T) {
		data := []int{70, 80, 60}
		got := Prepend(data, 50, 30)
		want := []int{50, 30, 70, 80, 60}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Prepend() = %v, want %v", got, want)
		}
	})
}

func TestPush(t *testing.T) {
	t.Run("Push-score", func(t *testing.T) {
		data := []int{70, 80, 60}
		got := Push(data, 50, 30)
		want := []int{70, 80, 60, 50, 30}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Push() = %v, want %v", got, want)
		}
	})
}

func TestPut(t *testing.T) {
	t.Run("Put-score-nil", func(t *testing.T) {
		var data map[string]int
		got := Put(data, "english", 60)
		want := map[string]int{"english": 60}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Put() = %v, want %v", got, want)
		}
	})
	t.Run("Put-score-nil", func(t *testing.T) {
		data := map[string]int{
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}
		got := Put(data, "english", 80)
		want := map[string]int{
			"english":     80,
			"mathematics": 70,
			"language":    80,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Put() = %v, want %v", got, want)
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("IsEmpty-score", func(t *testing.T) {
		var data int
		got := IsEmpty(data)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsEmpty(() = %v, want %v", got, want)
		}
	})
	t.Run("IsEmpty-score", func(t *testing.T) {
		var data float64
		got := IsEmpty(data)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsEmpty(() = %v, want %v", got, want)
		}
	})
	t.Run("IsEmpty-score", func(t *testing.T) {
		var data string
		got := IsEmpty(data)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsEmpty(() = %v, want %v", got, want)
		}
	})
	t.Run("IsEmpty-score", func(t *testing.T) {
		var data bool
		got := IsEmpty(data)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsEmpty(() = %v, want %v", got, want)
		}
	})
	t.Run("IsEmpty-score", func(t *testing.T) {
		data := 1
		got := IsEmpty(data)
		want := false
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsEmpty(() = %v, want %v", got, want)
		}
	})
	t.Run("IsEmpty-score", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("run time panic: %v", err)
			}
		}()
		data := []string{}
		got := IsEmpty(data)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsEmpty(() = %v, want %v", got, want)
		}
	})
}

func TestIsEmptySlice(t *testing.T) {
	t.Run("IsEmptySlice-score-nil", func(t *testing.T) {
		data := []string{}
		got := IsEmptySlice(data)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsEmptySlice(() = %v, want %v", got, want)
		}
	})
	t.Run("IsEmptySlice-score-nil", func(t *testing.T) {
		data := []string{"a"}
		got := IsEmptySlice(data)
		want := false
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsEmptySlice(() = %v, want %v", got, want)
		}
	})
}

func TestIsEmptyMap(t *testing.T) {
	t.Run("IsEmptyMap-score-nil", func(t *testing.T) {
		data := map[string]int{}
		got := IsEmptyMap(data)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsEmptyMap(() = %v, want %v", got, want)
		}
	})
	t.Run("IsEmptyMap-score-nil", func(t *testing.T) {
		data := map[string]int{"language": 80}
		got := IsEmptyMap(data)
		want := false
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsEmptyMap(() = %v, want %v", got, want)
		}
	})
}

func TestIsNotEmpty(t *testing.T) {
	t.Run("IsNotEmpty-score", func(t *testing.T) {
		data := 1
		got := IsNotEmpty(data)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsNotEmpty(() = %v, want %v", got, want)
		}
	})
}

func TestIsNotEmptySlice(t *testing.T) {
	t.Run("IsNotEmptySlice-score-nil", func(t *testing.T) {
		data := []string{"a"}
		got := IsNotEmptySlice(data)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsNotEmptySlice(() = %v, want %v", got, want)
		}
	})
}

func TestIsNotEmptyMap(t *testing.T) {
	t.Run("IsNotEmptyMap-score-nil", func(t *testing.T) {
		data := map[string]int{"language": 80}
		got := IsNotEmptyMap(data)
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("IsNotEmptyMap(() = %v, want %v", got, want)
		}
	})
}

func TestImplode(t *testing.T) {
	t.Run("Implode-score-nil", func(t *testing.T) {
		var data []map[string]int
		got := Implode(data, func(key int, value map[string]int) string {
			return fmt.Sprintf("username:%d,englist:%d", value["username"], value["english"])
		}, "---")
		want := ""
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Implode() = %v, want %v", got, want)
		}
	})
	t.Run("Implode-score", func(t *testing.T) {
		data := []map[string]int{{
			"username": 10001,
			"english":  60,
		}, {
			"username": 10004,
			"english":  80,
		}}
		got := Implode(data, func(key int, value map[string]int) string {
			return fmt.Sprintf("username:%d,englist:%d", value["username"], value["english"])
		}, "---")
		want := "username:10001,englist:60---username:10004,englist:80"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Implode() = %v, want %v", got, want)
		}
	})
}

func TestKeyBy(t *testing.T) {
	t.Run("KeyBy-score-nil", func(t *testing.T) {
		var data []map[string]int
		got := KeyBy(data, func(key int, value map[string]int) string {
			return fmt.Sprintf("username:%d,englist:%d", value["username"], value["english"])
		})
		var want map[string]map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("KeyBy() = %v, want %v", got, want)
		}
	})
	t.Run("KeyBy-score-nil", func(t *testing.T) {
		data := []map[string]int{{
			"username": 10001,
			"english":  60,
		}, {
			"username": 10004,
			"english":  80,
		}}
		got := KeyBy(data, func(key int, value map[string]int) string {
			return fmt.Sprintf("username:%d", value["username"])
		})
		want := map[string]map[string]int{
			"username:10001": {
				"english":  60,
				"username": 10001,
			},
			"username:10004": {
				"english":  80,
				"username": 10004,
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("KeyBy() = %v, want %v", got, want)
		}
	})
}

func TestMerge(t *testing.T) {
	t.Run("Merge-score-nil", func(t *testing.T) {
		var data []map[string]int
		got := Merge(data...)
		var want map[string]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Merge() = %v, want %v", got, want)
		}
	})
	t.Run("Merge-score-nil", func(t *testing.T) {
		data := []map[string]int{{
			"username": 10001,
			"english":  60,
		}, {
			"username1": 10004,
			"english2":  80,
		}}
		got := Merge(data...)
		want := map[string]int{
			"username":  10001,
			"english":   60,
			"username1": 10004,
			"english2":  80,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Merge() = %v, want %v", got, want)
		}
	})
}

func TestMergeRecursive(t *testing.T) {
	t.Run("MergeRecursive-score-nil", func(t *testing.T) {
		var data []map[string][]int
		got := MergeRecursive(data...)
		var want map[string][]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MergeRecursive() = %v, want %v", got, want)
		}
	})
	t.Run("MergeRecursive-score-nil", func(t *testing.T) {
		data := []map[string][]int{{
			"username": []int{10001},
			"english":  []int{60},
		}, {
			"username": []int{10004},
			"english":  []int{80},
		}}
		got := MergeRecursive(data...)

		want := map[string][]int{
			"username": {10001, 10004},
			"english":  {60, 80},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MergeRecursive() = %v, want %v", got, want)
		}
	})
}

func TestSkip(t *testing.T) {
	t.Run("Skip-score-nil", func(t *testing.T) {
		var data []int
		got := Skip(data, 3)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Skip() = %v, want %v", got, want)
		}
	})
	t.Run("Skip-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		got := Skip(data, 2)
		want := []int{3, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Skip() = %v, want %v", got, want)
		}
	})
	t.Run("Skip-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		got := Skip(data, 55)
		want := []int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Skip() = %v, want %v", got, want)
		}
	})
}

func TestSlice(t *testing.T) {
	t.Run("Slice-score-nil", func(t *testing.T) {
		var data []int
		got := Slice(data, 2, 2)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Slice() = %v, want %v", got, want)
		}
	})
	t.Run("Slice-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		got := Slice(data, 6, 2)
		want := []int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Slice() = %v, want %v", got, want)
		}
	})
	t.Run("Slice-score-nil", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		got := Slice(data, -2, 20)
		want := []int{4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Slice() = %v, want %v", got, want)
		}
	})
	t.Run("Slice-score-nil", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		got := Slice(data, 2, 20)
		want := []int{3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Slice() = %v, want %v", got, want)
		}
	})
}

func TestNth(t *testing.T) {
	t.Run("Nth-score-nil", func(t *testing.T) {
		var data []int
		got := Nth(data, 3)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Nth() = %v, want %v", got, want)
		}
	})
	t.Run("Nth-score-nil", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8}
		got := Nth(data, 3)
		want := []int{1, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Nth() = %v, want %v", got, want)
		}
	})
}

func TestPad(t *testing.T) {
	t.Run("Pad-score", func(t *testing.T) {
		data := []int{1, 2, 3}
		got := Pad(data, 2, -1)
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Pad() = %v, want %v", got, want)
		}
	})
	t.Run("Pad-score", func(t *testing.T) {
		data := []int{1, 2, 3}
		got := Pad(data, 5, -1)
		want := []int{1, 2, 3, -1, -1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Pad() = %v, want %v", got, want)
		}
	})
	t.Run("Pad-score", func(t *testing.T) {
		data := []int{1, 2, 3}
		got := Pad(data, -5, -1)
		want := []int{-1, -1, 1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Pad() = %v, want %v", got, want)
		}
	})
}

func TestPartition(t *testing.T) {
	t.Run("Partition-score", func(t *testing.T) {
		var data []int
		got, other := Partition(data, func(index int, value int) bool {
			return value > 3
		})
		var wantOther, want []int
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(wantOther, other) {
			t.Errorf("Partition() = %v, want %v", got, want)
		}
	})
	t.Run("Partition-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6}
		got, other := Partition(data, func(index int, value int) bool {
			return value > 3
		})
		want := []int{4, 5, 6}
		wantOther := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(wantOther, other) {
			t.Errorf("Partition() = %v, want %v", got, want)
		}
	})
}

func TestPluck(t *testing.T) {
	t.Run("Pluck-score", func(t *testing.T) {
		var data []map[string]int
		got := Pluck(data, "english")
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Pluck() = %v, want %v", got, want)
		}
	})
	t.Run("Pluck-score", func(t *testing.T) {
		data := []map[string]int{{
			"username":    10001,
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10003,
			"english":     70,
			"mathematics": 70,
			"language":    80,
		}, {
			"username":    10002,
			"english":     68,
			"mathematics": 70,
			"language":    80,
		}}
		got := Pluck(data, "english")
		want := []int{60, 80, 70, 68}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Pluck() = %v, want %v", got, want)
		}
	})
}

func TestRange(t *testing.T) {
	t.Run("Range-score", func(t *testing.T) {
		got := Range(3, 6)
		want := []int{3, 4, 5, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Range() = %v, want %v", got, want)
		}
	})
}

func TestForPage(t *testing.T) {
	t.Run("ForPage-score", func(t *testing.T) {
		var data []int
		got := ForPage(data, 2, 3)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ForPage() = %v, want %v", got, want)
		}
	})
	t.Run("ForPage-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7}
		got := ForPage(data, 2, 3)
		want := []int{4, 5, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ForPage() = %v, want %v", got, want)
		}
	})

}

func TestGroupBy(t *testing.T) {
	t.Run("GroupBy-score", func(t *testing.T) {
		var data []int
		got := GroupBy(data, func(key int, value int) string {
			return fmt.Sprintf("g-%d", value/3)
		})
		var want map[string][]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GroupBy() = %v, want %v", got, want)
		}
	})
	t.Run("GroupBy-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7}
		got := GroupBy(data, func(key int, value int) string {
			return fmt.Sprintf("g-%d", value%3)
		})
		want := map[string][]int{"g-0": {3, 6}, "g-1": {1, 4, 7}, "g-2": {2, 5}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GroupBy() = %v, want %v", got, want)
		}
	})
}

func TestShuffle(t *testing.T) {
	t.Run("Shuffle-score-nil", func(t *testing.T) {
		var data []int
		got := Shuffle(data)
		var want []int
		if !reflect.DeepEqual(Sort(got), want) {
			t.Errorf("Shuffle() = %v, want %v", got, want)
		}
	})
	t.Run("Shuffle-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7}
		got := Shuffle(data)
		want := []int{1, 2, 3, 4, 5, 6, 7}
		if !reflect.DeepEqual(Sort(got), want) {
			t.Errorf("Shuffle() = %v, want %v", got, want)
		}
	})

}

func TestRandom(t *testing.T) {
	t.Run("Random-score-nil", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7}
		got := Random(data, 2)
		want := 2
		if !reflect.DeepEqual(len(got), want) {
			t.Errorf("Random() = %v, want %v", got, want)
		}
	})
	t.Run("Random-score-nil", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7}
		got := Random(data, 10)
		want := 7
		if !reflect.DeepEqual(len(got), want) {
			t.Errorf("Random() = %v, want %v", got, want)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("Reduce-score-nil", func(t *testing.T) {
		var data []int
		got := Reduce(data, func(s string, i int, item int) string {
			return fmt.Sprintf("%s-%d", s, item)
		})
		var want string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Reduce() = %v, want %v", got, want)
		}
	})
	t.Run("Reduce-score-nil", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7}
		got := Reduce(data, func(s string, i int, item int) string {
			return fmt.Sprintf("%s-%d", s, item)
		})
		want := "-1-2-3-4-5-6-7"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Reduce() = %v, want %v", got, want)
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("Reverse-score-nil", func(t *testing.T) {
		var data map[string]int
		got := Reverse(data)
		var want map[int]string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Reverse() = %v, want %v", got, want)
		}
	})
	t.Run("Reverse-score-nil", func(t *testing.T) {
		data := map[string]int{
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    81,
		}
		got := Reverse(data)
		want := map[int]string{10004: "username", 70: "mathematics", 80: "english", 81: "language"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Reverse() = %v, want %v", got, want)
		}
	})
}

func TestSearchSlice(t *testing.T) {
	t.Run("SearchSlice-score-nil", func(t *testing.T) {
		var data []string
		got, ok := SearchSlice(data, "english")
		var want int
		var wantOk bool
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(wantOk, ok) {
			t.Errorf("SearchSlice() = %v, want %v", got, want)
		}
	})
	t.Run("SearchSlice-score-nil", func(t *testing.T) {
		data := []string{"a", "b", "c"}
		got, ok := SearchSlice(data, "b")
		want := 1
		wantOk := true
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(wantOk, ok) {
			t.Errorf("SearchSlice() = %v, want %v", got, want)
		}
	})
	t.Run("SearchSlice-score-nil", func(t *testing.T) {
		data := []string{"a", "b", "c"}
		got, ok := SearchSlice(data, "ab")
		want := 0
		wantOk := false
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(wantOk, ok) {
			t.Errorf("SearchSlice() = %v, want %v", got, want)
		}
	})
}

func TestSearchMap(t *testing.T) {
	t.Run("SearchMap-score-nil", func(t *testing.T) {
		var data map[string]int
		got, ok := SearchMap(data, 12)
		var want string
		var wantOk bool
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(wantOk, ok) {
			t.Errorf("SearchMap() = %v, want %v", got, want)
		}
	})
	t.Run("SearchMap-score-nil", func(t *testing.T) {
		data := map[string]int{
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    81,
		}
		got, ok := SearchMap(data, 80)
		want := "english"
		wantOk := true
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(wantOk, ok) {
			t.Errorf("SearchMap() = %v, want %v", got, want)
		}
	})
	t.Run("SearchMap-score-nil", func(t *testing.T) {
		data := map[string]int{
			"username":    10004,
			"english":     80,
			"mathematics": 70,
			"language":    81,
		}
		got, ok := SearchMap(data, 100)
		want := ""
		wantOk := false
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(wantOk, ok) {
			t.Errorf("SearchMap() = %v, want %v", got, want)
		}
	})
}

func TestShift(t *testing.T) {
	t.Run("Shift-score-nil", func(t *testing.T) {
		var data []int
		got, _ := Shift(data)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Shift() = %v, want %v", got, want)
		}
	})
	t.Run("Shift-score", func(t *testing.T) {
		data := []int{70, 80, 60}
		got, got1 := Shift(data)
		want := []int{80, 60}
		want1 := 70
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(got1, want1) {
			t.Errorf("Shift() = %v, want %v", got, want)
		}
	})
	t.Run("Shift-score", func(t *testing.T) {
		data := []int{}
		got, got1 := Shift(data)
		want := []int{}
		want1 := 0
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(got1, want1) {
			t.Errorf("Shift() = %v, want %v", got, want)
		}
	})
	t.Run("Shift-score", func(t *testing.T) {
		data := []int{70}
		got, got1 := Shift(data)
		want := []int{}
		want1 := 70
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(got1, want1) {
			t.Errorf("Shift() = %v, want %v", got, want)
		}
	})
}

func TestSliding(t *testing.T) {
	t.Run("Sliding-score-nil", func(t *testing.T) {
		var data []int
		got := Sliding(data, 3, 1)
		var want [][]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Sliding() = %v, want %v", got, want)
		}
	})
	t.Run("Sliding-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7}
		got := Sliding(data, 9, 3)
		var want [][]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Sliding() = %v, want %v", got, want)
		}
	})
	t.Run("Sliding-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7}
		got := Sliding(data, 3, 4)
		want := [][]int{{1, 2, 3}, {5, 6, 7}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Sliding() = %v, want %v", got, want)
		}
	})
}

func TestTimes(t *testing.T) {
	t.Run("Times-score", func(t *testing.T) {
		got := Times(-3, func(index int) string {
			return fmt.Sprintf("%d", index)
		})
		var want []string
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Times() = %v, want %v", got, want)
		}
	})
}

func TestSplice(t *testing.T) {
	t.Run("Splice-score-nil", func(t *testing.T) {
		var data []int
		got, other := Splice(data, 3, 1)
		var want []int
		var wantOther []int
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(other, wantOther) {
			t.Errorf("Sliding() = %v, want %v", got, want)
		}
	})
	t.Run("Splice-score", func(t *testing.T) {
		data := []string{"a", "b", "c", "d", "e"}
		got, other := Splice(data, 2, 2)
		want := []string{"c", "d"}
		wantOther := []string{"a", "b", "e"}
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(other, wantOther) {
			t.Errorf("Sliding() = %v, want %v", got, want)
		}
	})
	t.Run("Splice-score", func(t *testing.T) {
		data := []string{"a", "b", "c", "d", "e"}
		got, other := Splice(data, 2, 30)
		want := []string{"c", "d", "e"}
		wantOther := []string{"a", "b"}
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(other, wantOther) {
			t.Errorf("Sliding() = %v, want %v", got, want)
		}
	})
	t.Run("Splice-score", func(t *testing.T) {
		data := []string{"a", "b", "c", "d", "e"}
		got, other := Splice(data, -2, 30)
		want := []string{"d", "e"}
		wantOther := []string{"a", "b", "c"}
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(other, wantOther) {
			t.Errorf("Sliding() = %v, want %v", got, want)
		}
	})
	t.Run("Splice-score", func(t *testing.T) {
		data := []string{"a", "b", "c", "d", "e"}
		got, other := Splice(data, -20, 30)
		want := []string{}
		wantOther := []string{"a", "b", "c", "d", "e"}
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(other, wantOther) {
			t.Errorf("Sliding() = %v, want %v", got, want)
		}
	})
}

func TestSplit(t *testing.T) {
	t.Run("Split-score-nil", func(t *testing.T) {
		var data []int
		got := Split(data, 4)
		var want [][]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Split() = %v, want %v", got, want)
		}
	})
	t.Run("Split-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := Split(data, 2)
		want := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Split() = %v, want %v", got, want)
		}
	})
	t.Run("Split-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := Split(data, 4)
		want := [][]int{{1, 2, 3}, {4, 5}, {6, 7}, {8, 9}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Split() = %v, want %v", got, want)
		}
	})
}

func TestToJson(t *testing.T) {
	t.Run("ToJson-score", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("run time panic: %v", err)
			}
		}()
		data := TestToJson
		got := ToJson(data)
		want := "[1,2,3,4,5,6,7,8,9]"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ToJson() = %v, want %v", got, want)
		}
	})
	t.Run("ToJson-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := ToJson(data)
		want := "[1,2,3,4,5,6,7,8,9]"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ToJson() = %v, want %v", got, want)
		}
	})
}

func TestUnion(t *testing.T) {
	t.Run("Union-score", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		got := Union(data, data, data)
		want := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Union() = %v, want %v", got, want)
		}
	})
}

func TestUnique(t *testing.T) {
	t.Run("Unique-score-nil", func(t *testing.T) {
		var data []int
		got := Unique(data, func(i int, v int) bool {
			return i > 0
		})
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unique() = %v, want %v", got, want)
		}
	})
	t.Run("Unique-score", func(t *testing.T) {
		data := []string{"a", "b", "c", "d", "a", "c", "e"}
		got := Unique(data, func(_ int, v string) string {
			return v
		})
		want := []string{"b", "d", "a", "c", "e"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unique() = %v, want %v", got, want)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add-score", func(t *testing.T) {
		data := map[string]int{
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}
		got := Add(data, "english", 100)
		want := map[string]int{
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Add() = %v, want %v", got, want)
		}
	})
	t.Run("Add-score", func(t *testing.T) {
		data := map[string]int{
			"english":     60,
			"mathematics": 70,
			"language":    80,
		}
		got := Add(data, "english-plus", 100)
		want := map[string]int{
			"english":      60,
			"mathematics":  70,
			"language":     80,
			"english-plus": 100,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Add() = %v, want %v", got, want)
		}
	})
	t.Run("Add-score", func(t *testing.T) {
		var data map[string]int
		got := Add(data, "english-plus", 100)
		want := map[string]int{
			"english-plus": 100,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Add() = %v, want %v", got, want)
		}
	})
}

func TestWrap(t *testing.T) {
	t.Run("Wrap-score", func(t *testing.T) {
		data := 1
		got := Wrap[int, int](data)
		want := []int{1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wrap() = %v, want %v", got, want)
		}
	})
	t.Run("Wrap-score", func(t *testing.T) {
		data := []int{1}
		got := Wrap[[]int, int](data)
		want := []int{1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wrap() = %v, want %v", got, want)
		}
	})
	t.Run("Wrap-score", func(t *testing.T) {
		var data map[string]int
		got := Wrap[map[string]int, map[string]int](data)
		want := []map[string]int{nil}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wrap() = %v, want %v", got, want)
		}
	})
}

func TestZip(t *testing.T) {
	t.Run("Zip-score-nil", func(t *testing.T) {
		var data []int
		got := Zip(data, data)
		var want [][]int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Zip() = %v, want %v", got, want)
		}
	})
	t.Run("Zip-score-nil", func(t *testing.T) {
		data := []int{1, 2, 3}
		data2 := []int{4, 5, 6}
		got := Zip(data, data2)
		want := [][]int{{1, 4}, {2, 5}, {3, 6}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Zip() = %v, want %v", got, want)
		}
	})
	t.Run("Zip-score-nil", func(t *testing.T) {
		data := []int{1, 2, 3, 7, 8}
		data2 := []int{4, 5, 6}
		got := Zip(data, data2)
		want := [][]int{{1, 4}, {2, 5}, {3, 6}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Zip() = %v, want %v", got, want)
		}
	})
}
