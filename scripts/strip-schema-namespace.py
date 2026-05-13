#!/usr/bin/env python3
"""
strip-schema-namespace.py — preprocess a Keyfactor swagger file to strip the
`Keyfactor.Web.KeyfactorApi.Models.` namespace prefix from every schema
definition and every $ref pointer.

Why: the upstream v25 swagger fully qualifies schemas with that prefix
(e.g., `Keyfactor.Web.KeyfactorApi.Models.Templates.TemplateRetrievalResponse`).
openapi-generator translates the schema name directly into the Go type name,
producing verbose identifiers that change every existing consumer reference.
The v25 SDK released in commit 536f3e2 was generated from a swagger that did
not carry this prefix (model file: `model_certificate_authorities_certificate_authority_request.go`,
type: `CertificateAuthoritiesCertificateAuthorityRequest`).

Stripping the prefix preserves the existing naming convention and keeps the
hand-written v25/client.go and downstream consumers compatible.

NOTE: other `Keyfactor.*` namespaces are preserved (e.g., `Keyfactor.Common.Scheduling.*`,
`Keyfactor.Platform.Extensions.Enums.*`). Only `Keyfactor.Web.KeyfactorApi.Models.`
is stripped, matching the baseline v25 code's naming.

Usage:
    strip-schema-namespace.py <input.swagger.json> <output.swagger.json>

Documented for engineering: see SWAGGER_GAPS_FOR_ENGINEERING.md.
"""

from __future__ import annotations

import json
import sys
from collections import OrderedDict
from pathlib import Path

PREFIX = "Keyfactor.Web.KeyfactorApi.Models."
REF_PREFIX = f"#/components/schemas/{PREFIX}"


def strip_schema_keys(data: OrderedDict) -> int:
    """Rename schemas under components.schemas to drop the prefix.
    Returns the number of schemas renamed.
    Mutates `data` in place.
    """
    schemas = data.get("components", {}).get("schemas")
    if not isinstance(schemas, OrderedDict):
        return 0
    renamed = 0
    new_schemas = OrderedDict()
    seen = set()
    for name, body in schemas.items():
        if name.startswith(PREFIX):
            new_name = name[len(PREFIX):]
            if new_name in seen or new_name in schemas:
                # Collision with an existing schema name. Refuse to clobber.
                # This is a swagger bug we'd want to know about.
                raise ValueError(
                    f"Schema rename collision: stripping '{name}' would clash with existing schema '{new_name}'"
                )
            new_schemas[new_name] = body
            seen.add(new_name)
            renamed += 1
        else:
            if name in seen:
                raise ValueError(f"Duplicate schema name in input: {name}")
            new_schemas[name] = body
            seen.add(name)
    data["components"]["schemas"] = new_schemas
    return renamed


def rewrite_refs(node, counter: list) -> None:
    """Walk the entire tree rewriting any $ref pointing into the prefixed namespace."""
    if isinstance(node, dict):
        for k, v in node.items():
            if k == "$ref" and isinstance(v, str) and v.startswith(REF_PREFIX):
                node[k] = "#/components/schemas/" + v[len(REF_PREFIX):]
                counter[0] += 1
            else:
                rewrite_refs(v, counter)
    elif isinstance(node, list):
        for item in node:
            rewrite_refs(item, counter)


def main(argv: list[str]) -> int:
    if len(argv) != 3:
        print(f"usage: {argv[0]} <input.swagger.json> <output.swagger.json>", file=sys.stderr)
        return 2
    src = Path(argv[1])
    dst = Path(argv[2])

    with src.open() as f:
        data = json.load(f, object_pairs_hook=OrderedDict)

    renamed = strip_schema_keys(data)
    refs_counter = [0]
    rewrite_refs(data, refs_counter)
    refs = refs_counter[0]

    dst.parent.mkdir(parents=True, exist_ok=True)
    with dst.open("w") as f:
        json.dump(data, f, indent=2, ensure_ascii=False)
        f.write("\n")

    print(f"  stripped '{PREFIX}' from {renamed} schemas and {refs} $ref pointers")
    print(f"  wrote: {dst}")
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv))
