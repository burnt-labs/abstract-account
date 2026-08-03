# x/abstractaccount

Module that implements the `AbstractAccount` type and ante/post handler logics.

## Stable registration addresses

The module owns the stable account-address namespace while preserving
caller-selected, allowlisted account implementations:

- `bootstrap_code_id` is used for every `Instantiate2` address. Runtime
  governance may configure it once; after that, `MsgUpdateParams` rejects every
  attempted change so the address namespace is permanent.
- `registration_enabled` lets governance pause only new registrations without
  changing the address namespace or blocking address queries.
- `MsgRegisterAccount` retains the caller-selected `code_id`. The existing
  allowlist controls which final implementations may be selected.
- Registration instantiates the fixed bootstrap and, when the requested code ID
  differs, atomically migrates the new account to that implementation before
  making the account its own admin.
- `AccountAddress(sender, salt)` returns the registered address or predicts the
  stable address from the bootstrap code checksum. The requested implementation
  does not participate in address derivation.

The bootstrap must accept every supported registration instantiate message and
produce storage compatible with every allowed target implementation. Target
implementations selected during registration must accept the empty `{}`
migration message. If either step is incompatible, registration fails and the
entire transaction is rolled back.

The `(sender, salt) -> address` registry makes duplicate prevention a chain
invariant. Indexers may still provide history and analytics, but are not part
of registration correctness.

The v3 module migration leaves the bootstrap code ID at zero and
`registration_enabled` false. A chain upgrade must configure the bootstrap and
then explicitly enable registration because the reusable module cannot infer a
chain-specific address anchor.

Genesis and chain upgrade handlers call the keeper's trusted `SetParams` path,
which can initialize bootstrap configuration but does not enforce runtime
immutability or Wasm code existence. An upgrade must preserve an already
configured bootstrap ID. Changing it would create a different address namespace
and is not a supported operation.

AA-API must query `AccountAddress(sender, salt)` before constructing
address-bound authenticator credentials. It must not derive the address from
the requested implementation code ID because the bootstrap and requested IDs
may intentionally differ.

Existing registration clients retain the same `MsgRegisterAccount` field layout.
The semantic change is that `code_id` selects only the final implementation;
the fixed bootstrap controls the address.

## License

(c) larry0x, 2023 - [All rights reserved](../../LICENSE).
