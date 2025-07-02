package main

import (
	"fmt"

	"os"
	"papa/mathutils"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main(){
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	
	log.Debug().Msg(("DREGGTGHTHGNGHTRHTHTGH"))
	log.Info().Msg(("D54545454fdfdfd44hggHTGH"))
	log.Warn().Msg(("D54545454fdfdfd44hggHTGH"))
	log.Error().Msg(("D54545454fdfdfd44hggHTGH"))
	log.Fatal().Msg(("D54545454fdfdfd44hggHTGH"))
	fmt.Println(mathutils.PI)
	fmt.Println(mathutils.Subtract(56, 6))
}