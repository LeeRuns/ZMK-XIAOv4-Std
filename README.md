# ZMK-XIAOv4-Std

ZMK firmware for the Skree-style split keyboard on Seeeduino XIAO BLE controllers.

## Editing the keymap

Keymap edits are made through [nickcoutsos.github.io/keymap-editor](https://nickcoutsos.github.io/keymap-editor). The editor commits directly to a branch on this repo; GitHub Actions builds the firmware; download the `.uf2` artifacts from the run and flash each half.

The editor labels layers as **Layer 1 / Layer 2 / Layer 3 / Layer 4** (1-indexed). Their actual purpose in this keymap:

| Editor label | Layer index | Internal name in keymap | Purpose |
|---|---|---|---|
| Layer 1 | 0 | `default_layer` | **Base** — standard QWERTY typing layer |
| Layer 2 | 1 | `layer_1` | **System** — Bluetooth profiles, RGB, Studio unlock, Mac-mode toggle, bootloader |
| Layer 3 | 2 | `layer_2` | **Spare** — currently mostly empty; only holds a left-side bootloader trigger |
| Layer 4 | 3 | `mac_mode` | **Mac mode** — remaps the bottom thumb row so `Ctrl`/`Alt` act as `Cmd` on macOS |

### Layer details

**Base (Layer 1 in editor)** — numbers row, QWERTY alphas, standard punctuation. Bottom thumb row exposes `Shift / Space / Backspace` (left) and `Delete / Enter / RGUI` (right); inner thumbs are `Mo1` (system layer hold) and `Mo2` (spare layer hold).

**System (Layer 2 in editor)** — held with left inner thumb (`Mo1`).
- Top row: `BT_CLR`, `BT_SEL 0–4` (pair / select Bluetooth profiles)
- Second row: `studio_unlock` (lets ZMK Studio modify the keyboard live)
- Bottom row: `RGB_TOG`, `RGB_EFF`, and **`Mo1 + M` toggles Mac mode**
- Right inner thumb in this layer triggers `&bootloader` (puts the right half into UF2 flash mode)

**Spare (Layer 3 in editor)** — held with right inner thumb (`Mo2`). Currently mostly `&trans`. The left inner thumb in this layer triggers `&bootloader` (puts the left half into UF2 flash mode). Free real estate for adding screenshot keys, media keys, etc.

**Mac mode (Layer 4 in editor)** — toggled on/off via `Mo1 + M`. Only changes the bottom thumb row: `LCTRL`/`LALT` → `LGUI` (Cmd) and `RCTRL`/`RALT` → `RGUI`. Tap `Mo1 + M` again to return to Windows mode.

## Build pipeline

The CI workflow is currently pinned to a pre-Zephyr-4.1 snapshot of ZMK upstream (commit `abb64ba`, 2025-12-07) to keep builds reproducible. See `.github/workflows/blank.yml` and `config/west.yml`. To pick up newer ZMK features, both pins need to advance together and the board identifier may need updating to the new ZMK-variant naming.

## Hardware

- Two halves, each driven by a Seeeduino XIAO BLE (`seeeduino_xiao_ble`)
- Polyimide matrix; ghosting-fix timing applied in `config/skree-custom.conf`
- WS2812 RGB underglow
- ZMK Studio enabled on the left (central) half via `studio-rpc-usb-uart` snippet

## Flashing

1. Push or merge — GitHub Actions builds three artifacts: `skreecustom_left`, `skreecustom_right`, `settings_reset` (last is a recovery image).
2. Double-tap reset on a XIAO BLE → it mounts as a USB drive.
3. Drag the matching `.uf2` onto it. The board reboots automatically. Repeat for the other half.
