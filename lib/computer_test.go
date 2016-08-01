package computer

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestComputer(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Computer")
}

var _ = Describe("Computer", func() {
    It("should be return snake head", func() {
        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
        }
        Expect(getSnakeHead(snake)).To(Equal([2]int{0, 2}))
    })

    It("should check empty cell", func() {
        grid := [][]int{
            {1, 1, 1, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 2},
        }

        Expect(isEmptyCell(grid, [2]int{0, 2})).To(Equal(false))
        Expect(isEmptyCell(grid, [2]int{2, 2})).To(Equal(true))
    })

    It("should check outside bounding box", func() {
        grid := [][]int{
            {1, 1, 1, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 2},
        }

        Expect(isOutsideBoundingBox([2]int{-1, 0}, grid)).To(Equal(true))
        Expect(isOutsideBoundingBox([2]int{0, 5}, grid)).To(Equal(true))
        Expect(isOutsideBoundingBox([2]int{0, 0}, grid)).To(Equal(false))
    })

    It("should return adjacent position", func() {
        Expect(getAdjacentPosition([2]int{1, 0}, up)).To(Equal([2]int{0, 0}))
        Expect(getAdjacentPosition([2]int{0, 0}, right)).To(Equal([2]int{0, 1}))
        Expect(getAdjacentPosition([2]int{0, 0}, down)).To(Equal([2]int{1, 0}))
        Expect(getAdjacentPosition([2]int{0, 1}, left)).To(Equal([2]int{0, 0}))
    })

    It("should return possible moves", func() {
        grid := [][]int{
            {1, 1, 1, 0, 0},
            {0, 0, 1, 0, 0},
            {2, 1, 1, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
        }

        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
            {1, 2},
            {2, 2},
            {2, 1},
        }

        Expect(getPossibleMoves(grid, snake)).To(Equal([]int{0, 2, 3}))
    })

    It("should check if snake has free space", func() {
        grid := [][]int{
            {1, 1, 1, 0, 0},
            {1, 1, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
        }

        snake := [][2]int{
            {0, 2}, {0, 1}, {1, 1}, {1, 0}, {0, 0},
        }

        Expect(isSnakeHasFreeSpace(grid, snake)).To(Equal(false))
    })

    It("should move snake", func() {
        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
        }

        apple := [2]int{4, 0}

        newSnake := moveSnake(snake, apple, []int{right, right, down})

        Expect(newSnake).To(Equal([][2]int{
            {0, 3},
            {0, 4},
            {1, 4},
        }))
    })

    It("should initialize grid", func() {
        snake := [][2]int{
            {2, 0},
            {2, 1},
            {2, 2},
        }

        apple := [2]int{4, 4}

        grid := initializeGrid(5, 5, snake, apple)
        Expect(grid).To(Equal([][]int{
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0},
            {1, 1, 1, 0, 0},
            {0, 0, 0, 0, 0},
            {0, 0, 0, 0, 2},
        }))
    })

    /*
        {1, 1, 1, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 2, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0},
    */
    It("should find path", func() {
        snake := [][2]int{
            {0, 0},
            {0, 1},
            {0, 2},
        }

        apple := [2]int{2, 2}

        path := GetPath(5, 5, snake, apple)
        Expect(path[:2]).To(Equal([]int{down, down}))
    })

    /*
        {1, 1, 1, 1, 0},
        {1, 0, 0, 0, 0},
        {1, 0, 0, 0, 0},
        {1, 1, 1, 1, 1},
        {0, 2, 0, 0, 0},
    */
    It("should not enter in closed zone", func() {
        snake := [][2]int{
            {0, 3}, {0, 2}, {0, 1}, {0, 0}, {1, 0}, {2, 0}, {3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4},
        }

        apple := [2]int{4, 1}

        path := GetPath(5, 5, snake, apple)
        Expect(path[:1][0]).To(Equal(up))
    })

    /*
        {0, 0, 0, 1, 2},
        {0, 0, 0, 1, 1},
        {0, 0, 0, 0, 1},
        {0, 0, 0, 0, 1},
        {0, 0, 0, 0, 0},
    */
    It("should not eat apple if no free space", func() {
        snake := [][2]int{
            {3, 4}, {2, 4}, {1, 4}, {1, 3}, {0, 3},
        }

        apple := [2]int{0, 4}

        path := GetPath(5, 5, snake, apple)
        Expect(path[:1][0]).To(Equal(left))
    })
})
