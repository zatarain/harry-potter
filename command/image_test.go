package command

import "testing"

func TestImageCommandUse(t *testing.T) {
	if imageCommand.Use != "image" {
		t.Errorf("imageCommand.Use = %q, want 'image'", imageCommand.Use)
	}
}

func TestImageCommandHasImgAlias(t *testing.T) {
	if !hasAlias(imageCommand, "img") {
		t.Error("alias 'img' not found in imageCommand.Aliases")
	}
}

func TestImageSubcommands(t *testing.T) {
	for _, name := range []string{
		"build", "history", "import", "inspect", "load",
		"ls", "prune", "pull", "push", "rm", "save", "tag",
	} {
		if !hasSubcommand(imageCommand, name) {
			t.Errorf("subcommand %q not found on imageCommand", name)
		}
	}
}

func TestImageListAliases(t *testing.T) {
	if !hasAlias(imageListCommand, "list") {
		t.Error("alias 'list' not found in imageListCommand.Aliases")
	}
}

func TestImageRemoveAliases(t *testing.T) {
	if !hasAlias(imageRemoveCommand, "remove") {
		t.Error("alias 'remove' not found in imageRemoveCommand.Aliases")
	}
}

func TestImageBuildFlags(t *testing.T) {
	for _, flag := range []string{"file", "tag", "no-cache", "build-arg", "platform", "pull", "quiet", "target"} {
		if !hasFlag(imageBuildCommand, flag) {
			t.Errorf("flag --%s not registered on imageBuildCommand", flag)
		}
	}
}

func TestImageListFlags(t *testing.T) {
	for _, flag := range []string{"all", "digests", "filter", "format", "no-trunc", "quiet"} {
		if !hasFlag(imageListCommand, flag) {
			t.Errorf("flag --%s not registered on imageListCommand", flag)
		}
	}
}

func TestImagePullFlags(t *testing.T) {
	for _, flag := range []string{"all-tags", "disable-content-trust", "platform", "quiet"} {
		if !hasFlag(imagePullCommand, flag) {
			t.Errorf("flag --%s not registered on imagePullCommand", flag)
		}
	}
}

func TestImagePushFlags(t *testing.T) {
	for _, flag := range []string{"all-tags", "disable-content-trust", "quiet"} {
		if !hasFlag(imagePushCommand, flag) {
			t.Errorf("flag --%s not registered on imagePushCommand", flag)
		}
	}
}

func TestImageRemoveFlags(t *testing.T) {
	for _, flag := range []string{"force", "no-prune"} {
		if !hasFlag(imageRemoveCommand, flag) {
			t.Errorf("flag --%s not registered on imageRemoveCommand", flag)
		}
	}
}

func TestImageSaveFlags(t *testing.T) {
	if !hasFlag(imageSaveCommand, "output") {
		t.Error("flag --output not registered on imageSaveCommand")
	}
}

func TestImageLoadFlags(t *testing.T) {
	for _, flag := range []string{"input", "quiet"} {
		if !hasFlag(imageLoadCommand, flag) {
			t.Errorf("flag --%s not registered on imageLoadCommand", flag)
		}
	}
}
