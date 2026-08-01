# x/abstractaccount

Module that implements the `AbstractAccount` type and ante/post handler logics.

## Chain-owned registration

The module owns the stable account-address namespace and current contract
implementation:

- `bootstrap_code_id` is used for every `Instantiate2` address and becomes
  immutable once configured.
- `implementation_code_id` can be updated by governance. New accounts are
  atomically migrated to it before registration completes.
- `registration_enabled` lets governance pause only new registrations without
  changing the address namespace or blocking queries and existing-account
  migrations.
- `MsgRegisterAccount` accepts only the sender, instantiate message, funds, and
  salt. Callers never provide a Wasm code ID.
- `AccountAddress(sender, salt)` returns the registered address or predicts the
  canonical address from the bootstrap code checksum.
- `MsgMigrateAccount` migrates an existing account to the current
  implementation without exposing a code ID to the caller.

The `(sender, salt) -> address` registry makes duplicate prevention a chain
invariant. Indexers may still provide history and analytics, but are not part
of registration correctness.

The v3 module migration leaves both registration code IDs at zero and
`registration_enabled` false. A chain upgrade must configure the IDs and then
explicitly enable registration because the reusable module cannot infer
chain-specific code IDs.

AA-API must query `AccountAddress(sender, salt)` before constructing
address-bound authenticator credentials. In particular, it must not derive the
address from `implementation_code_id`: bootstrap and implementation code IDs may
intentionally differ.

## License

(c) larry0x, 2023 - [All rights reserved](../../LICENSE).
