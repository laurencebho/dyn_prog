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
    fmt.Println("Matrix chain dimensions:", A)
    return A
}

func make_matrix(n int) [][]Mult {
    M := make([][]Mult, n) //n x n matrix
    for i:=0; i<n; i++ {
        M[i] = make([]Mult, n)
    }
    return M
}

func print_solution(M [][]Mult, i, j int) {
    if i==j {
        fmt.Print(i)
    } else {
        fmt.Print("(")
        print_solution(M, i, M[i][j].split)
        print_solution(M, M[i][j].split + 1, j)
        fmt.Print(")")
    }
}

func bottom_up(n, min, max int) [][]Mult {
    A := get_problem(n, min, max)
    M := make_matrix(n) //n x n matrix

    for l:=1; l<n; l++ { //l = chain length, builds up from 2 to n
        for i:=0; i<n-l; i++ { //i = chain start point
            j := i + l //j = chain end point
            M[i][j].num = -1
            for k:=i; k<j; k++ {
                q := M[i][k].num + M[k+1][j].num + A[i] * A[k+1] * A[j+1]
                if q < M[i][j].num || M[i][j].num == -1 {
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

    //number of matrices in chain, min dimensions, max dimensions
    n := 10
    min := 3
    max := 12

    M := bottom_up(n, min, max)
    fmt.Println("Fewest multiplications:", M[0][n-1].num)
    fmt.Print("Parenthesisation: ")
    print_solution(M, 0, n-1)
    fmt.Println()
}
