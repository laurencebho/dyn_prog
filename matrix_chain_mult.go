package main
import (
    "fmt"
    "math/rand"
    "time"
)

type Mult struct {
    num, split int
}

func get_problem(n, min, max int) []int {
    var A []int
    for i:=0; i<n+1; i++ {
        A = append(A, rand.Intn(max - min + 1) + min)
    }
    fmt.Println(A)
    return A
}

func make_matrix(n int) [8][8]Mult {
    var M [8][8]Mult //n x n matrix
    for i:=0; i<n; i++ {
        for j:=0; j<n; j++ {
            M[i][j] = Mult{0, 0}
        }
    }
    return M
}

func print_solution(M [8][8]Mult, i, j int) {
    if i==j {
        fmt.Print(i)
    } else {
        fmt.Print("(")
        print_solution(M, i, M[i][j].split)
        if j+1 < len(M) {
            print_solution(M, M[i][j+1].split, j)
        } else {
            fmt.Println(j+1)
        }
        fmt.Print(")")
    }
}
func bottom_up(n, min, max int) [8][8]Mult {
    A := get_problem(n, min, max)
    var M [8][8]Mult = make_matrix(n) //n x n matrix

    for l:=1; l<n; l++ { //l is the splitting point - finding best one
        for i:=0; i<n-l; i++ { //max value is n-1
            j := i + l //from l to n-1
            M[i][j].num = 100000 //initialise to large number
            fmt.Println("initialising", i, j)
            fmt.Println("Product of ", j-i+1, i, j-1)
            for k:=i; k<j; k++ {
                fmt.Println(M[i][j])
                fmt.Println(M[i][k])
                fmt.Println(M[k+1][j])
                fmt.Println(i, k+1)
                q := M[i][k].num + M[k+1][j].num + A[i] * A[k+1] * A[j+1]
                if q < M[i][j].num {
                    M[i][j].num = q
                    M[i][j].split = k
                }
            }
        }
    }
    return M
}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

    M := bottom_up(8, 3, 8)
    fmt.Println(M)
    print_solution(M, 0, 7)
}
