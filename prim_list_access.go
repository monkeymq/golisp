// Copyright "2"014 SteelSeries ApS.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This package implements a basic LISP interpretor for embedding in a go program for scripting.
// This file contains the list access primitive functions.

package golisp

import (
	"errors"
	"fmt"
	"unsafe"
)

func RegisterListAccessPrimitives() {
	MakePrimitiveFunction("car", "1", CarImpl)
	MakePrimitiveFunction("head", "1", CarImpl)
	MakePrimitiveFunction("cdr", "1", CdrImpl)
	MakePrimitiveFunction("rest", "1", CdrImpl)
	MakePrimitiveFunction("tail", "1", CdrImpl)

	MakePrimitiveFunction("caar", "1", CaarImpl)
	MakePrimitiveFunction("cadr", "1", CadrImpl)
	MakePrimitiveFunction("cdar", "1", CdarImpl)
	MakePrimitiveFunction("cddr", "1", CddrImpl)

	MakePrimitiveFunction("caaar", "1", CaaarImpl)
	MakePrimitiveFunction("caadr", "1", CaadrImpl)
	MakePrimitiveFunction("cadar", "1", CadarImpl)
	MakePrimitiveFunction("caddr", "1", CaddrImpl)
	MakePrimitiveFunction("cdaar", "1", CdaarImpl)
	MakePrimitiveFunction("cdadr", "1", CdadrImpl)
	MakePrimitiveFunction("cddar", "1", CddarImpl)
	MakePrimitiveFunction("cdddr", "1", CdddrImpl)

	MakePrimitiveFunction("caaaar", "1", CaaaarImpl)
	MakePrimitiveFunction("caaadr", "1", CaaadrImpl)
	MakePrimitiveFunction("caadar", "1", CaadarImpl)
	MakePrimitiveFunction("caaddr", "1", CaaddrImpl)
	MakePrimitiveFunction("cadaar", "1", CadaarImpl)
	MakePrimitiveFunction("cadadr", "1", CadadrImpl)
	MakePrimitiveFunction("caddar", "1", CaddarImpl)
	MakePrimitiveFunction("cadddr", "1", CadddrImpl)
	MakePrimitiveFunction("cdaaar", "1", CdaaarImpl)
	MakePrimitiveFunction("cdaadr", "1", CdaadrImpl)
	MakePrimitiveFunction("cdadar", "1", CdadarImpl)
	MakePrimitiveFunction("cdaddr", "1", CdaddrImpl)
	MakePrimitiveFunction("cddaar", "1", CddaarImpl)
	MakePrimitiveFunction("cddadr", "1", CddadrImpl)
	MakePrimitiveFunction("cdddar", "1", CdddarImpl)
	MakePrimitiveFunction("cddddr", "1", CddddrImpl)
	MakePrimitiveFunction("general-car-cdr", "2", GeneralCarCdrImpl)

	MakePrimitiveFunction("first", "1", FirstImpl)
	MakePrimitiveFunction("second", "1", SecondImpl)
	MakePrimitiveFunction("third", "1", ThirdImpl)
	MakePrimitiveFunction("fourth", "1", FourthImpl)
	MakePrimitiveFunction("fifth", "1", FifthImpl)
	MakePrimitiveFunction("sixth", "1", SixthImpl)
	MakePrimitiveFunction("seventh", "1", SeventhImpl)
	MakePrimitiveFunction("eighth", "1", EighthImpl)
	MakePrimitiveFunction("ninth", "1", NinthImpl)
	MakePrimitiveFunction("tenth", "1", TenthImpl)
	MakePrimitiveFunction("last-pair", "1", LastPairImpl)
	MakePrimitiveFunction("last", "1", LastImpl)

	MakePrimitiveFunction("nth", "2", NthImpl)
	MakePrimitiveFunction("take", "2", TakeImpl)
	MakePrimitiveFunction("drop", "2", DropImpl)
}

func CarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "a"), nil
}

func CdrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "d"), nil
}

func CaarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "aa"), nil
}

func CadrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "ad"), nil
}

func CdarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "da"), nil
}

func CddrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "dd"), nil
}

func CaaarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "aaa"), nil
}

func CaadrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "aad"), nil
}

func CadarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "ada"), nil
}

func CaddrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "add"), nil
}

func CdaarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "daa"), nil
}

func CdadrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "dad"), nil
}

func CddarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "dda"), nil
}

func CdddrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "ddd"), nil
}

func CaaaarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "aaaa"), nil
}

func CaaadrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "aaad"), nil
}

func CaadarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "aada"), nil
}

func CaaddrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "aadd"), nil
}

func CadaarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "adaa"), nil
}

func CadadrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "adad"), nil
}

func CaddarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "adda"), nil
}

func CadddrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "addd"), nil
}

func CdaaarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "daaa"), nil
}

func CdaadrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "daad"), nil
}

func CdadarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "dada"), nil
}

func CdaddrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "dadd"), nil
}

func CddaarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "ddaa"), nil
}

func CddadrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "ddad"), nil
}

func CdddarImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "ddda"), nil
}

func CddddrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	return WalkList(Car(args), "dddd"), nil
}

func GeneralCarCdrImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	list := Car(args)
	path := IntegerValue(Cadr(args))
	if path == 0 {
		err = errors.New("general-car-cdr requires a non-zero path specifier")
		return
	}
	for path != 1 {
		code := path & 0x1
		if code == 0 {
			list = Cdr(list)
		} else {
			list = Car(list)
		}
		path = path >> 1
	}
	return list, nil
}

func FirstImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	if VectorP(Car(args)) {
		return VectorFirstImpl(args, env)
	}
	return First(Car(args)), nil
}

func SecondImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	if VectorP(Car(args)) {
		return VectorSecondImpl(args, env)
	}
	return Second(Car(args)), nil
}

func ThirdImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	if VectorP(Car(args)) {
		return VectorThirdImpl(args, env)
	}
	return Third(Car(args)), nil
}

func FourthImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	if VectorP(Car(args)) {
		return VectorFourthImpl(args, env)
	}
	return Fourth(Car(args)), nil
}

func FifthImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	if VectorP(Car(args)) {
		return VectorFifthImpl(args, env)
	}
	return Fifth(Car(args)), nil
}

func SixthImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	if VectorP(Car(args)) {
		return VectorSixthImpl(args, env)
	}
	return Sixth(Car(args)), nil
}

func SeventhImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	if VectorP(Car(args)) {
		return VectorSeventhImpl(args, env)
	}
	return Seventh(Car(args)), nil
}

func EighthImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	if VectorP(Car(args)) {
		return VectorEighthImpl(args, env)
	}
	return Eighth(Car(args)), nil
}

func NinthImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	if VectorP(Car(args)) {
		return VectorNinthImpl(args, env)
	}
	return Ninth(Car(args)), nil
}

func TenthImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	if VectorP(Car(args)) {
		return VectorTenthImpl(args, env)
	}
	return Tenth(Car(args)), nil
}

func LastPairImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	if !(ListP(Car(args)) || DottedListP(Car(args))) {
		err = ProcessError(fmt.Sprintf("last-pair requires a non-circular list but received %s.", String(Car(args))), env)
		return
	}

	if Length(Car(args)) == 0 {
		err = ProcessError("last-pair requires a non-empty list but received nil.", env)
		return
	}

	return LastPair(Car(args)), nil
}

func LastImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	if VectorP(Car(args)) {
		return VectorLastImpl(args, env)
	}
	if !ListP(Car(args)) || ListWithLoopP(Car(args)) {
		err = ProcessError(fmt.Sprintf("last requires a non-circular list but received %s.", String(Car(args))), env)
		return
	}
	return Car(LastPair(Car(args))), nil
}

func NthImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	index := First(args)
	if !IntegerP(index) {
		err = ProcessError(fmt.Sprintf("nth requires integer index but received %s.", String(index)), env)
		return
	}

	col := Second(args)
	if !ListP(col) && !VectorP(col) {
		err = ProcessError(fmt.Sprintf("nth required a list or vector but received %s.", String(col)), env)
		return
	}

	if VectorP(col) {
		return VectorRefImpl(InternalMakeList(col, index), env)
	}

	indexVal := int(IntegerValue(index))
	if indexVal < 0 || indexVal >= Length(col) {
		err = ProcessError(fmt.Sprintf("nth index out of bounds: %d.", indexVal), env)
		return
	}

	return Nth(col, indexVal), nil
}

func TakeImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	n := Car(args)
	if !IntegerP(n) {
		err = ProcessError("take requires a number as its first argument.", env)
	}
	size := int(IntegerValue(n))

	if VectorP(Second(args)) {
		return VectorHeadImpl(InternalMakeList(Second(args), First(args)), env)
	}

	l := Cadr(args)
	if ListP(l) {
		var items []*Data = make([]*Data, 0, Length(args))
		for i, cell := 0, l; i < size && NotNilP(cell); i, cell = i+1, Cdr(cell) {
			items = append(items, Car(cell))
		}
		result = ArrayToList(items)
	} else if ObjectP(l) && ObjectType(l) == "[]byte" {
		dataBytes := (*[]byte)(ObjectValue(l))
		var bytesToCopy []byte
		if size >= len(*dataBytes) {
			bytesToCopy = *dataBytes
		} else {
			bytesToCopy = (*dataBytes)[:size]
		}
		newBytes := make([]byte, len(bytesToCopy))
		for i, v := range bytesToCopy {
			newBytes[i] = v
		}
		result = ObjectWithTypeAndValue("[]byte", unsafe.Pointer(&newBytes))
	} else {
		err = ProcessError("take requires a list, vector, or bytearray as its second argument.", env)
	}
	return
}

func DropImpl(args *Data, env *SymbolTableFrame) (result *Data, err error) {
	n := Car(args)
	if !IntegerP(n) {
		err = ProcessError("drop requires a number as its first argument.", env)
	}
	size := int(IntegerValue(n))

	if VectorP(Second(args)) {
		return VectorTailImpl(InternalMakeList(Second(args), First(args)), env)
	}

	l := Cadr(args)

	if ListP(l) {
		var cell *Data
		var i int
		for i, cell = 0, l; i < size && NotNilP(cell); i, cell = i+1, Cdr(cell) {
		}
		result = cell
	} else if ObjectP(l) && ObjectType(l) == "[]byte" {
		dataBytes := (*[]byte)(ObjectValue(l))
		if size >= len(*dataBytes) {
			newBytes := make([]byte, 0)
			result = ObjectWithTypeAndValue("[]byte", unsafe.Pointer(&newBytes))
		} else {
			newBytes := make([]byte, len(*dataBytes)-size)
			for i, v := range (*dataBytes)[size:] {
				newBytes[i] = v
			}
			result = ObjectWithTypeAndValue("[]byte", unsafe.Pointer(&newBytes))
		}
	} else {
		err = ProcessError("drop requires a list, vector, or bytearray as its second argument.", env)
	}
	return
}
