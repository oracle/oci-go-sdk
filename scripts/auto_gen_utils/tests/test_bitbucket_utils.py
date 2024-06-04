import os

import config
import shared.bitbucket_utils


def test_printv():
    shared.bitbucket_utils.verbose = True
    shared.bitbucket_utils.printv("abc")
    shared.bitbucket_utils.printv(123)
    shared.bitbucket_utils.printv("")
    shared.bitbucket_utils.printv(None)
    shared.bitbucket_utils.printv(True)
    shared.bitbucket_utils.printv(False)


# Note: If you have the DEBUG_DEXREQ_BRANCH_PREFIX set, this test will exit out
def test_get_spec_pr_branch_reference_no_init_branches():
    if os.environ.get('DEBUG_DEXREQ_BRANCH_PREFIX'):
        return

    assert shared.bitbucket_utils.get_spec_pr_branch_reference("preview", "suffix") == "refs/heads/spec-auto-v2-preview-suffix-diff"
    assert shared.bitbucket_utils.get_spec_pr_branch_reference("public", "suffix") == "refs/heads/spec-auto-v2-public-suffix-diff"
    assert shared.bitbucket_utils.get_spec_pr_branch_reference("preview", "suffix2") == "refs/heads/spec-auto-v2-preview-suffix2-diff"
    assert shared.bitbucket_utils.get_spec_pr_branch_reference("public", "suffix2") == "refs/heads/spec-auto-v2-public-suffix2-diff"


def test_get_spec_pr_branch_reference(monkeypatch):
    previous = config.DEBUG_DEXREQ_BRANCH_PREFIX
    try:
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)

        config.init_branches()

        assert shared.bitbucket_utils.get_spec_pr_branch_reference("preview", "suffix") == "refs/heads/spec-auto-v2-preview-suffix-diff"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("public", "suffix") == "refs/heads/spec-auto-v2-public-suffix-diff"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("preview", "suffix2") == "refs/heads/spec-auto-v2-preview-suffix2-diff"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("public", "suffix2") == "refs/heads/spec-auto-v2-public-suffix2-diff"
    finally:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", previous)
        config.init_branches()
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)


def test_get_spec_pr_branch_reference_with_debug_prefix(monkeypatch):
    previous = config.DEBUG_DEXREQ_BRANCH_PREFIX
    try:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", "xyz-")

        config.init_branches()

        assert config.DEBUG_DEXREQ_BRANCH_PREFIX == "xyz-"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("preview", "suffix") == "refs/heads/spec-xyz-auto-v2-preview-suffix-diff"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("public", "suffix") == "refs/heads/spec-xyz-auto-v2-public-suffix-diff"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("preview", "suffix2") == "refs/heads/spec-xyz-auto-v2-preview-suffix2-diff"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("public", "suffix2") == "refs/heads/spec-xyz-auto-v2-public-suffix2-diff"
    finally:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", previous)
        config.init_branches()
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)


def test_get_spec_pr_branch_reference_with_debug_prefix_empty_quotes(monkeypatch):
    previous = config.DEBUG_DEXREQ_BRANCH_PREFIX
    try:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", '""')

        config.init_branches()

        assert config.DEBUG_DEXREQ_BRANCH_PREFIX == ""
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("preview", "suffix") == "refs/heads/spec-auto-v2-preview-suffix-diff"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("public", "suffix") == "refs/heads/spec-auto-v2-public-suffix-diff"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("preview", "suffix2") == "refs/heads/spec-auto-v2-preview-suffix2-diff"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("public", "suffix2") == "refs/heads/spec-auto-v2-public-suffix2-diff"
    finally:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", previous)
        config.init_branches()
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)


def test_get_spec_pr_branch_reference_with_debug_prefix_with_quotes(monkeypatch):
    previous = config.DEBUG_DEXREQ_BRANCH_PREFIX
    try:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", '"xyz-"')

        config.init_branches()

        assert config.DEBUG_DEXREQ_BRANCH_PREFIX == "xyz-"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("preview", "suffix") == "refs/heads/spec-xyz-auto-v2-preview-suffix-diff"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("public", "suffix") == "refs/heads/spec-xyz-auto-v2-public-suffix-diff"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("preview", "suffix2") == "refs/heads/spec-xyz-auto-v2-preview-suffix2-diff"
        assert shared.bitbucket_utils.get_spec_pr_branch_reference("public", "suffix2") == "refs/heads/spec-xyz-auto-v2-public-suffix2-diff"
    finally:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", previous)
        config.init_branches()
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)
