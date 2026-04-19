package command

import "github.com/spf13/cobra"

var imageCommand = &cobra.Command{
	Use:     "image",
	Short:   "Manage images",
	Aliases: []string{"img"},
}

// ── subcommands ──────────────────────────────────────────────────────────────

var imageBuildCommand = &cobra.Command{
	Use:   "build [OPTIONS] PATH | URL | -",
	Short: "Build an image from a Jailfile",
	Args:  cobra.MaximumNArgs(1),
	RunE:  notImplemented,
}

var imageHistoryCommand = &cobra.Command{
	Use:   "history [OPTIONS] IMAGE",
	Short: "Show the history of an image",
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var imageImportCommand = &cobra.Command{
	Use:   "import [OPTIONS] file|URL|- [REPOSITORY[:TAG]]",
	Short: "Import the contents from a tarball to create a filesystem image",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  notImplemented,
}

var imageInspectCommand = &cobra.Command{
	Use:   "inspect [OPTIONS] IMAGE [IMAGE...]",
	Short: "Display detailed information on one or more images",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var imageLoadCommand = &cobra.Command{
	Use:   "load [OPTIONS]",
	Short: "Load an image from a tar archive or STDIN",
	RunE:  notImplemented,
}

var imageListCommand = &cobra.Command{
	Use:     "ls [OPTIONS] [REPOSITORY[:TAG]]",
	Aliases: []string{"list"},
	Short:   "List images",
	RunE:    notImplemented,
}

var imagePruneCommand = &cobra.Command{
	Use:   "prune [OPTIONS]",
	Short: "Remove unused images",
	RunE:  notImplemented,
}

var imagePullCommand = &cobra.Command{
	Use:   "pull [OPTIONS] NAME[:TAG|@DIGEST]",
	Short: "Download an image from a registry",
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var imagePushCommand = &cobra.Command{
	Use:   "push [OPTIONS] NAME[:TAG]",
	Short: "Upload an image to a registry",
	Args:  cobra.ExactArgs(1),
	RunE:  notImplemented,
}

var imageRemoveCommand = &cobra.Command{
	Use:     "rm [OPTIONS] IMAGE [IMAGE...]",
	Aliases: []string{"remove"},
	Short:   "Remove one or more images",
	Args:    cobra.MinimumNArgs(1),
	RunE:    notImplemented,
}

var imageSaveCommand = &cobra.Command{
	Use:   "save [OPTIONS] IMAGE [IMAGE...]",
	Short: "Save one or more images to a tar archive (streamed to STDOUT by default)",
	Args:  cobra.MinimumNArgs(1),
	RunE:  notImplemented,
}

var imageTagCommand = &cobra.Command{
	Use:   "tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]",
	Short: "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE",
	Args:  cobra.ExactArgs(2),
	RunE:  notImplemented,
}

// ── flag helpers (shared with top-level shorthands) ──────────────────────────

func addImageBuildFlags(target *cobra.Command) {
	target.Flags().StringSlice("add-host", nil, "Add a custom host-to-IP mapping (host:ip)")
	target.Flags().StringSlice("build-arg", nil, "Set build-time variables")
	target.Flags().String("cache-from", "", "Images to consider as cache sources")
	target.Flags().StringP("file", "f", "", "Name of the Jailfile (default: PATH/Jailfile)")
	target.Flags().StringSlice("label", nil, "Set metadata for an image")
	target.Flags().Bool("no-cache", false, "Do not use cache when building the image")
	target.Flags().String("platform", "", "Set platform if image is multi-platform capable")
	target.Flags().Bool("pull", false, "Always attempt to pull all referenced images")
	target.Flags().BoolP("quiet", "q", false, "Suppress the build output and print image ID on success")
	target.Flags().StringP("tag", "t", "", "Name and optionally a tag (format: name:tag)")
	target.Flags().String("target", "", "Set the target build stage to build")
}

func addImageListFlags(target *cobra.Command) {
	target.Flags().BoolP("all", "a", false, "Show all images (default hides intermediate images)")
	target.Flags().Bool("digests", false, "Show digests")
	target.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	target.Flags().String("format", "", "Format output using a custom template")
	target.Flags().Bool("no-trunc", false, "Don't truncate output")
	target.Flags().BoolP("quiet", "q", false, "Only show image IDs")
}

func init() {
	addImageBuildFlags(imageBuildCommand)
	addImageListFlags(imageListCommand)

	imageHistoryCommand.Flags().String("format", "", "Format output using a custom template")
	imageHistoryCommand.Flags().Bool("human", true, "Print sizes and dates in human readable format")
	imageHistoryCommand.Flags().Bool("no-trunc", false, "Don't truncate output")
	imageHistoryCommand.Flags().BoolP("quiet", "q", false, "Only show image IDs")

	imageImportCommand.Flags().StringSliceP("change", "c", nil, "Apply Jailfile instruction to the created image")
	imageImportCommand.Flags().StringP("message", "m", "", "Set commit message for imported image")
	imageImportCommand.Flags().String("platform", "", "Set platform if image is multi-platform capable")

	imageInspectCommand.Flags().StringP("format", "f", "", "Format output using a custom template")

	imageLoadCommand.Flags().StringP("input", "i", "", "Read from tar archive file, instead of STDIN")
	imageLoadCommand.Flags().BoolP("quiet", "q", false, "Suppress the load output")

	imagePruneCommand.Flags().BoolP("all", "a", false, "Remove all unused images, not just dangling ones")
	imagePruneCommand.Flags().String("filter", "", "Provide filter values (e.g. 'until=<timestamp>')")
	imagePruneCommand.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")

	imagePullCommand.Flags().BoolP("all-tags", "a", false, "Download all tagged images in the repository")
	imagePullCommand.Flags().Bool("disable-content-trust", true, "Skip image verification")
	imagePullCommand.Flags().String("platform", "", "Set platform if image is multi-platform capable")
	imagePullCommand.Flags().BoolP("quiet", "q", false, "Suppress verbose output")

	imagePushCommand.Flags().BoolP("all-tags", "a", false, "Push all tags of an image to the registry")
	imagePushCommand.Flags().Bool("disable-content-trust", true, "Skip image signing")
	imagePushCommand.Flags().BoolP("quiet", "q", false, "Suppress verbose output")

	imageRemoveCommand.Flags().BoolP("force", "f", false, "Force removal of the image")
	imageRemoveCommand.Flags().Bool("no-prune", false, "Do not delete untagged parents")

	imageSaveCommand.Flags().StringP("output", "o", "", "Write to a file, instead of STDOUT")

	imageCommand.AddCommand(
		imageBuildCommand,
		imageHistoryCommand,
		imageImportCommand,
		imageInspectCommand,
		imageLoadCommand,
		imageListCommand,
		imagePruneCommand,
		imagePullCommand,
		imagePushCommand,
		imageRemoveCommand,
		imageSaveCommand,
		imageTagCommand,
	)
	rootCommand.AddCommand(imageCommand)
}
