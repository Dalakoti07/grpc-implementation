package sample

import (
	"github.com/golang/protobuf/ptypes"
	"pc_book/pb/pc_book"
)

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *pc_book.Keyboard {
	keyboard := &pc_book.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

// NewCPU returns a new sample CPU
func NewCPU() *pc_book.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pc_book.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *pc_book.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 6)

	gpu := &pc_book.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &pc_book.Memory{
			Value: uint64(memGB),
			Unit:  pc_book.Memory_GIGABYTE,
		},
	}

	return gpu
}

// NewRAM returns a new sample RAM
func NewRAM() *pc_book.Memory {
	memGB := randomInt(4, 64)

	ram := &pc_book.Memory{
		Value: uint64(memGB),
		Unit:  pc_book.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a new sample SSD
func NewSSD() *pc_book.Storage {
	memGB := randomInt(128, 1024)

	ssd := &pc_book.Storage{
		Driver: pc_book.Storage_SSD,
		Memory: &pc_book.Memory{
			Value: uint64(memGB),
			Unit:  pc_book.Memory_GIGABYTE,
		},
	}

	return ssd
}

// NewHDD returns a new sample HDD
func NewHDD() *pc_book.Storage {
	memTB := randomInt(1, 6)

	hdd := &pc_book.Storage{
		Driver: pc_book.Storage_HDD,
		Memory: &pc_book.Memory{
			Value: uint64(memTB),
			Unit:  pc_book.Memory_TERABYTE,
		},
	}

	return hdd
}

// NewScreen returns a new sample Screen
func NewScreen() *pc_book.Screen {
	screen := &pc_book.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// NewLaptop returns a new sample Laptop
func NewLaptop() *pc_book.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pc_book.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*pc_book.GPU{NewGPU()},
		Storages: []*pc_book.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pc_book.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   ptypes.TimestampNow(),
	}

	return laptop
}

// RandomLaptopScore returns a random laptop score
func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}
