package config

import (
	"log"
	"math/big"
	"os"
)

var (
	home              = GetUserHome()
	Base              = home + "/.primes/"
	Directory         = Base + "directory.txt"
	configurationFile = home + "/.primegenerator.yaml"

	LocalConfig   = Config{}
	StartingPrime string
	MaxFilesize   int
	MaxBufferSize int
	ShowFails     bool
	Host          string

	Port                 = "8080"
	Address              string
	AssignmentPoint      = "/"
	ReturnPoint          = "/finished"
	HeavyAssignmentPoint = "/heavy"
	HeavyReturnPoint     = "/heavy/finished"

	Id                 uint64
	LastPrimeGenerated *big.Int

	Logger = log.New(os.Stdout, "", log.LstdFlags)
)
