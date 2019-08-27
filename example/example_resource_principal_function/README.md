# Resource Principal example

_Resource principals_ offer credentials tied to specific OCI resources. Code running in the context of those
resources may be granted the rights to act "as the resource".

The example code in this directory can be assembled into an OCI Functions container. If that function is given
the permissions to read (or use) resources in a tenancy, it may do so. As with all rights grants, this involves two
steps:

1. Construct a _dynamic group_ whose membership includes the function, for example with:

   ```
   ### Example dynamic-group rule for dynamic group fnfunc-example-compartment
   
   ALL {resource.type = 'fnfunc', resource.compartment.name = 'example-compartment'}
   ```
    
2. Add the rights to that dynamic group with a suitable policy, such as:

   ```
   ### Policy statement for policy example-compartment/fnfunc-grant-reads
   
   allow dynamic-group fnfunc-example-compartment to read all-resources in compartment example-compartment
   ```

The function may then be deployed and invoked as usual.

References:

- [General overview of OCI Functions](https://docs.cloud.oracle.com/iaas/Content/Functions/Concepts/functionsoverview.htm)
- [Using Resource Principals from within OCI Functions](https://docs.cloud.oracle.com/iaas/Content/Functions/Tasks/functionsaccessingociresources.htm)
