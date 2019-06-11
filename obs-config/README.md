# OBS Configurations

If you are using Linux, follow the pre-configuration steps below first.

## Setting up the TGIK Streaming Profile

In OBS, select `Profile -> Import`, and select the
[profiles/TGIK](profiles/TGIK) directory.

This sets up the stream profile and basic video streaming profile.

Then verify hardware encoding is going to be used by going to `File -> Settings
-> Output` and checking the encoder. For Linux, this should be `FFMPEG VAAPI`.

## Setting up the TGIK scene

In OBS, select `Scene Collection -> Import`, then select either
[scenes/tgik-macos.json](scenes/tgik-macos.json) or
[scenes/tgik-linux-mate.json](scenes/tgik-linux-mate.json) depending on your
computer. You may need to tweak the `crop-bottom`, and `crop-top` values for the
display capture for your particular set up.

Using a resolution of 1600x1200 is recommended for Linux.

### Verify capture devices

Check that `Video Capture Device` and `Doc Cam` (if you have one) match your
actual cameras.

### Update the title card

Re-point `Title Card Image` at the latest TGIK graphics.

### Linux browser configuration

Whereas the MacOS browser plugin can use embedded CSS, the Linux equivalent must
load a file. Enter the configuration for YouTube chat, then change the `Custom
CSS` path to point to [scenes/youtube.css](scenes/youtube.css)

## Linux pre-configuration

After installing OBS, you will also need the [Linux Browser
plugin][linux-browser-plugin] installed.

You should also configure [VA-API][VA-API] as appropriate for your computer.
This is normally achieved by setting `LIBVA_DRIVER_NAME` environment variable as
appropriate in `/etc/profile.d/` and installing the correct driver package
(along the lines of `libva-intel-driver` / `libva-amdgpu`) for your system.

Once done, `vainfo` should report back something like:

```
libva info: VA-API version 1.4.1
libva info: va_getDriverName() returns 0
...
vainfo: VA-API version: 1.4 (libva 2.4.0)
vainfo: Driver version: Mesa Gallium driver 19.0.5 for Radeon RX 580 Series (POLARIS10, DRM 3.30.0, 5.1.7-300.fc30.x86_64, LLVM 8.0.0)
vainfo: Supported profile and entrypoints
...
      VAProfileHEVCMain               : VAEntrypointVLD
      VAProfileHEVCMain               : VAEntrypointEncSlice
      VAProfileHEVCMain10             : VAEntrypointVLD
```

[linux-browser-plugin]: https://github.com/bazukas/obs-linuxbrowser
[VA-API]: https://wiki.archlinux.org/index.php/Hardware_video_acceleration
