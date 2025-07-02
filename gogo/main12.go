package main

import (
	"fmt"
	"math/rand/v2"
)
func main(){
	mathutils.Add()

fmt.Println(Attack())

attackWithDamageBoost := DamageBoostDecorator(Attack)
fmt.Println(attackWithDamageBoost())

attackWithCriticalHit := CriticalHitDecorator(Attack)
fmt.Println(attackWithCriticalHit())

attackWithSlowEffect :=  SlowEffectDecorator(Attack)
fmt.Println(attackWithSlowEffect ())

fmt.Println(attackWithSlowEffect())
}

func Attack() string {
	return"Атака выполнена!"
}

func DamageBoostDecorator(attackFunc func() string) func() string  {
	
	return func()string{
     return "Вам улыбнулась удача, нанесение урона увеличено на 10%! " + attackFunc()
	}	
}

func CriticalHitDecorator(attackFunc func() string) func() string{
	return func()string{

		if rand.Float64() < 0.25 {
     return "Критический удар! Урон удвоен! " + attackFunc()
		}
		return attackFunc()   
	}
}

func SlowEffectDecorator(attackFunc func() string) func() string {
	return func()string{
		     return   attackFunc() + " Цель замедлена на 2 хода!"
	}
}
