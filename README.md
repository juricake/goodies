# Goodies
A module of Go essential structures and methods I wish were included in the std libs.

# Examples
The following examples will assume the existence of the Item structure. 
``` go
type Item struct {
	ID    string
	Value int
}
```

## Sets
``` go 
values := []int{1, 2, 3, 3, 2, 1}
s := set.New[int](values...)

fmt.Println(s.Size()) // output: 3
s.Put(5)
fmt.Println(s.Contains(3)) // output: true
s.Del(3)
fmt.Println(s.Contains(3)) // output: false
```

## Maps
``` go 
items := []Item{
		{ID: "one", Value: 10},
		{ID: "two", Value: 20},
		{ID: "three", Value: 30},
}

itemsByID := dict.New[string, Item]()
for _, item := range items {
    itemsByID.Put(item.ID, item)
}

fmt.Println(itemsByID.KeysSorted(alphabetically))   // out: [one three two]
fmt.Println(itemsByID.KeysSortedByValue(descValue)) // out: [three two one]
fmt.Println(itemsByID.Contains("seven"))            // out: false
fmt.Println(itemsByID.Size())                       // out: 3
itemsByID.Del("one")
fmt.Println(itemsByID.Get("two").Value) // out: 20

func alphabetically(a, b string) bool {
	return a < b
}

func descValue(a, b Item) bool {
	return a.Value > b.Value
}

```

## Multi-valued Maps
``` go 
items := []Item{
    {ID: "one", Value: 10},
    {ID: "two", Value: 20},
    {ID: "two", Value: 22},
    {ID: "two", Value: 23},
    {ID: "three", Value: 30},
    {ID: "three", Value: 35},
}

itemsByID := dict.NewSliceMap[string, int]()
for _, item := range items {
    itemsByID.Append(item.ID, item.Value)
}

fmt.Println(itemsByID.KeysSorted(alphabetically))        // out: [one three two]
fmt.Println(itemsByID.KeysSortedByValue(descValueCount)) // out: [two three one]
fmt.Println(itemsByID.Contains("seven"))                 // out: false
fmt.Println(itemsByID.ContainsWithin("two", 20))         // out: true
fmt.Println(itemsByID.Size())                            // out: 3
itemsByID.Del("one")
fmt.Println(itemsByID.Get("one")) // out: []
itemsByID.Put("two", []int{100, 101})
fmt.Println(itemsByID.Get("two")) // out: [100, 101]

func alphabetically(a, b string) bool {
	return a < b
}

func descValueCount(a, b []int) bool {
	return len(a) > len(b)
}

```

## transform.ToMap
``` go 
items := []Item{
    {ID: "one", Value: 10},
    {ID: "two", Value: 20},
    {ID: "three", Value: 30},
}

itemsByID := transform.ToMap[string, Item](items, getItemKey)
fmt.Println(itemsByID.Size()) // out: 3

func getItemKey(i Item) string {
	return i.ID
}
```

## transform.ToSliceMap
``` go 
items := []Item{
    {ID: "one", Value: 10},
    {ID: "two", Value: 20},
    {ID: "three", Value: 30},
    {ID: "three", Value: 35},
}

itemsByID := transform.ToSliceMap[string, Item](items, getItemKey)
fmt.Println(itemsByID.Size()) // out: 3

func getItemKey(i Item) string {
	return i.ID
}
```

## split.Split
``` go 
items := []Item{
    {ID: "one", Value: 10},
    {ID: "two", Value: 20},
    {ID: "three", Value: 30},
    {ID: "four", Value: 40},
}

bigger, smaller := split.Split[Item](items, valueBiggerThan(25))
fmt.Println(bigger)  // out: [{three 30} {four 40}]
fmt.Println(smaller) // out: [{one 10} {two 20}]

func valueBiggerThan(threshold int) func(item Item) bool {
	return func(item Item) bool {
		return item.Value > threshold
	}
}
```
