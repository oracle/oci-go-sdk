import os
import re


def parse_build_id(build_id):
    is_build_service = False
    buildsvc_build_number = None
    buildsvc_branch = None

    search_result = re.search("^buildsvc(-(.*))?-([0-9]+)*$", build_id)
    if search_result:
        is_build_service = True
        # - group 2 is branch (optional)
        # - group 3 is build number
        buildsvc_build_number = search_result.group(3)
        if search_result.group(2):
            buildsvc_branch = search_result.group(2)

    return is_build_service, build_id, buildsvc_branch, buildsvc_build_number


def parse_vcs_root(vcs_root):
    # looks like ssh://git@bitbucket.oci.oraclecorp.com:7999/sdk/dexreq-surfaces.git
    project = None
    repo = None
    if vcs_root.startswith("ssh://git@bitbucket.oci.oraclecorp.com:7999/"):
        search_result = re.search("^ssh://git@bitbucket.oci.oraclecorp.com:7999/([^/]*)/([^/]*).git$", vcs_root, re.IGNORECASE)

        if search_result:
            project = search_result.group(1)
            repo = search_result.group(2)

    if project:
        project = project.upper()

    return project, repo


# build_step_name=None means use value from os.environ.get('BLD_STEP')
# If you don't want to provide a build_step_name, use build_step_name=""
#
# project=None or repo=None means use the value from os.environ.get('BLD_VCS_ROOT')
# or os.environ.get('BLD_VSC_ROOT') (yes, VSC_ROOT... https://jira.oci.oraclecorp.com/browse/BLD-3445)
def build_log_link(build_id, text="here", project=None, repo=None, default_branch="main", build_step_name=None, additional_url_parts=""):
    vcs_root = os.environ.get('BLD_VCS_ROOT') or os.environ.get('BLD_VSC_ROOT')

    if project is None:
        if vcs_root:
            project, r = parse_vcs_root(vcs_root)

        if not project:
            project = "SDK"

    if repo is None:
        if vcs_root:
            p, repo = parse_vcs_root(vcs_root)

        if not repo:
            repo = "dexreq-surfaces"

    is_build_service, build_id, buildsvc_branch, buildsvc_build_number = parse_build_id(build_id)
    if is_build_service:
        if buildsvc_branch is None:
            buildsvc_branch = default_branch
        build_step_part = ""
        if build_step_name is None:
            build_step_name = os.environ.get('BLD_STEP')
        if build_step_name:
            build_step_part = "/steps/{build_step_name}".format(build_step_name=build_step_name)
        return "[{text}|https://devops.oci.oraclecorp.com/build/teams/{project}/projects/{repo}/branches/{buildsvc_branch}/builds/{buildsvc_build_number}{build_step_part}{additional_url_parts}]".format(
            text=text,
            project=project,
            repo=repo,
            buildsvc_branch=buildsvc_branch,
            buildsvc_build_number=buildsvc_build_number,
            build_step_part=build_step_part,
            additional_url_parts=additional_url_parts
        )

    return "[{text}|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId={build_id}{additional_url_parts}]".format(
        text=text,
        build_id=build_id,
        additional_url_parts=additional_url_parts
    )


# build_step_name=None means use value from os.environ.get('BLD_STEP')
# If you don't want to provide a build_step_name, use build_step_name=""
#
# project=None or repo=None means use the value from os.environ.get('BLD_VCS_ROOT')
# or os.environ.get('BLD_VSC_ROOT') (yes, VSC_ROOT... https://jira.oci.oraclecorp.com/browse/BLD-3445)
def build_artifacts_link(build_id, text="here", project=None, repo=None, default_branch="main", build_step_name=None):
    return build_log_link(build_id, text=text, project=project, repo=repo, default_branch=default_branch, build_step_name=build_step_name, additional_url_parts="&tab=artifacts")
