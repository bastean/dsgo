package model

type UseCase[Input, Output any] interface {
	Run(Input) (Output, error)
}
