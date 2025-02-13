package action

import "github.com/rsteube/carapace"

func ActionTargets() carapace.Action {
	return carapace.ActionValues(
		"copy_assets",
		"kernel_snapshot",
		"aot_elf_profile",
		"aot_elf_release",
		"aot_assembly_profile",
		"aot_assembly_release",
		"debug_macos_framework",
		"debug_macos_bundle_flutter_assets",
		"debug_bundle_linux_x64_assets",
		"debug_bundle_linux_arm64_assets",
		"profile_bundle_linux_x64_assets",
		"profile_bundle_linux_arm64_assets",
		"release_bundle_linux_x64_assets",
		"release_bundle_linux_arm64_assets",
		"web_service_worker",
		"release_android_application",
		"copy_flutter_bundle",
		"debug_android_application",
		"profile_android_application",
		"android_aot_bundle_profile_android-arm",
		"android_aot_bundle_profile_android-arm",
		"android_aot_bundle_profile_android-arm64",
		"android_aot_bundle_profile_android-x64",
		"android_aot_bundle_release_android-arm",
		"android_aot_bundle_release_android-arm64",
		"android_aot_bundle_release_android-x64",
		"android_aot_deferred_components_bundle_profile_android-arm",
		"android_aot_deferred_components_bundle_profile_android-arm64",
		"android_aot_deferred_components_bundle_profile_android-x64",
		"android_aot_deferred_components_bundle_release_android-arm",
		"android_aot_deferred_components_bundle_release_android-arm64",
		"android_aot_deferred_components_bundle_release_android-x64",
		"debug_ios_bundle_flutter_assets",
		"profile_ios_bundle_flutter_assets",
		"release_ios_bundle_flutter_assets",
		"unpack_windows",
		"debug_bundle_windows_assets",
		"profile_bundle_windows_assets",
		"release_bundle_windows_assets",
		"debug_bundle_windows_assets_uwp",
		"profile_bundle_windows_assets_uwp",
		"release_bundle_windows_assets_uwp",
	)
}
