from shared.buildsvc_tc_compatibility import parse_build_id, parse_vcs_root, build_log_link, build_artifacts_link


def test_parse_build_id():
    is_build_service, build_id, buildsvc_branch, buildsvc_build_number = parse_build_id("1234")
    assert is_build_service is False
    assert build_id == "1234"
    assert buildsvc_branch is None
    assert buildsvc_build_number is None

    is_build_service, build_id, buildsvc_branch, buildsvc_build_number = parse_build_id("buildsvc-1234")
    assert is_build_service is True
    assert build_id == "buildsvc-1234"
    assert buildsvc_branch is None
    assert buildsvc_build_number == "1234"

    is_build_service, build_id, buildsvc_branch, buildsvc_build_number = parse_build_id("buildsvc-prototype-1234")
    assert is_build_service is True
    assert build_id == "buildsvc-prototype-1234"
    assert buildsvc_branch == "prototype"
    assert buildsvc_build_number == "1234"


def test_parse_vcs_root():
    project, repo = parse_vcs_root("xyz")
    assert project is None
    assert repo is None

    project, repo = parse_vcs_root("ssh://git@bitbucket.oci.oraclecorp.com:7999/sdk/dexreq-surfaces.git")
    assert project == "SDK"
    assert repo == "dexreq-surfaces"


