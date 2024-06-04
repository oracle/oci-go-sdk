# Copyright (c) 2016, 2020, Oracle and/or its affiliates.  All rights reserved.
# This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
from oci_cli import cli, dynamic_loader
import sys
import inspect
import glob
import jinja2

dynamic_loader.load_all_services()

file_loader = jinja2.FileSystemLoader('python_cli/templates')
env = jinja2.Environment(loader=file_loader)


# Generates dict object required for renaming parameter template
# old_params and new_params need to be a mapping to each other by index
# [param-old-1, param-old-2], [param_new_1, param_new_2]
# options are optional custom parameter options if you want to replace default values such as required=True. It is a list of lists
def rename_parameter_dict(command, old_params, new_params, options=None):
    """
    @param command: e.g. [audit, list, event]
    @param params_old: [param-old, param-old2]
    @param params_new: [param-new, param-new2]
    @param options: [[required=True], [required=False]]
    """

    parent_and_child = get_parent_child_command(command)
    parent_group = parent_and_child[0]
    child_group = parent_and_child[1]

    service_cli = get_service_cli_from_command(parent_group)

    params_old_underscore = [param.replace("-", "_") for param in old_params]
    params_new_underscore = [param.replace("-", "_") for param in new_params]
    # Zips old and new parameters for a 1-to-1 mapping
    rename_parameters = [list(x) for x in zip(params_old_underscore, params_new_underscore)]

    param_dict = parameter_list(child_group, params_old_underscore)
    param_options = get_parameter_options(command[0], parent_group, child_group, old_params, options)
    param_dict = _combine_parameter_dicts(param_dict, param_options)
    param_dict = zip(param_dict, new_params)

    command_name = get_command_from_extended(get_extended_file(parent_group), child_group)

    json_skeleton_generation_handler = get_json_skeleton_generation_handler(command[0], parent_group, child_group).rstrip()
    rename_dict = {'service_cli': service_cli,
                   'parent_command': parent_group.callback.__name__,
                   'command': command_name,
                   'command_name': child_group.name,
                   'params_to_exclude': params_old_underscore,
                   'rename_parameters': rename_parameters,
                   'param_dict': param_dict,
                   'json_skeleton_generation_handler': json_skeleton_generation_handler}
    return rename_dict


# Generates dict object required for removing parameter template
def remove_parameter_dict(command, params):
    """
    @param command: e.g. [audit, list, event]
    @param params: [param1, param2]
    """

    parent_and_child = get_parent_child_command(command)
    parent_group = parent_and_child[0]
    child_group = parent_and_child[1]

    service_cli = get_service_cli_from_command(parent_group)

    params_underscore = [param.replace("-", "_") for param in params]

    command_name = get_command_from_extended(get_extended_file(parent_group), child_group)

    json_skeleton_generation_handler = get_json_skeleton_generation_handler(command[0], parent_group, child_group).rstrip()
    rename_dict = {'service_cli': service_cli,
                   'parent_command': parent_group.callback.__name__,
                   'command': command_name,
                   'params_to_exclude': params_underscore,
                   'json_skeleton_generation_handler': json_skeleton_generation_handler}
    return rename_dict


# Generates dict object required for flattening parameter template.
# params and params_flattened need to be a mapping to each other by index.
# params is a list of params to be flattened.
# params_flattened is a list of lists with each index corresponding to the original param.
# params_options is a dictionary with key being the name of the flattened param and value being the param options
def flatten_parameter_dict(command, params, params_flattened, params_options):
    """
    @param command: [audit, list, event]
    @param params: e.g. [param-1, param-2]
    @param params_flattened: [[param1-flattened1, param1-flattened2]]
    @param params_options: {param-flattened1: [required=True, help="Text"], param-flattened2: [help="Text"]}
    """

    parent_and_child = get_parent_child_command(command)
    parent_group = parent_and_child[0]
    child_group = parent_and_child[1]

    service_cli = get_service_cli_from_command(parent_group)

    params_underscore = [param.replace("-", "_") for param in params]
    params_flattened_underscore = [[param.replace("-", "_") for param in params] for params in params_flattened]
    flatten_parameters = [list(x) for x in zip(params_underscore, params_flattened_underscore)]

    json_skeleton_generation_handler = get_json_skeleton_generation_handler(command[0], child_group).rstrip()

    command_name = get_command_from_extended(get_extended_file(parent_group), child_group)

    flatten_dict = {'service_cli': service_cli,
                    'parent_command': parent_group.callback.__name__,
                    'command': command_name,
                    'command_name': child_group.name,
                    'params_to_exclude': params_underscore,
                    'json_skeleton_generation_handler': json_skeleton_generation_handler,
                    'flatten_parameters': flatten_parameters,
                    'params_options': params_options}
    return flatten_dict


