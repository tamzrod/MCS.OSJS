# Simulator MMA2 Projection Context

## Overview

Context node for Simulator MMA2 projection settings. Handles:
- `state_sealing` logic
- RBE Rules (RBE) integration  
- Extension hydration
- Device override semantics via `projectAdvancedSettings`

## Listener Matching

Listeners use the pair `{"port": 0, "unitId": x}` as keys to match specific projection contexts.

## Hydration Semantics

Hydrated values are explicitly overridden by device-provided values. Extra fields from hydration do not populate already-present Extra map keys.
