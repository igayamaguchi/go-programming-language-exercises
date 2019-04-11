package sample

type Sample struct {
  ID int // 公開される
  name string // 公開されない
}

func (sample *Sample) SetName(name string) {
  sample.name = name
}

func (sample Sample) GetName() string {
  return sample.name
}

type tree struct {
  value int
  left, right *tree
}