# Generates dict object required for renaming command template
def rename_command_dict(commands, new_names):
    renamed_commands = []
    for i, command in enumerate(commands):
        print("Renaming command: {}".format(' '.join(command)))
        parent_and_child = get_parent_child_command(command)
        parent_group = parent_and_child[0]
        child_group = parent_and_child[1]

        service_cli = get_service_cli_from_command(parent_group)
        child_service_cli = get_service_cli_from_command(child_group)

        new_command = command.copy()
        new_command[-1] = new_names[i]
        rename_dict = {'service_cli': service_cli,
                       'child_service_cli': child_service_cli,
                       'command_parent_group': parent_group.callback.__name__,
                       'command_group': child_group.callback.__name__,
                       'new_name': new_names[i],
                       'old_command': " ".join(command),
                       'new_command': " ".join(new_command)}
        renamed_commands.append(rename_dict)
    return {'renamed_commands': renamed_commands}


# Generates dict object required for removing command template
def remove_command_dict(commands):
    removed_commands = []
    for command in commands:
        print("Removing command: {}".format(' '.join(command)))

        parent_and_child = get_parent_child_command(command)
        parent_group = parent_and_child[0]
        child_group = parent_and_child[1]

        service_cli = get_service_cli_from_command(child_group)

        removed_command = {'service_cli': service_cli,
                           'parent_group': parent_group.callback.__name__,
                           'command_group': child_group.callback.__name__,
                           'full_command': " ".join(command[:-1]),
                           'command': command[-1]}
        removed_commands.append(removed_command)
    return {'removed_commands': removed_commands}


# Generates dict object required for moving command template
def move_command_dict(command, new_group):

    old_parent_and_child = get_parent_child_command(command)
    old_parent_group = old_parent_and_child[0]
    old_child_group = old_parent_and_child[1]

    service_cli = get_service_cli_from_command(old_child_group)

    new_parent_group = get_command_group(new_group)

    move_dict = {'old_command': " ".join(command),
                 'new_command': " ".join(new_group),
                 'service_cli': service_cli,
                 'old_parent_group': old_parent_group.callback.__name__,
                 'command_group': old_child_group.callback.__name__,
                 'new_parent_group': new_parent_group.callback.__name__}

    return move_dict


def move_group_dict(old_group, new_group):

    old_parent_and_child = get_parent_child_command(old_group)
    old_parent_group = old_parent_and_child[0]
    old_child_group = old_parent_and_child[1]

    service_cli = get_service_cli_from_command(old_child_group)

    new_parent_group = get_command_group(new_group)
    new_service_cli = get_service_cli_from_command(new_parent_group)

    moved_commands = []
    for command in old_child_group.commands.values():
        moved_commands.append(command.callback.__name__)

    group_dict = {'service_cli': service_cli,
                  'new_service_cli':new_service_cli,
                  'old_group': " ".join(old_group),
                  'new_group': " ".join(new_group),
                  'old_parent_group': old_parent_group.callback.__name__,
                  'old_command_group': old_child_group.callback.__name__,
                  'new_parent_group': new_parent_group.callback.__name__,
                  'moved_commands': moved_commands}
    return group_dict


# Returns the "@json_skeleton_utils.json_skeleton_generation_handler(input_params_to_complex_types={}, output_type={})"
# that is used for a command.
def get_json_skeleton_generation_handler(service, command, child_group):
    service_cli = get_generated_file_from_command(command)
    command_name = child_group.callback.__name__
    command_to_find = command_name + ".command_name"
    found_command = False
    json_skeleton_generation_handler = ""
    try:
        with open(service_cli, "r") as service_file:
            contents = service_file.readlines()
            for line in contents:
                if not found_command:
                    if line.__contains__(command_to_find):
                        found_command = True
                else:
                    if line.__contains__("@json_skeleton_utils.json_skeleton_generation_handler"):
                        json_skeleton_generation_handler = line
                        break
    except IOError:
        raise Exception("Encountered error while parsing service file.")
    finally:
        service_file.close()
    return json_skeleton_generation_handler


# Returns the optional flags of a command mapped to their option
# Searches through the generated file for the options in params
# @cli_util.option('--tenancy-id', required=True, help=u"""The OCID of the tenancy for which usage is being fetched.""")
# Will return only '{"tenancy_id": "required=True, "}' from the above string
def get_parameter_options(service, command, child_group, params, options=None):
    service_cli = get_generated_file_from_command(command)
    command_name = child_group.callback.__name__
    command_name = command_name.replace("_extended", "")
    command_to_find = command_name + ".command_name"
    param_dict = {}

    found_command = False

    try:
        with open(service_cli, "r") as service_file:
            contents = service_file.readlines()
            for line in contents:
                if not found_command:
                    if line.__contains__(command_to_find):
                        found_command = True
                else:
                    if len(param_dict) >= len(params):
                        break
                    for param in params:
                        if line.__contains__(param + "',"):
                            line_split = line.split("help=")
                            line_split = line_split[0].split(",")
                            param_options = ""
                            for i, option in enumerate(line_split):
                                if "=" in option:
                                    param_options = ",".join(line_split[i:])
                                    break
                            param_dict.update({param.replace("-", "_"): param_options})
    except Exception as e:
        print(e)

    if options is not None:
        for param, option in options.items():
            param_dict.update({param.replace("-", "_"): option})
    return param_dict


