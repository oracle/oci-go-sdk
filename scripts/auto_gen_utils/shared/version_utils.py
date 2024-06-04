import re


def is_int(s):
    try:
        int(s)
        return True
    except ValueError:
        return False


preview_markers = ["-RELEASEPREVIEW", "-PREVIEW"]


def is_version_increasing(old_version, version):
    # Versioning schemes we have to deal with:
    # 0.1.9
    # 0.0.15-releasePreview < 0.0.15 (database)
    # 0.2.17-PREVIEW < 0.2.17 (objectstorage, identity)

    old_v = old_version.upper()
    new_v = version.upper()

    old_has_preview_marker = False
    new_has_preview_marker = False
    for marker in preview_markers:
        if marker in old_v:
            old_has_preview_marker = True
            old_v = old_v.replace(marker, '')
        if marker in new_v:
            new_has_preview_marker = True
            new_v = new_v.replace(marker, '')

    # Now none of the versions have a preview marker anymore

    old_v = old_v.replace('-', '.')
    new_v = new_v.replace('-', '.')

    print('Old version: {} (preview marker? {})'.format(old_v, old_has_preview_marker))
    print('New version: {} (preview marker? {})'.format(new_v, new_has_preview_marker))

    old_parts = old_v.split('.')
    new_parts = new_v.split('.')

    common_length = min(len(old_parts), len(new_parts))

    for index in range(0, common_length):
        old_part = old_parts[index]
        new_part = new_parts[index]

        if is_int(old_part) and is_int(new_part):
            old_part = int(old_part)
            new_part = int(new_part)

        if old_part < new_part:
            # Old version is less
            return True
        elif new_part < old_part:
            # New version is less, this is a regression in version.
            return False

    # If we are back here, then the common segment has been all the same.
    # If the lengths are the same, then it comes down to the preview flag.
    if len(old_parts) == len(new_parts):
        if old_has_preview_marker and not new_has_preview_marker:
            # old version was a preview, new version is not.
            return True
        elif new_has_preview_marker and not old_has_preview_marker:
            # old version was not a preview, new version is. This is a regression in version.
            return False
        else:
            # either both were previews or both were not -- same version
            return False

    # Otherwise, one of them is longer. The longer version is the newer one.
    if len(old_parts) > len(new_parts):
        # The old version was longer. This is a regression in version.
        return False

    return True


# Result is None if the version _is_ acceptible
# otherwise, the error message is returned
def is_version_not_acceptable(version):
    # see https://confluence.oci.oraclecorp.com/display/DEX/Creating+a+Spec+Artifact+in+Artifactory#CreatingaSpecArtifactinArtifactory-VersionNumberSchemes
    #
    # Recommended:
    # 0.0.1
    #
    # Also acceptable:
    #
    # 1
    # 1.2
    # 1-2
    # 1.2.3-PREVIEW
    # 1.2.3-releasePreview
    # 1.2.3-4
    # 1.2.3-4-PREVIEW
    #
    # Not acceptable:
    # <anything>-SNAPSHOT
    #
    # But acceptable form for all the above:
    # <version>-20180409.232938-5
    original = version

    if version.upper().endswith("-SNAPSHOT"):
        return 'Version not acceptable for SDK generation: "{}" is a snapshot version.'.format(original)

    timed_snapshot_match = re.search("-[0-9]{4}[01][0-9][0123][0-9].[012][0-9][0-5][0-9][0-5][0-9]-[0-9]*$", version)
    if timed_snapshot_match:
        return 'Version not acceptable for SDK generation: "{}" is a timed snapshot version.'.format(original)

    dash_count = version.count("-")
    if dash_count > 2:
        return 'Version not acceptable for SDK generation: "{}" has more than 2 dashes'.format(version)

    suffix = None
    last_dash = version.rfind("-")
    if last_dash != -1:
        suffix = version[last_dash:]

        if not suffix.upper() in preview_markers and not is_int(suffix[1:]):
            # We do want to use the version without the timed snapshot here, not `original`
            return 'Version not acceptable for SDK generation: "{}" does not end in {}, or a build number'.format(version,
                ", ".join(preview_markers))

        second_to_last_dash = version.rfind("-", 0, last_dash)
        if second_to_last_dash != -1:
            # If there are two dashes, then the suffix has to be "-PREVIEW"
            # and the part between dashes has to be an integer build number
            between_dashes = version[second_to_last_dash + 1:last_dash]
            if not suffix.upper() == "-PREVIEW":
                return 'Version not acceptable for SDK generation: "{}" has two dashes, but it does not end with "-PREVIEW"'.format(version)
            if not is_int(between_dashes):
                return 'Version not acceptable for SDK generation: "{}" has two dashes, but the part between the dashes is not a build number'.format(version)
        version = version[0:version.find("-")]

    if dash_count == 1:
        dash_comment = " (after removing anything after the dash)"
    elif dash_count > 1:
        dash_comment = " (after removing anything after the first dash)"
    else:
        dash_comment = ""

    if not re.match(r"^[0-9]{1,}(\.[0-9]{1,}(\.[0-9]{1,})?)?$", version):
        return 'Version not acceptable for SDK generation: "{}"{} does not fit the patterns for one-, two-, or three-part numeric versions ("123", 1.2" or "1.2.3")'.format(version,
            dash_comment)

    if suffix and suffix.upper() in preview_markers:
        # If we had a preview marker, it has to be a three-part version
        # 1.2.3-PREVIEW
        # 1.2.3-releasePreview
        # 1.2.3-4-PREVIEW
        if version.count(".") != 2:
            return 'Version not acceptable for SDK generation: "{}"{} has suffix "{}", but is not a three-part numeric version ("1.2.3")'.format(version,
                dash_comment, suffix)

    return None
