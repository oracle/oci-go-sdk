# Copyright (c) 2016, 2020, Oracle and/or its affiliates.  All rights reserved.
# This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

# For all existing and new manual chages through self service.
# Deals with changing and adding onto existing manual changes and new manual changes.

from oci_cli import cli, dynamic_loader
import ast
import jinja2
import re
import inspect
import python_cli.self_service_manual_change_util as self_service_util
import subprocess
import os

dynamic_loader.load_all_services()

TEMPLATE_LOCATION = 'python_cli/templates'
file_loader = jinja2.FileSystemLoader(TEMPLATE_LOCATION)
env = jinja2.Environment(loader=file_loader)

LIST_REGEX = re.compile(r"\[.*?\]")

SERVICE = None


# command is a command invocation as a list [audit, event, list]
# old_params and new_params need to be a mapping to each other by index
# [param-old-1, param-old-2], [param_new_1, param_new_2]
# options are optional custom parameter options if you want to replace default values such as required=True
def rename_parameters(command, old_params, new_params, options=None):
    print("Renaming: " + ' '.join(command))
    parent_and_child = self_service_util.get_parent_child_command(command)
    parent_group = parent_and_child[0]

    extended_file = self_service_util.get_extended_file(parent_group)
    rename_param_dict = self_service_util.rename_parameter_dict(command, old_params, new_params, options)
    print("Generated param dictionary")

    existing_change = self_service_util.check_existing_manual_change(extended_file, rename_param_dict['command'])
    print("Checked for Existing Change")

    if existing_change:
        out_file = []
        command_found = False
        option_updated = False
        kwargs_updated = False
        try:
            with open(extended_file, 'r') as file:
                contents = file.readlines()
                for i, line in enumerate(contents):
                    if 'cli_util.copy_params_from_generated_command' in line and rename_param_dict['command'] in line:
                        command_found = True
                        excluded_params = LIST_REGEX.search(line)
                        if excluded_params:
                            excluded_params = excluded_params.group(0)
                            excluded_params = ast.literal_eval(excluded_params)
                            old_params_underscore = [param.replace("-", "_") for param in old_params]
                            for param in old_params_underscore:
                                excluded_params.append(param)
                        line = self_service_util.render_output({'service_cli': rename_param_dict['service_cli'],
                                                                'command': rename_param_dict['command'],
                                                                'params_to_exclude': excluded_params
                                                               },
                                                               "copy_params_from_generated.py.js2")
                    if command_found and "cli_util.option" in line and not option_updated:
                        option_updated = True
                        output = self_service_util.render_output(rename_param_dict, "rename_parameter_option.py.js2")
                        out_file.append(output)

                    out_file.append(line)

                    if option_updated and "ctx" in line and not kwargs_updated:
                        kwargs_updated = True
                        output = self_service_util.render_output(rename_param_dict, "rename_parameter_kwargs.py.js2")
                        out_file.append(output)

        except Exception as e:
            print(e)

        try:
            with open(extended_file, 'w') as file:
                file.writelines(out_file)
        except Exception as e:
            print(e)

    else:
        output = self_service_util.render_output(rename_param_dict, "rename_parameter.py.js2")
        output = output.replace(",help=", ", help=")
        # check for imports
        add_import_if_not_present(extended_file, "from oci_cli import json_skeleton_utils")
        append_to_file(extended_file, output)

        # json_skeleton_generation_handler = self_service_util.get_json_skeleton_generation_handler(command[0], child_group).rstrip()
    print("Successfuly renamed parameters!")
    return


def remove_parameters(command, params):
    print("Removing: " + ' '.join(command))
    parent_and_child = self_service_util.get_parent_child_command(command)
    parent_group = parent_and_child[0]

    extended_file = self_service_util.get_extended_file(parent_group)
    remove_param_dict = self_service_util.remove_parameter_dict(command, params)
    print("Generated param dictionary")

    existing_change = self_service_util.check_existing_manual_change(extended_file, remove_param_dict['command'])
    print("Checked for Existing Change")

    if existing_change:
        out_file = []
        try:
            with open(extended_file, 'r') as file:
                contents = file.readlines()
                for i, line in enumerate(contents):
                    if 'cli_util.copy_params_from_generated_command' in line and remove_param_dict['command'] in line:
                        excluded_params = LIST_REGEX.search(line)
                        if excluded_params:
                            excluded_params = excluded_params.group(0)
                            excluded_params = ast.literal_eval(excluded_params)
                            old_params_underscore = [param.replace("-", "_") for param in params]
                            for param in old_params_underscore:
                                excluded_params.append(param)
                        line = self_service_util.render_output({'service_cli': remove_param_dict['service_cli'],
                                                                'command': remove_param_dict['command'],
                                                                'params_to_exclude': excluded_params
                                                               },
                                                               "copy_params_from_generated.py.js2")

                    out_file.append(line)

        except Exception as e:
            print(e)

        try:
            with open(extended_file, 'w') as file:
                file.writelines(out_file)
        except Exception as e:
            print(e)

    else:
        output = self_service_util.render_output(remove_param_dict, "remove_parameter.py.js2")
        append_to_file(extended_file, output)

        # json_skeleton_generation_handler = self_service_util.get_json_skeleton_generation_handler(command[0], child_group).rstrip()
    print("Successfuly removed parameters!")
    return


