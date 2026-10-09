# LG Remote on Android

`scripts/android-lg.sh` is the debug build and install path. Run
`scripts/android-release.sh lg-remote` to create the signed arm64 release APK
at `artifacts/lg-remote-arm64-release.apk`.

All Android mobile example APKs are listed in
[the APK index](../../android-host/APPS.md). The release key and password
files live under the ignored `temp/android-release-signing/` directory. Keep
a secure backup of the key because updates must use the same signing identity.

The app uses the shared `examples/android-host` module for window insets and
Android lifecycle handling. The generated AAR and Gradle outputs are ignored;
the release APK is the deliverable.
