# Simulator advanced-settings projection

Parent: [simulator-device-config](simulator-device-config.md)
Zoom Out: [simulator-device-config](simulator-device-config.md)
Zoom In: none
Connector: [osjs-shared-mma-settings](osjs-shared-mma-settings.md) only when work explicitly crosses shared MMA settings.
Source dependencies: `simulator/advanced_projection.go`, `simulator/advanced_projection_test.go`, `simulator/advanced_simulator_test.go`, `simulator/listener.go`.

## Boundary and historical verified contract

`projectAdvancedSettings` selects effective MMA2 memory using listener port and unit ID. An explicit Simulator device value takes precedence; missing policy, state sealing, RBE and extension fields may be hydrated from the matching memory without overriding explicitly present values, including an explicit empty map. Recognized malformed inherited values must be rejected. This node covers projection, not Windows Replicator settings ownership.

The old context named `simulatormain.js` as a dependency without establishing its repository location. That entry has been removed rather than guessed. If a current callsite is needed, perform *bounded missing-territory discovery* from the projection function and register only the verified path.

Historical projection checkpoint: `c289f2f`. The earlier focused test claim is historical, not a fresh test run. Source files and uncommitted overlay have not been revalidated at remote `main` HEAD `45cd3d2d81831306c6943f844e49b7deccbfc5ad`; compare the four declared dependencies against that baseline before using this as current implementation truth. No workflow or product verification authority is conveyed.
