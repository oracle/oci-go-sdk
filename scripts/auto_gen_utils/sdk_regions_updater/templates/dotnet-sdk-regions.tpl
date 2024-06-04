/*
 * Copyright (c) 2020, 2024, Oracle and/or its affiliates. All rights reserved.
 * This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
 */

namespace Oci.Common
{
    /// <summary>
    /// Class containing all of the known Regions that can be contacted.
    /// <br/>
    /// Note, not all services may be available in all regions.
    /// </summary>
    public partial class Region
    {
%REGIONS_DEFINITIONS%
    }
}