// OS.js client configuration for the neutral desktop shell.
// Auto-login as the demo user so the desktop renders directly (no auth UI
// needed). No Governor connection.
// SHELL-001: default the desktop to the Nameless Workstation theme.
// SHELL-002: classic workstation taskbar anchored to the bottom edge, with a
// Menu (start/launcher) area, window buttons (active application task
// buttons), and a clock — Windows-2000-style desktop layout.
// SHELL-004: force the English (en_EN) locale so the industrial-workstation
// shell is deterministic regardless of the host browser locale, and wire the
// NamelessClassicIcons icon theme (original era-appropriate SVGs).

export default {
  auth: {
    login: {
      username: 'demo',
      password: 'demo'
    }
  },
  locale: {
    language: 'en_EN'
  },
  settings: {
    // OSUI-003: persist settings on the server (per-user VFS home) instead of
    // browser localStorage, so the Auto-Start selection is an operator setting
    // that travels with the workstation, not the browser profile.
    adapter: 'server'
  },
  desktop: {
    settings: {
      theme: 'NamelessWorkstationTheme',
      icons: 'NamelessClassicIcons',
      // OSFIN-003: set the desktop background to the intended restrained
      // dark-blue workstation field. OS.js's built-in default is
      // {src: <wallpaper.png>, color: "#572a79" (purple), style: "cover"}.
      // The config object is MERGED over that default, so merely overriding
      // `color` leaves the default `src` (the stock wallpaper PNG) intact —
      // the PNG renders on top of the color, which is the "still purple"
      // defect a fresh human runtime inspection found. The fix has two parts,
      // both using standard OS.js mechanisms:
      //   1. `style: "color"` — OS.js only sets background-image when style is
      //      NOT "color" (see client main.js background logic). With "color"
      //      it leaves background-image untouched in the inline style.
      //   2. The NamelessWorkstationTheme overrides `.osjs-root` background-
      //      image: none, so the stock wallpaper PNG in the base @osjs/client
      //      CSS never shows. The `color` here then fills the desktop.
      background: {
        color: '#3c6ea6',
        style: 'color'
      },
      panels: [
        {
          position: 'bottom',
          ontop: true,
          items: [
            // SHELL-004: the menu (launcher) button. The panel module renders
            // options.icon verbatim as the <img src>, so it must be a URL, not
            // a bare icon name (the fs.icon resolver is not applied here).
            // Point it at our original start-here SVG served from the
            // NamelessClassicIcons theme dist. This shell is dedicated to that
            // icon theme, so the explicit theme path is intentional.
            {name: 'menu', options: {icon: '/icons/NamelessClassicIcons/icons/start-here.svg'}},
            {name: 'windows'},
            {name: 'tray'},
            {name: 'clock'}
          ]
        }
      ]
    }
  }
};

