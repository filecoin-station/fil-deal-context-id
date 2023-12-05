package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-address"
	cborutil "github.com/filecoin-project/go-cbor-util"
	market "github.com/filecoin-project/go-state-types/builtin/v10/market"
	cid "github.com/ipfs/go-cid"
)

type MarketDeal struct {
	Proposal market.DealProposal
	State    market.DealState
}

func main() {
	jsonString := `{"Proposal":{"PieceCID":{"/":"baga6ea4seaqmh37khqbfmvt3phtfsxuz2ok3koypajjxn7gwml26uojmub54kny"},"PieceSize":34359738368,"VerifiedDeal":true,"Client":"f01880278","Provider":"f02118066","Label":"bafybeiadnyv6t65hnwx4as467n4p7k4u6bwsh3dhw24xrahfdwcmpddyo4","StartEpoch":2867966,"EndEpoch":4359806,"StoragePricePerEpoch":"0","ProviderCollateral":"11361143711469027","ClientCollateral":"0"},"State":{"SectorStartEpoch":2856167,"LastUpdatedEpoch":3106586,"SlashEpoch":-1,"VerifiedClaim":18771683}}`
	deal := MarketDeal{}
	err := json.Unmarshal([]byte(jsonString), &deal)
	if err != nil {
		panic(fmt.Sprintf("Error parsing JSON: %v", err))
	}

	dealProposal := deal.Proposal

	// dealProposal := market.DealProposal{
	// 	PieceCID:             mustCID("baga6ea4seaqgfxci3yrqtlrdznacareb3nysxhr2d6umu2bgxbrjgnjgt6ez6jy"),
	// 	PieceSize:            abi.PaddedPieceSize(34359738368),
	// 	VerifiedDeal:         true,
	// 	Client:               mustAddr("f01880278"),
	// 	Provider:             mustAddr("f02366527"),
	// 	Label:                mustDealLabel("bafybeigduxb5che3ov434qtl3ebv7h2tx36ogzz2cwhec2jnrvdkshp4vy"),
	// 	StartEpoch:           abi.ChainEpoch(3163634),
	// 	EndEpoch:             abi.ChainEpoch(4675634),
	// 	StoragePricePerEpoch: abi.NewTokenAmount(0),
	// 	ProviderCollateral:   abi.NewTokenAmount(9723097997285858),
	// 	ClientCollateral:     abi.NewTokenAmount(0),
	// }

	fmt.Printf("DealProposal: %#v\n", dealProposal)

	// Copy'n'paste from
	// https://github.com/filecoin-project/boost/blob/ca748c3c6916449524921efc8389f9b3898361fe/indexprovider/wrapper.go#L168-L172
	propnd, err := cborutil.AsIpld(&dealProposal)
	if err != nil {
		panic(fmt.Sprintf("failed to compute signed deal proposal ipld node: %v", err))
	}
	propCid := propnd.Cid()

	fmt.Printf("DAG-CBOR: %x\n", propnd.RawData())
	fmt.Printf("CID bytes: %v\n", base64.StdEncoding.EncodeToString(propCid.Bytes()))
}

func mustDealLabel(label string) market.DealLabel {
	l, err := market.NewLabelFromString(label)
	if err != nil {
		panic(err)
	}
	return l
}

// Helper function to create a new address from a string
func mustAddr(addr string) address.Address {
	a, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
	}
	return a
}

// Helper function to create a new CID from a string
func mustCID(c string) cid.Cid {
	cid, err := cid.Parse(c)
	if err != nil {
		panic(err)
	}
	return cid
}
