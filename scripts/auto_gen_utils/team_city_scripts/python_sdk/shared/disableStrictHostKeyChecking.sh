# must disable StrictHostKeyChecking so that we don't get an interactive
# prompt later asking to confirm the host key
# If `set -e`, must disable "fail on non-zero exit code" using `set +e`
# because ssh returns 255
ssh -o StrictHostKeyChecking=no git@bitbucket.oci.oraclecorp.com -p 7999

# Old way of doing that:
# ls -la ~/.ssh
#
# cat ~/.ssh/config
# 
# printf "\n\nHost * \n  StrictHostKeyChecking no\n" >> ~/.ssh/config
# 
# cat ~/.ssh/config