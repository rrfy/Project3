package main

import "fmt"

type AbstractFactory interface {
  CreateProductA() ProductA
  CreateProductB() ProductB
}

type ProductA interface {
  Use()
}

type ProductB interface {
  Operate()
}

type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProductA() ProductA {
  return &ConcreteProductA1{}
}

func (f *ConcreteFactory1) CreateProductB() ProductB {
  return &ConcreteProductB1{}
}

type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProductA() ProductA {
  return &ConcreteProductA2{}
}

func (f *ConcreteFactory2) CreateProductB() ProductB {
  return &ConcreteProductB2{}
}

type ConcreteProductA1 struct{}

func (p *ConcreteProductA1) Use() {
  fmt.Println("Using ConcreteProductA1")
}

type ConcreteProductB1 struct{}

func (p *ConcreteProductB1) Operate() {
  fmt.Println("Operating ConcreteProductB1")
}

type ConcreteProductA2 struct{}

func (p *ConcreteProductA2) Use() {
  fmt.Println("Using ConcreteProductA2")
}

type ConcreteProductB2 struct{}

func (p *ConcreteProductB2) Operate() {
  fmt.Println("Operating ConcreteProductB2")
}

func main() {
  factory1 := &ConcreteFactory1{}
  productA1 := factory1.CreateProductA()
  productB1 := factory1.CreateProductB()

  productA1.Use()
  productB1.Operate()

  factory2 := &ConcreteFactory2{}
  productA2 := factory2.CreateProductA()
  productB2 := factory2.CreateProductB()

  productA2.Use()
  productB2.Operate()
}