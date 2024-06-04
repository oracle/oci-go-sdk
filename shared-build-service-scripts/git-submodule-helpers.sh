#!/bin/bash

# Usage:
#
# source git-submodule-helpers.sh
#
# or
#
# . git-submodule-helpers.sh
#

SHARED_SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Usage:
#
# git_submodule_update_all
#
# To update all submodules, if all the submodules have a tracking branch set.
# Must have this "branch = ..." setting:
# cat .gitmodules
# [submodule "bmc-sdk-swagger-validator/test/python-sdk-preview-submodule"]
#         path = bmc-sdk-swagger-validator/test/python-sdk
#         url = ssh://git@bitbucket.oci.oraclecorp.com:7999/sdk/python-sdk.git
#         branch = preview
function git_submodule_update_all() {
	# Check that 
	if ! git submodule foreach 'git config -f ${toplevel}/.gitmodules --get submodule."$name".branch' > /dev/null 2>&1; then
		echo "Not all submodules had tracking branches"
		return 1
	fi
	git submodule foreach 'git pull origin `git config -f ${toplevel}/.gitmodules --get submodule."$name".branch`'
}

# Usage:
#
# git_submodule_update 'path/to/submodule'
#
# To update a single submodule, if it has a tracking branch set.
# Must have this "branch = ..." setting:
# cat .gitmodules
# [submodule "bmc-sdk-swagger-validator/test/python-sdk-preview-submodule"]
#         path = bmc-sdk-swagger-validator/test/python-sdk
#         url = ssh://git@bitbucket.oci.oraclecorp.com:7999/sdk/python-sdk.git
#         branch = preview
function git_submodule_update() {
	path_to_submodule="$1"

	if [ -z ${BLD_INPUT_DIR+x} ]; then
    	echo "BLD_INPUT_DIR is unset"
    	return 1
	fi

	if [ ! -d ${path_to_submodule} ]; then
		echo "Directory ${path_to_submodule} does not exist"
		return 1
	fi

	# Determine the path from the toplevel Git directory to the submodule
	cd ${path_to_submodule}
	abs_path_to_submodule=`pwd`
	cd - > /dev/null
	rel_path_to_submodule=`${SHARED_SCRIPT_DIR}/relpath.sh ${BLD_INPUT_DIR} ${abs_path_to_submodule}`
	# echo "rel_path_to_submodule: $rel_path_to_submodule"
	
	# Determine the name of the submodule, based on the rel_path_to_submodule
	cd ${BLD_INPUT_DIR}
	submodule_name=`git config -f .gitmodules --get-regexp "submodule\..*\.path" "^${rel_path_to_submodule}$" | sed 's_^submodule\.\(.*\)\.path .*$_\1_'`
	cd - > /dev/null
	# echo "submodule_name: $submodule_name"
	if [ "" == "${submodule_name}" ]; then
    	echo "Could not find a submodule in directory ${path_to_submodule}"
    	return 1
	fi

	# Determine the tracking branch of the submodule, based on submodule_name
	cd ${BLD_INPUT_DIR}
	submodule_tracking_branch=`git config -f .gitmodules --get submodule."${submodule_name}".branch`
	cd - > /dev/null
	# echo "submodule_tracking_branch: $submodule_tracking_branch"
	if [ "" == "${submodule_tracking_branch}" ]; then
    	echo "Could not find a tracking branch for the submodule in directory ${path_to_submodule}"
    	return 1
	fi

	# Change into the directory and pull the tracking branch
	cd ${path_to_submodule}
	git pull --ff-only origin ${submodule_tracking_branch}
	cd - > /dev/null	
}

# Usage:
#
# git_submodule_update_to_branch 'path/to/submodule' 'preview'
#
# To update a single submodule to a certain remote branch.
function git_submodule_update_to_branch() {
	path_to_submodule="$1"
	target_branch="$2"

	if [ ! -d ${path_to_submodule} ]; then
		echo "Directory ${path_to_submodule} does not exist"
		return 1
	fi

	cd ${path_to_submodule}
	git pull --ff-only origin ${target_branch}
	ret=$?
	cd - > /dev/null

	return $ret
}
