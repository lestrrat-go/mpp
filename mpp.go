package mpp

import (
	"github.com/pkg/errors"
)

type Step interface {
	Do(interface{}) (interface{}, error)
}

type Processor struct {
	src   interface{}
	steps []Step
}

func New(src interface{}) *Processor {
	return &Processor{
		src: src,
	}
}

func (p *Processor) Step(s Step) *Processor {
	p.steps = append(p.steps, s)
	return p
}

func (p *Processor) JSON() *Processor {
	return p.Step(NewJSON())
}

func (p *Processor) Do() (interface{}, error) {
	result := p.src
	var err error
	for i, step := range p.steps {
		result, err = step.Do(result)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to run step #%d: %s`, i+1, step)
		}
	}
	return result, nil
}
