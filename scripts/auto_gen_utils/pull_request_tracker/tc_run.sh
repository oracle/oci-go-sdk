set -e
set -x

ISSUE_FILTER_ARG=""
if [ ! -z "$ISSUES" ]; then
  ISSUE_FILTER_ARG="--issues "$ISSUES
fi

LIMIT_ARG=""
if [ ! -z "$LIMIT" ]; then
  LIMIT_ARG="--limit "$LIMIT
fi

UDX_STATUS_ARG=""
if [ ! -z "$UDX_STATUS" ]; then
  UDX_STATUS_ARG="--udx-status "$UDX_STATUS
fi

REPORT_NAME_ARG=""
if [ ! -z "$REPORT_NAME" ]; then
  REPORT_NAME_ARG="--report-name "$REPORT_NAME
fi

UPLOAD_TO_OBJECT_STORAGE_ARG=""
if [ ! -z "$UPLOAD_TO_OBJECT_STORAGE" ]; then
  UPLOAD_TO_OBJECT_STORAGE_ARG="--upload-to-object-storage"
fi

SDK_CLI_GA_DATE_ARG=""
if [ ! -z "$SDK_CLI_GA_DATE" ]; then
  SDK_CLI_GA_DATE_ARG="--sdk-cli-ga-date "$SDK_CLI_GA_DATE
fi

NEXT_RELEASE_N_ARG=""
if [ ! -z "$NEXT_RELEASE_N" ]; then
  NEXT_RELEASE_N_ARG="--next-release-n "$NEXT_RELEASE_N
fi

eval "$(pyenv init -)"
eval "$(pyenv init --path)"
pyenv shell cli-3

MOUNT_DIR=auto-gen-utils

cd $MOUNT_DIR/pull_request_tracker

pip ${PIP_TIMEOUT_PARAMETER} install -r ../requirements.txt

python dexreq_pr_tracker.py $ISSUE_FILTER_ARG $LIMIT_ARG $UDX_STATUS_ARG $REPORT_NAME_ARG $UPLOAD_TO_OBJECT_STORAGE_ARG $SDK_CLI_GA_DATE_ARG $NEXT_RELEASE_N_ARG
