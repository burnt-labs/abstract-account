# Security Policy

This policy covers the abstract account Cosmos SDK module and the authenticator
contract infrastructure in this repository — an asset in the
[Blockchain / DLT bug bounty program](https://github.com/burnt-labs/bug-bounty/blob/main/programs/blockchain.md).
This file summarizes repository-specific terms. The published
[`burnt-labs/bug-bounty`](https://github.com/burnt-labs/bug-bounty) program is
canonical; where the documents differ, the program terms govern.

## Reporting a Vulnerability

**Do not open a public GitHub issue for a security vulnerability.**

| Type of finding                  | How to report                                         |
| -------------------------------- | ----------------------------------------------------- |
| Security vulnerability           | **Security → Report a vulnerability** on this repository, or email [security@burnt.com](mailto:security@burnt.com) |
| Non-sensitive or operational bug | Open a GitHub issue on this repository                 |

Include the type of vulnerability, affected version, steps to reproduce, impact,
how an attacker would exploit it, and any known mitigations.

We acknowledge receipt within **5 business days** and provide a triage decision
within **14 days**. Active exploitation, or confirmed attacker awareness of an
unpatched vulnerability, escalates the issue to Critical **response handling**
— prioritization, coordination, and disclosure timing — regardless of its
original classification. That escalation does not change the finding's
severity assessment or reward eligibility.

## Proof of Concept Requirements

**Reports must include an end-to-end proof of concept.** Severity is assessed on
demonstrated impact under real-world constraints, not theoretical worst-case
scenarios.

Unit tests that construct keeper state directly bypass transaction encoding,
routing, and the ante handler chain, and do not demonstrate on-chain
exploitability on their own.

The proof of concept should run against a **locally running XION node configured
with mainnet parameters**, with the attack executed via standard transaction
broadcast (`BroadcastTxSync` or equivalent). Broadcast acceptance alone is not
sufficient: show inclusion in a block, the successful execution result, and the
resulting state change or security impact. Simulated environments that model
chain state without running a full node do not demonstrate exploitability.

## Authentication Impact Scope

This module implements account authentication, so the following boundary applies
directly.

Authentication weaknesses whose impact is limited to accounts created after the
attack is established — and which cannot affect the funds, state, or
authentication of any account funded and operational before the attack began —
are capped at **Medium**, regardless of the authentication mechanism involved.

A High or Critical authentication finding must demonstrate unauthorized impact
on a **pre-existing funded account**, using only attacker-controlled keys.

## Permissioned Chain Policy

XION mainnet operates with `code_upload_access: Nobody`. Contract deployment
requires a governance proposal.

Any attack vector requiring an attacker to deploy a malicious contract on
mainnet is out of scope, regardless of technical validity — including
authenticator contracts the attacker would need to deploy themselves.

## Privileged Actor Policy

Findings are classified at **Medium at most** when the attack must begin with
control of governance, a module authority, validator or operator credentials,
or another privileged role — or requires that holder to cooperate — and the
demonstrated impact depends on that holder acting self-destructively, outside
normal operation, or in collusion while using authority the role already has.

The cap does not apply when a flaw lets an attacker who starts without that
privilege obtain it or bypass its authorization check, or lets a legitimately
held limited role perform actions outside its intended permissions. Those
findings are assessed by demonstrated impact. This policy does not authorize
researchers to acquire or exercise production privileges they do not
legitimately control, or to test with production privileges they do control.

## Out of Scope

**Assets**

- The chain node and other Cosmos SDK modules — see the
  [Blockchain / DLT program](https://github.com/burnt-labs/bug-bounty/blob/main/programs/blockchain.md)
- Core protocol smart contracts — see the
  [Core Protocol Contracts program](https://github.com/burnt-labs/bug-bounty/blob/main/programs/contracts.md)
- Frontend applications and web properties
- Public blockchain RPC, REST, gRPC, and Tendermint RPC endpoints
- Upstream dependencies — vulnerabilities in CosmWasm, the Cosmos SDK, or IBC
  are not eligible here; only code originating in this repository is covered

**Vulnerability classes**

- Attacks requiring malicious contract deployment on mainnet
- Denial of service of any form recoverable via a software patch, coordinated
  validator restart, or governance parameter update
- Governance attacks requiring a malicious proposal to pass
- Theoretical vulnerabilities without a working end-to-end proof of concept
- Attacks where the attacker's cost to execute exceeds the demonstrable harm to
  the protocol or its users
- Findings affecting only deprecated versions, or already remediated in the
  currently deployed mainnet version, regardless of whether the fix was publicly
  announced
- Best practices, gas optimizations, missing events, and informational findings

## Severity Characterization

| Severity     | Description                                                                                                                        |
| ------------ | ----------------------------------------------------------------------------------------------------------------------------------- |
| **CRITICAL** | Complete bypass of abstract account authentication enabling arbitrary transaction authorization against pre-existing funded accounts |
| **HIGH**     | Theft or freezing of funds affecting individual accounts. Authentication bypass with demonstrated exploitability against an existing account |
| **MEDIUM**   | Partial authentication bypass requiring secondary conditions. Impact limited to accounts created after the attack. Attacks requiring privileged-party cooperation |
| **LOW**      | Valid, reproducible code-level issue with no direct risk to funds, representing a meaningful hardening opportunity                   |

Only **High** and **Critical** findings are reward eligible. The canonical
program governs KYC, duplicate handling, and all other reward terms.

Severity is assessed by Burnt Labs based on demonstrated impact. Reports
submitted at a severity that does not match the definitions above are assessed
as written; we do not reclassify or negotiate severity on a reporter's behalf.

## Responsible Disclosure

- Do not exploit a vulnerability beyond what is necessary to confirm it exists
- **Do not test against XION mainnet.** Testing that targets live production
  systems will disqualify the report
- Do not access, modify, or exfiltrate user data
- Do not disclose publicly before a fix is confirmed and deployed

## Safe Harbor

Burnt Labs will not pursue legal action against researchers who report
vulnerabilities in good faith under this policy, do not exploit beyond what is
necessary to confirm the finding, do not access or disclose user data, and do
not disrupt production systems.

Naming this repository as an asset establishes eligibility, not permission to
test a production deployment. Authorization to actively test extends only to a
local environment or infrastructure you control. This policy does not authorize
testing with production privileged credentials, including credentials you
legitimately hold. Reporting a vulnerability you encountered incidentally is
always welcome.
