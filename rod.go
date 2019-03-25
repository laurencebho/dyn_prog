package main
import (
    "fmt"
    "math/rand"
    "sort"
    "time"
)

type Rod struct {
    price, best int
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func get_random_prices(n, pmax int) []int {
    var prices []int
    var rp int
    for i:=0; i<n; i++ {
        rp = rand.Intn(pmax) + 1
        prices = append(prices, rp)
    }

    sort.Ints(prices)
    return prices
}

func get_rods(n, pmax int) []Rod {
    var rods []Rod
    prices := get_random_prices(n, pmax)
    for i:=0; i<n; i++ {
        r := Rod{prices[i], -1}
        rods = append(rods, r)
    }
    return rods
}


func top_down(rodsp *[]Rod, n int) int {
    rods := *rodsp
    var q int

    if n==0 {
        return 0
    }

    if rods[n-1].best >= 0 { //if done
        return rods[n-1].best
    }

    q = -1
    for i:=0; i<n; i++ {
        q = max(q, rods[i].price + top_down(rodsp, n-i-1))
    }
    rods[n-1].best = q
    return q
}


func bottom_up(rods []Rod, n int) int {
    var q int

    for i:=0; i<n; i++ {
        q = -1
        for j:=0; j<=i; j++ {
            var b int
            if i-j-1 < 0 {
                b = 0
            } else {
                b = rods[i-j-1].best
            }

            q = max(q, rods[j].price + b)
            //fmt.Println(i, j, q)
        }
        rods[i].best = q
    }
    return rods[n-1].best
}

func main() {
    //random seed
    rand.Seed(time.Now().UTC().UnixNano())

    n := 18
    pmax := 100
    var rods []Rod = get_rods(n, pmax)

    fmt.Println("Rod size:", n, "  Max price:", pmax)
    fmt.Println("Prices for each length:")
    for i, r := range rods {
        fmt.Print(i+1, "=", r.price, "  ")
    }
    fmt.Println()

    td_best := top_down(&rods, n)
    bu_best := bottom_up(rods, n)

    fmt.Println("Top down best:", td_best)
    fmt.Println("Bottom up best:", bu_best)
}
