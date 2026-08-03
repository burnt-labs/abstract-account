# Chain-owned abstract-account addresses

## Outcome

Make `x/abstractaccount` the authority for stable address derivation and
`(sender, salt)` uniqueness while retaining caller-selected, allowlisted account
implementations.

## Module contract

- Add an immutable `bootstrap_code_id` parameter.
- Add an independent `registration_enabled` parameter so governance can pause
  new registrations without changing the stable namespace or blocking queries.
- Permit `bootstrap_code_id` to transition from zero to a configured value
  once, then make it immutable.
- Instantiate with `bootstrap_code_id`, migrate with `{}` to
  the caller-selected, allowlisted `code_id` when the two differ, then make the
  contract its own admin and register it as an abstract account in one
  transaction.
- Retain `code_id` in `MsgRegisterAccount` as the requested final
  implementation. It does not affect the address.
- Store `(sender, salt) -> account address` for every new registration and
  reject duplicate registrations atomically.
- Add `AccountAddress(sender, salt)` query. It returns a stored address when
  registered and otherwise predicts from the bootstrap code checksum held by
  the Wasm module.
- Treat a historical abstract account already present at the canonical
  predicted address as registered even when it predates the registry.
- Require the bootstrap to accept supported registration messages and produce
  storage compatible with allowed target implementations; targets must accept
  the empty `{}` migration message.

## Responsibilities removed from applications and indexers

After consumers adopt the new message/query, the module replaces:

- local checksum address derivation in AA-API and xion.js;
- Numia point lookups used to decide whether registration is safe;
- RPC scans over current and legacy checksum-derived addresses;
- KV-based correctness locks for duplicate registration (KV may remain only as
  an optional request-coalescing optimization).

Indexers remain useful for history, analytics, and discovery by authenticator,
but are no longer in the account-creation correctness path.

## Upgrade boundary

The v3 store migration adds the bootstrap as zero and leaves
`registration_enabled` false because a reusable module cannot infer the
chain-specific address anchor. The XION chain upgrade must set the bootstrap ID
before explicitly enabling registration. The audited mainnet accounts must all
resolve at the selected bootstrap checksum. Any
noncanonical historical account would need an explicit chain-specific backfill;
testnet-only historical duplicates do not constrain the mainnet rollout.

## Verification

- Parameter validation and bootstrap immutability tests.
- Address prediction, registered lookup, duplicate rejection, and historical
  canonical-address recognition tests.
- Atomic instantiate/migrate/admin/account registration tests.
- v2-to-v3 migration and genesis round-trip tests.
- CLI/protobuf compatibility tests and full repository test/coverage gates.

## Security review

- Caller-selected final code IDs remain constrained by the existing allowlist.
- Bootstrap-to-requested-code migration is synchronous and registration fails
  atomically if it fails.
- Authenticator verification remains address-bound downstream; AA-API must
  obtain that address from `AccountAddress` before credential construction.
- Signer state remains transient; only the canonical `(sender, salt) -> address` registry is persistent.
- Code selection remains with callers while uniqueness is enforced atomically on chain.
