// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import "github.com/centrifuge/go-substrate-rpc-client/v4/scale"

type NetworkID struct {
	IsAny bool

	IsNamed      bool
	NamedNetwork []U8

	IsPolkadot bool

	IsKusama bool
}

func (n *NetworkID) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch b {
	case 0:
		n.IsAny = true
	case 1:
		n.IsNamed = true

		return decoder.Decode(&n.NamedNetwork)
	case 2:
		n.IsPolkadot = true
	case 3:
		n.IsKusama = true
	}

	return nil
}

func (n NetworkID) Encode(encoder scale.Encoder) error {
	switch {
	case n.IsAny:
		return encoder.PushByte(0)
	case n.IsNamed:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(n.NamedNetwork)
	case n.IsPolkadot:
		return encoder.PushByte(2)
	case n.IsKusama:
		return encoder.PushByte(3)
	}

	return nil
}

type OptionNetworkIDV3 struct {
	option
	value NetworkIDV3
}

func NewOptionNetworkIDV3Empty() OptionNetworkIDV3 {
	return OptionNetworkIDV3{option: option{hasValue: false}}
}

func NewOptionNetworkIDV3(value NetworkIDV3) OptionNetworkIDV3 {
	return OptionNetworkIDV3{option{hasValue: true}, value}
}

func (o *OptionNetworkIDV3) Decode(decoder scale.Decoder) error {
	return decoder.DecodeOption(&o.hasValue, &o.value)
}

func (o OptionNetworkIDV3) Encode(encoder scale.Encoder) error {
	return encoder.EncodeOption(o.hasValue, o.value)
}

type NetworkIDV3 struct {
	IsByGenesis bool
	ByGenesis   []U8

	IsByForked bool
	ByForked   []U8

	IsPolkadot bool

	IsKusama bool

	IsWestend bool

	IsRococo bool

	IsWococo bool
}

func (n *NetworkIDV3) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch b {
	case 0:
		n.IsByGenesis = true
		return decoder.Decode(&n.ByGenesis)
	case 1:
		n.IsByForked = true
		return decoder.Decode(&n.ByForked)
	case 2:
		n.IsPolkadot = true
	case 3:
		n.IsKusama = true
	case 4:
		n.IsWestend = true
	case 5:
		n.IsRococo = true
	case 6:
		n.IsWococo = true
	}

	return nil
}

func (n NetworkIDV3) Encode(encoder scale.Encoder) error {
	switch {
	case n.IsByGenesis:
		if err := encoder.PushByte(0); err != nil {
			return err
		}
		return encoder.Encode(n.ByGenesis)
	case n.IsByForked:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(n.ByForked)
	case n.IsPolkadot:
		return encoder.PushByte(2)
	case n.IsKusama:
		return encoder.PushByte(3)
	case n.IsWestend:
		return encoder.PushByte(4)
	case n.IsRococo:
		return encoder.PushByte(5)
	case n.IsWococo:
		return encoder.PushByte(6)
	}

	return nil
}
