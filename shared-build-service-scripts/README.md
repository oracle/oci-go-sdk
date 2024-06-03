# Shared Build Service Scripts

These scripts are meant to be used as a Git submodule inside another repository.

```
git submodule add -b main ssh://git@bitbucket.oci.oraclecorp.com:7999/sdk/shared-build-service-scripts.git
```

Make sure that in your ocibuild.conf file, you don't set cloneSubmodules: false (the [default is true](https://confluence.oci.oraclecorp.com/display/BLD/Build+Service+ocibuild.conf+Reference+Guide#BuildServiceocibuild.confReferenceGuide-GitStepProperties)).


Also, the Git submodule does not automatically update to the latest version. If we make changes to the shared-build-service-scripts repo or add new scripts, you have to update the Git submodule in your repository and commit the change:

```
git submodule update --remote
git commit -a -m "Submodule update."
git push origin HEAD
```
