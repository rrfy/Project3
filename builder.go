package main

import "fmt"

type Product struct {
    Part1 string
    Part2 string
    Part3 string
}

type Builder struct {
    product *Product
}

func NewBuilder() *Builder {
    return &Builder{product: &Product{}}
}

func (b *Builder) SetPart1(part1 string) *Builder {
    b.product.Part1 = part1
    return b
}

func (b *Builder) SetPart2(part2 string) *Builder {
    b.product.Part2 = part2
    return b
}

func (b *Builder) SetPart3(part3 string) *Builder {
    b.product.Part3 = part3
    return b
}

func (b *Builder) Build() *Product {
    return b.product
}

func main() {
    product := NewBuilder().
        SetPart1("Part 1").
        SetPart2("Part 2").
        SetPart3("Part 3").
        Build()

    fmt.Printf("Product: %+v\n", product)
}