def _combine_parameter_dicts(param_dict, param_options):
    for dict in param_dict:
        dict.update({dict['name']: param_options[dict['name']]})
    return param_dict


# Command is a list of a valid CLI command and returns the group of the parent and child
# e.g. [os, object, get] returns the groups for object and get
def get_parent_child_command(command):
    commands = cli.commands
    parent_group = commands
    for group in command[:-1]:
        if group in commands:
            parent_group = commands[group]
            commands = parent_group.commands
        else:
            raise Exception("Invalid command entered. {}".format(str(command.name)))
    return [parent_group, parent_group.commands[command[-1]]]


def get_command_group(command):
    commands = cli.commands
    command_group = commands
    for group in command:
        if group in commands:
            command_group = commands[group]
            commands = command_group.commands
        else:
            raise Exception("Invalid command entered. {}".format(str(command.name)))
    return command_group


# Gets the extended file for the service.
# If it does not exist, create one.
def get_extended_file(command, child_group=None):
    service_cli_file = get_generated_file_from_command(command)
    service_dir = service_cli_file.rsplit('/', 2)[0]

    extended_file_path = check_extended_file_exists(service_dir)

    service_cli = get_service_cli_from_command(command)
    module_path = ".".join(service_cli_file.rsplit('/', 6)[1:-1])

    # This step is if the parent and child commands have different source generated files.
    # oci usage list, list can be part of usage_list_cli, but usage can be in usage_cli
    child_service_cli = None
    child_module_path = None
    if child_group:
        child_service_cli = get_generated_file_from_command(child_group)
        if service_cli_file == child_service_cli:
            child_group = None
        else:
            child_module_path = ".".join(child_service_cli.rsplit('/', 6)[1:-1])
            child_service_cli = child_service_cli.rsplit('/', 1)[-1][:-3]

    if not extended_file_path:
        output = render_output({'service_cli': service_cli,
                                'module_path': module_path,
                                'child_group': child_group,
                                'child_service_cli': child_service_cli,
                                'child_module_path': child_module_path},
                               "new_extended_file.py.js2")
        extended_file_path = service_dir + "/" + service_cli + "_extended.py"
        try:
            with open(extended_file_path, "w") as extended_file:
                extended_file.write(output)
        except Exception as e:
            print(e)
    return extended_file_path


def check_extended_file_exists(service_dir):
    extended_file = glob.glob(service_dir + "/" + "*extended*")
    if len(extended_file) > 0:
        return extended_file[0]
    else:
        return False


# Returns the service cli for the particular service
# e.g audit returns audit_cli
def get_service_cli(service):
    service_cli = str(inspect.getfile(cli.commands[service].callback))
    service_cli = service_cli.rsplit('/', 1)[-1][:-3]
    return service_cli


# Returns the generated service cli for a command
# oci audit -> audit_cli
def get_service_cli_from_command(command):
    service_cli = get_generated_file_from_command(command)
    service_cli = service_cli.rsplit('/', 1)[-1][:-3]
    return service_cli


# Returns the parameter list as required to rename a parameter
def parameter_list(command, params):
    param_dict = []
    for p in params:
        param = get_param(command, p)
        param = vars(param)
        param_dict.append(param)
    return param_dict


# Gets the parameter object from the command
def get_param(command, param):
    for p in command.params:
        if p.name == param:
            return p

    raise Exception("Invalid param, does not exist. {} {}".format(str(command.name), param))


def render_output(dict, template):
    template = env.get_template(template)
    return template.render(dict)


def check_existing_manual_change(extended_file, command):
    existing_change = False
    try:
        with open(extended_file, 'r') as file:
            contents = file.readlines()
            for i, line in enumerate(contents):
                if 'cli_util.copy_params_from_generated_command' in line and command in line:
                    existing_change = True
                    break
    except Exception as e:
        print(e)

    return existing_change


def get_command_from_extended(extended_file, command):
    command_name = command.callback.__name__
    try:
        with open(extended_file, 'r') as file:
            name = ""
            contents = file.readlines()
            for line in contents:
                if "copy_params_from_generated_command" in line:
                    name = line
                if command_name in line and "def" in line:
                    command_name = name
                    command_name = command_name.split('.')[2]
                    command_name = command_name.split(',')[0]
                    break
    except Exception as e:
        print(e)

    return command_name


def get_generated_file_from_command(command):
    return str(sys.modules[command.callback.__module__].__file__)
