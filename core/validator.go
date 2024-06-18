package core

import (
	"math/rand"
	"time"
)

type Validator struct {
	Address string
	Stake   int
}
type ValidatorManager struct {
	Validators []*Validator
	TotalStake int
}

func NewValidatorManager() *ValidatorManager {
	return &ValidatorManager{Validators: []*Validator{}}
}
func (vm *ValidatorManager) AddValidator(address string, stake int) {
	validator := &Validator{Address: address, Stake: stake}
	vm.Validators = append(vm.Validators, validator)
	vm.TotalStake += stake
}
func (vm *ValidatorManager) SelectValidator() *Validator {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(vm.TotalStake)
	sum := 0
	for _, v := range vm.Validators {
		sum += v.Stake
		if sum > r {
			return v
		}
	}
	return nil
}
