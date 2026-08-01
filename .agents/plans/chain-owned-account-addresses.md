# Chain-owned abstract-account addresses

## Outcome

Make `x/abstractaccount` the sole authority for account code selection, stable
address derivation, and `(sender, salt)` uniqueness. Callers submit only the
instantiate message, funds, and salt. They never submit or derive from a Wasm
code ID or checksum.

## Module contract

- Add `bootstrap_code_id` and `implementation_code_id` parameters.
- Permit `bootstrap_code_id` to transition from zero to a configured value
  once, then make it immutable. Governance may update
  `implementation_code_id`.
- Instantiate with `bootstrap_code_id`, migrate with `{}` to
  `implementation_code_id` when the two differ, then make the contract its own
  admin and register it as an abstract account in one transaction.
- Remove `code_id` from `MsgRegisterAccount` and reserve its protobuf field.
- Store `(sender, salt) -> account address` for every new registration and
  reject duplicate registrations atomically.
- Add `AccountAddress(sender, salt)` query. It returns a stored address when
  registered and otherwise predicts from the bootstrap code checksum held by
  the Wasm module.
- Add `MigrateAccount(sender)` so an account can migrate itself to the current
  on-chain implementation without a caller supplying a code ID.
- Treat a historical abstract account already present at the canonical
  predicted address as registered even when it predates the registry.

## Responsibilities removed from applications and indexers

After consumers adopt the new message/query, the module replaces:

- AA-API `CODE_ID`, `EXPECTED_CHECKSUM`, and `LEGACY_CHECKSUMS` configuration;
- client-selected code IDs and allow-list validation;
- code IDs in existing-account migration messages;
- local checksum address derivation in AA-API and xion.js;
- Numia point lookups used to decide whether registration is safe;
- RPC scans over current and legacy checksum-derived addresses;
- KV-based correctness locks for duplicate registration (KV may remain only as
  an optional request-coalescing optimization).

Indexers remain useful for history, analytics, and discovery by authenticator,
but are no longer in the account-creation correctness path.

## Upgrade boundary

The v3 store migration adds the new parameters as zero (registration disabled)
because a reusable module cannot infer chain-specific Wasm code IDs. The XION
chain upgrade must set both IDs before reopening registration. The audited
mainnet accounts must all resolve at the selected bootstrap checksum. Any
noncanonical historical account would need an explicit chain-specific backfill;
testnet-only historical duplicates do not constrain the mainnet rollout.

## Verification

- Parameter validation and bootstrap immutability tests.
- Address prediction, registered lookup, duplicate rejection, and historical
  canonical-address recognition tests.
- Atomic instantiate/migrate/admin/account registration tests.
- v2-to-v3 migration and genesis round-trip tests.
- CLI/protobuf compatibility tests and full repository test/coverage gates.
