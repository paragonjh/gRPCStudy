package sample

import "gitlab.com/techschool/pcbook/pb"

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 6)

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &pb.Memory{
			Value: uint64(memGB),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return gpu
}

func NewRam() *pb.Memory {
	memGB := randomInt(4, 64)

	ram := &pb.Memory{
		Value: uint64(memGB),
		Unit:  pb.Memory_GIGABYTE,
	}
	return ram
}

func NewSSD() *pb.Storage {
	memGB := randomInt(128, 1024)

	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(memGB),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return ssd
}

func NewHDD() *pb.Storage {
	memTB := randomInt(1, 6)
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(memTB),
			Unit:  pb.Memory_TERABYTE,
		},
	}

	return hdd
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRam(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
	}
	return laptop
}

func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}
