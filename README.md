# fil-deal-context-id

Calculate IPNI ContextID for Filecoin StorageMarket deals

## Basic use

```js
import { buildContextId } from 'fil-deal-context-id'
import { CID } from 'multiformats/cid'

const dealProposal = {
  PieceCID: CID.parse('baga6ea4seaqa3onc6kbej6ekjydnbgjnqjzlnlmcafedswhoyn5ubrfmb34uida'),
  VerifiedDeal: true,
  Client: 'f01985718',
  Provider: 'f02008883',
  Label: 'bafybeieosqoxjt6xvwby33ts4gdcykd3fbj6ffj2olhl67lbpefccdnihe',
  StartEpoch: 2652977,
  EndEpoch: 4173617,
  StoragePricePerEpoch: '0',
  ProviderCollateral: '12325970877125150',
  ClientCollateral: '0',
}
const contextId = await buildContextId(dealProposal)
console.log(contextId)
```

## Go reference implementation

See [go/main.go](./go/main.go) for a reference implementation
calculating Context ID from a Deal JSON object found in `StateMarketDeals.json`
using Filecoin Go libraries.