# command is a command invocation as a list [audit, event, list]
# params is a list of complex parameters to flatten [param1, param2]
# params_flattened is a list of lists of the flattened params [[param1_flat, param1_flat], [param2_flat, param2_flat]]
# params_options is a dictionary with key being the flattened params and value being the parameter options {'param1_flat': [required=True]}
def flatten_parameters(command, params, params_flattened, params_options):
    print("Flattening: " + ' '.join(command))
    parent_and_child = self_service_util.get_parent_child_command(command)
    parent_group = parent_and_child[0]

    extended_file = self_service_util.get_extended_file(parent_group)
    rename_param_dict = self_service_util.flatten_parameter_dict(command, params, params_flattened, params_options)

    existing_change = self_service_util.check_existing_manual_change(extended_file, rename_param_dict['command'])

    if existing_change:
        out_file = []
        command_found = False
        option_updated = False
        kwargs_updated = False
        try:
            with open(extended_file, 'r') as file:
                contents = file.readlines()
                for i, line in enumerate(contents):
                    if 'cli_util.copy_params_from_generated_command' in line and rename_param_dict['command'] in line:
                        command_found = True
                        excluded_params = LIST_REGEX.search(line)
                        if excluded_params:
                            excluded_params = excluded_params.group(0)
                            excluded_params = ast.literal_eval(excluded_params)
                            params_underscore = [param.replace("-", "_") for param in params]
                            for param in params_underscore:
                                excluded_params.append(param)
                        line = self_service_util.render_output({'service_cli': rename_param_dict['service_cli'],
                                                                'command': rename_param_dict['command'],
                                                                'params_to_exclude': excluded_params
                                                               },
                                                               "copy_params_from_generated.py.js2")
                    if command_found and "cli_util.option" in line and not option_updated:
                        option_updated = True
                        output = self_service_util.render_output(rename_param_dict, "flatten_parameter_option.py.js2")
                        out_file.append(output)

                    out_file.append(line)

                    if option_updated and "ctx" in line and not kwargs_updated:
                        kwargs_updated = True
                        output = self_service_util.render_output(rename_param_dict, "flatten_parameter_kwargs.py.js2")
                        out_file.append(output)

        except Exception as e:
            print(e)

        try:
            with open(extended_file, 'w') as file:
                file.writelines(out_file)
        except Exception as e:
            print(e)

    else:
        output = self_service_util.render_output(rename_param_dict, "flatten_parameter.py.js2")
        append_to_file(extended_file, output)

        # json_skeleton_generation_handler = self_service_util.get_json_skeleton_generation_handler(command[0], child_group).rstrip()
    print("Successfuly flattened parameters!")
    return


# Renames all commands to the new name.
# Commands and new_names are a 1-to-1 mapping by index
# Commands are a list of lists and new_names is a list
# e.g. [[audit, list-long], [audit, create-long]], [list, create]
def rename_commands(commands, new_names):
    parent_and_child = self_service_util.get_parent_child_command(commands[0])
    parent_group = parent_and_child[0]
    child_group = parent_and_child[1]

    extended_file = self_service_util.get_extended_file(parent_group, child_group)
    rename_dict = self_service_util.rename_command_dict(commands, new_names)
    output = self_service_util.render_output(rename_dict, "rename_command.py.js2")

    # check for required imports
    add_import_if_not_present(extended_file, "from oci_cli import cli_util")
    append_to_file(extended_file, output)
    print("Successfully renamed commands!")
    return


# Removes all commands
# Commands are a list of lists of commands
# e.g. [[audit, list-long], [audit, create-long]]
def remove_commands(commands):
    parent_and_child = self_service_util.get_parent_child_command(commands[0])
    child_group = parent_and_child[1]

    extended_file = self_service_util.get_extended_file(child_group)
    remove_dict = self_service_util.remove_command_dict(commands)
    output = self_service_util.render_output(remove_dict, "remove_command.py.js2")

    append_to_file(extended_file, output)
    print("Successfully removed commands!")
    return


# Move command between groups
# e.g. [audit, event, list], [audit, config, list]
# oci audit event list -> oci audit config list
def move_command(command, new_group):
    print("Moving: " + ' '.join(command))
    old_parent_and_child = self_service_util.get_parent_child_command(command)
    old_child_group = old_parent_and_child[1]

    extended_file = self_service_util.get_extended_file(old_child_group)
    move_dict = self_service_util.move_command_dict(command, new_group)
    output = self_service_util.render_output(move_dict, "move_command.py.js2")

    append_to_file(extended_file, output)
    print("Successfully moved commands!")
    return


