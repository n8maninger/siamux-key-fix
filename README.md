SiaMux Key Fix

Fixes an issue with Sia hosts where the SiaMux public key will not be loaded from the host's public key after losing the siamux folder.

# Usage

1. Download a release or compile
2. Stop Sia
3. Find your Sia data path

+ Linux: `$HOME/.config/Sia-UI/sia/`
+ Mac: `$HOME/Library/Application Support/Sia-UI/sia/`
+ Windows: `%APPDATA%\Sia-UI\sia\`

4. Rename the `siamux` folder in your data path to `siamux.old`
5. Create a new `siamux` folder
6. Find your `host.json` file, it should be in the `host` folder of your data path
7. Open a terminal at the location you downloaded this program 
8. Type `siamuxfix "path to your host.json"` and press enter
9. Move the generated `siamux.json` and `siamux.json_temp` into the `siamux` folder
10. Start Sia and unlock your wallet

# Building

From the project directory:
```
go build -o siamuxfix .
```