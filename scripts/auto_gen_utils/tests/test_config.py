import config
import os


# Note: If you have the DEBUG_DEXREQ_BRANCH_PREFIX set, this test will exit out
def test_branch_names_no_init_branches():
    if os.environ.get('DEBUG_DEXREQ_BRANCH_PREFIX'):
        return

    assert config.DEBUG_DEXREQ_BRANCH_PREFIX == ""
    assert config.BULK_PREVIEW_BRANCH_PREFIX == "auto-v2-preview-bulk"
    assert config.BULK_PREVIEW_BRANCH_PREFIX == 'auto-v2-preview-bulk'
    assert config.BULK_PUBLIC_BRANCH_PREFIX == 'auto-v2-public-bulk'
    assert config.INDIVIDUAL_PREVIEW_BRANCH_PREFIX == 'auto-v2-preview'
    assert config.INDIVIDUAL_PUBLIC_BRANCH_PREFIX == 'auto-v2-public'

    assert config.V1_BULK_PREVIEW_BRANCH_PREFIX == 'auto-preview-bulk'
    assert config.V1_INDIVIDUAL_PREVIEW_BRANCH_PREFIX == 'auto-preview'
    assert config.V1_INDIVIDUAL_PUBLIC_BRANCH_PREFIX == 'auto-public'


def test_branch_names(monkeypatch):
    previous = config.DEBUG_DEXREQ_BRANCH_PREFIX
    try:
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)

        config.init_branches()

        assert config.DEBUG_DEXREQ_BRANCH_PREFIX == ""
        assert config.BULK_PREVIEW_BRANCH_PREFIX == "auto-v2-preview-bulk"
        assert config.BULK_PREVIEW_BRANCH_PREFIX == 'auto-v2-preview-bulk'
        assert config.BULK_PUBLIC_BRANCH_PREFIX == 'auto-v2-public-bulk'
        assert config.INDIVIDUAL_PREVIEW_BRANCH_PREFIX == 'auto-v2-preview'
        assert config.INDIVIDUAL_PUBLIC_BRANCH_PREFIX == 'auto-v2-public'

        assert config.V1_BULK_PREVIEW_BRANCH_PREFIX == 'auto-preview-bulk'
        assert config.V1_INDIVIDUAL_PREVIEW_BRANCH_PREFIX == 'auto-preview'
        assert config.V1_INDIVIDUAL_PUBLIC_BRANCH_PREFIX == 'auto-public'
    finally:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", previous)
        config.init_branches()
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)


def test_branch_names_with_debug_prefix(monkeypatch):
    previous = config.DEBUG_DEXREQ_BRANCH_PREFIX
    try:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", "xyz-")

        config.init_branches()

        assert config.DEBUG_DEXREQ_BRANCH_PREFIX == "xyz-"
        assert config.BULK_PREVIEW_BRANCH_PREFIX == "xyz-auto-v2-preview-bulk"
        assert config.BULK_PREVIEW_BRANCH_PREFIX == 'xyz-auto-v2-preview-bulk'
        assert config.BULK_PUBLIC_BRANCH_PREFIX == 'xyz-auto-v2-public-bulk'
        assert config.INDIVIDUAL_PREVIEW_BRANCH_PREFIX == 'xyz-auto-v2-preview'
        assert config.INDIVIDUAL_PUBLIC_BRANCH_PREFIX == 'xyz-auto-v2-public'

        assert config.V1_BULK_PREVIEW_BRANCH_PREFIX == 'xyz-auto-preview-bulk'
        assert config.V1_INDIVIDUAL_PREVIEW_BRANCH_PREFIX == 'xyz-auto-preview'
        assert config.V1_INDIVIDUAL_PUBLIC_BRANCH_PREFIX == 'xyz-auto-public'
    finally:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", previous)
        config.init_branches()
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)


def test_branch_names_with_debug_prefix_empty_quotes(monkeypatch):
    previous = config.DEBUG_DEXREQ_BRANCH_PREFIX
    try:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", '""')

        config.init_branches()

        assert config.DEBUG_DEXREQ_BRANCH_PREFIX == ""
        assert config.BULK_PREVIEW_BRANCH_PREFIX == "auto-v2-preview-bulk"
        assert config.BULK_PREVIEW_BRANCH_PREFIX == 'auto-v2-preview-bulk'
        assert config.BULK_PUBLIC_BRANCH_PREFIX == 'auto-v2-public-bulk'
        assert config.INDIVIDUAL_PREVIEW_BRANCH_PREFIX == 'auto-v2-preview'
        assert config.INDIVIDUAL_PUBLIC_BRANCH_PREFIX == 'auto-v2-public'

        assert config.V1_BULK_PREVIEW_BRANCH_PREFIX == 'auto-preview-bulk'
        assert config.V1_INDIVIDUAL_PREVIEW_BRANCH_PREFIX == 'auto-preview'
        assert config.V1_INDIVIDUAL_PUBLIC_BRANCH_PREFIX == 'auto-public'
    finally:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", previous)
        config.init_branches()
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)