# Move commands under one group to another
# e.g. [audit, event], [audit, config]
# oci audit event list -> oci audit config list, oci audit event get -> oci audit config get
# [audit, event], [audit]
# oci audit event list -> oci audit list
def move_group(old_group, new_group):
    print("Moving: " + ' '.join(old_group))
    old_parent_and_child = self_service_util.get_parent_child_command(old_group)
    old_child_group = old_parent_and_child[1]

    extended_file = self_service_util.get_extended_file(old_child_group)
    move_dict = self_service_util.move_group_dict(old_group, new_group)
    output = self_service_util.render_output(move_dict, "move_group.py.js2")

    append_to_file(extended_file, output)
    print("Successfully moved group commands!")
    return


# Renames the root group in the service pom.xml file with new_name
def rename_root_group(service, new_name):
    print("Renaming service: " + service)
    service_cli_file = str(inspect.getfile(cli.commands[service].callback))
    pom_file = service_cli_file.rsplit('/', 4)[0] + "/pom.xml"
    with open(pom_file, 'r') as infile:
        file_data = infile.read()
        file_data = file_data.replace(_pom_root_group(service), _pom_root_group(new_name))
    with open(pom_file, 'w') as outfile:
        outfile.write(file_data)

    service_dir = service_cli_file.rsplit('/', 4)[0]
    run_make_gen(service_dir)
    run_make_docs_dir(service_dir)
    print("Successfully renamed root group!")


def _pom_root_group(service):
    return "<generate-rootCliGroupOverride>{}</generate-rootCliGroupOverride>".format(service)


def run_make_gen(service_dir):
    # Switches to service directory for 'make gen' then switches back to original working directory
    owd = os.getcwd()

    os.chdir(service_dir)
    os.environ["OCI_CLI_SKIP_DOC_LINK_VALIDATION"] = "1"
    make_gen = subprocess.run(['make', 'gen'],
                              stdout=subprocess.PIPE,
                              stderr=subprocess.DEVNULL,
                              universal_newlines=True)
    if make_gen.returncode != 0:
        raise Exception("Error while running 'make gen'!")

    os.chdir(owd)


def run_make_docs(service):
    service_cli_file = str(inspect.getfile(cli.commands[service].callback))
    run_make_docs_dir(service_cli_file.rsplit('/', 4)[0])


def run_make_docs_dir(service_dir):
    print(service_dir)
    # Switches to service directory for 'make gen' then switches back to original working directory
    owd = os.getcwd()

    os.chdir(service_dir)
    os.environ["OCI_CLI_SKIP_DOC_LINK_VALIDATION"] = "1"
    make_docs = subprocess.run(['make', 'docs'],
                               stdout=subprocess.PIPE,
                               stderr=subprocess.DEVNULL,
                               universal_newlines=True)
    if make_docs.returncode != 0:
        raise Exception("Error while running 'make docs'!")

    os.chdir(owd)


def run_command(dir, command):
    # Switches to directory for command then switches back to original working directory
    owd = os.getcwd()

    os.chdir(dir)
    process = subprocess.run(command.split(),
                             stdout=subprocess.PIPE,
                             universal_newlines=True)
    if process.returncode != 0:
        raise Exception("Error while running " + command)

    os.chdir(owd)


def use_local_python_sdk(dir):
    setup_file = os.path.join(dir, 'setup.py')
    modify_python_sdk_version(setup_file, "preview.1',\n")
    requirements_file = os.path.join(dir, 'requirements.txt')
    modify_python_sdk_version(requirements_file, "preview.1\n")


def modify_python_sdk_version(file, replacement):
    outfile = []
    try:
        with open(file, 'r') as f:
            content = f.readlines()
            for line in content:
                if 'oci==' in line:
                    local_sdk = line.split('preview.1')[0]
                    outfile.append(local_sdk + replacement)
                else:
                    outfile.append(line)
    except Exception as e:
        print(e)

    try:
        with open(file, 'w') as f:
            f.writelines(outfile)
    except Exception as e:
        print(e)


def append_to_file(file, data):
    try:
        with open(file, "a") as file:
            file.write(data)
    except Exception as e:
        print(e)


def add_import_if_not_present(file, input_line):
    # check if file contains input line
    with open(file, 'r') as f:
        lines_list = f.readlines()
        data = "".join(lines_list)

    regex = get_regex(input_line)
    matches = re.search(regex, data, re.MULTILINE)
    if not matches:
        i = 0
        line = lines_list[0]
        # skip comments
        while line.startswith("#"):
            i = i + 1
            line = lines_list[i]
        try:
            insert_line = input_line + "  # noqa: F401\n"
            lines_list.insert(i, insert_line)
            with open(file, "w") as file:
                lines_list = "".join(lines_list)
                file.write(lines_list)
        except Exception as e:
            print(e)


def get_regex(input_line):
    noqa_statement = '\\s*# noqa:\\s*F401\\s*'
    return '^(\\s*' + input_line + '\\s*|\\s*' + input_line + noqa_statement + ')$'
