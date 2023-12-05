import * as codec from '@ipld/dag-cbor'
import { CID } from 'multiformats/cid'
import { sha256 } from 'multiformats/hashes/sha2'
import { base64 } from 'multiformats/bases/base64'

/**
 * @param {{
 *   PieceCID: {
 *     '/': string;
 *   };
 *   PieceSize: number;
 *   VerifiedDeal: boolean;
 *   Client: string;
 *   Provider: string;
 *   Label?: string;
 *   StartEpoch: number;
 *   EndEpoch: number;
 *   StoragePricePerEpoch: string;
 *   ProviderCollateral: string;
 *   ClientCollateral: string;
 * }} dealProposal
 */
export async function buildContextId (dealProposal) {
  const data = codec.encode(dealProposal)
  console.log(Buffer.from(data).toString('hex'))
  const hash = await sha256.digest(data)
  const cid = CID.createV1(codec.code, hash)
  return base64.baseEncode(cid.bytes)
}
