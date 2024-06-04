################################################################################
# Script to clean up featureId and enabledGroups files by consolidating them
#
# Consolidates all feature IDs referenced in files under codegenConfig/featuredIds
# into a single file: codegenConfig/featureIds/{CONSOLIDATED_FEATURE_IDS_FILENAME}.yaml
#
# Consolidates all enabled groups referenced in files under codegenConfig/enabledGroups
# into a single file: codegenConfig/enabledGroups/{CONSOLIDATED_ENABLED_GROUPS_FILENAME}.txt
#
# Example invocation:
# python consolidate_feature_id_enabled_groups_files.py --codegen-config-dir ~/dev/SDK/python-cli/codegenConfig/
#
################################################################################

import argparse
import os
import sys
import yaml
import util

ENABLED_GROUPS_SUBDIRECTORY = "enabledGroups"
FEATURE_IDS_SUBDIRECTORY = "featureIds"

CONSOLIDATED_ENABLED_GROUPS_FILENAME = "enabledGroups-combined"
CONSOLIDATED_FEATURE_IDS_FILENAME = "featureIds-combined"

COMMENT_CHARACTER = "#"
FEATURE_ID_FILE_LIST_HEADER = "whitelisted"

FILES_TO_SKIP = [
    '.empty',
    CONSOLIDATED_ENABLED_GROUPS_FILENAME,
    CONSOLIDATED_FEATURE_IDS_FILENAME
]


def consolidate_directory(dir, consolidated_filename, get_all_ids_from_file_func, add_ids_to_file_func):
    ids = []
    for filename in os.listdir(dir):
        filename_without_extension = filename
        try:
            filename_without_extension = os.path.splitext(os.path.basename(filename))[0]
        except Exception:
            pass

        if filename in FILES_TO_SKIP or filename_without_extension in FILES_TO_SKIP:
            continue

        file_path = os.path.join(dir, filename)
        ids.extend(get_all_ids_from_file_func(file_path))

        os.remove(file_path)

    consolidated_file_path = os.path.join(dir, consolidated_filename)
    add_ids_to_file_func(dir, ids, consolidated_filename)

    print('Created file: {}. With IDs: {}'.format(consolidated_file_path, ', '.join(ids)))


def get_all_ids_from_feature_id_file(file_path):
    feature_ids = []
    with open(file_path, 'r') as f:
        doc = yaml.load(f)
        if FEATURE_ID_FILE_LIST_HEADER in doc:
            feature_ids.extend(doc[FEATURE_ID_FILE_LIST_HEADER])

    return feature_ids


def get_all_ids_from_enabled_groups_file(file_path):
    enabled_groups = []

    with open(file_path, 'r') as f:
        lines = f.readlines()

        for line in lines:
            if not line.strip().startswith(COMMENT_CHARACTER):
                enabled_groups.append(line.strip())

    return enabled_groups


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Consolidate codegenConfig directory into a single file for feature Ids and a single file for enabled groups.')
    parser.add_argument('--codegen-config-dir',
                        required=True,
                        help='The codegenConfig directory to consolidate. Two directories "enabledGroups" and "featureIds" are expected as subdirectories')

    args = parser.parse_args()

    codegen_config_dir = args.codegen_config_dir
    codegen_config_dir = os.path.expandvars(os.path.expanduser(codegen_config_dir))
    if not os.path.exists(codegen_config_dir):
        sys.exit('Path did not exist: {}'.format(codegen_config_dir))
    
    if not os.path.isdir(codegen_config_dir):
        sys.exit('Path {} is not a directory'.format(codegen_config_dir))
    
    enabled_groups_directory = os.path.join(codegen_config_dir, ENABLED_GROUPS_SUBDIRECTORY)
    feature_ids_directory = os.path.join(codegen_config_dir, FEATURE_IDS_SUBDIRECTORY)

    consolidate_directory(
        enabled_groups_directory,
        CONSOLIDATED_ENABLED_GROUPS_FILENAME,
        get_all_ids_from_enabled_groups_file,
        util.update_pre_processor_file
    )

    consolidate_directory(
        feature_ids_directory,
        CONSOLIDATED_FEATURE_IDS_FILENAME,
        get_all_ids_from_feature_id_file,
        util.update_feature_id_file
    )