def test_branch_names_with_debug_prefix_with_quotes(monkeypatch):
    previous = config.DEBUG_DEXREQ_BRANCH_PREFIX
    try:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", '"xyz-"')

        config.init_branches()

        assert config.DEBUG_DEXREQ_BRANCH_PREFIX == "xyz-"
        assert config.BULK_PREVIEW_BRANCH_PREFIX == "xyz-auto-v2-preview-bulk"
        assert config.BULK_PREVIEW_BRANCH_PREFIX == 'xyz-auto-v2-preview-bulk'
        assert config.BULK_PUBLIC_BRANCH_PREFIX == 'xyz-auto-v2-public-bulk'
        assert config.INDIVIDUAL_PREVIEW_BRANCH_PREFIX == 'xyz-auto-v2-preview'
        assert config.INDIVIDUAL_PUBLIC_BRANCH_PREFIX == 'xyz-auto-v2-public'

        assert config.V1_BULK_PREVIEW_BRANCH_PREFIX == 'xyz-auto-preview-bulk'
        assert config.V1_INDIVIDUAL_PREVIEW_BRANCH_PREFIX == 'xyz-auto-preview'
        assert config.V1_INDIVIDUAL_PUBLIC_BRANCH_PREFIX == 'xyz-auto-public'
    finally:
        monkeypatch.setenv("DEBUG_DEXREQ_BRANCH_PREFIX", previous)
        config.init_branches()
        monkeypatch.delenv("DEBUG_DEXREQ_BRANCH_PREFIX", None)


# Note: If you have the DEXREQ_IGNORED_ISSUES set, this test will exit out
def test_should_ignore_issue_no_init():
    if os.environ.get('DEXREQ_IGNORED_ISSUES'):
        return

    assert config.DEXREQ_IGNORED_ISSUES == []
    assert not config.should_ignore_issue("ABC-1234")
    assert not config.should_ignore_issue("DEXREQ-4567")


def test_should_ignore_issue_names(monkeypatch):
    previous = ",".join(config.DEXREQ_IGNORED_ISSUES)
    try:
        monkeypatch.delenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, None)

        config.init_dexreq_ignored_issues()

        assert config.DEXREQ_IGNORED_ISSUES == []
        assert not config.should_ignore_issue("ABC-1234")
        assert not config.should_ignore_issue("DEXREQ-4567")
    finally:
        monkeypatch.setenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, previous)
        config.init_dexreq_ignored_issues()
        monkeypatch.delenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, None)


def test_should_ignore_issue_with_debug_prefix(monkeypatch):
    previous = ",".join(config.DEXREQ_IGNORED_ISSUES)
    try:
        monkeypatch.setenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, "  DEXREQ-1234 ,   DEXREQ-4567  ")

        config.init_dexreq_ignored_issues()

        assert config.DEXREQ_IGNORED_ISSUES == ["DEXREQ-1234", "DEXREQ-4567"]
        assert not config.should_ignore_issue("ABC-1234")
        assert config.should_ignore_issue("DEXREQ-1234")
        assert config.should_ignore_issue("DEXREQ-4567")
    finally:
        monkeypatch.setenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, previous)
        config.init_dexreq_ignored_issues()
        monkeypatch.delenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, None)


def test_should_ignore_issue_with_debug_prefix_single(monkeypatch):
    previous = ",".join(config.DEXREQ_IGNORED_ISSUES)
    try:
        monkeypatch.setenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, "  DEXREQ-1234 ")

        config.init_dexreq_ignored_issues()

        assert config.DEXREQ_IGNORED_ISSUES == ["DEXREQ-1234"]
        assert not config.should_ignore_issue("ABC-1234")
        assert config.should_ignore_issue("DEXREQ-1234")
        assert not config.should_ignore_issue("DEXREQ-4567")
    finally:
        monkeypatch.setenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, previous)
        config.init_dexreq_ignored_issues()
        monkeypatch.delenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, None)


def test_should_ignore_issue_with_debug_prefix_empty_quotes(monkeypatch):
    previous = ",".join(config.DEXREQ_IGNORED_ISSUES)
    try:
        monkeypatch.setenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, '""')

        config.init_dexreq_ignored_issues()

        assert config.DEXREQ_IGNORED_ISSUES == []
        assert not config.should_ignore_issue("ABC-1234")
        assert not config.should_ignore_issue("DEXREQ-4567")
    finally:
        monkeypatch.setenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, previous)
        config.init_dexreq_ignored_issues()
        monkeypatch.delenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, None)


def test_should_ignore_issue_with_debug_prefix_with_quotes(monkeypatch):
    previous = ",".join(config.DEXREQ_IGNORED_ISSUES)
    try:
        monkeypatch.setenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, '"  DEXREQ-1234 ,   DEXREQ-4567  "')

        config.init_dexreq_ignored_issues()

        assert config.DEXREQ_IGNORED_ISSUES == ["DEXREQ-1234", "DEXREQ-4567"]
        assert not config.should_ignore_issue("ABC-1234")
        assert config.should_ignore_issue("DEXREQ-1234")
        assert config.should_ignore_issue("DEXREQ-4567")
    finally:
        monkeypatch.setenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, previous)
        config.init_dexreq_ignored_issues()
        monkeypatch.delenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, None)


def test_should_ignore_issue_with_debug_prefix_with_quotes_single(monkeypatch):
    previous = ",".join(config.DEXREQ_IGNORED_ISSUES)
    try:
        monkeypatch.setenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, '"  DEXREQ-1234 "')

        config.init_dexreq_ignored_issues()

        assert config.DEXREQ_IGNORED_ISSUES == ["DEXREQ-1234"]
        assert not config.should_ignore_issue("ABC-1234")
        assert config.should_ignore_issue("DEXREQ-1234")
        assert not config.should_ignore_issue("DEXREQ-4567")
    finally:
        monkeypatch.setenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, previous)
        config.init_dexreq_ignored_issues()
        monkeypatch.delenv(config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME, None)
