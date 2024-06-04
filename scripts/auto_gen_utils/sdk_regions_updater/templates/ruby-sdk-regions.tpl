# Copyright (c) 2016, 2024, Oracle and/or its affiliates.  All rights reserved.
# This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

module OCI
  # rubocop:disable Style/MutableConstant

  # Module defining available regions
  module Regions
    REGION_ENUM = [
%REGIONS_DEFINITIONS%
    ]

    REGION_SHORT_NAMES_TO_LONG_NAMES = {
%SHORT_NAME_REGIONS%
    }

    # --- Start of region realm mapping ---
    REGION_REALM_MAPPING = {
%REGION_REALM_MAPPINGS%
    }
    # ---  end of region realm mapping  ---

    # --- Start of realm domain mapping ---
    REALM_DOMAIN_MAPPING = {
%REALMS_DEFINITIONS%
    }
    # ---  end of realm domain mapping  ---
  end
  # rubocop:enable Style/MutableConstant
end
