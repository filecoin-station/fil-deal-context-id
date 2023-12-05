import assert from 'node:assert'
import { test } from 'node:test'
import { buildContextId } from './index.js'
import { CID } from 'multiformats/cid'

test('smoke test', async () => {
  const contextId = await buildContextId({
    // PieceCID: {
    //   '/': 'baga6ea4seaqa3onc6kbej6ekjydnbgjnqjzlnlmcafedswhoyn5ubrfmb34uida'
    // },
    PieceCID: CID.parse('baga6ea4seaqa3onc6kbej6ekjydnbgjnqjzlnlmcafedswhoyn5ubrfmb34uida'),
    VerifiedDeal: true,
    Client: 'f01985718',
    Provider: 'f02008883',
    Label: 'bafybeieosqoxjt6xvwby33ts4gdcykd3fbj6ffj2olhl67lbpefccdnihe',
    StartEpoch: 2652977,
    EndEpoch: 4173617,
    StoragePricePerEpoch: '0',
    ProviderCollateral: '12325970877125150',
    ClientCollateral: '0'
  })

  assert.strictEqual(
    contextId,
    'AXESID6uXFBtGR4r+3yxTK9DqevBRwfAhtlnbqVqusVYErka'
  )
})
