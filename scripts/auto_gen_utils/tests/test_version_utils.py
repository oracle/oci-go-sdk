from shared.version_utils import is_version_not_acceptable, is_version_increasing


def test_is_version_not_acceptable():
    assert is_version_not_acceptable("1.2.3") is None
    assert is_version_not_acceptable("1.2") is None
    assert is_version_not_acceptable("1.2.3-PREVIEW") is None
    assert is_version_not_acceptable("1.2.3-releasePreview") is None
    assert is_version_not_acceptable("1.2.3-4") is None
    assert is_version_not_acceptable("1.2.3-4-PREVIEW") is None

    assert is_version_not_acceptable("1.2.3-20180409.232938-55") is not None
    assert is_version_not_acceptable("1.2-20180409.232938-55") is not None
    assert is_version_not_acceptable("1.2.3-PREVIEW-20180409.232938-55") is not None
    assert is_version_not_acceptable("1.2.3-releasePreview-20180409.232938-55") is not None
    assert is_version_not_acceptable("1.2.3-4-20180409.232938-55") is not None
    assert is_version_not_acceptable("1.2.3-4-PREVIEW-20180409.232938-55") is not None

    assert is_version_not_acceptable("1.2.3-SNAPSHOT") is not None
    assert is_version_not_acceptable("1.2-SNAPSHOT") is not None
    assert is_version_not_acceptable("1.2.3-PREVIEW-SNAPSHOT") is not None
    assert is_version_not_acceptable("1.2.3-releasePreview-SNAPSHOT") is not None
    assert is_version_not_acceptable("1.2.3-4-SNAPSHOT") is not None
    assert is_version_not_acceptable("1.2.3-4-PREVIEW-SNAPSHOT") is not None

    assert is_version_not_acceptable("1.a.3") is not None
    assert is_version_not_acceptable("1.a") is not None
    assert is_version_not_acceptable("1.2.3-foobar") is not None
    assert is_version_not_acceptable("1.2.3-1-2") is not None
    assert is_version_not_acceptable("1.2.3-x-PREVIEW") is not None

    assert is_version_not_acceptable("1.a.3-20180409.232938-55") is not None
    assert is_version_not_acceptable("1.a-20180409.232938-55") is not None
    assert is_version_not_acceptable("1.2.3-foobar-20180409.232938-55") is not None
    assert is_version_not_acceptable("1.2.3-1-2-20180409.232938-55") is not None
    assert is_version_not_acceptable("1-4-20180409.232938-55") is not None
    assert is_version_not_acceptable("1.2.3-x-PREVIEW-20180409.232938-55") is not None

    assert is_version_not_acceptable("1.3-PREVIEW") is not None
    assert is_version_not_acceptable("1.3-releasePreview") is not None
    assert is_version_not_acceptable("1.3-4-PREVIEW") is not None
    assert is_version_not_acceptable("1.2.2.3-PREVIEW") is not None
    assert is_version_not_acceptable("1.2.2.3-releasePreview") is not None
    assert is_version_not_acceptable("1.2.2.3-4-PREVIEW") is not None

    assert is_version_not_acceptable("1") is None
    assert is_version_not_acceptable("12") is None
    assert is_version_not_acceptable("123") is None
    assert is_version_not_acceptable("1-4") is None


def test_is_version_increasing():
    assert is_version_increasing("1", "2")
    assert not is_version_increasing("2", "1")

    assert is_version_increasing("1.0", "1.1")
    assert is_version_increasing("1.0", "2.0")
    assert is_version_increasing("1.1", "2.0")
    assert not is_version_increasing("1.1", "1.0")
    assert not is_version_increasing("2.0", "1.0")
    assert not is_version_increasing("2.0", "1.1")

    assert is_version_increasing("1-PREVIEW", "1")
    assert not is_version_increasing("1", "1-PREVIEW")
    assert not is_version_increasing("1", "1")
    assert not is_version_increasing("1-PREVIEW", "1-PREVIEW")

    assert is_version_increasing("1.0", "1.0.1")
    assert not is_version_increasing("1.0.1", "1.0")
    assert not is_version_increasing("1.0.0", "1.0")
