import os

import clean_auto_branches


# Note: If you have the DEBUG_DEXREQ_BRANCH_PREFIX set, this test will exit out
def test_branch_names_no_init_branches():
    if os.environ.get('DEBUG_DEXREQ_BRANCH_PREFIX'):
        return

    assert clean_auto_branches.GENERATED_AUTO_PREVIEW_BRANCH_PATTERN == 'refs/remotes/origin/generated-auto-v2-preview-*'
    assert clean_auto_branches.AUTO_PREVIEW_BRANCH_PATTERN == 'refs/remotes/origin/auto-v2-preview-*'
    assert clean_auto_branches.GENERATED_AUTO_PUBLIC_BRANCH_PATTERN == 'refs/remotes/origin/generated-auto-v2-public-*'
    assert clean_auto_branches.AUTO_PUBLIC_BRANCH_PATTERN == 'refs/remotes/origin/auto-v2-public-*'

    assert clean_auto_branches.AUTO_PREVIEW_BRANCH_PREFIX == 'auto-v2-preview-'
    assert clean_auto_branches.AUTO_PUBLIC_BRANCH_PREFIX == 'auto-v2-public-'


def test_branch_names(monkeypatch):
    previous = clean_auto_branches.DEBUG_DEXREQ_BRANCH_PREFIX
    try:
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)

        clean_auto_branches.init_branches()

        assert clean_auto_branches.DEBUG_DEXREQ_BRANCH_PREFIX == ""
        assert clean_auto_branches.GENERATED_AUTO_PREVIEW_BRANCH_PATTERN == 'refs/remotes/origin/generated-auto-v2-preview-*'
        assert clean_auto_branches.AUTO_PREVIEW_BRANCH_PATTERN == 'refs/remotes/origin/auto-v2-preview-*'
        assert clean_auto_branches.GENERATED_AUTO_PUBLIC_BRANCH_PATTERN == 'refs/remotes/origin/generated-auto-v2-public-*'
        assert clean_auto_branches.AUTO_PUBLIC_BRANCH_PATTERN == 'refs/remotes/origin/auto-v2-public-*'

        assert clean_auto_branches.AUTO_PREVIEW_BRANCH_PREFIX == 'auto-v2-preview-'
        assert clean_auto_branches.AUTO_PUBLIC_BRANCH_PREFIX == 'auto-v2-public-'
    finally:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", previous)
        clean_auto_branches.init_branches()
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)


def test_branch_names_with_debug_prefix(monkeypatch):
    previous = clean_auto_branches.DEBUG_DEXREQ_BRANCH_PREFIX
    try:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", "xyz-")

        clean_auto_branches.init_branches()

        assert clean_auto_branches.DEBUG_DEXREQ_BRANCH_PREFIX == "xyz-"
        assert clean_auto_branches.GENERATED_AUTO_PREVIEW_BRANCH_PATTERN == 'refs/remotes/origin/generated-xyz-auto-v2-preview-*'
        assert clean_auto_branches.AUTO_PREVIEW_BRANCH_PATTERN == 'refs/remotes/origin/xyz-auto-v2-preview-*'
        assert clean_auto_branches.GENERATED_AUTO_PUBLIC_BRANCH_PATTERN == 'refs/remotes/origin/generated-xyz-auto-v2-public-*'
        assert clean_auto_branches.AUTO_PUBLIC_BRANCH_PATTERN == 'refs/remotes/origin/xyz-auto-v2-public-*'

        assert clean_auto_branches.AUTO_PREVIEW_BRANCH_PREFIX == 'xyz-auto-v2-preview-'
        assert clean_auto_branches.AUTO_PUBLIC_BRANCH_PREFIX == 'xyz-auto-v2-public-'
    finally:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", previous)
        clean_auto_branches.init_branches()
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)


def test_branch_names_with_debug_prefix_empty_quotes(monkeypatch):
    previous = clean_auto_branches.DEBUG_DEXREQ_BRANCH_PREFIX
    try:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", '""')

        clean_auto_branches.init_branches()

        assert clean_auto_branches.DEBUG_DEXREQ_BRANCH_PREFIX == ""
        assert clean_auto_branches.GENERATED_AUTO_PREVIEW_BRANCH_PATTERN == 'refs/remotes/origin/generated-auto-v2-preview-*'
        assert clean_auto_branches.AUTO_PREVIEW_BRANCH_PATTERN == 'refs/remotes/origin/auto-v2-preview-*'
        assert clean_auto_branches.GENERATED_AUTO_PUBLIC_BRANCH_PATTERN == 'refs/remotes/origin/generated-auto-v2-public-*'
        assert clean_auto_branches.AUTO_PUBLIC_BRANCH_PATTERN == 'refs/remotes/origin/auto-v2-public-*'

        assert clean_auto_branches.AUTO_PREVIEW_BRANCH_PREFIX == 'auto-v2-preview-'
        assert clean_auto_branches.AUTO_PUBLIC_BRANCH_PREFIX == 'auto-v2-public-'
    finally:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", previous)
        clean_auto_branches.init_branches()
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)
        

def test_branch_names_with_debug_prefix_with_quotes(monkeypatch):
    previous = clean_auto_branches.DEBUG_DEXREQ_BRANCH_PREFIX
    try:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", '"xyz-"')

        clean_auto_branches.init_branches()

        assert clean_auto_branches.DEBUG_DEXREQ_BRANCH_PREFIX == "xyz-"
        assert clean_auto_branches.GENERATED_AUTO_PREVIEW_BRANCH_PATTERN == 'refs/remotes/origin/generated-xyz-auto-v2-preview-*'
        assert clean_auto_branches.AUTO_PREVIEW_BRANCH_PATTERN == 'refs/remotes/origin/xyz-auto-v2-preview-*'
        assert clean_auto_branches.GENERATED_AUTO_PUBLIC_BRANCH_PATTERN == 'refs/remotes/origin/generated-xyz-auto-v2-public-*'
        assert clean_auto_branches.AUTO_PUBLIC_BRANCH_PATTERN == 'refs/remotes/origin/xyz-auto-v2-public-*'

        assert clean_auto_branches.AUTO_PREVIEW_BRANCH_PREFIX == 'xyz-auto-v2-preview-'
        assert clean_auto_branches.AUTO_PUBLIC_BRANCH_PREFIX == 'xyz-auto-v2-public-'
    finally:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", previous)
        clean_auto_branches.init_branches()
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)
