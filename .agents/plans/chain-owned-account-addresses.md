# Chain-owned abstract-account addresses

## Outcome

Make `x/abstractaccount` the authority for stable address derivation and
`(sender, salt)` uniqueness while retaining direct, caller-selected account
implementations.

## Wasm boundary

- Add a keeper-only `Instantiate2WithAddressHash` extension in the XION Wasmd
  fork.
- Keep the public `Instantiate2` message and behavior unchanged.
- Reuse the existing instantiation pipeline, authorization policy, storage, and
  events; replace only the checksum input to predictable address derivation.
- Require an exact 32-byte address hash.

## Module contract

- Add immutable `address_derivation_hash` and independent
  `registration_enabled` parameters.
- Retain `MsgRegisterAccount.code_id` as the caller-selected implementation and
  enforce the existing allowlist before instantiation.
- Directly instantiate that code at the address derived from
  `(address_derivation_hash, sender, salt)`; do not bootstrap or migrate.
- Store `(sender, salt) -> account address` and reject duplicates atomically.
- Add `AccountAddress(sender, salt)` to return the stored address or predict the
  fixed-hash address.
- Recognize an abstract account already present at the predicted address even
  when it predates the registry.

## Consumer responsibilities

After adoption, applications should replace local checksum derivation and
indexer-based existence checks with the module query. Indexers remain useful for
history, analytics, and authenticator discovery, but not registration
correctness.

## Upgrade boundary

The v3 store migration leaves the hash empty and registration disabled. The
XION upgrade must configure an audited chain-specific hash and explicitly enable
registration. Existing deployed accounts and ordinary Wasm migrations are
unchanged.

## Verification

- Wasmd tests proving two different code IDs derive the same address from the
  same fixed hash, sender, and salt while instantiating the requested code.
- Parameter validation and address-hash immutability tests.
- Address prediction, registered lookup, duplicate rejection, direct
  instantiation, atomic failure, and historical-address recognition tests.
- Store migration, genesis round-trip, protobuf compatibility, and repository
  test gates.

## Security properties

- Only the abstract-account keeper can use the fixed-hash instantiation path.
- Caller-selected code IDs remain constrained by the existing allowlist.
- The immutable hash prevents governance from silently creating a second
  address namespace after activation.
- Registration succeeds or fails atomically without a bootstrap compatibility
  or migration-message dependency.