def test_build_log_link_no_vcs_root_env_var(monkeypatch):
    monkeypatch.delenv("BLD_VCS_ROOT", None)
    monkeypatch.delenv("BLD_VSC_ROOT", None)
    monkeypatch.delenv("BLD_STEP", None)

    assert build_log_link("1234") == "[here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234]"
    assert build_log_link("1234", text="click here") == "[click here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234]"

    assert build_log_link("buildsvc-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234]"

    assert build_log_link("buildsvc-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234]"

    assert build_log_link("buildsvc-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/foo]"

    assert build_log_link("buildsvc-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234]"

    monkeypatch.setenv("BLD_STEP", "xyz")

    assert build_log_link("1234") == "[here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234]"
    assert build_log_link("1234", text="click here") == "[click here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234]"

    assert build_log_link("buildsvc-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/xyz]"

    assert build_log_link("buildsvc-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/xyz]"

    assert build_log_link("buildsvc-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/foo]"

    assert build_log_link("buildsvc-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234]"

    monkeypatch.delenv("BLD_STEP", None)
    monkeypatch.delenv("BLD_VCS_ROOT", None)
    monkeypatch.delenv("BLD_VSC_ROOT", None)


def test_build_log_link_with_vcs_root_env_var(monkeypatch):
    monkeypatch.delenv("BLD_VCS_ROOT", None)
    monkeypatch.delenv("BLD_VSC_ROOT", None)
    monkeypatch.delenv("BLD_STEP", None)
    monkeypatch.setenv("BLD_VCS_ROOT", "ssh://git@bitbucket.oci.oraclecorp.com:7999/otherproject/otherrepo.git")

    assert build_log_link("1234") == "[here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234]"
    assert build_log_link("1234", text="click here") == "[click here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234]"

    assert build_log_link("buildsvc-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"

    assert build_log_link("buildsvc-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"

    assert build_log_link("buildsvc-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo]"

    assert build_log_link("buildsvc-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"

    monkeypatch.setenv("BLD_STEP", "xyz")

    assert build_log_link("1234") == "[here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234]"
    assert build_log_link("1234", text="click here") == "[click here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234]"

    assert build_log_link("buildsvc-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz]"

    assert build_log_link("buildsvc-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz]"

    assert build_log_link("buildsvc-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo]"

    assert build_log_link("buildsvc-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"

    monkeypatch.delenv("BLD_STEP", None)
    monkeypatch.delenv("BLD_VCS_ROOT", None)
    monkeypatch.delenv("BLD_VSC_ROOT", None)


def test_build_log_link_with_vsc_root_env_var(monkeypatch):
    monkeypatch.delenv("BLD_VCS_ROOT", None)
    monkeypatch.delenv("BLD_VSC_ROOT", None)
    monkeypatch.delenv("BLD_STEP", None)
    monkeypatch.setenv("BLD_VSC_ROOT", "ssh://git@bitbucket.oci.oraclecorp.com:7999/otherproject/otherrepo.git")

    assert build_log_link("1234") == "[here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234]"
    assert build_log_link("1234", text="click here") == "[click here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234]"

    assert build_log_link("buildsvc-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"

    assert build_log_link("buildsvc-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"

    assert build_log_link("buildsvc-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo]"

    assert build_log_link("buildsvc-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"

    monkeypatch.setenv("BLD_STEP", "xyz")

    assert build_log_link("1234") == "[here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234]"
    assert build_log_link("1234", text="click here") == "[click here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234]"

    assert build_log_link("buildsvc-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz]"

    assert build_log_link("buildsvc-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/xyz]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz]"

    assert build_log_link("buildsvc-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/foo]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo]"

    assert build_log_link("buildsvc-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234]"
    assert build_log_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234]"
    assert build_log_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234]"

    monkeypatch.delenv("BLD_STEP", None)
    monkeypatch.delenv("BLD_VCS_ROOT", None)
    monkeypatch.delenv("BLD_VSC_ROOT", None)


# Artifacts

def test_build_artifacts_link_no_vcs_root_env_var(monkeypatch):
    monkeypatch.delenv("BLD_VCS_ROOT", None)
    monkeypatch.delenv("BLD_VSC_ROOT", None)
    monkeypatch.delenv("BLD_STEP", None)

    assert build_artifacts_link("1234") == "[here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234&tab=artifacts]"
    assert build_artifacts_link("1234", text="click here") == "[click here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/foo&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234&tab=artifacts]"

    monkeypatch.setenv("BLD_STEP", "xyz")

    assert build_artifacts_link("1234") == "[here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234&tab=artifacts]"
    assert build_artifacts_link("1234", text="click here") == "[click here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234/steps/foo&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/master/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/dexreq-surfaces/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/java-sdk/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/SDK/projects/dexreq-surfaces/branches/prototype/builds/1234&tab=artifacts]"

    monkeypatch.delenv("BLD_STEP", None)
    monkeypatch.delenv("BLD_VCS_ROOT", None)
    monkeypatch.delenv("BLD_VSC_ROOT", None)


def test_build_artifacts_link_with_vcs_root_env_var(monkeypatch):
    monkeypatch.delenv("BLD_VCS_ROOT", None)
    monkeypatch.delenv("BLD_VSC_ROOT", None)
    monkeypatch.delenv("BLD_STEP", None)
    monkeypatch.setenv("BLD_VCS_ROOT", "ssh://git@bitbucket.oci.oraclecorp.com:7999/otherproject/otherrepo.git")

    assert build_artifacts_link("1234") == "[here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234&tab=artifacts]"
    assert build_artifacts_link("1234", text="click here") == "[click here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"

    monkeypatch.setenv("BLD_STEP", "xyz")

    assert build_artifacts_link("1234") == "[here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234&tab=artifacts]"
    assert build_artifacts_link("1234", text="click here") == "[click here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"

    monkeypatch.delenv("BLD_STEP", None)
    monkeypatch.delenv("BLD_VCS_ROOT", None)
    monkeypatch.delenv("BLD_VSC_ROOT", None)


def test_build_artifacts_link_with_vsc_root_env_var(monkeypatch):
    monkeypatch.delenv("BLD_VCS_ROOT", None)
    monkeypatch.delenv("BLD_VSC_ROOT", None)
    monkeypatch.delenv("BLD_STEP", None)
    monkeypatch.setenv("BLD_VSC_ROOT", "ssh://git@bitbucket.oci.oraclecorp.com:7999/otherproject/otherrepo.git")

    assert build_artifacts_link("1234") == "[here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234&tab=artifacts]"
    assert build_artifacts_link("1234", text="click here") == "[click here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"

    monkeypatch.setenv("BLD_STEP", "xyz")

    assert build_artifacts_link("1234") == "[here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234&tab=artifacts]"
    assert build_artifacts_link("1234", text="click here") == "[click here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId=1234&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", project="CLI") == "[here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", repo="java-sdk") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", default_branch="master") == "[here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/xyz&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234/steps/foo&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="foo") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234/steps/foo&tab=artifacts]"

    assert build_artifacts_link("buildsvc-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/main/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/master/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", project="CLI", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/CLI/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", repo="java-sdk", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/java-sdk/branches/prototype/builds/1234&tab=artifacts]"
    assert build_artifacts_link("buildsvc-prototype-1234", text="click here", default_branch="master", build_step_name="") == "[click here|https://devops.oci.oraclecorp.com/build/teams/OTHERPROJECT/projects/otherrepo/branches/prototype/builds/1234&tab=artifacts]"

    monkeypatch.delenv("BLD_STEP", None)
    monkeypatch.delenv("BLD_VCS_ROOT", None)
    monkeypatch.delenv("BLD_VSC_ROOT", None)
