# Replicator advanced settings

Replicator has Device Definition and Advanced Settings folder tabs. Advanced Settings reuses Memory's RBE Rules, State Sealing and Access Policy editor.

These settings apply to the Replicator's destination MMA2 memory, not to the source PLC. Destination ranges derive from the union of Pull Blocks per function code. A sealing control coil must lie in a configured destination coil range. Shared RBE TCP output is configured through Memory's MMA Settings popup.

New devices allow IPv4 and IPv6, Read/Write, and default to sealing off. Existing destination settings are inherited during runtime composition unless explicitly overridden. Changes persist as mma2_advanced in the Replicator document; legacy documents without that field retain existing destination configuration.

Both editors use an unfiltered preset dropdown beside the editable Source IP / CIDR box. It remains available after typing. Comma-separated addresses and CIDRs are trimmed and stored as separate source_ip entries; empty comma segments are ignored. Invalid addresses still fail authoritative MMA2 validation.

Rebuild/restart the Replicator backend with the UI. New advanced saves check backend capability before sending Apply, so an outdated backend cannot silently discard the new fields.
