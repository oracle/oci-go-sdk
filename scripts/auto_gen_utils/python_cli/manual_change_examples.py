#!/usr/bin/env python3
# Copyright (c) 2016, 2020, Oracle and/or its affiliates.  All rights reserved.
# This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
import python_cli.self_service_manual_change_util as self_service_util  # noqa: F401
import python_cli.self_service_manual_change as self_service  # noqa: F401

# Use this python script as an example to generate manual changes

# Rename Command
# @param commands: A list of list of commands. Takes in multiple commands. [[audit, event, list], [audit, event, create]]
# @param new_names: A list of names to replace the command with. [new_list, new_create]
# self_service.rename_commands([['usage', 'subscription-info', 'get']], ['getter'])


# Rename Parameter
# @param command: A command in list form, [audit, event, list]. Takes in only 1 command
# @param old_params: A list of parameters to rename, [param1, param2]
# @param new_params: A list of parameters to replace old_params. Index mapped to old_params. [new_param1, new_param2]
# @param options: An optional field, defaulted to None, to replace the options for the parameter. A dictionary for parameter options. {'param1': [required=True], 'param1': [required=True]}
# self_service.rename_parameters(['usage', 'usage-record', 'list'], ['granularity', 'tenancy-id'], ['grangran', 'tenten'])


# Remove command
# @param commands: A list of list of commands. Takes in multiple commands.
# self_service.remove_commands([['usage', 'usage-record', 'list'], ['usage', 'subscription-info', 'get']])


# Move group
# @param old_group: A group in list form, [os, bucket]. Takes in only 1 group
# @param new_group: The group to move all commands under. A group in list form, [os, new_bucket]. Takes in only 1 group
# self_service.move_group(['waas', 'address-list'], ['waas', 'access-rule'])


# Move command
# @param old_command: A command in list form, [audit, event, list]. Takes in only 1 command
# @param new_command: The group to move the command under. A group in list form. [audit, config] Takes in only 1 group
# self_service.move_command(['waas', 'address-list', 'create'], ['waas', 'access-rule'])


# Flatten command
# @param command: A command in list form, [audit, event, list]. Takes in only 1 command
# @param params: A list of complex parameters to flatten, [param1, param2]
# @param params_flattened: A list of lists of flattened parameters to replace old_params. Index mapped to old_params. [[param1_flat, param1_flat], [param2_flat, param2_flat]]
# @param params_options: A dictionary with key being the flattened params and value being the parameter options {'param1_flat': [required=True]}
# self_service.flatten_parameters(['usage', 'usage-record', 'list'], ['tenancy-id'], [['tenancy-renamed', 'gran-renamed']], {'tenancy-renamed': ['required=True', 'help="Text"'], 'gran-renamed': ['help="Text"']})


# Rename root group
# @param service: The root group to rename, takes in only 1
# @param new_name: The new name for the root group.
# self_service.rename_root_group("budgets", "budget")
